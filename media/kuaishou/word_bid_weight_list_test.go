package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestWordBidWeightList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.WordBidWeightListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 139494
	req.PageInfo = kuaishouModel.WordBidWeightPageInfo{
		CurrentPage: 1,
		PageSize:    10,
	}
	req.FilterParam = &kuaishouModel.WordBidWeightFilterParam{
		UnitId: 3891211196,
		Word:   "快手科技",
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.WordBidWeightList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
