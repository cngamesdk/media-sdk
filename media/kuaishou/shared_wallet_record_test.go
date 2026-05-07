package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestSharedWalletRecord(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.SharedWalletRecordReq{}
	req.AccessToken = "your_access_token"
	req.AgentId = 10000100
	req.WalletId = "100001"
	req.StartTradeTime = 1690000000000
	req.EndTradeTime = 1690086400000
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.SharedWalletRecord(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
