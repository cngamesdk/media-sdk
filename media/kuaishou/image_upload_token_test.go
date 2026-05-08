package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestImageUploadToken(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.ImageUploadTokenReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.FileType = "mp4"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ImageUploadToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("upload_token=%s endpoint=%v\n", resp.UploadToken, resp.Endpoint)
}

func TestImageUploadTokenApk(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.ImageUploadTokenReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.FileType = "apk"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ImageUploadToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("upload_token=%s endpoint=%v\n", resp.UploadToken, resp.Endpoint)
}

func TestImageUploadTokenValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.ImageUploadTokenReq{}
	req.AdvertiserId = 20000681
	req.FileType = "mp4"
	_, err := adapter.ImageUploadToken(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.ImageUploadTokenReq{}
	req2.AccessToken = "your_access_token"
	req2.FileType = "mp4"
	_, err2 := adapter.ImageUploadToken(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 file_type，预期返回校验错误
	req3 := &kuaishouModel.ImageUploadTokenReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 20000681
	_, err3 := adapter.ImageUploadToken(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty file_type")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
