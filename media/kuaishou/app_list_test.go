package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12078245
	req.ListType = 1 // 我创建的
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d current_page=%d page_size=%d list_len=%d\n",
		resp.TotalCount, resp.CurrentPage, resp.PageSize, len(resp.List))
	for _, item := range resp.List {
		fmt.Printf("  app_id=%d app_name=%s platform=%s status=%d source_type=%d\n",
			item.AppId, item.RealAppName, item.Platform, item.AppStatus, item.SourceType)
	}
}

func TestAppListShared(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12078245
	req.ListType = 2 // 共享给我的
	req.Platform = "android"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d list_len=%d\n", resp.TotalCount, len(resp.List))
	for _, item := range resp.List {
		fmt.Printf("  app_id=%d app_name=%s source_account=%s(%d)\n",
			item.AppId, item.RealAppName, item.AppSource.AccountName, item.AppSource.AccountId)
	}
}

func TestAppListWithFilter(t *testing.T) {
	ctx := context.Background()
	appStatus := 4 // 已发布
	req := &kuaishouModel.AppListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12078245
	req.ListType = 1
	req.Platform = "android"
	req.AppStatus = &appStatus
	req.KeyWord = "示例应用"
	req.Page = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d list_len=%d\n", resp.TotalCount, len(resp.List))
}

func TestAppListByAppIds(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12078245
	req.ListType = 1
	req.AppIds = []string{"2199123264333", "2199123264406"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d list_len=%d\n", resp.TotalCount, len(resp.List))
}

func TestAppListByDateRange(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12078245
	req.ListType = 1
	req.StartDate = "2024-01-01"
	req.EndDate = "2024-12-31"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d list_len=%d\n", resp.TotalCount, len(resp.List))
}

func TestAppListValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppListReq{}
	req.AdvertiserId = 12078245
	req.ListType = 1
	_, err := adapter.AppList(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppListReq{}
	req2.AccessToken = "your_access_token"
	req2.ListType = 1
	_, err2 := adapter.AppList(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// list_type 非法值，预期返回校验错误
	req3 := &kuaishouModel.AppListReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 12078245
	req3.ListType = 3 // 非法值
	_, err3 := adapter.AppList(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for invalid list_type")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
