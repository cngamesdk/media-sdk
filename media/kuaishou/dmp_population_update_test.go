package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestDmpPopulationUpdateIncrement(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.DmpPopulationUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.OrientationId = 110750443
	req.OperationType = 1 // 增量更新
	req.Type = 3          // IMEI_MD5
	req.FilePaths = []string{"20000681_3_0d0bdfaaab8df08cf052d0ea41b6e05f.txt"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpPopulationUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("orientation_id=%d orientation_name=%s status=%d record_size=%d failed_file_paths=%v\n",
		resp.OrientationId, resp.OrientationName, resp.Status, resp.RecordSize, resp.FailedFilePaths)
}

func TestDmpPopulationUpdateDecrement(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.DmpPopulationUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.OrientationId = 110750443
	req.OperationType = 3 // 缩量更新
	req.Type = 3          // IMEI_MD5
	req.FilePaths = []string{"20000681_3_aabbcc112233.txt"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpPopulationUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("orientation_id=%d status=%d record_size=%d\n", resp.OrientationId, resp.Status, resp.RecordSize)
}

func TestDmpPopulationUpdateFull(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.DmpPopulationUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.OrientationId = 110750443
	req.OperationType = 4 // 全量更新
	req.Type = 7          // OAID
	req.FilePaths = []string{
		"20000681_7_file1.txt",
		"20000681_7_file2.txt",
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpPopulationUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("orientation_id=%d type=%d population_type=%d status=%d advertiser_id=%d create_time=%d\n",
		resp.OrientationId, resp.Type, resp.PopulationType, resp.Status, resp.AdvertiserId, resp.CreateTime)
}

func TestDmpPopulationUpdateValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.DmpPopulationUpdateReq{}
	req.AdvertiserId = 20000681
	req.OrientationId = 110750443
	req.OperationType = 1
	req.Type = 3
	req.FilePaths = []string{"path.txt"}
	_, err := adapter.DmpPopulationUpdate(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.DmpPopulationUpdateReq{}
	req2.AccessToken = "your_access_token"
	req2.OrientationId = 110750443
	req2.OperationType = 1
	req2.Type = 3
	req2.FilePaths = []string{"path.txt"}
	_, err2 := adapter.DmpPopulationUpdate(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 orientation_id，预期返回校验错误
	req3 := &kuaishouModel.DmpPopulationUpdateReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 20000681
	req3.OperationType = 1
	req3.Type = 3
	req3.FilePaths = []string{"path.txt"}
	_, err3 := adapter.DmpPopulationUpdate(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty orientation_id")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// operation_type 非法值，预期返回校验错误
	req4 := &kuaishouModel.DmpPopulationUpdateReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 20000681
	req4.OrientationId = 110750443
	req4.OperationType = 2 // 非法值
	req4.Type = 3
	req4.FilePaths = []string{"path.txt"}
	_, err4 := adapter.DmpPopulationUpdate(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for invalid operation_type")
	}
	fmt.Printf("got expected error: %v\n", err4)

	// 缺少 type，预期返回校验错误
	req5 := &kuaishouModel.DmpPopulationUpdateReq{}
	req5.AccessToken = "your_access_token"
	req5.AdvertiserId = 20000681
	req5.OrientationId = 110750443
	req5.OperationType = 1
	req5.FilePaths = []string{"path.txt"}
	_, err5 := adapter.DmpPopulationUpdate(ctx, req5)
	if err5 == nil {
		t.Fatal("expected validation error for empty type")
	}
	fmt.Printf("got expected error: %v\n", err5)

	// 缺少 file_paths，预期返回校验错误
	req6 := &kuaishouModel.DmpPopulationUpdateReq{}
	req6.AccessToken = "your_access_token"
	req6.AdvertiserId = 20000681
	req6.OrientationId = 110750443
	req6.OperationType = 1
	req6.Type = 3
	_, err6 := adapter.DmpPopulationUpdate(ctx, req6)
	if err6 == nil {
		t.Fatal("expected validation error for empty file_paths")
	}
	fmt.Printf("got expected error: %v\n", err6)

	// file_paths 超过10个，预期返回校验错误
	req7 := &kuaishouModel.DmpPopulationUpdateReq{}
	req7.AccessToken = "your_access_token"
	req7.AdvertiserId = 20000681
	req7.OrientationId = 110750443
	req7.OperationType = 1
	req7.Type = 3
	req7.FilePaths = []string{"1.txt", "2.txt", "3.txt", "4.txt", "5.txt", "6.txt", "7.txt", "8.txt", "9.txt", "10.txt", "11.txt"}
	_, err7 := adapter.DmpPopulationUpdate(ctx, req7)
	if err7 == nil {
		t.Fatal("expected validation error for file_paths exceeding 10 items")
	}
	fmt.Printf("got expected error: %v\n", err7)
}
