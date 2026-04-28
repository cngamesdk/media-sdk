package model

import "errors"

// CampaignUpdateStatusReq 修改广告计划状态请求
type CampaignUpdateStatusReq struct {
	accessTokenReq
	AdvertiserId int64   `json:"advertiser_id"`          // 广告主ID，必填，在获取 access_token 时返回
	CampaignId   int64   `json:"campaign_id,omitempty"`  // 广告计划ID，与 campaign_ids 至少填一个，最多共20个
	CampaignIds  []int64 `json:"campaign_ids,omitempty"` // 广告计划ID数组，与 campaign_id 可同时填写，不可重复，最多共20个
	PutStatus    int     `json:"put_status"`             // 投放状态，必填：1=投放 2=暂停 3=删除（删除后计划及下属广告组、创意等均被删除）
}

func (receiver *CampaignUpdateStatusReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *CampaignUpdateStatusReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.CampaignId <= 0 && len(receiver.CampaignIds) == 0 {
		err = errors.New("campaign_id and campaign_ids must have at least one")
		return
	}
	if receiver.PutStatus <= 0 {
		err = errors.New("put_status is empty")
		return
	}
	return
}

// CampaignUpdateStatusResp 修改广告计划状态响应数据（仅data部分）
type CampaignUpdateStatusResp struct {
	CampaignId  int64   `json:"campaign_id"`  // 广告计划ID
	CampaignIds []int64 `json:"campaign_ids"` // 成功修改状态的广告计划ID列表
}
