package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestSharedWalletBalance(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.SharedWalletBalanceReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.AgentId = 10000100
	req.WalletIds = []int64{100001, 100002}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.SharedWalletBalance(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
