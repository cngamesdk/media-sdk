package tencent

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ========== 获取落地页列表接口调用测试用例 ==========

// TestPagesGetBasicSelf 测试基本查询（不带过滤条件）
func TestPagesGetBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.PagesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestPagesGetFilterByPageTypeSelf 测试按 page_type 过滤查询
func TestPagesGetFilterByPageTypeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.PagesGetFilteringItem{
		{
			Field:    model.PagesGetFilterFieldPageType,
			Operator: model.PagesGetFilterOperatorEquals,
			Values:   []string{model.PageTypeOfficial},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.PagesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestPagesGetFilterByPageIDSelf 测试按 page_id 精确查询
func TestPagesGetFilterByPageIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.PagesGetFilteringItem{
		{
			Field:    model.PagesGetFilterFieldPageID,
			Operator: model.PagesGetFilterOperatorEquals,
			Values:   []string{"12345678"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.PagesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestPagesGetFilterByPageNameContainsSelf 测试按 page_name 模糊查询
func TestPagesGetFilterByPageNameContainsSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.PagesGetFilteringItem{
		{
			Field:    model.PagesGetFilterFieldPageName,
			Operator: model.PagesGetFilterOperatorContains,
			Values:   []string{"落地页"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.PagesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestPagesGetFilterByPageStatusSelf 测试按 page_status 过滤查询
func TestPagesGetFilterByPageStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.PagesGetFilteringItem{
		{
			Field:    model.PagesGetFilterFieldPageStatus,
			Operator: model.PagesGetFilterOperatorEquals,
			Values:   []string{model.PageStatusNormal},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.PagesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestPagesGetWithPaginationSelf 测试带分页参数查询
func TestPagesGetWithPaginationSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 2
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.PagesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 获取落地页列表参数验证测试用例 ==========

// TestPagesGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestPagesGetValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestPagesGetValidateFilteringTooLongSelf 测试 filtering 超过10条
func TestPagesGetValidateFilteringTooLongSelf(t *testing.T) {
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	list := make([]*model.PagesGetFilteringItem, 11)
	for i := range list {
		list[i] = &model.PagesGetFilteringItem{
			Field:    model.PagesGetFilterFieldPageType,
			Operator: model.PagesGetFilterOperatorEquals,
			Values:   []string{model.PageTypeOfficial},
		}
	}
	req.Filtering = list
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering超过10条")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestPagesGetValidateNilFilterItemSelf 测试 filtering 中包含 nil 元素
func TestPagesGetValidateNilFilterItemSelf(t *testing.T) {
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.PagesGetFilteringItem{
		{Field: model.PagesGetFilterFieldPageType, Operator: model.PagesGetFilterOperatorEquals, Values: []string{model.PageTypeOfficial}},
		nil,
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering中包含nil元素")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestPagesGetValidateFilterMissingFieldSelf 测试 filtering.field 为空
func TestPagesGetValidateFilterMissingFieldSelf(t *testing.T) {
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.PagesGetFilteringItem{
		{Operator: model.PagesGetFilterOperatorEquals, Values: []string{"1"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.field为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestPagesGetValidateFilterMissingOperatorSelf 测试 filtering.operator 为空
func TestPagesGetValidateFilterMissingOperatorSelf(t *testing.T) {
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.PagesGetFilteringItem{
		{Field: model.PagesGetFilterFieldPageType, Values: []string{model.PageTypeOfficial}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.operator为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestPagesGetValidateFilterEmptyValuesSelf 测试 filtering.values 为空
func TestPagesGetValidateFilterEmptyValuesSelf(t *testing.T) {
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.PagesGetFilteringItem{
		{Field: model.PagesGetFilterFieldPageType, Operator: model.PagesGetFilterOperatorEquals, Values: []string{}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.values为空")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestPagesGetValidatePageNameValueTooLongSelf 测试 page_name 过滤值超过180字节
func TestPagesGetValidatePageNameValueTooLongSelf(t *testing.T) {
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.PagesGetFilteringItem{
		{
			Field:    model.PagesGetFilterFieldPageName,
			Operator: model.PagesGetFilterOperatorContains,
			Values:   []string{strings.Repeat("a", 181)},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_name过滤值超过180字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestPagesGetValidatePageOutOfRangeSelf 测试 page 超出范围
func TestPagesGetValidatePageOutOfRangeSelf(t *testing.T) {
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 100000
	req.PageSize = 10
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page超出范围")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestPagesGetValidatePageSizeOutOfRangeSelf 测试 page_size 超出范围
func TestPagesGetValidatePageSizeOutOfRangeSelf(t *testing.T) {
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 1
	req.PageSize = 101
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size超出范围")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestPagesGetValidateAdContextMissingMarketingGoalSelf 测试 ad_context 缺少 marketing_goal
func TestPagesGetValidateAdContextMissingMarketingGoalSelf(t *testing.T) {
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.AdContext = &model.PagesGetAdContext{
		MarketingCarrierType: "MARKETING_CARRIER_TYPE_APP_ANDROID",
		MarketingTargetType:  "MARKETING_TARGET_TYPE_APP",
		SiteSet:              []string{"SITE_SET_WECHAT"},
		CreativeTemplateID:   1001,
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：ad_context.marketing_goal为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestPagesGetValidateAdContextEmptySiteSetSelf 测试 ad_context.site_set 为空
func TestPagesGetValidateAdContextEmptySiteSetSelf(t *testing.T) {
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.AdContext = &model.PagesGetAdContext{
		MarketingGoal:        "MARKETING_GOAL_USER_GROWTH",
		MarketingCarrierType: "MARKETING_CARRIER_TYPE_APP_ANDROID",
		MarketingTargetType:  "MARKETING_TARGET_TYPE_APP",
		SiteSet:              []string{},
		CreativeTemplateID:   1001,
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：ad_context.site_set为空")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestPagesGetValidateDefaultPaginationSelf 测试 Format() 默认填充分页参数
func TestPagesGetValidateDefaultPaginationSelf(t *testing.T) {
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Format()
	if req.Page != 1 {
		t.Fatalf("期望默认page=1，实际=%d", req.Page)
	}
	if req.PageSize != 10 {
		t.Fatalf("期望默认page_size=10，实际=%d", req.PageSize)
	}
	err := req.Validate()
	if err != nil {
		t.Fatalf("默认分页参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("默认分页参数验证通过")
}

// TestPagesGetValidateAdContextWithCarrierDetailSelf 测试 ad_context 带完整 marketing_carrier_detail
func TestPagesGetValidateAdContextWithCarrierDetailSelf(t *testing.T) {
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.AdContext = &model.PagesGetAdContext{
		MarketingGoal:        model.MarketingGoalUserGrowth,
		MarketingCarrierType: model.MarketingCarrierTypeAppAndroid,
		MarketingTargetType:  "MARKETING_TARGET_TYPE_APP_ANDROID",
		SiteSet:              []string{model.SiteSetWechat},
		CreativeTemplateID:   1001,
		MarketingCarrierDetail: &model.PagesGetAdContextCarrierDetail{
			MarketingCarrierID:    "com.example.app",
			MarketingSubCarrierID: "sub_carrier_001",
			MarketingCarrierName:  "示例应用",
		},
	}
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整ad_context应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整ad_context验证通过")
}

// TestPagesGetValidateAdContextCarrierIDTooLongSelf 测试 marketing_carrier_id 超过2048字节
func TestPagesGetValidateAdContextCarrierIDTooLongSelf(t *testing.T) {
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.AdContext = &model.PagesGetAdContext{
		MarketingGoal:        model.MarketingGoalUserGrowth,
		MarketingCarrierType: model.MarketingCarrierTypeAppAndroid,
		MarketingTargetType:  "MARKETING_TARGET_TYPE_APP_ANDROID",
		SiteSet:              []string{model.SiteSetWechat},
		CreativeTemplateID:   1001,
		MarketingCarrierDetail: &model.PagesGetAdContextCarrierDetail{
			MarketingCarrierID: strings.Repeat("a", 2049),
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：marketing_carrier_id超过2048字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestPagesGetValidateAdContextWithMpaSpecSelf 测试 ad_context 带 mpa_spec
func TestPagesGetValidateAdContextWithMpaSpecSelf(t *testing.T) {
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.AdContext = &model.PagesGetAdContext{
		MarketingGoal:        model.MarketingGoalProductSales,
		MarketingCarrierType: model.MarketingCarrierTypeAppAndroid,
		MarketingTargetType:  "MARKETING_TARGET_TYPE_APP_ANDROID",
		SiteSet:              []string{model.SiteSetWechat},
		CreativeTemplateID:   2001,
		MpaSpec: &model.PagesGetAdContextMpaSpec{
			RecommendMethodIDs: []int64{1, 2, 3},
			ProductCatalogID:   "catalog_001",
			ProductSeriesID:    "series_001",
		},
	}
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("带mpa_spec的ad_context应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("带mpa_spec的ad_context验证通过")
}

// TestPagesGetValidateAdContextMpaSpecRecommendTooLongSelf 测试 recommend_method_ids 超过16个
func TestPagesGetValidateAdContextMpaSpecRecommendTooLongSelf(t *testing.T) {
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	ids := make([]int64, 17)
	for i := range ids {
		ids[i] = int64(i + 1)
	}
	req.AdContext = &model.PagesGetAdContext{
		MarketingGoal:        model.MarketingGoalProductSales,
		MarketingCarrierType: model.MarketingCarrierTypeAppAndroid,
		MarketingTargetType:  "MARKETING_TARGET_TYPE_APP_ANDROID",
		SiteSet:              []string{model.SiteSetWechat},
		CreativeTemplateID:   2001,
		MpaSpec: &model.PagesGetAdContextMpaSpec{
			RecommendMethodIDs: ids,
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：recommend_method_ids超过16个")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestPagesGetValidateAdContextWithMarketingAssetOuterSpecSelf 测试 ad_context 带 marketing_asset_outer_spec
func TestPagesGetValidateAdContextWithMarketingAssetOuterSpecSelf(t *testing.T) {
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.AdContext = &model.PagesGetAdContext{
		MarketingGoal:        model.MarketingGoalUserGrowth,
		MarketingCarrierType: model.MarketingCarrierTypeAppAndroid,
		MarketingTargetType:  "MARKETING_TARGET_TYPE_APP_ANDROID",
		SiteSet:              []string{model.SiteSetWechat},
		CreativeTemplateID:   1001,
		MarketingAssetOuterSpec: &model.PagesGetAdContextMarketingAssetOuterSpec{
			MarketingTargetType:      "MARKETING_TARGET_TYPE_APP_ANDROID",
			MarketingAssetOuterID:    "com.example.app",
			MarketingAssetOuterSubID: "channel_001",
			MarketingAssetOuterName:  "示例应用",
		},
	}
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("带marketing_asset_outer_spec的ad_context应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("带marketing_asset_outer_spec的ad_context验证通过")
}

// TestPagesGetValidateAdContextWithOptimizationGoalStructSelf 测试 ad_context 带 optimization_goal_struct
func TestPagesGetValidateAdContextWithOptimizationGoalStructSelf(t *testing.T) {
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.AdContext = &model.PagesGetAdContext{
		MarketingGoal:        model.MarketingGoalUserGrowth,
		MarketingCarrierType: model.MarketingCarrierTypeAppAndroid,
		MarketingTargetType:  "MARKETING_TARGET_TYPE_APP_ANDROID",
		SiteSet:              []string{model.SiteSetWechat},
		CreativeTemplateID:   1001,
		OptimizationGoalStruct: &model.PagesGetAdContextOptimizationGoalStruct{
			OptimizationGoal:               "OPTIMIZATIONGOAL_APP_DOWNLOAD",
			DeepOptimizationGoal:           "OPTIMIZATIONGOAL_APP_REGISTER",
			DeepConversionOptimizationGoal: "GOAL_7DAY_PURCHASE_ROAS",
		},
	}
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("带optimization_goal_struct的ad_context应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("带optimization_goal_struct的ad_context验证通过")
}

// TestPagesGetWithAdContextSelf 测试带完整 ad_context 的接口调用
func TestPagesGetWithAdContextSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.PagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.AdContext = &model.PagesGetAdContext{
		MarketingGoal:        model.MarketingGoalUserGrowth,
		MarketingCarrierType: model.MarketingCarrierTypeAppAndroid,
		MarketingTargetType:  "MARKETING_TARGET_TYPE_APP_ANDROID",
		SiteSet:              []string{model.SiteSetWechat},
		CreativeTemplateID:   1001,
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.PagesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}
