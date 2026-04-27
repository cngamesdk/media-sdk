package model

import "errors"

// AgentFinanceCanDecreaseQueryReq 广告主可转/退金额查询请求
type AgentFinanceCanDecreaseQueryReq struct {
	accessTokenReq
	AgentId      int64 `json:"agent_id"`      // 代理商id，必填
	AdvertiserId int64 `json:"advertiser_id"` // 广告主id，必填
	RefundType   int   `json:"refund_type"`   // 转退类型：1=现金转账 2=信用转账 3=框返转账 4=激励转账 6=信用+现金转账，必填
}

func (receiver *AgentFinanceCanDecreaseQueryReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentFinanceCanDecreaseQueryReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.RefundType <= 0 {
		err = errors.New("refund_type is empty")
		return
	}
	return
}

// AgentFinanceCanDecreaseQueryResp 广告主可转/退金额查询响应数据（仅data部分）
type AgentFinanceCanDecreaseQueryResp struct {
	Amount int64 `json:"amount"` // 可转/退金额，单位：厘
}
