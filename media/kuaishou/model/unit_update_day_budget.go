package model

import "errors"

// UnitUpdateDayBudgetReq 修改广告组预算请求
type UnitUpdateDayBudgetReq struct {
	accessTokenReq
	AdvertiserId      int64   `json:"advertiser_id"`                 // 广告主ID，必填
	UnitId            int64   `json:"unit_id"`                       // 广告组ID，必填
	UnitIds           []int64 `json:"unit_ids,omitempty"`            // 广告组ID列表
	DayBudget         *int64  `json:"day_budget,omitempty"`          // 单日预算金额，单位：厘，0=不限；不能与day_budget_schedule同时传
	DayBudgetSchedule []int64 `json:"day_budget_schedule,omitempty"` // 分日预算，单位：厘，0=不限；优先级高于day_budget，不能与day_budget同时传
}

func (receiver *UnitUpdateDayBudgetReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *UnitUpdateDayBudgetReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.UnitId <= 0 {
		err = errors.New("unit_id is empty")
		return
	}
	if receiver.DayBudget == nil && len(receiver.DayBudgetSchedule) == 0 {
		err = errors.New("day_budget or day_budget_schedule is required")
		return
	}
	return
}

// UnitUpdateDayBudgetError 批量修改预算错误信息
type UnitUpdateDayBudgetError struct {
	ErrorMsg string `json:"error_msg"` // 错误信息
	Id       int64  `json:"id"`        // 广告组ID
}

// UnitUpdateDayBudgetResp 修改广告组预算响应数据（仅data部分）
type UnitUpdateDayBudgetResp struct {
	Errors  []UnitUpdateDayBudgetError `json:"errors"`   // 错误信息列表
	UnitIds []int64                    `json:"unit_ids"` // 广告组ID列表
}
