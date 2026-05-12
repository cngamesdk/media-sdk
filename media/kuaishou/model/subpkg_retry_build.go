package model

import "errors"

// SubpkgRetryBuildReq 分包失败重新构建请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/retryBuildSubPackage
type SubpkgRetryBuildReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID，必填
	AppId        int64 `json:"app_id"`        // 应用ID，必填
}

func (receiver *SubpkgRetryBuildReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *SubpkgRetryBuildReq) Validate() (err error) {
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

// SubpkgRetryBuildResp 分包失败重新构建响应数据（仅data部分）
type SubpkgRetryBuildResp struct {
	RetryCnt int `json:"retry_cnt"` // 本次发起重新构建的应用分包数量，cnt=0表示没有需要重建的分包
}
