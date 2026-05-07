package model

import "errors"

// VideoBatchDeleteReq 批量删除视频请求
type VideoBatchDeleteReq struct {
	accessTokenReq
	AdvertiserId int64    `json:"advertiser_id"` // 广告主ID，必填
	PhotoIds     []string `json:"photo_ids"`     // 视频id列表，必填
}

func (receiver *VideoBatchDeleteReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *VideoBatchDeleteReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.PhotoIds) == 0 {
		err = errors.New("photo_ids is empty")
		return
	}
	return
}

// VideoBatchDeleteResp 批量删除视频响应数据（仅data部分）
type VideoBatchDeleteResp struct {
	PhotoIds []string `json:"photo_ids"` // 删除的视频id列表
}
