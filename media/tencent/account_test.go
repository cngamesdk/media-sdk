package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 查询组织下广告账户信息
func TestOrganizationAccountRelationGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.OrganizationAccountRelationGetReq{}
	req.AccessToken = "123"
	req.PaginationMode = model.PaginationModeNormal
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OrganizationAccountRelationGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 获取竞价广告账户日预算
func TestAdvertiserDailyBudgetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdvertiserDailyBudgetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Fields = []string{"account_id", "daily_budget"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdvertiserDailyBudgetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 获取广告
func TestAdgroupsGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}
