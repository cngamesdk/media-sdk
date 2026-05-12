package model

import "errors"

// AppReleaseLatestReq 获取最新未发布应用包请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/release/latest
type AppReleaseLatestReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID，必填
	AppId        int64 `json:"app_id"`        // 应用ID，必填
}

func (receiver *AppReleaseLatestReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppReleaseLatestReq) Validate() (err error) {
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
	return
}

// AppReleaseLatestResp 获取最新未发布应用包响应数据（仅data部分）
type AppReleaseLatestResp struct {
	AppId          int64  `json:"app_id"`           // 应用ID
	PackageId      int64  `json:"package_id"`       // 应用包ID
	RealAppVersion string `json:"real_app_version"` // 应用完整版本信息
	VersionCode    int64  `json:"version_code"`     // 应用版本号
}
