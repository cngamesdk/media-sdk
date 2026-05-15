package model

import (
	"errors"
	"strings"
)

const (
	// BaseUrlOAuth 百度营销授权链接基础URL
	BaseUrlOAuth = "https://u.baidu.com/oauth/page/index"
)

// accessTokenReq 通用access_token请求体
type accessTokenReq struct {
	AccessToken string `json:"access_token,omitempty"`
}

func (a *accessTokenReq) Format() {
	a.AccessToken = strings.TrimSpace(a.AccessToken)
}

func (a *accessTokenReq) Validate() (err error) {
	if len(a.AccessToken) <= 0 {
		err = errors.New("access token is empty")
		return
	}
	return
}

func (a *accessTokenReq) GetHeaders() headersMap {
	headers := make(headersMap)
	headers.AccessToken(a.AccessToken)
	a.AccessToken = ""
	return headers
}

type headersMap map[string]string

func (receiver headersMap) AccessToken(token string) {
	receiver["Access-Token"] = token
}

// BaseResp 百度API通用响应
type BaseResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
