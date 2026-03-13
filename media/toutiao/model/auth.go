package model

import (
	"errors"
	"github.com/cngamesdk/media-sdk/model"
	"strings"
	"time"
)

type AuthReq struct {
	model.AuthReq
}

func (receiver *AuthReq) Format() {
	receiver.AuthReq.Format()
}

func (receiver *AuthReq) Validate() (err error) {
	if validateErr := receiver.AuthReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AppId <= 0 {
		err = errors.New("app_id is not exists")
		return
	}
	if receiver.RedirectUri == "" {
		err = errors.New("redirect_uri is not exists")
		return
	}
	return
}

func (receiver *AuthReq) Convert(req *model.AuthReq) {
	receiver.AuthReq = *req
	return
}

type AuthResp string

type AccessTokenReq struct {
	model.AccessTokenReq
}

func (receiver *AccessTokenReq) Convert(req *model.AccessTokenReq) {
	receiver.AccessTokenReq = *req
}

func (receiver *AccessTokenReq) Format() {
	receiver.AccessTokenReq.Format()
}

func (receiver *AccessTokenReq) Validate() (err error) {
	if validateErr := receiver.AccessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AppId <= 0 {
		err = errors.New("app_id is empty")
		return
	}
	if len(receiver.Secret) <= 0 {
		err = errors.New("secret is empty")
		return
	}
	if len(receiver.AuthCode) <= 0 {
		err = errors.New("AuthCode为空")
		return
	}
	return
}

type AccessTokenResp struct {
	model.AccessTokenResp
}

func (receiver *AccessTokenResp) Convert() (*model.AccessTokenResp, error) {
	return &receiver.AccessTokenResp, nil
}

type RefreshTokenReq struct {
	model.RefreshTokenReq
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
		err = errors.New("refresh token is empty")
		return
	}
	return
}

func (receiver *RefreshTokenReq) Convert(req *model.RefreshTokenReq) {
	receiver.RefreshTokenReq = *req
	return
}

type RefreshTokenResp struct {
	model.RefreshTokenResp
}

func (receiver *RefreshTokenResp) Convert() (resp *model.RefreshTokenResp, err error) {
	receiver.ExpireTime = time.Now().Add(time.Duration(receiver.ExpiresIn) * time.Second)
	receiver.RefreshTokenExpireTime = time.Now().Add(time.Duration(receiver.RefreshTokenExpireIn) * time.Second)
	resp = &receiver.RefreshTokenResp
	return
}

type AuthAdvertiserGetReq struct {
	AccessToken string `json:"access_token"`
}

func (a *AuthAdvertiserGetReq) Format() {
	a.AccessToken = strings.TrimSpace(a.AccessToken)
}

func (a *AuthAdvertiserGetReq) Validate() (err error) {
	if len(a.AccessToken) <= 0 {
		err = errors.New("access token is empty")
		return
	}
	return
}

type AuthAdvertiserGetResp struct {
	List []AuthAdvertiserGetRespListItem `json:"list"`
}

type AuthAdvertiserGetRespListItem struct {
	AdvertiserId   int64  `json:"advertiser_id"`
	AdvertiserName string `json:"advertiser_name"`
	AccountRole    string `json:"account_role"`
	IsValid        bool   `json:"is_valid"`
	CompanyList    []struct {
		CustomerCompanyId   int64  `json:"customer_company_id"`
		CustomerCompanyName string `json:"customer_company_name"`
	}
	AccountStringId string `json:"account_string_id"`
}
