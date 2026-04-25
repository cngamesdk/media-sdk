package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentTradeDetailSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentTradeDetailReq{}
	req.AccessToken = "your_access_token"
	req.AgentId = 309
	req.TradeNo = "776975954695352320"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentTradeDetailSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
