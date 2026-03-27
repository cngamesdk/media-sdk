package model

import "errors"

type AdgroupsGetReq struct {
	GlobalReq
	CursorPageV2Req
	AccountID int64                 `json:"account_id"`           // 广告主帐号id (必填)
	IsDeleted bool                  `json:"is_deleted,omitempty"` // 是否已删除
	Fields    []string              `json:"fields,omitempty"`     // 指定返回的字段列表
	Filtering []*AdgroupQueryFilter `json:"filtering,omitempty"`  // 过滤条件
}

// AdgroupQueryFilter 广告组查询过滤条件
type AdgroupQueryFilter struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)
	Operator string   `json:"operator"` // 操作符 (必填)
	Values   []string `json:"values"`   // 字段取值 (必填)
}

// 常量定义 - 过滤字段
const (
	FieldAdgroupId                  = "adgroup_id"
	FieldAdgroupName                = "adgroup_name"
	FieldCreatedTime                = "created_time"
	FieldLastModifiedTime           = "last_modified_time"
	FieldMaterialPackageId          = "material_package_id"
	FieldJointBudgetRuleId          = "joint_budget_rule_id"
	FieldConfiguredStatus           = "configured_status"
	FieldAutoDerivedCreativeEnabled = "auto_derived_creative_enabled"
	FieldSmartDeliveryPlatform      = "smart_delivery_platform"
)

// 常量定义 - 操作符
const (
	OperatorEquals        = "EQUALS"
	OperatorIn            = "IN"
	OperatorContains      = "CONTAINS"
	OperatorLess          = "LESS"
	OperatorLessEquals    = "LESS_EQUALS"
	OperatorGreater       = "GREATER"
	OperatorGreaterEquals = "GREATER_EQUALS"
)

// 常量定义 - 配置状态
const (
	ConfiguredStatusNormal  = "AD_STATUS_NORMAL"  // 正常
	ConfiguredStatusSuspend = "AD_STATUS_SUSPEND" // 暂停
)

// 长度限制常量
const (
	MinFilteringCount  = 1
	MaxFilteringCount  = 255
	MinValuesCount     = 1
	MaxValuesCount     = 100
	CreatedTimeLength  = 10
	ModifiedTimeLength = 10
)

func (p *AdgroupsGetReq) Format() {
	p.GlobalReq.Format()
	p.CursorPageV2Req.Format()
}

// Validate 验证广告组查询参数
func (p *AdgroupsGetReq) Validate() error {
	// 如果过滤条件为空，视为无限制条件
	if len(p.Filtering) == 0 {
		return nil
	}

	// 验证过滤条件数量
	if len(p.Filtering) < MinFilteringCount || len(p.Filtering) > MaxFilteringCount {
		return errors.New("filtering数组长度必须在1-255之间")
	}

	// 验证每个过滤条件
	for i, filter := range p.Filtering {
		if err := filter.Validate(); err != nil {
			return errors.New("filtering[" + string(rune(i)) + "]验证失败: " + err.Error())
		}
	}

	// 验证account_id
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	if validateErr := p.CursorPageV2Req.Validate(); validateErr != nil {
		return validateErr
	}

	// 验证fields
	if err := p.validateFields(); err != nil {
		return err
	}

	return nil
}

// validateFields 验证字段列表
func (p *AdgroupsGetReq) validateFields() error {
	if len(p.Fields) == 0 {
		return nil
	}
	if len(p.Fields) < MinAdgroupFieldsCount || len(p.Fields) > MaxAdgroupFieldsCount {
		return errors.New("fields数组长度必须在1-1024之间")
	}
	for _, field := range p.Fields {
		if len(field) < MinAdgroupFieldLength || len(field) > MaxAdgroupFieldLength {
			return errors.New("fields中的字段长度必须在1-64字节之间")
		}
	}
	return nil
}

