package model

import "errors"

// WeightInfo 词提量信息
type WeightInfo struct {
	Weight float64 `json:"weight"` // 提量系数，必填；1.1-2，保留1位小数
	Word   string  `json:"word"`   // 关键词，必填
}

// AddBidWeightInfo 创建优词提量信息
type AddBidWeightInfo struct {
	CampaignId int64        `json:"campaign_id"` // 计划id
	UnitId     int64        `json:"unit_id"`     // 广告组id
	Info       []WeightInfo `json:"info"`        // 词信息，必填
}

// WordBidWeightCreateReq 创建优词提量信息请求
type WordBidWeightCreateReq struct {
	accessTokenReq
	AdvertiserId  int64              `json:"advertiser_id"`   // 广告主id，必填
	Scope         int                `json:"scope"`           // 添加范围，必填；1-账户维度，2-广告组维度
	BidWeightInfo []AddBidWeightInfo `json:"bid_weight_info"` // 优词信息，必填
}

func (receiver *WordBidWeightCreateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *WordBidWeightCreateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.Scope != 1 && receiver.Scope != 2 {
		err = errors.New("scope must be 1 or 2")
		return
	}
	if len(receiver.BidWeightInfo) == 0 {
		err = errors.New("bid_weight_info is empty")
		return
	}
	for _, item := range receiver.BidWeightInfo {
		if len(item.Info) == 0 {
			err = errors.New("bid_weight_info.info is empty")
			return
		}
	}
	return
}

// WordBidWeightCreateResp 创建优词提量信息响应数据（仅data部分）
type WordBidWeightCreateResp struct {
	WordBidWeightInfo
}
