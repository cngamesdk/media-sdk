package model

import "errors"

// PhotoPackageRemovePhotosReq 从素材包内删除视频请求
type PhotoPackageRemovePhotosReq struct {
	accessTokenReq
	AdvertiserId   int64    `json:"advertiser_id"`    // 广告主ID，必填
	PhotoPackageId int64    `json:"photo_package_id"` // 素材包id，必填
	PhotoIds       []string `json:"photo_ids"`        // 视频ids，必填
}

func (receiver *PhotoPackageRemovePhotosReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *PhotoPackageRemovePhotosReq) Validate() (err error) {
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

// PhotoPackageRemovePhotosResp 从素材包内删除视频响应数据（仅data部分）
type PhotoPackageRemovePhotosResp struct {
	PhotoIds []string `json:"photo_ids"` // 删除的视频ids
}
