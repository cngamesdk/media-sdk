package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestImageUploadTokenVerify(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.ImageUploadTokenVerifyReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.UploadToken = "your_upload_token"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ImageUploadTokenVerify(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("blob_store_key=%s\n", resp.BlobStoreKey)
}

func TestImageUploadTokenVerifyValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.ImageUploadTokenVerifyReq{}
	req.AdvertiserId = 20000681
	req.UploadToken = "your_upload_token"
	_, err := adapter.ImageUploadTokenVerify(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.ImageUploadTokenVerifyReq{}
	req2.AccessToken = "your_access_token"
	req2.UploadToken = "your_upload_token"
	_, err2 := adapter.ImageUploadTokenVerify(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 upload_token，预期返回校验错误
	req3 := &kuaishouModel.ImageUploadTokenVerifyReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 20000681
	_, err3 := adapter.ImageUploadTokenVerify(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty upload_token")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
