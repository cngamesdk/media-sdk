package model

import (
	"errors"
	"unicode/utf8"
)

const (
	// Operation 项目状态
	OperationEnable  = "ENABLE"  // 开启（默认值）
	OperationDisable = "DISABLE" // 关闭

	// LandingType 推广目的/落地页类型
	LandingTypeApp          = "APP"           // 应用
	LandingTypeLink         = "LINK"          // 销售线索投放
	LandingTypeMicroGame    = "MICRO_GAME"    // 小程序
	LandingTypeShop         = "SHOP"          // 电商
	LandingTypeQuickApp     = "QUICK_APP"     // 快应用
	LandingTypeNativeAction = "NATIVE_ACTION" // 原生互动
	LandingTypeDPA          = "DPA"           // 商品目录

	// DeliveryMode 投放模式（根据场景不同）
	DeliveryModeManual     = "MANUAL"     // 手动投放
	DeliveryModeProcedural = "PROCEDURAL" // 自动投放

	// AppPromotionType 应用推广子目标
	AppPromotionTypeDownload = "DOWNLOAD" // 应用下载
	AppPromotionTypeLaunch   = "LAUNCH"   // 应用调起
	AppPromotionTypeReserve  = "RESERVE"  // 预约下载

	// MarketingGoal 营销场景
	MarketingGoalVideoImage = "VIDEO_AND_IMAGE" // 短视频/图片
	MarketingGoalLive       = "LIVE"            // 直播

	// MarketingType 营销类型
	MarketingTypeAll    = "ALL"    // 通投
	MarketingTypeSearch = "SEARCH" // 搜索

	// AdType 广告类型
	AdTypeAll = "ALL" // 通投

	// InventoryType 广告位类型
	InventoryTypeHomedAggregate = "INVENTORY_HOMED_AGGREGATE" // 住小帮

	// DeliveryType 投放类型
	DeliveryTypeNormal         = "NORMAL"          // 常规投放（默认值）
	DeliveryTypeDuration       = "DURATION"        // 极速智投（仅支持搜索）
	DeliveryTypeUBXIntelligent = "UBX_INTELLIGENT" // 智能托管（仅支持UBA项目）

	// 项目名称长度限制
	MinNameLength = 1
	MaxNameLength = 50 // 注意：两个英文字符算1个字

	// 常量：预算组相关说明
	BudgetGroupNote              = "一个项目一生只能被一个预算组关联，项目绑定预算组ID后，不允许修改。如果预算组被删除，绑定该预算组的项目也不能再绑定新的预算组。"
	BudgetGroupWhiteListRequired = "该功能白名单开放，如需使用请联系销售"

	// AIGC动态创意开关
	AIGCDynamicCreativeON  = "ON"  // 开启
	AIGCDynamicCreativeOFF = "OFF" // 关闭
)

