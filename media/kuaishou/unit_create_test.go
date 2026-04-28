package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestUnitCreate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.UnitCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.CampaignId = 2342843
	req.UnitName = "测试广告组"
	req.BidType = 10 // OCPM
	req.SceneId = []string{"1"}
	req.UnitType = 4 // 常规自定义创意
	req.BeginTime = "2026-05-01"
	req.EndTime = "2026-05-31"
	req.OcpxActionType = 2
	req.CpaBid = 20000 // 2元，单位：分
	req.Target = &kuaishouModel.UnitCreateTarget{
		Gender:      0, // 不限
		Network:     0, // 不限
		AgesRangeV2: []string{"18", "24", "31"},
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.UnitCreate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
