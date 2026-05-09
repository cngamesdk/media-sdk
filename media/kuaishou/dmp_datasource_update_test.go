package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestDmpDatasourceUpdateFull(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.DmpDatasourceUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.OperationType = 0 // 全量更新
	req.DataSourceId = 100001
	req.FileKeys = []string{"20000681_imei_md5_new_abc123.csv"}
	req.MatchType = "IMEI_MD5"
	req.SchemaType = "DS"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpDatasourceUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("data_source_id=%d data_source_name=%s account_id=%d match_type=%d schema_type=%s all_file_size=%d create_time=%d calcu_status=%d error_message=%s\n",
		resp.DataSourceId, resp.DataSourceName, resp.AccountId, resp.MatchType,
		resp.SchemaType, resp.AllFileSize, resp.CreateTime, resp.CalcuStatus, resp.ErrorMessage)
}

func TestDmpDatasourceUpdateIncrement(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.DmpDatasourceUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.OperationType = 1 // 增量更新
	req.DataSourceId = 100001
	req.FileKeys = []string{
		"20000681_oaid_md5_file1.csv",
		"20000681_oaid_md5_file2.csv",
	}
	req.MatchType = "OAID_MD5"
	req.SchemaType = "YX"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpDatasourceUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("data_source_id=%d calcu_status=%d\n", resp.DataSourceId, resp.CalcuStatus)
}

func TestDmpDatasourceUpdateValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.DmpDatasourceUpdateReq{}
	req.AdvertiserId = 20000681
	req.OperationType = 0
	req.DataSourceId = 100001
	req.FileKeys = []string{"file.csv"}
	req.MatchType = "IMEI_MD5"
	req.SchemaType = "DS"
	_, err := adapter.DmpDatasourceUpdate(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.DmpDatasourceUpdateReq{}
	req2.AccessToken = "your_access_token"
	req2.OperationType = 0
	req2.DataSourceId = 100001
	req2.FileKeys = []string{"file.csv"}
	req2.MatchType = "IMEI_MD5"
	req2.SchemaType = "DS"
	_, err2 := adapter.DmpDatasourceUpdate(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// operation_type 非法值，预期返回校验错误
	req3 := &kuaishouModel.DmpDatasourceUpdateReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 20000681
	req3.OperationType = 2 // 非法值
	req3.DataSourceId = 100001
	req3.FileKeys = []string{"file.csv"}
	req3.MatchType = "IMEI_MD5"
	req3.SchemaType = "DS"
	_, err3 := adapter.DmpDatasourceUpdate(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for invalid operation_type")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// 缺少 data_source_id，预期返回校验错误
	req4 := &kuaishouModel.DmpDatasourceUpdateReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 20000681
	req4.OperationType = 0
	req4.FileKeys = []string{"file.csv"}
	req4.MatchType = "IMEI_MD5"
	req4.SchemaType = "DS"
	_, err4 := adapter.DmpDatasourceUpdate(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for empty data_source_id")
	}
	fmt.Printf("got expected error: %v\n", err4)

	// 缺少 file_keys，预期返回校验错误
	req5 := &kuaishouModel.DmpDatasourceUpdateReq{}
	req5.AccessToken = "your_access_token"
	req5.AdvertiserId = 20000681
	req5.OperationType = 0
	req5.DataSourceId = 100001
	req5.MatchType = "IMEI_MD5"
	req5.SchemaType = "DS"
	_, err5 := adapter.DmpDatasourceUpdate(ctx, req5)
	if err5 == nil {
		t.Fatal("expected validation error for empty file_keys")
	}
	fmt.Printf("got expected error: %v\n", err5)

	// file_keys 超过10个，预期返回校验错误
	req6 := &kuaishouModel.DmpDatasourceUpdateReq{}
	req6.AccessToken = "your_access_token"
	req6.AdvertiserId = 20000681
	req6.OperationType = 0
	req6.DataSourceId = 100001
	req6.FileKeys = []string{"1.csv", "2.csv", "3.csv", "4.csv", "5.csv", "6.csv", "7.csv", "8.csv", "9.csv", "10.csv", "11.csv"}
	req6.MatchType = "IMEI_MD5"
	req6.SchemaType = "DS"
	_, err6 := adapter.DmpDatasourceUpdate(ctx, req6)
	if err6 == nil {
		t.Fatal("expected validation error for file_keys exceeding 10 items")
	}
	fmt.Printf("got expected error: %v\n", err6)

	// match_type 非法值，预期返回校验错误
	req7 := &kuaishouModel.DmpDatasourceUpdateReq{}
	req7.AccessToken = "your_access_token"
	req7.AdvertiserId = 20000681
	req7.OperationType = 0
	req7.DataSourceId = 100001
	req7.FileKeys = []string{"file.csv"}
	req7.MatchType = "INVALID"
	req7.SchemaType = "DS"
	_, err7 := adapter.DmpDatasourceUpdate(ctx, req7)
	if err7 == nil {
		t.Fatal("expected validation error for invalid match_type")
	}
	fmt.Printf("got expected error: %v\n", err7)

	// schema_type 非法值，预期返回校验错误
	req8 := &kuaishouModel.DmpDatasourceUpdateReq{}
	req8.AccessToken = "your_access_token"
	req8.AdvertiserId = 20000681
	req8.OperationType = 0
	req8.DataSourceId = 100001
	req8.FileKeys = []string{"file.csv"}
	req8.MatchType = "IMEI_MD5"
	req8.SchemaType = "INVALID"
	_, err8 := adapter.DmpDatasourceUpdate(ctx, req8)
	if err8 == nil {
		t.Fatal("expected validation error for invalid schema_type")
	}
	fmt.Printf("got expected error: %v\n", err8)
}