// Validate 验证单个过滤条件
func (f *AdgroupQueryFilter) Validate() error {
	// 1. 验证字段
	if f.Field == "" {
		return errors.New("field为必填")
	}
	if !isValidField(f.Field) {
		return errors.New("field值无效，请参考文档中的允许值")
	}

	// 2. 验证操作符
	if f.Operator == "" {
		return errors.New("operator为必填")
	}
	if !isValidOperatorForField(f.Field, f.Operator) {
		return errors.New("operator值无效，当前字段不支持该操作符")
	}

	// 3. 验证values
	if len(f.Values) == 0 {
		return errors.New("values为必填")
	}
	if err := validateValuesForField(f); err != nil {
		return err
	}

	return nil
}

// isValidField 验证字段是否有效
func isValidField(field string) bool {
	validFields := map[string]bool{
		FieldAdgroupId:                  true,
		FieldAdgroupName:                true,
		FieldCreatedTime:                true,
		FieldLastModifiedTime:           true,
		FieldMaterialPackageId:          true,
		FieldJointBudgetRuleId:          true,
		FieldConfiguredStatus:           true,
		FieldAutoDerivedCreativeEnabled: true,
		FieldSmartDeliveryPlatform:      true,
	}
	return validFields[field]
}

// isValidOperatorForField 验证字段支持的操作符
func isValidOperatorForField(field, operator string) bool {
	switch field {
	case FieldAdgroupId:
		return operator == OperatorEquals || operator == OperatorIn
	case FieldAdgroupName:
		return operator == OperatorEquals || operator == OperatorContains
	case FieldCreatedTime, FieldLastModifiedTime:
		return operator == OperatorEquals || operator == OperatorLessEquals ||
			operator == OperatorLess || operator == OperatorGreaterEquals ||
			operator == OperatorGreater
	case FieldMaterialPackageId:
		return operator == OperatorEquals || operator == OperatorLessEquals ||
			operator == OperatorLess || operator == OperatorGreaterEquals ||
			operator == OperatorGreater
	case FieldJointBudgetRuleId:
		return operator == OperatorEquals || operator == OperatorIn
	case FieldConfiguredStatus:
		return operator == OperatorEquals
	case FieldAutoDerivedCreativeEnabled:
		return operator == OperatorEquals
	case FieldSmartDeliveryPlatform:
		return operator == OperatorEquals || operator == OperatorIn || operator == OperatorGreaterEquals
	default:
		return false
	}
}

// validateValuesForField 验证字段取值
func validateValuesForField(f *AdgroupQueryFilter) error {
	switch f.Field {
	case FieldAdgroupId:
		return validateAdgroupIdValues(f)
	case FieldAdgroupName:
		return validateAdgroupNameValues(f)
	case FieldCreatedTime, FieldLastModifiedTime:
		return validateTimeValues(f)
	case FieldMaterialPackageId:
		return validateMaterialPackageIdValues(f)
	case FieldJointBudgetRuleId:
		return validateJointBudgetRuleIdValues(f)
	case FieldConfiguredStatus:
		return validateConfiguredStatusValues(f)
	case FieldAutoDerivedCreativeEnabled:
		return validateAutoDerivedCreativeEnabledValues(f)
	case FieldSmartDeliveryPlatform:
		return validateSmartDeliveryPlatformValues(f)
	}
	return nil
}

// validateAdgroupIdValues 验证广告组ID取值
func validateAdgroupIdValues(f *AdgroupQueryFilter) error {
	switch f.Operator {
	case OperatorEquals:
		if len(f.Values) != 1 {
			return errors.New("operator为EQUALS时，values数组长度必须为1")
		}
	case OperatorIn:
		if len(f.Values) < MinValuesCount || len(f.Values) > MaxValuesCount {
			return errors.New("operator为IN时，values数组长度必须在1-100之间")
		}
	}
	return nil
}

// validateAdgroupNameValues 验证广告组名称取值
func validateAdgroupNameValues(f *AdgroupQueryFilter) error {
	if len(f.Values) != 1 {
		return errors.New("values数组长度必须为1")
	}
	for _, v := range f.Values {
		if len(v) < 1 || len(v) > MaxFieldLength {
			return errors.New("字段长度必须在1-180字节之间")
		}
	}
	return nil
}

// validateTimeValues 验证时间取值
func validateTimeValues(f *AdgroupQueryFilter) error {
	if len(f.Values) != 1 {
		return errors.New("values数组长度必须为1")
	}
	for _, v := range f.Values {
		if len(v) != CreatedTimeLength {
			return errors.New("时间字段长度必须为10字节")
		}
	}
	return nil
}

