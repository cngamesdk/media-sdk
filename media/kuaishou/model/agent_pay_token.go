package model

import "errors"

// AgentPayTokenReq 代理商-获取新的交易号请求
type AgentPayTokenReq struct {
	accessTokenReq
	AgentId int64  `json:"agent_id"` // 代理商id，必填
	Remark  string `json:"remark"`   // 获取交易号的场景说明，必填
}

func (receiver *AgentPayTokenReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentPayTokenReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	if len(receiver.Remark) == 0 {
		err = errors.New("remark is empty")
		return
	}
	return
}

// AgentPayTokenResp 代理商-获取新的交易号响应数据（仅data部分）
type AgentPayTokenResp struct {
	TradeNo string `json:"trade_no"` // 新分配的交易号
}
