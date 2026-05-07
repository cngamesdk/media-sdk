package model

import "errors"

// SharedWalletBindListReq 客户共享钱包绑定账户查询请求
type SharedWalletBindListReq struct {
	accessTokenReq
	AgentId  int64 `json:"agent_id"`            // 代理商ID，必填
	WalletId int64 `json:"wallet_id"`           // 钱包ID，必填
	PageSize int   `json:"page_size,omitempty"` // 分页大小
	PageNum  int   `json:"page_num,omitempty"`  // 分页页码
}

func (receiver *SharedWalletBindListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *SharedWalletBindListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	if receiver.WalletId <= 0 {
		err = errors.New("wallet_id is empty")
		return
	}
	return
}

// SharedWalletAccountInfo 共享钱包绑定账户信息
type SharedWalletAccountInfo struct {
	Id                int64  `json:"id"`                  // ID
	WalletId          int64  `json:"wallet_id"`           // 钱包ID
	WalletName        string `json:"wallet_name"`         // 钱包名称
	AgentId           int64  `json:"agent_id"`            // 代理商ID
	AgentName         string `json:"agent_name"`          // 代理商名称
	AccountId         int64  `json:"account_id"`          // 关联投放账户ID
	AccountName       string `json:"account_name"`        // 关联投放账户名称
	AccountType       int    `json:"account_type"`        // 关联投放账户类型
	AccountTypeDesc   string `json:"account_type_desc"`   // 账户类型描述
	AccountStatus     int    `json:"account_status"`      // 投放账户状态
	AccountStatusDesc string `json:"account_status_desc"` // 投放账户状态描述
	AccountBindStatus int    `json:"account_bind_status"` // 账户绑定状态
	UserId            int64  `json:"user_id"`             // 关联投放账户快手ID
	UserName          string `json:"user_name"`           // 关联投放账户快手名称
	Subject           string `json:"subject"`             // 公司主体，投放账户公司名称
	ProductId         int64  `json:"product_id"`          // 产品ID
	ProductName       string `json:"product_name"`        // 产品名称
	AppId             int    `json:"app_id"`              // 业务线
	AppIdList         []int  `json:"app_id_list"`         // 业务线列表
	OpenAccountName   string `json:"open_account_name"`   // 开户人
	YesterdayConsume  int64  `json:"yesterday_consume"`   // 昨日消耗
	BindTime          int64  `json:"bind_time"`           // 绑定时间
	CreateTime        int64  `json:"create_time"`         // 创建时间
	UpdateTime        int64  `json:"update_time"`         // 更新时间
}

// SharedWalletBindListResp 客户共享钱包绑定账户查询响应数据（仅data部分）
type SharedWalletBindListResp struct {
	TotalCnt                       int64                     `json:"total_cnt"`                          // 总数
	AccountSharedWalletAccountInfo []SharedWalletAccountInfo `json:"account_shared_wallet_account_info"` // 绑定账户信息列表
}