type ProjectCreateReq struct {
	accessTokenReq
	AdvertiserID                   int64   `json:"advertiser_id"`                                // 投放账户id (必填)
	Operation                      string  `json:"operation,omitempty"`                          // 项目状态，允许值：ENABLE 开启（默认值），DISABLE 关闭
	DeliveryMode                   string  `json:"delivery_mode,omitempty"`                      // 投放模式
	LandingType                    string  `json:"landing_type"`                                 // 推广目的/落地页类型 (必填)
	AppPromotionType               string  `json:"app_promotion_type,omitempty"`                 // 应用推广子目标
	MarketingGoal                  string  `json:"marketing_goal,omitempty"`                     // 营销场景
	MarketingType                  string  `json:"marketing_type,omitempty"`                     // 营销类型
	AdType                         string  `json:"ad_type,omitempty"`                            // 广告类型
	Name                           string  `json:"name"`                                         // 项目名称 (必填)
	DeliveryType                   string  `json:"delivery_type,omitempty"`                      // 投放类型，默认NORMAL
	BudgetGroupID                  int64   `json:"budget_group_id,omitempty"`                    // 预算组ID
	InventoryType                  string  `json:"inventory_type,omitempty"`                     // 广告位类型
	OptimizationGoal               string  `json:"optimization_goal,omitempty"`                  // 优化目标
	DeepOptimization               string  `json:"deep_optimization,omitempty"`                  // 深度优化方式
	SevenRoiGoal                   float64 `json:"seven_roi_goal,omitempty"`                     // 7日ROI系数，范围[0.01,5]
	AIGCDynamicCreativeSwitch      string  `json:"aigc_dynamic_creative_switch,omitempty"`       // 是否开启AIGC动态创意
	StarTaskID                     int64   `json:"star_task_id,omitempty"`                       // 星广联投任务ID
	StarAutoMaterialAdditionSwitch string  `json:"star_auto_material_addition_switch,omitempty"` // 星广联投自动优选素材开关
	StarAutoDeliverySwitch         string  `json:"star_auto_delivery_switch,omitempty"`          // 星广联投全自动化开关
	RtaID                          int64   `json:"rta_id,omitempty"`

	//搜索快投
	SearchQuick

	//搜索关键词
	SearchKeywords

	//搜索蓝海流量投放
	BlueFlowPackage *BlueFlowPackage `json:"blue_flow_package,omitempty"` // 搜索蓝海流量投放
	RelatedProduct  *RelatedProduct  `json:"related_product,omitempty"`   // 商品

	//营销产品与投放载体
	DownloadUrl        string `json:"download_url,omitempty"`  // 下载链接
	AppName            string `json:"app_name,omitempty"`      // 应用名称
	DownloadType       string `json:"download_type,omitempty"` // 下载方式
	DownloadMode       string `json:"download_mode,omitempty"`
	QuickAppID         int64  `json:"quick_app_id,omitempty"`          // 快应用资产id，从【查询快应用信息】接口获取，仅支持已通过审核的快应用资产
	LaunchType         string `json:"launch_type,omitempty"`           // 调起方式（仅对应用目标下的应用调起方式，快应用调起方式请使用download_type参数）
	PromotionType      string `json:"promotion_type,omitempty"`        // 投放内容
	OpenURLType        string `json:"open_url_type"`                   // 直达链接类型 (条件必填)
	OpenURLParams      string `json:"open_url_params,omitempty"`       // 直达链接检测参数
	OpenURLField       string `json:"open_url_field,omitempty"`        // 直达链接字段选择
	OpenUrl            string `json:"open_url,omitempty"`              // Deeplink直达链接
	UlinkUrlType       string `json:"ulink_url_type,omitempty"`        // ulink直达链接备用链接类型
	UlinkUrl           string `json:"ulink_url,omitempty"`             // ulink直达链接备用链接
	SubscribeUrl       string `json:"subscribe_url,omitempty"`         // 预约下载链接
	AssetType          string `json:"asset_type,omitempty"`            // 资产类型
	MultiAssetType     string `json:"multi_asset_type,omitempty"`      // 多投放载体
	MicroPromotionType string `json:"micro_promotion_type,omitempty"`  // 小程序类型
	DpaAdtype          string `json:"dpa_adtype,omitempty"`            // DPA营销类型
	MicroAppInstanceId int64  `json:"micro_app_instance_id,omitempty"` // 微信、字节小程序/小游戏资产id
	NativeSetting      struct {
		AwemeId string `json:"aweme_id,omitempty"`
	} `json:"native_setting,omitempty"` // 微信、字节小程序/小游戏资产id
	OptimizeGoal struct {
		AssetIds                   []int64 `json:"asset_ids,omitempty"`                       // 事件管理资产 id
		ExternalAction             string  `json:"external_action,omitempty"`                 // 优化目标
		GameAddictionID            string  `json:"game_addiction_id,omitempty"`               // 关键行为ID
		PaidSwitch                 int     `json:"paid_switch,omitempty"`                     // 付费开关
		DeepExternalAction         string  `json:"deep_external_action,omitempty"`            // 深度转化目标
		ValueOptimizedType         string  `json:"value_optimized_type,omitempty"`            // 目标优化类型
		LandingPageStayTime        int64   `json:"landing_page_stay_time,omitempty"`          // 店铺停留时长（毫秒）
		Yuntu5aBrandID             string  `json:"yuntu_5a_brand_id,omitempty"`               // 云图品牌ID
		Yuntu5aBrandMainIndustryID string  `json:"yuntu_5a_brand_main_industry_id,omitempty"` // 云图品牌行业ID
	} `json:"optimize_goal,omitempty"` // 优化目标
	DeliveryRange struct {
		InventoryCatalog string   `json:"inventory_catalog,omitempty"` // 投放版位大类
		InventoryType    []string `json:"inventory_type,omitempty"`    // 投放位置（首选媒体）
		UnionVideoType   string   `json:"union_video_type,omitempty"`  // 投放形式（穿山甲视频创意类型）
	} `json:"delivery_range,omitempty"` // 投放版位
	Audience        interface{}      `json:"audience,omitempty"`          // 定向
	DeliverySetting *DeliverySetting `json:"delivery_setting,omitempty"`  // 排期、预算、出价
	TrackUrlSetting *TrackUrlSetting `json:"track_url_setting,omitempty"` // 监测链接
}

// SearchQuick 搜索快投
type SearchQuick struct {
	SearchBidRatio float32 `json:"search_bid_ratio,omitempty"` // 出价系数
	AudienceExtend string  `json:"audience_extend,omitempty"`  // 定向拓展
}

