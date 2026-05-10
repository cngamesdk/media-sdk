package model

import "errors"

// AppIosCreateReq 创建iOS应用请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/create/ios
type AppIosCreateReq struct {
	accessTokenReq
	AdvertiserId   int64  `json:"advertiser_id"`             // 广告主ID，必填
	IosDownloadUrl string `json:"ios_download_url"`          // App Store下载链接，必填
	AppIconUrl     string `json:"app_icon_url,omitempty"`    // 应用图标URL(从图片上传接口获取，不填则取解析图标)
	AppId          int64  `json:"app_id,omitempty"`          // 应用ID
	PackageId      int64  `json:"package_id,omitempty"`      // 应用包ID
	AppPrivacyUrl  string `json:"app_privacy_url,omitempty"` // 隐私政策链接
}

func (receiver *AppIosCreateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppIosCreateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.IosDownloadUrl == "" {
		err = errors.New("ios_download_url is empty")
		return
	}
	return
}

// AppIosCreateResp 创建iOS应用响应数据（仅data部分）
type AppIosCreateResp struct {
	AppId       int64 `json:"app_id"`        // 应用ID
	GlobalAppId int64 `json:"global_app_id"` // 全局应用ID
	PackageId   int64 `json:"package_id"`    // 应用包ID(母包ID)
	PrivacyId   int64 `json:"privacy_id"`    // 隐私协议ID
}
