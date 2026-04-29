package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestCreativeReviewDetail(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.CreativeReviewDetailReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 7869843
	req.CreativeMold = 1 // 自定义创意
	req.Ids = []int64{38502062490, 38501947560, 38502062459}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.CreativeReviewDetail(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
