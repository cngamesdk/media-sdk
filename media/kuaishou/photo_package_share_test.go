package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestPhotoPackageShare(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.PhotoPackageShareReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.PhotoPackageIds = []int64{11135761638}
	req.ShareAdvertiserIds = []int64{21237604}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.PhotoPackageShare(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("success=%v partSuccess=%v failed=%v\n",
		resp.SuccessPhotoPackageIds, resp.PartSuccessPhotoPackageIds, resp.FailedPhotoPackageIds)
}

func TestPhotoPackageShareMultiple(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.PhotoPackageShareReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.PhotoPackageIds = []int64{11135761637, 11135761638}
	req.ShareAdvertiserIds = []int64{21237604, 21237605}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.PhotoPackageShare(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestPhotoPackageShareValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 photo_package_ids，预期返回校验错误
	req := &kuaishouModel.PhotoPackageShareReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.ShareAdvertiserIds = []int64{21237604}
	_, err := adapter.PhotoPackageShare(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty photo_package_ids")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 share_advertiser_ids，预期返回校验错误
	req2 := &kuaishouModel.PhotoPackageShareReq{}
	req2.AccessToken = "your_access_token"
	req2.AdvertiserId = 16859321
	req2.PhotoPackageIds = []int64{11135761638}
	_, err2 := adapter.PhotoPackageShare(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty share_advertiser_ids")
	}
	fmt.Printf("got expected error: %v\n", err2)
}
