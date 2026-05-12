package model

import "errors"

// AppReleaseReq 发布应用请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/release
type AppReleaseReq struct {
	accessTokenReq
	AdvertiserId int64   `json:"advertiser_id"` // 广告主ID，必填
	PackageIds   []int64 `json:"package_ids"`   // 应用包ID列表（母包ID），必填。只有安卓审核通过的母包才可以发布
}

func (receiver *AppReleaseReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppReleaseReq) Validate() (err error) {
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

// AppReleaseRespItem 发布应用响应条目
type AppReleaseRespItem struct {
	AppId          int64  `json:"app_id"`           // 应用ID
	PackageId      int64  `json:"package_id"`       // 应用包ID
	RealAppVersion string `json:"real_app_version"` // 应用完整版本信息
	VersionCode    int64  `json:"version_code"`     // 应用版本号
}

// AppReleaseResp 发布应用响应数据（仅data部分）
type AppReleaseResp []AppReleaseRespItem
