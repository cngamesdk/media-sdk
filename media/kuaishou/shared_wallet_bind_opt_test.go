package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestSharedWalletBindOpt(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.SharedWalletBindOptReq{}
	req.AccessToken = "your_access_token"
	req.WalletId = "100001"
	req.AgentId = "10000100"
	req.TradeNo = "mapi_100001_20000681"
	req.AppId = 7
	req.AccountId = []int64{20000681}
	req.AccountOperator = 1
	req.UserId = "your_user_id"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.SharedWalletBindOpt(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
