package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// buildOfficialLandingPageDeleteBaseReq 构建基础删除落地页请求
func buildOfficialLandingPageDeleteBaseReq() *model.OfficialLandingPageDeleteReq {
	req := &model.OfficialLandingPageDeleteReq{}
	req.AccessToken = "123"
	req.AccountId = 111111
	req.PageId = 12345678
	return req
}

// ========== 官方落地页删除接口调用测试用例 ==========

// TestOfficialLandingPageDeleteSelf 测试删除落地页
func TestOfficialLandingPageDeleteSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageDeleteBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 官方落地页删除参数验证测试用例 ==========

// TestOfficialLandingPageDeleteValidateMissingAccountIDSelf 测试缺少 account_id
func TestOfficialLandingPageDeleteValidateMissingAccountIDSelf(t *testing.T) {
	req := buildOfficialLandingPageDeleteBaseReq()
	req.AccountId = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageDeleteValidateMissingPageIDSelf 测试缺少 page_id
func TestOfficialLandingPageDeleteValidateMissingPageIDSelf(t *testing.T) {
	req := buildOfficialLandingPageDeleteBaseReq()
	req.PageId = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageDeleteValidateFullParamsSelf 测试完整合法参数通过验证
func TestOfficialLandingPageDeleteValidateFullParamsSelf(t *testing.T) {
	req := buildOfficialLandingPageDeleteBaseReq()
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}
