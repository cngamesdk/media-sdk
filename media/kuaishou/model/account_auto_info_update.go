package model

import "errors"

// AccountAutoInfoUpdateReq 更新账户智投(标准)配置信息请求
type AccountAutoInfoUpdateReq struct {
	accessTokenReq
	AdvertiserId             int                        `json:"advertiser_id"`               // 广告主账号ID，必填
	AccountAutoManage        int                        `json:"account_auto_manage"`         // 智投开关：0-关闭，1-开启
	OcpxActionTypeConstraint []OcpxActionTypeConstraint `json:"ocpx_action_type_constraint"` // 智投目标成本配置
	AutoCampaignNameRule     string                     `json:"auto_campaign_name_rule"`     // 广告计划命名规则
}

func (receiver *AccountAutoInfoUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AccountAutoInfoUpdateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.AccountAutoManage != 0 && receiver.AccountAutoManage != 1 {
		err = errors.New("account_auto_manage must be 0 or 1")
		return
	}
	return
}

// AccountAutoInfoUpdateResp 更新账户智投(标准)配置信息响应数据（仅data部分）
type AccountAutoInfoUpdateResp struct {
	AccountAutoManage        int                        `json:"account_auto_manage"`         // 账户智投开关
	OcpxActionTypeConstraint []OcpxActionTypeConstraint `json:"ocpx_action_type_constraint"` // 智投目标成本配置
	AutoCampaignNameRule     string                     `json:"auto_campaign_name_rule"`     // 广告计划命名规则
}