// validateMaterialPackageIdValues 验证素材包ID取值
func validateMaterialPackageIdValues(f *AdgroupQueryFilter) error {
	if len(f.Values) != 1 {
		return errors.New("values数组长度必须为1")
	}
	return nil
}

// validateJointBudgetRuleIdValues 验证联合预算规则ID取值
func validateJointBudgetRuleIdValues(f *AdgroupQueryFilter) error {
	switch f.Operator {
	case OperatorEquals, OperatorIn:
		if len(f.Values) < MinValuesCount || len(f.Values) > MaxValuesCount {
			return errors.New("values数组长度必须在1-100之间")
		}
	}
	return nil
}

// validateConfiguredStatusValues 验证配置状态取值
func validateConfiguredStatusValues(f *AdgroupQueryFilter) error {
	if len(f.Values) != 1 {
		return errors.New("values数组长度必须为1")
	}
	if f.Values[0] != ConfiguredStatusNormal && f.Values[0] != ConfiguredStatusSuspend {
		return errors.New("configured_status值无效，允许值：AD_STATUS_NORMAL、AD_STATUS_SUSPEND")
	}
	return nil
}

// validateAutoDerivedCreativeEnabledValues 验证自动衍生创意启用状态取值
func validateAutoDerivedCreativeEnabledValues(f *AdgroupQueryFilter) error {
	if len(f.Values) != 1 {
		return errors.New("values数组长度必须为1")
	}
	if len(f.Values[0]) < 1 || len(f.Values[0]) > MaxFieldLength {
		return errors.New("字段长度必须在1-180字节之间")
	}
	return nil
}

// validateSmartDeliveryPlatformValues 验证智能投放平台取值
func validateSmartDeliveryPlatformValues(f *AdgroupQueryFilter) error {
	switch f.Operator {
	case OperatorEquals:
		if len(f.Values) != 1 {
			return errors.New("operator为EQUALS时，values数组长度必须为1")
		}
		if len(f.Values[0]) < 1 || len(f.Values[0]) > MaxFieldLength {
			return errors.New("字段长度必须在1-180字节之间")
		}
	case OperatorIn:
		if len(f.Values) < MinValuesCount || len(f.Values) > MaxValuesCount {
			return errors.New("operator为IN时，values数组长度必须在1-100之间")
		}
		for _, v := range f.Values {
			if len(v) < 1 || len(v) > MaxFieldLength {
				return errors.New("字段长度必须在1-180字节之间")
			}
		}
	case OperatorGreaterEquals:
		if len(f.Values) != 1 {
			return errors.New("operator为GREATER_EQUALS时，values数组长度必须为1")
		}
		if len(f.Values[0]) < 1 || len(f.Values[0]) > MaxFieldLength {
			return errors.New("字段长度必须在1-180字节之间")
		}
	}
	return nil
}

// 分页限制常量
const (
	DefaultPage     = 1
	DefaultPageSize = 10

	MinAdgroupFieldsCount = 1
	MaxAdgroupFieldsCount = 1024
	MinAdgroupFieldLength = 1
	MaxAdgroupFieldLength = 64

	MaxCursorLength = 10
)

// 常量定义 - 地点类型
const (
	LocationTypeLiveIn    = "LIVE_IN"    // 常住
	LocationTypeVisitedIn = "VISITED_IN" // 去过
)

type AdgroupsGetResp struct {
	List []*AdgroupsGetListItem `json:"list,omitempty"` // 列表
	PageInfoContainer
	CursorPageInfoV2Container
}

