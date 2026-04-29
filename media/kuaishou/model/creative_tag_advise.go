package model

import "errors"

// CreativeTagAdviseReq 创意标签填写建议请求
type CreativeTagAdviseReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID，必填
}

func (receiver *CreativeTagAdviseReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *CreativeTagAdviseReq) Validate() (err error) {
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

// CreativeTagAdviseItem 创意标签建议详情
type CreativeTagAdviseItem struct {
	Industry       string   `json:"industry"`       // 一级行业
	SecondIndustry string   `json:"secondIndustry"` // 二级行业
	Tags           []string `json:"tags"`           // 推荐标签
}

// CreativeTagAdviseResp 创意标签填写建议响应数据（仅data部分）
type CreativeTagAdviseResp []CreativeTagAdviseItem
