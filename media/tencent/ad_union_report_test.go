package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// ========== 获取联盟广告位报表测试用例 ==========

// TestAdUnionReportsGetSelf 测试获取联盟广告位报表-基础查询
func TestAdUnionReportsGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdUnionReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.DateRange = &model.AdUnionReportDateRange{
		StartDate: "2024-03-19",
		EndDate:   "2024-03-19",
	}
	req.GroupBy = []string{model.AdUnionGroupByDate, model.AdUnionGroupByUnionPositionID, model.AdUnionGroupByPlacementType, model.AdUnionGroupByIndustryParentID}
	req.Fields = []string{"promoted_object_type", "union_position_id", "placement_name", "industry_parent_name",
		"cost", "view_count", "valid_click_count", "ctr", "cpc", "conversions_count", "conversions_cost"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdUnionReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdUnionReportsGetWithFilteringSelf 测试获取联盟广告位报表-带过滤条件
func TestAdUnionReportsGetWithFilteringSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdUnionReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.DateRange = &model.AdUnionReportDateRange{
		StartDate: "2024-03-01",
		EndDate:   "2024-03-07",
	}
	req.Filtering = &model.AdUnionReportFiltering{
		UnionPositionID: []int64{111111, 222222},
	}
	req.GroupBy = []string{model.AdUnionGroupByDate, model.AdUnionGroupByUnionPositionID}
	req.Fields = []string{"date", "union_position_id", "cost", "view_count", "valid_click_count",
		"activated_count", "download_count", "install_count"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdUnionReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdUnionReportsGetWithOrderBySelf 测试获取联盟广告位报表-带排序
func TestAdUnionReportsGetWithOrderBySelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdUnionReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.DateRange = &model.AdUnionReportDateRange{
		StartDate: "2024-03-19",
		EndDate:   "2024-03-19",
	}
	req.GroupBy = []string{model.AdUnionGroupByUnionPositionID}
	req.Fields = []string{"union_position_id", "cost", "view_count", "valid_click_count", "order_pv", "order_amount", "order_roi"}
	req.OrderBy = []*model.AdUnionReportOrderBy{
		{
			SortField: "cost",
			SortType:  model.SortTypeDescending,
		},
	}
	req.Page = 1
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdUnionReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestAdUnionReportsGetAppMetricsSelf 测试获取联盟广告位报表-APP相关指标
func TestAdUnionReportsGetAppMetricsSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdUnionReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.DateRange = &model.AdUnionReportDateRange{
		StartDate: "2024-03-19",
		EndDate:   "2024-03-19",
	}
	req.GroupBy = []string{model.AdUnionGroupByDate}
	req.Fields = []string{"date", "cost", "first_pay_count", "first_pay_cost", "activated_count",
		"deep_conversions_count", "download_count", "install_count", "reg_pv", "purchase_pv"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdUnionReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 验证测试用例 ==========

// TestAdUnionReportsGetValidateAccountIdEmptySelf 测试account_id为空
func TestAdUnionReportsGetValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.AdUnionReportsGetReq{}
	req.AccessToken = "123"
	req.DateRange = &model.AdUnionReportDateRange{
		StartDate: "2024-03-19",
		EndDate:   "2024-03-19",
	}
	req.Fields = []string{"date", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdUnionReportsGetValidateDateRangeNilSelf 测试date_range为空
func TestAdUnionReportsGetValidateDateRangeNilSelf(t *testing.T) {
	req := &model.AdUnionReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Fields = []string{"date", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：date_range为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdUnionReportsGetValidateDateRangeInvalidSelf 测试start_date > end_date
func TestAdUnionReportsGetValidateDateRangeInvalidSelf(t *testing.T) {
	req := &model.AdUnionReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.DateRange = &model.AdUnionReportDateRange{
		StartDate: "2024-03-31",
		EndDate:   "2024-03-01",
	}
	req.Fields = []string{"date", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：date_range.start_date必须小于等于end_date")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdUnionReportsGetValidateFieldsEmptySelf 测试fields为空
func TestAdUnionReportsGetValidateFieldsEmptySelf(t *testing.T) {
	req := &model.AdUnionReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.DateRange = &model.AdUnionReportDateRange{
		StartDate: "2024-03-19",
		EndDate:   "2024-03-19",
	}
	req.Fields = []string{}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：fields数组长度必须在1-1024之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdUnionReportsGetValidateFilteringExceedSelf 测试filtering超过20个
func TestAdUnionReportsGetValidateFilteringExceedSelf(t *testing.T) {
	req := &model.AdUnionReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.DateRange = &model.AdUnionReportDateRange{
		StartDate: "2024-03-19",
		EndDate:   "2024-03-19",
	}
	req.Fields = []string{"date", "cost"}
	ids := make([]int64, 21)
	for i := range ids {
		ids[i] = int64(i + 1)
	}
	req.Filtering = &model.AdUnionReportFiltering{
		UnionPositionID: ids,
	}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.union_position_id数组长度必须在1-20之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdUnionReportsGetValidateGroupByExceedSelf 测试group_by超过5个
func TestAdUnionReportsGetValidateGroupByExceedSelf(t *testing.T) {
	req := &model.AdUnionReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.DateRange = &model.AdUnionReportDateRange{
		StartDate: "2024-03-19",
		EndDate:   "2024-03-19",
	}
	req.Fields = []string{"date", "cost"}
	req.GroupBy = []string{"a", "b", "c", "d", "e", "f"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：group_by数组长度必须在1-5之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdUnionReportsGetValidatePageSizeExceedSelf 测试page_size超过100
func TestAdUnionReportsGetValidatePageSizeExceedSelf(t *testing.T) {
	req := &model.AdUnionReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.DateRange = &model.AdUnionReportDateRange{
		StartDate: "2024-03-19",
		EndDate:   "2024-03-19",
	}
	req.Fields = []string{"date", "cost"}
	req.PageSize = 101

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size必须在1-100之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestAdUnionReportsGetValidateInvalidSortTypeSelf 测试无效的排序方式
func TestAdUnionReportsGetValidateInvalidSortTypeSelf(t *testing.T) {
	req := &model.AdUnionReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.DateRange = &model.AdUnionReportDateRange{
		StartDate: "2024-03-19",
		EndDate:   "2024-03-19",
	}
	req.Fields = []string{"date", "cost"}
	req.OrderBy = []*model.AdUnionReportOrderBy{
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
