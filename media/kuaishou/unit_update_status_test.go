package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestUnitUpdateStatus(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.UnitUpdateStatusReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.UnitId = 2960188
	req.PutStatus = 2 // 暂停
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.UnitUpdateStatus(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
