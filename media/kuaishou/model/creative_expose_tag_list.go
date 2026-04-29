package model

import "errors"

// CreativeExposeTagListReq 查询创意推荐理由请求
type CreativeExposeTagListReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"`           // 广告主ID，必填
	CampaignType int   `json:"campaign_type,omitempty"` // 计划类型
}

func (receiver *CreativeExposeTagListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *CreativeExposeTagListReq) Validate() (err error) {
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

// CreativeExposeTagItem 推荐理由项
type CreativeExposeTagItem struct {
	Text string `json:"text"` // 推荐理由
}

// CreativeExposeTagListResp 查询创意推荐理由响应数据（仅data部分）
type CreativeExposeTagListResp struct {
	Details []CreativeExposeTagItem `json:"details"` // 推荐理由列表
}
