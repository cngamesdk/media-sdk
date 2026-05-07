package model

import "errors"

// PhotoPackageShareReq 素材包推送请求
type PhotoPackageShareReq struct {
	accessTokenReq
	AdvertiserId       int64   `json:"advertiser_id"`        // 广告主ID，必填
	PhotoPackageIds    []int64 `json:"photo_package_ids"`    // 素材包ids，必填
	ShareAdvertiserIds []int64 `json:"share_advertiser_ids"` // 要分享的广告主ids，必填
}

func (receiver *PhotoPackageShareReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *PhotoPackageShareReq) Validate() (err error) {
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
	if len(receiver.ShareAdvertiserIds) == 0 {
		err = errors.New("share_advertiser_ids is empty")
		return
	}
	return
}

// PhotoPackageShareResp 素材包推送响应数据（仅data部分）
type PhotoPackageShareResp struct {
	SuccessPhotoPackageIds     []int64 `json:"success_photo_package_ids"`      // 成功的素材包ID
	PartSuccessPhotoPackageIds []int64 `json:"part_success_photo_package_ids"` // 部分成功的素材包ID
	FailedPhotoPackageIds      []int64 `json:"failed_photo_package_ids"`       // 失败的素材包ID
}
