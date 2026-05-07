package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"os"
	"testing"
)

func TestImageUploadV2ByFile(t *testing.T) {
	ctx := context.Background()
	fileData, err := os.ReadFile("/path/to/your/image.jpg")
	if err != nil {
		t.Fatal(err)
	}
	req := &kuaishouModel.ImageUploadV2Req{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.Type = 1
	req.UploadType = 1
	req.File = fileData
	req.FileName = "image.jpg"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ImageUploadV2(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestImageUploadV2ByUrl(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.ImageUploadV2Req{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.Type = 1
	req.UploadType = 2
	req.Url = "https://example.com/image.jpg"
	req.Name = "test_image"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ImageUploadV2(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
