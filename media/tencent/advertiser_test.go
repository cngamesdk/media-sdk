package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ========== 批量修改广告主日限额测试用例 ==========

// TestAdvertiserUpdateDailyBudgetSingleSelf 测试修改单个广告主日限额
func TestAdvertiserUpdateDailyBudgetSingleSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdvertiserUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.UpdateDailyBudgetSpec = []*model.AdvertiserUpdateDailyBudgetSpec{
		{AccountID: 111111, DailyBudget: 1000000},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdvertiserUpdateDailyBudgetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdvertiserUpdateDailyBudgetMultipleSelf 测试批量修改多个广告主日限额
func TestAdvertiserUpdateDailyBudgetMultipleSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdvertiserUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.UpdateDailyBudgetSpec = []*model.AdvertiserUpdateDailyBudgetSpec{
		{AccountID: 111111, DailyBudget: 1000000},
		{AccountID: 222222, DailyBudget: 5000000},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdvertiserUpdateDailyBudgetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdvertiserUpdateDailyBudgetUnlimitedSelf 测试设置广告主日限额为不限（0）
func TestAdvertiserUpdateDailyBudgetUnlimitedSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdvertiserUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.UpdateDailyBudgetSpec = []*model.AdvertiserUpdateDailyBudgetSpec{
		{AccountID: 111111, DailyBudget: 0},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdvertiserUpdateDailyBudgetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdvertiserUpdateDailyBudgetWithUseMinSelf 测试启用自动最小值（use_min_daily_budget=true）
func TestAdvertiserUpdateDailyBudgetWithUseMinSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdvertiserUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.UpdateDailyBudgetSpec = []*model.AdvertiserUpdateDailyBudgetSpec{
		{AccountID: 111111, DailyBudget: 5000, UseMinDailyBudget: true},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdvertiserUpdateDailyBudgetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdvertiserUpdateDailyBudgetMaxBudgetSelf 测试设置最大日限额（40亿分）
func TestAdvertiserUpdateDailyBudgetMaxBudgetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdvertiserUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.UpdateDailyBudgetSpec = []*model.AdvertiserUpdateDailyBudgetSpec{
		{AccountID: 111111, DailyBudget: model.MaxAdvertiserDailyBudget},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdvertiserUpdateDailyBudgetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 验证测试用例 ==========

// TestAdvertiserUpdateDailyBudgetValidateEmptySpecSelf 测试spec为空
func TestAdvertiserUpdateDailyBudgetValidateEmptySpecSelf(t *testing.T) {
	req := &model.AdvertiserUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.UpdateDailyBudgetSpec = []*model.AdvertiserUpdateDailyBudgetSpec{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：update_daily_budget_spec至少包含1个条件")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdvertiserUpdateDailyBudgetValidateNilSpecSelf 测试spec为nil
func TestAdvertiserUpdateDailyBudgetValidateNilSpecSelf(t *testing.T) {
	req := &model.AdvertiserUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：update_daily_budget_spec至少包含1个条件")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdvertiserUpdateDailyBudgetValidateExceedMaxSelf 测试spec超过100条
func TestAdvertiserUpdateDailyBudgetValidateExceedMaxSelf(t *testing.T) {
	req := &model.AdvertiserUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	specs := make([]*model.AdvertiserUpdateDailyBudgetSpec, 101)
	for i := range specs {
		specs[i] = &model.AdvertiserUpdateDailyBudgetSpec{AccountID: int64(i + 1), DailyBudget: 5000}
	}
	req.UpdateDailyBudgetSpec = specs
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：spec超过100条")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdvertiserUpdateDailyBudgetValidateMissingAccountIDSelf 测试spec中缺少account_id
func TestAdvertiserUpdateDailyBudgetValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.AdvertiserUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.UpdateDailyBudgetSpec = []*model.AdvertiserUpdateDailyBudgetSpec{
		{DailyBudget: 5000},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdvertiserUpdateDailyBudgetValidateBudgetTooLowSelf 测试日限额低于最小值（非0）
func TestAdvertiserUpdateDailyBudgetValidateBudgetTooLowSelf(t *testing.T) {
	req := &model.AdvertiserUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.UpdateDailyBudgetSpec = []*model.AdvertiserUpdateDailyBudgetSpec{
		{AccountID: 111111, DailyBudget: 100},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：daily_budget低于5000分")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdvertiserUpdateDailyBudgetValidateBudgetTooHighSelf 测试日限额超过最大值（40亿分）
func TestAdvertiserUpdateDailyBudgetValidateBudgetTooHighSelf(t *testing.T) {
	req := &model.AdvertiserUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.UpdateDailyBudgetSpec = []*model.AdvertiserUpdateDailyBudgetSpec{
		{AccountID: 111111, DailyBudget: 4000000001},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：daily_budget超过4000000000分")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdvertiserUpdateDailyBudgetValidateDuplicateAccountIDSelf 测试account_id重复
func TestAdvertiserUpdateDailyBudgetValidateDuplicateAccountIDSelf(t *testing.T) {
	req := &model.AdvertiserUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.UpdateDailyBudgetSpec = []*model.AdvertiserUpdateDailyBudgetSpec{
		{AccountID: 111111, DailyBudget: 1000000},
		{AccountID: 111111, DailyBudget: 2000000},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id不允许重复")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdvertiserUpdateDailyBudgetValidateNilSpecItemSelf 测试spec含nil项
func TestAdvertiserUpdateDailyBudgetValidateNilSpecItemSelf(t *testing.T) {
	req := &model.AdvertiserUpdateDailyBudgetReq{}
	req.AccessToken = "123"
	req.UpdateDailyBudgetSpec = []*model.AdvertiserUpdateDailyBudgetSpec{
		{AccountID: 111111, DailyBudget: 1000000},
		nil,
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：spec项不能为nil")
	}
	fmt.Printf("验证错误: %v\n", err)
}
