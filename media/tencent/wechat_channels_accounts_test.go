package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// buildWechatChannelsAccountsBaseReq 构建基础获取视频号列表请求
func buildWechatChannelsAccountsBaseReq() *model.WechatChannelsAccountsGetReq {
	req := &model.WechatChannelsAccountsGetReq{}
	req.AccessToken = "123"
	req.AccountId = 111111
	return req
}

// ========== 获取视频号列表接口调用测试用例 ==========

// TestWechatChannelsAccountsGetSelf 测试获取视频号列表（基础）
func TestWechatChannelsAccountsGetSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAccountsBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAccountsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAccountsGetWithPageSelf 测试分页参数
func TestWechatChannelsAccountsGetWithPageSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAccountsBaseReq()
	req.Page = 2
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAccountsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAccountsGetWithSceneFeedsAdSelf 测试按场景筛选（动态推广广告）
func TestWechatChannelsAccountsGetWithSceneFeedsAdSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAccountsBaseReq()
	req.Scene = model.WechatChannelsAccountSceneFeedsAd
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAccountsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAccountsGetWithSceneLiveAdSelf 测试按场景筛选（直播广告）
func TestWechatChannelsAccountsGetWithSceneLiveAdSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAccountsBaseReq()
	req.Scene = model.WechatChannelsAccountSceneLiveAd
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAccountsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAccountsGetFilterByAccountIdSelf 测试按视频号账号 id 过滤
func TestWechatChannelsAccountsGetFilterByAccountIdSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAccountsBaseReq()
	req.Filtering = []*model.WechatChannelsAccountsFilter{
		{
			Field:    model.WechatChannelsAccountsFilterFieldAccountId,
			Operator: "EQUALS",
			Values:   []string{"fake_channels_account_id"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAccountsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAccountsGetFilterByAccountNameSelf 测试按视频号名称过滤（模糊匹配）
func TestWechatChannelsAccountsGetFilterByAccountNameSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAccountsBaseReq()
	req.Filtering = []*model.WechatChannelsAccountsFilter{
		{
			Field:    model.WechatChannelsAccountsFilterFieldAccountName,
			Operator: "CONTAINS",
			Values:   []string{"测试视频号"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAccountsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAccountsGetFilterByIsAdAcctSelf 测试按是否广告专用账户过滤
func TestWechatChannelsAccountsGetFilterByIsAdAcctSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAccountsBaseReq()
	req.Filtering = []*model.WechatChannelsAccountsFilter{
		{
			Field:    model.WechatChannelsAccountsFilterFieldIsAdAcct,
			Operator: "EQUALS",
			Values:   []string{"true"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAccountsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAccountsGetFilterByVideoIdSelf 测试按互选 video_id 过滤
func TestWechatChannelsAccountsGetFilterByVideoIdSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAccountsBaseReq()
	req.Scene = model.WechatChannelsAccountSceneFeedsCreative
	req.Filtering = []*model.WechatChannelsAccountsFilter{
		{
			Field:    model.WechatChannelsAccountsFilterFieldVideoId,
			Operator: "IN",
			Values:   []string{"video_id_1", "video_id_2"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAccountsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 获取视频号列表参数验证测试用例 ==========

// TestWechatChannelsAccountsGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestWechatChannelsAccountsGetValidateMissingAccountIDSelf(t *testing.T) {
	req := buildWechatChannelsAccountsBaseReq()
	req.AccountId = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAccountsGetValidateInvalidPageSelf 测试非法 page
func TestWechatChannelsAccountsGetValidateInvalidPageSelf(t *testing.T) {
	req := buildWechatChannelsAccountsBaseReq()
	req.Page = 100000
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page须在1-99999之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAccountsGetValidateInvalidPageSizeSelf 测试非法 page_size
func TestWechatChannelsAccountsGetValidateInvalidPageSizeSelf(t *testing.T) {
	req := buildWechatChannelsAccountsBaseReq()
	req.PageSize = 101
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size须在1-100之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAccountsGetValidateFilteringTooManySelf 测试 filtering 超过最大长度 3
func TestWechatChannelsAccountsGetValidateFilteringTooManySelf(t *testing.T) {
	req := buildWechatChannelsAccountsBaseReq()
	req.Filtering = []*model.WechatChannelsAccountsFilter{
		{Field: model.WechatChannelsAccountsFilterFieldAccountId, Operator: "EQUALS", Values: []string{"id1"}},
		{Field: model.WechatChannelsAccountsFilterFieldIsAdAcct, Operator: "EQUALS", Values: []string{"true"}},
		{Field: model.WechatChannelsAccountsFilterFieldAccountName, Operator: "EQUALS", Values: []string{"name1"}},
		{Field: model.WechatChannelsAccountsFilterFieldVideoId, Operator: "EQUALS", Values: []string{"vid1"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering数组长度不能超过3")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAccountsGetValidateFilterMissingFieldSelf 测试 filtering 缺少 field
func TestWechatChannelsAccountsGetValidateFilterMissingFieldSelf(t *testing.T) {
	req := buildWechatChannelsAccountsBaseReq()
	req.Filtering = []*model.WechatChannelsAccountsFilter{
		{Field: "", Operator: "EQUALS", Values: []string{"value1"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering[0].field为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAccountsGetValidateFilterMissingValuesSelf 测试 filtering 缺少 values
func TestWechatChannelsAccountsGetValidateFilterMissingValuesSelf(t *testing.T) {
	req := buildWechatChannelsAccountsBaseReq()
	req.Filtering = []*model.WechatChannelsAccountsFilter{
		{Field: model.WechatChannelsAccountsFilterFieldAccountId, Operator: "EQUALS", Values: []string{}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering[0].values为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAccountsGetValidateFullParamsSelf 测试完整合法参数通过验证
func TestWechatChannelsAccountsGetValidateFullParamsSelf(t *testing.T) {
	req := buildWechatChannelsAccountsBaseReq()
	req.Page = 1
	req.PageSize = 10
	req.Scene = model.WechatChannelsAccountSceneFeedsAd
	req.Filtering = []*model.WechatChannelsAccountsFilter{
		{
			Field:    model.WechatChannelsAccountsFilterFieldAccountName,
			Operator: "CONTAINS",
			Values:   []string{"测试"},
		},
	}
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}
