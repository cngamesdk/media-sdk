package model

import "errors"

// AppOnlineReq 应用上架请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/online
type AppOnlineReq struct {
	accessTokenReq
	AdvertiserId int64   `json:"advertiser_id"` // 广告主ID，必填
	PackageIds   []int64 `json:"package_ids"`   // 应用包ID列表，必填
}

func (receiver *AppOnlineReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppOnlineReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.PackageIds) == 0 {
		err = errors.New("package_ids is empty")
		return
	}
	return
}

// AppOnlineResp 应用上架响应数据（仅data部分）
type AppOnlineResp struct {
	Result bool `json:"result"` // 结果
}
