package model

import "errors"

// ========== 创建批量异步请求任务 ==========
// https://developers.e.qq.com/v3.0/docs/api/batch_async_requests/add

// 任务类型枚举常量
const (
	TaskTypeUpdateUnionPositionPackageNew                   = "TASK_TYPE_UPDATE_UNION_POSITION_PACKAGE_NEW"
	TaskTypeUpdateExcludeUnionPositionPackageNew            = "TASK_TYPE_UPDATE_EXCLUDE_UNION_POSITION_PACKAGE_NEW"
	TaskTypeUpdateDeepConversionBehaviorBidNew              = "TASK_TYPE_UPDATE_DEEP_CONVERSION_BEHAVIOR_BID_NEW"
	TaskTypeDeleteAdgroupNew                                = "TASK_TYPE_DELETE_ADGROUP_NEW"
	TaskTypeUpdateAdgroupDeepConversionWorthRateNew         = "TASK_TYPE_UPDATE_ADGROUP_DEEP_CONVERSION_WORTH_RATE_NEW"
	TaskTypeTargetingsShareNew                              = "TASK_TYPE_TARGETINGS_SHARE_NEW"
	TaskTypeUpdateAdgroupConfiguredStatusNew                = "TASK_TYPE_UPDATE_ADGROUP_CONFIGURED_STATUS_NEW"
	TaskTypeUpdateAdgroupDailyBudgetNew                     = "TASK_TYPE_UPDATE_ADGROUP_DAILY_BUDGET_NEW"
	TaskTypeUpdateAdgroupAutoAcquisitionNew                 = "TASK_TYPE_UPDATE_ADGROUP_AUTO_ACQUISITION_NEW"
	TaskTypeUpdateAdgroupDeepConversionWorthAdvancedRateNew = "TASK_TYPE_UPDATE_ADGROUP_DEEP_CONVERSION_WORTH_ADVANCED_RATE_NEW"
	TaskTypeUpdateDeepConversionBehaviorAdvancedBidNew      = "TASK_TYPE_UPDATE_DEEP_CONVERSION_BEHAVIOR_ADVANCED_BID_NEW"
	TaskTypeReplyFinderObjectCommentNew                     = "TASK_TYPE_REPLY_FINDER_OBJECT_COMMENT_NEW"
	TaskTypeDeleteFinderObjectCommentNew                    = "TASK_TYPE_DELETE_FINDER_OBJECT_COMMENT_NEW"
	TaskTypeUpdateFinderObjectCommentFlagNew                = "TASK_TYPE_UPDATE_FINDER_OBJECT_COMMENT_FLAG_NEW"
	TaskTypeUpdateAdgroupTimeNew                            = "TASK_TYPE_UPDATE_ADGROUP_TIME_NEW"
	TaskTypeUpdateAdgroupDateNew                            = "TASK_TYPE_UPDATE_ADGROUP_DATE_NEW"
	TaskTypeUpdateAdgroupBidAmountNew                       = "TASK_TYPE_UPDATE_ADGROUP_BID_AMOUNT_NEW"
	TaskTypeUpdateAdgroupBindRtaPolicyNew                   = "TASK_TYPE_UPDATE_ADGROUP_BIND_RTA_POLICY_NEW"
	TaskTypeUpdateAdcreativeObjectCommentFlagNew            = "TASK_TYPE_UPDATE_ADCREATIVE_OBJECT_COMMENT_FLAG_NEW"
	TaskTypeUpdateDynamicCreativeConfiguredStatusNew        = "TASK_TYPE_UPDATE_DYNAMIC_CREATIVE_CONFIGURED_STATUS_NEW"
	TaskTypeDeleteDynamicCreativeNew                        = "TASK_TYPE_DELETE_DYNAMIC_CREATIVE_NEW"
	TaskTypeProcessUserPageObjectNew                        = "TASK_TYPE_PROCESS_USER_PAGE_OBJECT_NEW"
	TaskTypeCreateScheduledUpdateAdgroupDailyBudgetNew      = "TASK_TYPE_CREATE_SCHEDULED_UPDATE_ADGROUP_DAILY_BUDGET_NEW"
	TaskTypeDeleteScheduledTaskNew                          = "TASK_TYPE_DELETE_SCHEDULED_TASK_NEW"
	TaskTypeUpdateAdgroupTargeting                          = "TASK_TYPE_UPDATE_ADGROUP_TARGETING"
	TaskTypeUpdateAdgroupEcomPkamNew                        = "TASK_TYPE_UPDATE_ADGROUP_ECOM_PKAM_NEW"
	TaskTypeUpdateAdgroupDeriveConfNew                      = "TASK_TYPE_UPDATE_ADGROUP_DERIVE_CONF_NEW"
	TaskTypeUpdateComponentShared                           = "TASK_TYPE_UPDATE_COMPONENT_SHARED"
	TaskTypeUpdateSmartDeliveryGoal                         = "TASK_TYPE_UPDATE_SMART_DELIVERY_GOAL"
)

