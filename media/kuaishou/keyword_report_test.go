package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestKeywordReport(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.KeywordReportReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.StartDate = "2022-02-07"
	req.EndDate = "2022-02-09"
	req.Page = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.KeywordReport(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
