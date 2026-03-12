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

type RefreshTokenReq struct {
	model.RefreshTokenReq
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
	BaseResp
	Data model.RefreshTokenResp `json:"data"`
}

func (receiver *RefreshTokenResp) Convert() (resp *model.RefreshTokenResp, err error) {
	receiver.Data.ExpireTime = time.Now().Add(time.Duration(receiver.Data.ExpiresIn) * time.Second)
	receiver.Data.RefreshTokenExpireTime = time.Now().Add(time.Duration(receiver.Data.RefreshTokenExpireIn) * time.Second)
	resp = &receiver.Data
	return
}
