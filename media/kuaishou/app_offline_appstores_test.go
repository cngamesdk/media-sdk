package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppOfflineAppstores(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppOfflineAppstoresReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015366
	req.AppIds = []int64{2199123262524}
	req.OfflineStores = []string{"xiaomi"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppOfflineAppstores(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result=%v\n", resp.Result)
}

func TestAppOfflineAppstoresValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppOfflineAppstoresReq{}
	req.AdvertiserId = 900015366
	req.AppIds = []int64{2199123262524}
	req.OfflineStores = []string{"xiaomi"}
	_, err := adapter.AppOfflineAppstores(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppOfflineAppstoresReq{}
	req2.AccessToken = "your_access_token"
	req2.AppIds = []int64{2199123262524}
	req2.OfflineStores = []string{"xiaomi"}
	_, err2 := adapter.AppOfflineAppstores(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 app_ids，预期返回校验错误
	req3 := &kuaishouModel.AppOfflineAppstoresReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015366
	req3.OfflineStores = []string{"xiaomi"}
	_, err3 := adapter.AppOfflineAppstores(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty app_ids")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// 缺少 offline_stores，预期返回校验错误
	req4 := &kuaishouModel.AppOfflineAppstoresReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 900015366
	req4.AppIds = []int64{2199123262524}
	_, err4 := adapter.AppOfflineAppstores(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for empty offline_stores")
	}
	fmt.Printf("got expected error: %v\n", err4)
}
