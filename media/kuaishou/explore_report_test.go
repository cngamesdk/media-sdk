package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestExploreReport(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.ExploreReportReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.UnitId = 123456
	req.Param = kuaishouModel.ExploreReportParam{
		Id:               1,
		ExploreType:      1,
		ExploreStartTime: 1636450616000,
		ExploreEndTime:   1636536616000,
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ExploreReport(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
