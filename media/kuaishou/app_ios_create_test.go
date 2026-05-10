package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppIosCreate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppIosCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015366
	req.IosDownloadUrl = "https://apps.apple.com/cn/app/id1454750038"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppIosCreate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("app_id=%d global_app_id=%d package_id=%d privacy_id=%d\n",
		resp.AppId, resp.GlobalAppId, resp.PackageId, resp.PrivacyId)
}

func TestAppIosCreateFull(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppIosCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015366
	req.IosDownloadUrl = "https://apps.apple.com/cn/app/id1454750038"
	req.AppIconUrl = "https://p1.kuaishou.com/app_icon.png"
	req.AppPrivacyUrl = "https://example.com/privacy"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppIosCreate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("app_id=%d global_app_id=%d package_id=%d privacy_id=%d\n",
		resp.AppId, resp.GlobalAppId, resp.PackageId, resp.PrivacyId)
}

func TestAppIosCreateValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppIosCreateReq{}
	req.AdvertiserId = 900015366
	req.IosDownloadUrl = "https://apps.apple.com/cn/app/id1454750038"
	_, err := adapter.AppIosCreate(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppIosCreateReq{}
	req2.AccessToken = "your_access_token"
	req2.IosDownloadUrl = "https://apps.apple.com/cn/app/id1454750038"
	_, err2 := adapter.AppIosCreate(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 ios_download_url，预期返回校验错误
	req3 := &kuaishouModel.AppIosCreateReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015366
	_, err3 := adapter.AppIosCreate(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty ios_download_url")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
