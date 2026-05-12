package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppVersionList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppVersionListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.AppId = 2199123264034
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppVersionList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d current_page=%d page_size=%d\n",
		resp.TotalCount, resp.CurrentPage, resp.PageSize)
	for _, item := range resp.List {
		fmt.Printf("real_app_version=%s version_code=%d version_status=%d update_time=%d\n",
			item.RealAppVersion, item.VersionCode, item.VersionStatus, item.UpdateTime)
	}
}

func TestAppVersionListWithPagination(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppVersionListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.AppId = 2199123264034
	req.Page = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppVersionList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d current_page=%d page_size=%d\n",
		resp.TotalCount, resp.CurrentPage, resp.PageSize)
	for _, item := range resp.List {
		fmt.Printf("real_app_version=%s version_code=%d version_status=%d update_time=%d\n",
			item.RealAppVersion, item.VersionCode, item.VersionStatus, item.UpdateTime)
	}
}

func TestAppVersionListValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppVersionListReq{}
	req.AdvertiserId = 900015364
	req.AppId = 2199123264034
	_, err := adapter.AppVersionList(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppVersionListReq{}
	req2.AccessToken = "your_access_token"
	req2.AppId = 2199123264034
	_, err2 := adapter.AppVersionList(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 app_id，预期返回校验错误
	req3 := &kuaishouModel.AppVersionListReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015364
	_, err3 := adapter.AppVersionList(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty app_id")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