// 评论精选操作类型枚举
const (
	SetObjectCommentFlagOpTypeOpen  = "SET_OBJECT_COMMENT_FLAG_OP_TYPE_OPEN"
	SetObjectCommentFlagOpTypeClose = "SET_OBJECT_COMMENT_FLAG_OP_TYPE_CLOSE"
)

// 一方跑量开关枚举
const (
	EcomPkamSwitchClose = "ECOM_PKAM_SWITCH_CLOSE"
	EcomPkamSwitchOpen  = "ECOM_PKAM_SWITCH_OPEN"
)

// 字段限制常量
const (
	MaxBatchAsyncTaskNameBytes                = 120      // task_name 最大字节数
	MaxUpdateUnionPositionPackageCount        = 20       // union_position_package 最大长度
	MaxUpdateExcludeUnionPositionPackageCount = 20       // exclude_union_position_package 最大长度
	MaxDeepConversionBehaviorBid              = 1000000  // deep_conversion_behavior_bid 最大值
	MinDeepConversionWorthRate                = 0.001    // deep_conversion_worth_rate 最小值
	MaxDeepConversionWorthRate                = 1000     // deep_conversion_worth_rate 最大值
	MinAutoAcquisitionBudget                  = 20000    // auto_acquisition_budget 最小值
	MaxAutoAcquisitionBudget                  = 10000000 // auto_acquisition_budget 最大值
	MaxDeepConversionBehaviorAdvancedBid      = 1000000  // deep_conversion_behavior_advanced_bid 最大值
)

// ---------- task_spec 子类型 ----------

// UpdateUnionPositionPackageItem 批量修改定投腾讯广告联盟流量包列表信息
type UpdateUnionPositionPackageItem struct {
	AdgroupID            int64   `json:"adgroup_id"`                       // 广告 id (必填)
	UnionPositionPackage []int64 `json:"union_position_package,omitempty"` // 定投腾讯广告联盟流量包 id 列表，最大20
}

// UpdateExcludeUnionPositionPackageItem 批量修改广告屏蔽腾讯广告联盟流量包列表信息
type UpdateExcludeUnionPositionPackageItem struct {
	AdgroupID                   int64   `json:"adgroup_id"`                               // 广告 id (必填)
	ExcludeUnionPositionPackage []int64 `json:"exclude_union_position_package,omitempty"` // 屏蔽腾讯广告联盟流量包 id 列表，最大20
}

// UpdateDeepConversionBehaviorBidItem 批量修改深度优化行为出价
type UpdateDeepConversionBehaviorBidItem struct {
	AdgroupID                 int64 `json:"adgroup_id"`                   // 广告 id (必填)
	DeepConversionBehaviorBid int64 `json:"deep_conversion_behavior_bid"` // 深度优化行为的出价 (必填)，最小值 0，最大值 1000000
}

// DeleteAdgroupItem 批量删除广告
type DeleteAdgroupItem struct {
	AdgroupID int64 `json:"adgroup_id"` // 广告 id (必填)
}

