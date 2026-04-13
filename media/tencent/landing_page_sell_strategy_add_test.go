package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// buildLandingPageSellStrategyAddBaseReq 构建基础创建售卖策略请求
func buildLandingPageSellStrategyAddBaseReq() *model.LandingPageSellStrategyAddReq {
	req := &model.LandingPageSellStrategyAddReq{}
	req.AccessToken = "123"
	req.AccountId = 111111
	req.StrategyName = "测试售卖策略"
	req.EpisodePrice = 0.10
	req.MinRechargeTier = 1.00
	req.RechargeNum = 5
	return req
}

// ========== 短剧售卖策略创建接口调用测试用例 ==========

// TestLandingPageSellStrategyAddSelf 测试创建售卖策略
func TestLandingPageSellStrategyAddSelf(t *testing.T) {
	ctx := context.Background()
	req := buildLandingPageSellStrategyAddBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LandingPageSellStrategyAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestLandingPageSellStrategyAddWithMinPriceSelf 测试最小单集价格（0.01元）
func TestLandingPageSellStrategyAddWithMinPriceSelf(t *testing.T) {
	ctx := context.Background()
	req := buildLandingPageSellStrategyAddBaseReq()
	req.EpisodePrice = 0.01
	req.MinRechargeTier = 0.01
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LandingPageSellStrategyAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestLandingPageSellStrategyAddWithZeroRechargeNumSelf 测试起充集数为0
func TestLandingPageSellStrategyAddWithZeroRechargeNumSelf(t *testing.T) {
	ctx := context.Background()
	req := buildLandingPageSellStrategyAddBaseReq()
	req.RechargeNum = 0
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LandingPageSellStrategyAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 短剧售卖策略创建参数验证测试用例 ==========

// TestLandingPageSellStrategyAddValidateMissingAccountIDSelf 测试缺少 account_id
func TestLandingPageSellStrategyAddValidateMissingAccountIDSelf(t *testing.T) {
	req := buildLandingPageSellStrategyAddBaseReq()
	req.AccountId = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestLandingPageSellStrategyAddValidateMissingStrategyNameSelf 测试缺少 strategy_name
func TestLandingPageSellStrategyAddValidateMissingStrategyNameSelf(t *testing.T) {
	req := buildLandingPageSellStrategyAddBaseReq()
	req.StrategyName = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：strategy_name长度须在1-60字节之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestLandingPageSellStrategyAddValidateStrategyNameTooLongSelf 测试 strategy_name 超长
func TestLandingPageSellStrategyAddValidateStrategyNameTooLongSelf(t *testing.T) {
	req := buildLandingPageSellStrategyAddBaseReq()
	req.StrategyName = "这个策略名称超过了六十个字节限制，用于测试名称长度验证是否生效，超长内容继续填充到足够长度超出"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：strategy_name长度须在1-60字节之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestLandingPageSellStrategyAddValidateEpisodePriceTooLowSelf 测试 episode_price 低于最小值
func TestLandingPageSellStrategyAddValidateEpisodePriceTooLowSelf(t *testing.T) {
	req := buildLandingPageSellStrategyAddBaseReq()
	req.EpisodePrice = 0.001
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：episode_price须在0.01-999999之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestLandingPageSellStrategyAddValidateEpisodePriceTooHighSelf 测试 episode_price 超出最大值
func TestLandingPageSellStrategyAddValidateEpisodePriceTooHighSelf(t *testing.T) {
	req := buildLandingPageSellStrategyAddBaseReq()
	req.EpisodePrice = 999999.01
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：episode_price须在0.01-999999之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestLandingPageSellStrategyAddValidateMinRechargeTierTooLowSelf 测试 min_recharge_tier 低于最小值
func TestLandingPageSellStrategyAddValidateMinRechargeTierTooLowSelf(t *testing.T) {
	req := buildLandingPageSellStrategyAddBaseReq()
	req.MinRechargeTier = 0.001
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：min_recharge_tier须在0.01-999999之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestLandingPageSellStrategyAddValidateRechargeNumTooHighSelf 测试 recharge_num 超出最大值
func TestLandingPageSellStrategyAddValidateRechargeNumTooHighSelf(t *testing.T) {
	req := buildLandingPageSellStrategyAddBaseReq()
	req.RechargeNum = 1000000
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：recharge_num须在0-999999之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestLandingPageSellStrategyAddValidateFullParamsSelf 测试完整合法参数通过验证
func TestLandingPageSellStrategyAddValidateFullParamsSelf(t *testing.T) {
	req := buildLandingPageSellStrategyAddBaseReq()
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}
