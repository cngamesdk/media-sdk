package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestVideoAiRecommend(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoAiRecommendReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoAiRecommend(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d details=%d\n", resp.TotalCount, len(resp.Details))
}

func TestVideoAiRecommendWithPaging(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoAiRecommendReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.Page = 1
	req.PageSize = 10
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoAiRecommend(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range resp.Details {
		fmt.Printf("photo_id=%s native_good_type=%d cursor=%d\n", d.PhotoId, d.NativeGoodType, d.Cursor)
	}
}

func TestVideoAiRecommendWithDpa(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoAiRecommendReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.DpaProductId = 123456
	req.Page = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoAiRecommend(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestVideoAiRecommendWithSeries(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoAiRecommendReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.SeriesId = 789012
	req.Page = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoAiRecommend(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestVideoAiRecommendValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 advertiser_id，预期返回校验错误
	req := &kuaishouModel.VideoAiRecommendReq{}
	req.AccessToken = "your_access_token"
	_, err := adapter.VideoAiRecommend(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err)
}