// SearchKeywords 搜索关键词
type SearchKeywords struct {
	Keywords struct {
		Word      string  `json:"word,omitempty"`       // 关键词
		BidType   string  `json:"bid_type,omitempty"`   // 出价类型
		MatchType string  `json:"match_type,omitempty"` // 匹配类型
		Bid       float32 `json:"bid,omitempty"`        // 出价
	} `json:"keywords,omitempty"` // 待添加搜索关键词列表
	AutoExtendTraffic string `json:"auto_extend_traffic,omitempty"` // 智能拓流
}

// 常量定义
const (
	// 蓝海流量设置
	BlueFlowON  = "ON"  // 启用
	BlueFlowOFF = "OFF" // 不启用

	// 白名单key
	WhiteListBlueFlow        = "blue_flow"              // 蓝海流量投放白名单
	WhiteListSearchMergeAuto = "search_merge_auto_blue" // 极速智投蓝海关键词白名单
)

// BlueFlowPackage 搜索蓝海流量投放相关参数
type BlueFlowPackage struct {
	BlueFlowPackageSetting string   `json:"blue_flow_package_setting"`        // 蓝海流量设置 (条件必填)
	BlueFlowPackageID      int64    `json:"blue_flow_package_id,omitempty"`   // 蓝海流量包ID (条件必填)
	BlueFlowKeywordName    []string `json:"blue_flow_keyword_name,omitempty"` // 蓝海关键词
}

// ValidateBlueFlowPackage 验证搜索蓝海流量投放配置
func (b *BlueFlowPackage) Validate(adType, deliveryType string, hasWhiteList, hasSearchMergeAutoWhiteList bool) error {
	// 1. 仅当ad_type=SEARCH时支持传入
	if adType != "SEARCH" {
		if b != nil {
			return errors.New("blue_flow_package仅当ad_type=SEARCH时支持传入")
		}
		return nil
	}

	// 如果没传blue_flow_package，不需要验证
	if b == nil {
		return nil
	}

	// 2. 验证蓝海流量设置
	if err := b.validateSetting(hasWhiteList); err != nil {
		return err
	}

	// 3. 根据投放类型验证
	if deliveryType == DeliveryTypeDuration {
		return b.validateForDuration(hasSearchMergeAutoWhiteList)
	} else {
		return b.validateForNormal()
	}
}

// validateSetting 验证蓝海流量设置
func (b *BlueFlowPackage) validateSetting(hasWhiteList bool) error {
	// blue_flow_package_setting为条件必填
	if b.BlueFlowPackageSetting == "" {
		return errors.New("blue_flow_package_setting为条件必填")
	}

	// 验证设置值有效性
	if b.BlueFlowPackageSetting != BlueFlowON && b.BlueFlowPackageSetting != BlueFlowOFF {
		return errors.New("blue_flow_package_setting值无效，允许值：ON、OFF")
	}

	// 蓝海流量投放为白名单功能
	if !hasWhiteList {
		return errors.New("蓝海流量投放为白名单功能，当前账户没有权限")
	}

	return nil
}

// validateForDuration 验证极速智投场景
func (b *BlueFlowPackage) validateForDuration(hasSearchMergeAutoWhiteList bool) error {
	// 极速智投场景下，通过blue_flow_keyword_name设置进行蓝海投放
	// 若不传入blue_flow_keyword_name即不进行蓝海投放

	if len(b.BlueFlowKeywordName) > 0 {
		// 极速智投蓝海关键词需要白名单
		if !hasSearchMergeAutoWhiteList {
			return errors.New("极速智投蓝海关键词需要search_merge_auto_blue白名单")
		}
	}

	// 极速智投场景不支持blue_flow_package_id
	if b.BlueFlowPackageID != 0 {
		return errors.New("极速智投场景下不支持blue_flow_package_id")
	}

	return nil
}

// validateForNormal 验证常规场景
func (b *BlueFlowPackage) validateForNormal() error {
	// 仅当blue_flow_package_setting=ON时，需要验证blue_flow_package_id
	if b.BlueFlowPackageSetting == BlueFlowON {
		if b.BlueFlowPackageID == 0 {
			return errors.New("启用蓝海流量投放时，blue_flow_package_id为必填")
		}
	}

	// 常规场景不支持blue_flow_keyword_name
	if len(b.BlueFlowKeywordName) > 0 {
		return errors.New("blue_flow_keyword_name仅极速智投场景下有效")
	}

	return nil
}

