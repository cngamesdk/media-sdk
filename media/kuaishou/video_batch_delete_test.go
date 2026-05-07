package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestVideoBatchDelete(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoBatchDeleteReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.PhotoIds = []string{"5238249419979025844"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoBatchDelete(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("photo_ids=%v\n", resp.PhotoIds)
}

func TestVideoBatchDeleteMultiple(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoBatchDeleteReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.PhotoIds = []string{"5238249419979025844", "5238249419979025845", "5238249419979025846"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoBatchDelete(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestVideoBatchDeleteValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 photo_ids，预期返回校验错误
	req := &kuaishouModel.VideoBatchDeleteReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	_, err := adapter.VideoBatchDelete(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty photo_ids")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.VideoBatchDeleteReq{}
	req2.AccessToken = "your_access_token"
	req2.PhotoIds = []string{"5238249419979025844"}
	_, err2 := adapter.VideoBatchDelete(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)
}
