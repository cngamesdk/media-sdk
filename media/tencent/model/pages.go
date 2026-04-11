package model

import "errors"

// ========== 获取落地页列表 ==========
// https://developers.e.qq.com/v3.0/docs/api/pages/get

// 落地页类型枚举（其余类型见 creative.go）
const (
	PageTypeOfficial = "PAGE_TYPE_OFFICIAL" // 官方落地页
)

// 落地页状态枚举
const (
	PageStatusNormal  = "NORMAL"  // 正常
	PageStatusDeleted = "DELETED" // 已删除
	PageStatusPending = "PENDING" // 待审核
)

// 落地页选择场景枚举
const (
	PageSelectorSceneDefault = "PAGE_SELECTOR_SCENE_DEFAULT" // 默认场景
	PageSelectorSceneLiving  = "PAGE_SELECTOR_SCENE_LIVING"  // 直播场景
)

// 过滤字段常量
const (
	PagesGetFilterFieldPageType          = "page_type"           // 落地页类型
	PagesGetFilterFieldPageID            = "page_id"             // 落地页 id
	PagesGetFilterFieldPageName          = "page_name"           // 落地页名称
	PagesGetFilterFieldPageStatus        = "page_status"         // 落地页状态
	PagesGetFilterFieldOwnerAccountID    = "owner_account_id"    // 落地页所属账户 id
	PagesGetFilterFieldPageSelectorScene = "page_selector_scene" // 落地页选择场景
)

// 过滤操作符常量
const (
	PagesGetFilterOperatorEquals   = "EQUALS"   // 精确匹配
	PagesGetFilterOperatorIn       = "IN"       // 包含匹配
	PagesGetFilterOperatorContains = "CONTAINS" // 模糊匹配
)

// 分页常量
const (
	MinPagesGetPage         = 1     // page 最小值
	MaxPagesGetPage         = 99999 // page 最大值
	MinPagesGetPageSize     = 1     // page_size 最小值
	MaxPagesGetPageSize     = 100   // page_size 最大值
	DefaultPagesGetPage     = 1     // page 默认值
	DefaultPagesGetPageSize = 10    // page_size 默认值

	MaxPagesGetFilteringCount = 10  // filtering 最大长度
	MaxPagesGetPageNameBytes  = 180 // page_name 值最大字节数
)

// ad_context 相关枚举常量
const (
	// 动态广告类型
	DynamicAdTypeDynamicContent = "DYNAMIC_AD_TYPE_DYNAMIC_CONTENT" // DCA 动态内容广告

	// 广告类型
	AdgroupTypeSearch = "ADGROUP_TYPE_SEARCH" // 搜索广告
	AdgroupTypeNormal = "ADGROUP_TYPE_NORMAL" // 普通广告
)

// ad_context 字段长度常量
const (
	MaxPagesGetAdContextCarrierIDBytes        = 2048 // marketing_carrier_id 最大长度
	MaxPagesGetAdContextAssetOuterIDBytes     = 1024 // marketing_asset_outer_id 最大长度
	MaxPagesGetAdContextAssetOuterSubIDBytes  = 1024 // marketing_asset_outer_sub_id 最大长度
	MaxPagesGetAdContextAssetOuterNameBytes   = 1024 // marketing_asset_outer_name 最大长度
	MinPagesGetAdContextSiteSetLen            = 1    // site_set 最小长度
	MaxPagesGetAdContextSiteSetLen            = 32   // site_set 最大长度
	MinPagesGetAdContextRecommendMethodIDsLen = 1    // recommend_method_ids 最小长度
	MaxPagesGetAdContextRecommendMethodIDsLen = 16   // recommend_method_ids 最大长度
)

// PagesGetFilteringItem 落地页过滤条件
type PagesGetFilteringItem struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)
	Operator string   `json:"operator"` // 操作符 (必填)
	Values   []string `json:"values"`   // 字段取值 (必填)
}

// Validate 验证单个过滤条件
func (f *PagesGetFilteringItem) Validate() error {
	if f.Field == "" {
		return errors.New("field为必填")
	}
	if f.Operator == "" {
		return errors.New("operator为必填")
	}
	if len(f.Values) == 0 {
		return errors.New("values为必填，至少包含1个值")
	}
	if f.Field == PagesGetFilterFieldPageName {
		if len(f.Values[0]) == 0 || len(f.Values[0]) > MaxPagesGetPageNameBytes {
			return errors.New("page_name过滤值长度须在1-180字节之间")
		}
	}
	return nil
}

