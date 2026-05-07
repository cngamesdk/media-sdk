package model

import "errors"

// SharedWalletBalanceReq 客户共享钱包余额查询请求
type SharedWalletBalanceReq struct {
	accessTokenReq
	AdvertiserId int64   `json:"advertiser_id"`        // 广告主ID，必填（主要用于鉴权，无实际业务含义）
	AgentId      int64   `json:"agent_id"`             // 代理商ID，必填
	WalletIds    []int64 `json:"wallet_ids"`           // 钱包ID列表，必填
	QueryTime    int64   `json:"query_time,omitempty"` // 日终余额查询起始时间戳，可选
}

func (receiver *SharedWalletBalanceReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *SharedWalletBalanceReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	if len(receiver.WalletIds) == 0 {
		err = errors.New("wallet_ids is empty")
		return
	}
	return
}

// SharedWalletBalanceInfo 共享钱包余额信息
type SharedWalletBalanceInfo struct {
	WalletId                  int64  `json:"wallet_id"`                    // 钱包ID
	WalletName                string `json:"wallet_name"`                  // 钱包名称
	AccountSharedWalletStatus int    `json:"account_shared_wallet_status"` // 钱包状态
	UserId                    int64  `json:"user_id"`                      // 操作人快手ID
	Operator                  string `json:"operator"`                     // 操作人快手昵称
	CreateUserId              int64  `json:"create_user_id"`               // 创建人快手ID
	CreateOperator            string `json:"create_operator"`              // 创建人快手昵称
	CreateTime                int64  `json:"create_time"`                  // 钱包创建时间
	AfterRebate               int64  `json:"after_rebate"`                 // 后返
	BeforeRebate              int64  `json:"before_rebate"`                // 前返
	Credit                    int64  `json:"credit"`                       // 信用
	Cash                      int64  `json:"cash"`                         // 现金
	TotalBalance              int64  `json:"total_balance"`                // 总金额
	AfterRebateEnd            int64  `json:"after_rebate_end"`             // 后返日终余额
	BeforeRebateEnd           int64  `json:"before_rebate_end"`            // 前返日终余额
	CreditEnd                 int64  `json:"credit_end"`                   // 信用日终余额
	CashEnd                   int64  `json:"cash_end"`                     // 现金日终余额
	CreditRefundBalance       int64  `json:"credit_refund_balance"`        // 信用可退金额
	CashRefundBalance         int64  `json:"cash_refund_balance"`          // 现金可退金额
}

// SharedWalletBalanceResp 客户共享钱包余额查询响应数据（仅data部分）
type SharedWalletBalanceResp []SharedWalletBalanceInfo