// UpdateAdgroupDeepConversionWorthRateItem 批量修改广告深度优化价值的期望 ROI
type UpdateAdgroupDeepConversionWorthRateItem struct {
	AdgroupID               int64   `json:"adgroup_id"`                 // 广告 id (必填)
	DeepConversionWorthRate float64 `json:"deep_conversion_worth_rate"` // 深度优化价值的期望 ROI (必填)，0.001-1000，最多4位小数
}

// UpdateAdgroupConfiguredStatusItem 批量修改广告客户设置的状态
type UpdateAdgroupConfiguredStatusItem struct {
	AdgroupID        int64  `json:"adgroup_id"`        // 广告 id (必填)
	ConfiguredStatus string `json:"configured_status"` // 客户设置的状态 (必填)，AD_STATUS_NORMAL/AD_STATUS_SUSPEND
}

// UpdateAdgroupDailyBudgetItem 批量修改广告日预算
type UpdateAdgroupDailyBudgetItem struct {
	AdgroupID   int64 `json:"adgroup_id"`   // 广告 id (必填)
	DailyBudget int64 `json:"daily_budget"` // 广告日预算 (必填)
}

// UpdateAdgroupAutoAcquisitionItem 批量修改广告一键起量
type UpdateAdgroupAutoAcquisitionItem struct {
	AdgroupID              int64 `json:"adgroup_id"`                        // 广告 id (必填)
	AutoAcquisitionEnabled bool  `json:"auto_acquisition_enabled"`          // 是否开启一键起量 (必填)
	AutoAcquisitionBudget  int64 `json:"auto_acquisition_budget,omitempty"` // 一键起量探索预算，20000-10000000
}

// UpdateAdgroupDeepConversionWorthAdvancedRateItem 批量修改广告深度优化价值的强化 ROI
type UpdateAdgroupDeepConversionWorthAdvancedRateItem struct {
	AdgroupID                       int64   `json:"adgroup_id"`                          // 广告 id (必填)
	DeepConversionWorthAdvancedRate float64 `json:"deep_conversion_worth_advanced_rate"` // 深度优化价值的强化 ROI (必填)，0.001-1000，最多4位小数
}

// UpdateDeepConversionBehaviorAdvancedBidItem 批量修改深度辅助行为出价
type UpdateDeepConversionBehaviorAdvancedBidItem struct {
	AdgroupID                         int64 `json:"adgroup_id"`                            // 广告 id (必填)
	DeepConversionBehaviorAdvancedBid int64 `json:"deep_conversion_behavior_advanced_bid"` // 深度辅助行为的出价 (必填)，最小值 0，最大值 1000000
}

// ReplyFinderObjectCommentItem 批量评论回复
type ReplyFinderObjectCommentItem struct {
	AccountID        int64  `json:"account_id"`          // 通用的用户 id (必填)
	ReplyCommentID   string `json:"reply_comment_id"`    // 回复评论 id (必填)，1-1024字节
	Content          string `json:"content"`             // 评论内容 (必填)，1-10240字节
	FinderAdObjectID int64  `json:"finder_ad_object_id"` // 广告动态 id (必填)
}

// DeleteFinderObjectCommentItem 批量删除评论
type DeleteFinderObjectCommentItem struct {
	AccountID        int64  `json:"account_id"`          // 通用的用户 id (必填)
	FinderAdObjectID int64  `json:"finder_ad_object_id"` // 广告动态 id (必填)
	CommentID        string `json:"comment_id"`          // 回复评论 id (必填)，1-1024字节
}

// UpdateFinderObjectCommentFlagItem 批量评论精选
type UpdateFinderObjectCommentFlagItem struct {
	FinderAdObjectID int64  `json:"finder_ad_object_id"`     // 广告动态 id (必填)
	OpType           string `json:"op_type"`                 // 操作类型 (必填)
	AccountID        int64  `json:"account_id"`              // 通用的用户 id (必填)
	CommentID        string `json:"comment_id,omitempty"`    // 评论 id，1-1024字节
	CommentLevel     int    `json:"comment_level,omitempty"` // 评论层级，1-2
}

