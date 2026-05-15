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

// TestAccessTokenSelf 测试换取授权令牌
func TestAccessTokenSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AccessTokenReq{
		AppID:     "test_app_id_123",
		AuthCode:  "test_auth_code_abc",
		SecretKey: "test_secret_key_xyz",
		GrantType: "auth_code",
		UserID:    123456,
	}
	resp, err := factory.AccessTokenSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestAccessTokenSelfDefaultGrantType 测试grantType默认值
func TestAccessTokenSelfDefaultGrantType(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AccessTokenReq{
		AppID:     "test_app_id_456",
		AuthCode:  "test_auth_code_def",
		SecretKey: "test_secret_key_uvw",
		UserID:    789012,
	}
	resp, err := factory.AccessTokenSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestAccessTokenSelfValidation 测试换取授权令牌参数校验
func TestAccessTokenSelfValidation(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())

	// 测试缺少appId
	req := &model.AccessTokenReq{
		AuthCode:  "test_auth_code",
		SecretKey: "test_secret_key",
		GrantType: "auth_code",
		UserID:    123456,
	}
	_, err := factory.AccessTokenSelf(ctx, req)
	if err == nil {
		t.Fatal("expected error for missing appId")
	}
	println(fmt.Sprintf("missing appId error: %s", err.Error()))

	// 测试缺少authCode
	req2 := &model.AccessTokenReq{
		AppID:     "test_app_id",
		SecretKey: "test_secret_key",
		GrantType: "auth_code",
		UserID:    123456,
	}
	_, err = factory.AccessTokenSelf(ctx, req2)
	if err == nil {
		t.Fatal("expected error for missing authCode")
	}
	println(fmt.Sprintf("missing authCode error: %s", err.Error()))

	// 测试缺少secretKey
	req3 := &model.AccessTokenReq{
		AppID:     "test_app_id",
		AuthCode:  "test_auth_code",
		GrantType: "auth_code",
		UserID:    123456,
	}
	_, err = factory.AccessTokenSelf(ctx, req3)
	if err == nil {
		t.Fatal("expected error for missing secretKey")
	}
	println(fmt.Sprintf("missing secretKey error: %s", err.Error()))

	// 测试缺少userId
	req4 := &model.AccessTokenReq{
		AppID:     "test_app_id",
		AuthCode:  "test_auth_code",
		SecretKey: "test_secret_key",
		GrantType: "auth_code",
	}
	_, err = factory.AccessTokenSelf(ctx, req4)
	if err == nil {
		t.Fatal("expected error for missing userId")
	}
	println(fmt.Sprintf("missing userId error: %s", err.Error()))
}

// TestRefreshTokenSelf 测试更新授权令牌
func TestRefreshTokenSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.RefreshTokenReq{
		AppID:        "test_app_id_123",
		RefreshToken: "test_refresh_token_abc",
		SecretKey:    "test_secret_key_xyz",
		UserID:       123456,
	}
	resp, err := factory.RefreshTokenSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestRefreshTokenSelfValidation 测试更新授权令牌参数校验
func TestRefreshTokenSelfValidation(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())

	// 测试缺少appId
	req := &model.RefreshTokenReq{
		RefreshToken: "test_refresh_token",
		SecretKey:    "test_secret_key",
		UserID:       123456,
	}
	_, err := factory.RefreshTokenSelf(ctx, req)
	if err == nil {
		t.Fatal("expected error for missing appId")
	}
	println(fmt.Sprintf("missing appId error: %s", err.Error()))

	// 测试缺少refreshToken
	req2 := &model.RefreshTokenReq{
		AppID:     "test_app_id",
		SecretKey: "test_secret_key",
		UserID:    123456,
	}
	_, err = factory.RefreshTokenSelf(ctx, req2)
	if err == nil {
		t.Fatal("expected error for missing refreshToken")
	}
	println(fmt.Sprintf("missing refreshToken error: %s", err.Error()))

	// 测试缺少secretKey
	req3 := &model.RefreshTokenReq{
		AppID:        "test_app_id",
		RefreshToken: "test_refresh_token",
		UserID:       123456,
	}
	_, err = factory.RefreshTokenSelf(ctx, req3)
	if err == nil {
		t.Fatal("expected error for missing secretKey")
	}
	println(fmt.Sprintf("missing secretKey error: %s", err.Error()))

	// 测试缺少userId
	req4 := &model.RefreshTokenReq{
		AppID:        "test_app_id",
		RefreshToken: "test_refresh_token",
		SecretKey:    "test_secret_key",
	}
	_, err = factory.RefreshTokenSelf(ctx, req4)
	if err == nil {
		t.Fatal("expected error for missing userId")
	}
	println(fmt.Sprintf("missing userId error: %s", err.Error()))
}