// RelatedProduct 商品
type RelatedProduct struct {
	ProductSetting    string `json:"product_setting,omitempty"`     // 商品库设置
	ProductPlatformId int64  `json:"product_platform_id,omitempty"` // 商品库ID
	ProductId         string `json:"product_id,omitempty"`          // 产品ID
	UniqueProductId   int64  `json:"unique_product_id,omitempty"`   // 升级版商品库产品ID
	Products          struct {
		ProductId         string `json:"product_id,omitempty"`          // 产品ID
		ProductPlatformId int64  `json:"product_platform_id,omitempty"` // 产品库ID
		UniqueProductId   int64  `json:"unique_product_id,omitempty"`   // 升级版商品库产品ID
		AssetId           int64  `json:"asset_id,omitempty"`            // 投放条件ID
	} `json:"products,omitempty"` // 产品ID列表
}

func (p *ProjectCreateReq) Validate() error {
	// 1. 验证应用推广子目标
	if p.LandingType == LandingTypeApp {
		if p.AppPromotionType == "" {
			return errors.New("当 landing_type = APP 时，app_promotion_type 为必填")
		}

		// 自动投放模式仅支持下载
		if p.DeliveryMode == DeliveryModeProcedural &&
			p.AppPromotionType != AppPromotionTypeDownload {
			return errors.New("当 delivery_mode = PROCEDURAL 时，app_promotion_type 仅支持 DOWNLOAD")
		}

		// 直播场景仅支持下载和调起
		if p.MarketingGoal == MarketingGoalLive &&
			p.AppPromotionType != AppPromotionTypeDownload &&
			p.AppPromotionType != AppPromotionTypeLaunch {
			return errors.New("当 marketing_goal = LIVE 时，app_promotion_type 仅支持 DOWNLOAD、LAUNCH")
		}
	}

	// 2. 验证营销场景
	if p.MarketingGoal == MarketingGoalLive {
		// 检查直播投放协议（需在业务层验证）

		// 验证 landing_type 是否支持直播
		supportedLandingTypes := map[string]bool{
			LandingTypeApp:          true,
			LandingTypeMicroGame:    true,
			LandingTypeLink:         true,
			LandingTypeNativeAction: true,
		}
		if !supportedLandingTypes[p.LandingType] {
			return errors.New("直播营销场景仅支持应用/小程序/销售线索/原生互动 landing_type")
		}

		// 自动投放 + LINK 场景需白名单（需在业务层验证）
	}

	// 3. 验证营销类型
	if p.MarketingGoal == MarketingGoalLive && p.MarketingType != MarketingTypeAll {
		return errors.New("当 marketing_goal = LIVE 时，marketing_type 仅支持 ALL")
	}

	// 搜索场景验证
	if p.MarketingType == MarketingTypeSearch {
		// 检查是否支持搜索
		supportedSearchLandingTypes := map[string]bool{
			LandingTypeApp:          true,
			LandingTypeLink:         true,
			LandingTypeNativeAction: true,
			LandingTypeShop:         true,
		}

		if !supportedSearchLandingTypes[p.LandingType] {
			return errors.New("搜索场景仅支持 APP/LINK/NATIVE_ACTION/SHOP landing_type")
		}

		if p.DeliveryMode != DeliveryModeManual {
			return errors.New("搜索场景仅支持 MANUAL 投放模式")
		}
	}

	// 快应用/小程序自动投放场景
	if (p.LandingType == LandingTypeQuickApp || p.LandingType == LandingTypeMicroGame) &&
		p.DeliveryMode == DeliveryModeProcedural &&
		p.MarketingType != MarketingTypeAll {
		return errors.New("快应用/小程序自动投放时，marketing_type 仅支持 ALL")
	}

	// 住小帮广告位验证
	if p.InventoryType == InventoryTypeHomedAggregate && p.AdType != AdTypeAll {
		return errors.New("住小帮广告位仅支持 ad_type = ALL")
	}

	if validateErr := p.validateName(); validateErr != nil {
		return validateErr
	}
	return nil
}

// validateName 验证项目名称长度
func (p *ProjectCreateReq) validateName() error {
	if p.Name == "" {
		return errors.New("项目名称为必填")
	}

	// 计算中文字符数（两个英文字符算1个字）
	// 简化处理：使用UTF-8字符计数，实际业务可能需要更精确的计算
	charCount := utf8.RuneCountInString(p.Name)

	if charCount < MinNameLength || charCount > MaxNameLength {
		return errors.New("项目名称长度必须在1-50个字之间（两个英文字符占1个字）")
	}

	return nil
}

// validateDeliveryType 验证投放类型
func (p *ProjectCreateReq) validateDeliveryType() error {
	switch p.DeliveryType {
	case "":
		// 默认值，不需要验证
		return nil
	case DeliveryTypeNormal:
		return nil
	case DeliveryTypeDuration:
		return p.validateDurationDelivery()
	case DeliveryTypeUBXIntelligent:
		return p.validateUBXIntelligentDelivery()
	default:
		return errors.New("无效的投放类型")
	}
}

