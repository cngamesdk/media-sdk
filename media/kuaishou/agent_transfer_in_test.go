package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentTransferInSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentTransferInReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 1
	req.CustomTransferAmount = 1000 // 单位：厘
	req.TransferType = 1            // 1=现金
	req.BizUniqueKey = "your_biz_unique_key"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentTransferInSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
