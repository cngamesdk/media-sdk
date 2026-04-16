package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// ========== 获取定向标签报表测试用例 ==========

// TestTargetingTagReportsGetGenderSelf 测试获取定向标签报表-性别-广告主级别
func TestTargetingTagReportsGetGenderSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.TargetingTagReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Type = model.TargetingTagTypeGender
	req.Level = model.TargetingTagLevelAdvertiser
	req.DateRange = &model.TargetingTagReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"date", "gender_id"}
	req.Fields = []string{"date", "gender_id", "account_id", "cost", "view_count", "valid_click_count", "ctr", "cpc"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.TargetingTagReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestTargetingTagReportsGetRegionAdgroupSelf 测试获取定向标签报表-地域-广告组级别（带filtering）
func TestTargetingTagReportsGetRegionAdgroupSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.TargetingTagReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Type = model.TargetingTagTypeRegion
	req.Level = model.TargetingTagLevelAdgroup
	req.DateRange = &model.TargetingTagReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.Filtering = []*model.TargetingTagReportFilter{
		{
			Field:    model.TargetingTagReportFilterAdgroupId,
			Operator: model.TargetingTagReportOperatorIn,
			Values:   []string{"111111", "222222"},
		},
	}
	req.GroupBy = []string{"date", "region_id", "adgroup_id"}
	req.Fields = []string{"date", "adgroup_id", "region_id", "view_count", "valid_click_count", "cost"}
	req.TimeLine = model.TimeLineRequestTime
	req.Page = 1
	req.PageSize = 10
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.TargetingTagReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestTargetingTagReportsGetAgeSelf 测试获取定向标签报表-年龄
func TestTargetingTagReportsGetAgeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.TargetingTagReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Type = model.TargetingTagTypeAge
	req.Level = model.TargetingTagLevelAdvertiser
	req.DateRange = &model.TargetingTagReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-07",
	}
	req.GroupBy = []string{"date", "age_range"}
	req.Fields = []string{"date", "age_range", "cost", "view_count", "valid_click_count", "conversions_count", "conversions_cost"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.TargetingTagReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestTargetingTagReportsGetCustomAudienceSelf 测试获取定向标签报表-自定义人群-动态创意级别
func TestTargetingTagReportsGetCustomAudienceSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.TargetingTagReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Type = model.TargetingTagTypeCustomAudience
	req.Level = model.TargetingTagLevelDynamicCreative
	req.DateRange = &model.TargetingTagReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.Filtering = []*model.TargetingTagReportFilter{
		{
			Field:    model.TargetingTagReportFilterAdgroupId,
			Operator: model.TargetingTagReportOperatorEquals,
			Values:   []string{"111111"},
		},
	}
	req.GroupBy = []string{"date", "custom_audience_id"}
	req.Fields = []string{"date", "custom_audience_id", "cost", "view_count", "valid_click_count"}
	req.OrderBy = []*model.TargetingTagReportOrderBy{
		{
			SortField: "cost",
			SortType:  model.SortTypeDescending,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.TargetingTagReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestTargetingTagReportsGetOsSelf 测试获取定向标签报表-操作系统
func TestTargetingTagReportsGetOsSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.TargetingTagReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Type = model.TargetingTagTypeOs
	req.Level = model.TargetingTagLevelAdvertiser
	req.DateRange = &model.TargetingTagReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"date", "os_platform"}
	req.Fields = []string{"date", "os_platform", "cost", "view_count", "valid_click_count", "activated_count", "activated_cost"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.TargetingTagReportsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 验证测试用例 ==========

// TestTargetingTagReportsGetValidateAccountIdEmptySelf 测试account_id为空
func TestTargetingTagReportsGetValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.TargetingTagReportsGetReq{}
	req.AccessToken = "123"
	req.Type = model.TargetingTagTypeGender
	req.Level = model.TargetingTagLevelAdvertiser
	req.DateRange = &model.TargetingTagReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"date"}
	req.Fields = []string{"date", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestTargetingTagReportsGetValidateTypeEmptySelf 测试type为空
func TestTargetingTagReportsGetValidateTypeEmptySelf(t *testing.T) {
	req := &model.TargetingTagReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Level = model.TargetingTagLevelAdvertiser
	req.DateRange = &model.TargetingTagReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"date"}
	req.Fields = []string{"date", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestTargetingTagReportsGetValidateInvalidTypeSelf 测试无效的type
func TestTargetingTagReportsGetValidateInvalidTypeSelf(t *testing.T) {
	req := &model.TargetingTagReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Type = "INVALID_TYPE"
	req.Level = model.TargetingTagLevelAdvertiser
	req.DateRange = &model.TargetingTagReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"date"}
	req.Fields = []string{"date", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：type值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestTargetingTagReportsGetValidateLevelEmptySelf 测试level为空
func TestTargetingTagReportsGetValidateLevelEmptySelf(t *testing.T) {
	req := &model.TargetingTagReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Type = model.TargetingTagTypeGender
	req.DateRange = &model.TargetingTagReportDateRange{
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

// TestTargetingTagReportsGetValidateInvalidLevelSelf 测试非CUSTOM_AUDIENCE类型使用DYNAMIC_CREATIVE级别
func TestTargetingTagReportsGetValidateInvalidLevelSelf(t *testing.T) {
	req := &model.TargetingTagReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Type = model.TargetingTagTypeGender
	req.Level = model.TargetingTagLevelDynamicCreative // GENDER不支持DYNAMIC_CREATIVE
	req.DateRange = &model.TargetingTagReportDateRange{
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

// TestTargetingTagReportsGetValidateFilteringRequiredSelf 测试ADGROUP级别时filtering必填
func TestTargetingTagReportsGetValidateFilteringRequiredSelf(t *testing.T) {
	req := &model.TargetingTagReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Type = model.TargetingTagTypeRegion
	req.Level = model.TargetingTagLevelAdgroup
	req.DateRange = &model.TargetingTagReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"date", "region_id"}
	req.Fields = []string{"date", "region_id", "cost"}
	// 未设置filtering

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：level为ADGROUP时filtering为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestTargetingTagReportsGetValidateDateRangeInvalidSelf 测试start_date > end_date
func TestTargetingTagReportsGetValidateDateRangeInvalidSelf(t *testing.T) {
	req := &model.TargetingTagReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Type = model.TargetingTagTypeGender
	req.Level = model.TargetingTagLevelAdvertiser
	req.DateRange = &model.TargetingTagReportDateRange{
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

// TestTargetingTagReportsGetValidateFilteringInvalidOperatorSelf 测试gender过滤使用无效操作符
func TestTargetingTagReportsGetValidateFilteringInvalidOperatorSelf(t *testing.T) {
	req := &model.TargetingTagReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Type = model.TargetingTagTypeGender
	req.Level = model.TargetingTagLevelAdgroup
	req.DateRange = &model.TargetingTagReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.Filtering = []*model.TargetingTagReportFilter{
		{
			Field:    model.TargetingTagReportFilterGender,
			Operator: model.TargetingTagReportOperatorIn, // gender只支持EQUALS
			Values:   []string{"1"},
		},
	}
	req.GroupBy = []string{"date"}
	req.Fields = []string{"date", "cost"}

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.field为gender时operator只允许EQUALS")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestTargetingTagReportsGetValidatePageExceedSelf 测试page超出范围
func TestTargetingTagReportsGetValidatePageExceedSelf(t *testing.T) {
	req := &model.TargetingTagReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Type = model.TargetingTagTypeGender
	req.Level = model.TargetingTagLevelAdvertiser
	req.DateRange = &model.TargetingTagReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"date"}
	req.Fields = []string{"date", "cost"}
	req.Page = 101 // 超出最大值100

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page必须在1-100之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestTargetingTagReportsGetValidatePageSizeExceedSelf 测试page*pageSize超过20000
func TestTargetingTagReportsGetValidatePageSizeExceedSelf(t *testing.T) {
	req := &model.TargetingTagReportsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Type = model.TargetingTagTypeGender
	req.Level = model.TargetingTagLevelAdvertiser
	req.DateRange = &model.TargetingTagReportDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-01",
	}
	req.GroupBy = []string{"date"}
	req.Fields = []string{"date", "cost"}
	req.Page = 100
	req.PageSize = 2000 // 100 * 2000 = 200000 > 20000

	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page * page_size必须小于等于20000")
	}
	fmt.Printf("验证错误: %v\n", err)
}
