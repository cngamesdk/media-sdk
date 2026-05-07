package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestVideoRelateCreatives(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoRelateCreativesReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.PhotoIds = []string{"5228679230762349823"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoRelateCreatives(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range resp.RelatedCreatives {
		fmt.Printf("photo_id=%s creative_count=%d\n", item.PhotoId, item.CreativeCount)
	}
}

func TestVideoRelateCreativesMultiple(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoRelateCreativesReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.PhotoIds = []string{"5228679230762349823", "5251760176597172988"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoRelateCreatives(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range resp.RelatedCreatives {
		fmt.Printf("photo_id=%s creative_count=%d advanced_count=%d smart_count=%d\n",
			item.PhotoId, item.CreativeCount, item.AdvancedCreativeCount, item.SmartCreativeCount)
		for _, c := range item.Creatives {
			fmt.Printf("  creative_id=%d creative_name=%s\n", c.CreativeId, c.CreativeName)
		}
	}
}

func TestVideoRelateCreativesValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 photo_ids，预期返回校验错误
	req := &kuaishouModel.VideoRelateCreativesReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	_, err := adapter.VideoRelateCreatives(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty photo_ids")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.VideoRelateCreativesReq{}
	req2.AccessToken = "your_access_token"
	req2.PhotoIds = []string{"5228679230762349823"}
	_, err2 := adapter.VideoRelateCreatives(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)
}
