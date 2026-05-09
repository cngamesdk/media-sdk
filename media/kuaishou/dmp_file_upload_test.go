package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"os"
	"testing"
)

func TestDmpFileUploadTxt(t *testing.T) {
	ctx := context.Background()
	fileData, err := os.ReadFile("/path/to/your/imei_md5.txt")
	if err != nil {
		t.Fatal(err)
	}
	req := &kuaishouModel.DmpFileUploadReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.Type = 3 // IMEI_MD5
	req.File = fileData
	req.FileName = "imei_md5.txt"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpFileUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("advertiser_id=%d match_type=%s file_path=%s file_size=%d upload_file_type=%s record_size=%d type=%d md5=%s\n",
		resp.AdvertiserId, resp.MatchType, resp.FilePath, resp.FileSize, resp.UploadFileType, resp.RecordSize, resp.Type, resp.Md5)
}

func TestDmpFileUploadZip(t *testing.T) {
	ctx := context.Background()
	fileData, err := os.ReadFile("/path/to/your/phones.zip")
	if err != nil {
		t.Fatal(err)
	}
	req := &kuaishouModel.DmpFileUploadReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.Type = 5 // 手机号-MD5
	req.File = fileData
	req.FileName = "phones.zip"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpFileUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("file_path=%s upload_file_type=%s record_size=%d\n", resp.FilePath, resp.UploadFileType, resp.RecordSize)
}

func TestDmpFileUploadValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.DmpFileUploadReq{}
	req.AdvertiserId = 20000681
	req.Type = 3
	req.File = []byte("mock")
	req.FileName = "test.txt"
	_, err := adapter.DmpFileUpload(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.DmpFileUploadReq{}
	req2.AccessToken = "your_access_token"
	req2.Type = 3
	req2.File = []byte("mock")
	req2.FileName = "test.txt"
	_, err2 := adapter.DmpFileUpload(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 type，预期返回校验错误
	req3 := &kuaishouModel.DmpFileUploadReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 20000681
	req3.File = []byte("mock")
	req3.FileName = "test.txt"
	_, err3 := adapter.DmpFileUpload(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty type")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// 缺少 file，预期返回校验错误
	req4 := &kuaishouModel.DmpFileUploadReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 20000681
	req4.Type = 3
	req4.FileName = "test.txt"
	_, err4 := adapter.DmpFileUpload(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for empty file")
	}
	fmt.Printf("got expected error: %v\n", err4)

	// 缺少 file_name，预期返回校验错误
	req5 := &kuaishouModel.DmpFileUploadReq{}
	req5.AccessToken = "your_access_token"
	req5.AdvertiserId = 20000681
	req5.Type = 3
	req5.File = []byte("mock")
	_, err5 := adapter.DmpFileUpload(ctx, req5)
	if err5 == nil {
		t.Fatal("expected validation error for empty file_name")
	}
	fmt.Printf("got expected error: %v\n", err5)
}
