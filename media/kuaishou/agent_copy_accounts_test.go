package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentCopyAccountsSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentCopyAccountsReq{}
	req.AccessToken = "your_access_token"
	req.AgentId = 309
	req.UcType = "DSP_MAPI"
	req.PageNo = 1
	req.PageSize = 10
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentCopyAccountsSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
