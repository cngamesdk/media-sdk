package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestImageRelateCreatives(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.ImageRelateCreativesReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.ImageToken = "marketd604af10412c48d6b2e44e0ed1b9b336.jpg"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ImageRelateCreatives(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestImageRelateCreativesCount(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.ImageRelateCreativesReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.ImageToken = "marketd604af10412c48d6b2e44e0ed1b9b336.jpg"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ImageRelateCreatives(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("image_token=%s creative_count=%d\n", resp.ImageToken, resp.CreativeCount)
	for _, c := range resp.Creatives {
		fmt.Printf("  creative_id=%d creative_name=%s\n", c.CreativeId, c.CreativeName)
	}
}

func TestImageRelateCreativesValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 image_token，预期返回校验错误
	req := &kuaishouModel.ImageRelateCreativesReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	_, err := adapter.ImageRelateCreatives(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty image_token")
	}
	fmt.Printf("got expected error: %v\n", err)
}
