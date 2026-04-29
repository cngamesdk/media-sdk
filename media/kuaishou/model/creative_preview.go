package model

import "errors"

// CreativePreviewReq 创意体验请求
type CreativePreviewReq struct {
	accessTokenReq
	AdvertiserId int64    `json:"advertiser_id"`         // 广告主ID，必填
	UnitType     int      `json:"unit_type"`             // 广告组类型，必填：4=自定义创意 7=程序化创意2.0
	UserIds      []string `json:"user_ids,omitempty"`    // 快手ID列表，一次不超过10个；与phones二选一必填
	Phones       []string `json:"phones,omitempty"`      // 手机号列表，与user_ids二选一必填
	CreativeId   int64    `json:"creative_id,omitempty"` // 创意ID，unit_type=4时必填
	UnitId       int64    `json:"unit_id,omitempty"`     // 广告组ID，unit_type=7时必填
}

func (receiver *CreativePreviewReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *CreativePreviewReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.UnitType <= 0 {
		err = errors.New("unit_type is empty")
		return
	}
	if len(receiver.UserIds) == 0 && len(receiver.Phones) == 0 {
		err = errors.New("user_ids or phones is required")
		return
	}
	if receiver.UnitType == 4 && receiver.CreativeId <= 0 {
		err = errors.New("creative_id is required when unit_type=4")
		return
	}
	if receiver.UnitType == 7 && receiver.UnitId <= 0 {
		err = errors.New("unit_id is required when unit_type=7")
		return
	}
	return
}

// CreativePreviewResp 创意体验响应数据（仅data部分）
type CreativePreviewResp struct {
	CreativeId int64  `json:"creative_id"` // 创意ID
	UnitId     string `json:"unit_id"`     // 广告组ID
}
