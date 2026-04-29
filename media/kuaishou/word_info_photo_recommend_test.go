package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestWordInfoPhotoRecommend(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.WordInfoPhotoRecommendReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 90000344
	req.PhotoIds = []string{"5189272782789606404"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.WordInfoPhotoRecommend(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
