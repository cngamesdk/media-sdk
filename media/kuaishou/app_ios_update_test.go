package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppIosUpdate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppIosUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015366
	req.PackageId = 1099611636631
	req.AppId = 2199123264406
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppIosUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("app_id=%d global_app_id=%d package_id=%d privacy_id=%d\n",
		resp.AppId, resp.GlobalAppId, resp.PackageId, resp.PrivacyId)
}

func TestAppIosUpdateFull(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppIosUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015366
	req.PackageId = 1099611636631
	req.AppId = 2199123264406
	req.IosDownloadUrl = "https://apps.apple.com/app/id9921801"
	req.AppIconUrl = "https://p1.kuaishou.com/app_icon.png"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppIosUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("app_id=%d global_app_id=%d package_id=%d privacy_id=%d\n",
		resp.AppId, resp.GlobalAppId, resp.PackageId, resp.PrivacyId)
}

func TestAppIosUpdateValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppIosUpdateReq{}
	req.AdvertiserId = 900015366
	req.PackageId = 1099611636631
	req.AppId = 2199123264406
	_, err := adapter.AppIosUpdate(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppIosUpdateReq{}
	req2.AccessToken = "your_access_token"
	req2.PackageId = 1099611636631
	req2.AppId = 2199123264406
	_, err2 := adapter.AppIosUpdate(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 package_id，预期返回校验错误
	req3 := &kuaishouModel.AppIosUpdateReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015366
	req3.AppId = 2199123264406
	_, err3 := adapter.AppIosUpdate(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty package_id")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// 缺少 app_id，预期返回校验错误
	req4 := &kuaishouModel.AppIosUpdateReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 900015366
	req4.PackageId = 1099611636631
	_, err4 := adapter.AppIosUpdate(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for empty app_id")
	}
	fmt.Printf("got expected error: %v\n", err4)
}