// UpdateAdgroupTimeItem 批量修改广告投放时间
type UpdateAdgroupTimeItem struct {
	AdgroupID         int64  `json:"adgroup_id"`                     // 广告 id (必填)
	AccountID         int64  `json:"account_id,omitempty"`           // 通用的用户 id
	TimeSeries        string `json:"time_series,omitempty"`          // 投放时间段，336字节
	FirstDayBeginTime string `json:"first_day_begin_time,omitempty"` // 首日开始投放时间，0-8字节
}

// UpdateAdgroupDateItem 批量修改广告投放日期
type UpdateAdgroupDateItem struct {
	AdgroupID int64  `json:"adgroup_id"`           // 广告 id (必填)
	BeginDate string `json:"begin_date"`           // 开始投放日期 (必填)，10字节
	EndDate   string `json:"end_date"`             // 结束投放日期 (必填)，0-10字节
	AccountID int64  `json:"account_id,omitempty"` // 通用的用户 id
}

// UpdateAdgroupBidAmountItem 批量修改广告出价
type UpdateAdgroupBidAmountItem struct {
	AdgroupID           int64  `json:"adgroup_id"`                      // 广告 id (必填)
	BidAmount           int64  `json:"bid_amount"`                      // 广告出价 (必填)
	AccountID           int64  `json:"account_id,omitempty"`            // 通用的用户 id
	IsPotential         bool   `json:"is_potential,omitempty"`          // 是否是潜力广告
	ReportPotentialData string `json:"report_potential_data,omitempty"` // 潜力广告改价上报的信息，2-1000字节
}

// UpdateAdgroupBindRtaPolicyItem 批量修改广告绑定的 rta 策略
type UpdateAdgroupBindRtaPolicyItem struct {
	AdgroupID           int64  `json:"adgroup_id"`             // 广告 id (必填)
	OriginRtaPolicyUUID string `json:"origin_rta_policy_uuid"` // 源 rta 策略 (必填)，0-1024字节
	TargetRtaPolicyUUID string `json:"target_rta_policy_uuid"` // 目标 rta 策略 (必填)，0-1024字节
	AccountID           int64  `json:"account_id,omitempty"`   // 推广帐号 id
}

// UpdateAdcreativeObjectCommentFlagItem 批量评论管理
type UpdateAdcreativeObjectCommentFlagItem struct {
	FinderAdObjectID int64  `json:"finder_ad_object_id"`  // 广告动态 id (必填)
	OpType           string `json:"op_type"`              // 操作类型 (必填)
	AccountID        int64  `json:"account_id,omitempty"` // 通用的用户 id
}

// UpdateDynamicCreativeConfiguredStatusItem 批量修改广告创意客户设置的状态
type UpdateDynamicCreativeConfiguredStatusItem struct {
	DynamicCreativeID int64  `json:"dynamic_creative_id"` // 广告创意 id (必填)
	ConfiguredStatus  string `json:"configured_status"`   // 客户设置的状态 (必填)
}

// DeleteDynamicCreativeItem 批量删除广告创意
type DeleteDynamicCreativeItem struct {
	DynamicCreativeID int64 `json:"dynamic_creative_id"` // 广告创意 id (必填)
}

// ProcessUserPageObjectItem 批量沉淀动态
type ProcessUserPageObjectItem struct {
	AdExportID string `json:"ad_export_id"`         // 视频号动态 id (必填)，0-1024字节
	AccountID  int64  `json:"account_id,omitempty"` // 通用的用户 id
}

// CreateScheduledUpdateAdgroupDailyBudgetItem 批量预设置广告次日预算
type CreateScheduledUpdateAdgroupDailyBudgetItem struct {
	AdgroupID   int64 `json:"adgroup_id"`           // 广告 id (必填)
	DailyBudget int64 `json:"daily_budget"`         // 广告日预算 (必填)
	AccountID   int64 `json:"account_id,omitempty"` // 通用的用户 id
}

// DeleteScheduledTaskItem 批量删除预设置任务
type DeleteScheduledTaskItem struct {
	TaskID    int64 `json:"task_id"`              // 预设置任务 id (必填)
	AccountID int64 `json:"account_id,omitempty"` // 通用的用户 id
}

