package model

import "errors"

// AdvCardShareReq 推送高级创意请求
type AdvCardShareReq struct {
	accessTokenReq
	AdvertiserId       int64   `json:"advertiser_id"`        // 广告主ID，必填
	ShareAdvertiserIds []int64 `json:"share_advertiser_ids"` // 推送目标账户ID列表，必填
	AdvCardIds         []int64 `json:"adv_card_ids"`         // 推送卡片ID列表，必填
}

func (receiver *AdvCardShareReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AdvCardShareReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.ShareAdvertiserIds) == 0 {
		err = errors.New("share_advertiser_ids is empty")
		return
	}
	if len(receiver.AdvCardIds) == 0 {
		err = errors.New("adv_card_ids is empty")
		return
	}
	return
}

// AdvCardShareItem 推送高级创意响应单条项
type AdvCardShareItem struct {
	AdvertiserId int64 `json:"advertiser_id"` // 分享后的账户ID
	AdvCardId    int64 `json:"adv_card_id"`   // 分享后的卡片ID
}

// AdvCardShareResp 推送高级创意响应数据（仅data部分）
type AdvCardShareResp []AdvCardShareItem