// validateDurationDelivery 验证极速智投投放类型
func (p *ProjectCreateReq) validateDurationDelivery() error {
	// 极速智投仅支持搜索
	if p.MarketingType != MarketingTypeSearch {
		return errors.New("极速智投仅支持搜索营销类型")
	}

	if p.AdType != "SEARCH" {
		return errors.New("极速智投仅支持 ad_type=SEARCH")
	}

	// 验证支持的组合
	supported := false

	// 应用营销目的 && 直播/短视频营销场景
	if p.LandingType == LandingTypeApp &&
		(p.MarketingGoal == MarketingGoalLive || p.MarketingGoal == MarketingGoalVideoImage) {
		supported = true
	}

	// 销售线索营销目的 && 直播/短视频营销场景
	if p.LandingType == LandingTypeLink &&
		(p.MarketingGoal == MarketingGoalLive || p.MarketingGoal == MarketingGoalVideoImage) {
		supported = true
	}

	// 小程序营销目的 && 直播/短视频营销场景
	if p.LandingType == LandingTypeMicroGame &&
		(p.MarketingGoal == MarketingGoalLive || p.MarketingGoal == MarketingGoalVideoImage) {
		supported = true
	}

	// 电商营销目的 && 短视频营销场景
	if p.LandingType == LandingTypeShop && p.MarketingGoal == MarketingGoalVideoImage {
		supported = true
	}

	if !supported {
		return errors.New("极速智投不支持当前营销目的和营销场景的组合")
	}

	return nil
}

// validateUBXIntelligentDelivery 验证智能托管投放类型
func (p *ProjectCreateReq) validateUBXIntelligentDelivery() error {
	// 智能托管仅支持UBA项目（需在业务层验证）

	// 营销目的为小程序
	if p.LandingType != LandingTypeMicroGame {
		return errors.New("智能托管仅支持营销目的为小程序")
	}

	// 营销场景为短视频&图文
	if p.MarketingGoal != MarketingGoalVideoImage {
		return errors.New("智能托管仅支持营销场景为短视频&图文")
	}

	// 投放模式为自动投放
	if p.DeliveryMode != DeliveryModeProcedural {
		return errors.New("智能托管仅支持自动投放模式")
	}

	// 营销产品（投放载体）为字节小游戏（需在业务层验证）

	// 优化目标&深度优化方式为变现ROI-7R && 7日ROI系数
	if p.OptimizationGoal != "ROI-7R" {
		return errors.New("智能托管仅支持优化目标为变现ROI-7R")
	}

	if p.SevenRoiGoal <= 0 {
		return errors.New("智能托管需要设置有效的ROI系数")
	}

	return nil
}

// IsBudgetGroupImmutable 判断预算组ID是否不可修改
func (p *ProjectCreateReq) IsBudgetGroupImmutable(existingBudgetGroupID int64) bool {
	// 如果已有预算组ID且与新传入的不同，则不允许修改
	return existingBudgetGroupID != 0 && existingBudgetGroupID != p.BudgetGroupID
}

// IsLiveWhiteListRequired 判断是否需要直播白名单
func (p *ProjectCreateReq) IsLiveWhiteListRequired() bool {
	return p.MarketingGoal == MarketingGoalLive &&
		p.LandingType == LandingTypeLink &&
		p.DeliveryMode == DeliveryModeProcedural
}

func (p *ProjectCreateReq) GetHeaders() headersMap {
	headers := p.accessTokenReq.GetHeaders()
	headers.Json()
	return headers
}

// DPAAudience DPA商品定向（当landing_type=DPA时投放目标参数）
type DPAAudience struct {
	DPACity             string  `json:"dpa_city,omitempty"`               // 地域匹配-商品所在城市
	DPALbs              string  `json:"dpa_lbs,omitempty"`                // 地域匹配-适地性服务
	DPARtaSwitch        string  `json:"dpa_rta_switch,omitempty"`         // RTA重定向开关
	RtaID               int64   `json:"rta_id,omitempty"`                 // RTA策略ID（条件必填）
	DPARtaRecommendType string  `json:"dpa_rta_recommend_type,omitempty"` // RTA推荐逻辑（条件必填）
	DpaCategories       []int64 `json:"dpa_categories,omitempty"`         // 商品投放范围
	DpaProductTarget    struct {
		Title string `json:"title,omitempty"` // 筛选字段
		Rule  string `json:"rule,omitempty"`  // 定向规则
		Type  string `json:"type,omitempty"`  // 字段类型
		Value string `json:"value,omitempty"` // 规则值
	} `json:"dpa_product_target,omitempty"` // 自定义筛选条件（商品投放条件）
}

