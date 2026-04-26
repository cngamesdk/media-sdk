package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentDeliveryCertSubmitSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentDeliveryCertSubmitReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.AgentId = 309
	req.CertList = []kuaishouModel.CertParam{
		{
			CertId:       0,
			CertCategory: 5, // 投放资质
			FileToken:    "your_file_token",
			ExpireTime:   "", // 永久有效
		},
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentDeliveryCertSubmitSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
