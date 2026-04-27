package model

import "errors"

// AgentAccountListReq 代理商-账户列表请求
type AgentAccountListReq struct {
	accessTokenReq
	AgentId         int64  `json:"agent_id"`                    // 代理商id，必填
	Page            int    `json:"page"`                        // 当前页码，必填
	PageSize        int    `json:"page_size"`                   // 页码大小，必填
	SelectType      int    `json:"select_type"`                 // 搜索类型：0=不搜索 1=全部分类 2=广告主ID 3=快手ID 4=广告主昵称 5=企业名称 9=广告主ID批量，必填
	SelectValue     string `json:"select_value"`                // 搜索值，必填
	CreateTimeBegin int64  `json:"create_time_begin,omitempty"` // 广告主创建开始时间(时间戳)
	CreateTimeEnd   int64  `json:"create_time_end,omitempty"`   // 广告主创建结束时间(时间戳)
}

func (receiver *AgentAccountListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentAccountListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	if receiver.Page <= 0 {
		err = errors.New("page is empty")
		return
	}
	if receiver.PageSize <= 0 {
		err = errors.New("page_size is empty")
		return
	}
	return
}

// AuthenticationReviewInfoSnake 账户真实性认证信息
type AuthenticationReviewInfoSnake struct {
	AuthenticationStatus int    `json:"authentication_status"` // 账户真实性认证状态：0=无状态 1=待认证 2=认证中 3=认证通过 4=认证失败 5=认证失效
	AuthenticationDetail string `json:"authentication_detail"` // 认证失败原因
}

// ContractReviewInfoSnake 合同签约信息
type ContractReviewInfoSnake struct {
	ContractReviewStatus int    `json:"contract_review_status"` // 合同签约状态：0=无效 1=签约中 2=已签约 3=已终止 4=已过期 5=已撤销 6=已拒签 7=已删除 8=提前终止
	ContractReviewDetail string `json:"contract_review_detail"` // 合同签约状态详情
}

// UserReviewInfoSnake 账户审核信息
type UserReviewInfoSnake struct {
	ReviewStatus int    `json:"review_status"` // 账户审核状态：0=待提交 1=审核中 2=审核通过 3=审核拒绝
	ReviewDetail string `json:"review_detail"` // 审核拒绝原因
}

// ChildReviewStatusInfoSnake 子审核状态信息
type ChildReviewStatusInfoSnake struct {
	AuthenticationReviewInfo *AuthenticationReviewInfoSnake `json:"authentication_review_info"` // 真实性认证信息
	ContractReviewInfo       *ContractReviewInfoSnake       `json:"contract_review_info"`       // 合同签约信息
	UserReviewInfo           *UserReviewInfoSnake           `json:"user_review_info"`           // 账户审核信息
}

// CertReviewDetail 审核详情
type CertReviewDetail struct {
	Id   int    `json:"id"`   // id
	Desc string `json:"desc"` // 描述
}

// AgentAccountInfo 广告主账户信息
type AgentAccountInfo struct {
	ResponsiblePerson     string                      `json:"responsible_person"`       // 销售责任人
	UserId                int64                       `json:"user_id"`                  // 创建账户时使用的快手user_id
	AccountId             int64                       `json:"account_id"`               // 广告主ID
	UcType                string                      `json:"uc_type"`                  // 账户类型
	PaymentType           string                      `json:"payment_type"`             // 付款类型
	AccountName           string                      `json:"account_name"`             // 快手昵称
	CreateTime            int64                       `json:"create_time"`              // 创建时间
	Balance               int64                       `json:"balance"`                  // 现金余额
	CreditBalance         int64                       `json:"credit_balance"`           // 信用账户余额
	ExtendedBalance       int64                       `json:"extended_balance"`         // 预留账户余额
	Rebate                int64                       `json:"rebate"`                   // 后返余额
	PreRebate             int64                       `json:"pre_rebate"`               // 前返余额
	ContractRebate        int64                       `json:"contract_rebate"`          // 框返余额
	TotalBalance          int64                       `json:"total_balance"`            // 总余额
	LoLimit               int64                       `json:"lo_limit"`                 // 账户最低余额
	SingleOut             int64                       `json:"single_out"`               // 单次转账金额
	AutoOut               bool                        `json:"auto_out"`                 // 自动转账状态
	BalanceWarn           bool                        `json:"balance_warn"`             // 余额不足提醒
	ProductName           string                      `json:"product_name"`             // 产品名称
	FirstCostDay          string                      `json:"first_cost_day"`           // 首日消耗日期
	Industry              string                      `json:"industry"`                 // 一级行业
	SecondIndustry        string                      `json:"second_industry"`          // 二级行业
	KuaibiBalance         int64                       `json:"kuaibi_balance"`           // kuaibi余额
	PushBalance           int64                       `json:"push_balance"`             // push余额
	Recharged             bool                        `json:"recharged"`                // 是否已充值
	CorporationName       string                      `json:"corporation_name"`         // 企业名称
	ReviewStatus          int                         `json:"review_status"`            // 账户审核状态：0=待提交 1=审核中 2=审核通过 3=审核拒绝
	FrozenStatus          int                         `json:"frozen_status"`            // 账户冻结状态：1=未冻结 2=冻结
	TransferAccountStatus bool                        `json:"transfer_account_status"`  // 转账状态
	AppId                 int                         `json:"app_id"`                   // appId
	CreateSource          int                         `json:"create_source"`            // 创建来源
	CopyAccount           bool                        `json:"copy_account"`             // 是否为复制账户
	DirectRebate          int64                       `json:"direct_rebate"`            // 激励余额
	OptimizerOwner        string                      `json:"optimizer_owner"`          // 优化师责任人
	Banned                bool                        `json:"banned"`                   // 是否封禁
	FrozenReason          string                      `json:"frozen_reason"`            // 冻结原因
	ChildReviewStatusInfo *ChildReviewStatusInfoSnake `json:"child_review_status_info"` // 子审核状态信息
	ReviewDetail          []CertReviewDetail          `json:"review_detail"`            // 审核详情
}

// AgentAccountListResp 代理商-账户列表响应数据（仅data部分）
type AgentAccountListResp struct {
	TotalCount int64              `json:"total_count"` // 总数
	PageNo     int64              `json:"page_no"`     // 页码
	PageSize   int64              `json:"page_size"`   // 页大小
	Details    []AgentAccountInfo `json:"details"`     // 广告主列表
}
