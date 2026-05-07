package model

import "errors"

// PhotoPackageUpdateReq 编辑素材包请求
type PhotoPackageUpdateReq struct {
	accessTokenReq
	AdvertiserId   int64    `json:"advertiser_id"`       // 广告主ID，必填
	PhotoPackageId int64    `json:"photo_package_id"`    // 素材包id，必填
	Name           string   `json:"name,omitempty"`      // 素材包名称
	PhotoIds       []string `json:"photo_ids,omitempty"` // 视频ids
}

func (receiver *PhotoPackageUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *PhotoPackageUpdateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.PhotoPackageId <= 0 {
		err = errors.New("photo_package_id is empty")
		return
	}
	return
}

// PhotoPackageUpdateResp 编辑素材包响应数据（仅data部分）
type PhotoPackageUpdateResp struct {
	PhotoPackageId int64 `json:"photo_package_id"` // 素材包id
}
