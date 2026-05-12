package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppShareCanShareAccountList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppShareCanShareAccountListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12078245
	req.CurrentPage = 1
	req.PageSize = 10
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppShareCanShareAccountList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d current_page=%d page_size=%d list_len=%d\n",
		resp.TotalCount, resp.CurrentPage, resp.PageSize, len(resp.List))
	for _, item := range resp.List {
		fmt.Printf("  account_id=%d account_name=%s\n", item.AccountId, item.AccountName)
	}
}

func TestAppShareCanShareAccountListWithSearchIds(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppShareCanShareAccountListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12078245
	req.CurrentPage = 1
	req.PageSize = 10
	req.SearchAccountId = []int64{16855298, 13918416}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppShareCanShareAccountList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d list_len=%d\n", resp.TotalCount, len(resp.List))
	for _, item := range resp.List {
		fmt.Printf("  account_id=%d account_name=%s\n", item.AccountId, item.AccountName)
	}
}

func TestAppShareCanShareAccountListValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppShareCanShareAccountListReq{}
	req.AdvertiserId = 12078245
	req.CurrentPage = 1
	req.PageSize = 10
	_, err := adapter.AppShareCanShareAccountList(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppShareCanShareAccountListReq{}
	req2.AccessToken = "your_access_token"
	req2.CurrentPage = 1
	req2.PageSize = 10
	_, err2 := adapter.AppShareCanShareAccountList(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// current_page 未填，预期返回校验错误
	req3 := &kuaishouModel.AppShareCanShareAccountListReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 12078245
	req3.PageSize = 10
	_, err3 := adapter.AppShareCanShareAccountList(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty current_page")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// page_size 未填，预期返回校验错误
	req4 := &kuaishouModel.AppShareCanShareAccountListReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 12078245
	req4.CurrentPage = 1
	_, err4 := adapter.AppShareCanShareAccountList(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for empty page_size")
	}
	fmt.Printf("got expected error: %v\n", err4)

	// search_account_id 超过500个，预期返回校验错误
	req5 := &kuaishouModel.AppShareCanShareAccountListReq{}
	req5.AccessToken = "your_access_token"
	req5.AdvertiserId = 12078245
	req5.CurrentPage = 1
	req5.PageSize = 10
	ids := make([]int64, 501)
	for i := range ids {
		ids[i] = int64(i + 1)
	}
	req5.SearchAccountId = ids
	_, err5 := adapter.AppShareCanShareAccountList(ctx, req5)
	if err5 == nil {
		t.Fatal("expected validation error for search_account_id exceeding 500")
	}
	fmt.Printf("got expected error: %v\n", err5)
}
