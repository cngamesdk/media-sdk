package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"os"
	"testing"
)

func TestAppPicUploadIcon(t *testing.T) {
	ctx := context.Background()
	fileData, err := os.ReadFile("/path/to/your/icon.png")
	if err != nil {
		t.Fatal(err)
	}
	req := &kuaishouModel.AppPicUploadReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.File = fileData
	req.FileName = "icon.png"
	req.Type = 1 // 应用图标(450x450,<1MB,PNG/JPG/JPEG)
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppPicUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("url=%s\n", resp.Url)
}

func TestAppPicUploadAppImage(t *testing.T) {
	ctx := context.Background()
	fileData, err := os.ReadFile("/path/to/your/app_image.png")
	if err != nil {
		t.Fatal(err)
	}
	req := &kuaishouModel.AppPicUploadReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.File = fileData
	req.FileName = "app_image.png"
	req.Type = 2 // 应用图片(9:20,≥720x1280,<2MB,PNG/JPG/JPEG)
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppPicUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("url=%s\n", resp.Url)
}

func TestAppPicUploadValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppPicUploadReq{}
	req.AdvertiserId = 20000681
	req.File = []byte("mock")
	req.FileName = "icon.png"
	req.Type = 1
	_, err := adapter.AppPicUpload(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppPicUploadReq{}
	req2.AccessToken = "your_access_token"
	req2.File = []byte("mock")
	req2.FileName = "icon.png"
	req2.Type = 1
	_, err2 := adapter.AppPicUpload(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 file，预期返回校验错误
	req3 := &kuaishouModel.AppPicUploadReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 20000681
	req3.FileName = "icon.png"
	req3.Type = 1
	_, err3 := adapter.AppPicUpload(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty file")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// 缺少 file_name，预期返回校验错误
	req4 := &kuaishouModel.AppPicUploadReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 20000681
	req4.File = []byte("mock")
	req4.Type = 1
	_, err4 := adapter.AppPicUpload(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for empty file_name")
	}
	fmt.Printf("got expected error: %v\n", err4)

	// type 非法值，预期返回校验错误
	req5 := &kuaishouModel.AppPicUploadReq{}
	req5.AccessToken = "your_access_token"
	req5.AdvertiserId = 20000681
	req5.File = []byte("mock")
	req5.FileName = "icon.png"
	req5.Type = 3 // 非法值
	_, err5 := adapter.AppPicUpload(ctx, req5)
	if err5 == nil {
		t.Fatal("expected validation error for invalid type")
	}
	fmt.Printf("got expected error: %v\n", err5)
}
