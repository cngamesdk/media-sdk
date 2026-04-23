package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestClueOptimizeSwitchModSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.ClueOptimizeSwitchModReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.ClueOptimizeSwitchTypes = []kuaishouModel.ClueOptimizeSwitchModItem{
		{ClueOptimizeType: 53, Status: true},
		{ClueOptimizeType: 786, Status: false},
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ClueOptimizeSwitchModSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