type AdgroupsGetListItem struct {
	Targeting *Targeting `json:"targeting,omitempty"` // 定向详细设置

	AdgroupID                         int64                   `json:"adgroup_id"`                                      // 广告id
	TargetingTranslation              string                  `json:"targeting_translation,omitempty"`                 // 已选择定向条件的描述
	ConfiguredStatus                  string                  `json:"configured_status"`                               // 客户设置的状态
	CreatedTime                       int64                   `json:"created_time"`                                    // 创建时间，时间戳
	LastModifiedTime                  int64                   `json:"last_modified_time"`                              // 最后修改时间，时间戳
	IsDeleted                         bool                    `json:"is_deleted"`                                      // 是否已删除
	SystemStatus                      string                  `json:"system_status"`                                   // 广告在系统中的状态
	AdgroupName                       string                  `json:"adgroup_name"`                                    // 广告名称
	MarketingGoal                     string                  `json:"marketing_goal"`                                  // 营销目的类型
	MarketingSubGoal                  string                  `json:"marketing_sub_goal,omitempty"`                    // 二级营销目的类型
	MarketingCarrierType              string                  `json:"marketing_carrier_type"`                          // 营销载体类型
	MarketingCarrierDetail            *MarketingCarrierDetail `json:"marketing_carrier_detail,omitempty"`              // 营销载体详情
	MarketingTargetType               string                  `json:"marketing_target_type"`                           // 推广产品类型
	MarketingTargetDetail             *MarketingTargetDetail  `json:"marketing_target_detail,omitempty"`               // 营销对象详情
	MarketingTargetID                 int64                   `json:"marketing_target_id,omitempty"`                   // 营销对象id
	BeginDate                         string                  `json:"begin_date"`                                      // 开始投放日期 (必填)
	EndDate                           string                  `json:"end_date"`                                        // 结束投放日期 (必填)
	FirstDayBeginTime                 string                  `json:"first_day_begin_time,omitempty"`                  // 首日开始投放时间
	BidAmount                         int64                   `json:"bid_amount"`                                      // 广告出价，单位分 (必填)
	OptimizationGoal                  string                  `json:"optimization_goal"`                               // 广告优化目标类型 (必填)
	TimeSeries                        string                  `json:"time_series"`                                     // 投放时间段 (必填)
	AutomaticSiteEnabled              bool                    `json:"automatic_site_enabled"`                          // 是否开启智能版位功能
	SiteSet                           []string                `json:"site_set,omitempty"`                              // 投放站点集合
	DailyBudget                       int64                   `json:"daily_budget"`                                    // 日预算，单位分
	SceneSpec                         *SceneSpec              `json:"scene_spec,omitempty"`                            // 场景定向
	UserActionSets                    []*UserActionSet        `json:"user_action_sets,omitempty"`                      // 用户行为数据源
	DeepConversionSpec                *DeepConversionSpec     `json:"deep_conversion_spec,omitempty"`                  // oCPA 深度优化内容
	ConversionID                      int64                   `json:"conversion_id,omitempty"`                         // 转化id
	DeepConversionBehaviorBid         int64                   `json:"deep_conversion_behavior_bid,omitempty"`          // 深度优化行为出价，单位分
	DeepConversionWorthRate           float64                 `json:"deep_conversion_worth_rate,omitempty"`            // 深度优化价值出价
	DeepConversionWorthAdvancedRate   float64                 `json:"deep_conversion_worth_advanced_rate,omitempty"`   // 强化优化价值的期望ROI
	DeepConversionBehaviorAdvancedBid int64                   `json:"deep_conversion_behavior_advanced_bid,omitempty"` // 深度辅助优化OG出价，单位分
	BidMode                           string                  `json:"bid_mode"`                                        // 出价方式 (必填)
	AutoAcquisitionEnabled            bool                    `json:"auto_acquisition_enabled,omitempty"`              // 一键起量开关
	AutoAcquisitionBudget             int64                   `json:"auto_acquisition_budget,omitempty"`               // 一键起量预算，单位分
	SmartBidType                      string                  `json:"smart_bid_type,omitempty"`                        // 出价类型
	SmartCostCap                      int64                   `json:"smart_cost_cap,omitempty"`                        // 自动出价下预计成本上限，单位分
	AutoDerivedCreativeEnabled        bool                    `json:"auto_derived_creative_enabled,omitempty"`         // 创意增强MAX开关

	AutoDerivedCreativePreference *AutoDerivedCreativePreference `json:"auto_derived_creative_preference,omitempty"` // 创意增强MAX偏好设置
	SearchExpandTargetingSwitch   string                         `json:"search_expand_targeting_switch,omitempty"`   // 搜索定向拓展开关
	AutoDerivedLandingPageSwitch  bool                           `json:"auto_derived_landing_page_switch,omitempty"` // 是否开启自动衍生落地页开关
	DataModelVersion              int64                          `json:"data_model_version,omitempty"`               // 数据版本号
	BidScene                      string                         `json:"bid_scene,omitempty"`                        // 出价场景
	MarketingTargetExt            *MarketingTargetExt            `json:"marketing_target_ext,omitempty"`             // 营销对象扩展数据
	DeepOptimizationType          string                         `json:"deep_optimization_type,omitempty"`           // 深度优化策略类型
	FlowOptimizationEnabled       bool                           `json:"flow_optimization_enabled,omitempty"`        // 是否使用自动流量优选（已废弃）

	MarketingTargetAttachment *MarketingTargetAttachment `json:"marketing_target_attachment,omitempty"` // 营销对象附加信息
	NegativeWordCnt           *NegativeWordCnt           `json:"negative_word_cnt,omitempty"`           // 否定词个数
	SearchExpansionSwitch     string                     `json:"search_expansion_switch,omitempty"`     // 搜索引擎开关
	MarketingAssetID          int64                      `json:"marketing_asset_id,omitempty"`          // 产品id
	PromotedAssetType         string                     `json:"promoted_asset_type,omitempty"`         // 推广内容类型
	MaterialPackageID         int64                      `json:"material_package_id,omitempty"`         // 素材标签id

	MarketingAssetOuterSpec *MarketingAssetOuterSpec `json:"marketing_asset_outer_spec,omitempty"` // 产品外部id数据
	PoiList                 []string                 `json:"poi_list,omitempty"`                   // 门店id列表
	MarketingScene          string                   `json:"marketing_scene,omitempty"`            // 营销目标
	ExplorationStrategy     string                   `json:"exploration_strategy,omitempty"`       // 探索策略
	PrioritySiteSet         []string                 `json:"priority_site_set,omitempty"`          // 投放站点集合
	EcomPkamSwitch          string                   `json:"ecom_pkam_switch,omitempty"`           // 一方人群跑量加强开关状态
	ForwardLinkAssist       string                   `json:"forward_link_assist,omitempty"`        // 助攻行为目标
	ConversionName          string                   `json:"conversion_name,omitempty"`            // 转化名称

	AutoAcquisitionStatus string   `json:"auto_acquisition_status,omitempty"` // 一键起量状态
	CostConstraintScene   string   `json:"cost_constraint_scene,omitempty"`   // 成本控制场景
	CustomCostCap         int64    `json:"custom_cost_cap,omitempty"`         // 用户输入的成本上限，单位分
	MpaSpec               *MpaSpec `json:"mpa_spec,omitempty"`                // 动态商品广告属性
	ShortPlayPayType      string   `json:"short_play_pay_type,omitempty"`     // 售卖方式类型
	SellStrategyID        int64    `json:"sell_strategy_id,omitempty"`        // 售卖策略id
	OgCompletionType      string   `json:"og_completion_type,omitempty"`      // 达成类型
	DcaSpec               *DcaSpec `json:"dca_spec,omitempty"`                // 动态内容广告属性

	AoiOptimizationStrategy      *AoiOptimizationStrategy `json:"aoi_optimization_strategy,omitempty"`       // 高价值范围探索
	CostGuaranteeStatus          string                   `json:"cost_guarantee_status,omitempty"`           // 成本保障状态
	CostGuaranteeMoney           int64                    `json:"cost_guarantee_money,omitempty"`            // 成本保障赔付金额，单位分
	AdditionalProductSpec        *AdditionalProductSpec   `json:"additional_product_spec,omitempty"`         // 附加商品属性
	EnableBreakthroughSiteset    bool                     `json:"enable_breakthrough_siteset,omitempty"`     // 是否支持版位突破
	LiveRecommendStrategyEnabled bool                     `json:"live_recommend_strategy_enabled,omitempty"` // 直播种草人群探索
	CustomCostRolCap             float64                  `json:"custom_cost_rol_cap,omitempty"`             // 控制成本的期望ROI
	EnableSteadyExploration      bool                     `json:"enable_steady_exploration,omitempty"`       // 是否稳步探索更多版位
	AdxRealtimeType              string                   `json:"adx_realtime_type,omitempty"`               // ADX程序化广告素材实时回复类型
	SmartTargetingStatus         string                   `json:"smart_targeting_status,omitempty"`          // 广告智能定向状态
}

