package model

import "errors"

// AdvertiserUpdateBudgetReq 修改账户预算请求
type AdvertiserUpdateBudgetReq struct {
	accessTokenReq
	AdvertiserId      int64   `json:"advertiser_id"`                 // 广告主ID，必填，在获取 access_token 的时候返回
	DayBudgetSchedule []int64 `json:"day_budget_schedule,omitempty"` // 分日预算，单位：厘；0表示不限预算（默认0）；每日最低500元，最高1亿元；优先级高于 day_budget；传空数组可清除分日预算
	DayBudget         int64   `json:"day_budget,omitempty"`          // 单日预算，单位：厘；0表示不限预算（默认0）；最低500元，最高1亿元；优先级低于 day_budget_schedule
}

func (receiver *AdvertiserUpdateBudgetReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AdvertiserUpdateBudgetReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.DayBudget <= 0 && len(receiver.DayBudgetSchedule) <= 0 {
		err = errors.New("day_budget and day_budget_schedule must select one")
		return
	}
	return
}

// AdvertiserUpdateBudgetResp 修改账户预算响应数据（仅data部分）
type AdvertiserUpdateBudgetResp struct {
}
