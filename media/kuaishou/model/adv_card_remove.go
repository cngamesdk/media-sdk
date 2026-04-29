package model

import "errors"

// AdvCardRemoveReq 删除高级创意请求
type AdvCardRemoveReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID，必填
	AdvCardId    int64 `json:"adv_card_id"`   // 卡片ID，必填
}

func (receiver *AdvCardRemoveReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AdvCardRemoveReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.AdvCardId <= 0 {
		err = errors.New("adv_card_id is empty")
		return
	}
	return
}

// AdvCardRemoveResp 删除高级创意响应数据（仅data部分）
type AdvCardRemoveResp struct {
	AdvCardId []int64 `json:"adv_card_id"` // 卡片ID数组
}
