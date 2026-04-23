package model

import (
	"errors"
	"net/url"
	"strings"

	genericModel "github.com/cngamesdk/media-sdk/model"
)

const (
	OauthTypeAdvertiser = "advertiser"
	OauthTypeAgent      = "agent"
	OauthTypeAdSocial   = "ad_social"
	OauthTypeSeries     = "series"
)

type AuthReq struct {
	AppId       string   `json:"app_id,omitempty"`
	Scope       []string `json:"scope,omitempty"`
	RedirectUri string   `json:"redirect_uri,omitempty"`
	State       string   `json:"state,omitempty"`
	OauthType   string   `json:"oauth_type,omitempty"`
}

func (receiver *AuthReq) Format() {
	receiver.State = strings.TrimSpace(receiver.State)
	receiver.AppId = strings.TrimSpace(receiver.AppId)
	receiver.OauthType = strings.TrimSpace(receiver.OauthType)
	receiver.RedirectUri = strings.TrimSpace(receiver.RedirectUri)
	receiver.RedirectUri = url.QueryEscape(receiver.RedirectUri)
}

func (receiver *AuthReq) Validate() (err error) {
	if receiver.OauthType != OauthTypeAdvertiser &&
		receiver.OauthType != OauthTypeAgent &&
		receiver.OauthType != OauthTypeAdSocial &&
		receiver.OauthType != OauthTypeSeries {
		err = errors.New("oauth_type is not invalid." + receiver.OauthType)
		return
	}
	if len(receiver.AppId) <= 0 {
		err = errors.New("app_id is empty")
		return
	}
	if len(receiver.Scope) <= 0 {
		err = errors.New("scope is empty")
		return
	}
	if receiver.RedirectUri == "" {
		err = errors.New("redirect_uri is not exists")
		return
	}
	return
}

type AuthResp string

// AccessTokenReq 获取token请求
type AccessTokenReq struct {
	AppId    int64  `json:"app_id,omitempty"`
	Secret   string `json:"secret,omitempty"`
	AuthCode string `json:"auth_code,omitempty"`
}

func (receiver *AccessTokenReq) Format() {
	receiver.AuthCode = strings.TrimSpace(receiver.AuthCode)
	receiver.Secret = strings.TrimSpace(receiver.Secret)
}

func (receiver *AccessTokenReq) Convert(req *genericModel.AccessTokenReq) {
	receiver.AppId = req.AppId
	receiver.Secret = req.Secret
	receiver.AuthCode = req.AuthCode
}

func (receiver *AccessTokenReq) Validate() (err error) {
	if receiver.AppId <= 0 {
		err = errors.New("app_id is empty")
		return
	}
	if len(receiver.Secret) <= 0 {
		err = errors.New("secret is empty")
		return
	}
	if len(receiver.AuthCode) <= 0 {
		err = errors.New("auth_code is empty")
		return
	}
	return
}

// AccessTokenResp 获取token响应数据（仅data部分）
type AccessTokenResp struct {
	AccessToken           string `json:"access_token"`
	AccessTokenExpiresIn  int64  `json:"access_token_expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int64  `json:"refresh_token_expires_in"`
}

func (receiver *AccessTokenResp) Convert() (*genericModel.AccessTokenResp, error) {
	resp := &genericModel.AccessTokenResp{}
	resp.AccessToken = receiver.AccessToken
	resp.RefreshToken = receiver.RefreshToken
	resp.ExpiresIn = receiver.AccessTokenExpiresIn
	resp.RefreshTokenExpireIn = receiver.RefreshTokenExpiresIn
	return resp, nil
}

// RefreshTokenReq 刷新token请求
type RefreshTokenReq struct {
	AppId        int64  `json:"app_id,omitempty"`
	Secret       string `json:"secret,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (receiver *RefreshTokenReq) Format() {
	receiver.Secret = strings.TrimSpace(receiver.Secret)
	receiver.RefreshToken = strings.TrimSpace(receiver.RefreshToken)
}

func (receiver *RefreshTokenReq) Validate() (err error) {
	if receiver.AppId <= 0 {
		err = errors.New("app_id is empty")
		return
	}
	if len(receiver.Secret) <= 0 {
		err = errors.New("secret is empty")
		return
	}
	if len(receiver.RefreshToken) <= 0 {
		err = errors.New("refresh_token is empty")
		return
	}
	return
}

func (receiver *RefreshTokenReq) Convert(req *genericModel.RefreshTokenReq) {
	receiver.AppId = req.AppId
	receiver.Secret = req.Secret
	receiver.RefreshToken = req.RefreshToken
}

// RefreshTokenResp 刷新token响应数据（仅data部分）
type RefreshTokenResp struct {
	AccessToken           string `json:"access_token"`
	AccessTokenExpiresIn  int64  `json:"access_token_expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int64  `json:"refresh_token_expires_in"`
}

func (receiver *RefreshTokenResp) Convert() (*genericModel.RefreshTokenResp, error) {
	resp := &genericModel.RefreshTokenResp{}
	resp.AccessToken = receiver.AccessToken
	resp.RefreshToken = receiver.RefreshToken
	resp.ExpiresIn = receiver.AccessTokenExpiresIn
	resp.RefreshTokenExpireIn = receiver.RefreshTokenExpiresIn
	return resp, nil
}
