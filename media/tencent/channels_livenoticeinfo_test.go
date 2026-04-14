package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// buildChannelsLivenoticeinfoBaseReq 构建基础获取视频号预约直播信息请求
func buildChannelsLivenoticeinfoBaseReq() *model.ChannelsLivenoticeinfoGetReq {
	req := &model.ChannelsLivenoticeinfoGetReq{}
	req.AccessToken = "123"
	req.AccountId = 111111
	return req
}

// ========== 获取视频号预约直播信息接口调用测试用例 ==========

// TestChannelsLivenoticeinfoGetSelf 测试获取视频号预约直播信息（仅 account_id）
func TestChannelsLivenoticeinfoGetSelf(t *testing.T) {
	ctx := context.Background()
	req := buildChannelsLivenoticeinfoBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ChannelsLivenoticeinfoGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestChannelsLivenoticeinfoGetWithNicknameSelf 测试带视频号名称参数
func TestChannelsLivenoticeinfoGetWithNicknameSelf(t *testing.T) {
	ctx := context.Background()
	req := buildChannelsLivenoticeinfoBaseReq()
	req.Nickname = "测试视频号"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ChannelsLivenoticeinfoGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestChannelsLivenoticeinfoGetWithWechatChannelsAccountIdSelf 测试带视频号账号 id 参数
func TestChannelsLivenoticeinfoGetWithWechatChannelsAccountIdSelf(t *testing.T) {
	ctx := context.Background()
	req := buildChannelsLivenoticeinfoBaseReq()
	req.WechatChannelsAccountId = "fake_wechat_channels_account_id"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ChannelsLivenoticeinfoGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestChannelsLivenoticeinfoGetWithFinderUsernameSelf 测试带废弃字段 finder_username 参数
func TestChannelsLivenoticeinfoGetWithFinderUsernameSelf(t *testing.T) {
	ctx := context.Background()
	req := buildChannelsLivenoticeinfoBaseReq()
	req.FinderUsername = "fake_finder_username"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ChannelsLivenoticeinfoGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 获取视频号预约直播信息参数验证测试用例 ==========

// TestChannelsLivenoticeinfoGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestChannelsLivenoticeinfoGetValidateMissingAccountIDSelf(t *testing.T) {
	req := buildChannelsLivenoticeinfoBaseReq()
	req.AccountId = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestChannelsLivenoticeinfoGetValidateFullParamsSelf 测试完整合法参数通过验证
func TestChannelsLivenoticeinfoGetValidateFullParamsSelf(t *testing.T) {
	req := buildChannelsLivenoticeinfoBaseReq()
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}
