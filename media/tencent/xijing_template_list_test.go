package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ========== 获取蹊径落地页模板列表接口调用测试用例 ==========

// TestXijingTemplateListGetBasicSelf 测试基本查询（不带可选参数）
func TestXijingTemplateListGetBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingTemplateListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingTemplateListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingTemplateListGetByTemplateIDSelf 测试按 page_template_id 查询
func TestXijingTemplateListGetByTemplateIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingTemplateListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageTemplateID = "1996"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingTemplateListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingTemplateListGetGrantSourceSelf 测试查询授权模板
func TestXijingTemplateListGetGrantSourceSelf(t *testing.T) {
	ctx := context.Background()
	isPublic := true
	req := &model.XijingTemplateListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.TemplateSource = model.XijingTemplateSourceGrant
	req.IsPublic = &isPublic
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingTemplateListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingTemplateListGetInteractionSelf 测试查询互动模板
func TestXijingTemplateListGetInteractionSelf(t *testing.T) {
	ctx := context.Background()
	isInteraction := true
	req := &model.XijingTemplateListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.IsInteraction = &isInteraction
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingTemplateListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingTemplateListGetWithPaginationSelf 测试带分页参数查询
func TestXijingTemplateListGetWithPaginationSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingTemplateListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 2
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingTemplateListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 获取蹊径落地页模板列表参数验证测试用例 ==========

// TestXijingTemplateListGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestXijingTemplateListGetValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.XijingTemplateListGetReq{}
	req.AccessToken = "123"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingTemplateListGetValidateTemplateIDTooLongSelf 测试 page_template_id 超过32字节
func TestXijingTemplateListGetValidateTemplateIDTooLongSelf(t *testing.T) {
	req := &model.XijingTemplateListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	id := ""
	for i := 0; i < 33; i++ {
		id += "a"
	}
	req.PageTemplateID = id
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_template_id超过32字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingTemplateListGetValidatePageOutOfRangeSelf 测试 page 超出范围
func TestXijingTemplateListGetValidatePageOutOfRangeSelf(t *testing.T) {
	req := &model.XijingTemplateListGetReq{}
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

// TestXijingTemplateListGetValidatePageSizeOutOfRangeSelf 测试 page_size 超出范围
func TestXijingTemplateListGetValidatePageSizeOutOfRangeSelf(t *testing.T) {
	req := &model.XijingTemplateListGetReq{}
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

// TestXijingTemplateListGetValidateDefaultPaginationSelf 测试 Format() 默认填充分页参数
func TestXijingTemplateListGetValidateDefaultPaginationSelf(t *testing.T) {
	req := &model.XijingTemplateListGetReq{}
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

// TestXijingTemplateListGetValidateFullParamsSelf 测试完整合法参数通过验证
func TestXijingTemplateListGetValidateFullParamsSelf(t *testing.T) {
	isInteraction := false
	isPublic := true
	req := &model.XijingTemplateListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageTemplateID = "1996"
	req.IsInteraction = &isInteraction
	req.IsPublic = &isPublic
	req.TemplateSource = model.XijingTemplateSourceOwner
	req.Page = 1
	req.PageSize = 10
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}
