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
