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
