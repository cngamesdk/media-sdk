package baidu

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/baidu/model"
)

// TestAuthorizationLinkSelf 测试获取授权链接
func TestAuthorizationLinkSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AuthorizationLinkReq{
		AppID:    "test_app_id_123",
		Scope:    "BASIC_INFO,AD_MANAGEMENT",
		State:    "custom_state_value",
		Callback: "https://www.example.com/oauth/callback",
	}
	resp, err := factory.AuthorizationLinkSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
	println(fmt.Sprintf("authorization url: %s", resp.AuthorizationURL))
}

// TestAuthorizationLinkSelfWithoutState 测试无state参数的授权链接
func TestAuthorizationLinkSelfWithoutState(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AuthorizationLinkReq{
		AppID:    "test_app_id_456",
		Scope:    "BASIC_INFO",
		Callback: "https://www.example.com/oauth/callback",
	}
	resp, err := factory.AuthorizationLinkSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
	println(fmt.Sprintf("authorization url: %s", resp.AuthorizationURL))
}

// TestAuthorizationLinkSelfValidation 测试参数校验
func TestAuthorizationLinkSelfValidation(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())

	// 测试缺少app_id
	req := &model.AuthorizationLinkReq{
		Scope:    "BASIC_INFO",
		Callback: "https://www.example.com/oauth/callback",
	}
	_, err := factory.AuthorizationLinkSelf(ctx, req)
	if err == nil {
		t.Fatal("expected error for missing app_id")
	}
	println(fmt.Sprintf("missing app_id error: %s", err.Error()))

	// 测试缺少scope
	req2 := &model.AuthorizationLinkReq{
		AppID:    "test_app_id",
		Callback: "https://www.example.com/oauth/callback",
	}
	_, err = factory.AuthorizationLinkSelf(ctx, req2)
	if err == nil {
		t.Fatal("expected error for missing scope")
	}
	println(fmt.Sprintf("missing scope error: %s", err.Error()))

	// 测试缺少callback
	req3 := &model.AuthorizationLinkReq{
		AppID: "test_app_id",
		Scope: "BASIC_INFO",
	}
	_, err = factory.AuthorizationLinkSelf(ctx, req3)
	if err == nil {
		t.Fatal("expected error for missing callback")
	}
	println(fmt.Sprintf("missing callback error: %s", err.Error()))
}

// TestAuthorizationLinkSelfStateLength 测试state长度限制
func TestAuthorizationLinkSelfStateLength(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())

	// 构造超过512字符的state
	longState := ""
	for i := 0; i < 513; i++ {
		longState += "a"
	}
	req := &model.AuthorizationLinkReq{
		AppID:    "test_app_id",
		Scope:    "BASIC_INFO",
		State:    longState,
		Callback: "https://www.example.com/oauth/callback",
	}
	_, err := factory.AuthorizationLinkSelf(ctx, req)
	if err == nil {
		t.Fatal("expected error for state too long")
	}
	println(fmt.Sprintf("state too long error: %s", err.Error()))
}
