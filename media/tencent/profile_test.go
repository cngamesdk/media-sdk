package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ========== 获取朋友圈头像昵称跳转页测试用例 ==========

// TestProfileGetByAccountIDSelf 测试按 account_id 获取跳转页列表
func TestProfileGetByAccountIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ProfileGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ProfileGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestProfileGetWithPaginationSelf 测试自定义分页参数
func TestProfileGetWithPaginationSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ProfileGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 2
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ProfileGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestProfileGetWithMaxPageSizeSelf 测试 page_size 最大值 100
func TestProfileGetWithMaxPageSizeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ProfileGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 1
	req.PageSize = 100
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ProfileGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestProfileGetFilterByProfileIDSelf 测试按 profile_id 过滤
func TestProfileGetFilterByProfileIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ProfileGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ProfileFilteringItem{
		{
			Field:    model.ProfileFilterFieldProfileID,
			Operator: "EQUALS",
			Values:   []string{"11111"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ProfileGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestProfileGetFilterByProfileTypeSelf 测试按 profile_type 过滤（自定义类型）
func TestProfileGetFilterByProfileTypeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ProfileGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ProfileFilteringItem{
		{
			Field:    model.ProfileFilterFieldProfileType,
			Operator: "EQUALS",
			Values:   []string{model.ProfileTypeDefinition},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ProfileGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestProfileGetFilterByAutoGenerateTypeSelf 测试按 profile_type 过滤（自动填充类型）
func TestProfileGetFilterByAutoGenerateTypeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ProfileGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ProfileFilteringItem{
		{
			Field:    model.ProfileFilterFieldProfileType,
			Operator: "EQUALS",
			Values:   []string{model.ProfileTypeAutoGenerate},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ProfileGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestProfileGetWithOrganizationIDSelf 测试传入 organization_id
func TestProfileGetWithOrganizationIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ProfileGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.OrganizationID = 12345
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ProfileGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 获取朋友圈头像昵称跳转页验证测试用例 ==========

// TestProfileGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestProfileGetValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.ProfileGetReq{}
	req.AccessToken = "123"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileGetValidatePageTooLargeSelf 测试 page 超过最大值 99999
func TestProfileGetValidatePageTooLargeSelf(t *testing.T) {
	req := &model.ProfileGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Format()
	req.Page = 100000
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page超过99999")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileGetValidatePageSizeTooLargeSelf 测试 page_size 超过最大值 100
func TestProfileGetValidatePageSizeTooLargeSelf(t *testing.T) {
	req := &model.ProfileGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Format()
	req.PageSize = 101
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size超过100")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileGetValidatePageTooSmallSelf 测试 page 小于最小值 1
func TestProfileGetValidatePageTooSmallSelf(t *testing.T) {
	req := &model.ProfileGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Format()
	req.Page = 0
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page小于1")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileGetValidateFilteringTooManySelf 测试 filtering 超过最大4条
func TestProfileGetValidateFilteringTooManySelf(t *testing.T) {
	req := &model.ProfileGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ProfileFilteringItem{
		{Field: model.ProfileFilterFieldProfileID, Operator: "EQUALS", Values: []string{"1"}},
		{Field: model.ProfileFilterFieldProfileType, Operator: "EQUALS", Values: []string{model.ProfileTypeDefinition}},
		{Field: model.ProfileFilterFieldMarketingGoal, Operator: "EQUALS", Values: []string{"MARKETING_GOAL_APP_PROMOTION"}},
		{Field: model.ProfileFilterFieldMarketingSubGoal, Operator: "EQUALS", Values: []string{"MARKETING_SUB_GOAL_APP_INSTALL"}},
		{Field: model.ProfileFilterFieldMarketingCarrierType, Operator: "EQUALS", Values: []string{"MARKETING_CARRIER_TYPE_APP_ANDROID"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering超过4条")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileGetValidateFilteringMissingFieldSelf 测试 filtering 缺少 field
func TestProfileGetValidateFilteringMissingFieldSelf(t *testing.T) {
	req := &model.ProfileGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ProfileFilteringItem{
		{Operator: "EQUALS", Values: []string{"11111"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering field为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileGetValidateFilteringMissingOperatorSelf 测试 filtering 缺少 operator
func TestProfileGetValidateFilteringMissingOperatorSelf(t *testing.T) {
	req := &model.ProfileGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ProfileFilteringItem{
		{Field: model.ProfileFilterFieldProfileID, Values: []string{"11111"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering operator为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileGetValidateFilteringMissingValuesSelf 测试 filtering 缺少 values
func TestProfileGetValidateFilteringMissingValuesSelf(t *testing.T) {
	req := &model.ProfileGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ProfileFilteringItem{
		{Field: model.ProfileFilterFieldProfileID, Operator: "EQUALS"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering values为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileGetDefaultPaginationSelf 测试默认分页（不传 page/page_size）
func TestProfileGetDefaultPaginationSelf(t *testing.T) {
	req := &model.ProfileGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Format()
	if req.Page != model.DefaultProfileGetPage {
		t.Fatalf("期望 page 默认值为 %d，实际为 %d", model.DefaultProfileGetPage, req.Page)
	}
	if req.PageSize != model.DefaultProfileGetPageSize {
		t.Fatalf("期望 page_size 默认值为 %d，实际为 %d", model.DefaultProfileGetPageSize, req.PageSize)
	}
	fmt.Printf("默认分页: page=%d, page_size=%d\n", req.Page, req.PageSize)
}