// AoiOptimizationStrategy 高价值范围探索
type AoiOptimizationStrategy struct {
	AoiOptimizationStrategyEnabled bool    `json:"aoi_optimization_strategy_enabled"` // 是否开启高价值范围探索
	AoiIDList                      []int64 `json:"aoi_id_list,omitempty"`             // AOI ID列表
}

// AdditionalProductSpec 附加商品属性
type AdditionalProductSpec struct {
	ProductCatalogID string `json:"product_catalog_id"` // 商品库id (必填)
	ProductOuterID   string `json:"product_outer_id"`   // 商品id (必填)
}

// MpaSpec 动态商品广告属性
type MpaSpec struct {
	RecommendMethodIDs []int64 `json:"recommend_method_ids,omitempty"` // 动态创意广告的商品推荐方式
	ProductCatalogID   string  `json:"product_catalog_id"`             // 商品库id (必填)
	ProductSeriesID    string  `json:"product_series_id,omitempty"`    // 商品集合id
}

// DcaSpec 动态内容广告属性
type DcaSpec struct {
	RecommendMethodIDs []int64 `json:"recommend_method_ids"` // 动态内容广告的优选方式
	SetID              string  `json:"set_id"`               // 动态内容广告的素材集合id (必填)
}

// MarketingAssetOuterSpec 产品外部id数据
type MarketingAssetOuterSpec struct {
	MarketingTargetType      string `json:"marketing_target_type"`                  // 推广产品类型 (必填)
	MarketingAssetOuterID    string `json:"marketing_asset_outer_id,omitempty"`     // 推广产品外部id
	MarketingAssetOuterSubID string `json:"marketing_asset_outer_sub_id,omitempty"` // 推广产品外部子id
	MarketingAssetOuterName  string `json:"marketing_asset_outer_name,omitempty"`   // 推广产品外部名称
}

