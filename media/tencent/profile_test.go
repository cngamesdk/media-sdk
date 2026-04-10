package tencent

import (
	"context"
	"fmt"
	"strings"
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

// ========== 删除朋友圈头像昵称跳转页测试用例 ==========

// TestProfileDeleteByProfileIDSelf 测试按 profile_id 删除跳转页
func TestProfileDeleteByProfileIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ProfileDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.ProfileID = 11111
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ProfileDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestProfileDeleteWithOrganizationIDSelf 测试携带 organization_id 删除跳转页
func TestProfileDeleteWithOrganizationIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ProfileDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.ProfileID = 11111
	req.OrganizationID = 12345
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ProfileDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 删除朋友圈头像昵称跳转页验证测试用例 ==========

// TestProfileDeleteValidateMissingProfileIDSelf 测试缺少 profile_id
func TestProfileDeleteValidateMissingProfileIDSelf(t *testing.T) {
	req := &model.ProfileDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：profile_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileDeleteValidateOrganizationIDTooLargeSelf 测试 organization_id 超过最大值
func TestProfileDeleteValidateOrganizationIDTooLargeSelf(t *testing.T) {
	req := &model.ProfileDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.ProfileID = 11111
	req.OrganizationID = 10000000000 // 超过9999999999
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：organization_id超过9999999999")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileDeleteValidateMissingAccessTokenSelf 测试缺少 access_token
func TestProfileDeleteValidateMissingAccessTokenSelf(t *testing.T) {
	req := &model.ProfileDeleteReq{}
	req.AccountID = 111111
	req.ProfileID = 11111
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：access_token为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// ========== 创建朋友圈头像昵称跳转页测试用例 ==========

// TestProfileAddDefinitionTypeSelf 测试创建自定义类型跳转页（PROFILE_TYPE_DEFINITION）
func TestProfileAddDefinitionTypeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ProfileAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.ProfileType = model.ProfileTypeDefinition
	req.HeadImageID = "abc123headimageid"
	req.ProfileName = "测试昵称"
	req.Description = "这是一段朋友圈头像昵称跳转页简介"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ProfileAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestProfileAddAutoGenerateTypeSelf 测试创建自动填充类型跳转页（PROFILE_TYPE_AUTO_GENERATE）
func TestProfileAddAutoGenerateTypeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ProfileAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.ProfileType = model.ProfileTypeAutoGenerate
	req.MarketingCarrierType = model.ProfileMarketingCarrierTypeAppAndroid
	req.MarketingCarrierID = "com.example.testapp"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ProfileAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestProfileAddWithAllFieldsSelf 测试携带全部可选字段创建跳转页
func TestProfileAddWithAllFieldsSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ProfileAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.ProfileType = model.ProfileTypeDefinition
	req.HeadImageID = "abc123headimageid"
	req.ProfileName = "全字段测试昵称"
	req.Description = "全字段测试简介内容"
	req.MarketingGoal = model.ProfileMarketingGoalUserGrowth
	req.MarketingTargetType = "MARKETING_TARGET_TYPE_APP_ANDROID"
	req.OrganizationID = 12345
	req.MdmID = 67890
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ProfileAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 创建朋友圈头像昵称跳转页验证测试用例 ==========

// TestProfileAddValidateMissingProfileTypeSelf 测试缺少 profile_type
func TestProfileAddValidateMissingProfileTypeSelf(t *testing.T) {
	req := &model.ProfileAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：profile_type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileAddValidateInvalidProfileTypeSelf 测试 profile_type 非法值
func TestProfileAddValidateInvalidProfileTypeSelf(t *testing.T) {
	req := &model.ProfileAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.ProfileType = "INVALID_TYPE"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：profile_type非法值")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileAddValidateMissingHeadImageIDSelf 测试 DEFINITION 类型缺少 head_image_id
func TestProfileAddValidateMissingHeadImageIDSelf(t *testing.T) {
	req := &model.ProfileAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.ProfileType = model.ProfileTypeDefinition
	req.ProfileName = "测试昵称"
	req.Description = "测试简介"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：DEFINITION类型head_image_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileAddValidateMissingProfileNameSelf 测试 DEFINITION 类型缺少 profile_name
func TestProfileAddValidateMissingProfileNameSelf(t *testing.T) {
	req := &model.ProfileAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.ProfileType = model.ProfileTypeDefinition
	req.HeadImageID = "abc123headimageid"
	req.Description = "测试简介"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：DEFINITION类型profile_name为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileAddValidateMissingDescriptionSelf 测试 DEFINITION 类型缺少 description
func TestProfileAddValidateMissingDescriptionSelf(t *testing.T) {
	req := &model.ProfileAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.ProfileType = model.ProfileTypeDefinition
	req.HeadImageID = "abc123headimageid"
	req.ProfileName = "测试昵称"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：DEFINITION类型description为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileAddValidateHeadImageIDTooLongSelf 测试 head_image_id 超过64字节
func TestProfileAddValidateHeadImageIDTooLongSelf(t *testing.T) {
	req := &model.ProfileAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.ProfileType = model.ProfileTypeDefinition
	req.HeadImageID = strings.Repeat("a", 65) // 超过64字节
	req.ProfileName = "测试昵称"
	req.Description = "测试简介"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：head_image_id超过64字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileAddValidateProfileNameTooLongSelf 测试 profile_name 超过30字节
func TestProfileAddValidateProfileNameTooLongSelf(t *testing.T) {
	req := &model.ProfileAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.ProfileType = model.ProfileTypeDefinition
	req.HeadImageID = "abc123headimageid"
	req.ProfileName = strings.Repeat("a", 31) // 超过30字节
	req.Description = "测试简介"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：profile_name超过30字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileAddValidateDescriptionTooLongSelf 测试 description 超过240字节
func TestProfileAddValidateDescriptionTooLongSelf(t *testing.T) {
	req := &model.ProfileAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.ProfileType = model.ProfileTypeDefinition
	req.HeadImageID = "abc123headimageid"
	req.ProfileName = "测试昵称"
	req.Description = strings.Repeat("a", 241) // 超过240字节
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：description超过240字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileAddValidateMarketingCarrierIDTooLongSelf 测试 marketing_carrier_id 超过2048字节
func TestProfileAddValidateMarketingCarrierIDTooLongSelf(t *testing.T) {
	req := &model.ProfileAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.ProfileType = model.ProfileTypeAutoGenerate
	req.MarketingCarrierID = strings.Repeat("a", 2049) // 超过2048字节
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：marketing_carrier_id超过2048字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileAddValidateOrganizationIDTooLargeSelf 测试 organization_id 超过最大值
func TestProfileAddValidateOrganizationIDTooLargeSelf(t *testing.T) {
	req := &model.ProfileAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.ProfileType = model.ProfileTypeAutoGenerate
	req.OrganizationID = 10000000000 // 超过9999999999
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：organization_id超过9999999999")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestProfileAddValidateAutoGenerateNoExtraFieldsSelf 测试 AUTO_GENERATE 类型不需要 DEFINITION 三要素
func TestProfileAddValidateAutoGenerateNoExtraFieldsSelf(t *testing.T) {
	req := &model.ProfileAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.ProfileType = model.ProfileTypeAutoGenerate
	req.Format()
	err := req.Validate()
	// AUTO_GENERATE 不需要 head_image_id/profile_name/description，验证应通过
	if err != nil {
		t.Fatalf("AUTO_GENERATE类型不应要求DEFINITION三要素，但返回了错误: %v", err)
	}
	fmt.Println("AUTO_GENERATE类型验证通过")
}
