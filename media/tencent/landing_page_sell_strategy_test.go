package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// buildLandingPageSellStrategyBaseReq 构建基础获取售卖策略列表请求
func buildLandingPageSellStrategyBaseReq() *model.LandingPageSellStrategyGetReq {
	req := &model.LandingPageSellStrategyGetReq{}
	req.AccessToken = "123"
	req.AccountId = 111111
	return req
}

// ========== 短剧售卖策略获取列表接口调用测试用例 ==========

// TestLandingPageSellStrategyGetSelf 测试获取售卖策略列表（基础）
func TestLandingPageSellStrategyGetSelf(t *testing.T) {
	ctx := context.Background()
	req := buildLandingPageSellStrategyBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LandingPageSellStrategyGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestLandingPageSellStrategyGetWithStrategyIdSelf 测试按策略ID筛选
func TestLandingPageSellStrategyGetWithStrategyIdSelf(t *testing.T) {
	ctx := context.Background()
	req := buildLandingPageSellStrategyBaseReq()
	req.StrategyId = 100001
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LandingPageSellStrategyGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestLandingPageSellStrategyGetWithStrategyNameSelf 测试按策略名称模糊查询
func TestLandingPageSellStrategyGetWithStrategyNameSelf(t *testing.T) {
	ctx := context.Background()
	req := buildLandingPageSellStrategyBaseReq()
	req.StrategyName = "策略"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LandingPageSellStrategyGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestLandingPageSellStrategyGetWithFullStrategyNameSelf 测试按策略名称精准查询
func TestLandingPageSellStrategyGetWithFullStrategyNameSelf(t *testing.T) {
	ctx := context.Background()
	req := buildLandingPageSellStrategyBaseReq()
	req.FullStrategyName = "测试售卖策略"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LandingPageSellStrategyGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestLandingPageSellStrategyGetWithStatusInvalidSelf 测试按策略状态筛选（无效）
func TestLandingPageSellStrategyGetWithStatusInvalidSelf(t *testing.T) {
	ctx := context.Background()
	req := buildLandingPageSellStrategyBaseReq()
	status := model.LandingPageSellStrategyStatusInvalid
	req.StrategyStatus = &status
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LandingPageSellStrategyGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestLandingPageSellStrategyGetWithStatusValidSelf 测试按策略状态筛选（有效）
func TestLandingPageSellStrategyGetWithStatusValidSelf(t *testing.T) {
	ctx := context.Background()
	req := buildLandingPageSellStrategyBaseReq()
	status := model.LandingPageSellStrategyStatusValid
	req.StrategyStatus = &status
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LandingPageSellStrategyGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestLandingPageSellStrategyGetWithPageSelf 测试分页参数
func TestLandingPageSellStrategyGetWithPageSelf(t *testing.T) {
	ctx := context.Background()
	req := buildLandingPageSellStrategyBaseReq()
	req.Page = 2
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LandingPageSellStrategyGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 短剧售卖策略获取列表参数验证测试用例 ==========

// TestLandingPageSellStrategyGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestLandingPageSellStrategyGetValidateMissingAccountIDSelf(t *testing.T) {
	req := buildLandingPageSellStrategyBaseReq()
	req.AccountId = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestLandingPageSellStrategyGetValidateInvalidStrategyStatusSelf 测试非法策略状态
func TestLandingPageSellStrategyGetValidateInvalidStrategyStatusSelf(t *testing.T) {
	req := buildLandingPageSellStrategyBaseReq()
	status := 2
	req.StrategyStatus = &status
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：strategy_status须为0或1")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestLandingPageSellStrategyGetValidateStrategyNameTooLongSelf 测试策略名称超长（模糊查询）
func TestLandingPageSellStrategyGetValidateStrategyNameTooLongSelf(t *testing.T) {
	req := buildLandingPageSellStrategyBaseReq()
	req.StrategyName = "这个策略名称超过了六十个字节限制，用于测试模糊查询名称长度验证是否生效，超长内容继续填充到足够长度"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：strategy_name长度须在1-60字节之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestLandingPageSellStrategyGetValidateFullStrategyNameTooLongSelf 测试策略名称超长（精准查询）
func TestLandingPageSellStrategyGetValidateFullStrategyNameTooLongSelf(t *testing.T) {
	req := buildLandingPageSellStrategyBaseReq()
	req.FullStrategyName = "这个策略名称超过了六十个字节限制，用于测试精准查询名称长度验证是否生效，超长内容继续填充到足够长度"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：full_strategy_name长度须在1-60字节之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestLandingPageSellStrategyGetValidateInvalidPageSelf 测试非法 page
func TestLandingPageSellStrategyGetValidateInvalidPageSelf(t *testing.T) {
	req := buildLandingPageSellStrategyBaseReq()
	req.Page = 100000
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page须在1-99999之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestLandingPageSellStrategyGetValidateInvalidPageSizeSelf 测试非法 page_size
func TestLandingPageSellStrategyGetValidateInvalidPageSizeSelf(t *testing.T) {
	req := buildLandingPageSellStrategyBaseReq()
	req.PageSize = 101
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size须在1-100之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestLandingPageSellStrategyGetValidateFullParamsSelf 测试完整合法参数通过验证
func TestLandingPageSellStrategyGetValidateFullParamsSelf(t *testing.T) {
	req := buildLandingPageSellStrategyBaseReq()
	req.StrategyId = 100001
	req.StrategyName = "测试策略"
	req.Page = 1
	req.PageSize = 10
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}
