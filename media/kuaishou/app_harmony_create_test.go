package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppHarmonyCreate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppHarmonyCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.AppIconUrl = "https://p1.kuaishou.com/app_icon.png"
	req.AppName = "示例应用"
	req.Developer = "北京示例科技有限公司"
	req.PackageName = "com.example.app"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppHarmonyCreate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("app_id=%d global_app_id=%d package_id=%d\n",
		resp.AppId, resp.GlobalAppId, resp.PackageId)
}

func TestAppHarmonyCreateValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppHarmonyCreateReq{}
	req.AdvertiserId = 20000681
	req.AppIconUrl = "https://p1.kuaishou.com/app_icon.png"
	req.AppName = "示例应用"
	req.Developer = "北京示例科技有限公司"
	req.PackageName = "com.example.app"
	_, err := adapter.AppHarmonyCreate(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppHarmonyCreateReq{}
	req2.AccessToken = "your_access_token"
	req2.AppIconUrl = "https://p1.kuaishou.com/app_icon.png"
	req2.AppName = "示例应用"
	req2.Developer = "北京示例科技有限公司"
	req2.PackageName = "com.example.app"
	_, err2 := adapter.AppHarmonyCreate(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 app_icon_url，预期返回校验错误
	req3 := &kuaishouModel.AppHarmonyCreateReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 20000681
	req3.AppName = "示例应用"
	req3.Developer = "北京示例科技有限公司"
	req3.PackageName = "com.example.app"
	_, err3 := adapter.AppHarmonyCreate(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty app_icon_url")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// 缺少 app_name，预期返回校验错误
	req4 := &kuaishouModel.AppHarmonyCreateReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 20000681
	req4.AppIconUrl = "https://p1.kuaishou.com/app_icon.png"
	req4.Developer = "北京示例科技有限公司"
	req4.PackageName = "com.example.app"
	_, err4 := adapter.AppHarmonyCreate(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for empty app_name")
	}
	fmt.Printf("got expected error: %v\n", err4)

	// 缺少 developer，预期返回校验错误
	req5 := &kuaishouModel.AppHarmonyCreateReq{}
	req5.AccessToken = "your_access_token"
	req5.AdvertiserId = 20000681
	req5.AppIconUrl = "https://p1.kuaishou.com/app_icon.png"
	req5.AppName = "示例应用"
	req5.PackageName = "com.example.app"
	_, err5 := adapter.AppHarmonyCreate(ctx, req5)
	if err5 == nil {
		t.Fatal("expected validation error for empty developer")
	}
	fmt.Printf("got expected error: %v\n", err5)

	// 缺少 package_name，预期返回校验错误
	req6 := &kuaishouModel.AppHarmonyCreateReq{}
	req6.AccessToken = "your_access_token"
	req6.AdvertiserId = 20000681
	req6.AppIconUrl = "https://p1.kuaishou.com/app_icon.png"
	req6.AppName = "示例应用"
	req6.Developer = "北京示例科技有限公司"
	_, err6 := adapter.AppHarmonyCreate(ctx, req6)
	if err6 == nil {
		t.Fatal("expected validation error for empty package_name")
	}
	fmt.Printf("got expected error: %v\n", err6)
}
