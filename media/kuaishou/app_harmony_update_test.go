package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppHarmonyUpdate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppHarmonyUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.AppId = 123456789
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppHarmonyUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("app_id=%d package_id=%d\n", resp.AppId, resp.PackageId)
}

func TestAppHarmonyUpdateFull(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppHarmonyUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.AppId = 123456789
	req.AppName = "示例应用"
	req.Developer = "北京示例科技有限公司"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppHarmonyUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("app_id=%d package_id=%d\n", resp.AppId, resp.PackageId)
}

func TestAppHarmonyUpdateValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppHarmonyUpdateReq{}
	req.AdvertiserId = 20000681
	_, err := adapter.AppHarmonyUpdate(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppHarmonyUpdateReq{}
	req2.AccessToken = "your_access_token"
	_, err2 := adapter.AppHarmonyUpdate(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)
}
