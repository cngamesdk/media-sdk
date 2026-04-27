package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentAdvertiserDepositSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentAdvertiserDepositReq{}
	req.AccessToken = "your_access_token"
	req.AgentId = 309
	req.StartTime = 1714492800000
	req.EndTime = 1714579200000
	req.OperationType = 0
	req.AccountSearchType = 0
	req.UcType = ""
	req.PageInfo = &kuaishouModel.AgentAdvertiserDepositPageInfo{
		CurrentPage: 1,
		PageSize:    10,
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentAdvertiserDepositSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
