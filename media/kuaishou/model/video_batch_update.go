package model

import "errors"

// VideoBatchUpdateReq 批量更新视频请求
type VideoBatchUpdateReq struct {
	accessTokenReq
	AdvertiserId int64    `json:"advertiser_id"`        // 广告主ID，必填
	PhotoIds     []string `json:"photo_ids"`            // 视频ids，必填，不超过100个
	PhotoName    string   `json:"photo_name,omitempty"` // 视频名称，最多100字符，与PhotoTag填其一
	PhotoTag     []string `json:"photo_tag,omitempty"`  // 视频标签，单个标签最多10字符，只支持一个标签，与PhotoName填其一
}

func (receiver *VideoBatchUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *VideoBatchUpdateReq) Validate() (err error) {
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
	if receiver.PhotoName == "" && len(receiver.PhotoTag) == 0 {
		err = errors.New("photo_name or photo_tag is required")
		return
	}
	return
}

// VideoBatchUpdateResp 批量更新视频响应数据（仅data部分）
type VideoBatchUpdateResp struct {
	PhotoIds []string `json:"photo_ids"` // 更新成功的视频ids
}
