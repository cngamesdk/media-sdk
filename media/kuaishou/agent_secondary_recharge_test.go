package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentSecondaryRechargeSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentSecondaryRechargeReq{}
	req.AccessToken = "your_access_token"
	req.AgentId = 309
	req.SecondaryAgentId = 310
	req.TransferType = 1 // 1=现金
	req.Amount = 1000    // 单位：厘
	req.BizUniqueKey = "your_biz_unique_key"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentSecondaryRechargeSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
