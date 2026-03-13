package model

import (
	"strings"
	"time"
)

type AuthReq struct {
	AppId        int64  `json:"app_id,omitempty"`
	State        string `json:"state,omitempty"`
	Scope        []int  `json:"scope,omitempty"`
	RedirectUri  string `json:"redirect_uri,omitempty"`
	MaterialAuth int    `json:"material_auth,omitempty"`
	AuthType     string `json:"auth_type,omitempty"`
}

func (receiver *AuthReq) Format() {

}

func (receiver *AuthReq) Validate() (err error) {
	return
}

type tokenData struct {
	AccessToken            string    `json:"access_token"`
	RefreshToken           string    `json:"refresh_token"`
	ExpiresIn              int64     `json:"expires_in"`
	ExpireTime             time.Time `json:"expire_time"`
	RefreshTokenExpireIn   int64     `json:"refresh_token_expire_in"`
	RefreshTokenExpireTime time.Time `json:"refresh_token_expire_time"`
	Extension
}

type RefreshTokenReq struct {
	AppId        int64  `json:"app_id,omitempty"`
	Secret       string `json:"secret,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type RefreshTokenResp struct {
	tokenData
}

type AccessTokenReq struct {
	AppId    int64  `json:"app_id,omitempty"`
	Secret   string `json:"secret,omitempty"`
	AuthCode string `json:"auth_code,omitempty"`
}

func (receiver *AccessTokenReq) Format() {
	receiver.AuthCode = strings.TrimSpace(receiver.AuthCode)
	receiver.Secret = strings.TrimSpace(receiver.Secret)
}

func (receiver AccessTokenReq) Validate() (err error) {
	return
}

type AccessTokenResp struct {
	tokenData
}
