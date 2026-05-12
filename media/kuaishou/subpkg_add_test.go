package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestSubpkgAdd(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.SubpkgAddReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015391
	req.ParentPackageId = 1099611757910
	req.Type = 2
	req.ChannelColumns = []kuaishouModel.ChannelColumn{
		{ChannelName: "测试备注1", Description: "备注-1"},
		{ChannelName: "测试备注2", Description: "备注-2"},
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.SubpkgAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range *resp {
		fmt.Printf("package_id=%d build_status=%d channel_id=%s parent_package_id=%d description=%s\n",
			item.PackageId, item.BuildStatus, item.ChannelId, item.ParentPackageId, item.Description)
	}
}

func TestSubpkgAddAuto(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.SubpkgAddReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015391
	req.ParentPackageId = 1099611757910
	req.Type = 1
	req.Count = 10
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.SubpkgAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range *resp {
		fmt.Printf("package_id=%d build_status=%d channel_id=%s parent_package_id=%d description=%s\n",
			item.PackageId, item.BuildStatus, item.ChannelId, item.ParentPackageId, item.Description)
	}
}

func TestSubpkgAddValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.SubpkgAddReq{}
	req.AdvertiserId = 900015391
	req.ParentPackageId = 1099611757910
	req.Type = 2
	_, err := adapter.SubpkgAdd(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.SubpkgAddReq{}
	req2.AccessToken = "your_access_token"
	req2.ParentPackageId = 1099611757910
	req2.Type = 2
	_, err2 := adapter.SubpkgAdd(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 parent_package_id，预期返回校验错误
	req3 := &kuaishouModel.SubpkgAddReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015391
	req3.Type = 2
	_, err3 := adapter.SubpkgAdd(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty parent_package_id")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// type 无效，预期返回校验错误
	req4 := &kuaishouModel.SubpkgAddReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 900015391
	req4.ParentPackageId = 1099611757910
	req4.Type = 0
	_, err4 := adapter.SubpkgAdd(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for invalid type")
	}
	fmt.Printf("got expected error: %v\n", err4)
}
