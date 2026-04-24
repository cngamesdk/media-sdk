package model

import "errors"

// AgentSmsCodeSendReq 代理商-创建广告主发送验证码请求
type AgentSmsCodeSendReq struct {
	accessTokenReq
	AdvertiserUserId int64 `json:"advertiser_user_id"` // 快手id
	AgentId          int64 `json:"agent_id"`           // 代理商ID
}

func (receiver *AgentSmsCodeSendReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentSmsCodeSendReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserUserId <= 0 {
		err = errors.New("advertiser_user_id is empty")
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	return
}

// AgentSmsCodeSendResp 代理商-创建广告主发送验证码响应数据（仅data部分）
type AgentSmsCodeSendResp struct {
	CheckSmsCode bool `json:"check_sms_code"` // 是否需要校验验证码
	Result       bool `json:"result"`         // 是否成功
}
