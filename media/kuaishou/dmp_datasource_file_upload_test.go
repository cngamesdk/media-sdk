package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"os"
	"testing"
)

func TestDmpDatasourceFileUpload(t *testing.T) {
	ctx := context.Background()
	fileData, err := os.ReadFile("/path/to/your/data.csv")
	if err != nil {
		t.Fatal(err)
	}
	req := &kuaishouModel.DmpDatasourceFileUploadReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.MatchType = "IMEI_MD5"
	req.SchemaType = "DS" // 电商行业模板
	req.File = fileData
	req.FileName = "data.csv"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpDatasourceFileUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("account_id=%d match_type=%s file_key=%s file_size=%d upload_file_type=%s record_size=%d match_size=%d schema_type=%d md5=%s\n",
		resp.AccountId, resp.MatchType, resp.FileKey, resp.FileSize,
		resp.UploadFileType, resp.RecordSize, resp.MatchSize, resp.SchemaType, resp.Md5)
}

func TestDmpDatasourceFileUploadMobileSha256(t *testing.T) {
	ctx := context.Background()
	fileData, err := os.ReadFile("/path/to/your/mobile_sha256.csv")
	if err != nil {
		t.Fatal(err)
	}
	req := &kuaishouModel.DmpDatasourceFileUploadReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.MatchType = "MOBILE_SHA256"
	req.SchemaType = "QT" // 其他通用模板
	req.File = fileData
	req.FileName = "mobile_sha256.csv"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpDatasourceFileUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("file_key=%s record_size=%d match_size=%d\n", resp.FileKey, resp.RecordSize, resp.MatchSize)
}

func TestDmpDatasourceFileUploadValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.DmpDatasourceFileUploadReq{}
	req.AdvertiserId = 20000681
	req.MatchType = "IMEI_MD5"
	req.SchemaType = "DS"
	req.File = []byte("mock")
	req.FileName = "test.csv"
	_, err := adapter.DmpDatasourceFileUpload(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.DmpDatasourceFileUploadReq{}
	req2.AccessToken = "your_access_token"
	req2.MatchType = "IMEI_MD5"
	req2.SchemaType = "DS"
	req2.File = []byte("mock")
	req2.FileName = "test.csv"
	_, err2 := adapter.DmpDatasourceFileUpload(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// match_type 非法值，预期返回校验错误
	req3 := &kuaishouModel.DmpDatasourceFileUploadReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 20000681
	req3.MatchType = "INVALID"
	req3.SchemaType = "DS"
	req3.File = []byte("mock")
	req3.FileName = "test.csv"
	_, err3 := adapter.DmpDatasourceFileUpload(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for invalid match_type")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// schema_type 非法值，预期返回校验错误
	req4 := &kuaishouModel.DmpDatasourceFileUploadReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 20000681
	req4.MatchType = "IMEI_MD5"
	req4.SchemaType = "INVALID"
	req4.File = []byte("mock")
	req4.FileName = "test.csv"
	_, err4 := adapter.DmpDatasourceFileUpload(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for invalid schema_type")
	}
	fmt.Printf("got expected error: %v\n", err4)

	// 缺少 file，预期返回校验错误
	req5 := &kuaishouModel.DmpDatasourceFileUploadReq{}
	req5.AccessToken = "your_access_token"
	req5.AdvertiserId = 20000681
	req5.MatchType = "IMEI_MD5"
	req5.SchemaType = "DS"
	req5.FileName = "test.csv"
	_, err5 := adapter.DmpDatasourceFileUpload(ctx, req5)
	if err5 == nil {
		t.Fatal("expected validation error for empty file")
	}
	fmt.Printf("got expected error: %v\n", err5)

	// 缺少 file_name，预期返回校验错误
	req6 := &kuaishouModel.DmpDatasourceFileUploadReq{}
	req6.AccessToken = "your_access_token"
	req6.AdvertiserId = 20000681
	req6.MatchType = "IMEI_MD5"
	req6.SchemaType = "DS"
	req6.File = []byte("mock")
	_, err6 := adapter.DmpDatasourceFileUpload(ctx, req6)
	if err6 == nil {
		t.Fatal("expected validation error for empty file_name")
	}
	fmt.Printf("got expected error: %v\n", err6)
}
