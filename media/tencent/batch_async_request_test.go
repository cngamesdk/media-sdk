package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ========== 创建批量异步请求任务测试用例 ==========

// TestBatchAsyncRequestAddDeleteAdgroupSelf 测试批量删除广告任务
func TestBatchAsyncRequestAddDeleteAdgroupSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchAsyncRequestAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TaskName = "批量删除广告任务"
	req.TaskType = model.TaskTypeDeleteAdgroupNew
	req.TaskSpec = &model.BatchAsyncTaskSpec{
		DeleteAdgroupSpec: []*model.DeleteAdgroupItem{
			{AdgroupID: 10001},
			{AdgroupID: 10002},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchAsyncRequestAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBatchAsyncRequestAddUpdateConfiguredStatusSelf 测试批量修改广告状态任务
func TestBatchAsyncRequestAddUpdateConfiguredStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchAsyncRequestAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TaskName = "批量修改广告状态任务"
	req.TaskType = model.TaskTypeUpdateAdgroupConfiguredStatusNew
	req.TaskSpec = &model.BatchAsyncTaskSpec{
		UpdateAdgroupConfiguredStatusSpec: []*model.UpdateAdgroupConfiguredStatusItem{
			{AdgroupID: 10001, ConfiguredStatus: model.ConfiguredStatusNormal},
			{AdgroupID: 10002, ConfiguredStatus: model.ConfiguredStatusSuspend},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchAsyncRequestAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBatchAsyncRequestAddUpdateDailyBudgetSelf 测试批量修改广告日预算任务
func TestBatchAsyncRequestAddUpdateDailyBudgetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchAsyncRequestAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TaskName = "批量修改广告日预算任务"
	req.TaskType = model.TaskTypeUpdateAdgroupDailyBudgetNew
	req.TaskSpec = &model.BatchAsyncTaskSpec{
		UpdateAdgroupDailyBudgetSpec: []*model.UpdateAdgroupDailyBudgetItem{
			{AdgroupID: 10001, DailyBudget: 500000},
			{AdgroupID: 10002, DailyBudget: 1000000},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchAsyncRequestAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBatchAsyncRequestAddUpdateBidAmountSelf 测试批量修改广告出价任务
func TestBatchAsyncRequestAddUpdateBidAmountSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchAsyncRequestAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TaskName = "批量修改广告出价任务"
	req.TaskType = model.TaskTypeUpdateAdgroupBidAmountNew
	req.TaskSpec = &model.BatchAsyncTaskSpec{
		UpdateAdgroupBidAmountSpec: []*model.UpdateAdgroupBidAmountItem{
			{AdgroupID: 10001, BidAmount: 3000},
			{AdgroupID: 10002, BidAmount: 5000},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchAsyncRequestAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBatchAsyncRequestAddUpdateDeepConversionBidSelf 测试批量修改深度优化行为出价任务
func TestBatchAsyncRequestAddUpdateDeepConversionBidSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchAsyncRequestAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TaskName = "批量修改深度优化出价任务"
	req.TaskType = model.TaskTypeUpdateDeepConversionBehaviorBidNew
	req.TaskSpec = &model.BatchAsyncTaskSpec{
		UpdateDeepConversionBehaviorBidSpec: []*model.UpdateDeepConversionBehaviorBidItem{
			{AdgroupID: 10001, DeepConversionBehaviorBid: 50000},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchAsyncRequestAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBatchAsyncRequestAddAutoAcquisitionSelf 测试批量修改广告一键起量任务
func TestBatchAsyncRequestAddAutoAcquisitionSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchAsyncRequestAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TaskName = "批量修改广告一键起量任务"
	req.TaskType = model.TaskTypeUpdateAdgroupAutoAcquisitionNew
	req.TaskSpec = &model.BatchAsyncTaskSpec{
		UpdateAdgroupAutoAcquisitionSpec: []*model.UpdateAdgroupAutoAcquisitionItem{
			{AdgroupID: 10001, AutoAcquisitionEnabled: true, AutoAcquisitionBudget: 50000},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchAsyncRequestAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 验证测试用例 ==========

// TestBatchAsyncRequestAddValidateMissingAccountIDSelf 测试缺少account_id
func TestBatchAsyncRequestAddValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.BatchAsyncRequestAddReq{}
	req.AccessToken = "123"
	req.TaskName = "测试任务"
	req.TaskType = model.TaskTypeDeleteAdgroupNew
	req.TaskSpec = &model.BatchAsyncTaskSpec{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBatchAsyncRequestAddValidateMissingTaskNameSelf 测试缺少task_name
func TestBatchAsyncRequestAddValidateMissingTaskNameSelf(t *testing.T) {
	req := &model.BatchAsyncRequestAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TaskType = model.TaskTypeDeleteAdgroupNew
	req.TaskSpec = &model.BatchAsyncTaskSpec{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：task_name为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBatchAsyncRequestAddValidateTaskNameTooLongSelf 测试task_name超长（超过120字节）
func TestBatchAsyncRequestAddValidateTaskNameTooLongSelf(t *testing.T) {
	req := &model.BatchAsyncRequestAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	// 构造超过120字节的task_name
	longName := "这是一个非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常长的任务名称超过了限制"
	req.TaskName = longName
	req.TaskType = model.TaskTypeDeleteAdgroupNew
	req.TaskSpec = &model.BatchAsyncTaskSpec{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：task_name超过120字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBatchAsyncRequestAddValidateMissingTaskTypeSelf 测试缺少task_type
func TestBatchAsyncRequestAddValidateMissingTaskTypeSelf(t *testing.T) {
	req := &model.BatchAsyncRequestAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TaskName = "测试任务"
	req.TaskSpec = &model.BatchAsyncTaskSpec{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：task_type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBatchAsyncRequestAddValidateMissingTaskSpecSelf 测试缺少task_spec
func TestBatchAsyncRequestAddValidateMissingTaskSpecSelf(t *testing.T) {
	req := &model.BatchAsyncRequestAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TaskName = "测试任务"
	req.TaskType = model.TaskTypeDeleteAdgroupNew
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：task_spec为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
