package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentSmsCodeSendSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentSmsCodeSendReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserUserId = 1
	req.AgentId = 309
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentSmsCodeSendSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
