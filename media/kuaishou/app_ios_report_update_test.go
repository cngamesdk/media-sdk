package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppIosReportUpdate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppIosReportUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015366
	req.AppId = 2199123264406
	req.IosAppId = 1321803705
	req.PackageId = 1099611636631
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppIosReportUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result=%v\n", resp.Result)
}

func TestAppIosReportUpdateValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppIosReportUpdateReq{}
	req.AdvertiserId = 900015366
	req.AppId = 2199123264406
	req.IosAppId = 1321803705
	req.PackageId = 1099611636631
	_, err := adapter.AppIosReportUpdate(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppIosReportUpdateReq{}
	req2.AccessToken = "your_access_token"
	req2.AppId = 2199123264406
	req2.IosAppId = 1321803705
	req2.PackageId = 1099611636631
	_, err2 := adapter.AppIosReportUpdate(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 app_id，预期返回校验错误
	req3 := &kuaishouModel.AppIosReportUpdateReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015366
	req3.IosAppId = 1321803705
	req3.PackageId = 1099611636631
	_, err3 := adapter.AppIosReportUpdate(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty app_id")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// 缺少 ios_app_id，预期返回校验错误
	req4 := &kuaishouModel.AppIosReportUpdateReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 900015366
	req4.AppId = 2199123264406
	req4.PackageId = 1099611636631
	_, err4 := adapter.AppIosReportUpdate(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for empty ios_app_id")
	}
	fmt.Printf("got expected error: %v\n", err4)

	// 缺少 package_id，预期返回校验错误
	req5 := &kuaishouModel.AppIosReportUpdateReq{}
	req5.AccessToken = "your_access_token"
	req5.AdvertiserId = 900015366
	req5.AppId = 2199123264406
	req5.IosAppId = 1321803705
	_, err5 := adapter.AppIosReportUpdate(ctx, req5)
	if err5 == nil {
		t.Fatal("expected validation error for empty package_id")
	}
	fmt.Printf("got expected error: %v\n", err5)
}
