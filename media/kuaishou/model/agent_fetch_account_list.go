package model

import "errors"

// AgentFetchAccountListReq 批量拉取代理商下账户列表请求
type AgentFetchAccountListReq struct {
	accessTokenReq
	AgentId         int64  `json:"agent_id"`                   // 代理商id，必填
	BeginAccountId  int64  `json:"begin_account_id"`           // 分页起始账户ID，首次传0，后续传上次最后一条的account_id，必填
	BatchSize       int    `json:"batch_size"`                 // 每批拉取数量，最大1000，必填
	CreateTimeBegin int64  `json:"create_time_begin"`          // 账户创建开始时间(时间戳毫秒)，必填
	CreateTimeEnd   int64  `json:"create_time_end"`            // 账户创建结束时间(时间戳毫秒)，必填
	CorporationName string `json:"corporation_name,omitempty"` // 企业名称，精确匹配
}

func (receiver *AgentFetchAccountListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentFetchAccountListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	if receiver.BatchSize <= 0 {
		err = errors.New("batch_size is empty")
		return
	}
	if receiver.CreateTimeBegin <= 0 {
		err = errors.New("create_time_begin is empty")
		return
	}
	if receiver.CreateTimeEnd <= 0 {
		err = errors.New("create_time_end is empty")
		return
	}
	return
}

// AgentFetchAccountListResp 批量拉取代理商下账户列表响应数据（仅data部分）
type AgentFetchAccountListResp struct {
	Details    []AgentAccountInfo `json:"details"`     // 广告主账户列表
	TotalCount int64              `json:"total_count"` // 总数
}
