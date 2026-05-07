package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestPhotoPackageRemovePhotos(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.PhotoPackageRemovePhotosReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.PhotoPackageId = 11135761637
	req.PhotoIds = []string{"5196028205791837691"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.PhotoPackageRemovePhotos(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("PhotoIds=%v\n", resp.PhotoIds)
}

func TestPhotoPackageRemovePhotosMultiple(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.PhotoPackageRemovePhotosReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.PhotoPackageId = 11135761637
	req.PhotoIds = []string{"5196028205791837691", "5199405903511832514"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.PhotoPackageRemovePhotos(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestPhotoPackageRemovePhotosValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 photo_package_id，预期返回校验错误
	req := &kuaishouModel.PhotoPackageRemovePhotosReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.PhotoIds = []string{"5196028205791837691"}
	_, err := adapter.PhotoPackageRemovePhotos(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty photo_package_id")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 photo_ids，预期返回校验错误
	req2 := &kuaishouModel.PhotoPackageRemovePhotosReq{}
	req2.AccessToken = "your_access_token"
	req2.AdvertiserId = 16859321
	req2.PhotoPackageId = 11135761637
	_, err2 := adapter.PhotoPackageRemovePhotos(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty photo_ids")
	}
	fmt.Printf("got expected error: %v\n", err2)
}
