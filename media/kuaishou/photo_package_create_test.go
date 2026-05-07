package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestPhotoPackageCreate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.PhotoPackageCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.Name = "MAPI创建素材包"
	req.PhotoIds = []string{"5196028205791837691", "5199405903511832514"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.PhotoPackageCreate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestPhotoPackageCreateWithoutPhotos(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.PhotoPackageCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.Name = "空素材包"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.PhotoPackageCreate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("photo_package_id=%d name=%s status=%d photo_add_quota=%d\n",
		resp.PhotoPackageId, resp.Name, resp.Status, resp.PhotoAddQuota)
}

func TestPhotoPackageCreateValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 name，预期返回校验错误
	req := &kuaishouModel.PhotoPackageCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	_, err := adapter.PhotoPackageCreate(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty name")
	}
	fmt.Printf("got expected error: %v\n", err)
}
