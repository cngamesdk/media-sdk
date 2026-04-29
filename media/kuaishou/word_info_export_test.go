package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestWordInfoExport(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.WordInfoExportReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 90000344
	req.StartTime = 1698825600000 // 2023-11-01 00:00:00
	req.EndTime = 1701421200000   // 2023-12-01 00:00:00
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.WordInfoExport(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
