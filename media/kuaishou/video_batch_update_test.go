package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestVideoBatchUpdateByName(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoBatchUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.PhotoIds = []string{"5238249419979025844"}
	req.PhotoName = "视频名称"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoBatchUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("photo_ids=%v\n", resp.PhotoIds)
}

func TestVideoBatchUpdateByTag(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoBatchUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.PhotoIds = []string{"5238249419979025844", "5238249419979025845"}
	req.PhotoTag = []string{"标签1"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoBatchUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestVideoBatchUpdateWithAll(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoBatchUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.PhotoIds = []string{"5238249419979025844"}
	req.PhotoName = "视频名称"
	req.PhotoTag = []string{"标签1"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoBatchUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestVideoBatchUpdateValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 photo_ids，预期返回校验错误
	req := &kuaishouModel.VideoBatchUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.PhotoName = "视频名称"
	_, err := adapter.VideoBatchUpdate(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty photo_ids")
	}
	fmt.Printf("got expected error: %v\n", err)

	// photo_name 和 photo_tag 均为空，预期返回校验错误
	req2 := &kuaishouModel.VideoBatchUpdateReq{}
	req2.AccessToken = "your_access_token"
	req2.AdvertiserId = 11311124
	req2.PhotoIds = []string{"5238249419979025844"}
	_, err2 := adapter.VideoBatchUpdate(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty photo_name and photo_tag")
	}
	fmt.Printf("got expected error: %v\n", err2)
}
