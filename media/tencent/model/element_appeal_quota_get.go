package model

import "errors"

// ========== 获取元素申诉复审配额 ==========
// https://developers.e.qq.com/v3.0/docs/api/element_appeal_quota/get

// ElementAppealQuotaGetReq 获取元素申诉复审配额请求
type ElementAppealQuotaGetReq struct {
	GlobalReq
	AccountID int64 `json:"account_id"` // 推广帐号 id，有操作权限的帐号 id，包括代理商和广告主帐号 id (必填)
}

func (p *ElementAppealQuotaGetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证获取元素申诉复审配额请求参数
func (p *ElementAppealQuotaGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	return nil
}

// ElementAppealQuotaGetResp 获取元素申诉复审配额响应
type ElementAppealQuotaGetResp struct {
	AccountID          int64  `json:"account_id,omitempty"`           // 推广帐号 id，有操作权限的帐号 id，包括代理商和广告主帐号 id
	HasPrivilege       int    `json:"has_privilege,omitempty"`        // 是否有权限，取值为 1 表示有权限，0 表示无权限
	DailyQuota         int64  `json:"daily_quota,omitempty"`          // 每日配额
	LeaveQuota         int64  `json:"leave_quota,omitempty"`          // 剩余配额
	QuotaCalculateRule string `json:"quota_calculate_rule,omitempty"` // 申诉复审配额计算规则
}
