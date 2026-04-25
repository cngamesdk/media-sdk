package model

import "errors"

// AgentAdvertiserCreateReq 代理商-创建广告主请求
type AgentAdvertiserCreateReq struct {
	accessTokenReq
	CorporationName  string `json:"corporation_name"`       // 公司名称，必填
	AdvertiserUserId int64  `json:"advertiser_user_id"`     // 快手id，必填
	AgentId          int64  `json:"agent_id"`               // 代理商ID，必填
	SmsCode          string `json:"sms_code,omitempty"`     // 验证码
	UcType           string `json:"uc_type,omitempty"`      // 账户类型：DSP=DSP-MAPI CID=ESP_CID
	AccountName      string `json:"account_name,omitempty"` // 账户名称
}

func (receiver *AgentAdvertiserCreateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AgentAdvertiserCreateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if len(receiver.CorporationName) == 0 {
		err = errors.New("corporation_name is empty")
		return
	}
	if receiver.AdvertiserUserId <= 0 {
		err = errors.New("advertiser_user_id is empty")
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("agent_id is empty")
		return
	}
	return
}

// AgentAdvertiserCreateResp 代理商-创建广告主响应数据（仅data部分）
type AgentAdvertiserCreateResp struct {
	Id int64 `json:"id"` // 广告主id
}
