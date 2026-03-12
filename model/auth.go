package model

import "time"

type AuthReq struct {
	AppId        int64  `json:"app_id"`
	State        string `json:"state"`
	Scope        []int  `json:"scope"`
	RedirectUri  string `json:"redirect_uri"`
	MaterialAuth int    `json:"material_auth"`
	AuthType     string `json:"auth_type"`
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
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResp struct {
	tokenData
}
