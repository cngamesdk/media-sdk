package model

import "errors"

// SharedWalletConsumeQueryReq 客户共享钱包消耗明细查询请求
type SharedWalletConsumeQueryReq struct {
	accessTokenReq
	AgentId          int64  `json:"agent_id"`           // 代理商ID，必填
	WalletId         string `json:"wallet_id"`          // 钱包ID，必填
	StartConsumeTime int64  `json:"start_consume_time"` // 消耗起始时间，必填
	EndConsumeTime   int64  `json:"end_consume_time"`   // 消耗结束时间，必填
}

func (receiver *SharedWalletConsumeQueryReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *SharedWalletConsumeQueryReq) Validate() (err error) {
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
	if receiver.StartConsumeTime <= 0 {
		err = errors.New("start_consume_time is empty")
		return
	}
	if receiver.EndConsumeTime <= 0 {
		err = errors.New("end_consume_time is empty")
		return
	}
	return
}

// SharedWalletConsumeDetail 共享钱包消耗明细
type SharedWalletConsumeDetail struct {
	BizTradeTime     string `json:"biz_trade_time"`     // 消耗时间
	AccountId        string `json:"account_id"`         // 广告主ID
	AccountName      string `json:"account_name"`       // 广告主名称
	SignCompany      string `json:"sign_company"`       // 签约公司主体
	AgentCompanyName string `json:"agent_company_name"` // 代理商名称
	WalletId         string `json:"wallet_id"`          // 钱包ID
	WalletName       string `json:"wallet_name"`        // 钱包名称
	TotalBalance     string `json:"total_balance"`      // 总消耗
	Cash             string `json:"cash"`               // 现金消耗
	Credit           string `json:"credit"`             // 信用消耗
	BeforeRebate     string `json:"before_rebate"`      // 前返消耗
	AfterRebate      string `json:"after_rebate"`       // 后返消耗
	Remark           string `json:"remark"`             // 备注
}

// SharedWalletConsumeQueryResp 客户共享钱包消耗明细查询响应数据（仅data部分）
type SharedWalletConsumeQueryResp []SharedWalletConsumeDetail
