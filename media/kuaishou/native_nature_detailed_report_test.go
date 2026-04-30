package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestNativeNatureDetailedReport(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.NativeNatureDetailedReportReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12345678
	req.PageInfo = kuaishouModel.NatureDetailedPageInfo{
		CurrentPage: 1,
		PageSize:    10,
	}
	req.SearchParam = kuaishouModel.NatureDetailedSearchParam{
		ViewType:       4,
		ReportStartDay: 1719744000000,
		ReportEndDay:   1719820800000,
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.NativeNatureDetailedReport(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
