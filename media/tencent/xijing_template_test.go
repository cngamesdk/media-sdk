package tencent

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ========== 获取蹊径落地页模板接口调用测试用例 ==========

// TestXijingTemplateGetBasicSelf 测试基本查询
func TestXijingTemplateGetBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingTemplateGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TemplateID = "1006"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingTemplateGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingTemplateGetAndroidSelf 测试查询 Android 模板
func TestXijingTemplateGetAndroidSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingTemplateGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TemplateID = "android_template_001"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingTemplateGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 获取蹊径落地页模板参数验证测试用例 ==========

// TestXijingTemplateGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestXijingTemplateGetValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.XijingTemplateGetReq{}
	req.AccessToken = "123"
	req.TemplateID = "1006"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingTemplateGetValidateMissingTemplateIDSelf 测试缺少 template_id
func TestXijingTemplateGetValidateMissingTemplateIDSelf(t *testing.T) {
	req := &model.XijingTemplateGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：template_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingTemplateGetValidateTemplateIDTooLongSelf 测试 template_id 超过32字节
func TestXijingTemplateGetValidateTemplateIDTooLongSelf(t *testing.T) {
	req := &model.XijingTemplateGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TemplateID = strings.Repeat("a", 33)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：template_id超过32字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingTemplateGetValidateFullParamsSelf 测试完整合法参数通过验证
func TestXijingTemplateGetValidateFullParamsSelf(t *testing.T) {
	req := &model.XijingTemplateGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TemplateID = "1006"
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}

// TestXijingTemplateGetValidateTemplateIDMaxLengthSelf 测试 template_id 恰好32字节通过验证
func TestXijingTemplateGetValidateTemplateIDMaxLengthSelf(t *testing.T) {
	req := &model.XijingTemplateGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TemplateID = strings.Repeat("a", 32)
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("32字节template_id应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("32字节template_id验证通过")
}
