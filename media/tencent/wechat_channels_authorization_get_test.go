package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// buildWechatChannelsAuthorizationGetBaseReq 构建基础获取视频号授权记录列表请求
func buildWechatChannelsAuthorizationGetBaseReq() *model.WechatChannelsAuthorizationGetReq {
	req := &model.WechatChannelsAuthorizationGetReq{}
	req.AccessToken = "123"
	req.AccountId = 111111
	return req
}

// ========== 获取视频号授权记录列表接口调用测试用例 ==========

// TestWechatChannelsAuthorizationGetSelf 测试获取视频号授权记录列表（仅 account_id）
func TestWechatChannelsAuthorizationGetSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAuthorizationGetBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAuthorizationGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAuthorizationGetWithAccountNameSelf 测试带视频号名称过滤
func TestWechatChannelsAuthorizationGetWithAccountNameSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAuthorizationGetBaseReq()
	req.WechatChannelsAccountName = "测试视频号"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAuthorizationGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAuthorizationGetWithFilteringByAccountIdSelf 测试按视频号账号 id 过滤
func TestWechatChannelsAuthorizationGetWithFilteringByAccountIdSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAuthorizationGetBaseReq()
	req.Filtering = []*model.WechatChannelsAuthorizationFilter{
		{
			Field:    model.WechatChannelsAuthorizationFilterFieldAccountId,
			Operator: "EQUALS",
			Values:   []string{"fake_wechat_channels_account_id"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAuthorizationGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAuthorizationGetWithFilteringByAuthorizationIdSelf 测试按授权 id 过滤
func TestWechatChannelsAuthorizationGetWithFilteringByAuthorizationIdSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAuthorizationGetBaseReq()
	req.Filtering = []*model.WechatChannelsAuthorizationFilter{
		{
			Field:    model.WechatChannelsAuthorizationFilterFieldAuthorizationId,
			Operator: "EQUALS",
			Values:   []string{"fake_authorization_id_001"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAuthorizationGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAuthorizationGetWithFilteringByStatusSelf 测试按授权状态过滤
func TestWechatChannelsAuthorizationGetWithFilteringByStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAuthorizationGetBaseReq()
	req.Filtering = []*model.WechatChannelsAuthorizationFilter{
		{
			Field:    model.WechatChannelsAuthorizationFilterFieldAuthorizationStatus,
			Operator: "EQUALS",
			Values:   []string{"AUTHORIZED"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAuthorizationGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAuthorizationGetWithPaginationSelf 测试带分页参数
func TestWechatChannelsAuthorizationGetWithPaginationSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAuthorizationGetBaseReq()
	req.Page = 2
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAuthorizationGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 获取视频号授权记录列表参数验证测试用例 ==========

// TestWechatChannelsAuthorizationGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestWechatChannelsAuthorizationGetValidateMissingAccountIDSelf(t *testing.T) {
	req := buildWechatChannelsAuthorizationGetBaseReq()
	req.AccountId = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAuthorizationGetValidatePageOutOfRangeSelf 测试 page 超出范围
func TestWechatChannelsAuthorizationGetValidatePageOutOfRangeSelf(t *testing.T) {
	req := buildWechatChannelsAuthorizationGetBaseReq()
	req.Page = 100000
	req.PageSize = 10
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page须在1-99999之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAuthorizationGetValidatePageSizeOutOfRangeSelf 测试 page_size 超出范围
func TestWechatChannelsAuthorizationGetValidatePageSizeOutOfRangeSelf(t *testing.T) {
	req := buildWechatChannelsAuthorizationGetBaseReq()
	req.Page = 1
	req.PageSize = 101
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size须在1-100之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAuthorizationGetValidateFilteringTooManySelf 测试 filtering 数量超限
func TestWechatChannelsAuthorizationGetValidateFilteringTooManySelf(t *testing.T) {
	req := buildWechatChannelsAuthorizationGetBaseReq()
	req.Filtering = []*model.WechatChannelsAuthorizationFilter{
		{Field: "wechat_channels_account_id", Operator: "EQUALS", Values: []string{"id1"}},
		{Field: "authorization_status", Operator: "EQUALS", Values: []string{"AUTHORIZED"}},
		{Field: "authorization_id", Operator: "EQUALS", Values: []string{"auth1"}},
		{Field: "authorization_id", Operator: "EQUALS", Values: []string{"auth2"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering数组长度不能超过3")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAuthorizationGetValidateMissingFilterFieldSelf 测试 filtering 缺少 field
func TestWechatChannelsAuthorizationGetValidateMissingFilterFieldSelf(t *testing.T) {
	req := buildWechatChannelsAuthorizationGetBaseReq()
	req.Filtering = []*model.WechatChannelsAuthorizationFilter{
		{Operator: "EQUALS", Values: []string{"fake_id"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering[0].field为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAuthorizationGetValidateFullParamsSelf 测试完整合法参数通过验证
func TestWechatChannelsAuthorizationGetValidateFullParamsSelf(t *testing.T) {
	req := buildWechatChannelsAuthorizationGetBaseReq()
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}
