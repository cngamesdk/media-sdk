package model

import "errors"

// CampaignUpdateReq 修改广告计划请求
type CampaignUpdateReq struct {
	accessTokenReq
	AdvertiserId               int64                            `json:"advertiser_id"`                           // 账号ID，必填，在获取 access_token 时返回
	CampaignId                 int64                            `json:"campaign_id"`                             // 广告计划ID，必填
	CampaignName               string                           `json:"campaign_name,omitempty"`                 // 广告计划名称，1-100个字符，账户内唯一
	DayBudget                  *int64                           `json:"day_budget,omitempty"`                    // 日预算，单位：厘；0=不限；最低500元，最高1亿元；与 day_budget_schedule 不可同时使用；传0可清除预算；null=不更新
	DayBudgetSchedule          []int64                          `json:"day_budget_schedule,omitempty"`           // 分日预算（7日数组），单位：厘；0=不限；与 day_budget 不可同时使用；传空数组可清除；null=不更新
	RangeBudget                int64                            `json:"range_budget,omitempty"`                  // 周期稳定投放总预算，仅 periodic_delivery_type=1 时有效，单位：元，只增不减
	CapBid                     int64                            `json:"cap_bid,omitempty"`                       // 成本约束（非ROI单/双约束），单位：厘；范围：0-8000元
	CapRoiRatio                float64                          `json:"cap_roi_ratio,omitempty"`                 // 成本约束ROI约束；范围：0-100
	ConstraintCpa              int64                            `json:"constraint_cpa,omitempty"`                // 浅层成本约束（非ROI双约束），单位：厘；范围：0-8000元；需同时填写 cap_bid
	PutStatus                  int                              `json:"put_status,omitempty"`                    // 计划状态：1=投放 2=暂停（删除请使用修改广告计划状态接口）
	AutoAdjust                 int                              `json:"auto_adjust,omitempty"`                   // 自动调节开关：0=关闭 1=开启
	AutoBuild                  int                              `json:"auto_build,omitempty"`                    // 自动建组开关：0=关闭 1=开启
	AutoBuildNameRule          *CampaignCreateAutoBuildNameRule `json:"auto_build_name_rule,omitempty"`          // 自动建组命名规则，auto_build=1 时有效；广告组和创意命名规则必须包含[日期]和[序号]宏
	UnitNameRule               string                           `json:"unit_name_rule,omitempty"`                // 广告组命名规则，必须包含[日期]和[序号]宏；示例：系统自动搭建_[日期][序号]
	CreativeNameRule           string                           `json:"creative_name_rule,omitempty"`            // 创意命名规则，必须包含[日期]和[序号]宏；示例：系统自动搭建_[日期][序号]
	AutoManage                 int                              `json:"auto_manage,omitempty"`                   // 智能优化：0=关闭 1=开启；开启时 auto_adjust 和 auto_build 也必须开启；不支持周期稳定投放计划
	CampaignOcpxActionType     int                              `json:"campaign_ocpx_action_type,omitempty"`     // 智能优化浅层目标
	CampaignDeepConversionType int                              `json:"campaign_deep_conversion_type,omitempty"` // 智能优化深层目标
}

func (receiver *CampaignUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *CampaignUpdateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.CampaignId <= 0 {
		err = errors.New("campaign_id is empty")
		return
	}
	return
}

// CampaignUpdateResp 修改广告计划响应数据（仅data部分）
type CampaignUpdateResp struct {
	CampaignId int64 `json:"campaign_id"` // 广告计划ID
}
