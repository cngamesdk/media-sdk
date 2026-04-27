package model

import "errors"

// AgentAdvertiserDepositSearchMap 搜索条件
type AgentAdvertiserDepositSearchMap struct {
	BatchAccountIdExact string `json:"batch_account_id_exact,omitempty"` // 账户id搜索
	BatchUserIdExact    string `json:"batch_user_id_exact,omitempty"`    // 快手userId搜索
	BatchCorpName       string `json:"batch_corp_name,omitempty"`        // 公司名搜索
	BatchProductName    string `json:"batch_product_name,omitempty"`     // 产品名搜索
}

// AgentAdvertiserDepositPageInfo 分页信息
type AgentAdvertiserDepositPageInfo struct {
	CurrentPage int   `json:"current_page,omitempty"` // 当前页号
	PageSize    int   `json:"page_size,omitempty"`    // 页内记录数
	TotalCount  int64 `json:"total_count,omitempty"`  // 总条数
}

// AgentAdvertiserDepositReq 代理商-广告主流水列表请求
type AgentAdvertiserDepositReq struct {
	accessTokenReq
	SearchTypeValueMap *AgentAdvertiserDepositSearchMap `json:"search_type_value_map,omitempty"` // 搜索条件
	StartTime          int64                            `json:"start_time"`                      // 开始时间(时间戳)，必填
	EndTime            int64                            `json:"end_time"`                        // 结束时间(时间戳)，必填
	OperationType      int                              `json:"operation_type"`                  // 操作类型：0=全部 1=充值 2=退款 11=信用账户充值 14=保证金充值 15=保证金减款，必填
	AccountSearchType  int                              `json:"account_search_type"`             // 搜索类型：0=全部分类 1=广告主ID 2=快手ID 3=产品名称 4=企业名称，必填
	Keyword            string                           `json:"keyword,omitempty"`               // 检索字段
	UcType             string                           `json:"uc_type"`                         // 业务线筛选：空字符串=全部 DSP=信息流 ADX=ADX ESP=电商，必填
	AgentId            int64                            `json:"agent_id"`                        // 代理商id，必填
	PageInfo           *AgentAdvertiserDepositPageInfo  `json:"page_info,omitempty"`             // 分页
}

func (receiver *AgentAdvertiserDepositReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentAdvertiserDepositReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	if receiver.StartTime <= 0 {
		err = errors.New("start_time is empty")
		return
	}
	if receiver.EndTime <= 0 {
		err = errors.New("end_time is empty")
		return
	}
	return
}

// AgentAdvertiserDepositRecord 广告主流水记录
type AgentAdvertiserDepositRecord struct {
	OperatorName         string `json:"operator_name"`          // 操作人名称
	OperatorId           int64  `json:"operator_id"`            // 操作人id
	OperationType        string `json:"operation_type"`         // 操作类型
	AdvertiserId         int64  `json:"advertiser_id"`          // 广告主id
	UserId               int64  `json:"user_id"`                // 快手userId
	ProductName          string `json:"product_name"`           // 产品名
	CorporationName      string `json:"corporation_name"`       // 公司名称
	FirstCostDay         string `json:"first_cost_day"`         // 首次消费日期
	Industry             string `json:"industry"`               // 行业
	SecondIndustry       string `json:"second_industry"`        // 二级行业
	Amount               int64  `json:"amount"`                 // 现金金额
	CreditAmount         int64  `json:"credit_amount"`          // 信用金额
	RebateAmount         int64  `json:"rebate_amount"`          // 后返金额
	PreRebateAmount      int64  `json:"pre_rebate_amount"`      // 前返金额
	ContractRebateAmount int64  `json:"contract_rebate_amount"` // 合同返点金额
	DirectRebateAmount   int64  `json:"direct_rebate_amount"`   // 直客返点金额
	BeforeBalance        int64  `json:"before_balance"`         // 操作前余额
	AfterBalance         int64  `json:"after_balance"`          // 操作后余额
	BankSn               string `json:"bank_sn"`                // 银行流水号
	Remark               string `json:"remark"`                 // 备注
	Date                 string `json:"date"`                   // 日期
	CreateTime           int64  `json:"create_time"`            // 创建时间
	SerialId             int64  `json:"serial_id"`              // 流水id
	Status               int    `json:"status"`                 // 状态
	InstructionId        string `json:"instruction_id"`         // 指令id
	TradeNo              string `json:"trade_no"`               // 交易号
}

// AgentAdvertiserDepositResp 代理商-广告主流水列表响应数据（仅data部分）
type AgentAdvertiserDepositResp struct {
	Details  []AgentAdvertiserDepositRecord  `json:"details"`   // 流水记录列表
	PageInfo *AgentAdvertiserDepositPageInfo `json:"page_info"` // 分页信息
}
