package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// buildOfficialLandingPageSubmitBaseReq 构建基础送审落地页请求
func buildOfficialLandingPageSubmitBaseReq() *model.OfficialLandingPageSubmitUpdateReq {
	req := &model.OfficialLandingPageSubmitUpdateReq{}
	req.AccessToken = "123"
	req.AccountId = 111111
	req.PageId = 12345678
	return req
}

// ========== 官方落地页送审接口调用测试用例 ==========

// TestOfficialLandingPageSubmitWithPageIdSelf 测试指定 page_id 送审
func TestOfficialLandingPageSubmitWithPageIdSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageSubmitBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageSubmitUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageSubmitWithoutPageIdSelf 测试不传 page_id 送审
func TestOfficialLandingPageSubmitWithoutPageIdSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageSubmitBaseReq()
	req.PageId = 0
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageSubmitUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 官方落地页送审参数验证测试用例 ==========

// TestOfficialLandingPageSubmitValidateMissingAccountIDSelf 测试缺少 account_id
func TestOfficialLandingPageSubmitValidateMissingAccountIDSelf(t *testing.T) {
	req := buildOfficialLandingPageSubmitBaseReq()
	req.AccountId = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageSubmitValidateFullParamsSelf 测试完整合法参数通过验证
func TestOfficialLandingPageSubmitValidateFullParamsSelf(t *testing.T) {
	req := buildOfficialLandingPageSubmitBaseReq()
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}

// TestOfficialLandingPageSubmitValidateOnlyAccountIDSelf 测试仅传 account_id（page_id 可选）通过验证
func TestOfficialLandingPageSubmitValidateOnlyAccountIDSelf(t *testing.T) {
	req := buildOfficialLandingPageSubmitBaseReq()
	req.PageId = 0
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("仅传 account_id 应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("仅传 account_id 验证通过")
}
