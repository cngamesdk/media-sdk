package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"os"
	"testing"
)

func TestAtlasAudioUpload(t *testing.T) {
	ctx := context.Background()
	fileData, err := os.ReadFile("/path/to/your/audio.mp4")
	if err != nil {
		t.Fatal(err)
	}
	req := &kuaishouModel.AtlasAudioUploadReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12078245
	req.File = fileData
	req.FileName = "audio.mp4"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AtlasAudioUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("blob_store_key=%s\n", resp.BlobStoreKey)
}

func TestAtlasAudioUploadWithBucketName(t *testing.T) {
	ctx := context.Background()
	fileData, err := os.ReadFile("/path/to/your/audio.mp4")
	if err != nil {
		t.Fatal(err)
	}
	req := &kuaishouModel.AtlasAudioUploadReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12078245
	req.File = fileData
	req.FileName = "audio_1.mp4"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AtlasAudioUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("blob_store_key=%s bucket_name=%s\n", resp.BlobStoreKey, resp.BucketName)
}

func TestAtlasAudioUploadValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 file，预期返回校验错误
	req := &kuaishouModel.AtlasAudioUploadReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12078245
	req.FileName = "audio.mp4"
	_, err := adapter.AtlasAudioUpload(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty file")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AtlasAudioUploadReq{}
	req2.AccessToken = "your_access_token"
	req2.File = []byte("fake")
	req2.FileName = "audio.mp4"
	_, err2 := adapter.AtlasAudioUpload(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)
}
