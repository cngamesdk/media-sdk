package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestVideoList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 139494
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d details=%d\n", resp.TotalCount, len(resp.Details))
}

func TestVideoListByPhotoIds(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 139494
	req.PhotoIds = []string{"5196591116855324734"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range resp.Details {
		fmt.Printf("photo_id=%s photo_name=%s new_status=%d\n", d.PhotoId, d.PhotoName, d.NewStatus)
	}
}

func TestVideoListByDateRange(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 139494
	req.StartDate = "2024-01-01"
	req.EndDate = "2024-01-31"
	req.Page = 1
	req.PageSize = 50
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d\n", resp.TotalCount)
	for _, d := range resp.Details {
		fmt.Printf("%+v\n", d)
	}
}

func TestVideoListWithAllFilters(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 139494
	req.PhotoName = "测试视频"
	req.Signature = "c922a641cc5dc03a497e540996d12198"
	req.OuterLoopNative = 0
	req.UpdateStartDate = "2024-01-01"
	req.UpdateEndDate = "2024-01-31"
	req.Page = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestVideoListValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 advertiser_id，预期返回校验错误
	req := &kuaishouModel.VideoListReq{}
	req.AccessToken = "your_access_token"
	_, err := adapter.VideoList(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err)
}
