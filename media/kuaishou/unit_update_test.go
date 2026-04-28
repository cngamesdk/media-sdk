package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestUnitUpdate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.UnitUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.UnitId = 2959088
	req.UnitName = "测试广告组-修改"
	req.BidType = 10       // OCPM
	req.CpaBid = 30000     // 3元，单位：分
	req.DayBudget = 100000 // 1000元，单位：分
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.UnitUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
