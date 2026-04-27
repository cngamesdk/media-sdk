package model

import "errors"

// AgentTransferOutReq 代理商给广告主账户转账请求
type AgentTransferOutReq struct {
	accessTokenReq
	BizUniqueKey    string `json:"biz_unique_key,omitempty"`    // 交易号，可用于查询交易状态
	AdvertiserId    int64  `json:"advertiser_id"`               // 广告主ID，必填
	Amount          int64  `json:"amount,omitempty"`            // 现金金额，转现金的时候，金额放在这个字段上
	CreditAmount    int64  `json:"credit_amount,omitempty"`     // 信用账户金额，转信用、现金+信用的时候，金额放在这个字段
	ExtendedAmount  int64  `json:"extended_amount,omitempty"`   // 预留账户金额
	RebateAmount    int64  `json:"rebate_amount,omitempty"`     // 后返金额
	PreRebateAmount int64  `json:"pre_rebate_amount,omitempty"` // 前返金额
	TransferType    int64  `json:"transfer_type,omitempty"`     // 转账类型 1:现金 2:信用 6:现金+信用
}

func (receiver *AgentTransferOutReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentTransferOutReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	return
}

// AgentTransferOutResp 代理商给广告主账户转账响应数据（仅data部分）
type AgentTransferOutResp struct {
	TradeNo     string `json:"trade_no"`     // 转账请求对应的交易号（业务唯一键）
	Success     bool   `json:"success"`      // 转账是否成功
	TradeStatus string `json:"trade_status"` // 转账交易状态：SUCCESS/DOING/FAILED/EXCEPTION/IGNORED
	FailReason  string `json:"fail_reason"`  // 转账失败原因
}
