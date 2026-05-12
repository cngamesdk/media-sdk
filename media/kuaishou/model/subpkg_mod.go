package model

import "errors"

// SubpkgModReq 更新/恢复/删除应用分包请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/subpkg/mod
type SubpkgModReq struct {
	accessTokenReq
	AdvertiserId int64   `json:"advertiser_id"` // 广告主ID，必填
	PackageId    []int64 `json:"package_id"`    // 子包ID列表，支持批量，必填
	PutStatus    int     `json:"put_status"`    // 更改子包类型：0-更新，1-恢复，2-删除，必填
}

func (receiver *SubpkgModReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *SubpkgModReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.PackageId) == 0 {
		err = errors.New("package_id is empty")
		return
	}
	if receiver.PutStatus < 0 || receiver.PutStatus > 2 {
		err = errors.New("put_status must be 0, 1 or 2")
		return
	}
	return
}

// SubpkgModResp 更新/恢复/删除应用分包响应数据（仅data部分）
type SubpkgModResp struct {
	Result bool `json:"result"` // 修改是否成功
}
