package model

import (
	"errors"
	"strings"
)

const (
	DevelopersUrl = "https://developers.e.kuaishou.com"
	AdUrl         = "https://ad.e.kuaishou.com"
)

type BaseResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type accessTokenReq struct {
	AccessToken string `json:"access_token,omitempty"`
}

func (a *accessTokenReq) Format() {
	a.AccessToken = strings.TrimSpace(a.AccessToken)
}

func (a *accessTokenReq) Validate() (err error) {
	if len(a.AccessToken) <= 0 {
		err = errors.New("access_token is empty")
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
