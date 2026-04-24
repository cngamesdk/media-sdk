package model

import "errors"

// AgentSecondaryRechargeReq 一代给二代充值请求
type AgentSecondaryRechargeReq struct {
	accessTokenReq
	AgentId          int64  `json:"agent_id"`                 // 代理商id，必填
	SecondaryAgentId int64  `json:"secondary_agent_id"`       // 子端口id（二代代理商id），必填
	TransferType     int    `json:"transfer_type"`            // 转账类型：1=现金 2=信用，必填
	BizUniqueKey     string `json:"biz_unique_key,omitempty"` // 唯一交易号，可不填
	Amount           int64  `json:"amount,omitempty"`         // 现金，单位厘
	CreditAmount     int64  `json:"credit_amount,omitempty"`  // 信用，单位厘
}

func (receiver *AgentSecondaryRechargeReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentSecondaryRechargeReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	if receiver.SecondaryAgentId <= 0 {
		err = errors.New("secondary_agent_id is empty")
		return
	}
	if receiver.TransferType <= 0 {
		err = errors.New("transfer_type is empty")
		return
	}
	return
}

// AgentSecondaryRechargeResp 一代给二代充值响应数据（仅data部分）
type AgentSecondaryRechargeResp struct {
	Data bool `json:"data"` // 转账结果：true=成功 false=失败
}
