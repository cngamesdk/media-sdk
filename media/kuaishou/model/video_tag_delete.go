package model

import "errors"

// VideoTagDeleteReq 删除视频标签请求
type VideoTagDeleteReq struct {
	accessTokenReq
	AdvertiserId int64    `json:"advertiser_id"` // 广告主ID，必填
	PhotoTag     []string `json:"photo_tag"`     // 视频标签列表，必填，最多10个
}

func (receiver *VideoTagDeleteReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *VideoTagDeleteReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.PhotoTag) == 0 {
		err = errors.New("photo_tag is empty")
		return
	}
	return
}

// VideoTagDeleteResp 删除视频标签响应数据（仅data部分）
type VideoTagDeleteResp struct {
	PhotoTag []string `json:"photo_tag"` // 已删除的视频标签列表
}
