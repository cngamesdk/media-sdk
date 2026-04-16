package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// ========== 创建异步报表任务测试用例 ==========

// TestAsyncReportsAddDailySelf 测试创建异步报表任务-日报粒度
func TestAsyncReportsAddDailySelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AsyncReportsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.TaskName = "daily_report_task_1"
	req.Level = model.ReportLevelAdvertiser
	req.Granularity = model.AsyncReportGranularityDaily
	req.Date = "2024-01-01"
	req.ReportFields = []string{"date", "account_id", "cost", "view_count", "valid_click_count",
		"ctr", "cpc", "conversions_count", "conversions_cost"}
	req.TimeLine = model.TimeLineRequestTime
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AsyncReportsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAsyncReportsAddHourlySelf 测试创建异步报表任务-小时粒度
func TestAsyncReportsAddHourlySelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AsyncReportsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.TaskName = "hourly_report_task_1"
	req.Level = model.ReportLevelAdgroup
	req.Granularity = model.AsyncReportGranularityHourly
	req.Date = "2024-01-01"
	req.ReportFields = []string{"date", "adgroup_id", "hour", "cost", "view_count", "valid_click_count"}
	req.GroupBy = []string{"site_set"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AsyncReportsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAsyncReportsAddWithFilteringSelf 测试创建异步报表任务-带过滤条件
func TestAsyncReportsAddWithFilteringSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AsyncReportsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.TaskName = "filtered_report_task"
	req.Level = model.ReportLevelAdgroup
	req.Granularity = model.AsyncReportGranularityDaily
	req.Date = "2024-01-01"
	req.ReportFields = []string{"date", "adgroup_id", "cost", "view_count", "valid_click_count"}
	req.Filtering = []*model.AsyncReportFilter{
		{
			Field:    model.AsyncReportFilterAdgroupId,
			Operator: model.AsyncReportOperatorIn,
			Values:   []string{"111111", "222222"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AsyncReportsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAsyncReportsAddChannelLevelSelf 测试创建异步报表任务-渠道级别
func TestAsyncReportsAddChannelLevelSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AsyncReportsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.TaskName = "915225_169104810011_1"
	req.Level = model.ReportLevelChannel
	req.TimeLine = model.TimeLineRequestTime
	req.Granularity = model.AsyncReportGranularityDaily
	req.Date = "2023-11-24"
	req.ReportFields = []string{"date", "channel_id", "cost", "view_count", "valid_click_count"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AsyncReportsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 验证测试用例 ==========

// TestAsyncReportsAddValidateTaskNameEmptySelf 测试task_name为空
func TestAsyncReportsAddValidateTaskNameEmptySelf(t *testing.T) {
	req := &model.AsyncReportsAddReq{}
	req.AccessToken = "123"
	req.Level = model.ReportLevelAdvertiser
	req.Granularity = model.AsyncReportGranularityDaily
	req.Date = "2024-01-01"
	req.ReportFields = []string{"date", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：task_name为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAsyncReportsAddValidateLevelEmptySelf 测试level为空
func TestAsyncReportsAddValidateLevelEmptySelf(t *testing.T) {
	req := &model.AsyncReportsAddReq{}
	req.AccessToken = "123"
	req.TaskName = "test_task"
	req.Granularity = model.AsyncReportGranularityDaily
	req.Date = "2024-01-01"
	req.ReportFields = []string{"date", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：level为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAsyncReportsAddValidateInvalidLevelSelf 测试无效的level
func TestAsyncReportsAddValidateInvalidLevelSelf(t *testing.T) {
	req := &model.AsyncReportsAddReq{}
	req.AccessToken = "123"
	req.TaskName = "test_task"
	req.Level = "INVALID_LEVEL"
	req.Granularity = model.AsyncReportGranularityDaily
	req.Date = "2024-01-01"
	req.ReportFields = []string{"date", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：level值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAsyncReportsAddValidateGranularityEmptySelf 测试granularity为空
func TestAsyncReportsAddValidateGranularityEmptySelf(t *testing.T) {
	req := &model.AsyncReportsAddReq{}
	req.AccessToken = "123"
	req.TaskName = "test_task"
	req.Level = model.ReportLevelAdvertiser
	req.Date = "2024-01-01"
	req.ReportFields = []string{"date", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：granularity为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAsyncReportsAddValidateInvalidGranularitySelf 测试无效的granularity
func TestAsyncReportsAddValidateInvalidGranularitySelf(t *testing.T) {
	req := &model.AsyncReportsAddReq{}
	req.AccessToken = "123"
	req.TaskName = "test_task"
	req.Level = model.ReportLevelAdvertiser
	req.Granularity = "INVALID"
	req.Date = "2024-01-01"
	req.ReportFields = []string{"date", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：granularity值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAsyncReportsAddValidateDateEmptySelf 测试date为空
func TestAsyncReportsAddValidateDateEmptySelf(t *testing.T) {
	req := &model.AsyncReportsAddReq{}
	req.AccessToken = "123"
	req.TaskName = "test_task"
	req.Level = model.ReportLevelAdvertiser
	req.Granularity = model.AsyncReportGranularityDaily
	req.ReportFields = []string{"date", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：date为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAsyncReportsAddValidateReportFieldsEmptySelf 测试report_fields为空
func TestAsyncReportsAddValidateReportFieldsEmptySelf(t *testing.T) {
	req := &model.AsyncReportsAddReq{}
	req.AccessToken = "123"
	req.TaskName = "test_task"
	req.Level = model.ReportLevelAdvertiser
	req.Granularity = model.AsyncReportGranularityDaily
	req.Date = "2024-01-01"
	req.ReportFields = []string{}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：report_fields数组长度必须在1-1024之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAsyncReportsAddValidateFilteringExceedSelf 测试filtering超过5个
func TestAsyncReportsAddValidateFilteringExceedSelf(t *testing.T) {
	req := &model.AsyncReportsAddReq{}
	req.AccessToken = "123"
	req.TaskName = "test_task"
	req.Level = model.ReportLevelAdgroup
	req.Granularity = model.AsyncReportGranularityDaily
	req.Date = "2024-01-01"
	req.ReportFields = []string{"date", "cost"}
	req.Filtering = []*model.AsyncReportFilter{
		{Field: "a1", Operator: "IN", Values: []string{"1"}},
		{Field: "a2", Operator: "IN", Values: []string{"1"}},
		{Field: "a3", Operator: "IN", Values: []string{"1"}},
		{Field: "a4", Operator: "IN", Values: []string{"1"}},
		{Field: "a5", Operator: "IN", Values: []string{"1"}},
		{Field: "a6", Operator: "IN", Values: []string{"1"}},
	}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering数组长度必须在1-5之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAsyncReportsAddValidateInvalidTimeLineSelf 测试无效的time_line
func TestAsyncReportsAddValidateInvalidTimeLineSelf(t *testing.T) {
	req := &model.AsyncReportsAddReq{}
	req.AccessToken = "123"
	req.TaskName = "test_task"
	req.Level = model.ReportLevelAdvertiser
	req.Granularity = model.AsyncReportGranularityDaily
	req.Date = "2024-01-01"
	req.ReportFields = []string{"date", "cost"}
	req.TimeLine = "INVALID_TIME"

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：time_line值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// ========== 获取异步报表任务测试用例 ==========

// TestAsyncReportsGetSelf 测试获取异步报表任务-按task_id查询
func TestAsyncReportsGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AsyncReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.AsyncReportsGetFilter{
		{
			Field:    model.AsyncReportsGetFilterTaskId,
			Operator: model.AsyncReportOperatorEquals,
			Values:   []string{"53181057839"},
		},
	}
	req.Page = 1
	req.PageSize = 10
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AsyncReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAsyncReportsGetByTaskNameSelf 测试获取异步报表任务-按task_name查询
func TestAsyncReportsGetByTaskNameSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AsyncReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.AsyncReportsGetFilter{
		{
			Field:    model.AsyncReportsGetFilterTaskName,
			Operator: model.AsyncReportOperatorContains,
			Values:   []string{"daily_report"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AsyncReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAsyncReportsGetNoFilterSelf 测试获取异步报表任务-无过滤条件
func TestAsyncReportsGetNoFilterSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AsyncReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Page = 1
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AsyncReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 获取异步报表任务验证测试用例 ==========

// TestAsyncReportsGetValidateAccountIdEmptySelf 测试account_id为空
func TestAsyncReportsGetValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.AsyncReportsGetReq{}
	req.AccessToken = "123"

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAsyncReportsGetValidateInvalidFilterFieldSelf 测试无效的过滤字段
func TestAsyncReportsGetValidateInvalidFilterFieldSelf(t *testing.T) {
	req := &model.AsyncReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.AsyncReportsGetFilter{
		{
			Field:    "invalid_field",
			Operator: model.AsyncReportOperatorEquals,
			Values:   []string{"123"},
		},
	}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.field值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAsyncReportsGetValidatePageSizeExceedSelf 测试page_size超过100
func TestAsyncReportsGetValidatePageSizeExceedSelf(t *testing.T) {
	req := &model.AsyncReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.PageSize = 101

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size必须在1-100之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAsyncReportsGetValidateFilteringExceedSelf 测试filtering超过5个
func TestAsyncReportsGetValidateFilteringExceedSelf(t *testing.T) {
	req := &model.AsyncReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.AsyncReportsGetFilter{
		{Field: model.AsyncReportsGetFilterTaskId, Operator: model.AsyncReportOperatorEquals, Values: []string{"1"}},
		{Field: model.AsyncReportsGetFilterTaskId, Operator: model.AsyncReportOperatorEquals, Values: []string{"2"}},
		{Field: model.AsyncReportsGetFilterTaskId, Operator: model.AsyncReportOperatorEquals, Values: []string{"3"}},
		{Field: model.AsyncReportsGetFilterTaskId, Operator: model.AsyncReportOperatorEquals, Values: []string{"4"}},
		{Field: model.AsyncReportsGetFilterTaskId, Operator: model.AsyncReportOperatorEquals, Values: []string{"5"}},
		{Field: model.AsyncReportsGetFilterTaskId, Operator: model.AsyncReportOperatorEquals, Values: []string{"6"}},
	}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering数组长度必须在1-5之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// ========== 获取文件接口测试用例 ==========

// TestAsyncReportFilesGetSelf 测试获取文件
func TestAsyncReportFilesGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AsyncReportFilesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.TaskID = 53181057839
	req.FileID = 831530232
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AsyncReportFilesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("file data length: %d\n", len(result.FileData))
}

// TestAsyncReportFilesGetWithOrgIdSelf 测试获取文件-带业务单元id
func TestAsyncReportFilesGetWithOrgIdSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AsyncReportFilesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.TaskID = 25610
	req.FileID = 831530232
	req.OrganizationID = 100
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AsyncReportFilesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("file data length: %d\n", len(result.FileData))
}

// ========== 获取文件接口验证测试用例 ==========

// TestAsyncReportFilesGetValidateTaskIdEmptySelf 测试task_id为空
func TestAsyncReportFilesGetValidateTaskIdEmptySelf(t *testing.T) {
	req := &model.AsyncReportFilesGetReq{}
	req.AccessToken = "123"
	req.FileID = 831530232

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：task_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAsyncReportFilesGetValidateFileIdEmptySelf 测试file_id为空
func TestAsyncReportFilesGetValidateFileIdEmptySelf(t *testing.T) {
	req := &model.AsyncReportFilesGetReq{}
	req.AccessToken = "123"
	req.TaskID = 53181057839

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：file_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
