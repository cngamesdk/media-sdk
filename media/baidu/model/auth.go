package model

import (
	"errors"
	"net/url"
)

const (
	// PlatformID 百度营销平台ID（固定值）
	PlatformID = "4960345965958561794"
)

// AuthorizationLinkReq 获取授权链接请求
type AuthorizationLinkReq struct {
	AppID    string `json:"app_id"`          // 应用ID，可从应用管理界面获取
	Scope    string `json:"scope"`           // 应用权限代码，建议从应用管理界面系统生成的授权链接中获取
	State    string `json:"state,omitempty"` // 开发者自定义参数，长度限制512个字符，特殊字符需要URLEncode
	Callback string `json:"callback"`        // 应用回调链接
}

// Format 格式化请求参数
func (r *AuthorizationLinkReq) Format() {
	// 对state做URLEncode（如果未编码）
	if len(r.State) > 0 {
		r.State = url.QueryEscape(r.State)
	}
}

// Validate 校验请求参数
func (r *AuthorizationLinkReq) Validate() error {
	if len(r.AppID) == 0 {
		return errors.New("app_id为必填")
	}
	if len(r.Scope) == 0 {
		return errors.New("scope为必填")
	}
	if len(r.State) > 512 {
		return errors.New("state长度不能超过512个字符")
	}
	if len(r.Callback) == 0 {
		return errors.New("callback为必填")
	}
	return nil
}

// BuildURL 构建授权链接
// 模板: https://u.baidu.com/oauth/page/index?platformId={platformId}&appId={appId}&scope={scope}&state={state}&callback={callback}
func (r *AuthorizationLinkReq) BuildURL() string {
	params := url.Values{}
	params.Set("platformId", PlatformID)
	params.Set("appId", r.AppID)
	params.Set("scope", r.Scope)
	if len(r.State) > 0 {
		params.Set("state", r.State)
	}
	params.Set("callback", r.Callback)
	return BaseUrlOAuth + "?" + params.Encode()
}

// AuthorizationLinkResp 获取授权链接响应
type AuthorizationLinkResp struct {
	AuthorizationURL string `json:"authorization_url"` // 授权链接
	PlatformID       string `json:"platform_id"`       // 平台ID
	AppID            string `json:"app_id"`            // 应用ID
	Scope            string `json:"scope"`             // 权限代码
	State            string `json:"state,omitempty"`   // 开发者自定义参数
	Callback         string `json:"callback"`          // 回调地址
}

// AccessTokenReq 换取授权令牌请求
// POST https://u.baidu.com/oauth/accessToken
type AccessTokenReq struct {
	AppID     string `json:"appId"`     // 应用appid
	AuthCode  string `json:"authCode"`  // 临时授权码（通过回调接口获取）
	SecretKey string `json:"secretKey"` // 应用的secretKey
	GrantType string `json:"grantType"` // 获取token方式，限定为：auth_code（授权码模式）
	UserID    int64  `json:"userId"`    // 授权推广账户ID（通过回调接口获取）
}

// Format 格式化请求参数
func (r *AccessTokenReq) Format() {
	if len(r.GrantType) == 0 {
		r.GrantType = "auth_code"
	}
}

// Validate 校验请求参数
func (r *AccessTokenReq) Validate() error {
	if len(r.AppID) == 0 {
		return errors.New("appId为必填")
	}
	if len(r.AuthCode) == 0 {
		return errors.New("authCode为必填")
	}
	if len(r.SecretKey) == 0 {
		return errors.New("secretKey为必填")
	}
	if len(r.GrantType) == 0 {
		return errors.New("grantType为必填")
	}
	if r.UserID <= 0 {
		return errors.New("userId为必填")
	}
	return nil
}

// AccessTokenResp 换取授权令牌响应（data字段）
type AccessTokenResp struct {
	AccessToken        string `json:"accessToken"`        // 授权令牌
	RefreshToken       string `json:"refreshToken"`       // 更新令牌
	OpenID             string `json:"openId"`             // 获取授权用户信息标识
	ExpiresTime        string `json:"expiresTime"`        // 授权令牌到期时间
	RefreshExpiresTime string `json:"refreshExpiresTime"` // 更新令牌到期时间
	ExpiresIn          int    `json:"expiresIn"`          // 授权令牌剩余有效期（秒）
	RefreshExpiresIn   int    `json:"refreshExpiresIn"`   // 更新令牌剩余有效期（秒）
	UserID             int    `json:"userId"`             // 授权账号ucid
}

// RefreshTokenReq 更新授权令牌请求
// POST https://u.baidu.com/oauth/refreshToken
type RefreshTokenReq struct {
	AppID        string `json:"appId"`        // 应用appid
	RefreshToken string `json:"refreshToken"` // 已有的更新令牌
	SecretKey    string `json:"secretKey"`    // 应用的secretKey
	UserID       int64  `json:"userId"`       // 授权推广账户ID
}

// Format 格式化请求参数
func (r *RefreshTokenReq) Format() {}

// Validate 校验请求参数
func (r *RefreshTokenReq) Validate() error {
	if len(r.AppID) == 0 {
		return errors.New("appId为必填")
	}
	if len(r.RefreshToken) == 0 {
		return errors.New("refreshToken为必填")
	}
	if len(r.SecretKey) == 0 {
		return errors.New("secretKey为必填")
	}
	if r.UserID <= 0 {
		return errors.New("userId为必填")
	}
	return nil
}

// RefreshTokenResp 更新授权令牌响应（data字段）
type RefreshTokenResp struct {
	AccessToken        string `json:"accessToken"`        // 授权令牌
	RefreshToken       string `json:"refreshToken"`       // 更新令牌
	ExpiresIn          int    `json:"expiresIn"`          // 授权令牌剩余有效期（秒）
	RefreshExpiresIn   int    `json:"refreshExpiresIn"`   // 更新令牌剩余有效期（秒）
	ExpiresTime        string `json:"expiresTime"`        // 授权令牌到期时间
	RefreshExpiresTime string `json:"refreshExpiresTime"` // 更新令牌到期时间
}
