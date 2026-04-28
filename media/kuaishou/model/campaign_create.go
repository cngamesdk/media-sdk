package model

import "errors"

// CampaignCreateAutoBuildNameRule 自动建组命名规则
type CampaignCreateAutoBuildNameRule struct {
	UnitNameRule     string `json:"unit_name_rule"`     // 广告组命名规则，必须包含[日期]和[序号]宏
	CreativeNameRule string `json:"creative_name_rule"` // 创意命名规则，必须包含[日期]和[序号]宏
}

// CampaignCreateReq 创建广告计划请求
type CampaignCreateReq struct {
	accessTokenReq
	AdvertiserId            int64                            `json:"advertiser_id"`                        // 账号ID，必填，在获取 access_token 时返回
	CampaignName            string                           `json:"campaign_name"`                        // 广告计划名称，必填，1-100个字符，账户内唯一
	Type                    int                              `json:"type"`                                 // 营销目标类型，必填：2=应用下载 4=品牌活动 5=线索收集 7=应用促活 9=商品目录/DPA 16=粉丝/直播 19=小程序/游戏 30=短剧 32=微信小程序 34=小说 35=电商
	SmartMaterialSupply     int                              `json:"smart_material_supply,omitempty"`      // 智能素材补充开关
	PutStatus               int                              `json:"put_status,omitempty"`                 // 投放状态：1=投放（默认）2=暂停
	PeriodicDeliveryType    int                              `json:"periodic_delivery_type,omitempty"`     // 周期稳定投放：0=关闭（默认）1=开启，需白名单
	RangeBudget             int64                            `json:"range_budget,omitempty"`               // 周期总预算，单位：厘；最低3500元，最高100000元
	ContinuePeriodType      int                              `json:"continue_period_type,omitempty"`       // 周期延续：1=关闭 2=开启
	PeriodicDays            int                              `json:"periodic_days,omitempty"`              // 投放周期，默认7天，白名单用户可选5天
	AdType                  int                              `json:"ad_type,omitempty"`                    // 计划类型：0=信息流（默认）1=搜索
	AutoAdjust              int                              `json:"auto_adjust,omitempty"`                // 自动调节：0=关闭 1=开启
	BidType                 int                              `json:"bid_type,omitempty"`                   // 出价类型：0=默认 1=最大转化量，创建后不可修改
	DayBudget               int64                            `json:"day_budget,omitempty"`                 // 日预算，单位：厘；最低500元，最高1亿元，0=不限；与 day_budget_schedule 不可同时使用
	DayBudgetSchedule       []int64                          `json:"day_budget_schedule,omitempty"`        // 分日预算（7日数组），单位：厘；与 day_budget 不可同时使用
	AutoBuild               int                              `json:"auto_build,omitempty"`                 // 自动建组：0=关闭 1=开启
	AutoBuildNameRule       *CampaignCreateAutoBuildNameRule `json:"auto_build_name_rule,omitempty"`       // 自动建组命名规则，auto_build=1 时必填
	CapRoiRatio             float64                          `json:"cap_roi_ratio,omitempty"`              // ROI约束，默认0=不限
	CapBid                  int64                            `json:"cap_bid,omitempty"`                    // 成本约束，默认0=不限
	ConstraintCpa           int64                            `json:"constraint_cpa,omitempty"`             // 浅层约束目标成本，默认0=不限
	AutoManage              int                              `json:"auto_manage,omitempty"`                // 智能投放：0=关闭 1=开启，需 auto_adjust=1 且 auto_build=1
	AutoPhotoScope          int                              `json:"auto_photo_scope,omitempty"`           // 建组取材范围：0=系统最优 1=素材包
	PhotoPackageInfo        []int64                          `json:"photo_package_info,omitempty"`         // 素材包ID列表，auto_photo_scope=1 时必填
	PeriodicDeliveryPutType int                              `json:"periodic_delivery_put_type,omitempty"` // 夜间投放标识：2=夜间稳定，不传=常规
	UnitNameRule            string                           `json:"unit_name_rule,omitempty"`             // 广告组命名规则，必须包含[日期]和[序号]
	CreativeNameRule        string                           `json:"creative_name_rule,omitempty"`         // 创意命名规则，必须包含[日期]和[序号]
}

func (receiver *CampaignCreateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *CampaignCreateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.CampaignName) == 0 {
		err = errors.New("campaign_name is empty")
		return
	}
	if receiver.Type <= 0 {
		err = errors.New("type is empty")
		return
	}
	return
}

// CampaignCreateResp 创建广告计划响应数据（仅data部分）
type CampaignCreateResp struct {
	CampaignId int64 `json:"campaign_id"` // 广告计划ID
}