// MarketingTargetAttachment 营销对象附加信息
type MarketingTargetAttachment struct {
	AndroidChannelID string `json:"android_channel_id,omitempty"` // 安卓应用渠道包id
}

// NegativeWordCnt 否定词个数
type NegativeWordCnt struct {
	ExactNegativeWordCnt  int `json:"exact_negative_word_cnt"`  // 精确否定词个数
	PhraseNegativeWordCnt int `json:"phrase_negative_word_cnt"` // 短语否定词个数
}

// AutoDerivedCreativePreference 创意增强MAX偏好设置
type AutoDerivedCreativePreference struct {
	AutoDerivedCreativeMethodTypeList []string `json:"auto_derived_creative_method_type_list,omitempty"` // 创意增强MAX偏好设置列表
}

// MarketingTargetExt 营销对象扩展数据
type MarketingTargetExt struct {
	MarketingTargetName string `json:"marketing_target_name,omitempty"` // 营销对象名称
	CategoryName1       string `json:"category_name1,omitempty"`        // 一级类目名称
	CategoryName2       string `json:"category_name2,omitempty"`        // 二级类目名称
	CategoryName3       string `json:"category_name3,omitempty"`        // 三级类目名称
}

// SceneSpec 场景定向
type SceneSpec struct {
	MobileUnion                 []string     `json:"mobile_union,omitempty"`                   // 移动联盟场景定向
	ExcludeMobileUnion          []string     `json:"exclude_mobile_union,omitempty"`           // 移动联盟场景屏蔽定向
	UnionPositionPackage        []int64      `json:"union_position_package,omitempty"`         // 定投联盟流量包列表
	ExcludeUnionPositionPackage []int64      `json:"exclude_union_position_package,omitempty"` // 屏蔽联盟流量包列表
	TencentNews                 []string     `json:"tencent_news,omitempty"`                   // 腾讯新闻流量场景定向
	DisplayScene                []string     `json:"display_scene,omitempty"`                  // 广告展示场景
	WechatScene                 *WechatScene `json:"wechat_scene,omitempty"`                   // 微信场景定向
	WechatPosition              []int64      `json:"wechat_position,omitempty"`                // 微信公众号与小程序定投
	MobileUnionCategory         []int64      `json:"mobile_union_category,omitempty"`          // 腾讯广告联盟媒体类型场景定向
	QbsearchScene               []string     `json:"qbsearch_scene,omitempty"`                 // QQ浏览器、应用宝流量场景
	WechatChannelsScene         []int64      `json:"wechat_channels_scene,omitempty"`          // 微信视频号定投
	PcScene                     []string     `json:"pc_scene,omitempty"`                       // PC端定投
	WechatSearchScene           []string     `json:"wechat_search_scene,omitempty"`            // 搜一搜流量场景
}

