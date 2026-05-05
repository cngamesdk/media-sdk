package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestUnitReport(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.UnitReportReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.StartDate = "2021-12-01"
	req.EndDate = "2021-12-01"
	req.Page = 1
	req.PageSize = 20
	req.TemporalGranularity = "DAILY"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.UnitReport(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
