package model

import "errors"

// AgentAccountTransferReq 广告主转账请求
type AgentAccountTransferReq struct {
	accessTokenReq
	FromAccountId  int64  `json:"from_account_id"` // 转账广告主ID，必填
	ToAccountId    int64  `json:"to_account_id"`   // 目标广告主ID，必填
	OperatorName   string `json:"operator_name"`   // 操作人名称，必填
	TransferType   int    `json:"transfer_type"`   // 操作类型：1=现金 2=信用 3=框返 4=激励 6=现金+信用，必填
	TransferAmount int64  `json:"transfer_amount"` // 转账金额，单位厘，必填
	AgentId        int64  `json:"agent_id"`        // 代理商ID，必填
	BizUniqueKey   string `json:"biz_unique_key"`  // 转账交易唯一串（可通过getPayToken接口获取），必填
}

func (receiver *AgentAccountTransferReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentAccountTransferReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.FromAccountId <= 0 {
		err = errors.New("from_account_id is empty")
		return
	}
	if receiver.ToAccountId <= 0 {
		err = errors.New("to_account_id is empty")
		return
	}
	if len(receiver.OperatorName) == 0 {
		err = errors.New("operator_name is empty")
		return
	}
	if receiver.TransferType <= 0 {
		err = errors.New("transfer_type is empty")
		return
	}
	if receiver.TransferAmount <= 0 {
		err = errors.New("transfer_amount must be greater than 0")
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	if len(receiver.BizUniqueKey) == 0 {
		err = errors.New("biz_unique_key is empty")
		return
	}
	return
}

// AgentAccountTransferResp 广告主转账响应数据（仅data部分）
type AgentAccountTransferResp struct {
	TradeNo     string `json:"trade_no"`     // 转账请求对应的交易号（业务唯一键）
	Success     bool   `json:"success"`      // 转账是否成功
	TradeStatus string `json:"trade_status"` // 交易状态：SUCCESS=成功 DOING=处理中 FAILED=失败 EXCEPTION=异常 IGNORED=无效
	FailReason  string `json:"fail_reason"`  // 转账失败原因
}
