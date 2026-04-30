package model

import "errors"

// AccountAutoInfoReq 查询账户智投(标准)配置信息请求
type AccountAutoInfoReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主账号ID，必填
}

func (receiver *AccountAutoInfoReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AccountAutoInfoReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	return
}

// OcpxActionTypeConstraint 智投目标成本配置
type OcpxActionTypeConstraint struct {
	OcpxActionType int     `json:"ocpx_action_type"` // 转化目标：191-首日ROI, 190-付费, 180-激活, 394-订单提交, 53-表单优化, 773-关键行为, 324-唤端
	Value          float64 `json:"value"`            // 目标成本约束（非ROI类单位：元），精确至0.001
}

// AccountAutoInfoResp 查询账户智投(标准)配置信息响应数据（仅data部分）
type AccountAutoInfoResp struct {
	AutoManageType           int                        `json:"auto_manage_type"`            // 智投模式：0-标准，1-小说，2-短剧
	AccountAutoManage        int                        `json:"account_auto_manage"`         // 智投开关：0-关闭，1-开启
	OcpxActionTypeConstraint []OcpxActionTypeConstraint `json:"ocpx_action_type_constraint"` // 智投目标成本配置（仅auto_manage_type=0时有值）
	AutoCampaignNameRule     string                     `json:"auto_campaign_name_rule"`     // 广告计划命名规则（仅auto_manage_type=0时有值）
}
