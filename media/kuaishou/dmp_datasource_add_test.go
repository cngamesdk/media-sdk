package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestDmpDatasourceAdd(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.DmpDatasourceAddReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.DataSourceName = "测试数据源"
	req.MatchType = "IMEI_MD5"
	req.SchemaType = "DS" // 电商行业模板
	req.FileKeys = []string{"20000681_imei_md5_abc123.csv"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpDatasourceAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("data_source_id=%d data_source_name=%s account_id=%d match_type=%d schema_type=%s all_file_size=%d create_time=%d calcu_status=%d error_message=%s\n",
		resp.DataSourceId, resp.DataSourceName, resp.AccountId, resp.MatchType,
		resp.SchemaType, resp.AllFileSize, resp.CreateTime, resp.CalcuStatus, resp.ErrorMessage)
}

func TestDmpDatasourceAddMultiFile(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.DmpDatasourceAddReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.DataSourceName = "OAID游戏数据源"
	req.MatchType = "OAID_MD5"
	req.SchemaType = "YX" // 游戏行业模板
	req.FileKeys = []string{
		"20000681_oaid_md5_file1.csv",
		"20000681_oaid_md5_file2.csv",
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpDatasourceAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("data_source_id=%d calcu_status=%d\n", resp.DataSourceId, resp.CalcuStatus)
}

func TestDmpDatasourceAddValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.DmpDatasourceAddReq{}
	req.AdvertiserId = 20000681
	req.DataSourceName = "test"
	req.MatchType = "IMEI_MD5"
	req.SchemaType = "DS"
	req.FileKeys = []string{"file.csv"}
	_, err := adapter.DmpDatasourceAdd(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.DmpDatasourceAddReq{}
	req2.AccessToken = "your_access_token"
	req2.DataSourceName = "test"
	req2.MatchType = "IMEI_MD5"
	req2.SchemaType = "DS"
	req2.FileKeys = []string{"file.csv"}
	_, err2 := adapter.DmpDatasourceAdd(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 data_source_name，预期返回校验错误
	req3 := &kuaishouModel.DmpDatasourceAddReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 20000681
	req3.MatchType = "IMEI_MD5"
	req3.SchemaType = "DS"
	req3.FileKeys = []string{"file.csv"}
	_, err3 := adapter.DmpDatasourceAdd(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty data_source_name")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// data_source_name 超过20字符，预期返回校验错误
	req4 := &kuaishouModel.DmpDatasourceAddReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 20000681
	req4.DataSourceName = "超过二十个字符的数据源名称超过限制了"
	req4.MatchType = "IMEI_MD5"
	req4.SchemaType = "DS"
	req4.FileKeys = []string{"file.csv"}
	_, err4 := adapter.DmpDatasourceAdd(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for data_source_name exceeding 20 characters")
	}
	fmt.Printf("got expected error: %v\n", err4)

	// match_type 非法值，预期返回校验错误
	req5 := &kuaishouModel.DmpDatasourceAddReq{}
	req5.AccessToken = "your_access_token"
	req5.AdvertiserId = 20000681
	req5.DataSourceName = "test"
	req5.MatchType = "INVALID"
	req5.SchemaType = "DS"
	req5.FileKeys = []string{"file.csv"}
	_, err5 := adapter.DmpDatasourceAdd(ctx, req5)
	if err5 == nil {
		t.Fatal("expected validation error for invalid match_type")
	}
	fmt.Printf("got expected error: %v\n", err5)

	// schema_type 非法值，预期返回校验错误
	req6 := &kuaishouModel.DmpDatasourceAddReq{}
	req6.AccessToken = "your_access_token"
	req6.AdvertiserId = 20000681
	req6.DataSourceName = "test"
	req6.MatchType = "IMEI_MD5"
	req6.SchemaType = "INVALID"
	req6.FileKeys = []string{"file.csv"}
	_, err6 := adapter.DmpDatasourceAdd(ctx, req6)
	if err6 == nil {
		t.Fatal("expected validation error for invalid schema_type")
	}
	fmt.Printf("got expected error: %v\n", err6)

	// 缺少 file_keys，预期返回校验错误
	req7 := &kuaishouModel.DmpDatasourceAddReq{}
	req7.AccessToken = "your_access_token"
	req7.AdvertiserId = 20000681
	req7.DataSourceName = "test"
	req7.MatchType = "IMEI_MD5"
	req7.SchemaType = "DS"
	_, err7 := adapter.DmpDatasourceAdd(ctx, req7)
	if err7 == nil {
		t.Fatal("expected validation error for empty file_keys")
	}
	fmt.Printf("got expected error: %v\n", err7)

	// file_keys 超过10个，预期返回校验错误
	req8 := &kuaishouModel.DmpDatasourceAddReq{}
	req8.AccessToken = "your_access_token"
	req8.AdvertiserId = 20000681
	req8.DataSourceName = "test"
	req8.MatchType = "IMEI_MD5"
	req8.SchemaType = "DS"
	req8.FileKeys = []string{"1.csv", "2.csv", "3.csv", "4.csv", "5.csv", "6.csv", "7.csv", "8.csv", "9.csv", "10.csv", "11.csv"}
	_, err8 := adapter.DmpDatasourceAdd(ctx, req8)
	if err8 == nil {
		t.Fatal("expected validation error for file_keys exceeding 10 items")
	}
	fmt.Printf("got expected error: %v\n", err8)
}
