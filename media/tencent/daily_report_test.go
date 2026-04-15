package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// ========== 获取日报表测试用例 ==========

// TestDailyReportsGetSelf 测试获取日报表-广告主级别
func TestDailyReportsGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DailyReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Level = model.ReportLevelAdvertiser
	req.DateRange = &model.DailyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"date"}
	req.Fields = []string{"date", "account_id", "cost", "view_count", "valid_click_count", "ctr", "cpc", "conversions_count", "conversions_cost"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DailyReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestDailyReportsGetAdgroupLevelSelf 测试获取日报表-广告级别
func TestDailyReportsGetAdgroupLevelSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DailyReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Level = model.ReportLevelAdgroup
	req.DateRange = &model.DailyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-07",
	}
	req.GroupBy = []string{"adgroup_id", "date"}
	req.Fields = []string{"date", "adgroup_id", "cost", "view_count", "valid_click_count"}
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DailyReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestDailyReportsGetWithFilteringSelf 测试获取日报表-带过滤条件
func TestDailyReportsGetWithFilteringSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DailyReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Level = model.ReportLevelAdgroup
	req.DateRange = &model.DailyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"adgroup_id"}
	req.Fields = []string{"adgroup_id", "cost", "view_count", "valid_click_count"}
	req.Filtering = []*model.DailyReportFilter{
		{
			Field:    model.DailyReportFilterAdgroupId,
			Operator: "IN",
			Values:   []string{"111111", "222222"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DailyReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestDailyReportsGetWithOrderBySelf 测试获取日报表-带排序
func TestDailyReportsGetWithOrderBySelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DailyReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Level = model.ReportLevelAdvertiser
	req.DateRange = &model.DailyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-31",
	}
	req.GroupBy = []string{"date"}
	req.Fields = []string{"date", "cost", "view_count"}
	req.OrderBy = []*model.DailyReportOrderBy{
		{
			SortField: "cost",
			SortType:  model.SortTypeDescending,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DailyReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestDailyReportsGetMaterialVideoSelf 测试获取日报表-素材视频级别
func TestDailyReportsGetMaterialVideoSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DailyReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Level = model.ReportLevelMaterialVideo
	req.TimeLine = model.TimeLineRequestTime
	req.DateRange = &model.DailyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"video_id"}
	req.Fields = []string{"video_id", "cost", "view_count", "valid_click_count", "video_outer_play_count"}
	req.Page = 1
	req.PageSize = 10
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DailyReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 验证测试用例 ==========

// TestDailyReportsGetValidateLevelEmptySelf 测试level为空
func TestDailyReportsGetValidateLevelEmptySelf(t *testing.T) {
	req := &model.DailyReportsGetReq{}
	req.AccessToken = "123"
	req.DateRange = &model.DailyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"date"}
	req.Fields = []string{"date", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：level为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestDailyReportsGetValidateDateRangeNilSelf 测试date_range为空
func TestDailyReportsGetValidateDateRangeNilSelf(t *testing.T) {
	req := &model.DailyReportsGetReq{}
	req.AccessToken = "123"
	req.Level = model.ReportLevelAdvertiser
	req.GroupBy = []string{"date"}
	req.Fields = []string{"date", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：date_range为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestDailyReportsGetValidateDateRangeInvalidSelf 测试start_date > end_date
func TestDailyReportsGetValidateDateRangeInvalidSelf(t *testing.T) {
	req := &model.DailyReportsGetReq{}
	req.AccessToken = "123"
	req.Level = model.ReportLevelAdvertiser
	req.DateRange = &model.DailyReportDateRange{
		StartDate: "2024-01-31",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"date"}
	req.Fields = []string{"date", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：date_range.start_date必须小于等于end_date")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestDailyReportsGetValidateGroupByEmptySelf 测试group_by为空
func TestDailyReportsGetValidateGroupByEmptySelf(t *testing.T) {
	req := &model.DailyReportsGetReq{}
	req.AccessToken = "123"
	req.Level = model.ReportLevelAdvertiser
	req.DateRange = &model.DailyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{}
	req.Fields = []string{"date", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：group_by数组长度必须在1-10之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestDailyReportsGetValidateFieldsEmptySelf 测试fields为空
func TestDailyReportsGetValidateFieldsEmptySelf(t *testing.T) {
	req := &model.DailyReportsGetReq{}
	req.AccessToken = "123"
	req.Level = model.ReportLevelAdvertiser
	req.DateRange = &model.DailyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"date"}
	req.Fields = []string{}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：fields数组长度必须在1-1024之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestDailyReportsGetValidatePageSizeExceedSelf 测试page*pageSize超过20000
func TestDailyReportsGetValidatePageSizeExceedSelf(t *testing.T) {
	req := &model.DailyReportsGetReq{}
	req.AccessToken = "123"
	req.Level = model.ReportLevelAdvertiser
	req.DateRange = &model.DailyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"date"}
	req.Fields = []string{"date", "cost"}
	req.Page = 100
	req.PageSize = 2000

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page * page_size必须小于等于20000")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestDailyReportsGetValidateInvalidLevelSelf 测试无效的level
func TestDailyReportsGetValidateInvalidLevelSelf(t *testing.T) {
	req := &model.DailyReportsGetReq{}
	req.AccessToken = "123"
	req.Level = "INVALID_LEVEL"
	req.DateRange = &model.DailyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"date"}
	req.Fields = []string{"date", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：level值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestDailyReportsGetValidateInvalidSortTypeSelf 测试无效的排序方式
func TestDailyReportsGetValidateInvalidSortTypeSelf(t *testing.T) {
	req := &model.DailyReportsGetReq{}
	req.AccessToken = "123"
	req.Level = model.ReportLevelAdvertiser
	req.DateRange = &model.DailyReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"date"}
	req.Fields = []string{"date", "cost"}
	req.OrderBy = []*model.DailyReportOrderBy{
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
