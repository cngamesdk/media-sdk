package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppAndroidUpdate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppAndroidUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015366
	req.AppInfo = kuaishouModel.AppAndroidUpdateAppInfo{
		AppId:       2199123264333,
		ReleaseType: 1, // 手动发版
	}
	req.AppPrivacyInfo = kuaishouModel.AppAndroidPrivacyInfo{
		PrivacyId: 4801049,
		Url:       "https://example.com/privacy",
	}
	req.PackageInfo = kuaishouModel.AppAndroidUpdatePackageInfo{
		PackageId:    1099611636558,
		BlobStoreKey: "your_blob_store_key_from_apk_upload",
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppAndroidUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("app_id=%d global_app_id=%d package_id=%d privacy_id=%d task_id=%d\n",
		resp.AppId, resp.GlobalAppId, resp.PackageId, resp.PrivacyId, resp.TaskId)
}

func TestAppAndroidUpdateFull(t *testing.T) {
	ctx := context.Background()
	applyAge := 1
	category := 1
	onlineEarnType := 1
	useSdk := 1
	networkType := 1
	req := &kuaishouModel.AppAndroidUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015366
	req.AppInfo = kuaishouModel.AppAndroidUpdateAppInfo{
		AppId:                  2199123264333,
		ReleaseType:            2, // 自动发版
		AppDetailImg:           "https://p1.kuaishou.com/app_detail.png",
		AppIconUrl:             "https://p1.kuaishou.com/app_icon.png",
		ApplyAge:               &applyAge,
		Category:               &category,
		ContactEmail:           "test@kuaishou.com",
		ContactName:            "后端测试联系人",
		ContactTel:             "18888888888",
		Developer:              "后端测试",
		Location:               `["北京","北京","海淀"]`,
		OfflineAppStores:       `["xiaomi"]`,
		OnlineEarnType:         &onlineEarnType,
		UseSdk:                 &useSdk,
		FunctionIntroduction:   "本应用提供xxx功能，支持用户进行xxx操作，帮助用户解决xxx问题，提升用户体验。",
		RecordNumber:           "911823145670",
		DocumentNumber:         "563178963123",
		ServiceCategory:        "[500,21]",
		NetworkType:            &networkType,
		RecordCorpName:         "快手科技",
		AppRecordScreenshotUrl: "https://static.yximgs.com/udata/pkg/3ba2fb.png",
		RecordCorpLicenseUrl:   "https://static.yximgs.com/udata/pkg/3ba2fb.png",
	}
	req.AppPrivacyInfo = kuaishouModel.AppAndroidPrivacyInfo{
		PrivacyId: 4801049,
		Url:       "https://example.com/privacy",
	}
	req.PackageInfo = kuaishouModel.AppAndroidUpdatePackageInfo{
		PackageId:    1099611636558,
		AppName:      "示例应用",
		BlobStoreKey: "your_blob_store_key_from_apk_upload",
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppAndroidUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("app_id=%d global_app_id=%d package_id=%d privacy_id=%d task_id=%d\n",
		resp.AppId, resp.GlobalAppId, resp.PackageId, resp.PrivacyId, resp.TaskId)
}

func TestAppAndroidUpdateValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppAndroidUpdateReq{}
	req.AdvertiserId = 900015366
	req.AppInfo = kuaishouModel.AppAndroidUpdateAppInfo{AppId: 2199123264333, ReleaseType: 1}
	req.AppPrivacyInfo = kuaishouModel.AppAndroidPrivacyInfo{PrivacyId: 4801049}
	req.PackageInfo = kuaishouModel.AppAndroidUpdatePackageInfo{PackageId: 1099611636558}
	_, err := adapter.AppAndroidUpdate(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppAndroidUpdateReq{}
	req2.AccessToken = "your_access_token"
	req2.AppInfo = kuaishouModel.AppAndroidUpdateAppInfo{AppId: 2199123264333, ReleaseType: 1}
	req2.AppPrivacyInfo = kuaishouModel.AppAndroidPrivacyInfo{PrivacyId: 4801049}
	req2.PackageInfo = kuaishouModel.AppAndroidUpdatePackageInfo{PackageId: 1099611636558}
	_, err2 := adapter.AppAndroidUpdate(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 app_info.app_id，预期返回校验错误
	req3 := &kuaishouModel.AppAndroidUpdateReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015366
	req3.AppInfo = kuaishouModel.AppAndroidUpdateAppInfo{ReleaseType: 1}
	req3.AppPrivacyInfo = kuaishouModel.AppAndroidPrivacyInfo{PrivacyId: 4801049}
	req3.PackageInfo = kuaishouModel.AppAndroidUpdatePackageInfo{PackageId: 1099611636558}
	_, err3 := adapter.AppAndroidUpdate(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty app_id")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// release_type 非法值，预期返回校验错误
	req4 := &kuaishouModel.AppAndroidUpdateReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 900015366
	req4.AppInfo = kuaishouModel.AppAndroidUpdateAppInfo{AppId: 2199123264333, ReleaseType: 3}
	req4.AppPrivacyInfo = kuaishouModel.AppAndroidPrivacyInfo{PrivacyId: 4801049}
	req4.PackageInfo = kuaishouModel.AppAndroidUpdatePackageInfo{PackageId: 1099611636558}
	_, err4 := adapter.AppAndroidUpdate(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for invalid release_type")
	}
	fmt.Printf("got expected error: %v\n", err4)

	// 缺少 app_privacy_info.privacy_id，预期返回校验错误
	req5 := &kuaishouModel.AppAndroidUpdateReq{}
	req5.AccessToken = "your_access_token"
	req5.AdvertiserId = 900015366
	req5.AppInfo = kuaishouModel.AppAndroidUpdateAppInfo{AppId: 2199123264333, ReleaseType: 1}
	req5.PackageInfo = kuaishouModel.AppAndroidUpdatePackageInfo{PackageId: 1099611636558}
	_, err5 := adapter.AppAndroidUpdate(ctx, req5)
	if err5 == nil {
		t.Fatal("expected validation error for empty privacy_id")
	}
	fmt.Printf("got expected error: %v\n", err5)

	// 缺少 package_info.package_id，预期返回校验错误
	req6 := &kuaishouModel.AppAndroidUpdateReq{}
	req6.AccessToken = "your_access_token"
	req6.AdvertiserId = 900015366
	req6.AppInfo = kuaishouModel.AppAndroidUpdateAppInfo{AppId: 2199123264333, ReleaseType: 1}
	req6.AppPrivacyInfo = kuaishouModel.AppAndroidPrivacyInfo{PrivacyId: 4801049}
	_, err6 := adapter.AppAndroidUpdate(ctx, req6)
	if err6 == nil {
		t.Fatal("expected validation error for empty package_id")
	}
	fmt.Printf("got expected error: %v\n", err6)
}
