package model

import (
	"errors"
	"time"
)

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

// AdgroupCommonStruct 广告组通用结构体
type AdgroupCommonStruct struct {
	Targeting                         *Targeting                     `json:"targeting,omitempty"`                             // 定向详细设置
	ConfiguredStatus                  string                         `json:"configured_status"`                               // 客户设置的状态
	AdgroupName                       string                         `json:"adgroup_name"`                                    // 广告名称
	MarketingGoal                     string                         `json:"marketing_goal"`                                  // 营销目的类型
	MarketingSubGoal                  string                         `json:"marketing_sub_goal,omitempty"`                    // 二级营销目的类型
	MarketingCarrierType              string                         `json:"marketing_carrier_type"`                          // 营销载体类型
	MarketingCarrierDetail            *MarketingCarrierDetail        `json:"marketing_carrier_detail,omitempty"`              // 营销载体详情
	MarketingTargetType               string                         `json:"marketing_target_type"`                           // 推广产品类型
	BeginDate                         string                         `json:"begin_date"`                                      // 开始投放日期 (必填)
	EndDate                           string                         `json:"end_date"`                                        // 结束投放日期 (必填)
	FirstDayBeginTime                 string                         `json:"first_day_begin_time,omitempty"`                  // 首日开始投放时间
	BidAmount                         int64                          `json:"bid_amount"`                                      // 广告出价，单位分 (必填)
	OptimizationGoal                  string                         `json:"optimization_goal"`                               // 广告优化目标类型 (必填)
	TimeSeries                        string                         `json:"time_series"`                                     // 投放时间段 (必填)
	AutomaticSiteEnabled              bool                           `json:"automatic_site_enabled"`                          // 是否开启智能版位功能
	SiteSet                           []string                       `json:"site_set,omitempty"`                              // 投放站点集合
	DailyBudget                       int64                          `json:"daily_budget"`                                    // 日预算，单位分
	SceneSpec                         *SceneSpec                     `json:"scene_spec,omitempty"`                            // 场景定向
	UserActionSets                    []*UserActionSet               `json:"user_action_sets,omitempty"`                      // 用户行为数据源
	DeepConversionSpec                *DeepConversionSpec            `json:"deep_conversion_spec,omitempty"`                  // oCPA 深度优化内容
	ConversionID                      int64                          `json:"conversion_id,omitempty"`                         // 转化id
	DeepConversionBehaviorBid         int64                          `json:"deep_conversion_behavior_bid,omitempty"`          // 深度优化行为出价，单位分
	DeepConversionWorthRate           float64                        `json:"deep_conversion_worth_rate,omitempty"`            // 深度优化价值出价
	DeepConversionWorthAdvancedRate   float64                        `json:"deep_conversion_worth_advanced_rate,omitempty"`   // 强化优化价值的期望ROI
	DeepConversionBehaviorAdvancedBid int64                          `json:"deep_conversion_behavior_advanced_bid,omitempty"` // 深度辅助优化OG出价，单位分
	BidMode                           string                         `json:"bid_mode"`                                        // 出价方式 (必填)
	AutoAcquisitionEnabled            bool                           `json:"auto_acquisition_enabled,omitempty"`              // 一键起量开关
	AutoAcquisitionBudget             int64                          `json:"auto_acquisition_budget,omitempty"`               // 一键起量预算，单位分
	SmartBidType                      string                         `json:"smart_bid_type,omitempty"`                        // 出价类型
	SmartCostCap                      int64                          `json:"smart_cost_cap,omitempty"`                        // 自动出价下预计成本上限，单位分
	AutoDerivedCreativeEnabled        bool                           `json:"auto_derived_creative_enabled,omitempty"`         // 创意增强MAX开关
	AutoDerivedCreativePreference     *AutoDerivedCreativePreference `json:"auto_derived_creative_preference,omitempty"`      // 创意增强MAX偏好设置
	SearchExpandTargetingSwitch       string                         `json:"search_expand_targeting_switch,omitempty"`        // 搜索定向拓展开关
	AutoDerivedLandingPageSwitch      bool                           `json:"auto_derived_landing_page_switch,omitempty"`      // 是否开启自动衍生落地页开关
	BidScene                          string                         `json:"bid_scene,omitempty"`                             // 出价场景
	FlowOptimizationEnabled           bool                           `json:"flow_optimization_enabled,omitempty"`             // 是否使用自动流量优选（已废弃）
	SearchExpansionSwitch             string                         `json:"search_expansion_switch,omitempty"`               // 搜索引擎开关
	MarketingAssetID                  int64                          `json:"marketing_asset_id,omitempty"`                    // 产品id
	MaterialPackageID                 int64                          `json:"material_package_id,omitempty"`                   // 素材标签id
	MarketingAssetOuterSpec           *MarketingAssetOuterSpec       `json:"marketing_asset_outer_spec,omitempty"`            // 产品外部id数据
	PoiList                           []string                       `json:"poi_list,omitempty"`                              // 门店id列表
	ExplorationStrategy               string                         `json:"exploration_strategy,omitempty"`                  // 探索策略
	PrioritySiteSet                   []string                       `json:"priority_site_set,omitempty"`                     // 投放站点集合
	EcomPkamSwitch                    string                         `json:"ecom_pkam_switch,omitempty"`                      // 一方人群跑量加强开关状态
	ForwardLinkAssist                 string                         `json:"forward_link_assist,omitempty"`                   // 助攻行为目标
	AutoAcquisitionStatus             string                         `json:"auto_acquisition_status,omitempty"`               // 一键起量状态
	CustomCostCap                     int64                          `json:"custom_cost_cap,omitempty"`                       // 用户输入的成本上限，单位分
	MpaSpec                           *MpaSpec                       `json:"mpa_spec,omitempty"`                              // 动态商品广告属性
	ShortPlayPayType                  string                         `json:"short_play_pay_type,omitempty"`                   // 售卖方式类型
	SellStrategyID                    int64                          `json:"sell_strategy_id,omitempty"`                      // 售卖策略id
	DcaSpec                           *DcaSpec                       `json:"dca_spec,omitempty"`                              // 动态内容广告属性
	AoiOptimizationStrategy           *AoiOptimizationStrategy       `json:"aoi_optimization_strategy,omitempty"`             // 高价值范围探索
	AdditionalProductSpec             *AdditionalProductSpec         `json:"additional_product_spec,omitempty"`               // 附加商品属性
	LiveRecommendStrategyEnabled      bool                           `json:"live_recommend_strategy_enabled,omitempty"`       // 直播种草人群探索
	EnableSteadyExploration           bool                           `json:"enable_steady_exploration,omitempty"`             // 是否稳步探索更多版位
	AdxRealtimeType                   string                         `json:"adx_realtime_type,omitempty"`                     // ADX程序化广告素材实时回复类型
}

