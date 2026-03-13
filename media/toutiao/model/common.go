package model

import (
	"errors"
	"strings"
)

const (
	BaseUrlOpen = "https://open.oceanengine.com"
	BaseUrlApi  = "https://api.oceanengine.com"
	BaseUrlAd   = "https://ad.oceanengine.com"
)

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

func (a *accessTokenReq) GetHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Access-Token"] = a.AccessToken
	a.AccessToken = ""
	return headers
}

type BaseResp struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	RequestId string      `json:"request_id"`
}
