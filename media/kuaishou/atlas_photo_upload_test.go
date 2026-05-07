package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAtlasPhotoUpload(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AtlasPhotoUploadReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 139494
	req.PicIds = []string{"5215449848999249535"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AtlasPhotoUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("photo_id=%s\n", resp.PhotoId)
}

func TestAtlasPhotoUploadWithOptions(t *testing.T) {
	ctx := context.Background()
	shieldBackward := false
	waitForTranscode := true
	req := &kuaishouModel.AtlasPhotoUploadReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 139494
	req.PicIds = []string{"5215449848999249535", "5220516399508008924", "5206724126146862162"}
	req.AudioBsKey = "55a4aebd6b544b6aa4ac99ef9c004c18"
	req.ShieldBackwardSwitch = &shieldBackward
	req.WaitForTranscode = &waitForTranscode
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AtlasPhotoUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestAtlasPhotoUploadValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 pic_ids，预期返回校验错误
	req := &kuaishouModel.AtlasPhotoUploadReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 139494
	_, err := adapter.AtlasPhotoUpload(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty pic_ids")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AtlasPhotoUploadReq{}
	req2.AccessToken = "your_access_token"
	req2.PicIds = []string{"5215449848999249535"}
	_, err2 := adapter.AtlasPhotoUpload(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)
}
