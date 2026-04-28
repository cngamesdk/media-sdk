package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestCampaignContinuePeriodUpdate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.CampaignContinuePeriodUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.CampaignId = 2960188
	req.ContinuePeriodType = 2 // 开启续投
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.CampaignContinuePeriodUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
