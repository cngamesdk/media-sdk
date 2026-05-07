package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestVideoCursorList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoCursorListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.Cursor = 0
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoCursorList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d details=%d\n", resp.TotalCount, len(resp.Details))
}

func TestVideoCursorListWithLimit(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoCursorListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.Cursor = 0
	req.Limit = 2
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoCursorList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range resp.Details {
		fmt.Printf("cursor=%d photo_id=%s photo_name=%s\n", d.Cursor, d.PhotoId, d.PhotoName)
	}
}

func TestVideoCursorListPaging(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 第一轮：从0开始
	req := &kuaishouModel.VideoCursorListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.Cursor = 0
	req.Limit = 2
	resp, err := adapter.VideoCursorList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("第一轮: total_count=%d count=%d\n", resp.TotalCount, len(resp.Details))

	if len(resp.Details) == 0 {
		return
	}

	// 取最大cursor用于下一轮
	var maxCursor int64
	for _, d := range resp.Details {
		if d.Cursor > maxCursor {
			maxCursor = d.Cursor
		}
	}
	fmt.Printf("最大cursor=%d\n", maxCursor)

	// 第二轮：从上一轮最大cursor继续
	req2 := &kuaishouModel.VideoCursorListReq{}
	req2.AccessToken = "your_access_token"
	req2.AdvertiserId = 11311124
	req2.Cursor = maxCursor
	req2.Limit = 2
	resp2, err := adapter.VideoCursorList(ctx, req2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("第二轮: count=%d\n", len(resp2.Details))
	for _, d := range resp2.Details {
		fmt.Printf("cursor=%d photo_id=%s\n", d.Cursor, d.PhotoId)
	}
}

func TestVideoCursorListValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 advertiser_id，预期返回校验错误
	req := &kuaishouModel.VideoCursorListReq{}
	req.AccessToken = "your_access_token"
	req.Cursor = 0
	_, err := adapter.VideoCursorList(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err)
}
