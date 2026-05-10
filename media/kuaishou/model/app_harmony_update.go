package model

import "errors"

// AppHarmonyUpdateReq 更新Harmony应用请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/update/harmony
type AppHarmonyUpdateReq struct {
	accessTokenReq
	AdvertiserId int64  `json:"advertiser_id"`       // 广告主ID，必填
	AppId        int64  `json:"app_id,omitempty"`    // 应用ID
	AppName      string `json:"app_name,omitempty"`  // 应用名称
	Developer    string `json:"developer,omitempty"` // 开发者名称
}

func (receiver *AppHarmonyUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppHarmonyUpdateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	return
}

// AppHarmonyUpdateResp 更新Harmony应用响应数据（仅data部分）
type AppHarmonyUpdateResp struct {
	AppId     int64 `json:"app_id"`     // 应用ID
	PackageId int64 `json:"package_id"` // 应用包ID
}
