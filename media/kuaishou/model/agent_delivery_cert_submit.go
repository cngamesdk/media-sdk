package model

import "errors"

// AgentDeliveryCertSubmitReq 创建或更新投放资质请求
type AgentDeliveryCertSubmitReq struct {
	accessTokenReq
	AdvertiserId int64       `json:"advertiser_id"` // 广告主id，必填
	AgentId      int64       `json:"agent_id"`      // 代理商id，必填，和access_token保持一致
	CertList     []CertParam `json:"cert_list"`     // 投放资质列表，只传投放资质（cert_category为5），不填代表删除所有投放资质
}

func (receiver *AgentDeliveryCertSubmitReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentDeliveryCertSubmitReq) Validate() (err error) {
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
	return
}

// AgentDeliveryCertSubmitResp 创建或更新投放资质响应数据（仅data部分）
type AgentDeliveryCertSubmitResp struct {
	Result bool `json:"result"` // 是否成功
}