// CommonAudience 常规人群定向（其余情况均按下表传入)
type CommonAudience struct {
	AudiencePackageID      int64          `json:"audience_package_id,omitempty"`      // 定向包ID
	District               string         `json:"district"`                           // 地域类型 (必填)
	Geolocation            []*Geolocation `json:"geolocation,omitempty"`              // 地图位置 (商圈专用)
	RegionVersion          string         `json:"region_version,omitempty"`           // 行政区版本号 (条件必填)
	City                   []int64        `json:"city,omitempty"`                     // 地域定向省市或者区县列表 (条件必填)
	LocationType           string         `json:"location_type,omitempty"`            // 位置类型 (条件必填)
	Gender                 string         `json:"gender,omitempty"`                   // 性别
	Age                    []string       `json:"age,omitempty"`                      // 年龄
	RetargetingTagsInclude []int64        `json:"retargeting_tags_include,omitempty"` // 定向人群包列表
	RetargetingTagsExclude []int64        `json:"retargeting_tags_exclude,omitempty"` // 排除人群包列表
	InterestActionMode     string         `json:"interest_action_mode,omitempty"`     // 行为兴趣模式
	ActionDays             []int          `json:"action_days,omitempty"`              // 用户发生行为天数
	ActionCategories       []int64        `json:"action_categories,omitempty"`        // 行为类目词
	ActionWords            []int64        `json:"action_words,omitempty"`             // 行为关键词
	InterestCategories     []int64        `json:"interest_categories,omitempty"`      // 兴趣类目
	InterestWords          []int64        `json:"interest_words,omitempty"`           // 兴趣关键词

	// 抖音达人定向相关
	AwesomeFanBehaviors       []string `json:"awesome_fan_behaviors,omitempty"`        // 抖音达人互动用户行为类型
	AwesomeFanTimeScope       string   `json:"awesome_fan_time_scope,omitempty"`       // 抖音达人互动行为时间范围
	AwesomeFanCategories      []int64  `json:"awesome_fan_categories,omitempty"`       // 抖音达人分类ID列表
	AwesomeFanAccounts        []int64  `json:"awesome_fan_accounts,omitempty"`         // 抖音达人ID列表
	SuperiorPopularityType    string   `json:"superior_popularity_type,omitempty"`     // 媒体定向
	FlowPackage               []int64  `json:"flow_package,omitempty"`                 // 定向流量包
	ExcludeFlowPackage        []int64  `json:"exclude_flow_package,omitempty"`         // 排除流量包
	Platform                  []string `json:"platform,omitempty"`                     // 投放平台列表
	AndroidOsv                string   `json:"android_osv,omitempty"`                  // 最低安卓版本
	IosOsv                    string   `json:"ios_osv,omitempty"`                      // 最低IOS版本
	HarmonyOsv                string   `json:"harmony_osv,omitempty"`                  // 鸿蒙版本
	DeviceType                []string `json:"device_type,omitempty"`                  // 设备类型
	Ac                        []string `json:"ac,omitempty"`                           // 网络类型
	Carrier                   []string `json:"carrier,omitempty"`                      // 运营商
	CarrierRegionOptimize     string   `json:"carrier_region_optimize,omitempty"`      // 运营商号段开关
	HideIfExists              string   `json:"hide_if_exists,omitempty"`               // 过滤已安装
	HideIfConverted           string   `json:"hide_if_converted,omitempty"`            // 过滤已转化用户
	ConvertedTimeDuration     string   `json:"converted_time_duration,omitempty"`      // 过滤时间范围
	FilterEvent               []string `json:"filter_event,omitempty"`                 // 自定义过滤事件
	FilterAwemeAbnormalActive string   `json:"filter_aweme_abnormal_active,omitempty"` // 过滤高活跃用户
	FilterAwemeFansCount      int64    `json:"filter_aweme_fans_count,omitempty"`      // 过滤高关注数用户
	FilterOwnAwemeFans        string   `json:"filter_own_aweme_fans,omitempty"`        // 过滤自己的粉丝
	DeviceBrand               []string `json:"device_brand,omitempty"`                 // 手机品牌
	LaunchPrice               []int64  `json:"launch_price,omitempty"`                 // 手机价格区间
	AutoExtendTargets         []string `json:"auto_extend_targets,omitempty"`          // 可放开定向（智能放量）
	DPACity                   string   `json:"dpa_city,omitempty"`                     // 地域匹配-商品所在城市
	DPALbs                    string   `json:"dpa_lbs,omitempty"`                      // 地域匹配-适地性服务
	DPARtaSwitch              string   `json:"dpa_rta_switch,omitempty"`               // RTA重定向开关
	RtaID                     int64    `json:"rta_id,omitempty"`                       // RTA策略ID（条件必填）
	DPARtaRecommendType       string   `json:"dpa_rta_recommend_type,omitempty"`       // RTA推荐逻辑（条件必填）
}

