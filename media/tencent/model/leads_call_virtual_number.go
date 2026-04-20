package model

import "errors"

// ========== 获取中间号 ==========
// https://developers.e.qq.com/v3.0/docs/api/leads_call_virtual_number/get

// LeadsCallVirtualNumberGetReq 获取中间号请求
type LeadsCallVirtualNumberGetReq struct {
	GlobalReq
	AccountID    int64  `json:"account_id"`               // 广告主账号id (必填)
	LeadsId      int64  `json:"leads_id,omitempty"`       // 线索id，与outer_leads_id二选一必填
	OuterLeadsId string `json:"outer_leads_id,omitempty"` // 外部线索id，与leads_id二选一必填
	Caller       string `json:"caller"`                   // 主叫号码 (必填)，11位手机号或座机号
	Callee       string `json:"callee"`                   // 被叫号码 (必填)，11位手机号或座机号
	RequestId    string `json:"request_id,omitempty"`     // 唯一业务请求id
}

func (p *LeadsCallVirtualNumberGetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证获取中间号请求参数
func (p *LeadsCallVirtualNumberGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	if p.LeadsId == 0 && p.OuterLeadsId == "" {
		return errors.New("leads_id和outer_leads_id二选一必填")
	}

	if p.Caller == "" {
		return errors.New("caller为必填")
	}

	if p.Callee == "" {
		return errors.New("callee为必填")
	}

	return nil
}

// LeadsCallVirtualNumberGetResp 获取中间号响应
type LeadsCallVirtualNumberGetResp struct {
	VirtualNumber string `json:"virtual_number,omitempty"` // 虚拟中间号
	RequestId     string `json:"request_id,omitempty"`     // 唯一业务请求id
	ContactId     string `json:"contact_id,omitempty"`     // 标识一次外呼行为
}
