package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"os"
	"testing"
)

func TestAppApkUpload(t *testing.T) {
	ctx := context.Background()
	fileData, err := os.ReadFile("/path/to/your/app.apk")
	if err != nil {
		t.Fatal(err)
	}
	req := &kuaishouModel.AppApkUploadReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.File = fileData
	req.FileName = "app.apk"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppApkUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("blob_store_key=%s url=%s test_apk_url=%s\n", resp.BlobStoreKey, resp.Url, resp.TestApkUrl)
}

func TestAppApkUploadWithTestPackage(t *testing.T) {
	ctx := context.Background()
	fileData, err := os.ReadFile("/path/to/your/app.apk")
	if err != nil {
		t.Fatal(err)
	}
	apkType := 1 // 上传母包并生成测试分包下载链接
	req := &kuaishouModel.AppApkUploadReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.File = fileData
	req.FileName = "app.apk"
	req.Type = &apkType
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppApkUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("blob_store_key=%s url=%s test_apk_url=%s\n", resp.BlobStoreKey, resp.Url, resp.TestApkUrl)
}

func TestAppApkUploadValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppApkUploadReq{}
	req.AdvertiserId = 20000681
	req.File = []byte("mock")
	req.FileName = "app.apk"
	_, err := adapter.AppApkUpload(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppApkUploadReq{}
	req2.AccessToken = "your_access_token"
	req2.File = []byte("mock")
	req2.FileName = "app.apk"
	_, err2 := adapter.AppApkUpload(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 file，预期返回校验错误
	req3 := &kuaishouModel.AppApkUploadReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 20000681
	req3.FileName = "app.apk"
	_, err3 := adapter.AppApkUpload(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty file")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// 缺少 file_name，预期返回校验错误
	req4 := &kuaishouModel.AppApkUploadReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 20000681
	req4.File = []byte("mock")
	_, err4 := adapter.AppApkUpload(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for empty file_name")
	}
	fmt.Printf("got expected error: %v\n", err4)
}