// UBLAudience UBL人群定向（线索智投场景）
type UBLAudience struct {
	AudiencePackageID      int64          `json:"audience_package_id,omitempty"`      // 定向包ID
	District               string         `json:"district"`                           // 地域类型 (必填)
	Geolocation            []*Geolocation `json:"geolocation,omitempty"`              // 地图位置 (商圈专用)
	RegionRecommend        string         `json:"region_recommend,omitempty"`         // 地域智能放量定向
	Age                    []string       `json:"age,omitempty"`                      // 年龄
	RetargetingTagsInclude []int64        `json:"retargeting_tags_include,omitempty"` // 定向人群包列表
	RetargetingTagsExclude []int64        `json:"retargeting_tags_exclude,omitempty"` // 排除人群包列表
	HideIfConverted        string         `json:"hide_if_converted,omitempty"`        // 过滤已转化用户
}

// Geolocation 地图位置信息
type Geolocation struct {
	Radius int     `json:"radius,omitempty"` // 半径，单位m，默认6000
	Name   string  `json:"name,omitempty"`   // 地点名称
	Long   float64 `json:"long,omitempty"`   // 经度
	Lat    float64 `json:"lat,omitempty"`    // 纬度
}

// 常量定义
const (
	// ScheduleType 投放时间类型
	ScheduleTypeFromNow  = "SCHEDULE_FROM_NOW"  // 从今天起长期投放
	ScheduleTypeStartEnd = "SCHEDULE_START_END" // 设置开始和结束日期
	ScheduleType7Days    = "SCHEDULE_7_DAYS"    // 7日稳投 (new)

	// 时间格式
	DateFormat = "2006-01-02" // 精确到天
)

// DeliverySetting 排期、预算、出价
type DeliverySetting struct {
	ScheduleType           string             `json:"schedule_type"`                      // 投放时间类型 (必填)
	StartTime              string             `json:"start_time,omitempty"`               // 投放起始时间 (条件必填)
	EndTime                string             `json:"end_time,omitempty"`                 // 投放结束时间 (条件必填)
	ScheduleTime           string             `json:"schedule_time,omitempty"`            // 投放时间段
	LiveDuration           int64              `json:"live_duration,omitempty"`            // 直播时长
	FilterNightSwitch      string             `json:"filter_night_switch,omitempty"`      // 过滤夜间时段开关
	DeepBidType            string             `json:"deep_bid_type,omitempty"`            // 深度优化方式
	BidType                string             `json:"bid_type,omitempty"`                 // 竞价策略
	ProjectCustom          string             `json:"project_custom,omitempty"`           // 项目成本稳投
	Bid                    float32            `json:"bid,omitempty"`                      // 点击出价/展示出价
	BudgetMode             string             `json:"budget_mode,omitempty"`              // 项目预算类型
	Budget                 float32            `json:"budget,omitempty"`                   // 项目预算，单位为元
	Pricing                float32            `json:"pricing,omitempty"`                  // 付费方式
	CpaBid                 float32            `json:"cpa_bid,omitempty"`                  // 目标转化出价/预期成本
	DeepCpabid             float32            `json:"deep_cpabid,omitempty"`              // 深度优化出价
	RoiGoal                float32            `json:"roi_goal,omitempty"`                 // ROI系数
	FirstRoiGoal           float64            `json:"first_roi_goal,omitempty"`           // 首日roi系数
	SevenRoiGoal           float64            `json:"seven_roi_goal,omitempty"`           // 7日ROI系数
	ShopMultiRoiGoals      *ShopMultiRoiGoals `json:"shop_multi_roi_goals,omitempty"`     // 电商平台多ROI系数，指引流电商多平台投放ROI系数及平台信息，可按照电商平台分别确定ROI系数，分平台调控出价，白名单开放。
	BudgetOptimizeSwitch   string             `json:"budget_optimize_switch,omitempty"`   // 支持预算择优分配
	SearchContinueDelivery string             `json:"search_continue_delivery,omitempty"` // 续投，仅当delivery_type = DURATION搜索极速智投时必填
}

type ShopMultiRoiGoals struct {
	RoiGoal      float32 `json:"roi_goal,omitempty"`      // ROI系数
	ShopPlatform string  `json:"shop_platform,omitempty"` // ROI系数所属平台
}

