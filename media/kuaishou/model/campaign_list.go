package model

import "errors"

// CampaignListReq 查询广告计划请求
type CampaignListReq struct {
	accessTokenReq
	AdvertiserId   int64   `json:"advertiser_id"`              // 广告主ID，必填，在获取 access_token 时返回
	CampaignId     int64   `json:"campaign_id,omitempty"`      // 广告计划ID，筛选条件，不传或为空则不限制
	CampaignIds    []int64 `json:"campaign_ids,omitempty"`     // 广告计划ID数组，最多200个
	CampaignName   string  `json:"campaign_name,omitempty"`    // 广告计划名称，筛选条件，不传或为空则不限制
	AdType         int     `json:"ad_type,omitempty"`          // 广告类型：0=信息流 1=搜索
	CampaignType   int     `json:"campaign_type,omitempty"`    // 计划类型
	Status         int     `json:"status,omitempty"`           // 计划状态：1=暂停 4=投放中 5=已删除 -2=不限
	PutStatusList  []int   `json:"put_status_list,omitempty"`  // 投放状态筛选：1=投放 2=暂停 3=删除；传入时覆盖 status 参数
	StartDate      string  `json:"start_date,omitempty"`       // 开始日期，格式：yyyy-MM-dd，需与 end_date 同时传或同时不传
	EndDate        string  `json:"end_date,omitempty"`         // 结束日期，格式：yyyy-MM-dd，需与 start_date 同时传或同时不传
	TimeFilterType int     `json:"time_filter_type,omitempty"` // 时间筛选类型：0/不传=更新时间 1=创建时间
	Page           int     `json:"page,omitempty"`             // 页码，默认1
	PageSize       int     `json:"page_size,omitempty"`        // 每页数量，默认20
}

func (receiver *CampaignListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *CampaignListReq) Validate() (err error) {
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

// CampaignListPhotoPackageDetail 素材包详情
type CampaignListPhotoPackageDetail struct {
	PhotoPackageId int64    `json:"photo_package_id"` // 素材包ID
	Name           string   `json:"name"`             // 素材包名称
	Status         int      `json:"status"`           // 素材包状态：0=已删除 1=有效
	PhotoIdInfo    []string `json:"photo_id_info"`    // 素材包中的视频ID列表
}

// CampaignDetail 广告计划详情
type CampaignDetail struct {
	CampaignId                     int64                            `json:"campaign_id"`                        // 广告计划ID
	CampaignName                   string                           `json:"campaign_name"`                      // 广告计划名称
	CampaignType                   int                              `json:"campaign_type"`                      // 计划类型
	CampaignSubType                int                              `json:"campaign_sub_type"`                  // 计划子类型：4=DPA 5=SDPA
	AdType                         int                              `json:"ad_type"`                            // 广告类型：0=信息流 1=搜索
	CreateTime                     string                           `json:"create_time"`                        // 创建时间，格式：yyyy-MM-dd HH:mm:ss
	UpdateTime                     string                           `json:"update_time"`                        // 更新时间，格式：yyyy-MM-dd HH:mm:ss
	PutStatus                      int                              `json:"put_status"`                         // 投放状态：1=投放 2=暂停 3=删除
	Status                         int                              `json:"status"`                             // 计划状态：1=暂停 3=超预算 4=投放中 5=已删除 6=余额不足 -2=不限
	DayBudget                      int64                            `json:"day_budget"`                         // 日预算，单位：厘
	DayBudgetSchedule              []int64                          `json:"day_budget_schedule"`                // 分日预算（7日数组），单位：厘，优先级高于 day_budget
	BidType                        int                              `json:"bid_type"`                           // 出价类型
	ConstraintActionType           int                              `json:"constraint_action_type"`             // 成本约束目标：0=无 1=千次曝光 2=点击 53=表单 180=激活 190=付费 191=ROI
	ConstraintCpa                  int64                            `json:"constraint_cpa"`                     // 行为成本约束，单位：厘，范围：0-8000元
	CapRoiRatio                    float64                          `json:"cap_roi_ratio"`                      // ROI成本约束，范围：0-100
	CapBid                         int64                            `json:"cap_bid"`                            // 成本上限约束，单位：厘，范围：0-8000元
	AutoManage                     int                              `json:"auto_manage"`                        // 智能优化开关：0=关闭 1=开启
	CampaignOcpxActionType         int                              `json:"campaign_ocpx_action_type"`          // 智能优化浅层目标
	CampaignOcpxActionTypeName     string                           `json:"campaign_ocpx_action_type_name"`     // 智能优化浅层目标名称
	CampaignDeepConversionType     int                              `json:"campaign_deep_conversion_type"`      // 智能优化深层目标
	CampaignDeepConversionTypeName string                           `json:"campaign_deep_conversion_type_name"` // 智能优化深层目标名称
	AutoAdjust                     int                              `json:"auto_adjust"`                        // 自动调节开关：0=关闭 1=开启
	AutoBuild                      int                              `json:"auto_build"`                         // 自动建组开关：0=关闭 1=开启
	AutoBuildNameRule              *CampaignCreateAutoBuildNameRule `json:"auto_build_name_rule"`               // 自动建组命名规则
	UnitNameRule                   string                           `json:"unit_name_rule"`                     // 广告组命名规则，包含[日期]和[序号]宏
	CreativeNameRule               string                           `json:"creative_name_rule"`                 // 创意命名规则，包含[日期]和[序号]宏
	DspVersion                     int                              `json:"dsp_version"`                        // 版本号
	AutoPhotoScope                 int                              `json:"auto_photo_scope"`                   // 建组取材范围：0=系统最优 1=素材包
	PhotoPackageDetails            []CampaignListPhotoPackageDetail `json:"photo_package_details"`              // 素材包详情列表
	SmartMaterialSupply            int                              `json:"smart_material_supply"`              // 智能素材补充开关
	PeriodicDeliveryType           int                              `json:"periodic_delivery_type"`             // 周期稳定投放开关
	PeriodicDeliveryPutType        int                              `json:"periodic_delivery_put_type"`         // 夜间投放标识：1=常规 2=夜间稳投
	ContinuePeriodType             int                              `json:"continue_period_type"`               // 周期延续开关
	RangeBudget                    int64                            `json:"range_budget"`                       // 周期总预算，单位：元
	PeriodicDays                   int                              `json:"periodic_days"`                      // 投放周期天数
}

// CampaignListResp 查询广告计划响应数据（仅data部分）
type CampaignListResp struct {
	TotalCount int64            `json:"total_count"` // 广告计划总数量
	Details    []CampaignDetail `json:"details"`     // 广告计划详情列表
}
