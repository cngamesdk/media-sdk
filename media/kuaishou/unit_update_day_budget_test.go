package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestUnitUpdateDayBudget(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.UnitUpdateDayBudgetReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.UnitId = 2960079
	dayBudget := int64(100000) // 10元，单位：厘
	req.DayBudget = &dayBudget
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.UnitUpdateDayBudget(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
