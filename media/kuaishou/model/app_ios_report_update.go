package model

import "errors"

// AppIosReportUpdateReq iOS应用上报更新请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/ios/update
type AppIosReportUpdateReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID，必填
	AppId        int64 `json:"app_id"`        // 应用ID，必填
	IosAppId     int64 `json:"ios_app_id"`    // 苹果商店 iOS App Id，必填
	PackageId    int64 `json:"package_id"`    // 应用包ID，必填
}

func (receiver *AppIosReportUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppIosReportUpdateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.AppId <= 0 {
		err = errors.New("app_id is empty")
		return
	}
	if receiver.IosAppId <= 0 {
		err = errors.New("ios_app_id is empty")
		return
	}
	if receiver.PackageId <= 0 {
		err = errors.New("package_id is empty")
		return
	}
	return
}

// AppIosReportUpdateResp iOS应用上报更新响应数据（仅data部分）
type AppIosReportUpdateResp struct {
	Result bool `json:"result"` // 上报更新结果
}
