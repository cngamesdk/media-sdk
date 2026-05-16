package model

import (
	"errors"
	"strings"
)

const (
	// BaseUrlOAuth 百度营销授权链接基础URL
	BaseUrlOAuth = "https://u.baidu.com/oauth/page/index"

	// BaseUrlOAuthAPI 百度营销OAuth API基础URL
	BaseUrlOAuthAPI = "https://u.baidu.com/oauth"

	// BaseUrlAPI 百度营销业务API基础URL
	BaseUrlAPI = "https://api.baidu.com"
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

// BaseResp 百度OAuth API通用响应
type BaseResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// ApiReqHeader 业务API请求头
type ApiReqHeader struct {
	UserName    string `json:"userName"`    // 被操作账户的用户名
	AccessToken string `json:"accessToken"` // 应用授权令牌
}

// ApiRespHeader 业务API响应头
type ApiRespHeader struct {
	Status   int       `json:"status"`   // 0:成功, 1:部分失败, 2:全部失败, 3:系统错误
	Desc     string    `json:"desc"`     // 描述
	Rquota   int       `json:"rquota"`   // 剩余的请求配额
	Quota    int       `json:"quota"`    // 本次请求发送的数据条数
	Failures []Failure `json:"failures"` // 错误信息
	Oprs     int       `json:"oprs"`     // 成功操作数据条数
	Oprtime  int       `json:"oprtime"`  // 操作时间描述
}

// Failure 错误信息
type Failure struct {
	Code     int    `json:"code"`     // 错误码
	Message  string `json:"message"`  // 错误信息
	Position string `json:"position"` // 错误位置
}

// ApiReq 业务API通用请求体
type ApiReq struct {
	Header ApiReqHeader `json:"header"`
	Body   any          `json:"body"`
}

// ApiResp 业务API通用响应体
type ApiResp struct {
	Header ApiRespHeader `json:"header"`
	Body   any           `json:"body"`
}
