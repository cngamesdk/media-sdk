package model

import "errors"

// AtlasPhotoUploadReq 上传图文视频请求
type AtlasPhotoUploadReq struct {
	accessTokenReq
	AdvertiserId         int64    `json:"advertiser_id"`                    // 广告主ID，必填
	PicIds               []string `json:"pic_ids"`                          // 图片ID，必填，封面使用第一张图片
	AudioBsKey           string   `json:"audio_bs_key,omitempty"`           // 音频bs_key
	ShieldBackwardSwitch *bool    `json:"shield_backward_switch,omitempty"` // 上传视频后是否自动同步至快手个人主页，默认false
	WaitForTranscode     *bool    `json:"wait_for_transcode,omitempty"`     // 同步/异步上传视频，默认false
}

func (receiver *AtlasPhotoUploadReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AtlasPhotoUploadReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.PicIds) == 0 {
		err = errors.New("pic_ids is empty")
		return
	}
	return
}

// AtlasPhotoUploadResp 上传图文视频响应数据（仅data部分）
type AtlasPhotoUploadResp struct {
	PhotoId string `json:"photo_id"` // 图文视频ID
}
