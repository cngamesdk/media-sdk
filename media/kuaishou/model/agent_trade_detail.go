package model

import "errors"

// AgentTradeDetailReq 获取交易状态信息请求
type AgentTradeDetailReq struct {
	accessTokenReq
	AgentId int64  `json:"agent_id"` // 代理商id，必填
	TradeNo string `json:"trade_no"` // 要查询的交易号，必填
}

func (receiver *AgentTradeDetailReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentTradeDetailReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	if len(receiver.TradeNo) == 0 {
		err = errors.New("trade_no is empty")
		return
	}
	return
}

// AgentTradeDetailResp 获取交易状态信息响应数据（仅data部分）
type AgentTradeDetailResp struct {
	TradeNo              string `json:"trade_no"`                // 交易号
	TradeStatus          string `json:"trade_status"`            // 交易状态
	TradeType            string `json:"trade_type"`              // 交易类型：AGENT_TO_ACCOUNT=广告主充值 ACCOUNT_TO_AGENT=广告主退款 ACCOUNT_TO_ACCOUNT=广告主互转
	PayerSideRecordId    string `json:"payer_side_record_id"`    // 支付侧（转出方一侧）产生流水号
	ReceiverSideRecordId string `json:"receiver_side_record_id"` // 收款侧（转入方一侧）产生流水号
}
