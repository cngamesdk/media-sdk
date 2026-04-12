package tencent

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// 构建基础互动落地页请求
func buildXijingPageInteractiveBaseReq() *model.XijingPageInteractiveAddReq {
	req := &model.XijingPageInteractiveAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.IsAutoSubmit = 0
	req.PageType = model.XijingTemplatePageTypeXijingAndroid
	req.InteractivePageType = model.XijingInteractivePageTypeCompressedPackage
	req.PageTitle = "测试互动标题"
	req.PageName = "测试互动落地页"
	req.MobileAppID = "1104790111"
	return req
}

// ========== 蹊径创建互动落地页接口调用测试用例 ==========

// TestXijingPageInteractiveAddBasicSelf 测试基本创建（无文件上传）
func TestXijingPageInteractiveAddBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := buildXijingPageInteractiveBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageInteractiveAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageInteractiveAddWithTransformTypeSelf 测试带转化类型
func TestXijingPageInteractiveAddWithTransformTypeSelf(t *testing.T) {
	ctx := context.Background()
	req := buildXijingPageInteractiveBaseReq()
	req.TransformType = model.XijingInteractiveTransformTypeAppDownload
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageInteractiveAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageInteractiveAddWithPageConfigSelf 测试带页面配置
func TestXijingPageInteractiveAddWithPageConfigSelf(t *testing.T) {
	ctx := context.Background()
	req := buildXijingPageInteractiveBaseReq()
	req.PageConfig = `{"key":"value"}`
	req.TransformType = model.XijingInteractiveTransformTypeWebsiteLink
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageInteractiveAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageInteractiveAddWithFileSelf 测试带文件上传
func TestXijingPageInteractiveAddWithFileSelf(t *testing.T) {
	ctx := context.Background()
	req := buildXijingPageInteractiveBaseReq()
	req.FileName = "interactive.zip"
	req.FileData = []byte("PK mock zip content")
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageInteractiveAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageInteractiveAddIosSelf 测试创建 iOS 互动落地页
func TestXijingPageInteractiveAddIosSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageInteractiveAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.IsAutoSubmit = 1
	req.PageType = model.XijingTemplatePageTypeXijingIos
	req.InteractivePageType = model.XijingInteractivePageTypeCompressedPackage
	req.PageTitle = "iOS互动标题"
	req.PageName = "iOS互动落地页"
	req.MobileAppID = "987654321"
	req.TransformType = model.XijingInteractiveTransformTypeAppDownload
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageInteractiveAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 蹊径创建互动落地页参数验证测试用例 ==========

// TestXijingPageInteractiveAddValidateMissingAccountIDSelf 测试缺少 account_id
func TestXijingPageInteractiveAddValidateMissingAccountIDSelf(t *testing.T) {
	req := buildXijingPageInteractiveBaseReq()
	req.AccountID = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageInteractiveAddValidateInvalidAutoSubmitSelf 测试 is_auto_submit 非法值
func TestXijingPageInteractiveAddValidateInvalidAutoSubmitSelf(t *testing.T) {
	req := buildXijingPageInteractiveBaseReq()
	req.IsAutoSubmit = 2
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：is_auto_submit非法")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageInteractiveAddValidateMissingPageTypeSelf 测试缺少 page_type
func TestXijingPageInteractiveAddValidateMissingPageTypeSelf(t *testing.T) {
	req := buildXijingPageInteractiveBaseReq()
	req.PageType = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageInteractiveAddValidateMissingInteractivePageTypeSelf 测试缺少 interactive_page_type
func TestXijingPageInteractiveAddValidateMissingInteractivePageTypeSelf(t *testing.T) {
	req := buildXijingPageInteractiveBaseReq()
	req.InteractivePageType = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：interactive_page_type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageInteractiveAddValidatePageTitleTooLongSelf 测试 page_title 超过20字节
func TestXijingPageInteractiveAddValidatePageTitleTooLongSelf(t *testing.T) {
	req := buildXijingPageInteractiveBaseReq()
	req.PageTitle = strings.Repeat("a", 21)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_title超过20字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageInteractiveAddValidatePageNameTooLongSelf 测试 page_name 超过20字节
func TestXijingPageInteractiveAddValidatePageNameTooLongSelf(t *testing.T) {
	req := buildXijingPageInteractiveBaseReq()
	req.PageName = strings.Repeat("a", 21)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_name超过20字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageInteractiveAddValidateMissingMobileAppIDSelf 测试缺少 mobile_app_id
func TestXijingPageInteractiveAddValidateMissingMobileAppIDSelf(t *testing.T) {
	req := buildXijingPageInteractiveBaseReq()
	req.MobileAppID = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：mobile_app_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageInteractiveAddValidateFileTooLargeSelf 测试文件超过7MB
func TestXijingPageInteractiveAddValidateFileTooLargeSelf(t *testing.T) {
	req := buildXijingPageInteractiveBaseReq()
	req.FileName = "large.zip"
	req.FileData = make([]byte, model.MaxXijingInteractiveFileSize+1)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：file超过7MB")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageInteractiveAddValidateFileNameTooLongSelf 测试文件名超过32字节
func TestXijingPageInteractiveAddValidateFileNameTooLongSelf(t *testing.T) {
	req := buildXijingPageInteractiveBaseReq()
	req.FileName = strings.Repeat("a", 33)
	req.FileData = []byte("PK content")
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：文件名超过32字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageInteractiveAddValidatePageConfigTooLongSelf 测试 page_config 超过8000字节
func TestXijingPageInteractiveAddValidatePageConfigTooLongSelf(t *testing.T) {
	req := buildXijingPageInteractiveBaseReq()
	req.PageConfig = strings.Repeat("a", 8001)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_config超过8000字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageInteractiveAddValidateFullParamsSelf 测试完整合法参数通过验证
func TestXijingPageInteractiveAddValidateFullParamsSelf(t *testing.T) {
	req := buildXijingPageInteractiveBaseReq()
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}
