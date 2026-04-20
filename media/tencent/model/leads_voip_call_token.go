package model

import "errors"

// ========== 获取网络电话 token ==========
// https://developers.e.qq.com/v3.0/docs/api/leads_voip_call_token/get

// LeadsVoipCallTokenGetReq 获取网络电话 token 请求
type LeadsVoipCallTokenGetReq struct {
	GlobalReq
	AccountID int64  `json:"account_id"`           // 广告主账号 id，直客账号或子客账号 (必填)
	UserID    int64  `json:"user_id"`              // 客服 id，平台下客服 id 不能重复 (必填)
	RequestId string `json:"request_id,omitempty"` // 唯一业务请求 id，不填由线索平台生成后返回
}

func (p *LeadsVoipCallTokenGetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证获取网络电话 token 请求参数
func (p *LeadsVoipCallTokenGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	if p.UserID == 0 {
		return errors.New("user_id为必填")
	}

	return nil
}

// LeadsVoipCallTokenGetResp 获取网络电话 token 响应
type LeadsVoipCallTokenGetResp struct {
	Token     string `json:"token,omitempty"`      // 坐席 token
	RequestId string `json:"request_id,omitempty"` // 唯一业务请求 id
}
