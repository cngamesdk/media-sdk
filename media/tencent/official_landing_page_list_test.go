package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// buildOfficialLandingPageListBaseReq 构建基础获取落地页列表请求
func buildOfficialLandingPageListBaseReq() *model.OfficialLandingPageListGetReq {
	req := &model.OfficialLandingPageListGetReq{}
	req.AccessToken = "123"
	req.AccountId = 111111
	return req
}

// ========== 官方落地页获取落地页列表接口调用测试用例 ==========

// TestOfficialLandingPageListGetBasicSelf 测试基本获取落地页列表（默认分页）
func TestOfficialLandingPageListGetBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageListBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageListGetWithPageSelf 测试指定分页参数
func TestOfficialLandingPageListGetWithPageSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageListBaseReq()
	req.Page = 1
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageListGetFilterByPageIdSelf 测试按 page_id 过滤
func TestOfficialLandingPageListGetFilterByPageIdSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageListBaseReq()
	req.Filtering = []*model.OfficialLandingPageListFilter{
		{
			Field:    model.OfficialLandingPageFilterFieldPageId,
			Operator: "EQUALS",
			Values:   []string{"12345678"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageListGetFilterByPageNameSelf 测试按 page_name 过滤
func TestOfficialLandingPageListGetFilterByPageNameSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageListBaseReq()
	req.Filtering = []*model.OfficialLandingPageListFilter{
		{
			Field:    model.OfficialLandingPageFilterFieldPageName,
			Operator: "EQUALS",
			Values:   []string{"测试落地页"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageListGetFilterByStatusSelf 测试按 page_status 过滤（单状态）
func TestOfficialLandingPageListGetFilterByStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageListBaseReq()
	req.Filtering = []*model.OfficialLandingPageListFilter{
		{
			Field:    model.OfficialLandingPageFilterFieldPageStatus,
			Operator: "EQUALS",
			Values:   []string{model.OfficialLandingPageStatusApproved},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageListGetFilterByMultiStatusSelf 测试按 page_status 过滤（多状态）
func TestOfficialLandingPageListGetFilterByMultiStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageListBaseReq()
	req.Filtering = []*model.OfficialLandingPageListFilter{
		{
			Field:    model.OfficialLandingPageFilterFieldPageStatus,
			Operator: "EQUALS",
			Values:   []string{model.OfficialLandingPageStatusEditing, model.OfficialLandingPageStatusPending},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 官方落地页获取落地页列表参数验证测试用例 ==========

// TestOfficialLandingPageListGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestOfficialLandingPageListGetValidateMissingAccountIDSelf(t *testing.T) {
	req := buildOfficialLandingPageListBaseReq()
	req.AccountId = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageListGetValidateInvalidPageSelf 测试非法的 page 值
func TestOfficialLandingPageListGetValidateInvalidPageSelf(t *testing.T) {
	req := buildOfficialLandingPageListBaseReq()
	req.Page = 100000
	req.Format()
	// Format 会将 page_size 默认设为 10，但不覆盖已设置的 page
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page超过最大值99999")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageListGetValidateInvalidPageSizeSelf 测试非法的 page_size 值
func TestOfficialLandingPageListGetValidateInvalidPageSizeSelf(t *testing.T) {
	req := buildOfficialLandingPageListBaseReq()
	req.PageSize = 101
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size超过最大值100")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageListGetValidateFilteringTooManySelf 测试 filtering 超过10个
func TestOfficialLandingPageListGetValidateFilteringTooManySelf(t *testing.T) {
	req := buildOfficialLandingPageListBaseReq()
	for i := 0; i < 11; i++ {
		req.Filtering = append(req.Filtering, &model.OfficialLandingPageListFilter{
			Field:    model.OfficialLandingPageFilterFieldPageStatus,
			Operator: "EQUALS",
			Values:   []string{model.OfficialLandingPageStatusEditing},
		})
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering超过10个")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageListGetValidateMissingFilterFieldSelf 测试 filtering 缺少 field
func TestOfficialLandingPageListGetValidateMissingFilterFieldSelf(t *testing.T) {
	req := buildOfficialLandingPageListBaseReq()
	req.Filtering = []*model.OfficialLandingPageListFilter{
		{
			Field:    "",
			Operator: "EQUALS",
			Values:   []string{model.OfficialLandingPageStatusEditing},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering[0].field为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageListGetValidateMissingFilterOperatorSelf 测试 filtering 缺少 operator
func TestOfficialLandingPageListGetValidateMissingFilterOperatorSelf(t *testing.T) {
	req := buildOfficialLandingPageListBaseReq()
	req.Filtering = []*model.OfficialLandingPageListFilter{
		{
			Field:    model.OfficialLandingPageFilterFieldPageStatus,
			Operator: "",
			Values:   []string{model.OfficialLandingPageStatusEditing},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering[0].operator为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageListGetValidateMissingFilterValuesSelf 测试 filtering 缺少 values
func TestOfficialLandingPageListGetValidateMissingFilterValuesSelf(t *testing.T) {
	req := buildOfficialLandingPageListBaseReq()
	req.Filtering = []*model.OfficialLandingPageListFilter{
		{
			Field:    model.OfficialLandingPageFilterFieldPageStatus,
			Operator: "EQUALS",
			Values:   []string{},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering[0].values为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageListGetValidateFullParamsSelf 测试完整合法参数通过验证
func TestOfficialLandingPageListGetValidateFullParamsSelf(t *testing.T) {
	req := buildOfficialLandingPageListBaseReq()
	req.Page = 1
	req.PageSize = 10
	req.Filtering = []*model.OfficialLandingPageListFilter{
		{
			Field:    model.OfficialLandingPageFilterFieldPageStatus,
			Operator: "EQUALS",
			Values:   []string{model.OfficialLandingPageStatusApproved},
		},
	}
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}

// TestOfficialLandingPageListGetValidateNoFilteringSelf 测试无过滤条件通过验证
func TestOfficialLandingPageListGetValidateNoFilteringSelf(t *testing.T) {
	req := buildOfficialLandingPageListBaseReq()
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("无过滤条件应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("无过滤条件验证通过")
}
