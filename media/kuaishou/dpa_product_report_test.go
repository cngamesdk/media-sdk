package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestDpaProductReport(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.DpaProductReportReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.LibraryId = 1
	req.StartDate = "2022-01-01"
	req.EndDate = "2022-04-01"
	req.Page = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DpaProductReport(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
