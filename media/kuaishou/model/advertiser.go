package model

import "errors"

// AdvertiserInfoReq 获取广告主资质信息请求
type AdvertiserInfoReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"`
}

func (receiver *AdvertiserInfoReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AdvertiserInfoReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	return
}

// AdvertiserInfoResp 获取广告主资质信息响应数据（仅data部分）
type AdvertiserInfoResp struct {
	PrimaryIndustryName string `json:"primary_industry_name"` // 一级行业名称
	IndustryId          int64  `json:"industry_id"`           // 二级行业ID
	IndustryName        string `json:"industry_name"`         // 二级行业名称
	UserId              int64  `json:"user_id"`               // 快手账户ID
	UserName            string `json:"user_name"`             // 快手账户名称
	DeliveryType        int    `json:"delivery_type"`         // 投放方式：0默认 1优先效果
	PrimaryIndustryId   int64  `json:"primary_industry_id"`   // 一级行业ID
	EffectFirst         int    `json:"effect_first"`          // 优先效果策略生效状态：1开启
	CorporationName     string `json:"corporation_name"`      // 公司名称
	ProductName         string `json:"product_name"`          // 账户产品名称
}
