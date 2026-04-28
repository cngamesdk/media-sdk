package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestCampaignUpdate(t *testing.T) {
	ctx := context.Background()
	dayBudget := int64(50000000) // 500元，单位：厘
	req := &kuaishouModel.CampaignUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.CampaignId = 2342843
	req.CampaignName = "信息类_收集线索销售0"
	req.DayBudget = &dayBudget
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.CampaignUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
