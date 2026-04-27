package model

import "errors"

// AgentCopyAccountsReq 代理商-可复制账户列表请求
type AgentCopyAccountsReq struct {
	accessTokenReq
	AgentId  int64  `json:"agent_id"`          // 代理商id，必填
	Keyword  string `json:"keyword,omitempty"` // 模糊查询关键字
	UcType   string `json:"uc_type"`           // 广告主复制类型：DSP_MAPI，必填
	PageNo   int    `json:"page_no"`           // 页码，必填
	PageSize int    `json:"page_size"`         // 页面大小，必填
}

func (receiver *AgentCopyAccountsReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentCopyAccountsReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	if len(receiver.UcType) == 0 {
		err = errors.New("uc_type is empty")
		return
	}
	if receiver.PageNo <= 0 {
		err = errors.New("page_no is empty")
		return
	}
	if receiver.PageSize <= 0 {
		err = errors.New("page_size is empty")
		return
	}
	return
}

// AgentCanCopyAccount 可复制广告主
type AgentCanCopyAccount struct {
	AccountId       int64  `json:"account_id"`       // 广告主id
	UserId          int64  `json:"user_id"`          // 广告主快手id
	AccountName     string `json:"account_name"`     // 广告主名称
	CorporationName string `json:"corporation_name"` // 公司名称
	UcType          string `json:"uc_type"`          // 账户类型
	ProductName     string `json:"product_name"`     // 产品名
}

// AgentCopyAccountsResp 代理商-可复制账户列表响应数据（仅data部分）
type AgentCopyAccountsResp struct {
	AgentCanCopyAccounts []AgentCanCopyAccount `json:"agent_can_copy_accounts"` // 可复制广告主列表
	Total                int64                 `json:"total"`                   // 总条数
}
