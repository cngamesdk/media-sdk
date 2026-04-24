package model

import "errors"

// AgentDepositPageInfo 分页请求体
type AgentDepositPageInfo struct {
	CurrentPage int   `json:"current_page,omitempty"` // 当前页号
	PageSize    int   `json:"page_size,omitempty"`    // 页内记录数
	TotalCount  int64 `json:"total_count,omitempty"`  // 总条数
}

// AgentDepositListReq 代理商流水列表请求
type AgentDepositListReq struct {
	accessTokenReq
	AgentId       int64                 `json:"agent_id"`                 // 代理商id，必填
	StartTime     int64                 `json:"start_time"`               // 开始时间(时间戳)，必填
	EndTime       int64                 `json:"end_time"`                 // 结束时间(时间戳)，必填
	InvoiceStatus int                   `json:"invoice_status,omitempty"` // 开票状态：0-全部 1-未开票 2-申请中 3-已开票 4-已邮寄 5-已拒绝 6-不可开票 7-系统重新申请
	OperationType int                   `json:"operation_type,omitempty"` // 操作类型：0-全部 1-充值 2-退款 11-信用账户充值 12-信用账户还款 13-信用还款和现金充值 14-保证金充值 15-保证金减款
	IsPage        bool                  `json:"is_page,omitempty"`        // 是否分页
	PageInfo      *AgentDepositPageInfo `json:"page_info,omitempty"`      // 分页请求体
}

func (receiver *AgentDepositListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentDepositListReq) Validate() (err error) {
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

// AgentDepositRecordItem 代理商流水明细
type AgentDepositRecordItem map[string]interface{}

// AgentDepositListResp 代理商流水列表响应数据（仅data部分）
type AgentDepositListResp struct {
	Details  []AgentDepositRecordItem `json:"details"`   // 流水明细列表
	PageSize int                      `json:"page_size"` // 页大小
	PageNo   int                      `json:"page_no"`   // 当前页号
	Total    int64                    `json:"total"`     // 总条数
}