// UpdateAdgroupEcomPkamItem 批量修改广告一方跑量
type UpdateAdgroupEcomPkamItem struct {
	AdgroupID      int64  `json:"adgroup_id"`       // 广告 id (必填)
	EcomPkamSwitch string `json:"ecom_pkam_switch"` // 是否开启一方跑量 (必填)
}

// DeriveTemplateConfItem 衍生模版配置项
type DeriveTemplateConfItem struct {
	TemplateList         []int64 `json:"template_list,omitempty"` // 创意衍生自定义模版列表，最大256
	AdcreativeTemplateID int64   `json:"adcreative_template_id"`  // 创意衍生使用的原规格 id (必填)
}

// DeriveTemplateConf 衍生模版配置
type DeriveTemplateConf struct {
	TemplateConfList []*DeriveTemplateConfItem `json:"template_conf_list"` // 衍生配置列表 (必填)，最大10
}

// UpdateAdgroupDeriveConfItem 批量修改广告自动衍生视频创意
type UpdateAdgroupDeriveConfItem struct {
	AdgroupID                  int64               `json:"adgroup_id"`                     // 广告 id (必填)
	AutoDerivedCreativeEnabled bool                `json:"auto_derived_creative_enabled"`  // 是否开启自动衍生视频创意 (必填)
	DeriveTemplateConf         *DeriveTemplateConf `json:"derive_template_conf,omitempty"` // 衍生模版配置
	AccountID                  int64               `json:"account_id,omitempty"`           // 通用的用户 id
}

// SharedAccountItem 被共享账号信息
type SharedAccountItem struct {
	SharedAccountID   int64  `json:"shared_account_id"`   // 被共享账号 id (必填)，0-9999999999
	SharedAccountType string `json:"shared_account_type"` // 被共享账号类型 (必填)
}

// UpdateComponentSharingItem 更新组件共享
type UpdateComponentSharingItem struct {
	OrganizationID    int64                `json:"organization_id"`     // 业务单元 id (必填)，0-9999999999
	SharedAccountList []*SharedAccountItem `json:"shared_account_list"` // 被共享账号信息 (必填)，最大100
	ComponentID       int64                `json:"component_id"`        // 组件 id (必填)
}

