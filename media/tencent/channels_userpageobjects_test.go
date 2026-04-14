package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// buildChannelsUserpageobjectsBaseReq 构建基础获取视频号动态列表请求
func buildChannelsUserpageobjectsBaseReq() *model.ChannelsUserpageobjectsGetReq {
	req := &model.ChannelsUserpageobjectsGetReq{}
	req.AccessToken = "123"
	req.AccountId = 111111
	return req
}

// ========== 获取视频号动态列表接口调用测试用例 ==========

// TestChannelsUserpageobjectsGetSelf 测试获取视频号动态列表（基础）
func TestChannelsUserpageobjectsGetSelf(t *testing.T) {
	ctx := context.Background()
	req := buildChannelsUserpageobjectsBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ChannelsUserpageobjectsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestChannelsUserpageobjectsGetWithAccountIdSelf 测试指定视频号账号 id
func TestChannelsUserpageobjectsGetWithAccountIdSelf(t *testing.T) {
	ctx := context.Background()
	req := buildChannelsUserpageobjectsBaseReq()
	req.WechatChannelsAccountId = "fake_channels_account_id"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ChannelsUserpageobjectsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestChannelsUserpageobjectsGetWithNicknameSelf 测试按视频号名称筛选
func TestChannelsUserpageobjectsGetWithNicknameSelf(t *testing.T) {
	ctx := context.Background()
	req := buildChannelsUserpageobjectsBaseReq()
	req.Nickname = "测试视频号"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ChannelsUserpageobjectsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestChannelsUserpageobjectsGetWithCountSelf 测试指定返回条数
func TestChannelsUserpageobjectsGetWithCountSelf(t *testing.T) {
	ctx := context.Background()
	req := buildChannelsUserpageobjectsBaseReq()
	req.Count = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ChannelsUserpageobjectsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestChannelsUserpageobjectsGetWithLastBufferSelf 测试连续翻页（携带 last_buffer）
func TestChannelsUserpageobjectsGetWithLastBufferSelf(t *testing.T) {
	ctx := context.Background()
	req := buildChannelsUserpageobjectsBaseReq()
	req.LastBuffer = "fake_last_buffer_token"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ChannelsUserpageobjectsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestChannelsUserpageobjectsGetWithAdContextSelf 测试携带广告上下文信息
func TestChannelsUserpageobjectsGetWithAdContextSelf(t *testing.T) {
	ctx := context.Background()
	req := buildChannelsUserpageobjectsBaseReq()
	req.AdContext = &model.ChannelsAdContext{
		MarketingGoal:        "MARKETING_GOAL_USER_GROWTH",
		MarketingCarrierType: "MARKETING_CARRIER_TYPE_WECHAT_CHANNELS",
		MarketingTargetType:  "MARKETING_TARGET_TYPE_WECHAT_CHANNELS",
		SiteSet:              []string{"SITE_SET_CHANNELS"},
		CreativeTemplateId:   10001,
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ChannelsUserpageobjectsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestChannelsUserpageobjectsGetWithAdContextFullSelf 测试携带完整广告上下文（含嵌套结构）
func TestChannelsUserpageobjectsGetWithAdContextFullSelf(t *testing.T) {
	ctx := context.Background()
	req := buildChannelsUserpageobjectsBaseReq()
	req.AdContext = &model.ChannelsAdContext{
		MarketingGoal:        "MARKETING_GOAL_USER_GROWTH",
		MarketingSubGoal:     "MARKETING_SUB_GOAL_APP_ACQUISITION",
		MarketingCarrierType: "MARKETING_CARRIER_TYPE_APP_ANDROID",
		MarketingTargetType:  "MARKETING_TARGET_TYPE_APP_ANDROID",
		MarketingCarrierDetail: &model.ChannelsAdContextMarketingCarrierDetail{
			MarketingCarrierId: "com.example.app",
		},
		SiteSet:            []string{"SITE_SET_WECHAT"},
		CreativeTemplateId: 10002,
		OptimizationGoalStruct: &model.ChannelsAdContextOptimizationGoalStruct{
			OptimizationGoal: "OPTIMIZATIONGOAL_APP_DOWNLOAD",
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ChannelsUserpageobjectsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 获取视频号动态列表参数验证测试用例 ==========

// TestChannelsUserpageobjectsGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestChannelsUserpageobjectsGetValidateMissingAccountIDSelf(t *testing.T) {
	req := buildChannelsUserpageobjectsBaseReq()
	req.AccountId = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestChannelsUserpageobjectsGetValidateCountTooHighSelf 测试 count 超出最大值
func TestChannelsUserpageobjectsGetValidateCountTooHighSelf(t *testing.T) {
	req := buildChannelsUserpageobjectsBaseReq()
	req.Count = 31
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：count须在0-30之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestChannelsUserpageobjectsGetValidateAdContextMissingMarketingGoalSelf 测试 ad_context 缺少 marketing_goal
func TestChannelsUserpageobjectsGetValidateAdContextMissingMarketingGoalSelf(t *testing.T) {
	req := buildChannelsUserpageobjectsBaseReq()
	req.AdContext = &model.ChannelsAdContext{
		MarketingCarrierType: "MARKETING_CARRIER_TYPE_WECHAT_CHANNELS",
		MarketingTargetType:  "MARKETING_TARGET_TYPE_WECHAT_CHANNELS",
		SiteSet:              []string{"SITE_SET_CHANNELS"},
		CreativeTemplateId:   10001,
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：ad_context.marketing_goal为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestChannelsUserpageobjectsGetValidateAdContextMissingSiteSetSelf 测试 ad_context 缺少 site_set
func TestChannelsUserpageobjectsGetValidateAdContextMissingSiteSetSelf(t *testing.T) {
	req := buildChannelsUserpageobjectsBaseReq()
	req.AdContext = &model.ChannelsAdContext{
		MarketingGoal:        "MARKETING_GOAL_USER_GROWTH",
		MarketingCarrierType: "MARKETING_CARRIER_TYPE_WECHAT_CHANNELS",
		MarketingTargetType:  "MARKETING_TARGET_TYPE_WECHAT_CHANNELS",
		SiteSet:              []string{},
		CreativeTemplateId:   10001,
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：ad_context.site_set为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestChannelsUserpageobjectsGetValidateAdContextMissingCreativeTemplateIdSelf 测试 ad_context 缺少 creative_template_id
func TestChannelsUserpageobjectsGetValidateAdContextMissingCreativeTemplateIdSelf(t *testing.T) {
	req := buildChannelsUserpageobjectsBaseReq()
	req.AdContext = &model.ChannelsAdContext{
		MarketingGoal:        "MARKETING_GOAL_USER_GROWTH",
		MarketingCarrierType: "MARKETING_CARRIER_TYPE_WECHAT_CHANNELS",
		MarketingTargetType:  "MARKETING_TARGET_TYPE_WECHAT_CHANNELS",
		SiteSet:              []string{"SITE_SET_CHANNELS"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：ad_context.creative_template_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestChannelsUserpageobjectsGetValidateFullParamsSelf 测试完整合法参数通过验证
func TestChannelsUserpageobjectsGetValidateFullParamsSelf(t *testing.T) {
	req := buildChannelsUserpageobjectsBaseReq()
	req.WechatChannelsAccountId = "fake_account_id"
	req.Count = 20
	req.AdContext = &model.ChannelsAdContext{
		MarketingGoal:        "MARKETING_GOAL_USER_GROWTH",
		MarketingCarrierType: "MARKETING_CARRIER_TYPE_WECHAT_CHANNELS",
		MarketingTargetType:  "MARKETING_TARGET_TYPE_WECHAT_CHANNELS",
		SiteSet:              []string{"SITE_SET_CHANNELS"},
		CreativeTemplateId:   10001,
	}
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}
