package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppOffline(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppOfflineReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015366
	req.PackageIds = []int64{1099611634749}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppOffline(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result=%v\n", resp.Result)
}

func TestAppOfflineValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppOfflineReq{}
	req.AdvertiserId = 900015366
	req.PackageIds = []int64{1099611634749}
	_, err := adapter.AppOffline(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppOfflineReq{}
	req2.AccessToken = "your_access_token"
	req2.PackageIds = []int64{1099611634749}
	_, err2 := adapter.AppOffline(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 package_ids，预期返回校验错误
	req3 := &kuaishouModel.AppOfflineReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015366
	_, err3 := adapter.AppOffline(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty package_ids")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
