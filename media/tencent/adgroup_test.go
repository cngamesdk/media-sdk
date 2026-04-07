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

// ========== 批量修改广告开启/暂停状态测试用例 ==========

// TestAdgroupsUpdateConfiguredStatusSuspendSelf 测试批量暂停单个广告
func TestAdgroupsUpdateConfiguredStatusSuspendSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupsUpdateConfiguredStatusReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateConfiguredStatusSpec = []*model.UpdateConfiguredStatusSpec{
		{AdgroupID: 13397328752, ConfiguredStatus: model.ConfiguredStatusSuspend},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupsUpdateConfiguredStatusSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdgroupsUpdateConfiguredStatusNormalSelf 测试批量开启单个广告
func TestAdgroupsUpdateConfiguredStatusNormalSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupsUpdateConfiguredStatusReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateConfiguredStatusSpec = []*model.UpdateConfiguredStatusSpec{
		{AdgroupID: 13397328752, ConfiguredStatus: model.ConfiguredStatusNormal},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupsUpdateConfiguredStatusSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdgroupsUpdateConfiguredStatusMultipleSelf 测试批量修改多个广告状态（混合开启/暂停）
func TestAdgroupsUpdateConfiguredStatusMultipleSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupsUpdateConfiguredStatusReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateConfiguredStatusSpec = []*model.UpdateConfiguredStatusSpec{
		{AdgroupID: 111, ConfiguredStatus: model.ConfiguredStatusNormal},
		{AdgroupID: 222, ConfiguredStatus: model.ConfiguredStatusSuspend},
		{AdgroupID: 333, ConfiguredStatus: model.ConfiguredStatusNormal},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupsUpdateConfiguredStatusSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 批量修改广告状态验证测试用例 ==========

// TestAdgroupsUpdateConfiguredStatusValidateMissingAccountIDSelf 测试缺少account_id
func TestAdgroupsUpdateConfiguredStatusValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.AdgroupsUpdateConfiguredStatusReq{}
	req.AccessToken = "123"
	req.UpdateConfiguredStatusSpec = []*model.UpdateConfiguredStatusSpec{
		{AdgroupID: 111, ConfiguredStatus: model.ConfiguredStatusNormal},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateConfiguredStatusValidateEmptySpecSelf 测试spec为空
func TestAdgroupsUpdateConfiguredStatusValidateEmptySpecSelf(t *testing.T) {
	req := &model.AdgroupsUpdateConfiguredStatusReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateConfiguredStatusSpec = []*model.UpdateConfiguredStatusSpec{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：update_configured_status_spec至少包含1个条件")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateConfiguredStatusValidateExceedMaxSelf 测试spec超过100条
func TestAdgroupsUpdateConfiguredStatusValidateExceedMaxSelf(t *testing.T) {
	req := &model.AdgroupsUpdateConfiguredStatusReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	specs := make([]*model.UpdateConfiguredStatusSpec, 101)
	for i := range specs {
		specs[i] = &model.UpdateConfiguredStatusSpec{
			AdgroupID:        int64(i + 1),
			ConfiguredStatus: model.ConfiguredStatusNormal,
		}
	}
	req.UpdateConfiguredStatusSpec = specs
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：spec超过100条")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateConfiguredStatusValidateMissingAdgroupIDSelf 测试spec中缺少adgroup_id
func TestAdgroupsUpdateConfiguredStatusValidateMissingAdgroupIDSelf(t *testing.T) {
	req := &model.AdgroupsUpdateConfiguredStatusReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateConfiguredStatusSpec = []*model.UpdateConfiguredStatusSpec{
		{ConfiguredStatus: model.ConfiguredStatusNormal},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：adgroup_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateConfiguredStatusValidateMissingStatusSelf 测试spec中缺少configured_status
func TestAdgroupsUpdateConfiguredStatusValidateMissingStatusSelf(t *testing.T) {
	req := &model.AdgroupsUpdateConfiguredStatusReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateConfiguredStatusSpec = []*model.UpdateConfiguredStatusSpec{
		{AdgroupID: 111},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：configured_status为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateConfiguredStatusValidateInvalidStatusSelf 测试configured_status值无效
func TestAdgroupsUpdateConfiguredStatusValidateInvalidStatusSelf(t *testing.T) {
	req := &model.AdgroupsUpdateConfiguredStatusReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateConfiguredStatusSpec = []*model.UpdateConfiguredStatusSpec{
		{AdgroupID: 111, ConfiguredStatus: "AD_STATUS_UNKNOWN"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：configured_status值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateConfiguredStatusValidateDuplicateAdgroupIDSelf 测试adgroup_id重复
func TestAdgroupsUpdateConfiguredStatusValidateDuplicateAdgroupIDSelf(t *testing.T) {
	req := &model.AdgroupsUpdateConfiguredStatusReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateConfiguredStatusSpec = []*model.UpdateConfiguredStatusSpec{
		{AdgroupID: 111, ConfiguredStatus: model.ConfiguredStatusNormal},
		{AdgroupID: 111, ConfiguredStatus: model.ConfiguredStatusSuspend},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：adgroup_id不允许重复")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateConfiguredStatusValidateNilSpecItemSelf 测试spec含nil项
func TestAdgroupsUpdateConfiguredStatusValidateNilSpecItemSelf(t *testing.T) {
	req := &model.AdgroupsUpdateConfiguredStatusReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateConfiguredStatusSpec = []*model.UpdateConfiguredStatusSpec{
		{AdgroupID: 111, ConfiguredStatus: model.ConfiguredStatusNormal},
		nil,
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：spec项不能为nil")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// ========== 批量修改广告出价测试用例 ==========

// TestAdgroupsUpdateBidAmountSingleSelf 测试修改单个广告出价
func TestAdgroupsUpdateBidAmountSingleSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupsUpdateBidAmountReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateBidAmountSpec = []*model.UpdateBidAmountSpec{
		{AdgroupID: 13397328752, BidAmount: 5000},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupsUpdateBidAmountSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdgroupsUpdateBidAmountMultipleSelf 测试批量修改多个广告出价
func TestAdgroupsUpdateBidAmountMultipleSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupsUpdateBidAmountReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateBidAmountSpec = []*model.UpdateBidAmountSpec{
		{AdgroupID: 111, BidAmount: 200},
		{AdgroupID: 222, BidAmount: 1000},
		{AdgroupID: 333, BidAmount: 50000},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupsUpdateBidAmountSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdgroupsUpdateBidAmountAdxDefaultSelf 测试 ADX 程序化广告默认出价（200分）
func TestAdgroupsUpdateBidAmountAdxDefaultSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupsUpdateBidAmountReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateBidAmountSpec = []*model.UpdateBidAmountSpec{
		{AdgroupID: 13397328752, BidAmount: 200},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupsUpdateBidAmountSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 批量修改广告出价验证测试用例 ==========

// TestAdgroupsUpdateBidAmountValidateMissingAccountIDSelf 测试缺少account_id
func TestAdgroupsUpdateBidAmountValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.AdgroupsUpdateBidAmountReq{}
	req.AccessToken = "123"
	req.UpdateBidAmountSpec = []*model.UpdateBidAmountSpec{
		{AdgroupID: 111, BidAmount: 5000},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateBidAmountValidateEmptySpecSelf 测试spec为空
func TestAdgroupsUpdateBidAmountValidateEmptySpecSelf(t *testing.T) {
	req := &model.AdgroupsUpdateBidAmountReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateBidAmountSpec = []*model.UpdateBidAmountSpec{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：update_bid_amount_spec至少包含1个条件")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateBidAmountValidateExceedMaxSelf 测试spec超过100条
func TestAdgroupsUpdateBidAmountValidateExceedMaxSelf(t *testing.T) {
	req := &model.AdgroupsUpdateBidAmountReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	specs := make([]*model.UpdateBidAmountSpec, 101)
	for i := range specs {
		specs[i] = &model.UpdateBidAmountSpec{AdgroupID: int64(i + 1), BidAmount: 5000}
	}
	req.UpdateBidAmountSpec = specs
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：spec超过100条")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateBidAmountValidateMissingAdgroupIDSelf 测试spec中缺少adgroup_id
func TestAdgroupsUpdateBidAmountValidateMissingAdgroupIDSelf(t *testing.T) {
	req := &model.AdgroupsUpdateBidAmountReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateBidAmountSpec = []*model.UpdateBidAmountSpec{
		{BidAmount: 5000},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：adgroup_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateBidAmountValidateZeroBidAmountSelf 测试bid_amount为0
func TestAdgroupsUpdateBidAmountValidateZeroBidAmountSelf(t *testing.T) {
	req := &model.AdgroupsUpdateBidAmountReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateBidAmountSpec = []*model.UpdateBidAmountSpec{
		{AdgroupID: 111, BidAmount: 0},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：bid_amount必须大于0")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateBidAmountValidateNegativeBidAmountSelf 测试bid_amount为负数
func TestAdgroupsUpdateBidAmountValidateNegativeBidAmountSelf(t *testing.T) {
	req := &model.AdgroupsUpdateBidAmountReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateBidAmountSpec = []*model.UpdateBidAmountSpec{
		{AdgroupID: 111, BidAmount: -100},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：bid_amount必须大于0")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateBidAmountValidateDuplicateAdgroupIDSelf 测试adgroup_id重复
func TestAdgroupsUpdateBidAmountValidateDuplicateAdgroupIDSelf(t *testing.T) {
	req := &model.AdgroupsUpdateBidAmountReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateBidAmountSpec = []*model.UpdateBidAmountSpec{
		{AdgroupID: 111, BidAmount: 5000},
		{AdgroupID: 111, BidAmount: 8000},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：adgroup_id不允许重复")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdgroupsUpdateBidAmountValidateNilSpecItemSelf 测试spec含nil项
func TestAdgroupsUpdateBidAmountValidateNilSpecItemSelf(t *testing.T) {
	req := &model.AdgroupsUpdateBidAmountReq{}
	req.AccessToken = "123"
	req.AccountID = 20458
	req.UpdateBidAmountSpec = []*model.UpdateBidAmountSpec{
		{AdgroupID: 111, BidAmount: 5000},
		nil,
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：spec项不能为nil")
	}
	fmt.Printf("验证错误: %v\n", err)
}
