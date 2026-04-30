package model

import "errors"

// KolUserInfoReq 授权快手号信息（请求）
type KolUserInfoReq struct {
	KolUserType int   `json:"kol_user_type"` // 授权快手号类型
	KolUserId   int64 `json:"kol_user_id"`   // 授权快手号ID
}

// AutoOcpxConstraintReq 智投目标成本配置（请求）
type AutoOcpxConstraintReq struct {
	OcpxActionType int     `json:"ocpx_action_type"` // 转化目标
	Value          float64 `json:"value"`            // 目标成本
}

// AutoProjectConfigUpdateReq 智投配置更新请求
type AutoProjectConfigUpdateReq struct {
	accessTokenReq
	AdvertiserId             int                     `json:"advertiser_id"`               // 广告主账号ID，必填
	AccountAutoManage        int                     `json:"account_auto_manage"`         // 智投开关：0-关闭，1-开启
	AutoManageType           int                     `json:"auto_manage_type"`            // 智投模式：0-标准，1-小说，2-短剧
	KolUserInfo              *KolUserInfoReq         `json:"kol_user_info"`               // 授权快手号信息（仅auto_manage_type=1时生效）
	OcpxActionTypeConstraint []AutoOcpxConstraintReq `json:"ocpx_action_type_constraint"` // 智投目标成本配置（仅auto_manage_type=0/1时生效）
	AutoCampaignNameRule     string                  `json:"auto_campaign_name_rule"`     // 广告计划命名规则（仅auto_manage_type=0时生效）
}

func (receiver *AutoProjectConfigUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AutoProjectConfigUpdateReq) Validate() (err error) {
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

// KolUserInfoResp 授权快手号信息（响应）
type KolUserInfoResp struct {
	UserId   int64  `json:"user_id"`   // 授权快手号ID
	UserName string `json:"user_name"` // 授权快手号名称
}

// AutoOcpxConstraintResp 智投目标成本配置（响应）
type AutoOcpxConstraintResp struct {
	OcpxActionType int     `json:"ocpx_action_type"` // 转化目标
	RoiRatio       float64 `json:"roi_ratio"`        // 目标成本
}

// AutoProjectConfigUpdateResp 智投配置更新响应数据（仅data部分）
type AutoProjectConfigUpdateResp struct {
	AccountAutoManage        int                      `json:"account_auto_manage"`         // 智投开关
	AutoManageType           int                      `json:"auto_manage_type"`            // 智投模式
	KolUserInfo              KolUserInfoResp          `json:"kol_user_info"`               // 授权快手号信息（auto_manage_type=1时生效）
	KolUserType              int                      `json:"kol_user_type"`               // 授权快手号类型：1-普通快手号，2-蓝V服务号，3-聚星达人
	OcpxActionTypeConstraint []AutoOcpxConstraintResp `json:"ocpx_action_type_constraint"` // 智投目标成本配置（auto_manage_type=0/1时生效）
	AutoCampaignNameRule     string                   `json:"auto_campaign_name_rule"`     // 广告计划命名规则（auto_manage_type=0时生效）
}
