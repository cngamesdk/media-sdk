package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppShareAddByAccount(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppShareAddReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12078245
	req.AppId = 2199123262731
	req.ShareType = 1 // 按账号共享
	req.ShareAdvertiserIds = []int64{11, 22, 33}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppShareAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result=%v\n", resp.Result)
}

func TestAppShareAddByCorp(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppShareAddReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12078245
	req.AppId = 2199123262731
	req.ShareType = 2 // 按主体共享（异步操作）
	req.ShareCorpIds = []int64{15280}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppShareAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result=%v\n", resp.Result)
}

func TestAppShareAddValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppShareAddReq{}
	req.AdvertiserId = 12078245
	req.AppId = 2199123262731
	req.ShareType = 1
	req.ShareAdvertiserIds = []int64{11}
	_, err := adapter.AppShareAdd(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppShareAddReq{}
	req2.AccessToken = "your_access_token"
	req2.AppId = 2199123262731
	req2.ShareType = 1
	req2.ShareAdvertiserIds = []int64{11}
	_, err2 := adapter.AppShareAdd(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 app_id，预期返回校验错误
	req3 := &kuaishouModel.AppShareAddReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 12078245
	req3.ShareType = 1
	req3.ShareAdvertiserIds = []int64{11}
	_, err3 := adapter.AppShareAdd(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty app_id")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// share_type 非法值，预期返回校验错误
	req4 := &kuaishouModel.AppShareAddReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 12078245
	req4.AppId = 2199123262731
	req4.ShareType = 3 // 非法值
	_, err4 := adapter.AppShareAdd(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for invalid share_type")
	}
	fmt.Printf("got expected error: %v\n", err4)

	// share_type=1 但未填 share_advertiser_ids，预期返回校验错误
	req5 := &kuaishouModel.AppShareAddReq{}
	req5.AccessToken = "your_access_token"
	req5.AdvertiserId = 12078245
	req5.AppId = 2199123262731
	req5.ShareType = 1
	_, err5 := adapter.AppShareAdd(ctx, req5)
	if err5 == nil {
		t.Fatal("expected validation error for missing share_advertiser_ids")
	}
	fmt.Printf("got expected error: %v\n", err5)

	// share_advertiser_ids 超过200个，预期返回校验错误
	req6 := &kuaishouModel.AppShareAddReq{}
	req6.AccessToken = "your_access_token"
	req6.AdvertiserId = 12078245
	req6.AppId = 2199123262731
	req6.ShareType = 1
	ids := make([]int64, 201)
	for i := range ids {
		ids[i] = int64(i + 1)
	}
	req6.ShareAdvertiserIds = ids
	_, err6 := adapter.AppShareAdd(ctx, req6)
	if err6 == nil {
		t.Fatal("expected validation error for share_advertiser_ids exceeding 200 items")
	}
	fmt.Printf("got expected error: %v\n", err6)

	// share_type=2 但未填 share_corp_ids，预期返回校验错误
	req7 := &kuaishouModel.AppShareAddReq{}
	req7.AccessToken = "your_access_token"
	req7.AdvertiserId = 12078245
	req7.AppId = 2199123262731
	req7.ShareType = 2
	_, err7 := adapter.AppShareAdd(ctx, req7)
	if err7 == nil {
		t.Fatal("expected validation error for missing share_corp_ids")
	}
	fmt.Printf("got expected error: %v\n", err7)
}
