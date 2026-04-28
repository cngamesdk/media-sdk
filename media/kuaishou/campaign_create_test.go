package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestCampaignCreate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.CampaignCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.CampaignName = "测试广告计划"
	req.Type = 5 // 线索收集
	req.PutStatus = 1
	req.DayBudget = 50000000 // 500元，单位：厘
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.CampaignCreate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