// UserActionSet 用户行为数据源
type UserActionSet struct {
	Type         string `json:"type"`                     // 数据类型 (必填)
	ID           int64  `json:"id"`                       // 数据源id (必填)
	DataSourceID int64  `json:"data_source_id,omitempty"` // DN数据源id
}

// DeepConversionBehaviorSpec oCPA优化转化行为配置
type DeepConversionBehaviorSpec struct {
	Goal      string `json:"goal"`       // 优化转化行为目标 (必填)
	BidAmount int64  `json:"bid_amount"` // 深度优化行为的出价，单位分 (必填)
}

// DeepConversionWorthSpec oCPA优化ROI配置
type DeepConversionWorthSpec struct {
	Goal        string  `json:"goal"`         // 优化ROI目标 (必填)
	ExpectedRoi float64 `json:"expected_roi"` // 深度优化价值效果值 (必填)
}

// DeepConversionWorthAdvanceSpec oCPC/oCPM优化ROI配置
type DeepConversionWorthAdvanceSpec struct {
	Goal        string  `json:"goal"`         // 优化ROI目标 (必填)
	ExpectedRoi float64 `json:"expected_roi"` // 深度优化价值效果值 (必填)
}

// DeepConversionBehaviorAdvancedSpec oCPX深度辅助配置
type DeepConversionBehaviorAdvancedSpec struct {
	Goal      string `json:"goal"`       // 深度辅助优化OG目标 (必填)
	BidAmount int64  `json:"bid_amount"` // 深度辅助优化OG出价，单位分 (必填)
}

// DeepConversionSpec oCPA深度优化内容
type DeepConversionSpec struct {
	DeepConversionType                 string                              `json:"deep_conversion_type"`                             // oCPA深度优化价值配置 (必填)
	DeepConversionBehaviorSpec         *DeepConversionBehaviorSpec         `json:"deep_conversion_behavior_spec,omitempty"`          // oCPA优化转化行为配置
	DeepConversionWorthSpec            *DeepConversionWorthSpec            `json:"deep_conversion_worth_spec,omitempty"`             // oCPA优化ROI配置
	DeepConversionWorthAdvanceSpec     *DeepConversionWorthAdvanceSpec     `json:"deep_conversion_worth_advance_spec,omitempty"`     // oCPC/oCPM优化ROI配置
	DeepConversionBehaviorAdvancedSpec *DeepConversionBehaviorAdvancedSpec `json:"deep_conversion_behavior_advanced_spec,omitempty"` // oCPX深度辅助配置
}

// WechatScene 微信场景定向
type WechatScene struct {
	OfficialAccountMediaCategory []int64 `json:"official_account_media_category,omitempty"` // 公众号媒体类型
	MiniProgramAndMiniGame       []int64 `json:"mini_program_and_mini_game,omitempty"`      // 小程序小游戏流量类型
	PayScene                     []int64 `json:"pay_scene,omitempty"`                       // 订单详情页消费场景
}

// MarketingTargetDetail 营销对象详情
type MarketingTargetDetail struct {
	MarketingTargetDetailID    string `json:"marketing_target_detail_id,omitempty"`     // 推广内容资产详情id
	MarketingTargetSubDetailID string `json:"marketing_target_sub_detail_id,omitempty"` // 二级推广内容资产详情id
}

