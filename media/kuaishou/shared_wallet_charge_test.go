package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestSharedWalletCharge(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.SharedWalletChargeReq{}
	req.AccessToken = "your_access_token"
	req.WalletId = "100001"
	req.WalletName = "测试钱包"
	req.AgentId = "10000100"
	req.AgentName = "测试代理商"
	req.TradeNo = "mapi_100001_10000100_001"
	req.TradeType = 16
	req.RelatedFlowNo = "related_001"
	req.BizTradeTime = "1690000000000"
	req.TotalAmount = 10000
	req.UserId = "your_user_id"
	req.Operator = "your_operator"
	req.AppId = 7
	req.SignCompany = "your_company"
	req.FundsOpType = 1
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.SharedWalletCharge(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
