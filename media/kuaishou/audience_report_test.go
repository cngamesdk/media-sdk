package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAudienceReport(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AudienceReportReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.StartDate = "2021-12-01"
	req.EndDate = "2021-12-01"
	req.ReportDims = "province"
	req.Page = 1
	req.PageSize = 100
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AudienceReport(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
