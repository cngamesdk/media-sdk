package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentAccountTransferSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentAccountTransferReq{}
	req.AccessToken = "your_access_token"
	req.FromAccountId = 100001
	req.ToAccountId = 100002
	req.OperatorName = "operator"
	req.TransferType = 1      // 1=现金
	req.TransferAmount = 1000 // 单位：厘
	req.AgentId = 309
	req.BizUniqueKey = "your_biz_unique_key"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentAccountTransferSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
