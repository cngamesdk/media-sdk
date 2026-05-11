package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppShareCorpAccountList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppShareCorpAccountListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12078245
	req.AppId = 2199123262731
	req.CorpId = 15280
	req.PageInfo = kuaishouModel.AppShareCorpAccountPageInfo{
		CurrentPage: 1,
		PageSize:    10,
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppShareCorpAccountList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("list_len=%d\n", len(resp.List))
	for _, item := range resp.List {
		fmt.Printf("  corp_id=%d corp_name=%s total_account_cnt=%d accounts=%d\n",
			item.CorpId, item.CorpName, item.TotalAccountCnt, len(item.ShareAccountVos))
		for _, acc := range item.ShareAccountVos {
			fmt.Printf("    account_id=%d account_name=%s\n", acc.AccountId, acc.AccountName)
		}
	}
}

func TestAppShareCorpAccountListValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppShareCorpAccountListReq{}
	req.AdvertiserId = 12078245
	req.AppId = 2199123262731
	req.CorpId = 15280
	req.PageInfo = kuaishouModel.AppShareCorpAccountPageInfo{CurrentPage: 1, PageSize: 10}
	_, err := adapter.AppShareCorpAccountList(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppShareCorpAccountListReq{}
	req2.AccessToken = "your_access_token"
	req2.AppId = 2199123262731
	req2.CorpId = 15280
	req2.PageInfo = kuaishouModel.AppShareCorpAccountPageInfo{CurrentPage: 1, PageSize: 10}
	_, err2 := adapter.AppShareCorpAccountList(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 app_id，预期返回校验错误
	req3 := &kuaishouModel.AppShareCorpAccountListReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 12078245
	req3.CorpId = 15280
	req3.PageInfo = kuaishouModel.AppShareCorpAccountPageInfo{CurrentPage: 1, PageSize: 10}
	_, err3 := adapter.AppShareCorpAccountList(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty app_id")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// 缺少 corp_id，预期返回校验错误
	req4 := &kuaishouModel.AppShareCorpAccountListReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 12078245
	req4.AppId = 2199123262731
	req4.PageInfo = kuaishouModel.AppShareCorpAccountPageInfo{CurrentPage: 1, PageSize: 10}
	_, err4 := adapter.AppShareCorpAccountList(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for empty corp_id")
	}
	fmt.Printf("got expected error: %v\n", err4)

	// page_info.current_page 未填，预期返回校验错误
	req5 := &kuaishouModel.AppShareCorpAccountListReq{}
	req5.AccessToken = "your_access_token"
	req5.AdvertiserId = 12078245
	req5.AppId = 2199123262731
	req5.CorpId = 15280
	req5.PageInfo = kuaishouModel.AppShareCorpAccountPageInfo{PageSize: 10}
	_, err5 := adapter.AppShareCorpAccountList(ctx, req5)
	if err5 == nil {
		t.Fatal("expected validation error for empty current_page")
	}
	fmt.Printf("got expected error: %v\n", err5)

	// page_info.page_size 未填，预期返回校验错误
	req6 := &kuaishouModel.AppShareCorpAccountListReq{}
	req6.AccessToken = "your_access_token"
	req6.AdvertiserId = 12078245
	req6.AppId = 2199123262731
	req6.CorpId = 15280
	req6.PageInfo = kuaishouModel.AppShareCorpAccountPageInfo{CurrentPage: 1}
	_, err6 := adapter.AppShareCorpAccountList(ctx, req6)
	if err6 == nil {
		t.Fatal("expected validation error for empty page_size")
	}
	fmt.Printf("got expected error: %v\n", err6)
}