type TrackUrlSetting struct {
	TrackUrlType               string   `json:"track_url_type,omitempty"`                 // 监测链接类型，区分使用监测链接组或者自定义链接
	TrackUrlGroupId            int64    `json:"track_url_group_id,omitempty"`             // 监测链接组id
	TrackUrl                   []string `json:"track_url,omitempty"`                      // 展示（监测链接）
	ActionTrackUrl             []string `json:"action_track_url,omitempty"`               // 点击（监测链接），只允许传入1个
	ActiveTrackUrl             []string `json:"active_track_url,omitempty"`               // 激活检测链接，只允许传入1个
	VideoPlayEffectiveTrackUrl []string `json:"video_play_effective_track_url,omitempty"` // 视频有效播放（监测链接）
	VideoPlayDoneTrackUrl      []string `json:"video_play_done_track_url,omitempty"`      // 视频播完（监测链接）
	VideoPlayFirstTrackUrl     []string `json:"video_play_first_track_url,omitempty"`     // 视频开始播放（监测链接）
	SendType                   string   `json:"send_type,omitempty"`                      // 数据发送方式
}

// 辅助方法：检查LandingType是否支持自动投放
func IsDeliveryModeSupportedForLandingType(landingType string) bool {
	supportedTypes := map[string]bool{
		LandingTypeApp:       true,
		LandingTypeMicroGame: true,
		LandingTypeLink:      true,
		LandingTypeShop:      true,
		LandingTypeQuickApp:  true,
	}
	return supportedTypes[landingType]
}

// 辅助方法：检查是否直播场景支持自动投放
func IsLiveAutoDeliverySupported(marketingGoal string) bool {
	return marketingGoal == "LIVE"
}

type ProjectCreateResp struct {
	ProjectID                  int64              `json:"project_id"`                             // 项目id
	SupplementaryAgreementInfo string             `json:"supplementary_agreement_info,omitempty"` // 星广联投投放协议查看地址
	AutoProjectCount           int                `json:"auto_project_count,omitempty"`           // 账户下的剩余可建自动投放项目额度数量
	ErrorKeywordsList          []ErrorKeywordInfo `json:"error_keywords_list,omitempty"`          // 极速智投场景下自定义关键词上传失败的关键词list
	AutoDurationProjectCount   int                `json:"auto_duration_project_count,omitempty"`  // 账户下剩余可建智能托管项目额度数量
}

// ErrorKeywordInfo 失败的关键词信息
type ErrorKeywordInfo struct {
	ErrorKeyword string `json:"error_keyword"` // 失败的关键词
	ErrorMessage string `json:"error_message"` // 失败原因
}

// ProjectUpdateReq 更新项目请求参数
type ProjectUpdateReq struct {
	accessTokenReq
	AdvertiserId                   int64  `json:"advertiser_id,omitempty"`                      // 投放账户id
	ProjectId                      int64  `json:"project_id,omitempty"`                         // 项目id
	Name                           string `json:"name,omitempty"`                               // 项目名称
	AigcDynamicCreativeSwitch      string `json:"aigc_dynamic_creative_switch,omitempty"`       // 是否开启AIGC动态创意
	StarAutoDeliverySwitch         string `json:"star_auto_delivery_switch,omitempty"`          // 星广联投全自动化开关
	StarAutoMaterialAdditionSwitch string `json:"star_auto_material_addition_switch,omitempty"` // 星广联投自动优化素材开关
	SearchQuick
	SearchKeywords

	//搜索蓝海流量投放
	BlueFlowPackage *BlueFlowPackage `json:"blue_flow_package,omitempty"` // 搜索蓝海流量投放

	//营销产品与投放载体
	RelatedProduct *RelatedProduct `json:"related_product,omitempty"` // 商品

	//用户定向
	Audience interface{} `json:"audience,omitempty"` // 定向

	DeliverySetting *DeliverySetting `json:"delivery_setting,omitempty"`  // 排期、预算、出价
	TrackUrlSetting *TrackUrlSetting `json:"track_url_setting,omitempty"` // 监测链接
}

func (receiver *ProjectUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *ProjectUpdateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.ProjectId <= 0 {
		err = errors.New("project_id is empty")
		return
	}
	return
}

func (receiver *ProjectUpdateReq) GetHeaders() map[string]string {
	headers := receiver.accessTokenReq.GetHeaders()
	headers.Json()
	return headers
}

type ProjectUpdateResp struct {
	ProjectId int64 `json:"project_id,omitempty"` // 项目ID
	ErrorList []struct {
		ObjectType   string `json:"object_type,omitempty"`   // 错误对象类型，返回值：BUDGET 预算、PROJECT_SETTING 项目设置
		ErrorCode    string `json:"error_code,omitempty"`    // 错误码
		ErrorMessage string `json:"error_message,omitempty"` // 错误信息
	} `json:"error_list,omitempty"` // 错误list，项目为分块更新，存在部分内容更新失败，部分内容更新成功
	ErrorKeywordsList []struct {
		ErrorKeyword string `json:"error_keyword,omitempty"` // 失败的关键词
		ErrorMessage string `json:"error_message,omitempty"` // 失败原因
	}
}
