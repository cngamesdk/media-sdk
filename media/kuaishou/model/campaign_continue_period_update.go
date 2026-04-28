package model

import "errors"

// CampaignContinuePeriodUpdateReq 修改周期稳投计划续投状态请求
type CampaignContinuePeriodUpdateReq struct {
	accessTokenReq
	AdvertiserId       int64 `json:"advertiser_id"`        // 广告主ID，必填，在获取 access_token 时返回
	CampaignId         int64 `json:"campaign_id"`          // 广告计划ID，必填
	ContinuePeriodType int   `json:"continue_period_type"` // 周期稳投续投开关，必填：1=关闭 2=开启
}

func (receiver *CampaignContinuePeriodUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *CampaignContinuePeriodUpdateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.CampaignId <= 0 {
		err = errors.New("campaign_id is empty")
		return
	}
	if receiver.ContinuePeriodType <= 0 {
		err = errors.New("continue_period_type is empty")
		return
	}
	return
}

// CampaignContinuePeriodUpdateResp 修改周期稳投计划续投状态响应数据（仅data部分）
type CampaignContinuePeriodUpdateResp struct {
	Id int64 `json:"id"` // 修改成功的广告计划ID
}
