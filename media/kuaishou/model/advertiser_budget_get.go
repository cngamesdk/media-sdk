package model

import "errors"

// AdvertiserBudgetGetReq 账户日预算查询请求
type AdvertiserBudgetGetReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID，必填，在获取 access_token 的时候返回
}

func (receiver *AdvertiserBudgetGetReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AdvertiserBudgetGetReq) Validate() (err error) {
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

// AdvertiserBudgetGetResp 账户日预算查询响应数据（仅data部分）
type AdvertiserBudgetGetResp struct {
	DayBudget         int64 `json:"day_budget"`          // 单日预算，单位：厘
	DayBudgetSchedule int64 `json:"day_budget_schedule"` // 分日预算，单位：厘；与 day_budget 同时存在时优先级更高
}
