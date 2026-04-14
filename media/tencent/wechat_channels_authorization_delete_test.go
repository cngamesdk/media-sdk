package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// buildWechatChannelsAuthorizationDeleteBaseReq 构建基础删除视频号授权请求
func buildWechatChannelsAuthorizationDeleteBaseReq() *model.WechatChannelsAuthorizationDeleteReq {
	req := &model.WechatChannelsAuthorizationDeleteReq{}
	req.AccessToken = "123"
	req.AccountId = 111111
	return req
}

// ========== 删除视频号授权接口调用测试用例 ==========

// TestWechatChannelsAuthorizationDeleteSelf 测试删除视频号授权（仅 account_id）
func TestWechatChannelsAuthorizationDeleteSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAuthorizationDeleteBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAuthorizationDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAuthorizationDeleteByAuthorizationIdSelf 测试按授权 id 删除
func TestWechatChannelsAuthorizationDeleteByAuthorizationIdSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAuthorizationDeleteBaseReq()
	req.AuthorizationId = "fake_authorization_id_001"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAuthorizationDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAuthorizationDeleteByWechatChannelsAccountIdSelf 测试按视频号账号 id 删除
func TestWechatChannelsAuthorizationDeleteByWechatChannelsAccountIdSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAuthorizationDeleteBaseReq()
	req.WechatChannelsAccountId = "fake_wechat_channels_account_id"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAuthorizationDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAuthorizationDeleteByFinderUsernameSelf 测试按废弃字段 finder_username 删除
func TestWechatChannelsAuthorizationDeleteByFinderUsernameSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAuthorizationDeleteBaseReq()
	req.FinderUsername = "fake_finder_username"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAuthorizationDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 删除视频号授权参数验证测试用例 ==========

// TestWechatChannelsAuthorizationDeleteValidateMissingAccountIDSelf 测试缺少 account_id
func TestWechatChannelsAuthorizationDeleteValidateMissingAccountIDSelf(t *testing.T) {
	req := buildWechatChannelsAuthorizationDeleteBaseReq()
	req.AccountId = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAuthorizationDeleteValidateFullParamsSelf 测试完整合法参数通过验证
func TestWechatChannelsAuthorizationDeleteValidateFullParamsSelf(t *testing.T) {
	req := buildWechatChannelsAuthorizationDeleteBaseReq()
	req.AuthorizationId = "fake_authorization_id_001"
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}
