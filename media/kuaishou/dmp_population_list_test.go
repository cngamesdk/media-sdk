package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestDmpPopulationList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.DmpPopulationListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpPopulationList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d details=%d\n", resp.TotalCount, len(resp.Details))
	for _, item := range resp.Details {
		fmt.Printf("  orientation_id=%d name=%s type=%d population_type=%d src_type=%d status=%d record_size=%d match_size=%d cover_num=%d create_time=%d verify_time=%d third_platform_code=%d third_platform_name=%s\n",
			item.OrientationId, item.OrientationName, item.Type, item.PopulationType, item.SrcType,
			item.Status, item.RecordSize, item.MatchSize, item.CoverNum, item.CreateTime, item.VerifyTime,
			item.ThirdPlatformCode, item.ThirdPlatformName)
	}
}

func TestDmpPopulationListByStatus(t *testing.T) {
	ctx := context.Background()
	status := 4 // 已上线
	req := &kuaishouModel.DmpPopulationListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.Status = &status
	req.Page = 1
	req.PageSize = 50
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpPopulationList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d details=%d\n", resp.TotalCount, len(resp.Details))
}

func TestDmpPopulationListByIds(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.DmpPopulationListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.OrientationIds = []int64{110750443, 110839912}
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpPopulationList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d\n", resp.TotalCount)
	for _, item := range resp.Details {
		fmt.Printf("  orientation_id=%d name=%s status=%d\n", item.OrientationId, item.OrientationName, item.Status)
	}
}

func TestDmpPopulationListValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.DmpPopulationListReq{}
	req.AdvertiserId = 20000681
	_, err := adapter.DmpPopulationList(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.DmpPopulationListReq{}
	req2.AccessToken = "your_access_token"
	_, err2 := adapter.DmpPopulationList(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// page_size 超过500，预期返回校验错误
	req3 := &kuaishouModel.DmpPopulationListReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 20000681
	req3.PageSize = 501
	_, err3 := adapter.DmpPopulationList(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for page_size exceeding 500")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// orientation_ids 数量不小于 page_size，预期返回校验错误
	req4 := &kuaishouModel.DmpPopulationListReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 20000681
	req4.PageSize = 2
	req4.OrientationIds = []int64{110750443, 110839912}
	_, err4 := adapter.DmpPopulationList(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for orientation_ids count >= page_size")
	}
	fmt.Printf("got expected error: %v\n", err4)
}
