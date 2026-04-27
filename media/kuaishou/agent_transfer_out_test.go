package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentTransferOutSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentTransferOutReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.Amount = 10000
	req.TransferType = 1
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentTransferOutSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
