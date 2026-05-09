package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppAndroidCreate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppAndroidCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.AppInfo = kuaishouModel.AppAndroidAppInfo{
		AppId:       123456789,
		ReleaseType: 1, // 手动发版
	}
	req.AppPrivacyInfo = kuaishouModel.AppAndroidPrivacyInfo{
		PrivacyId: 100001,
		Url:       "https://example.com/privacy",
	}
	req.PackageInfo = kuaishouModel.AppAndroidPackageInfo{
		BlobStoreKey: "your_blob_store_key_from_apk_upload",
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppAndroidCreate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("app_id=%d global_app_id=%d package_id=%d privacy_id=%d task_id=%d\n",
		resp.AppId, resp.GlobalAppId, resp.PackageId, resp.PrivacyId, resp.TaskId)
}

func TestAppAndroidCreateFull(t *testing.T) {
	ctx := context.Background()
	applyAge := 1       // 全年龄
	category := 1       // 软件
	onlineEarnType := 2 // 非盈利
	useSdk := 1         // 已接入快手SDK
	createSource := 1   // 文件创编
	networkType := 1    // 线上
	req := &kuaishouModel.AppAndroidCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.AppInfo = kuaishouModel.AppAndroidAppInfo{
		AppId:                123456789,
		ReleaseType:          2, // 自动发版
		AppDetailImg:         "https://p1.kuaishou.com/app_detail.jpg",
		AppIconUrl:           "https://p1.kuaishou.com/app_icon.png",
		ApplyAge:             &applyAge,
		Category:             &category,
		ContactEmail:         "dev@example.com",
		ContactName:          "张三",
		ContactTel:           "13800138000",
		Description:          "这是一款优秀的应用",
		Developer:            "北京示例科技有限公司",
		Location:             `["北京市","北京市","海淀区"]`,
		OnlineEarnType:       &onlineEarnType,
		UseSdk:               &useSdk,
		CreateSource:         &createSource,
		FunctionIntroduction: "本应用提供xxx功能，支持用户进行xxx操作，帮助用户解决xxx问题，提升用户体验。",
		RecordNumber:         "京ICP备12345678号",
		NetworkType:          &networkType,
		RecordCorpName:       "北京示例科技有限公司",
	}
	req.AppPrivacyInfo = kuaishouModel.AppAndroidPrivacyInfo{
		PrivacyId: 100001,
		Url:       "https://example.com/privacy",
	}
	req.PackageInfo = kuaishouModel.AppAndroidPackageInfo{
		AppName:      "示例应用",
		BlobStoreKey: "your_blob_store_key_from_apk_upload",
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppAndroidCreate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("app_id=%d global_app_id=%d package_id=%d privacy_id=%d task_id=%d\n",
		resp.AppId, resp.GlobalAppId, resp.PackageId, resp.PrivacyId, resp.TaskId)
}

func TestAppAndroidCreateValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppAndroidCreateReq{}
	req.AdvertiserId = 20000681
	req.AppInfo = kuaishouModel.AppAndroidAppInfo{AppId: 123456789, ReleaseType: 1}
	req.AppPrivacyInfo = kuaishouModel.AppAndroidPrivacyInfo{PrivacyId: 100001}
	_, err := adapter.AppAndroidCreate(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppAndroidCreateReq{}
	req2.AccessToken = "your_access_token"
	req2.AppInfo = kuaishouModel.AppAndroidAppInfo{AppId: 123456789, ReleaseType: 1}
	req2.AppPrivacyInfo = kuaishouModel.AppAndroidPrivacyInfo{PrivacyId: 100001}
	_, err2 := adapter.AppAndroidCreate(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 app_info.app_id，预期返回校验错误
	req3 := &kuaishouModel.AppAndroidCreateReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 20000681
	req3.AppInfo = kuaishouModel.AppAndroidAppInfo{ReleaseType: 1}
	req3.AppPrivacyInfo = kuaishouModel.AppAndroidPrivacyInfo{PrivacyId: 100001}
	_, err3 := adapter.AppAndroidCreate(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty app_id")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// release_type 非法值，预期返回校验错误
	req4 := &kuaishouModel.AppAndroidCreateReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 20000681
	req4.AppInfo = kuaishouModel.AppAndroidAppInfo{AppId: 123456789, ReleaseType: 3} // 非法值
	req4.AppPrivacyInfo = kuaishouModel.AppAndroidPrivacyInfo{PrivacyId: 100001}
	_, err4 := adapter.AppAndroidCreate(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for invalid release_type")
	}
	fmt.Printf("got expected error: %v\n", err4)

	// 缺少 app_privacy_info.privacy_id，预期返回校验错误
	req5 := &kuaishouModel.AppAndroidCreateReq{}
	req5.AccessToken = "your_access_token"
	req5.AdvertiserId = 20000681
	req5.AppInfo = kuaishouModel.AppAndroidAppInfo{AppId: 123456789, ReleaseType: 1}
	_, err5 := adapter.AppAndroidCreate(ctx, req5)
	if err5 == nil {
		t.Fatal("expected validation error for empty privacy_id")
	}
	fmt.Printf("got expected error: %v\n", err5)
}
