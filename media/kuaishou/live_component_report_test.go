package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestLiveComponentReport(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.LiveComponentReportReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.StartDate = "2022-05-24"
	req.EndDate = "2022-05-27"
	req.Page = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.LiveComponentReport(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
