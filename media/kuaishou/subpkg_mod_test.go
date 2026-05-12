package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestSubpkgMod(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.SubpkgModReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.PackageId = []int64{100009135}
	req.PutStatus = 2 // 删除
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.SubpkgMod(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result=%v\n", resp.Result)
}

func TestSubpkgModUpdate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.SubpkgModReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.PackageId = []int64{100009135}
	req.PutStatus = 0 // 更新
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.SubpkgMod(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result=%v\n", resp.Result)
}

func TestSubpkgModRestore(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.SubpkgModReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.PackageId = []int64{100009135}
	req.PutStatus = 1 // 恢复
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.SubpkgMod(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result=%v\n", resp.Result)
}

func TestSubpkgModValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.SubpkgModReq{}
	req.AdvertiserId = 900015364
	req.PackageId = []int64{100009135}
	req.PutStatus = 2
	_, err := adapter.SubpkgMod(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.SubpkgModReq{}
	req2.AccessToken = "your_access_token"
	req2.PackageId = []int64{100009135}
	req2.PutStatus = 2
	_, err2 := adapter.SubpkgMod(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 package_id，预期返回校验错误
	req3 := &kuaishouModel.SubpkgModReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015364
	req3.PutStatus = 2
	_, err3 := adapter.SubpkgMod(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty package_id")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// put_status 无效，预期返回校验错误
	req4 := &kuaishouModel.SubpkgModReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 900015364
	req4.PackageId = []int64{100009135}
	req4.PutStatus = 3
	_, err4 := adapter.SubpkgMod(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for invalid put_status")
	}
	fmt.Printf("got expected error: %v\n", err4)
}
