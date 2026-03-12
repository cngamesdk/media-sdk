package model

import (
	"errors"
	"github.com/cngamesdk/media-sdk/model"
	"time"
)

type AuthReq struct {
	*model.AuthReq
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
	receiver.AuthReq = req
}

type AuthResp string

type AccessTokenReq struct {
	*model.AccessTokenReq
	AppId  int64  `json:"app_id"`
	Secret string `json:"secret"`
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
	return
}

type AccessTokenResp struct {
	model.AccessTokenResp
}

func (receiver *AccessTokenResp) Convert() (*model.AccessTokenResp, error) {
	return &receiver.AccessTokenResp, nil
}

type RefreshTokenReq struct {
	*model.RefreshTokenReq
}

func (receiver *RefreshTokenReq) Format() {

}

func (receiver *RefreshTokenReq) Validate() (err error) {
	return
}

func (receiver *RefreshTokenReq) Convert() (resp interface{}, err error) {
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
