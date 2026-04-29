package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestUnitMonitorUrlsUpdate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.UnitMonitorUrlsUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 13949412321
	req.UnitMonitorUrls = []kuaishouModel.UnitMonitorUrlUpdateItem{
		{
			UnitId:       12321312,
			ClickUrl:     "https://www.baidu.com",
			LiveTrackUrl: "https://www.baidu.com",
		},
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.UnitMonitorUrlsUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