type AdgroupsGetListItem struct {
	AdgroupCommonStruct
	AdgroupID                 int64                      `json:"adgroup_id"`                            // 广告id
	TargetingTranslation      string                     `json:"targeting_translation,omitempty"`       // 已选择定向条件的描述
	CreatedTime               int64                      `json:"created_time"`                          // 创建时间，时间戳
	LastModifiedTime          int64                      `json:"last_modified_time"`                    // 最后修改时间，时间戳
	IsDeleted                 bool                       `json:"is_deleted"`                            // 是否已删除
	SystemStatus              string                     `json:"system_status"`                         // 广告在系统中的状态
	MarketingTargetDetail     *MarketingTargetDetail     `json:"marketing_target_detail,omitempty"`     // 营销对象详情
	MarketingTargetID         int64                      `json:"marketing_target_id,omitempty"`         // 营销对象id
	DataModelVersion          int64                      `json:"data_model_version,omitempty"`          // 数据版本号
	MarketingTargetExt        *MarketingTargetExt        `json:"marketing_target_ext,omitempty"`        // 营销对象扩展数据
	DeepOptimizationType      string                     `json:"deep_optimization_type,omitempty"`      // 深度优化策略类型
	MarketingTargetAttachment *MarketingTargetAttachment `json:"marketing_target_attachment,omitempty"` // 营销对象附加信息
	NegativeWordCnt           *NegativeWordCnt           `json:"negative_word_cnt,omitempty"`           // 否定词个数
	PromotedAssetType         string                     `json:"promoted_asset_type,omitempty"`         // 推广内容类型
	MarketingScene            string                     `json:"marketing_scene,omitempty"`             // 营销目标
	ConversionName            string                     `json:"conversion_name,omitempty"`             // 转化名称
	AutoAcquisitionStatus     string                     `json:"auto_acquisition_status,omitempty"`     // 一键起量状态
	OgCompletionType          string                     `json:"og_completion_type,omitempty"`          // 达成类型
	CostGuaranteeStatus       string                     `json:"cost_guarantee_status,omitempty"`       // 成本保障状态
	CostGuaranteeMoney        int64                      `json:"cost_guarantee_money,omitempty"`        // 成本保障赔付金额，单位分
	EnableBreakthroughSiteset bool                       `json:"enable_breakthrough_siteset,omitempty"` // 是否支持版位突破
	CustomCostRolCap          float64                    `json:"custom_cost_rol_cap,omitempty"`         // 控制成本的期望ROI
	SmartTargetingStatus      string                     `json:"smart_targeting_status,omitempty"`      // 广告智能定向状态
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

// 常量定义 - 营销目的类型
const (
	MarketingGoalUnknown                 = "MARKETING_GOAL_UNKNOWN"                   // 未知
	MarketingGoalUserGrowth              = "MARKETING_GOAL_USER_GROWTH"               // 用户增长
	MarketingGoalProductSales            = "MARKETING_GOAL_PRODUCT_SALES"             // 商品销售
	MarketingGoalLeadRetention           = "MARKETING_GOAL_LEAD_RETENTION"            // 销售线索收集
	MarketingGoalBrandPromotion          = "MARKETING_GOAL_BRAND_PROMOTION"           // 品牌推广
	MarketingGoalIncreaseFansInteraction = "MARKETING_GOAL_INCREASE_FANS_INTERACTION" // 提升粉丝互动
)

// 常量定义 - 二级营销目的类型
const (
	MarketingSubGoalUnknown                          = "MARKETING_SUB_GOAL_UNKNOWN"                              // 未知
	MarketingSubGoalNewGameReserve                   = "MARKETING_SUB_GOAL_NEW_GAME_RESERVE"                     // 新游戏预约
	MarketingSubGoalNewGameTest                      = "MARKETING_SUB_GOAL_NEW_GAME_TEST"                        // 新游戏测试
	MarketingSubGoalNewGameLaunch                    = "MARKETING_SUB_GOAL_NEW_GAME_LAUNCH"                      // 新游戏首发
	MarketingSubGoalPlateauPhaseLaunch               = "MARKETING_SUB_GOAL_PLATEAU_PHASE_LAUNCH"                 // 平推期投放
	MarketingSubGoalMinigameNewCustomerGrowth        = "MARKETING_SUB_GOAL_MINIGAME_NEW_CUSTOMER_GROWTH"         // 小游戏新客增长
	MarketingSubGoalMiniGameReturnCustomerEngagement = "MARKETING_SUB_GOAL_MINI_GAME_RETURN_CUSTOMER_ENGAGEMENT" // 小游戏老客召回
	MarketingSubGoalAppAcquisition                   = "MARKETING_SUB_GOAL_APP_ACQUISITION"                      // 应用获取
	MarketingSubGoalAppActivation                    = "MARKETING_SUB_GOAL_APP_ACTIVATION"                       // 应用激活
	MarketingSubGoalNotInstallUser                   = "MARKETING_SUB_GOAL_NOT_INSTALL_USER"                     // 未安装用户
	MarketingSubGoalPreInstallUser                   = "MARKETING_SUB_GOAL_PRE_INSTALL_USER"                     // 预安装用户
	MarketingSubGoalUnloadedUser                     = "MARKETING_SUB_GOAL_UNLOADED_USER"                        // 已卸载用户
	MarketingSubGoalShortInactiveUser                = "MARKETING_SUB_GOAL_SHORT_INACTIVE_USER"                  // 短期不活跃用户
	MarketingSubGoalLongInactiveUser                 = "MARKETING_SUB_GOAL_LONG_INACTIVE_USER"                   // 长期不活跃用户
	MarketingSubGoalGameVersionUpgrade               = "MARKETING_SUB_GOAL_GAME_VERSION_UPGRADE"                 // 游戏版本升级
	MarketingSubGoalNewStoreOpening                  = "MARKETING_SUB_GOAL_NEW_STORE_OPENING"                    // 新店开业
	MarketingSubGoalEveningPromotion                 = "MARKETING_SUB_GOAL_EVENING_PROMOTION"                    // 晚间促销
	MarketingSubGoalSpecialRelease                   = "MARKETING_SUB_GOAL_SPECIAL_RELEASE"                      // 特殊发布
)

// 创建广告
type AdgroupsAddReq struct {
	GlobalReq
	AdgroupCommonStruct
	AccountID              int64                   `json:"account_id"`                         // 广告主帐号id (必填)
	AdgroupName            string                  `json:"adgroup_name"`                       // 广告名称 (必填)
	MarketingGoal          string                  `json:"marketing_goal"`                     // 营销目的类型 (必填)
	MarketingSubGoal       string                  `json:"marketing_sub_goal"`                 // 二级营销目的类型
	MarketingCarrierType   string                  `json:"marketing_carrier_type"`             // 营销载体类型 (必填)
	MarketingCarrierDetail *MarketingCarrierDetail `json:"marketing_carrier_detail,omitempty"` // 营销载体详情
	BeginDate              string                  `json:"begin_date"`                         // 开始投放日期 (必填)
	EndDate                string                  `json:"end_date"`                           // 结束投放日期 (必填)
	FirstDayBeginTime      string                  `json:"first_day_begin_time,omitempty"`     // 首日开始投放时间
	BidAmount              int64                   `json:"bid_amount"`                         // 广告出价，单位分 (必填)
	OptimizationGoal       string                  `json:"optimization_goal,omitempty"`        // 广告优化目标类型
	TimeSeries             string                  `json:"time_series"`                        // 投放时间段 (必填)
	AutomaticSiteEnabled   bool                    `json:"automatic_site_enabled"`             // 是否开启智能版位功能
	SiteSet                []string                `json:"site_set,omitempty"`                 // 投放版位集合
	ExplorationStrategy    string                  `json:"exploration_strategy,omitempty"`     // 自动版位探索策略
	PrioritySiteSet        []string                `json:"priority_site_set,omitempty"`        // 优先级位集合
	DailyBudget            int64                   `json:"daily_budget"`                       // 日预算，单位分
	SmartTargetingMode     string                  `json:"smart_targeting_mode,omitempty"`     // 广告智能定向功能
	SmartCouponMode        string                  `json:"smart_coupon_mode,omitempty"`        // 小店智券开关
}

func (p *AdgroupsAddReq) Format() {
	p.GlobalReq.Format()
}

// 常量定义 - 广告智能定向功能
const (
	SmartTargetingManual = "SMART_TARGETING_MANUAL" // 手动定向
	// 其他值根据文档补充
)

// 常量定义 - 小店智券开关
const (
	SwitchStatusOff = "SWITCH_STATUS_OFF" // 关闭（默认）
	SwitchStatusOn  = "SWITCH_STATUS_ON"  // 开启
)

// 常量定义 - 投放时间段
const (
	TimeSeriesLength = 336 // 48*7，以半小时为粒度
)

// 常量定义 - 站点集合长度限制
const (
	MinSiteSetCount = 1
	MaxSiteSetCount = 32
)

// 常量定义 - 日预算限制
const (
	DailyBudgetMin = 5000      // 最小日预算 50元
	DailyBudgetMax = 400000000 // 最大日预算 4,000,000元
)

// 常量定义 - 自动版位探索策略
const (
	ExplorationStrategySteady = "STEADY_EXPLORATION" // 稳步探索
	ExplorationStrategyFast   = "FAST_EXPLORATION"   // 快速探索
)

// 常量定义 - 投放版位
const (
	SiteSetMobileUnion       = "SITE_SET_MOBILE_UNION"        // 移动联盟
	SiteSetWechat            = "SITE_SET_WECHAT"              // 微信
	SiteSetTencentNews       = "SITE_SET_TENCENT_NEWS"        // 腾讯新闻
	SiteSetTencentVideo      = "SITE_SET_TENCENT_VIDEO"       // 腾讯视频
	SiteSetMobileYyb         = "SITE_SET_MOBILE_YYB"          // 应用宝
	SiteSetPcqq              = "SITE_SET_PCQQ"                // PCQQ
	SiteSetKandian           = "SITE_SET_KANDIAN"             // 看点
	SiteSetQqMusicGame       = "SITE_SET_QQ_MUSIC_GAME"       // QQ音乐游戏
	SiteSetMoments           = "SITE_SET_MOMENTS"             // 朋友圈
	SiteSetChannels          = "SITE_SET_CHANNELS"            // 视频号
	SiteSetWechatSearch      = "SITE_SET_WECHAT_SEARCH"       // 微信搜一搜
	SiteSetWechatPlugin      = "SITE_SET_WECHAT_PLUGIN"       // 微信插件
	SiteSetQbsearch          = "SITE_SET_QBSEARCH"            // QQ浏览器搜索
	SiteSetSearchScene       = "SITE_SET_SEARCH_SCENE"        // 搜索场景
	SiteSetSearchMobileUnion = "SITE_SET_SEARCH_MOBILE_UNION" // 搜索移动联盟
	SiteSetSmart             = "SITE_SET_SMART"               // 智能版位
)

// 常量定义 - 营销载体类型
const (
	MarketingCarrierTypeUnknown                       = "MARKETING_CARRIER_TYPE_UNKNOWN"                          // 未知
	MarketingCarrierTypeAppAndroid                    = "MARKETING_CARRIER_TYPE_APP_ANDROID"                      // 安卓应用
	MarketingCarrierTypeAppIOS                        = "MARKETING_CARRIER_TYPE_APP_IOS"                          // iOS应用
	MarketingCarrierTypeWechatOfficialAccount         = "MARKETING_CARRIER_TYPE_WECHAT_OFFICIAL_ACCOUNT"          // 微信公众号
	MarketingCarrierTypeJumpPage                      = "MARKETING_CARRIER_TYPE_JUMP_PAGE"                        // 跳转页面
	MarketingCarrierTypeWechatMiniGame                = "MARKETING_CARRIER_TYPE_WECHAT_MINIGAME"                  // 微信小游戏
	MarketingCarrierTypeWechatChannelsLive            = "MARKETING_CARRIER_TYPE_WECHAT_CHANNELS_LIVE"             // 微信视频号直播
	MarketingCarrierTypeWechatChannels                = "MARKETING_CARRIER_TYPE_WECHAT_CHANNELS"                  // 微信视频号
	MarketingCarrierTypeWechatChannelsLiveReservation = "MARKETING_CARRIER_TYPE_WECHAT_CHANNELS_LIVE_RESERVATION" // 微信视频号直播预约
	MarketingCarrierTypeMiniProgramWechat             = "MARKETING_CARRIER_TYPE_MINI_PROGRAM_WECHAT"              // 微信小程序
	MarketingCarrierTypeAppQuickApp                   = "MARKETING_CARRIER_TYPE_APP_QUICK_APP"                    // 快应用
	MarketingCarrierTypePCGame                        = "MARKETING_CARRIER_TYPE_PC_GAME"                          // PC游戏
	MarketingCarrierTypeQQMiniGame                    = "MARKETING_CARRIER_TYPE_QQ_MINIGAME"                      // QQ小游戏
	MarketingCarrierTypeAppHarmony                    = "MARKETING_CARRIER_TYPE_APP_HARMONY"                      // 鸿蒙应用
)

// 常量定义 - 首日开始投放时间格式
const TimeFormat = "15:04:05"

// 常量定义 - 首日开始投放时间长度限制
const (
	MinFirstDayBeginTimeLength = 0
	MaxFirstDayBeginTimeLength = 8
)

// 长度限制常量
const (
	MinAdgroupNameLength = 1
	MaxAdgroupNameLength = 60 // 等宽字符，即60个中文字或120个英文字
)

// 日期格式常量
const DateFormat = "2006-01-02"

// Validate 验证广告组创建请求
func (p *AdgroupsAddReq) Validate() error {
	// 1. 验证account_id
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	// 2. 验证广告名称
	if p.AdgroupName == "" {
		return errors.New("adgroup_name为必填")
	}
	if len(p.AdgroupName) > MaxAdgroupNameLength*3 {
		return errors.New("adgroup_name长度不能超过60个等宽字符")
	}

	// 3. 验证营销目的类型
	if p.MarketingGoal == "" {
		return errors.New("marketing_goal为必填")
	}
	if !isValidMarketingGoal(p.MarketingGoal) {
		return errors.New("marketing_goal值无效，请参考文档中的枚举值")
	}

	// 4. 验证二级营销目的类型
	if p.MarketingSubGoal != "" && !isValidMarketingSubGoal(p.MarketingSubGoal) {
		return errors.New("marketing_sub_goal值无效，请参考文档中的枚举值")
	}

	// 5. 验证营销载体类型
	if p.MarketingCarrierType == "" {
		return errors.New("marketing_carrier_type为必填")
	}
	if !isValidMarketingCarrierType(p.MarketingCarrierType) {
		return errors.New("marketing_carrier_type值无效，请参考文档中的枚举值")
	}

	// 6. 验证营销载体详情
	if p.MarketingCarrierDetail != nil {
		if err := p.MarketingCarrierDetail.Validate(); err != nil {
			return err
		}
	}

	// 7. 验证开始和结束日期
	if err := p.validateDates(); err != nil {
		return err
	}

	// 8. 验证首日开始投放时间
	if err := p.validateFirstDayBeginTime(); err != nil {
		return err
	}

	// 9. 验证广告出价
	if err := p.validateBidAmount(); err != nil {
		return err
	}

	// 10. 验证投放时间段
	if err := p.validateTimeSeries(); err != nil {
		return err
	}

	// 11. 验证站点集合
	if err := p.validateSiteSet(); err != nil {
		return err
	}

	// 12. 验证探索策略
	if err := p.validateExplorationStrategy(); err != nil {
		return err
	}

	// 13. 验证日预算
	if err := p.validateDailyBudget(); err != nil {
		return err
	}

	// 14. 验证智能定向模式
	if err := p.validateSmartTargetingMode(); err != nil {
		return err
	}

	// 15. 验证小店智券开关
	if err := p.validateSmartCouponMode(); err != nil {
		return err
	}

	return nil
}

// validateSmartTargetingMode 验证广告智能定向功能
func (p *AdgroupsAddReq) validateSmartTargetingMode() error {
	if p.SmartTargetingMode == "" {
		return nil
	}

	validModes := map[string]bool{
		SmartTargetingManual: true,
		// 可根据文档补充其他模式
	}

	if !validModes[p.SmartTargetingMode] {
		return errors.New("smart_targeting_mode值无效，允许值：SMART_TARGETING_MANUAL")
	}

	return nil
}

// validateSmartCouponMode 验证小店智券开关
func (p *AdgroupsAddReq) validateSmartCouponMode() error {
	if p.SmartCouponMode == "" {
		p.SmartCouponMode = SwitchStatusOff // 默认关闭
		return nil
	}

	if p.SmartCouponMode != SwitchStatusOff && p.SmartCouponMode != SwitchStatusOn {
		return errors.New("smart_coupon_mode值无效，允许值：SWITCH_STATUS_OFF、SWITCH_STATUS_ON")
	}

	return nil
}

// validateTimeSeries 验证投放时间段
func (p *AdgroupsAddReq) validateTimeSeries() error {
	if p.TimeSeries == "" {
		return errors.New("time_series为必填")
	}
	if len(p.TimeSeries) != TimeSeriesLength {
		return errors.New("time_series长度必须为336")
	}
	// 检查是否只包含0和1
	for _, c := range p.TimeSeries {
		if c != '0' && c != '1' {
			return errors.New("time_series只能包含0和1")
		}
	}
	// 不允许全部传0
	allZero := true
	for _, c := range p.TimeSeries {
		if c == '1' {
			allZero = false
			break
		}
	}
	if allZero {
		return errors.New("time_series不允许全部为0")
	}
	return nil
}

// validateSiteSet 验证站点集合
func (p *AdgroupsAddReq) validateSiteSet() error {
	// 使用智能版位时无需传site_set
	if p.AutomaticSiteEnabled {
		if len(p.SiteSet) > 0 {
			return errors.New("使用智能版位时，site_set字段无需传值")
		}
		return nil
	}

	// 非智能版位时，site_set为必填
	if len(p.SiteSet) == 0 {
		return errors.New("未开启智能版位时，site_set为必填")
	}
	if len(p.SiteSet) < MinSiteSetCount || len(p.SiteSet) > MaxSiteSetCount {
		return errors.New("site_set数组长度必须在1-32之间")
	}

	// 验证站点集合值是否有效
	validSiteSets := map[string]bool{
		SiteSetMobileUnion:       true,
		SiteSetWechat:            true,
		SiteSetTencentNews:       true,
		SiteSetTencentVideo:      true,
		SiteSetMobileYyb:         true,
		SiteSetPcqq:              true,
		SiteSetKandian:           true,
		SiteSetQqMusicGame:       true,
		SiteSetMoments:           true,
		SiteSetChannels:          true,
		SiteSetWechatSearch:      true,
		SiteSetWechatPlugin:      true,
		SiteSetQbsearch:          true,
		SiteSetSearchScene:       true,
		SiteSetSearchMobileUnion: true,
		SiteSetSmart:             true,
	}
	for _, site := range p.SiteSet {
		if !validSiteSets[site] {
			return errors.New("site_set包含无效值，请参考文档中的枚举值")
		}
	}
	return nil
}

// validateExplorationStrategy 验证自动版位探索策略
func (p *AdgroupsAddReq) validateExplorationStrategy() error {
	if p.ExplorationStrategy == "" {
		return nil
	}
	if p.ExplorationStrategy != ExplorationStrategySteady && p.ExplorationStrategy != ExplorationStrategyFast {
		return errors.New("exploration_strategy值无效，允许值：STEADY_EXPLORATION、FAST_EXPLORATION")
	}

	// 稳步探索策略需要设置优先级位集合
	if p.ExplorationStrategy == ExplorationStrategySteady && len(p.PrioritySiteSet) == 0 {
		return errors.New("exploration_strategy为STEADY_EXPLORATION时，priority_site_set为必填")
	}

	// 验证优先级位集合
	if len(p.PrioritySiteSet) > MaxSiteSetCount {
		return errors.New("priority_site_set数组长度不能超过32")
	}
	return nil
}

// validateDailyBudget 验证日预算
func (p *AdgroupsAddReq) validateDailyBudget() error {
	if p.DailyBudget == 0 {
		return nil // 0表示不设预算
	}
	if p.DailyBudget < DailyBudgetMin {
		return errors.New("daily_budget不能小于5000分（50元）")
	}
	if p.DailyBudget > DailyBudgetMax {
		return errors.New("daily_budget不能大于400000000分（4,000,000元）")
	}
	return nil
}

// validateFirstDayBeginTime 验证首日开始投放时间
func (p *AdgroupsAddReq) validateFirstDayBeginTime() error {
	if p.FirstDayBeginTime == "" {
		return nil
	}

	// 验证长度
	if len(p.FirstDayBeginTime) < MinFirstDayBeginTimeLength || len(p.FirstDayBeginTime) > MaxFirstDayBeginTimeLength {
		return errors.New("first_day_begin_time长度必须在0-8字节之间")
	}

	// 验证时间格式 HH:ii:ss
	_, err := time.Parse(TimeFormat, p.FirstDayBeginTime)
	if err != nil {
		return errors.New("first_day_begin_time格式错误，应为HH:ii:ss")
	}

	return nil
}

// validateBidAmount 验证广告出价
func (p *AdgroupsAddReq) validateBidAmount() error {
	if p.BidAmount <= 0 {
		return errors.New("bid_amount必须大于0")
	}
	return nil
}

// Validate 验证营销载体详情
func (m *MarketingCarrierDetail) Validate() error {
	// 营销载体id长度验证
	if len(m.MarketingCarrierID) > 2048 {
		return errors.New("marketing_carrier_id长度不能超过2048字节")
	}
	return nil
}

// validateDates 验证日期
func (p *AdgroupsAddReq) validateDates() error {
	// 验证开始日期
	if p.BeginDate == "" {
		return errors.New("begin_date为必填")
	}
	if len(p.BeginDate) != 10 {
		return errors.New("begin_date格式错误，应为YYYY-MM-DD")
	}
	begin, err := time.Parse(DateFormat, p.BeginDate)
	if err != nil {
		return errors.New("begin_date格式错误，应为YYYY-MM-DD")
	}

	// 验证结束日期
	if p.EndDate == "" {
		return errors.New("end_date为必填")
	}
	if len(p.EndDate) > 10 {
		return errors.New("end_date长度不能超过10字节")
	}
	end, err := time.Parse(DateFormat, p.EndDate)
	if err != nil {
		return errors.New("end_date格式错误，应为YYYY-MM-DD")
	}

	// 开始日期 <= 结束日期
	if begin.After(end) {
		return errors.New("begin_date不能大于end_date")
	}

	// 结束日期 >= 今天
	today := time.Now().Truncate(24 * time.Hour)
	if end.Before(today) {
		return errors.New("end_date不能小于今天")
	}

	return nil
}

// isValidMarketingCarrierType 验证营销载体类型是否有效
func isValidMarketingCarrierType(carrierType string) bool {
	validTypes := map[string]bool{
		MarketingCarrierTypeUnknown:                       true,
		MarketingCarrierTypeAppAndroid:                    true,
		MarketingCarrierTypeAppIOS:                        true,
		MarketingCarrierTypeWechatOfficialAccount:         true,
		MarketingCarrierTypeJumpPage:                      true,
		MarketingCarrierTypeWechatMiniGame:                true,
		MarketingCarrierTypeWechatChannelsLive:            true,
		MarketingCarrierTypeWechatChannels:                true,
		MarketingCarrierTypeWechatChannelsLiveReservation: true,
		MarketingCarrierTypeMiniProgramWechat:             true,
		MarketingCarrierTypeAppQuickApp:                   true,
		MarketingCarrierTypePCGame:                        true,
		MarketingCarrierTypeQQMiniGame:                    true,
		MarketingCarrierTypeAppHarmony:                    true,
	}
	return validTypes[carrierType]
}

// isValidMarketingGoal 验证营销目的类型是否有效
func isValidMarketingGoal(goal string) bool {
	validGoals := map[string]bool{
		MarketingGoalUnknown:                 true,
		MarketingGoalUserGrowth:              true,
		MarketingGoalProductSales:            true,
		MarketingGoalLeadRetention:           true,
		MarketingGoalBrandPromotion:          true,
		MarketingGoalIncreaseFansInteraction: true,
	}
	return validGoals[goal]
}

// isValidMarketingSubGoal 验证二级营销目的类型是否有效
func isValidMarketingSubGoal(subGoal string) bool {
	validSubGoals := map[string]bool{
		MarketingSubGoalUnknown:                          true,
		MarketingSubGoalNewGameReserve:                   true,
		MarketingSubGoalNewGameTest:                      true,
		MarketingSubGoalNewGameLaunch:                    true,
		MarketingSubGoalPlateauPhaseLaunch:               true,
		MarketingSubGoalMinigameNewCustomerGrowth:        true,
		MarketingSubGoalMiniGameReturnCustomerEngagement: true,
		MarketingSubGoalAppAcquisition:                   true,
		MarketingSubGoalAppActivation:                    true,
		MarketingSubGoalNotInstallUser:                   true,
		MarketingSubGoalPreInstallUser:                   true,
		MarketingSubGoalUnloadedUser:                     true,
		MarketingSubGoalShortInactiveUser:                true,
		MarketingSubGoalLongInactiveUser:                 true,
		MarketingSubGoalGameVersionUpgrade:               true,
		MarketingSubGoalNewStoreOpening:                  true,
		MarketingSubGoalEveningPromotion:                 true,
		MarketingSubGoalSpecialRelease:                   true,
	}
	return validSubGoals[subGoal]
}

type AdgroupsAddResp struct {
	AdgroupId int64 `json:"adgroup_id,omitempty"` // 广告 id
}

type AdgroupsDeleteReq struct {
	GlobalReq
	AccountID int64 `json:"account_id"` // 广告主账号 id (必填)
	AdgroupID int64 `json:"adgroup_id"` // 广告 id (必填)
}

func (p *AdgroupsDeleteReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证广告组更新请求
func (p *AdgroupsDeleteReq) Validate() error {
	// 1. 验证account_id
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	// 2. 验证adgroup_id
	if p.AdgroupID == 0 {
		return errors.New("adgroup_id为必填")
	}

	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	return nil
}

type AdgroupsDeleteResp struct {
	AdgroupId int64 `json:"adgroup_id,omitempty"` // 广告 id
}

// AdgroupsUpdateReq 广告组更新请求
type AdgroupsUpdateReq struct {
	GlobalReq
	AccountID                         int64                          `json:"account_id"`                     // 广告主帐号id (必填)
	AdgroupID                         int64                          `json:"adgroup_id"`                     // 广告id (必填)
	AdgroupName                       string                         `json:"adgroup_name,omitempty"`         // 广告名称
	BeginDate                         string                         `json:"begin_date,omitempty"`           // 开始投放日期
	EndDate                           string                         `json:"end_date,omitempty"`             // 结束投放日期
	FirstDayBeginTime                 string                         `json:"first_day_begin_time,omitempty"` // 首日开始投放时间
	BidAmount                         int64                          `json:"bid_amount,omitempty"`           // 广告出价，单位分
	OptimizationGoal                  string                         `json:"optimization_goal,omitempty"`    // 广告优化目标类型
	TimeSeries                        string                         `json:"time_series"`                    // 投放时间段 (必填)
	DailyBudget                       int64                          `json:"daily_budget"`                   // 日预算，单位分
	Targeting                         *Targeting                     `json:"targeting,omitempty"`
	SceneSpec                         *SceneSpec                     `json:"scene_spec,omitempty"`                            // 场景定向
	UserActionSets                    []*UserActionSet               `json:"user_action_sets,omitempty"`                      // 用户行为数据源
	DeepConversionSpec                *DeepConversionSpec            `json:"deep_conversion_spec,omitempty"`                  // oCPA 深度优化内容
	ConversionID                      int64                          `json:"conversion_id,omitempty"`                         // 转化id
	DeepConversionBehaviorBid         int64                          `json:"deep_conversion_behavior_bid,omitempty"`          // 深度优化行为出价，单位分
	DeepConversionWorthRate           float64                        `json:"deep_conversion_worth_rate,omitempty"`            // 深度优化价值出价
	DeepConversionWorthAdvancedRate   float64                        `json:"deep_conversion_worth_advanced_rate,omitempty"`   // 强化优化价值的期望ROI
	DeepConversionBehaviorAdvancedBid int64                          `json:"deep_conversion_behavior_advanced_bid,omitempty"` // 深度辅助优化OG出价，单位分
	BidMode                           string                         `json:"bid_mode"`                                        // 出价方式 (必填)
	AutoAcquisitionEnabled            bool                           `json:"auto_acquisition_enabled,omitempty"`              // 一键起量开关
	AutoAcquisitionBudget             int64                          `json:"auto_acquisition_budget,omitempty"`               // 一键起量预算，单位分
	SmartBidType                      string                         `json:"smart_bid_type,omitempty"`                        // 出价类型
	SmartCostCap                      int64                          `json:"smart_cost_cap,omitempty"`                        // 自动出价下预计成本上限，单位分
	AutoDerivedCreativeEnabled        bool                           `json:"auto_derived_creative_enabled,omitempty"`         // 创意增强MAX开关
	AutoDerivedCreativePreference     *AutoDerivedCreativePreference `json:"auto_derived_creative_preference,omitempty"`
	ConfiguredStatus                  string                         `json:"configured_status"`
	FlowOptimizationEnabled           bool                           `json:"flow_optimization_enabled,omitempty"`
	PoiList                           []string                       `json:"poi_list,omitempty"`
	EcomPkamSwitch                    string                         `json:"ecom_pkam_switch,omitempty"`
	RtaId                             int64                          `json:"rta_id,omitempty"`                          // RTA 客户 id
	RtaTargetId                       string                         `json:"rta_target_id,omitempty"`                   // RTA 策略 id
	CostConstraintScene               string                         `json:"cost_constraint_scene,omitempty"`           // 成本控制场景
	CustomCostCap                     int64                          `json:"custom_cost_cap,omitempty"`                 // 用户输入的成本上限
	FeedbackId                        int64                          `json:"feedback_id,omitempty"`                     // 监测链接组 id
	AoiOptimizationStrategy           *AoiOptimizationStrategy       `json:"aoi_optimization_strategy,omitempty"`       // 高价值范围探索
	SearchExpandTargetingSwitch       string                         `json:"search_expand_targeting_switch,omitempty"`  // 搜索定向拓展开关
	CloudUnionSpec                    *CloudUnionSpec                `json:"cloud_union_spec,omitempty"`                // 云选相关参数
	LiveRecommendStrategyEnabled      bool                           `json:"live_recommend_strategy_enabled,omitempty"` // 直播种草人群探索
	CustomCostRoiCap                  float32                        `json:"custom_cost_roi_cap,omitempty"`             // 控制成本的期望 ROI
	SmartTargetingMode                string                         `json:"smart_targeting_mode,omitempty"`            // 广告智能定向功能
	SmartCouponMode                   string                         `json:"smart_coupon_mode,omitempty"`               // 小店智券开关
}

type CloudUnionSpec struct {
	RoiGoal     string  `json:"roi_goal,omitempty"`     // 优化 ROI 目标
	ExpectedRoi float32 `json:"expected_roi,omitempty"` // 深度优化价值效果值
}

func (p *AdgroupsUpdateReq) Format() {
	p.GlobalReq.Format()
}

// 长度限制常量
const (
	MinUpdateAdgroupNameLength = 1
	MaxUpdateAdgroupNameLength = 180 // 字节长度
	BeginDateLength            = 10
	MaxEndDateLength           = 10
)

// Validate 验证广告组更新请求
func (p *AdgroupsUpdateReq) Validate() error {

	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	// 1. 验证account_id
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	// 2. 验证adgroup_id
	if p.AdgroupID == 0 {
		return errors.New("adgroup_id为必填")
	}

	// 3. 验证广告名称（如果传了）
	if p.AdgroupName != "" {
		if err := p.validateAdgroupName(); err != nil {
			return err
		}
	}

	// 4. 验证开始日期（如果传了）
	if p.BeginDate != "" {
		if err := p.validateBeginDate(); err != nil {
			return err
		}
	}

	// 5. 验证结束日期（如果传了）
	if p.EndDate != "" {
		if err := p.validateEndDate(); err != nil {
			return err
		}
	}

	// 6. 验证日期范围关系（如果两者都传了）
	if p.BeginDate != "" && p.EndDate != "" {
		if err := p.validateDateRange(); err != nil {
			return err
		}
	}

	// 7. 验证首日开始投放时间（如果传了）
	if p.FirstDayBeginTime != "" {
		if err := p.validateFirstDayBeginTime(); err != nil {
			return err
		}
	}

	// 8. 验证广告出价（如果传了）
	if p.BidAmount != 0 {
		if err := p.validateBidAmount(); err != nil {
			return err
		}
	}

	return nil
}

// validateAdgroupName 验证广告名称
func (p *AdgroupsUpdateReq) validateAdgroupName() error {
	// 验证字节长度
	if len(p.AdgroupName) < MinUpdateAdgroupNameLength || len(p.AdgroupName) > MaxUpdateAdgroupNameLength {
		return errors.New("adgroup_name长度必须在1-180字节之间")
	}
	return nil
}

// validateBeginDate 验证开始日期
func (p *AdgroupsUpdateReq) validateBeginDate() error {
	if len(p.BeginDate) != BeginDateLength {
		return errors.New("begin_date长度必须为10字节")
	}

	begin, err := time.Parse(DateFormat, p.BeginDate)
	if err != nil {
		return errors.New("begin_date格式错误，应为YYYY-MM-DD")
	}

	_ = begin
	return nil
}

// validateEndDate 验证结束日期
func (p *AdgroupsUpdateReq) validateEndDate() error {
	if len(p.EndDate) > MaxEndDateLength {
		return errors.New("end_date长度不能超过10字节")
	}

	end, err := time.Parse(DateFormat, p.EndDate)
	if err != nil {
		return errors.New("end_date格式错误，应为YYYY-MM-DD")
	}

	// 结束日期 >= 今天
	today := time.Now().Truncate(24 * time.Hour)
	if end.Before(today) {
		return errors.New("end_date不能小于今天")
	}

	return nil
}

// validateDateRange 验证日期范围关系
func (p *AdgroupsUpdateReq) validateDateRange() error {
	begin, err := time.Parse(DateFormat, p.BeginDate)
	if err != nil {
		return err
	}

	end, err := time.Parse(DateFormat, p.EndDate)
	if err != nil {
		return err
	}

	// 开始日期 <= 结束日期
	if begin.After(end) {
		return errors.New("begin_date不能大于end_date")
	}

	return nil
}

// validateFirstDayBeginTime 验证首日开始投放时间
func (p *AdgroupsUpdateReq) validateFirstDayBeginTime() error {
	if len(p.FirstDayBeginTime) > MaxFirstDayBeginTimeLength {
		return errors.New("first_day_begin_time长度不能超过8字节")
	}

	_, err := time.Parse(TimeFormat, p.FirstDayBeginTime)
	if err != nil {
		return errors.New("first_day_begin_time格式错误，应为HH:ii:ss")
	}

	return nil
}

// validateBidAmount 验证广告出价
func (p *AdgroupsUpdateReq) validateBidAmount() error {
	if p.BidAmount <= 0 {
		return errors.New("bid_amount必须大于0")
	}
	return nil
}

type AdgroupsUpdateResp struct {
	AdgroupId int64 `json:"adgroup_id,omitempty"` // 广告 id
}

// ========== 批量修改广告日限额 ==========
// https://developers.e.qq.com/v3.0/docs/api/adgroups/update_daily_budget

// 字段限制常量
const (
	MinUpdateDailyBudgetSpecCount = 1         // update_daily_budget_spec 最小长度
	MinUpdateDailyBudget          = 5000      // 日预算最小值（分），50元
	MaxUpdateDailyBudget          = 400000000 // 日预算最大值（分），4000000元
)

// UpdateDailyBudgetSpec 更新日限额条件
type UpdateDailyBudgetSpec struct {
	AdgroupID   int64 `json:"adgroup_id"`   // 广告 id (必填)
	DailyBudget int   `json:"daily_budget"` // 日预算，单位为分 (必填)，0=不限，否则 5000-400000000
}

// Validate 验证单个日限额条件
func (s *UpdateDailyBudgetSpec) Validate() error {
	if s.AdgroupID == 0 {
		return errors.New("adgroup_id为必填")
	}
	if s.DailyBudget != 0 && (s.DailyBudget < MinUpdateDailyBudget || s.DailyBudget > MaxUpdateDailyBudget) {
		return errors.New("daily_budget设置为0表示不限，否则须在5000-400000000分之间")
	}
	return nil
}

// AdgroupsUpdateDailyBudgetReq 批量修改广告日限额请求
// https://developers.e.qq.com/v3.0/docs/api/adgroups/update_daily_budget
type AdgroupsUpdateDailyBudgetReq struct {
	GlobalReq
	AccountID             int64                    `json:"account_id"`               // 广告主帐号 id (必填)
	UpdateDailyBudgetSpec []*UpdateDailyBudgetSpec `json:"update_daily_budget_spec"` // 更新日限额条件列表 (必填)
}

func (p *AdgroupsUpdateDailyBudgetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证批量修改广告日限额请求参数
func (p *AdgroupsUpdateDailyBudgetReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if len(p.UpdateDailyBudgetSpec) < MinUpdateDailyBudgetSpecCount {
		return errors.New("update_daily_budget_spec为必填，至少包含1个条件")
	}
	seen := make(map[int64]bool)
	for i, spec := range p.UpdateDailyBudgetSpec {
		if spec == nil {
			return errors.New("update_daily_budget_spec[" + itoa(i) + "]不能为空")
		}
		if err := spec.Validate(); err != nil {
			return errors.New("update_daily_budget_spec[" + itoa(i) + "]: " + err.Error())
		}
		if seen[spec.AdgroupID] {
			return errors.New("update_daily_budget_spec中adgroup_id不允许重复：" + itoa(int(spec.AdgroupID)))
		}
		seen[spec.AdgroupID] = true
	}
	return p.GlobalReq.Validate()
}

// UpdateDailyBudgetResultItem 批量修改日限额响应列表项
type UpdateDailyBudgetResultItem struct {
	Code      int    `json:"code"`       // 返回码
	Message   string `json:"message"`    // 英文返回消息
	MessageCn string `json:"message_cn"` // 中文返回消息
	AdgroupID int64  `json:"adgroup_id"` // 广告 id
}

// AdgroupsUpdateDailyBudgetResp 批量修改广告日限额响应
// https://developers.e.qq.com/v3.0/docs/api/adgroups/update_daily_budget
type AdgroupsUpdateDailyBudgetResp struct {
	List       []*UpdateDailyBudgetResultItem `json:"list"`         // 返回信息列表，顺序与请求一致
	FailIDList []int64                        `json:"fail_id_list"` // 失败的 id 集合
}

// ========== 批量修改广告开启/暂停状态 ==========
// https://developers.e.qq.com/v3.0/docs/api/adgroups/update_configured_status

// 字段限制常量
const (
	MaxUpdateConfiguredStatusSpecCount = 100 // update_configured_status_spec 最大长度
)

// UpdateConfiguredStatusSpec 更新客户设置状态条件
type UpdateConfiguredStatusSpec struct {
	AdgroupID        int64  `json:"adgroup_id"`        // 广告 id (必填)
	ConfiguredStatus string `json:"configured_status"` // 客户设置的状态 (必填)：AD_STATUS_NORMAL / AD_STATUS_SUSPEND
}

// Validate 验证单个状态更新条件
func (s *UpdateConfiguredStatusSpec) Validate() error {
	if s.AdgroupID == 0 {
		return errors.New("adgroup_id为必填")
	}
	if s.ConfiguredStatus == "" {
		return errors.New("configured_status为必填")
	}
	if s.ConfiguredStatus != ConfiguredStatusNormal && s.ConfiguredStatus != ConfiguredStatusSuspend {
		return errors.New("configured_status值无效，允许值：AD_STATUS_NORMAL、AD_STATUS_SUSPEND")
	}
	return nil
}

// AdgroupsUpdateConfiguredStatusReq 批量修改广告开启/暂停状态请求
// https://developers.e.qq.com/v3.0/docs/api/adgroups/update_configured_status
type AdgroupsUpdateConfiguredStatusReq struct {
	GlobalReq
	AccountID                  int64                         `json:"account_id"`                    // 广告主帐号 id (必填)
	UpdateConfiguredStatusSpec []*UpdateConfiguredStatusSpec `json:"update_configured_status_spec"` // 更新状态条件列表 (必填)，最大100
}

func (p *AdgroupsUpdateConfiguredStatusReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证批量修改广告状态请求参数
func (p *AdgroupsUpdateConfiguredStatusReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if len(p.UpdateConfiguredStatusSpec) == 0 {
		return errors.New("update_configured_status_spec为必填，至少包含1个条件")
	}
	if len(p.UpdateConfiguredStatusSpec) > MaxUpdateConfiguredStatusSpecCount {
		return errors.New("update_configured_status_spec数组长度不能超过100")
	}
	seen := make(map[int64]bool)
	for i, spec := range p.UpdateConfiguredStatusSpec {
		if spec == nil {
			return errors.New("update_configured_status_spec[" + itoa(i) + "]不能为空")
		}
		if err := spec.Validate(); err != nil {
			return errors.New("update_configured_status_spec[" + itoa(i) + "]: " + err.Error())
		}
		if seen[spec.AdgroupID] {
			return errors.New("update_configured_status_spec中adgroup_id不允许重复：" + itoa(int(spec.AdgroupID)))
		}
		seen[spec.AdgroupID] = true
	}
	return p.GlobalReq.Validate()
}

// AdgroupsUpdateConfiguredStatusResp 批量修改广告开启/暂停状态响应
// https://developers.e.qq.com/v3.0/docs/api/adgroups/update_configured_status
type AdgroupsUpdateConfiguredStatusResp struct {
	List       []*UpdateDailyBudgetResultItem `json:"list"`         // 返回信息列表，顺序与请求一致
	FailIDList []int64                        `json:"fail_id_list"` // 失败的 id 集合
}

// ========== 批量修改广告出价 ==========
// https://developers.e.qq.com/v3.0/docs/api/adgroups/update_bid_amount

// 字段限制常量
const (
	MaxUpdateBidAmountSpecCount = 100 // update_bid_amount_spec 最大长度
)

// UpdateBidAmountSpec 更新广告出价条件
type UpdateBidAmountSpec struct {
	AdgroupID int64 `json:"adgroup_id"` // 广告 id (必填)
	BidAmount int64 `json:"bid_amount"` // 广告出价，单位为分 (必填)，ADX 程序化广告默认填写 200
}

// Validate 验证单个出价条件
func (s *UpdateBidAmountSpec) Validate() error {
	if s.AdgroupID == 0 {
		return errors.New("adgroup_id为必填")
	}
	if s.BidAmount <= 0 {
		return errors.New("bid_amount必须大于0")
	}
	return nil
}

// AdgroupsUpdateBidAmountReq 批量修改广告出价请求
// https://developers.e.qq.com/v3.0/docs/api/adgroups/update_bid_amount
type AdgroupsUpdateBidAmountReq struct {
	GlobalReq
	AccountID           int64                  `json:"account_id"`             // 广告主帐号 id (必填)
	UpdateBidAmountSpec []*UpdateBidAmountSpec `json:"update_bid_amount_spec"` // 更新出价条件列表 (必填)，最大100
}

func (p *AdgroupsUpdateBidAmountReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证批量修改广告出价请求参数
func (p *AdgroupsUpdateBidAmountReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if len(p.UpdateBidAmountSpec) == 0 {
		return errors.New("update_bid_amount_spec为必填，至少包含1个条件")
	}
	if len(p.UpdateBidAmountSpec) > MaxUpdateBidAmountSpecCount {
		return errors.New("update_bid_amount_spec数组长度不能超过100")
	}
	seen := make(map[int64]bool)
	for i, spec := range p.UpdateBidAmountSpec {
		if spec == nil {
			return errors.New("update_bid_amount_spec[" + itoa(i) + "]不能为空")
		}
		if err := spec.Validate(); err != nil {
			return errors.New("update_bid_amount_spec[" + itoa(i) + "]: " + err.Error())
		}
		if seen[spec.AdgroupID] {
			return errors.New("update_bid_amount_spec中adgroup_id不允许重复：" + itoa(int(spec.AdgroupID)))
		}
		seen[spec.AdgroupID] = true
	}
	return p.GlobalReq.Validate()
}

// AdgroupsUpdateBidAmountResp 批量修改广告出价响应
// https://developers.e.qq.com/v3.0/docs/api/adgroups/update_bid_amount
type AdgroupsUpdateBidAmountResp struct {
	List       []*UpdateDailyBudgetResultItem `json:"list"`         // 返回信息列表，顺序与请求一致
	FailIDList []int64                        `json:"fail_id_list"` // 失败的 id 集合
}

// ========== 批量修改广告投放起止时间 ==========
// https://developers.e.qq.com/v3.0/docs/api/adgroups/update_datetime

// 字段限制常量
const (
	MaxUpdateDatetimeSpecCount = 100 // update_datetime_spec 最大长度
	DateLength                 = 10  // 日期字段长度（YYYY-MM-DD）
)

// UpdateDatetimeSpec 更新投放时间条件
type UpdateDatetimeSpec struct {
	AdgroupID  int64  `json:"adgroup_id"`            // 广告 id (必填)
	BeginDate  string `json:"begin_date,omitempty"`  // 开始投放日期，YYYY-MM-DD，<= end_date
	EndDate    string `json:"end_date,omitempty"`    // 结束投放日期，YYYY-MM-DD，>= 今天，>= begin_date；长度 0-10 字节
	TimeSeries string `json:"time_series,omitempty"` // 投放时间段，336字节的0/1字符串，不允许全为0
}

// Validate 验证单个投放时间更新条件
func (s *UpdateDatetimeSpec) Validate() error {
	if s.AdgroupID == 0 {
		return errors.New("adgroup_id为必填")
	}
	if s.BeginDate == "" && s.EndDate == "" && s.TimeSeries == "" {
		return errors.New("begin_date、end_date、time_series中至少需要传入一个参数")
	}
	if s.BeginDate != "" {
		if len(s.BeginDate) != DateLength {
			return errors.New("begin_date格式错误，应为YYYY-MM-DD")
		}
		if _, err := time.Parse(DateFormat, s.BeginDate); err != nil {
			return errors.New("begin_date格式错误，应为YYYY-MM-DD")
		}
	}
	if s.EndDate != "" {
		if len(s.EndDate) != DateLength {
			return errors.New("end_date格式错误，应为YYYY-MM-DD")
		}
		end, err := time.Parse(DateFormat, s.EndDate)
		if err != nil {
			return errors.New("end_date格式错误，应为YYYY-MM-DD")
		}
		today := time.Now().Truncate(24 * time.Hour)
		if end.Before(today) {
			return errors.New("end_date不能早于今天")
		}
		if s.BeginDate != "" {
			begin, _ := time.Parse(DateFormat, s.BeginDate)
			if end.Before(begin) {
				return errors.New("end_date不能早于begin_date")
			}
		}
	}
	if s.TimeSeries != "" {
		if len(s.TimeSeries) != TimeSeriesLength {
			return errors.New("time_series长度必须为336字节")
		}
		allZero := true
		for _, c := range s.TimeSeries {
			if c != '0' && c != '1' {
				return errors.New("time_series只能包含0和1")
			}
			if c == '1' {
				allZero = false
			}
		}
		if allZero {
			return errors.New("time_series不允许全部为0")
		}
	}
	return nil
}

// AdgroupsUpdateDatetimeReq 批量修改广告投放起止时间请求
// https://developers.e.qq.com/v3.0/docs/api/adgroups/update_datetime
type AdgroupsUpdateDatetimeReq struct {
	GlobalReq
	AccountID          int64                 `json:"account_id"`           // 广告主帐号 id (必填)
	UpdateDatetimeSpec []*UpdateDatetimeSpec `json:"update_datetime_spec"` // 更新投放时间条件列表 (必填)，最大100
}

func (p *AdgroupsUpdateDatetimeReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证批量修改广告投放时间请求参数
func (p *AdgroupsUpdateDatetimeReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if len(p.UpdateDatetimeSpec) == 0 {
		return errors.New("update_datetime_spec为必填，至少包含1个条件")
	}
	if len(p.UpdateDatetimeSpec) > MaxUpdateDatetimeSpecCount {
		return errors.New("update_datetime_spec数组长度不能超过100")
	}
	seen := make(map[int64]bool)
	for i, spec := range p.UpdateDatetimeSpec {
		if spec == nil {
			return errors.New("update_datetime_spec[" + itoa(i) + "]不能为空")
		}
		if err := spec.Validate(); err != nil {
			return errors.New("update_datetime_spec[" + itoa(i) + "]: " + err.Error())
		}
		if seen[spec.AdgroupID] {
			return errors.New("update_datetime_spec中adgroup_id不允许重复：" + itoa(int(spec.AdgroupID)))
		}
		seen[spec.AdgroupID] = true
	}
	return p.GlobalReq.Validate()
}

// AdgroupsUpdateDatetimeResp 批量修改广告投放起止时间响应
// https://developers.e.qq.com/v3.0/docs/api/adgroups/update_datetime
type AdgroupsUpdateDatetimeResp struct {
	List       []*UpdateDailyBudgetResultItem `json:"list"`         // 返回信息列表，顺序与请求一致
	FailIDList []int64                        `json:"fail_id_list"` // 失败的 id 集合
}
