package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentAdvertiserCertSubmitSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentAdvertiserCertSubmitReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.AgentId = 309
	req.IndustryId = 1251
	req.CorporationName = "测试企业"
	req.WebSite = "https://www.example.com"
	req.ProductName = "测试产品"
	req.MarketingContentType = 1 // 1=推广内容链接
	req.LicenceId = "91110000MA00ABCDE1"
	req.LicenceLocation = "北京市-北京市"
	req.CertList = []kuaishouModel.CertParam{
		{
			CertId:       0,
			CertCategory: 1, // 营业执照
			FileToken:    "your_file_token",
			ExpireTime:   "2025-12-31",
		},
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentAdvertiserCertSubmitSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
