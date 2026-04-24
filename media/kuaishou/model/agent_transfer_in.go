package model

import "errors"

// AgentTransferInReq 广告主退钱给代理商请求
type AgentTransferInReq struct {
	accessTokenReq
	AdvertiserId         int64  `json:"advertiser_id"`            // 广告主ID，必填
	CustomTransferAmount int64  `json:"custom_transfer_amount"`   // 退款金额（厘），必填
	TransferType         int    `json:"transfer_type"`            // 退款类型：1=现金 2=信用 6=现金+信用，必填
	BizUniqueKey         string `json:"biz_unique_key,omitempty"` // 交易号，可用于查询交易状态
}

func (receiver *AgentTransferInReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentTransferInReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.CustomTransferAmount <= 0 {
		err = errors.New("custom_transfer_amount must be greater than 0")
		return
	}
	if receiver.TransferType <= 0 {
		err = errors.New("transfer_type is empty")
		return
	}
	return
}

// AgentTransferInResp 广告主退钱给代理商响应数据（仅data部分）
type AgentTransferInResp struct {
	TradeNo     string `json:"trade_no"`     // 转账请求对应的交易号（业务唯一键）
	Success     bool   `json:"success"`      // 转账是否成功
	TradeStatus string `json:"trade_status"` // 交易状态：SUCCESS=成功 DOING=处理中 FAILED=失败 EXCEPTION=异常 IGNORED=无效
	FailReason  string `json:"fail_reason"`  // 转账失败原因
}
