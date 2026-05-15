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
