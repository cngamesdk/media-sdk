package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestWordBidWeightCreate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.WordBidWeightCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 139494
	req.Scope = 2
	req.BidWeightInfo = []kuaishouModel.AddBidWeightInfo{
		{
			UnitId: 3891211196,
			Info: []kuaishouModel.WeightInfo{
				{Word: "yiulld04", Weight: 1.1},
				{Word: "yiulld05", Weight: 1.1},
			},
		},
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.WordBidWeightCreate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
