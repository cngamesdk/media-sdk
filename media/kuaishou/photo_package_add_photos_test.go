package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestPhotoPackageAddPhotos(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.PhotoPackageAddPhotosReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.PhotoPackageId = 11135761637
	req.PhotoIds = []string{"5196028205791837691"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.PhotoPackageAddPhotos(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("success=%v dup=%v\n", resp.SuccessPhotoIds, resp.DupPhotoIds)
}

func TestPhotoPackageAddPhotosMultiple(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.PhotoPackageAddPhotosReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.PhotoPackageId = 11135761637
	req.PhotoIds = []string{"5196028205791837691", "5199405903511832514"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.PhotoPackageAddPhotos(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestPhotoPackageAddPhotosValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 photo_package_id，预期返回校验错误
	req := &kuaishouModel.PhotoPackageAddPhotosReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.PhotoIds = []string{"5196028205791837691"}
	_, err := adapter.PhotoPackageAddPhotos(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty photo_package_id")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 photo_ids，预期返回校验错误
	req2 := &kuaishouModel.PhotoPackageAddPhotosReq{}
	req2.AccessToken = "your_access_token"
	req2.AdvertiserId = 16859321
	req2.PhotoPackageId = 11135761637
	_, err2 := adapter.PhotoPackageAddPhotos(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty photo_ids")
	}
	fmt.Printf("got expected error: %v\n", err2)
}
