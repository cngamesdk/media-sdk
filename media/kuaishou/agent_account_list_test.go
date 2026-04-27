package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentAccountListSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentAccountListReq{}
	req.AccessToken = "your_access_token"
	req.AgentId = 309
	req.Page = 1
	req.PageSize = 10
	req.SelectType = 0
	req.SelectValue = ""
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentAccountListSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
