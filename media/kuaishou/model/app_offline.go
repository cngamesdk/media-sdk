package model

import "errors"

// AppOfflineReq 应用下架请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/offline
type AppOfflineReq struct {
	accessTokenReq
	AdvertiserId int64   `json:"advertiser_id"` // 广告主ID，必填
	PackageIds   []int64 `json:"package_ids"`   // 应用包ID列表，必填
}

func (receiver *AppOfflineReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppOfflineReq) Validate() (err error) {
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

// AppOfflineResp 应用下架响应数据（仅data部分）
type AppOfflineResp struct {
	Result bool `json:"result"` // 结果
}
