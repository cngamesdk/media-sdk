package tencent

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ========== 获取蹊径落地页互动模板配置接口调用测试用例 ==========

// TestXijingComplexTemplateGetBasicSelf 测试基本查询
func TestXijingComplexTemplateGetBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingComplexTemplateGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageTemplateID = "1739"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingComplexTemplateGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingComplexTemplateGetByPrivateIDSelf 测试查询私有模板
func TestXijingComplexTemplateGetByPrivateIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingComplexTemplateGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageTemplateID = "private_template_001"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingComplexTemplateGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 获取蹊径落地页互动模板配置参数验证测试用例 ==========

// TestXijingComplexTemplateGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestXijingComplexTemplateGetValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.XijingComplexTemplateGetReq{}
	req.AccessToken = "123"
	req.PageTemplateID = "1739"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingComplexTemplateGetValidateMissingTemplateIDSelf 测试缺少 page_template_id
func TestXijingComplexTemplateGetValidateMissingTemplateIDSelf(t *testing.T) {
	req := &model.XijingComplexTemplateGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_template_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingComplexTemplateGetValidateTemplateIDTooLongSelf 测试 page_template_id 超过32字节
func TestXijingComplexTemplateGetValidateTemplateIDTooLongSelf(t *testing.T) {
	req := &model.XijingComplexTemplateGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageTemplateID = strings.Repeat("a", 33)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_template_id超过32字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingComplexTemplateGetValidateTemplateIDMaxLengthSelf 测试 page_template_id 恰好32字节通过验证
func TestXijingComplexTemplateGetValidateTemplateIDMaxLengthSelf(t *testing.T) {
	req := &model.XijingComplexTemplateGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageTemplateID = strings.Repeat("a", 32)
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("32字节page_template_id应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("32字节page_template_id验证通过")
}

// TestXijingComplexTemplateGetValidateFullParamsSelf 测试完整合法参数通过验证
func TestXijingComplexTemplateGetValidateFullParamsSelf(t *testing.T) {
	req := &model.XijingComplexTemplateGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageTemplateID = "1739"
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}
