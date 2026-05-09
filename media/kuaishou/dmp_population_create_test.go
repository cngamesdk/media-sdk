package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestDmpPopulationCreate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.DmpPopulationCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.OrientationName = "test-imei-md5"
	req.Type = 3 // IMEI_MD5
	req.FilePaths = []string{"20000681_3_99cf1fd0aa868613039ff19be7698e71.txt"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpPopulationCreate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("orientation_id=%d orientation_name=%s type=%d population_type=%d status=%d create_time=%d record_size=%d advertiser_id=%d failed_file_paths=%v\n",
		resp.OrientationId, resp.OrientationName, resp.Type, resp.PopulationType, resp.Status, resp.CreateTime, resp.RecordSize, resp.AdvertiserId, resp.FailedFilePaths)
}

func TestDmpPopulationCreateWithGidMatch(t *testing.T) {
	ctx := context.Background()
	gidMatch := true
	req := &kuaishouModel.DmpPopulationCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.OrientationName = "test-oaid-gid"
	req.Type = 7 // OAID
	req.FilePaths = []string{
		"20000681_7_aaa111.txt",
		"20000681_7_bbb222.txt",
	}
	req.GidMatch = &gidMatch
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpPopulationCreate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("orientation_id=%d status=%d record_size=%d\n", resp.OrientationId, resp.Status, resp.RecordSize)
}

func TestDmpPopulationCreateValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.DmpPopulationCreateReq{}
	req.AdvertiserId = 20000681
	req.OrientationName = "test"
	req.Type = 3
	req.FilePaths = []string{"path.txt"}
	_, err := adapter.DmpPopulationCreate(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.DmpPopulationCreateReq{}
	req2.AccessToken = "your_access_token"
	req2.OrientationName = "test"
	req2.Type = 3
	req2.FilePaths = []string{"path.txt"}
	_, err2 := adapter.DmpPopulationCreate(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 orientation_name，预期返回校验错误
	req3 := &kuaishouModel.DmpPopulationCreateReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 20000681
	req3.Type = 3
	req3.FilePaths = []string{"path.txt"}
	_, err3 := adapter.DmpPopulationCreate(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty orientation_name")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// orientation_name 超过50字符，预期返回校验错误
	req4 := &kuaishouModel.DmpPopulationCreateReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 20000681
	req4.OrientationName = "这个名称超过了五十个字符限制这个名称超过了五十个字符限制这个名称超过了五十个字符限制这个名称超过了五十个字符限制"
	req4.Type = 3
	req4.FilePaths = []string{"path.txt"}
	_, err4 := adapter.DmpPopulationCreate(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for orientation_name exceeding 50 characters")
	}
	fmt.Printf("got expected error: %v\n", err4)

	// 缺少 type，预期返回校验错误
	req5 := &kuaishouModel.DmpPopulationCreateReq{}
	req5.AccessToken = "your_access_token"
	req5.AdvertiserId = 20000681
	req5.OrientationName = "test"
	req5.FilePaths = []string{"path.txt"}
	_, err5 := adapter.DmpPopulationCreate(ctx, req5)
	if err5 == nil {
		t.Fatal("expected validation error for empty type")
	}
	fmt.Printf("got expected error: %v\n", err5)

	// 缺少 file_paths，预期返回校验错误
	req6 := &kuaishouModel.DmpPopulationCreateReq{}
	req6.AccessToken = "your_access_token"
	req6.AdvertiserId = 20000681
	req6.OrientationName = "test"
	req6.Type = 3
	_, err6 := adapter.DmpPopulationCreate(ctx, req6)
	if err6 == nil {
		t.Fatal("expected validation error for empty file_paths")
	}
	fmt.Printf("got expected error: %v\n", err6)

	// file_paths 超过10个，预期返回校验错误
	req7 := &kuaishouModel.DmpPopulationCreateReq{}
	req7.AccessToken = "your_access_token"
	req7.AdvertiserId = 20000681
	req7.OrientationName = "test"
	req7.Type = 3
	req7.FilePaths = []string{"1.txt", "2.txt", "3.txt", "4.txt", "5.txt", "6.txt", "7.txt", "8.txt", "9.txt", "10.txt", "11.txt"}
	_, err7 := adapter.DmpPopulationCreate(ctx, req7)
	if err7 == nil {
		t.Fatal("expected validation error for file_paths exceeding 10 items")
	}
	fmt.Printf("got expected error: %v\n", err7)
}
