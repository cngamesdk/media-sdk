package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentAdvertiserCertGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentAdvertiserCertGetReq{}
	req.AccessToken = "your_access_token"
	req.AgentId = 309
	req.AdvertiserId = 20000800
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentAdvertiserCertGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
