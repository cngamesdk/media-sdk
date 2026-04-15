package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// ========== 获取小时报表测试用例 ==========

// TestHourlyReportsGetSelf 测试获取小时报表-广告主级别
func TestHourlyReportsGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.HourlyReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Level = model.ReportLevelAdvertiser
	req.DateRange = &model.HourlyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"hour"}
	req.Fields = []string{"hour", "account_id", "cost", "view_count", "valid_click_count", "ctr", "cpc", "conversions_count", "conversions_cost"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.HourlyReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestHourlyReportsGetAdgroupLevelSelf 测试获取小时报表-广告级别
func TestHourlyReportsGetAdgroupLevelSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.HourlyReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Level = model.ReportLevelAdgroup
	req.DateRange = &model.HourlyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"adgroup_id", "hour"}
	req.Fields = []string{"hour", "adgroup_id", "cost", "view_count", "valid_click_count"}
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.HourlyReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestHourlyReportsGetWithFilteringSelf 测试获取小时报表-带过滤条件
func TestHourlyReportsGetWithFilteringSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.HourlyReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Level = model.ReportLevelAdgroup
	req.DateRange = &model.HourlyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"adgroup_id", "hour"}
	req.Fields = []string{"hour", "adgroup_id", "cost", "view_count", "valid_click_count"}
	req.Filtering = []*model.HourlyReportFilter{
		{
			Field:    model.HourlyReportFilterAdgroupId,
			Operator: "IN",
			Values:   []string{"111111", "222222"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.HourlyReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestHourlyReportsGetWithOrderBySelf 测试获取小时报表-带排序
func TestHourlyReportsGetWithOrderBySelf(t *testing.T) {
	ctx := context.Background()
	req := &model.HourlyReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Level = model.ReportLevelAdvertiser
	req.DateRange = &model.HourlyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"hour"}
	req.Fields = []string{"hour", "cost", "view_count"}
	req.OrderBy = []*model.HourlyReportOrderBy{
		{
			SortField: "cost",
			SortType:  model.SortTypeDescending,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.HourlyReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestHourlyReportsGetWithTimeLineSelf 测试获取小时报表-指定时间口径
func TestHourlyReportsGetWithTimeLineSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.HourlyReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Level = model.ReportLevelAdvertiser
	req.TimeLine = model.TimeLineRequestTime
	req.DateRange = &model.HourlyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"hour"}
	req.Fields = []string{"hour", "cost", "view_count", "video_outer_play_count"}
	req.Page = 1
	req.PageSize = 10
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.HourlyReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 验证测试用例 ==========

// TestHourlyReportsGetValidateAccountIdEmptySelf 测试account_id为空
func TestHourlyReportsGetValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.HourlyReportsGetReq{}
	req.AccessToken = "123"
	req.Level = model.ReportLevelAdvertiser
	req.DateRange = &model.HourlyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"hour"}
	req.Fields = []string{"hour", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestHourlyReportsGetValidateLevelEmptySelf 测试level为空
func TestHourlyReportsGetValidateLevelEmptySelf(t *testing.T) {
	req := &model.HourlyReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.DateRange = &model.HourlyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"hour"}
	req.Fields = []string{"hour", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：level为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestHourlyReportsGetValidateInvalidLevelSelf 测试无效的level（小时报表不支持的级别）
func TestHourlyReportsGetValidateInvalidLevelSelf(t *testing.T) {
	req := &model.HourlyReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Level = model.ReportLevelComponent // 小时报表不支持Component级别
	req.DateRange = &model.HourlyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"hour"}
	req.Fields = []string{"hour", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：level值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestHourlyReportsGetValidateDateRangeNotEqualSelf 测试start_date != end_date
func TestHourlyReportsGetValidateDateRangeNotEqualSelf(t *testing.T) {
	req := &model.HourlyReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Level = model.ReportLevelAdvertiser
	req.DateRange = &model.HourlyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-02",
	}
	req.GroupBy = []string{"hour"}
	req.Fields = []string{"hour", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：小时报表date_range.start_date必须等于end_date")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestHourlyReportsGetValidateDateRangeNilSelf 测试date_range为空
func TestHourlyReportsGetValidateDateRangeNilSelf(t *testing.T) {
	req := &model.HourlyReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Level = model.ReportLevelAdvertiser
	req.GroupBy = []string{"hour"}
	req.Fields = []string{"hour", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：date_range为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestHourlyReportsGetValidateGroupByEmptySelf 测试group_by为空
func TestHourlyReportsGetValidateGroupByEmptySelf(t *testing.T) {
	req := &model.HourlyReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Level = model.ReportLevelAdvertiser
	req.DateRange = &model.HourlyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{}
	req.Fields = []string{"hour", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：group_by数组长度必须在1-10之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestHourlyReportsGetValidateFieldsEmptySelf 测试fields为空
func TestHourlyReportsGetValidateFieldsEmptySelf(t *testing.T) {
	req := &model.HourlyReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Level = model.ReportLevelAdvertiser
	req.DateRange = &model.HourlyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"hour"}
	req.Fields = []string{}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：fields数组长度必须在1-1024之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestHourlyReportsGetValidatePageExceedSelf 测试page超过100
func TestHourlyReportsGetValidatePageExceedSelf(t *testing.T) {
	req := &model.HourlyReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Level = model.ReportLevelAdvertiser
	req.DateRange = &model.HourlyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"hour"}
	req.Fields = []string{"hour", "cost"}
	req.Page = 101

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page必须在1-100之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestHourlyReportsGetValidateInvalidSortTypeSelf 测试无效的排序方式
func TestHourlyReportsGetValidateInvalidSortTypeSelf(t *testing.T) {
	req := &model.HourlyReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Level = model.ReportLevelAdvertiser
	req.DateRange = &model.HourlyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"hour"}
	req.Fields = []string{"hour", "cost"}
	req.OrderBy = []*model.HourlyReportOrderBy{
		{
			SortField: "cost",
			SortType:  "INVALID_SORT",
		},
	}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：order_by.sort_type值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}
