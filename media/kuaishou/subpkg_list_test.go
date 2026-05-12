package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestSubpkgList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.SubpkgListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.AppId = 2199123264588
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.SubpkgList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d current_page=%d page_size=%d\n",
		resp.TotalCount, resp.CurrentPage, resp.PageSize)
	for _, item := range resp.List {
		fmt.Printf("package_id=%d channel_id=%s sub_package_status=%d real_app_version=%s description=%s url=%s\n",
			item.PackageId, item.ChannelId, item.SubPackageStatus, item.RealAppVersion, item.Description, item.Url)
	}
}

func TestSubpkgListRecycleBin(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.SubpkgListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.AppId = 2199123264588
	req.ListType = 2 // 分包回收列表
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.SubpkgList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d current_page=%d page_size=%d\n",
		resp.TotalCount, resp.CurrentPage, resp.PageSize)
	for _, item := range resp.List {
		fmt.Printf("package_id=%d channel_id=%s sub_package_status=%d can_recycle=%v delete_time=%d\n",
			item.PackageId, item.ChannelId, item.SubPackageStatus, item.CanRecycle, item.DeleteTime)
	}
}

func TestSubpkgListWithFilters(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.SubpkgListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.AppId = 2199123264588
	req.Status = 4 // 已发布
	req.Page = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.SubpkgList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d current_page=%d page_size=%d\n",
		resp.TotalCount, resp.CurrentPage, resp.PageSize)
	for _, item := range resp.List {
		fmt.Printf("package_id=%d channel_id=%s sub_package_status=%d can_update=%v update_time=%d\n",
			item.PackageId, item.ChannelId, item.SubPackageStatus, item.CanUpdate, item.UpdateTime)
	}
}

func TestSubpkgListValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.SubpkgListReq{}
	req.AdvertiserId = 900015364
	req.AppId = 2199123264588
	_, err := adapter.SubpkgList(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.SubpkgListReq{}
	req2.AccessToken = "your_access_token"
	req2.AppId = 2199123264588
	_, err2 := adapter.SubpkgList(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 app_id，预期返回校验错误
	req3 := &kuaishouModel.SubpkgListReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015364
	_, err3 := adapter.SubpkgList(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty app_id")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
