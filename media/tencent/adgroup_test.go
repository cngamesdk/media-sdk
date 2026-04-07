package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

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

// 创建广告
func TestAdgroupsAddSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupName = "test"
	req.MarketingGoal = model.MarketingGoalUserGrowth
	req.MarketingCarrierType = model.MarketingCarrierTypeAppAndroid
	req.BeginDate = "2026-03-28"
	req.EndDate = "2026-03-28"
	req.BidAmount = 1
	req.TimeSeries = "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
	req.SiteSet = []string{model.SiteSetChannels}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 删除广告
func TestAdgroupsDeleteSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupsDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 123
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupsDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 更新广告
func TestAdgroupsUpdateSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupsUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 123
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupsUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// ========== 批量修改广告日限额测试用例 ==========

// TestAdgroupsUpdateDailyBudgetSingleSelf 测试修改单个广告日限额
func TestAdgroupsUpdateDailyBudgetSingleSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupsUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateDailyBudgetSpec = []*model.UpdateDailyBudgetSpec{
		{AdgroupID: 13397328752, DailyBudget: 5000},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupsUpdateDailyBudgetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdgroupsUpdateDailyBudgetMultipleSelf 测试批量修改多个广告日限额
func TestAdgroupsUpdateDailyBudgetMultipleSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupsUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateDailyBudgetSpec = []*model.UpdateDailyBudgetSpec{
		{AdgroupID: 111, DailyBudget: 5000},
		{AdgroupID: 222, DailyBudget: 100000},
		{AdgroupID: 333, DailyBudget: 400000000},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupsUpdateDailyBudgetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdgroupsUpdateDailyBudgetUnlimitedSelf 测试设置日限额为不限（0）
func TestAdgroupsUpdateDailyBudgetUnlimitedSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupsUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateDailyBudgetSpec = []*model.UpdateDailyBudgetSpec{
		{AdgroupID: 13397328752, DailyBudget: 0},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupsUpdateDailyBudgetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 批量修改广告日限额验证测试用例 ==========

// TestAdgroupsUpdateDailyBudgetValidateMissingAccountIDSelf 测试缺少account_id
func TestAdgroupsUpdateDailyBudgetValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.AdgroupsUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.UpdateDailyBudgetSpec = []*model.UpdateDailyBudgetSpec{
		{AdgroupID: 111, DailyBudget: 5000},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateDailyBudgetValidateEmptySpecSelf 测试spec为空
func TestAdgroupsUpdateDailyBudgetValidateEmptySpecSelf(t *testing.T) {
	req := &model.AdgroupsUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateDailyBudgetSpec = []*model.UpdateDailyBudgetSpec{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：update_daily_budget_spec至少包含1个条件")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateDailyBudgetValidateMissingAdgroupIDSelf 测试spec中缺少adgroup_id
func TestAdgroupsUpdateDailyBudgetValidateMissingAdgroupIDSelf(t *testing.T) {
	req := &model.AdgroupsUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateDailyBudgetSpec = []*model.UpdateDailyBudgetSpec{
		{DailyBudget: 5000},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：adgroup_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateDailyBudgetValidateBudgetTooLowSelf 测试日限额低于最小值
func TestAdgroupsUpdateDailyBudgetValidateBudgetTooLowSelf(t *testing.T) {
	req := &model.AdgroupsUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateDailyBudgetSpec = []*model.UpdateDailyBudgetSpec{
		{AdgroupID: 111, DailyBudget: 100},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：daily_budget低于5000分")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateDailyBudgetValidateBudgetTooHighSelf 测试日限额超过最大值
func TestAdgroupsUpdateDailyBudgetValidateBudgetTooHighSelf(t *testing.T) {
	req := &model.AdgroupsUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateDailyBudgetSpec = []*model.UpdateDailyBudgetSpec{
		{AdgroupID: 111, DailyBudget: 400000001},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：daily_budget超过400000000分")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateDailyBudgetValidateDuplicateAdgroupIDSelf 测试adgroup_id重复
func TestAdgroupsUpdateDailyBudgetValidateDuplicateAdgroupIDSelf(t *testing.T) {
	req := &model.AdgroupsUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateDailyBudgetSpec = []*model.UpdateDailyBudgetSpec{
		{AdgroupID: 111, DailyBudget: 5000},
		{AdgroupID: 111, DailyBudget: 10000},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：adgroup_id不允许重复")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateDailyBudgetValidateNilSpecItemSelf 测试spec含nil项
func TestAdgroupsUpdateDailyBudgetValidateNilSpecItemSelf(t *testing.T) {
	req := &model.AdgroupsUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateDailyBudgetSpec = []*model.UpdateDailyBudgetSpec{
		{AdgroupID: 111, DailyBudget: 5000},
		nil,
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：spec项不能为nil")
	}
	fmt.Printf("验证错误: %v\n", err)
}
