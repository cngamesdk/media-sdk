package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestSharedWalletConsumeQuery(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.SharedWalletConsumeQueryReq{}
	req.AccessToken = "your_access_token"
	req.AgentId = 10000100
	req.WalletId = "100001"
	req.StartConsumeTime = 1690000000000
	req.EndConsumeTime = 1690086400000
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.SharedWalletConsumeQuery(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