// PagesGetAdContextCarrierDetail 营销载体详情
type PagesGetAdContextCarrierDetail struct {
	MarketingCarrierID    string `json:"marketing_carrier_id,omitempty"`     // 营销载体 id (必填)，0-2048字节
	MarketingSubCarrierID string `json:"marketing_sub_carrier_id,omitempty"` // 二级营销载体 id
	MarketingCarrierName  string `json:"marketing_carrier_name,omitempty"`   // 营销载体名称
}

// PagesGetAdContextOptimizationGoalStruct 优化目标组合
type PagesGetAdContextOptimizationGoalStruct struct {
	OptimizationGoal               string `json:"optimization_goal,omitempty"`                 // 优化目标类型
	DeepOptimizationGoal           string `json:"deep_optimization_goal,omitempty"`            // 深度优化目标类型
	DeepConversionOptimizationGoal string `json:"deep_conversion_optimization_goal,omitempty"` // ROI 目标
}

// PagesGetAdContextMpaSpec 动态商品广告属性
type PagesGetAdContextMpaSpec struct {
	RecommendMethodIDs []int64 `json:"recommend_method_ids,omitempty"` // 商品推荐方式 id 列表，1-16个
	ProductCatalogID   string  `json:"product_catalog_id,omitempty"`   // 商品库 id
	ProductSeriesID    string  `json:"product_series_id,omitempty"`    // 商品集合 id
}

// PagesGetAdContextMarketingAssetOuterSpec 产品外部 id 数据
type PagesGetAdContextMarketingAssetOuterSpec struct {
	MarketingTargetType      string `json:"marketing_target_type,omitempty"`        // 推广产品类型
	MarketingAssetOuterID    string `json:"marketing_asset_outer_id,omitempty"`     // 推广产品外部 id，1-1024字节
	MarketingAssetOuterSubID string `json:"marketing_asset_outer_sub_id,omitempty"` // 推广产品外部子 id，1-1024字节
	MarketingAssetOuterName  string `json:"marketing_asset_outer_name,omitempty"`   // 推广产品外部名称，1-1024字节
}

// PagesGetAdContext 广告上下文信息
type PagesGetAdContext struct {
	MarketingGoal           string                                    `json:"marketing_goal"`                       // 营销目的类型 (必填)
	MarketingSubGoal        string                                    `json:"marketing_sub_goal,omitempty"`         // 二级营销目的类型
	MarketingCarrierType    string                                    `json:"marketing_carrier_type"`               // 营销载体类型 (必填)
	MarketingTargetType     string                                    `json:"marketing_target_type"`                // 推广产品类型 (必填)
	MarketingCarrierDetail  *PagesGetAdContextCarrierDetail           `json:"marketing_carrier_detail,omitempty"`   // 营销载体详情
	MarketingAssetID        int64                                     `json:"marketing_asset_id,omitempty"`         // 产品 id
	SiteSet                 []string                                  `json:"site_set"`                             // 投放站点集合 (必填)，1-32个
	CreativeTemplateID      int64                                     `json:"creative_template_id"`                 // 创意形式 id (必填)
	PromotedAssetType       string                                    `json:"promoted_asset_type,omitempty"`        // 推广内容类型
	ComponentType           string                                    `json:"component_type,omitempty"`             // 创意组件类型
	OptimizationGoalStruct  *PagesGetAdContextOptimizationGoalStruct  `json:"optimization_goal_struct,omitempty"`   // 优化目标组合
	MpaSpec                 *PagesGetAdContextMpaSpec                 `json:"mpa_spec,omitempty"`                   // 动态商品广告属性
	MarketingAssetOuterSpec *PagesGetAdContextMarketingAssetOuterSpec `json:"marketing_asset_outer_spec,omitempty"` // 产品外部 id 数据
	DynamicAdType           string                                    `json:"dynamic_ad_type,omitempty"`            // 动态广告类型
	AdgroupType             string                                    `json:"adgroup_type,omitempty"`               // 广告类型
}

