package model

import (
	"errors"
	"strings"
)

type AdvertiserBudgetGetReq struct {
	accessTokenReq
	AdvertiserIds []int64 `json:"advertiser_ids,omitempty"`
}

func (a *AdvertiserBudgetGetReq) Format() {
	a.accessTokenReq.Format()
	return
}

func (a *AdvertiserBudgetGetReq) Validate() (err error) {
	if validateErr := a.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if len(a.AdvertiserIds) <= 0 {
		err = errors.New("advertiser_ids is empty")
		return
	}
	return
}

// AdvertiserBudgetGetResp 获取账户日预算响应
type AdvertiserBudgetGetResp struct {
	List []AdvertiserBudget `json:"list,omitempty"` // 预算列表
}

// AdvertiserBudget 广告主预算信息
type AdvertiserBudget struct {
	AdvertiserID int64   `json:"advertiser_id"` // 客户ID
	Budget       float64 `json:"budget"`        // 预算，单位：元；精度：小数点后两位
	BudgetMode   string  `json:"budget_mode"`   // 预算类型，详见【附录-预算类型】
}

// BudgetMode 预算类型常量
const (
	BudgetModeDay      = "BUDGET_MODE_DAY"      // 日预算
	BudgetModeInfinite = "BUDGET_MODE_INFINITE" // 不限
)

type AdvertiserBudgetUpdateReq struct {
	accessTokenReq
	AdvertiserId int64   `json:"advertiser_id,omitempty"`
	BudgetMode   string  `json:"budget_mode,omitempty"`
	Budget       float32 `json:"budget,omitempty"`
}

func (a *AdvertiserBudgetUpdateReq) Format() {
	a.accessTokenReq.Format()
	a.BudgetMode = strings.TrimSpace(a.BudgetMode)
	return
}

func (a *AdvertiserBudgetUpdateReq) Validate() (err error) {
	if validateErr := a.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if a.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(a.BudgetMode) <= 0 {
		err = errors.New("budget_mode is empty")
		return
	}
	if a.BudgetMode == BudgetModeDay {
		if a.Budget <= 0 {
			err = errors.New("budget is empty")
			return
		}
	}
	return
}

func (a *AdvertiserBudgetUpdateReq) GetHeaders() headersMap {
	header := a.accessTokenReq.GetHeaders()
	header.Json()
	return header
}

type AdvertiserBudgetUpdateResp struct {
}
