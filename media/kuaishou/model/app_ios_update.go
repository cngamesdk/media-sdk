package model

import "errors"

// AppIosUpdateReq 更新iOS应用请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/update/ios
type AppIosUpdateReq struct {
	accessTokenReq
	AdvertiserId   int64  `json:"advertiser_id"`              // 广告主ID，必填
	PackageId      int64  `json:"package_id"`                 // 应用包ID，必填
	AppId          int64  `json:"app_id"`                     // 应用ID，必填
	IosDownloadUrl string `json:"ios_download_url,omitempty"` // App Store下载链接
	AppIconUrl     string `json:"app_icon_url,omitempty"`     // 应用图标(从图片上传接口获取)
}

func (receiver *AppIosUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppIosUpdateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.PackageId <= 0 {
		err = errors.New("package_id is empty")
		return
	}
	if receiver.AppId <= 0 {
		err = errors.New("app_id is empty")
		return
	}
	return
}

// AppIosUpdateResp 更新iOS应用响应数据（仅data部分）
type AppIosUpdateResp struct {
	AppId       int64 `json:"app_id"`        // 应用ID
	GlobalAppId int64 `json:"global_app_id"` // 全局应用ID
	PackageId   int64 `json:"package_id"`    // 应用包ID
	PrivacyId   int64 `json:"privacy_id"`    // 隐私声明ID
}
