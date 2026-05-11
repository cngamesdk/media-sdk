package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppSubPackageReleaseList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppSubPackageReleaseListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.AppId = 2199123264588
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppSubPackageReleaseList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d current_page=%d page_size=%d list_len=%d\n",
		resp.TotalCount, resp.CurrentPage, resp.PageSize, len(resp.List))
	for _, item := range resp.List {
		fmt.Printf("  package_id=%d channel_id=%s app_name=%s platform=%s parent_package_id=%d\n",
			item.PackageId, item.ChannelId, item.RealAppName, item.Platform, item.ParentPackageId)
	}
}

func TestAppSubPackageReleaseListWithKeyword(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppSubPackageReleaseListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.AppId = 2199123264588
	req.KeyWord = "channel_001"
	req.Page = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppSubPackageReleaseList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d list_len=%d\n", resp.TotalCount, len(resp.List))
}

func TestAppSubPackageReleaseListValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppSubPackageReleaseListReq{}
	req.AdvertiserId = 900015364
	req.AppId = 2199123264588
	_, err := adapter.AppSubPackageReleaseList(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppSubPackageReleaseListReq{}
	req2.AccessToken = "your_access_token"
	req2.AppId = 2199123264588
	_, err2 := adapter.AppSubPackageReleaseList(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 app_id，预期返回校验错误
	req3 := &kuaishouModel.AppSubPackageReleaseListReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015364
	_, err3 := adapter.AppSubPackageReleaseList(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty app_id")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
