package model

import "errors"

// SharedWalletAlertReq 客户共享钱包告警信息查询请求
type SharedWalletAlertReq struct {
	accessTokenReq
	AgentId int64 `json:"agent_id"` // 代理商ID，必填
}

func (receiver *SharedWalletAlertReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *SharedWalletAlertReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	return
}

// SharedWalletAlertBalanceInfo 告警钱包余额信息
type SharedWalletAlertBalanceInfo struct {
	WalletId                  int64   `json:"wallet_id"`                    // 钱包ID
	WalletName                string  `json:"wallet_name"`                  // 钱包名称
	AccountSharedWalletStatus int     `json:"account_shared_wallet_status"` // 钱包状态
	UserId                    int64   `json:"user_id"`                      // 操作人快手ID
	Operator                  string  `json:"operator"`                     // 操作人快手昵称
	CreateTime                int64   `json:"create_time"`                  // 钱包创建时间
	AfterRebate               int64   `json:"after_rebate"`                 // 后返
	BeforeRebate              int64   `json:"before_rebate"`                // 前返
	Credit                    int64   `json:"credit"`                       // 信用
	Cash                      int64   `json:"cash"`                         // 现金
	TotalBalance              int64   `json:"total_balance"`                // 总金额
	DailyTotalCharge          int64   `json:"daily_total_charge"`           // 当日累计总消耗
	AppId                     int     `json:"app_id"`                       // 业务线
	AppIdList                 []int   `json:"app_id_list"`                  // 业务线列表
	BindAccountId             []int64 `json:"bind_account_id"`              // 绑定账户ID列表
	BindAccountCnt            int     `json:"bind_account_cnt"`             // 绑定账户数量
	BindAccountCorpInfo       int     `json:"bind_account_corp_info"`       // 绑定账户公司详情
}

// SharedWalletAlertInfo 共享钱包告警信息
type SharedWalletAlertInfo struct {
	Type                           int                            `json:"type"`                               // 警告类型
	HitNumber                      int                            `json:"hit_number"`                         // 告警命中钱包数量
	Status                         int                            `json:"status"`                             // 警告状态：1=开启 0=关闭
	AccountSharedWalletBalanceInfo []SharedWalletAlertBalanceInfo `json:"account_shared_wallet_balance_info"` // 钱包余额信息列表
}

// SharedWalletAlertResp 客户共享钱包告警信息查询响应数据（仅data部分）
type SharedWalletAlertResp []SharedWalletAlertInfo
