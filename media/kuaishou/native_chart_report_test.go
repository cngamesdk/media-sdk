package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestNativeChartReport(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.NativeChartReportReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12345678
	req.SearchParam = kuaishouModel.NativeChartReportSearchParam{
		ReportStartDay: 1719744000000,
		ReportEndDay:   1719820800000,
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.NativeChartReport(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
