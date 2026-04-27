package model

import "errors"

// AgentAdvertiserModifyNameReq 代理商-修改账户名称请求
type AgentAdvertiserModifyNameReq struct {
	accessTokenReq
	AgentId        int64  `json:"agent_id"`         // 代理商id，必填
	AdvertiserId   int64  `json:"advertiser_id"`    // 广告主id，必填
	UserId         int64  `json:"user_id"`          // 广告主对应的快手id，必填
	NewAccountName string `json:"new_account_name"` // 新的广告主名称，必填
}

func (receiver *AgentAdvertiserModifyNameReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentAdvertiserModifyNameReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.UserId <= 0 {
		err = errors.New("user_id is empty")
		return
	}
	if len(receiver.NewAccountName) == 0 {
		err = errors.New("new_account_name is empty")
		return
	}
	return
}

// AgentAdvertiserModifyNameResp 代理商-修改账户名称响应数据（仅data部分）
type AgentAdvertiserModifyNameResp struct {
	Success bool `json:"success"` // 是否成功
}
