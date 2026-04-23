package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAdvertiserFundDailyFlowsSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AdvertiserFundDailyFlowsReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.StartDate = "2024-01-01"
	req.EndDate = "2024-01-31"
	req.Page = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AdvertiserFundDailyFlowsSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
