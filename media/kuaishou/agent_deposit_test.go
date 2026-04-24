package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentDepositListSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentDepositListReq{}
	req.AccessToken = "your_access_token"
	req.AgentId = 309
	req.StartTime = 1700000000
	req.EndTime = 1700086400
	req.IsPage = true
	req.PageInfo = &kuaishouModel.AgentDepositPageInfo{
		CurrentPage: 1,
		PageSize:    20,
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentDepositListSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
