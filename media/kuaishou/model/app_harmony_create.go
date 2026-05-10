package model

import "errors"

// AppHarmonyCreateReq 创建Harmony应用请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/create/harmony
type AppHarmonyCreateReq struct {
	accessTokenReq
	AdvertiserId int64  `json:"advertiser_id"` // 广告主ID，必填
	AppIconUrl   string `json:"app_icon_url"`  // 应用图标(从图片上传接口获取)，必填
	AppName      string `json:"app_name"`      // 应用名称，必填
	Developer    string `json:"developer"`     // 开发者名称(需与软著保持一致)，必填
	PackageName  string `json:"package_name"`  // 应用包名，必填
}

func (receiver *AppHarmonyCreateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppHarmonyCreateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.AppIconUrl == "" {
		err = errors.New("app_icon_url is empty")
		return
	}
	if receiver.AppName == "" {
		err = errors.New("app_name is empty")
		return
	}
	if receiver.Developer == "" {
		err = errors.New("developer is empty")
		return
	}
	if receiver.PackageName == "" {
		err = errors.New("package_name is empty")
		return
	}
	return
}

// AppHarmonyCreateResp 创建Harmony应用响应数据（仅data部分）
type AppHarmonyCreateResp struct {
	AppId       int64 `json:"app_id"`        // 应用ID
	GlobalAppId int64 `json:"global_app_id"` // 全局应用ID
	PackageId   int64 `json:"package_id"`    // 应用包ID(母包ID)
}
