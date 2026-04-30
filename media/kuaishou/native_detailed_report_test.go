package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestNativeDetailedReport(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.NativeDetailedReportReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12345678
	req.PageInfo = kuaishouModel.NatureDetailedPageInfo{
		CurrentPage: 1,
		PageSize:    10,
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.NativeDetailedReport(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
