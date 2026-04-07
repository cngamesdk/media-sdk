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

// ========== 获取批量异步请求任务列表测试用例 ==========

// TestBatchAsyncRequestGetNoFilterSelf 测试不带过滤条件查询任务列表
func TestBatchAsyncRequestGetNoFilterSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchAsyncRequestGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchAsyncRequestGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBatchAsyncRequestGetByTaskIDSelf 测试按 task_id 过滤
func TestBatchAsyncRequestGetByTaskIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchAsyncRequestGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.BatchAsyncTaskFilteringItem{
		{Field: model.BatchAsyncTaskFilterFieldTaskID, Operator: model.OperatorEquals, Values: []string{"10001"}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchAsyncRequestGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBatchAsyncRequestGetByTaskTypeSelf 测试按 task_type 过滤
func TestBatchAsyncRequestGetByTaskTypeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchAsyncRequestGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.BatchAsyncTaskFilteringItem{
		{Field: model.BatchAsyncTaskFilterFieldTaskType, Operator: model.OperatorIn, Values: []string{
			model.TaskTypeDeleteAdgroupNew,
			model.TaskTypeUpdateAdgroupConfiguredStatusNew,
		}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchAsyncRequestGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBatchAsyncRequestGetByResultStatusSelf 测试按 result_status 过滤
func TestBatchAsyncRequestGetByResultStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchAsyncRequestGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.BatchAsyncTaskFilteringItem{
		{Field: model.BatchAsyncTaskFilterFieldResultStatus, Operator: model.OperatorIn, Values: []string{
			model.TaskResultStatusSuccess,
			model.TaskResultStatusPartialFail,
		}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchAsyncRequestGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBatchAsyncRequestGetByStatusSelf 测试按 status 过滤
func TestBatchAsyncRequestGetByStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchAsyncRequestGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.BatchAsyncTaskFilteringItem{
		{Field: model.BatchAsyncTaskFilterFieldStatus, Operator: model.OperatorEquals, Values: []string{model.TaskStatusCompleted}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchAsyncRequestGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBatchAsyncRequestGetWithPaginationSelf 测试分页参数
func TestBatchAsyncRequestGetWithPaginationSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchAsyncRequestGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 2
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchAsyncRequestGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 获取任务列表验证测试用例 ==========

// TestBatchAsyncRequestGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestBatchAsyncRequestGetValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.BatchAsyncRequestGetReq{}
	req.AccessToken = "123"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBatchAsyncRequestGetValidateFilteringExceedMaxSelf 测试 filtering 超过20条
func TestBatchAsyncRequestGetValidateFilteringExceedMaxSelf(t *testing.T) {
	req := &model.BatchAsyncRequestGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	filters := make([]*model.BatchAsyncTaskFilteringItem, 21)
	for i := range filters {
		filters[i] = &model.BatchAsyncTaskFilteringItem{
			Field:    model.BatchAsyncTaskFilterFieldTaskID,
			Operator: model.OperatorEquals,
			Values:   []string{"1"},
		}
	}
	req.Filtering = filters
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering超过20条")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBatchAsyncRequestGetValidateFilteringMissingFieldSelf 测试 filtering 缺少 field
func TestBatchAsyncRequestGetValidateFilteringMissingFieldSelf(t *testing.T) {
	req := &model.BatchAsyncRequestGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.BatchAsyncTaskFilteringItem{
		{Operator: model.OperatorEquals, Values: []string{"1"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：field为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBatchAsyncRequestGetValidateFilteringEmptyValuesSelf 测试 filtering values 为空
func TestBatchAsyncRequestGetValidateFilteringEmptyValuesSelf(t *testing.T) {
	req := &model.BatchAsyncRequestGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.BatchAsyncTaskFilteringItem{
		{Field: model.BatchAsyncTaskFilterFieldStatus, Operator: model.OperatorEquals, Values: []string{}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：values为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBatchAsyncRequestGetValidatePageTooLargeSelf 测试 page 超过最大值
func TestBatchAsyncRequestGetValidatePageTooLargeSelf(t *testing.T) {
	req := &model.BatchAsyncRequestGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 100000
	req.PageSize = 10
	req.Format()
	// reset Page after Format to bypass default assignment
	req.Page = 100000
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page超过99999")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBatchAsyncRequestGetValidatePageSizeTooLargeSelf 测试 page_size 超过最大值100
func TestBatchAsyncRequestGetValidatePageSizeTooLargeSelf(t *testing.T) {
	req := &model.BatchAsyncRequestGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 1
	req.PageSize = 101
	req.Format()
	// reset PageSize after Format to bypass default assignment
	req.PageSize = 101
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size超过100")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// ========== 获取批量异步请求任务详情测试用例 ==========

// TestBatchAsyncRequestSpecGetDefaultPageSelf 测试默认分页获取任务详情
func TestBatchAsyncRequestSpecGetDefaultPageSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchAsyncRequestSpecGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TaskID = 10001
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchAsyncRequestSpecGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBatchAsyncRequestSpecGetWithPaginationSelf 测试自定义分页获取任务详情
func TestBatchAsyncRequestSpecGetWithPaginationSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchAsyncRequestSpecGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TaskID = 10001
	req.Page = 2
	req.PageSize = 50
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchAsyncRequestSpecGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBatchAsyncRequestSpecGetMaxPageSizeSelf 测试最大 page_size（100）
func TestBatchAsyncRequestSpecGetMaxPageSizeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchAsyncRequestSpecGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TaskID = 10001
	req.Page = 1
	req.PageSize = 100
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchAsyncRequestSpecGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 获取任务详情验证测试用例 ==========

// TestBatchAsyncRequestSpecGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestBatchAsyncRequestSpecGetValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.BatchAsyncRequestSpecGetReq{}
	req.AccessToken = "123"
	req.TaskID = 10001
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBatchAsyncRequestSpecGetValidateMissingTaskIDSelf 测试缺少 task_id
func TestBatchAsyncRequestSpecGetValidateMissingTaskIDSelf(t *testing.T) {
	req := &model.BatchAsyncRequestSpecGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：task_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBatchAsyncRequestSpecGetValidatePageTooLargeSelf 测试 page 超过最大值 99999
func TestBatchAsyncRequestSpecGetValidatePageTooLargeSelf(t *testing.T) {
	req := &model.BatchAsyncRequestSpecGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TaskID = 10001
	req.Page = 1
	req.PageSize = 10
	req.Format()
	req.Page = 100000
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page超过99999")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBatchAsyncRequestSpecGetValidatePageSizeTooLargeSelf 测试 page_size 超过最大值 100
func TestBatchAsyncRequestSpecGetValidatePageSizeTooLargeSelf(t *testing.T) {
	req := &model.BatchAsyncRequestSpecGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TaskID = 10001
	req.Page = 1
	req.PageSize = 101
	req.Format()
	req.PageSize = 101
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size超过100")
	}
	fmt.Printf("验证错误: %v\n", err)
}
