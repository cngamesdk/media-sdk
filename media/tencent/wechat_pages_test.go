package tencent

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ========== 获取微信落地页列表接口调用测试用例 ==========

// TestWechatPagesGetBasicSelf 测试基本查询（不带过滤条件）
func TestWechatPagesGetBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.WechatPagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatPagesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatPagesGetWithOwnerUIDSelf 测试指定 owner_uid 查询
func TestWechatPagesGetWithOwnerUIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.WechatPagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.OwnerUID = 999999
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatPagesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatPagesGetFilterByPageIDSelf 测试按 page_id 精确查询
func TestWechatPagesGetFilterByPageIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.WechatPagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.WechatPagesGetFilteringItem{
		{
			Field:    model.WechatPagesGetFilterFieldPageID,
			Operator: model.WechatPagesGetFilterOperatorEquals,
			Values:   []string{"12345678"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatPagesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatPagesGetFilterByPageNameSelf 测试按 page_name 模糊查询
func TestWechatPagesGetFilterByPageNameSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.WechatPagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.WechatPagesGetFilteringItem{
		{
			Field:    model.WechatPagesGetFilterFieldPageName,
			Operator: model.WechatPagesGetFilterOperatorContains,
			Values:   []string{"测试落地页"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatPagesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatPagesGetFilterByPageTypeSelf 测试按 page_type 查询
func TestWechatPagesGetFilterByPageTypeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.WechatPagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.WechatPagesGetFilteringItem{
		{
			Field:    model.WechatPagesGetFilterFieldPageType,
			Operator: model.WechatPagesGetFilterOperatorEquals,
			Values:   []string{model.PageTypeWechatCanvas},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatPagesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatPagesGetFilterByPageStatusSelf 测试按 page_status IN 查询
func TestWechatPagesGetFilterByPageStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.WechatPagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.WechatPagesGetFilteringItem{
		{
			Field:    model.WechatPagesGetFilterFieldPageStatus,
			Operator: model.WechatPagesGetFilterOperatorIn,
			Values:   []string{model.PageStatusNormal, model.PageStatusDeleted},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatPagesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatPagesGetFilterByMarketingGoalSelf 测试按 marketing_goal 查询
func TestWechatPagesGetFilterByMarketingGoalSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.WechatPagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.WechatPagesGetFilteringItem{
		{
			Field:    model.WechatPagesGetFilterFieldMarketingGoal,
			Operator: model.WechatPagesGetFilterOperatorEquals,
			Values:   []string{model.MarketingGoalUserGrowth},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatPagesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatPagesGetWithPaginationSelf 测试带分页参数查询
func TestWechatPagesGetWithPaginationSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.WechatPagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 2
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatPagesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 获取微信落地页列表参数验证测试用例 ==========

// TestWechatPagesGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestWechatPagesGetValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.WechatPagesGetReq{}
	req.AccessToken = "123"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatPagesGetValidateFilteringTooLongSelf 测试 filtering 超过10条
func TestWechatPagesGetValidateFilteringTooLongSelf(t *testing.T) {
	req := &model.WechatPagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	list := make([]*model.WechatPagesGetFilteringItem, 11)
	for i := range list {
		list[i] = &model.WechatPagesGetFilteringItem{
			Field:    model.WechatPagesGetFilterFieldPageType,
			Operator: model.WechatPagesGetFilterOperatorEquals,
			Values:   []string{model.PageTypeWechatCanvas},
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

// TestWechatPagesGetValidateNilFilterItemSelf 测试 filtering 中包含 nil 元素
func TestWechatPagesGetValidateNilFilterItemSelf(t *testing.T) {
	req := &model.WechatPagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.WechatPagesGetFilteringItem{
		{Field: model.WechatPagesGetFilterFieldPageType, Operator: model.WechatPagesGetFilterOperatorEquals, Values: []string{model.PageTypeWechatCanvas}},
		nil,
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering中包含nil元素")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatPagesGetValidateFilterMissingFieldSelf 测试 filtering.field 为空
func TestWechatPagesGetValidateFilterMissingFieldSelf(t *testing.T) {
	req := &model.WechatPagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.WechatPagesGetFilteringItem{
		{Operator: model.WechatPagesGetFilterOperatorEquals, Values: []string{"1"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.field为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatPagesGetValidateFilterMissingOperatorSelf 测试 filtering.operator 为空
func TestWechatPagesGetValidateFilterMissingOperatorSelf(t *testing.T) {
	req := &model.WechatPagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.WechatPagesGetFilteringItem{
		{Field: model.WechatPagesGetFilterFieldPageType, Values: []string{model.PageTypeWechatCanvas}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.operator为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatPagesGetValidateFilterEmptyValuesSelf 测试 filtering.values 为空
func TestWechatPagesGetValidateFilterEmptyValuesSelf(t *testing.T) {
	req := &model.WechatPagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.WechatPagesGetFilteringItem{
		{Field: model.WechatPagesGetFilterFieldPageType, Operator: model.WechatPagesGetFilterOperatorEquals, Values: []string{}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.values为空")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatPagesGetValidatePageNameValueTooLongSelf 测试 page_name 过滤值超过120字节
func TestWechatPagesGetValidatePageNameValueTooLongSelf(t *testing.T) {
	req := &model.WechatPagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.WechatPagesGetFilteringItem{
		{
			Field:    model.WechatPagesGetFilterFieldPageName,
			Operator: model.WechatPagesGetFilterOperatorContains,
			Values:   []string{strings.Repeat("a", 121)},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_name过滤值超过120字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatPagesGetValidateLiveNoticeIDTooLongSelf 测试 live_notice_id 过滤值超过1024字节
func TestWechatPagesGetValidateLiveNoticeIDTooLongSelf(t *testing.T) {
	req := &model.WechatPagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.WechatPagesGetFilteringItem{
		{
			Field:    model.WechatPagesGetFilterFieldLiveNoticeID,
			Operator: model.WechatPagesGetFilterOperatorEquals,
			Values:   []string{strings.Repeat("a", 1025)},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：live_notice_id过滤值超过1024字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatPagesGetValidateProductSourceTooLongSelf 测试 product_source 过滤值超过128字节
func TestWechatPagesGetValidateProductSourceTooLongSelf(t *testing.T) {
	req := &model.WechatPagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.WechatPagesGetFilteringItem{
		{
			Field:    model.WechatPagesGetFilterFieldProductSource,
			Operator: model.WechatPagesGetFilterOperatorEquals,
			Values:   []string{strings.Repeat("a", 129)},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：product_source过滤值超过128字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatPagesGetValidatePageOutOfRangeSelf 测试 page 超出范围
func TestWechatPagesGetValidatePageOutOfRangeSelf(t *testing.T) {
	req := &model.WechatPagesGetReq{}
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

// TestWechatPagesGetValidatePageSizeOutOfRangeSelf 测试 page_size 超出范围
func TestWechatPagesGetValidatePageSizeOutOfRangeSelf(t *testing.T) {
	req := &model.WechatPagesGetReq{}
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

// TestWechatPagesGetValidateDefaultPaginationSelf 测试 Format() 默认填充分页参数
func TestWechatPagesGetValidateDefaultPaginationSelf(t *testing.T) {
	req := &model.WechatPagesGetReq{}
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

// TestWechatPagesGetValidateMultipleFiltersSelf 测试同时使用多个过滤条件
func TestWechatPagesGetValidateMultipleFiltersSelf(t *testing.T) {
	req := &model.WechatPagesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.WechatPagesGetFilteringItem{
		{
			Field:    model.WechatPagesGetFilterFieldPageType,
			Operator: model.WechatPagesGetFilterOperatorEquals,
			Values:   []string{model.PageTypeWechatCanvas},
		},
		{
			Field:    model.WechatPagesGetFilterFieldMarketingGoal,
			Operator: model.WechatPagesGetFilterOperatorEquals,
			Values:   []string{model.MarketingGoalUserGrowth},
		},
	}
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("多过滤条件应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("多过滤条件验证通过")
}
