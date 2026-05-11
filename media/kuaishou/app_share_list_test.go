package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppShareList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppShareListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.AppId = 2199123264636
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppShareList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d current_page=%d page_size=%d list_len=%d\n",
		resp.TotalCount, resp.CurrentPage, resp.PageSize, len(resp.List))
	for _, item := range resp.List {
		fmt.Printf("  account_id=%d account_name=%s\n", item.AccountId, item.AccountName)
	}
}

func TestAppShareListWithKeyword(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppShareListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.AppId = 2199123264636
	req.KeyWord = "快手用户"
	req.Page = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppShareList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d list_len=%d\n", resp.TotalCount, len(resp.List))
}

func TestAppShareListValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppShareListReq{}
	req.AdvertiserId = 900015364
	req.AppId = 2199123264636
	_, err := adapter.AppShareList(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppShareListReq{}
	req2.AccessToken = "your_access_token"
	req2.AppId = 2199123264636
	_, err2 := adapter.AppShareList(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 app_id，预期返回校验错误
	req3 := &kuaishouModel.AppShareListReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015364
	_, err3 := adapter.AppShareList(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty app_id")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
