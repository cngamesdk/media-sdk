package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppShareCancelByAccount(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppShareCancelReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.AppId = 2199123264636
	req.ShareType = 1 // 按账号取消共享
	req.CancelShareAdvertiserIds = []int64{900164519}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppShareCancel(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result=%v\n", resp.Result)
}

func TestAppShareCancelByCorp(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppShareCancelReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.AppId = 2199123264636
	req.ShareType = 2 // 按主体取消共享
	req.CancelShareCorpIds = []int64{15280}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppShareCancel(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result=%v\n", resp.Result)
}

func TestAppShareCancelValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppShareCancelReq{}
	req.AdvertiserId = 900015364
	req.AppId = 2199123264636
	req.ShareType = 1
	req.CancelShareAdvertiserIds = []int64{900164519}
	_, err := adapter.AppShareCancel(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppShareCancelReq{}
	req2.AccessToken = "your_access_token"
	req2.AppId = 2199123264636
	req2.ShareType = 1
	req2.CancelShareAdvertiserIds = []int64{900164519}
	_, err2 := adapter.AppShareCancel(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 app_id，预期返回校验错误
	req3 := &kuaishouModel.AppShareCancelReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015364
	req3.ShareType = 1
	req3.CancelShareAdvertiserIds = []int64{900164519}
	_, err3 := adapter.AppShareCancel(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty app_id")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// share_type 非法值，预期返回校验错误
	req4 := &kuaishouModel.AppShareCancelReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 900015364
	req4.AppId = 2199123264636
	req4.ShareType = 3 // 非法值
	_, err4 := adapter.AppShareCancel(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for invalid share_type")
	}
	fmt.Printf("got expected error: %v\n", err4)

	// share_type=1 但未填 cancel_share_advertiser_ids，预期返回校验错误
	req5 := &kuaishouModel.AppShareCancelReq{}
	req5.AccessToken = "your_access_token"
	req5.AdvertiserId = 900015364
	req5.AppId = 2199123264636
	req5.ShareType = 1
	_, err5 := adapter.AppShareCancel(ctx, req5)
	if err5 == nil {
		t.Fatal("expected validation error for missing cancel_share_advertiser_ids")
	}
	fmt.Printf("got expected error: %v\n", err5)

	// cancel_share_advertiser_ids 超过200个，预期返回校验错误
	req6 := &kuaishouModel.AppShareCancelReq{}
	req6.AccessToken = "your_access_token"
	req6.AdvertiserId = 900015364
	req6.AppId = 2199123264636
	req6.ShareType = 1
	ids := make([]int64, 201)
	for i := range ids {
		ids[i] = int64(i + 1)
	}
	req6.CancelShareAdvertiserIds = ids
	_, err6 := adapter.AppShareCancel(ctx, req6)
	if err6 == nil {
		t.Fatal("expected validation error for cancel_share_advertiser_ids exceeding 200 items")
	}
	fmt.Printf("got expected error: %v\n", err6)

	// share_type=2 但未填 cancel_share_corp_ids，预期返回校验错误
	req7 := &kuaishouModel.AppShareCancelReq{}
	req7.AccessToken = "your_access_token"
	req7.AdvertiserId = 900015364
	req7.AppId = 2199123264636
	req7.ShareType = 2
	_, err7 := adapter.AppShareCancel(ctx, req7)
	if err7 == nil {
		t.Fatal("expected validation error for missing cancel_share_corp_ids")
	}
	fmt.Printf("got expected error: %v\n", err7)
}
