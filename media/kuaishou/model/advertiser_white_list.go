package model

import "errors"

// AdvertiserWhiteListReq 获取创意分类标签白名单客户请求
type AdvertiserWhiteListReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID，必填
}

func (receiver *AdvertiserWhiteListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AdvertiserWhiteListReq) Validate() (err error) {
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

// AdvertiserWhiteListResp 获取创意分类标签白名单客户响应数据（仅data部分）
type AdvertiserWhiteListResp struct {
	CreativeCategorySwitch int `json:"creative_category_switch"` // 账户能否使用创意标签分类：1=能使用 0=不能使用
}