// MarketingCarrierDetail 营销载体详情
type MarketingCarrierDetail struct {
	MarketingCarrierID    string `json:"marketing_carrier_id,omitempty"`     // 营销载体id
	MarketingSubCarrierID string `json:"marketing_sub_carrier_id,omitempty"` // 二级营销载体id
	MarketingCarrierName  string `json:"marketing_carrier_name,omitempty"`   // 营销载体名称
}

type Targeting struct {
	GeoLocation               *GeoLocation               `json:"geo_location,omitempty"`                // 地理位置定向
	Gender                    []string                   `json:"gender,omitempty"`                      // 性别定向
	Age                       []*AgeRange                `json:"age,omitempty"`                         // 年龄定向
	Education                 []string                   `json:"education,omitempty"`                   // 用户学历（即将下线）
	AppInstallStatus          []string                   `json:"app_install_status,omitempty"`          // 应用安装
	MaritalStatus             []string                   `json:"marital_status,omitempty"`              // 婚恋状态（即将下线）
	ExcludedConvertedAudience *ExcludedConvertedAudience `json:"excluded_converted_audience,omitempty"` // 排除已转化人群
	CustomAudience            []int64                    `json:"custom_audience,omitempty"`             // 定向用户群
	ExcludedCustomAudience    []int64                    `json:"excluded_custom_audience,omitempty"`    // 排除用户群
	DeviceBrandModel          *DeviceBrandModel          `json:"device_brand_model,omitempty"`          // 设备品牌型号定向（即将下线）
	UserOs                    []string                   `json:"user_os,omitempty"`                     // 操作系统定向（即将下线）
	NetworkType               []string                   `json:"network_type,omitempty"`                // 联网方式定向（即将下线）
	DevicePrice               []string                   `json:"device_price,omitempty"`                // 设备价格定向（即将下线）
	WechatAdBehavior          *WechatAdBehavior          `json:"wechat_ad_behavior,omitempty"`          // 微信广告行为定向
	GameConsumptionLevel      []string                   `json:"game_consumption_level,omitempty"`      // 游戏消费能力（即将下线）
	ExcludedOs                []string                   `json:"excluded_os,omitempty"`                 // 排除操作系统定向
}

type WechatAdBehavior struct {
	Actions         []string `json:"actions,omitempty"`          // 微信再营销类型（定向）
	ExcludedActions []string `json:"excluded_actions,omitempty"` // 微信再营销类型（排除）
	CorpID          []string `json:"corp_id,omitempty"`          // 企业微信id列表
}

// DeviceBrandModel 设备品牌型号定向
type DeviceBrandModel struct {
	IncludedList []int64 `json:"included_list,omitempty"` // 设备品牌型号定向列表
	ExcludedList []int64 `json:"excluded_list,omitempty"` // 排除设备品牌型号列表
}

// ExcludedConvertedAudience 排除已转化人群配置
type ExcludedConvertedAudience struct {
	ExcludedDimension      string   `json:"excluded_dimension,omitempty"`       // 排除已转化人群的数据维度
	ConversionBehaviorList []string `json:"conversion_behavior_list,omitempty"` // 转化行为
	ExcludedDay            string   `json:"excluded_day,omitempty"`             // 排除天数
}

// AgeRange 年龄范围
type AgeRange struct {
	Min int `json:"min"` // 年龄下限
	Max int `json:"max"` // 年龄上限
}

type GeoLocation struct {
	LocationTypes           []string          `json:"location_types,omitempty"`             // 地点类型
	Regions                 []int64           `json:"regions,omitempty"`                    // 省市区县列表
	BusinessDistricts       []int64           `json:"business_districts,omitempty"`         // 商圈id列表
	CustomLocations         []*CustomLocation `json:"custom_locations,omitempty"`           // 自定义地理位置列表
	GeoLocationAutoAudience bool              `json:"geo_location_auto_audience,omitempty"` // 是否使用地域优选
}

// CustomLocation 自定义地理位置
type CustomLocation struct {
	Longitude float64 `json:"longitude"` // 经度，单位度
	Latitude  float64 `json:"latitude"`  // 纬度，单位度
	Radius    int     `json:"radius"`    // 半径，单位米
}
