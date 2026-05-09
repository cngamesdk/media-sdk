package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestDmpSystemRecommendTarget(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.DmpSystemRecommendReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.Type = 1 // 定向
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpSystemRecommend(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("industry_target=%d consume_target=%d react_target=%d interest_target=%d negative_exclude=%d\n",
		len(resp.IndustryTarget), len(resp.ConsumeTarget), len(resp.ReactTarget),
		len(resp.InterestTarget), len(resp.NegativeExclude))
	for _, item := range resp.IndustryTarget {
		fmt.Printf("  [industry] orientation_id=%d name=%s type=%d population_type=%d status=%d cover_num=%d src_type=%d\n",
			item.OrientationId, item.OrientationName, item.Type, item.PopulationType, item.Status, item.CoverNum, item.SrcType)
	}
}

func TestDmpSystemRecommendExclude(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.DmpSystemRecommendReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.Type = 2 // 排除
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpSystemRecommend(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("negative_exclude=%d\n", len(resp.NegativeExclude))
	for _, item := range resp.NegativeExclude {
		fmt.Printf("  [negative] orientation_id=%d name=%s can_exclude=%d is_exclude_population=%v\n",
			item.OrientationId, item.OrientationName, item.CanExclude, item.IsExcludePopulation)
	}
}

func TestDmpSystemRecommendWithPopulationSource(t *testing.T) {
	ctx := context.Background()
	source := 1 // 联盟人群覆盖数
	req := &kuaishouModel.DmpSystemRecommendReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.Type = 1
	req.PopulationSource = &source
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpSystemRecommend(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("industry_target=%d consume_target=%d react_target=%d interest_target=%d\n",
		len(resp.IndustryTarget), len(resp.ConsumeTarget), len(resp.ReactTarget), len(resp.InterestTarget))
}

func TestDmpSystemRecommendValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.DmpSystemRecommendReq{}
	req.AdvertiserId = 20000681
	req.Type = 1
	_, err := adapter.DmpSystemRecommend(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.DmpSystemRecommendReq{}
	req2.AccessToken = "your_access_token"
	req2.Type = 1
	_, err2 := adapter.DmpSystemRecommend(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// type 为非法值，预期返回校验错误
	req3 := &kuaishouModel.DmpSystemRecommendReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 20000681
	req3.Type = 3 // 非法值，只允许1或2
	_, err3 := adapter.DmpSystemRecommend(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for invalid type")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
