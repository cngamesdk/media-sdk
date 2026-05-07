package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestPhotoPackageDel(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.PhotoPackageDelReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.PhotoPackageIds = []int64{11135761637}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.PhotoPackageDel(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("deleted photo_package_ids=%v\n", resp.PhotoPackageIds)
}

func TestPhotoPackageDelMultiple(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.PhotoPackageDelReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.PhotoPackageIds = []int64{11135761637, 11135761638}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.PhotoPackageDel(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestPhotoPackageDelValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 photo_package_ids，预期返回校验错误
	req := &kuaishouModel.PhotoPackageDelReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	_, err := adapter.PhotoPackageDel(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty photo_package_ids")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.PhotoPackageDelReq{}
	req2.AccessToken = "your_access_token"
	req2.PhotoPackageIds = []int64{11135761637}
	_, err2 := adapter.PhotoPackageDel(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)
}
