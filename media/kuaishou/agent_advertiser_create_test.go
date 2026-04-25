package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentAdvertiserCreateSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentAdvertiserCreateReq{}
	req.AccessToken = "your_access_token"
	req.CorporationName = "测试企业"
	req.AdvertiserUserId = 1
	req.AgentId = 309
	req.SmsCode = "123456"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentAdvertiserCreateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
