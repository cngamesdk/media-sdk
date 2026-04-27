package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentCheckAndCopySelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentCheckAndCopyReq{}
	req.AccessToken = "your_access_token"
	req.AgentId = 309
	req.AccountList = []int64{20000800}
	req.UcType = "DSP_MAPI"
	req.CopyNumber = 1
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentCheckAndCopySelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
