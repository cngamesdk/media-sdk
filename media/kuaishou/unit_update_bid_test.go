package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestUnitUpdateBid(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.UnitUpdateBidReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.UnitId = 2960079
	req.Bid = 8930 // 0.893元，单位：厘
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.UnitUpdateBid(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
