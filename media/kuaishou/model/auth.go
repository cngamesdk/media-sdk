package model

import (
	"errors"
	"net/url"
	"strings"
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
