package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAdvertiserUpdateBudget(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AdvertiserUpdateBudgetReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.DayBudget = 50000000 // 500元，单位：厘
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AdvertiserUpdateBudget(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
