package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestSubpkgDescription(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.SubpkgDescriptionReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.PackageId = 100009135
	req.Description = "应用分包备注"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.SubpkgDescription(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result=%v\n", resp.Result)
}

func TestSubpkgDescriptionValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.SubpkgDescriptionReq{}
	req.AdvertiserId = 900015364
	req.PackageId = 100009135
	req.Description = "应用分包备注"
	_, err := adapter.SubpkgDescription(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.SubpkgDescriptionReq{}
	req2.AccessToken = "your_access_token"
	req2.PackageId = 100009135
	req2.Description = "应用分包备注"
	_, err2 := adapter.SubpkgDescription(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 package_id，预期返回校验错误
	req3 := &kuaishouModel.SubpkgDescriptionReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015364
	req3.Description = "应用分包备注"
	_, err3 := adapter.SubpkgDescription(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty package_id")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// 缺少 description，预期返回校验错误
	req4 := &kuaishouModel.SubpkgDescriptionReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 900015364
	req4.PackageId = 100009135
	_, err4 := adapter.SubpkgDescription(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for empty description")
	}
	fmt.Printf("got expected error: %v\n", err4)
}
