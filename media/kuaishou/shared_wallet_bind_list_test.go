package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestSharedWalletBindList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.SharedWalletBindListReq{}
	req.AccessToken = "your_access_token"
	req.AgentId = 10000100
	req.WalletId = 100001
	req.PageNum = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.SharedWalletBindList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