// BatchAsyncTaskSpec 批量异步任务所需条件（根据 task_type 设置对应字段）
type BatchAsyncTaskSpec struct {
	UpdateUnionPositionPackageSpec                   []*UpdateUnionPositionPackageItem                   `json:"update_union_position_package_spec,omitempty"`
	UpdateExcludeUnionPositionPackageSpec            []*UpdateExcludeUnionPositionPackageItem            `json:"update_exclude_union_position_package_spec,omitempty"`
	UpdateDeepConversionBehaviorBidSpec              []*UpdateDeepConversionBehaviorBidItem              `json:"update_deep_conversion_behavior_bid_spec,omitempty"`
	DeleteAdgroupSpec                                []*DeleteAdgroupItem                                `json:"delete_adgroup_spec,omitempty"`
	UpdateAdgroupDeepConversionWorthRateSpec         []*UpdateAdgroupDeepConversionWorthRateItem         `json:"update_adgroup_deep_conversion_worth_rate_spec,omitempty"`
	UpdateAdgroupConfiguredStatusSpec                []*UpdateAdgroupConfiguredStatusItem                `json:"update_adgroup_configured_status_spec,omitempty"`
	UpdateAdgroupDailyBudgetSpec                     []*UpdateAdgroupDailyBudgetItem                     `json:"update_adgroup_daily_budget_spec,omitempty"`
	UpdateAdgroupAutoAcquisitionSpec                 []*UpdateAdgroupAutoAcquisitionItem                 `json:"update_adgroup_auto_acquisition_spec,omitempty"`
	UpdateAdgroupDeepConversionWorthAdvancedRateSpec []*UpdateAdgroupDeepConversionWorthAdvancedRateItem `json:"update_adgroup_deep_conversion_worth_advanced_rate_spec,omitempty"`
	UpdateDeepConversionBehaviorAdvancedBidSpec      []*UpdateDeepConversionBehaviorAdvancedBidItem      `json:"update_deep_conversion_behavior_advanced_bid_spec,omitempty"`
	ReplyFinderObjectCommentSpec                     []*ReplyFinderObjectCommentItem                     `json:"reply_finder_object_comment_spec,omitempty"`
	DeleteFinderObjectCommentSpec                    []*DeleteFinderObjectCommentItem                    `json:"delete_finder_object_comment_spec,omitempty"`
	UpdateFinderObjectCommentFlagSpec                []*UpdateFinderObjectCommentFlagItem                `json:"update_finder_object_comment_flag_spec,omitempty"`
	UpdateAdgroupTimeSpec                            []*UpdateAdgroupTimeItem                            `json:"update_adgroup_time_spec,omitempty"`
	UpdateAdgroupDateSpec                            []*UpdateAdgroupDateItem                            `json:"update_adgroup_date_spec,omitempty"`
	UpdateAdgroupBidAmountSpec                       []*UpdateAdgroupBidAmountItem                       `json:"update_adgroup_bid_amount_spec,omitempty"`
	UpdateAdgroupBindRtaPolicySpec                   []*UpdateAdgroupBindRtaPolicyItem                   `json:"update_adgroup_bind_rta_policy_spec,omitempty"`
	UpdateAdcreativeObjectCommentFlagSpec            []*UpdateAdcreativeObjectCommentFlagItem            `json:"update_adcreative_object_comment_flag_spec,omitempty"`
	UpdateDynamicCreativeConfiguredStatusSpec        []*UpdateDynamicCreativeConfiguredStatusItem        `json:"update_dynamic_creative_configured_status_spec,omitempty"`
	DeleteDynamicCreativeSpec                        []*DeleteDynamicCreativeItem                        `json:"delete_dynamic_creative_spec,omitempty"`
	ProcessUserPageObjectSpec                        []*ProcessUserPageObjectItem                        `json:"process_user_page_object_spec,omitempty"`
	CreateScheduledUpdateAdgroupDailyBudgetSpec      []*CreateScheduledUpdateAdgroupDailyBudgetItem      `json:"create_scheduled_update_adgroup_daily_budget_spec,omitempty"`
	DeleteScheduledTaskSpec                          []*DeleteScheduledTaskItem                          `json:"delete_scheduled_task_spec,omitempty"`
	UpdateAdgroupEcomPkamSpec                        []*UpdateAdgroupEcomPkamItem                        `json:"update_adgroup_ecom_pkam_spec,omitempty"`
	UpdateAdgroupDeriveConfSpec                      []*UpdateAdgroupDeriveConfItem                      `json:"update_adgroup_derive_conf_spec,omitempty"`
	UpdateComponentSharingSpec                       []*UpdateComponentSharingItem                       `json:"update_component_sharing_spec,omitempty"`
}

// BatchAsyncRequestAddReq 创建批量异步请求任务请求
// https://developers.e.qq.com/v3.0/docs/api/batch_async_requests/add
type BatchAsyncRequestAddReq struct {
	GlobalReq
	AccountID int64               `json:"account_id"` // 广告主帐号 id (必填)
	TaskName  string              `json:"task_name"`  // 任务名称 (必填)，1-120字节
	TaskType  string              `json:"task_type"`  // 任务类型 (必填)
	TaskSpec  *BatchAsyncTaskSpec `json:"task_spec"`  // 任务所需条件 (必填)
}

func (p *BatchAsyncRequestAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证创建批量异步请求任务请求参数
func (p *BatchAsyncRequestAddReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.TaskName == "" {
		return errors.New("task_name为必填")
	}
	if len(p.TaskName) > MaxBatchAsyncTaskNameBytes {
		return errors.New("task_name长度不能超过120字节")
	}
	if p.TaskType == "" {
		return errors.New("task_type为必填")
	}
	if p.TaskSpec == nil {
		return errors.New("task_spec为必填")
	}
	return p.GlobalReq.Validate()
}

// BatchAsyncRequestAddResp 创建批量异步请求任务响应
// https://developers.e.qq.com/v3.0/docs/api/batch_async_requests/add
type BatchAsyncRequestAddResp struct {
	TaskID int64 `json:"task_id"` // 任务 id
}