// Validate 验证广告上下文信息
func (a *PagesGetAdContext) Validate() error {
	if a.MarketingGoal == "" {
		return errors.New("ad_context.marketing_goal为必填")
	}
	if a.MarketingCarrierType == "" {
		return errors.New("ad_context.marketing_carrier_type为必填")
	}
	if a.MarketingTargetType == "" {
		return errors.New("ad_context.marketing_target_type为必填")
	}
	if len(a.SiteSet) < MinPagesGetAdContextSiteSetLen || len(a.SiteSet) > MaxPagesGetAdContextSiteSetLen {
		return errors.New("ad_context.site_set数组长度须在1-32之间")
	}
	if a.CreativeTemplateID == 0 {
		return errors.New("ad_context.creative_template_id为必填")
	}
	if a.MarketingCarrierDetail != nil {
		if len(a.MarketingCarrierDetail.MarketingCarrierID) > MaxPagesGetAdContextCarrierIDBytes {
			return errors.New("ad_context.marketing_carrier_detail.marketing_carrier_id长度不能超过2048字节")
		}
	}
	if a.MpaSpec != nil {
		if len(a.MpaSpec.RecommendMethodIDs) > MaxPagesGetAdContextRecommendMethodIDsLen {
			return errors.New("ad_context.mpa_spec.recommend_method_ids数组长度不能超过16")
		}
	}
	if a.MarketingAssetOuterSpec != nil {
		if len(a.MarketingAssetOuterSpec.MarketingAssetOuterID) > MaxPagesGetAdContextAssetOuterIDBytes {
			return errors.New("ad_context.marketing_asset_outer_spec.marketing_asset_outer_id长度不能超过1024字节")
		}
		if len(a.MarketingAssetOuterSpec.MarketingAssetOuterSubID) > MaxPagesGetAdContextAssetOuterSubIDBytes {
			return errors.New("ad_context.marketing_asset_outer_spec.marketing_asset_outer_sub_id长度不能超过1024字节")
		}
		if len(a.MarketingAssetOuterSpec.MarketingAssetOuterName) > MaxPagesGetAdContextAssetOuterNameBytes {
			return errors.New("ad_context.marketing_asset_outer_spec.marketing_asset_outer_name长度不能超过1024字节")
		}
	}
	return nil
}

// PagesGetReq 获取落地页列表请求（GET）
// https://developers.e.qq.com/v3.0/docs/api/pages/get
type PagesGetReq struct {
	GlobalReq
	AccountID int64                    `json:"account_id"`           // 广告主帐号 id (必填)
	Filtering []*PagesGetFilteringItem `json:"filtering,omitempty"`  // 过滤条件，最大10条
	Page      int                      `json:"page,omitempty"`       // 搜索页码，1-99999，默认1
	PageSize  int                      `json:"page_size,omitempty"`  // 每页条数，1-100，默认10
	AdContext *PagesGetAdContext       `json:"ad_context,omitempty"` // 广告上下文信息
}

func (p *PagesGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page == 0 {
		p.Page = DefaultPagesGetPage
	}
	if p.PageSize == 0 {
		p.PageSize = DefaultPagesGetPageSize
	}
}

// Validate 验证获取落地页列表请求参数
func (p *PagesGetReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if len(p.Filtering) > MaxPagesGetFilteringCount {
		return errors.New("filtering数组长度不能超过10")
	}
	for i, f := range p.Filtering {
		if f == nil {
			return errors.New("filtering[" + itoa(i) + "]不能为空")
		}
		if err := f.Validate(); err != nil {
			return errors.New("filtering[" + itoa(i) + "]: " + err.Error())
		}
	}
	if p.Page < MinPagesGetPage || p.Page > MaxPagesGetPage {
		return errors.New("page须在1-99999之间")
	}
	if p.PageSize < MinPagesGetPageSize || p.PageSize > MaxPagesGetPageSize {
		return errors.New("page_size须在1-100之间")
	}
	if p.AdContext != nil {
		if err := p.AdContext.Validate(); err != nil {
			return err
		}
	}
	return p.GlobalReq.Validate()
}

// PagesGetItem 落地页列表项
type PagesGetItem struct {
	PageType         string `json:"page_type"`          // 落地页类型
	PageID           int64  `json:"page_id"`            // 落地页 id
	PageName         string `json:"page_name"`          // 落地页名称
	PageURL          string `json:"page_url"`           // 落地页 url
	PageStatus       string `json:"page_status"`        // 落地页状态
	OwnerAccountID   int64  `json:"owner_account_id"`   // 所属账户 id
	CreatedTime      int64  `json:"created_time"`       // 创建时间，时间戳
	LastModifiedTime int64  `json:"last_modified_time"` // 最后修改时间，时间戳
	DisableCode      int64  `json:"disable_code"`       // 不可用错误码
	DisableMessage   string `json:"disable_message"`    // 不可用错误信息
}

// PagesGetResp 获取落地页列表响应
// https://developers.e.qq.com/v3.0/docs/api/pages/get
type PagesGetResp struct {
	List     []*PagesGetItem `json:"list"`      // 返回信息列表
	PageInfo *PageInfo       `json:"page_info"` // 分页配置信息
}
