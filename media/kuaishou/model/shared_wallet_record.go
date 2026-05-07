package model

import "errors"

// SharedWalletRecordReq 客户共享钱包交易明细查询请求
type SharedWalletRecordReq struct {
	accessTokenReq
	AgentId        int64  `json:"agent_id"`             // 代理商ID，必填
	WalletId       string `json:"wallet_id"`            // 钱包ID，必填
	StartTradeTime int64  `json:"start_trade_time"`     // 交易开始时间戳，必填
	EndTradeTime   int64  `json:"end_trade_time"`       // 交易结束时间戳，必填
	TradeType      int64  `json:"trade_type,omitempty"` // 交易类型：16=转入 17=转出，可选
}

func (receiver *SharedWalletRecordReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *SharedWalletRecordReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	if receiver.WalletId == "" {
		err = errors.New("wallet_id is empty")
		return
	}
	if receiver.StartTradeTime <= 0 {
		err = errors.New("start_trade_time is empty")
		return
	}
	if receiver.EndTradeTime <= 0 {
		err = errors.New("end_trade_time is empty")
		return
	}
	return
}

// SharedWalletRecordDetail 共享钱包交易明细
type SharedWalletRecordDetail struct {
	TradeType        int64  `json:"trade_type"`         // 交易类型：16=转入共享钱包 17=从共享钱包转出
	TradeNo          string `json:"trade_no"`           // 交易流水号
	Operator         string `json:"operator"`           // 操作人
	UserId           string `json:"user_id"`            // 快手userId
	BizTradeTime     string `json:"biz_trade_time"`     // 交易时间戳
	SignCompany      string `json:"sign_company"`       // 签约公司主体
	AgentCompanyName string `json:"agent_company_name"` // 代理商公司名称
	AppId            int64  `json:"app_id"`             // 快手业务线ID
	WalletId         string `json:"wallet_id"`          // 钱包ID
	WalletName       string `json:"wallet_name"`        // 钱包名称
	TotalBalance     string `json:"total_balance"`      // 交易总金额
	Cash             string `json:"cash"`               // 交易现金金额
	Credit           string `json:"credit"`             // 交易信用金额
	BeforeRebate     string `json:"before_rebate"`      // 交易前返金额
	AfterRebate      string `json:"after_rebate"`       // 交易后返金额
	Remark           string `json:"remark"`             // 交易备注信息
}

// SharedWalletRecordResp 客户共享钱包交易明细查询响应数据（仅data部分）
type SharedWalletRecordResp []SharedWalletRecordDetail
