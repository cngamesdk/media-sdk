package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppShareCanShareCorpList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppShareCanShareCorpListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12078245
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppShareCanShareCorpList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("list_len=%d\n", len(resp.List))
	for _, item := range resp.List {
		fmt.Printf("  corp_id=%d corp_name=%s\n", item.CorpId, item.CorpName)
	}
}

func TestAppShareCanShareCorpListValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppShareCanShareCorpListReq{}
	req.AdvertiserId = 12078245
	_, err := adapter.AppShareCanShareCorpList(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppShareCanShareCorpListReq{}
	req2.AccessToken = "your_access_token"
	_, err2 := adapter.AppShareCanShareCorpList(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)
}
