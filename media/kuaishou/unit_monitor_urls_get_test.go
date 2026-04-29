package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestUnitMonitorUrlsGet(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.UnitMonitorUrlsGetReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.UnitIds = []int64{123124, 325344}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.UnitMonitorUrlsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
