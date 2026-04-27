package model

import "errors"

// AgentAccountCanCopyQueryReq 拉取账户可复制信息列表请求
type AgentAccountCanCopyQueryReq struct {
	accessTokenReq
	AgentId       int64   `json:"agent_id"`        // 代理商id，必填
	AccountIdList []int64 `json:"account_id_list"` // 广告主id列表，必填
	UcType        string  `json:"uc_type"`         // 业务线：DSP_MAPI 或 ESP_CID，必填
}

func (receiver *AgentAccountCanCopyQueryReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentAccountCanCopyQueryReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	if len(receiver.AccountIdList) == 0 {
		err = errors.New("account_id_list is empty")
		return
	}
	if len(receiver.UcType) == 0 {
		err = errors.New("uc_type is empty")
		return
	}
	return
}

// AgentAccountCanCopyInfo 账户可复制信息
type AgentAccountCanCopyInfo struct {
	AccountId        int64  `json:"account_id"`         // 账户id
	UserId           int64  `json:"user_id"`            // 快手id
	AccountName      string `json:"account_name"`       // 账户名称
	CorporationName  string `json:"corporation_name"`   // 公司名称
	ProductName      string `json:"product_name"`       // 产品名称
	CanCopy          int    `json:"can_copy"`           // 是否可复制：0=不可复制 1=可复制
	CannotCopyReason string `json:"cannot_copy_reason"` // 不可复制原因
}

// AgentAccountCanCopyQueryResp 拉取账户可复制信息列表响应数据（仅data部分）
type AgentAccountCanCopyQueryResp []AgentAccountCanCopyInfo
