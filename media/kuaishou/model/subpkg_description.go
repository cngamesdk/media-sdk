package model

import "errors"

// SubpkgDescriptionReq 修改应用分包备注请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/subpkg/description
type SubpkgDescriptionReq struct {
	accessTokenReq
	AdvertiserId int64  `json:"advertiser_id"` // 广告主ID，必填
	PackageId    int64  `json:"package_id"`    // 子包ID，必填
	Description  string `json:"description"`   // 子包描述，必填
}

func (receiver *SubpkgDescriptionReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *SubpkgDescriptionReq) Validate() (err error) {
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
	if receiver.Description == "" {
		err = errors.New("description is empty")
		return
	}
	return
}

// SubpkgDescriptionResp 修改应用分包备注响应数据（仅data部分）
type SubpkgDescriptionResp struct {
	Result bool `json:"result"` // 修改分包备注是否成功
}
