package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentPayTokenSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentPayTokenReq{}
	req.AccessToken = "your_access_token"
	req.AgentId = 309
	req.Remark = "transfer"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentPayTokenSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
