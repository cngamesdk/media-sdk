package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentAccountCanCopyQuerySelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentAccountCanCopyQueryReq{}
	req.AccessToken = "your_access_token"
	req.AgentId = 309
	req.AccountIdList = []int64{20000800, 20000801}
	req.UcType = "DSP_MAPI"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentAccountCanCopyQuerySelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
