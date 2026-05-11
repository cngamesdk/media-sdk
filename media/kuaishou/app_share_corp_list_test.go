package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppShareCorpList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppShareCorpListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12078245
	req.AppId = 2199123262731
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppShareCorpList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("list_len=%d\n", len(resp.List))
	for _, item := range resp.List {
		fmt.Printf("  corp_id=%d corp_name=%s total_account_cnt=%d\n",
			item.CorpId, item.CorpName, item.TotalAccountCnt)
	}
}

func TestAppShareCorpListValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppShareCorpListReq{}
	req.AdvertiserId = 12078245
	req.AppId = 2199123262731
	_, err := adapter.AppShareCorpList(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppShareCorpListReq{}
	req2.AccessToken = "your_access_token"
	req2.AppId = 2199123262731
	_, err2 := adapter.AppShareCorpList(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 app_id，预期返回校验错误
	req3 := &kuaishouModel.AppShareCorpListReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 12078245
	_, err3 := adapter.AppShareCorpList(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty app_id")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
