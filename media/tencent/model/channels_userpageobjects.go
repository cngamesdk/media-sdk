package model

import "errors"

// ========== 视频号-获取视频号动态列表 ==========
// https://developers.e.qq.com/v3.0/docs/api/channels_userpageobjects/get

// count 常量
const (
	MinChannelsUserpageobjectsCount     = 0  // count 最小值
	MaxChannelsUserpageobjectsCount     = 30 // count 最大值
	DefaultChannelsUserpageobjectsCount = 10 // count 默认值
)

// ChannelsAdContextMarketingCarrierDetail 营销载体详情
type ChannelsAdContextMarketingCarrierDetail struct {
	MarketingCarrierId    string `json:"marketing_carrier_id"`               // 营销载体 id (必填)
	MarketingSubCarrierId string `json:"marketing_sub_carrier_id,omitempty"` // 二级营销载体 id
	MarketingCarrierName  string `json:"marketing_carrier_name,omitempty"`   // 营销载体名称
}

// ChannelsAdContextOptimizationGoalStruct og 组合
type ChannelsAdContextOptimizationGoalStruct struct {
	OptimizationGoal               string `json:"optimization_goal,omitempty"`                 // 优化目标类型（枚举）
	DeepOptimizationGoal           string `json:"deep_optimization_goal,omitempty"`            // 深度优化目标类型（枚举）
	DeepConversionOptimizationGoal string `json:"deep_conversion_optimization_goal,omitempty"` // ROI 目标（枚举）
}

// ChannelsAdContextMpaSpec 动态商品广告属性
type ChannelsAdContextMpaSpec struct {
	RecommendMethodIds []int64 `json:"recommend_method_ids,omitempty"` // 商品推荐方式，1-16个
	ProductCatalogId   string  `json:"product_catalog_id,omitempty"`   // 商品库 id
	ProductSeriesId    string  `json:"product_series_id,omitempty"`    // 商品集合 id
}

// ChannelsAdContextMarketingAssetOuterSpec 产品外部 id 数据
type ChannelsAdContextMarketingAssetOuterSpec struct {
	MarketingTargetType      string `json:"marketing_target_type,omitempty"`        // 推广产品类型（枚举）
	MarketingAssetOuterId    string `json:"marketing_asset_outer_id,omitempty"`     // 推广产品外部 id，1-1024字节
	MarketingAssetOuterSubId string `json:"marketing_asset_outer_sub_id,omitempty"` // 推广产品外部子 id，1-1024字节
	MarketingAssetOuterName  string `json:"marketing_asset_outer_name,omitempty"`   // 推广产品外部名称，1-1024字节
}

// ChannelsAdContext 广告上下文信息
type ChannelsAdContext struct {
	MarketingGoal           string                                    `json:"marketing_goal"`                       // 营销目的类型（枚举，必填）
	MarketingSubGoal        string                                    `json:"marketing_sub_goal,omitempty"`         // 二级营销目的类型（枚举）
	MarketingCarrierType    string                                    `json:"marketing_carrier_type"`               // 营销载体类型（枚举，必填）
	MarketingTargetType     string                                    `json:"marketing_target_type"`                // 推广产品类型（枚举，必填）
	MarketingCarrierDetail  *ChannelsAdContextMarketingCarrierDetail  `json:"marketing_carrier_detail,omitempty"`   // 营销载体详情
	MarketingAssetId        int64                                     `json:"marketing_asset_id,omitempty"`         // 产品 id
	SiteSet                 []string                                  `json:"site_set"`                             // 投放站点集合（枚举[]，必填），1-32个
	CreativeTemplateId      int64                                     `json:"creative_template_id"`                 // 创意形式 id（必填）
	PromotedAssetType       string                                    `json:"promoted_asset_type,omitempty"`        // 推广内容类型（枚举）
	ComponentType           string                                    `json:"component_type,omitempty"`             // 创意组件类型（枚举）
	OptimizationGoalStruct  *ChannelsAdContextOptimizationGoalStruct  `json:"optimization_goal_struct,omitempty"`   // og 组合
	MpaSpec                 *ChannelsAdContextMpaSpec                 `json:"mpa_spec,omitempty"`                   // 动态商品广告属性
	MarketingAssetOuterSpec *ChannelsAdContextMarketingAssetOuterSpec `json:"marketing_asset_outer_spec,omitempty"` // 产品外部 id 数据
	DynamicAdType           string                                    `json:"dynamic_ad_type,omitempty"`            // 动态广告类型（枚举）
	AdgroupType             string                                    `json:"adgroup_type,omitempty"`               // 广告类型（枚举）
}

