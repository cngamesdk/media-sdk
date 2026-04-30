package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAccountAutoInfoUpdate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AccountAutoInfoUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.AccountAutoManage = 1
	req.AutoCampaignNameRule = "系统创建_[日期]_[序号]"
	req.OcpxActionTypeConstraint = []kuaishouModel.OcpxActionTypeConstraint{
		{OcpxActionType: 191, Value: 5},
		{OcpxActionType: 180, Value: 4.2},
		{OcpxActionType: 394, Value: 4.231},
		{OcpxActionType: 53, Value: 3.421},
		{OcpxActionType: 773, Value: 5.423},
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AccountAutoInfoUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
