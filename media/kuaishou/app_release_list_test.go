package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppReleaseList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppReleaseListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppReleaseList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d current_page=%d page_size=%d list_len=%d\n",
		resp.TotalCount, resp.CurrentPage, resp.PageSize, len(resp.List))
	for _, item := range resp.List {
		fmt.Printf("  app_id=%d app_name=%s platform=%s package_name=%s source_type=%d\n",
			item.AppId, item.RealAppName, item.Platform, item.PackageName, item.SourceType)
	}
}

func TestAppReleaseListByPlatform(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppReleaseListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.Platform = "android"
	req.Page = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppReleaseList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d list_len=%d\n", resp.TotalCount, len(resp.List))
}

func TestAppReleaseListByKeyword(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppReleaseListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.KeyWord = "示例应用"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppReleaseList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d list_len=%d\n", resp.TotalCount, len(resp.List))
}

func TestAppReleaseListByListType(t *testing.T) {
	ctx := context.Background()
	listType := 1 // 我创建的
	req := &kuaishouModel.AppReleaseListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.ListType = &listType
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppReleaseList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d list_len=%d\n", resp.TotalCount, len(resp.List))
}

func TestAppReleaseListByAppIds(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppReleaseListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.AppIds = []int64{2199123264333, 2199123264406}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppReleaseList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d list_len=%d\n", resp.TotalCount, len(resp.List))
}

func TestAppReleaseListValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppReleaseListReq{}
	req.AdvertiserId = 900015364
	_, err := adapter.AppReleaseList(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppReleaseListReq{}
	req2.AccessToken = "your_access_token"
	_, err2 := adapter.AppReleaseList(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// app_ids 超过100个，预期返回校验错误
	req3 := &kuaishouModel.AppReleaseListReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015364
	ids := make([]int64, 101)
	for i := range ids {
		ids[i] = int64(i + 1)
	}
	req3.AppIds = ids
	_, err3 := adapter.AppReleaseList(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for app_ids exceeding 100 items")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
