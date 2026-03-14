package toutiao

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/toutiao/model"
	"testing"
)

// 获取账户日预算
func TestAdvertiserBudgetGetSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewToutiaoAdapter(config.DefaultConfig())
	req := &model.AdvertiserBudgetGetReq{}
	req.AccessToken = "test"
	req.AdvertiserIds = []int64{123}
	resp, err := factory.AdvertiserBudgetGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// 更新账户日预算
func TestAdvertiserBudgetUpdateSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewToutiaoAdapter(config.DefaultConfig())
	req := &model.AdvertiserBudgetUpdateReq{}
	req.AccessToken = "test"
	req.AdvertiserId = 123
	req.BudgetMode = model.BudgetModeDay
	req.Budget = 100
	resp, err := factory.AdvertiserBudgetUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}
