package model

import "errors"

// PhotoPackageAddPhotosReq 添加视频至素材包请求
type PhotoPackageAddPhotosReq struct {
	accessTokenReq
	AdvertiserId   int64    `json:"advertiser_id"`    // 广告主ID，必填
	PhotoPackageId int64    `json:"photo_package_id"` // 素材包id，必填
	PhotoIds       []string `json:"photo_ids"`        // 视频ids，必填
}

func (receiver *PhotoPackageAddPhotosReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *PhotoPackageAddPhotosReq) Validate() (err error) {
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
	if len(receiver.PhotoIds) == 0 {
		err = errors.New("photo_ids is empty")
		return
	}
	return
}

// PhotoPackageAddPhotosResp 添加视频至素材包响应数据（仅data部分）
type PhotoPackageAddPhotosResp struct {
	SuccessPhotoIds []string `json:"success_photo_ids"` // 成功视频ids
	DupPhotoIds     []string `json:"dup_photo_ids"`     // 重复视频ids
}
