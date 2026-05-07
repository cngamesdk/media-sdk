package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestVideoTagDelete(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoTagDeleteReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.PhotoTag = []string{"标签1"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoTagDelete(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("photo_tag=%v\n", resp.PhotoTag)
}

func TestVideoTagDeleteMultiple(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoTagDeleteReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.PhotoTag = []string{"标签1", "标签2"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoTagDelete(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestVideoTagDeleteValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 photo_tag，预期返回校验错误
	req := &kuaishouModel.VideoTagDeleteReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	_, err := adapter.VideoTagDelete(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty photo_tag")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.VideoTagDeleteReq{}
	req2.AccessToken = "your_access_token"
	req2.PhotoTag = []string{"标签1"}
	_, err2 := adapter.VideoTagDelete(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)
}
