package model

import "errors"

// NativeOpenAccountReq 开启原生扩量开关请求
type NativeOpenAccountReq struct {
	accessTokenReq
	AdvertiserId      int64 `json:"advertiser_id"`       // 广告主id，必填
	OpenAccountNative int   `json:"open_account_native"` // 开启/关闭：1-开启，0-关闭
}

func (receiver *NativeOpenAccountReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *NativeOpenAccountReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.OpenAccountNative != 0 && receiver.OpenAccountNative != 1 {
		err = errors.New("open_account_native must be 0 or 1")
		return
	}
	return
}

// NativeOpenAccountResp 开启原生扩量开关响应数据（仅data部分）
type NativeOpenAccountResp struct {
	OpenAccountNative bool `json:"open_account_native"` // 操作结果：true-成功，false-失败
}
