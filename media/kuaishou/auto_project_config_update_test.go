package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAutoProjectConfigUpdate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AutoProjectConfigUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12345678
	req.AccountAutoManage = 1
	req.AutoManageType = 1
	req.KolUserInfo = &kuaishouModel.KolUserInfoReq{
		KolUserType: 1,
		KolUserId:   1234567890,
	}
	req.OcpxActionTypeConstraint = []kuaishouModel.AutoOcpxConstraintReq{
		{OcpxActionType: 191, Value: 3},
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AutoProjectConfigUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
