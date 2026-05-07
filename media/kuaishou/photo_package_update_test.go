package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestPhotoPackageUpdate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.PhotoPackageUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.PhotoPackageId = 11135761638
	req.Name = "素材包MAPI编辑"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.PhotoPackageUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("photo_package_id=%d\n", resp.PhotoPackageId)
}

func TestPhotoPackageUpdateWithPhotos(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.PhotoPackageUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.PhotoPackageId = 11135761638
	req.Name = "素材包MAPI编辑"
	req.PhotoIds = []string{"5196028205791837691"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.PhotoPackageUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestPhotoPackageUpdateValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 photo_package_id，预期返回校验错误
	req := &kuaishouModel.PhotoPackageUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.Name = "素材包MAPI编辑"
	_, err := adapter.PhotoPackageUpdate(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty photo_package_id")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.PhotoPackageUpdateReq{}
	req2.AccessToken = "your_access_token"
	req2.PhotoPackageId = 11135761638
	_, err2 := adapter.PhotoPackageUpdate(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)
}
