package model

import "errors"

// PhotoPackageDelReq 删除素材包请求
type PhotoPackageDelReq struct {
	accessTokenReq
	AdvertiserId    int64   `json:"advertiser_id"`     // 广告主ID，必填
	PhotoPackageIds []int64 `json:"photo_package_ids"` // 素材包ids，必填
}

func (receiver *PhotoPackageDelReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *PhotoPackageDelReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.PhotoPackageIds) == 0 {
		err = errors.New("photo_package_ids is empty")
		return
	}
	return
}

// PhotoPackageDelResp 删除素材包响应数据（仅data部分）
type PhotoPackageDelResp struct {
	PhotoPackageIds []int64 `json:"photo_package_ids"` // 已删除的素材包ids
}
