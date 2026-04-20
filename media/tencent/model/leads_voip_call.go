package model

import "errors"

// ========== 网络电话呼叫 ==========
// https://developers.e.qq.com/v3.0/docs/api/leads_voip_call/add

// LeadsVoipCallAddReq 网络电话呼叫请求
type LeadsVoipCallAddReq struct {
	GlobalReq
	AccountID    int64  `json:"account_id"`               // 广告主账号 id，直客账号或子客账号 (必填)
	LeadsId      int64  `json:"leads_id,omitempty"`       // 线索 id，与 outer_leads_id 二选一必填
	OuterLeadsId string `json:"outer_leads_id,omitempty"` // 外部线索 id，与 leads_id 二选一必填
	UserID       int64  `json:"user_id"`                  // 客服 id，平台下客服 id 不能重复 (必填)
	CalleeNumber string `json:"callee_number"`            // 线索号码 (必填)
	RequestId    string `json:"request_id,omitempty"`     // 唯一业务请求 id，不填由线索平台生成后返回
	Version      string `json:"version,omitempty"`        // 版本号，不填就是最新版
}

func (p *LeadsVoipCallAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证网络电话呼叫请求参数
func (p *LeadsVoipCallAddReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	if p.LeadsId == 0 && p.OuterLeadsId == "" {
		return errors.New("leads_id和outer_leads_id二选一必填")
	}

	if p.UserID == 0 {
		return errors.New("user_id为必填")
	}

	if p.CalleeNumber == "" {
		return errors.New("callee_number为必填")
	}

	return nil
}

// LeadsVoipCallAddResp 网络电话呼叫响应
type LeadsVoipCallAddResp struct {
	ContactId string `json:"contact_id,omitempty"` // 用来标识一次外呼行为
	RequestId string `json:"request_id,omitempty"` // 唯一业务请求 id
}
