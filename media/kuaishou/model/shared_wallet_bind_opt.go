package model

import "errors"

// SharedWalletBindOptReq 客户共享钱包账户绑定/解绑请求
type SharedWalletBindOptReq struct {
	accessTokenReq
	WalletId        string  `json:"wallet_id"`        // 钱包ID，必填
	AgentId         string  `json:"agent_id"`         // 代理商ID，必填
	TradeNo         string  `json:"trade_no"`         // 交易流水号，必填。格式：mapi_{共享钱包ID}_{绑定accountId}
	AppId           int64   `json:"app_id"`           // 业务线标识，必填，磁力智投默认填7
	AccountId       []int64 `json:"account_id"`       // 绑定/解绑账户ID集合，必填
	AccountOperator int64   `json:"account_operator"` // 账户操作，必填：1=绑定 2=解绑
	UserId          string  `json:"user_id"`          // 操作人快手ID，必填
}

func (receiver *SharedWalletBindOptReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *SharedWalletBindOptReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.WalletId == "" {
		err = errors.New("wallet_id is empty")
		return
	}
	if receiver.AgentId == "" {
		err = errors.New("agent_id is empty")
		return
	}
	if receiver.TradeNo == "" {
		err = errors.New("trade_no is empty")
		return
	}
	if len(receiver.AccountId) == 0 {
		err = errors.New("account_id is empty")
		return
	}
	if receiver.AccountOperator <= 0 {
		err = errors.New("account_operator is empty")
		return
	}
	if receiver.UserId == "" {
		err = errors.New("user_id is empty")
		return
	}
	return
}

// SharedWalletBindOptResp 客户共享钱包账户绑定/解绑响应数据（仅data部分）
type SharedWalletBindOptResp struct {
	AccountId   string `json:"account_id"`   // 关联投放账户ID
	FailMessage string `json:"fail_message"` // 操作失败描述
}
