package model

import "errors"

// AppPrivacyDetailReq 获取应用中心隐私详情请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/privacy/detail
type AppPrivacyDetailReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID，必填
	PrivacyId    int64 `json:"privacy_id"`    // 隐私ID，必填
}

func (receiver *AppPrivacyDetailReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppPrivacyDetailReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.PrivacyId <= 0 {
		err = errors.New("privacy_id is empty")
		return
	}
	return
}

// AppPrivacyDetailResp 获取应用中心隐私详情响应数据（仅data部分）
type AppPrivacyDetailResp struct {
	AccountId int64  `json:"account_id"` // 账号ID
	PrivacyId int64  `json:"privacy_id"` // 隐私ID
	Url       string `json:"url"`        // 隐私链接
}