// ChannelsUserpageobjectsGetReq 获取视频号动态列表请求（GET）
// https://developers.e.qq.com/v3.0/docs/api/channels_userpageobjects/get
type ChannelsUserpageobjectsGetReq struct {
	GlobalReq
	AccountId               int64              `json:"account_id"`                           // 广告主帐号 id (必填)，不支持代理商 id
	FinderUsername          string             `json:"finder_username,omitempty"`            // 视频号账号 id（已废弃），1-1024字节
	Nickname                string             `json:"nickname,omitempty"`                   // 视频号名称，1-1024字节
	LastBuffer              string             `json:"last_buffer,omitempty"`                // 上次返回的 buffer，1-3145728字节
	Count                   int                `json:"count,omitempty"`                      // 数据条数，0-30，默认10
	WechatChannelsAccountId string             `json:"wechat_channels_account_id,omitempty"` // 视频号账号 id，1-1024字节
	AdContext               *ChannelsAdContext `json:"ad_context,omitempty"`                 // 广告上下文信息
}

func (r *ChannelsUserpageobjectsGetReq) Format() {
	r.GlobalReq.Format()
	if r.Count == 0 {
		r.Count = DefaultChannelsUserpageobjectsCount
	}
}

// Validate 验证获取视频号动态列表请求参数
func (r *ChannelsUserpageobjectsGetReq) Validate() error {
	if r.AccountId == 0 {
		return errors.New("account_id为必填")
	}
	if r.Count < MinChannelsUserpageobjectsCount || r.Count > MaxChannelsUserpageobjectsCount {
		return errors.New("count须在0-30之间")
	}
	if r.AdContext != nil {
		if r.AdContext.MarketingGoal == "" {
			return errors.New("ad_context.marketing_goal为必填")
		}
		if r.AdContext.MarketingCarrierType == "" {
			return errors.New("ad_context.marketing_carrier_type为必填")
		}
		if r.AdContext.MarketingTargetType == "" {
			return errors.New("ad_context.marketing_target_type为必填")
		}
		if len(r.AdContext.SiteSet) == 0 {
			return errors.New("ad_context.site_set为必填")
		}
		if r.AdContext.CreativeTemplateId == 0 {
			return errors.New("ad_context.creative_template_id为必填")
		}
	}
	return r.GlobalReq.Validate()
}

// ChannelsUserpageobjectsMedia 媒体信息
type ChannelsUserpageobjectsMedia struct {
	ThumbUrl     string  `json:"thumb_url"`      // 首帧图片 URL，地址非长期有效
	CoverUrl     string  `json:"cover_url"`      // 封面图 URL，地址非长期有效
	MediaType    int     `json:"media_type"`     // 动态资源类型
	Width        float64 `json:"width"`          // 宽度
	Height       float64 `json:"height"`         // 高度
	Url          string  `json:"url"`            // 视频 URL，地址非长期有效
	VideoPlayLen int     `json:"video_play_len"` // 视频时长
}

// ChannelsUserpageobjectsComponent 视频号动态组件
type ChannelsUserpageobjectsComponent struct {
	ComponentType string `json:"component_type"` // 视频号原生动态组件类型（枚举）
	Icon          string `json:"icon"`           // 视频号动态组件图标
	Wording       string `json:"wording"`        // 视频号动态组件文案
	ComponentUuid string `json:"component_uuid"` // 视频号动态组件 id
}

// ChannelsUserpageobjectsItem 视频号动态列表项
type ChannelsUserpageobjectsItem struct {
	ExportId                 string                              `json:"export_id"`                  // 动态 id
	CreateTime               int64                               `json:"create_time"`                // 创建时间
	DeleteFlag               int                                 `json:"delete_flag"`                // 删除标记，0 表示未删除
	Description              string                              `json:"description"`                // 动态标题
	Medias                   []*ChannelsUserpageobjectsMedia     `json:"medias"`                     // 媒体信息列表
	FinderUsername           string                              `json:"finder_username"`            // 视频号账号 id（已废弃）
	WechatChannelsAccountId  string                              `json:"wechat_channels_account_id"` // 视频号账号 id
	CreatedSource            string                              `json:"created_source"`             // 视频号创建来源（枚举）
	WechatChannelsComponents []*ChannelsUserpageobjectsComponent `json:"wechat_channels_components"` // 视频号动态组件列表
	FeedsSourceType          string                              `json:"feeds_source_type"`          // 视频号动态类型（枚举）
	IsDisable                bool                                `json:"is_disable"`                 // 是否可使用
	DisableMessage           string                              `json:"disable_message"`            // 禁用原因
	AuditStatus              string                              `json:"audit_status"`               // 视频号动态审核状态（枚举）
}

// ChannelsUserpageobjectsGetResp 获取视频号动态列表响应
type ChannelsUserpageobjectsGetResp struct {
	Objects      []*ChannelsUserpageobjectsItem `json:"objects"`       // 动态列表
	LastBuffer   string                         `json:"last_buffer"`   // 上次返回的 buffer，用于连续翻页
	ContinueFlag int                            `json:"continue_flag"` // 1 表示还有数据，带上 last_buffer 接着拉取
}
