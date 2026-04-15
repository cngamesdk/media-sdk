package model

import (
	"errors"
)

// ========== 获取日报表 ==========
// https://developers.e.qq.com/v3.0/docs/api/daily_reports/get

// DailyReportsGetReq 获取日报表请求
type DailyReportsGetReq struct {
	GlobalReq
	AccountID      int64                 `json:"account_id,omitempty"`      // 广告主帐号id，有操作权限的帐号id，不支持代理商id
	Level          string                `json:"level"`                     // 获取报表类型级别 (必填)
	DateRange      *DailyReportDateRange `json:"date_range"`                // 日期范围 (必填)
	Filtering      []*DailyReportFilter  `json:"filtering,omitempty"`       // 过滤条件
	GroupBy        []string              `json:"group_by"`                  // 聚合参数 (必填)
	OrderBy        []*DailyReportOrderBy `json:"order_by,omitempty"`        // 排序字段
	TimeLine       string                `json:"time_line,omitempty"`       // 时间口径
	Page           int                   `json:"page,omitempty"`            // 搜索页码，默认值：1
	PageSize       int                   `json:"page_size,omitempty"`       // 一页显示的数据条数，默认值：10
	Fields         []string              `json:"fields"`                    // 指定返回的字段列表 (必填)
	OrganizationID int64                 `json:"organization_id,omitempty"` // 业务单元id
}

// DailyReportDateRange 日期范围
type DailyReportDateRange struct {
	StartDate string `json:"start_date"` // 开始日期，格式：YYYY-MM-DD (必填)
	EndDate   string `json:"end_date"`   // 结束日期，格式：YYYY-MM-DD (必填)
}

// DailyReportFilter 过滤条件
type DailyReportFilter struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)
	Operator string   `json:"operator"` // 操作符 (必填)
	Values   []string `json:"values"`   // 字段取值 (必填)
}

// DailyReportOrderBy 排序字段
type DailyReportOrderBy struct {
	SortField string `json:"sort_field"` // 排序字段 (必填)
	SortType  string `json:"sort_type"`  // 排序方式 (必填)
}

// 报表级别常量
const (
	ReportLevelAdvertiser              = "REPORT_LEVEL_ADVERTISER"
	ReportLevelAdgroup                 = "REPORT_LEVEL_ADGROUP"
	ReportLevelDynamicCreative         = "REPORT_LEVEL_DYNAMIC_CREATIVE"
	ReportLevelComponent               = "REPORT_LEVEL_COMPONENT"
	ReportLevelChannel                 = "REPORT_LEVEL_CHANNEL"
	ReportLevelBidword                 = "REPORT_LEVEL_BIDWORD"
	ReportLevelQueryword               = "REPORT_LEVEL_QUERYWORD"
	ReportLevelMaterialImage           = "REPORT_LEVEL_MATERIAL_IMAGE"
	ReportLevelMaterialVideo           = "REPORT_LEVEL_MATERIAL_VIDEO"
	ReportLevelMarketingAsset          = "REPORT_LEVEL_MARKETING_ASSET"
	ReportLevelProductCatalog          = "REPORT_LEVEL_PRODUCT_CATALOG"
	ReportLevelProject                 = "REPORT_LEVEL_PROJECT"
	ReportLevelProjectCreative         = "REPORT_LEVEL_PROJECT_CREATIVE"
	ReportLevelVideoHighlight          = "REPORT_LEVEL_VIDEO_HIGHLIGHT"
	ReportLevelProductCreativeTemplate = "REPORT_LEVEL_PRODUCT_CREATIVE_TEMPLATE"
	ReportLevelWechatShopProduct       = "REPORT_LEVEL_WECHAT_SHOP_PRODUCT"
)

// 时间口径常量
const (
	TimeLineRequestTime   = "REQUEST_TIME"
	TimeLineReportingTime = "REPORTING_TIME"
	TimeLineActiveTime    = "ACTIVE_TIME"
)

// 排序方式常量
const (
	SortTypeAscending  = "ASCENDING"
	SortTypeDescending = "DESCENDING"
)

// 过滤字段常量
const (
	DailyReportFilterAdgroupId               = "adgroup_id"
	DailyReportFilterDynamicCreativeId       = "dynamic_creative_id"
	DailyReportFilterComponentId             = "component_id"
	DailyReportFilterComponentType           = "component_type"
	DailyReportFilterBidwordId               = "bidword_id"
	DailyReportFilterChannelId               = "channel_id"
	DailyReportFilterImageId                 = "image_id"
	DailyReportFilterVideoId                 = "video_id"
	DailyReportFilterMarketingTargetType     = "marketing_target_type"
	DailyReportFilterMarketingAssetId        = "marketing_asset_id"
	DailyReportFilterSmartDeliveryPlatform   = "smart_delivery_platform"
	DailyReportFilterMd5                     = "md5"
	DailyReportFilterProductCatalogId        = "product_catalog_id"
	DailyReportFilterProductSeriesId         = "product_series_id"
	DailyReportFilterProductOuterId          = "product_outer_id"
	DailyReportFilterCreativeTemplateGroupId = "creative_template_group_id"
)

// 日报表限制常量
const (
	DailyReportMinFilteringCount = 1
	DailyReportMaxFilteringCount = 40
	DailyReportMinValuesCount    = 1
	DailyReportMaxValuesCount    = 100
	DailyReportMaxValuesLength   = 64
	DailyReportMinGroupByCount   = 1
	DailyReportMaxGroupByCount   = 10
	DailyReportMaxGroupByLength  = 255
	DailyReportMinOrderByCount   = 1
	DailyReportMaxOrderByCount   = 2
	DailyReportMinPage           = 1
	DailyReportMaxPage           = 99999
	DailyReportMinPageSize       = 1
	DailyReportMaxPageSize       = 2000
	DailyReportDefaultPage       = 1
	DailyReportDefaultPageSize   = 10
	DailyReportMinFieldsCount    = 1
	DailyReportMaxFieldsCount    = 1024
	DailyReportMinFieldLength    = 1
	DailyReportMaxFieldLength    = 64
	DailyReportDateLength        = 10
	DailyReportMaxDataLimit      = 20000 // page * pageSize 最大值
)

func (p *DailyReportsGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page <= 0 {
		p.Page = DailyReportDefaultPage
	}
	if p.PageSize <= 0 {
		p.PageSize = DailyReportDefaultPageSize
	}
}

func (p *DailyReportsGetReq) Validate() error {
	// 验证全局参数
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	// 验证level (必填)
	if p.Level == "" {
		return errors.New("level为必填")
	}
	if !isValidReportLevel(p.Level) {
		return errors.New("level值无效")
	}

	// 验证date_range (必填)
	if p.DateRange == nil {
		return errors.New("date_range为必填")
	}
	if p.DateRange.StartDate == "" {
		return errors.New("date_range.start_date为必填")
	}
	if len(p.DateRange.StartDate) != DailyReportDateLength {
		return errors.New("date_range.start_date长度必须为10字节")
	}
	if p.DateRange.EndDate == "" {
		return errors.New("date_range.end_date为必填")
	}
	if len(p.DateRange.EndDate) != DailyReportDateLength {
		return errors.New("date_range.end_date长度必须为10字节")
	}
	if p.DateRange.StartDate > p.DateRange.EndDate {
		return errors.New("date_range.start_date必须小于等于end_date")
	}

	// 验证filtering
	if len(p.Filtering) > 0 {
		if len(p.Filtering) < DailyReportMinFilteringCount || len(p.Filtering) > DailyReportMaxFilteringCount {
			return errors.New("filtering数组长度必须在1-40之间")
		}
		for _, f := range p.Filtering {
			if f.Field == "" {
				return errors.New("filtering.field为必填")
			}
			if f.Operator == "" {
				return errors.New("filtering.operator为必填")
			}
			if len(f.Values) < DailyReportMinValuesCount || len(f.Values) > DailyReportMaxValuesCount {
				return errors.New("filtering.values数组长度必须在1-100之间")
			}
			for _, v := range f.Values {
				if len(v) > DailyReportMaxValuesLength {
					return errors.New("filtering.values字段长度不能超过64字节")
				}
			}
		}
	}

	// 验证group_by (必填)
	if len(p.GroupBy) < DailyReportMinGroupByCount || len(p.GroupBy) > DailyReportMaxGroupByCount {
		return errors.New("group_by数组长度必须在1-10之间")
	}
	for _, g := range p.GroupBy {
		if len(g) > DailyReportMaxGroupByLength {
			return errors.New("group_by字段长度不能超过255字节")
		}
	}

	// 验证order_by
	if len(p.OrderBy) > 0 {
		if len(p.OrderBy) < DailyReportMinOrderByCount || len(p.OrderBy) > DailyReportMaxOrderByCount {
			return errors.New("order_by数组长度必须在1-2之间")
		}
		for _, o := range p.OrderBy {
			if o.SortField == "" {
				return errors.New("order_by.sort_field为必填")
			}
			if o.SortType == "" {
				return errors.New("order_by.sort_type为必填")
			}
			if o.SortType != SortTypeAscending && o.SortType != SortTypeDescending {
				return errors.New("order_by.sort_type值无效，允许值：ASCENDING、DESCENDING")
			}
		}
	}

	// 验证time_line
	if p.TimeLine != "" {
		if p.TimeLine != TimeLineRequestTime && p.TimeLine != TimeLineReportingTime && p.TimeLine != TimeLineActiveTime {
			return errors.New("time_line值无效，允许值：REQUEST_TIME、REPORTING_TIME、ACTIVE_TIME")
		}
	}

	// 验证page
	if p.Page < DailyReportMinPage || p.Page > DailyReportMaxPage {
		return errors.New("page必须在1-99999之间")
	}

	// 验证page_size
	if p.PageSize < DailyReportMinPageSize || p.PageSize > DailyReportMaxPageSize {
		return errors.New("page_size必须在1-2000之间")
	}

	// 验证page * pageSize <= 20000
	if p.Page*p.PageSize > DailyReportMaxDataLimit {
		return errors.New("page * page_size必须小于等于20000")
	}

	// 验证fields (必填)
	if len(p.Fields) < DailyReportMinFieldsCount || len(p.Fields) > DailyReportMaxFieldsCount {
		return errors.New("fields数组长度必须在1-1024之间")
	}
	for _, f := range p.Fields {
		if len(f) < DailyReportMinFieldLength || len(f) > DailyReportMaxFieldLength {
			return errors.New("fields字段长度必须在1-64字节之间")
		}
	}

	return nil
}

// isValidReportLevel 验证报表级别
func isValidReportLevel(level string) bool {
	validLevels := map[string]bool{
		ReportLevelAdvertiser:              true,
		ReportLevelAdgroup:                 true,
		ReportLevelDynamicCreative:         true,
		ReportLevelComponent:               true,
		ReportLevelChannel:                 true,
		ReportLevelBidword:                 true,
		ReportLevelQueryword:               true,
		ReportLevelMaterialImage:           true,
		ReportLevelMaterialVideo:           true,
		ReportLevelMarketingAsset:          true,
		ReportLevelProductCatalog:          true,
		ReportLevelProject:                 true,
		ReportLevelProjectCreative:         true,
		ReportLevelVideoHighlight:          true,
		ReportLevelProductCreativeTemplate: true,
		ReportLevelWechatShopProduct:       true,
	}
	return validLevels[level]
}

// DailyReportsGetResp 获取日报表响应
type DailyReportsGetResp struct {
	List     []*DailyReportListItem `json:"list,omitempty"`
	PageInfo *PageInfo              `json:"page_info,omitempty"`
}

// DailyReportListItem 日报表列表项
type DailyReportListItem struct {
	// 基础维度字段
	Date      string `json:"date,omitempty"`       // 日期
	AccountID int64  `json:"account_id,omitempty"` // 账号ID

	// ========== 曝光与点击 ==========
	ViewCount          int64   `json:"view_count,omitempty"`           // 曝光次数
	ViewUserCount      int64   `json:"view_user_count,omitempty"`      // 曝光人数
	AvgViewPerUser     float64 `json:"avg_view_per_user,omitempty"`    // 人均曝光次数
	ValidClickCount    int64   `json:"valid_click_count,omitempty"`    // 点击次数
	ClickUserCount     int64   `json:"click_user_count,omitempty"`     // 点击人数
	Cpc                int64   `json:"cpc,omitempty"`                  // 点击均价
	Ctr                float64 `json:"ctr,omitempty"`                  // 点击率
	ValuableClickCount int64   `json:"valuable_click_count,omitempty"` // 可转化点击次数
	ValuableClickCost  int64   `json:"valuable_click_cost,omitempty"`  // 可转化点击成本
	ValuableClickRate  float64 `json:"valuable_click_rate,omitempty"`  // 可转化点击率

	// ========== 花费与成本 ==========
	Cost                 int64 `json:"cost,omitempty"`                   // 花费
	AcquisitionCost      int64 `json:"acquisition_cost,omitempty"`       // 一键起量消耗
	ThousandDisplayPrice int64 `json:"thousand_display_price,omitempty"` // 千次展现均价

	// ========== 转化 ==========
	ConversionsCount            int64   `json:"conversions_count,omitempty"`              // 目标转化量
	ConversionsRate             float64 `json:"conversions_rate,omitempty"`               // 目标转化率
	ConversionsCost             int64   `json:"conversions_cost,omitempty"`               // 目标转化成本
	DeepConversionsCount        int64   `json:"deep_conversions_count,omitempty"`         // 深度目标转化量
	DeepConversionsRate         float64 `json:"deep_conversions_rate,omitempty"`          // 深度目标转化率
	DeepConversionsCost         int64   `json:"deep_conversions_cost,omitempty"`          // 深度转化成本
	ConversionsByDisplayCount   int64   `json:"conversions_by_display_count,omitempty"`   // 目标转化量（曝光归因）
	ConversionsByDisplayRate    float64 `json:"conversions_by_display_rate,omitempty"`    // 目标转化率（曝光归因）
	ConversionsByDisplayCost    int64   `json:"conversions_by_display_cost,omitempty"`    // 目标转化成本（曝光归因）
	ConversionsByClickCount     int64   `json:"conversions_by_click_count,omitempty"`     // 目标转化量（点击归因）
	ConversionsByClickRate      float64 `json:"conversions_by_click_rate,omitempty"`      // 目标转化率（点击归因）
	ConversionsByClickCost      int64   `json:"conversions_by_click_cost,omitempty"`      // 目标转化成本（点击归因）
	PreviewConversionsCount     int64   `json:"preview_conversions_count,omitempty"`      // 目标转化量（在线预览）
	PreviewDeepConversionsCount int64   `json:"preview_deep_conversions_count,omitempty"` // 深度目标转化量（在线预览）

	// ========== 视频播放 ==========
	VideoOuterPlayCount       int64   `json:"video_outer_play_count,omitempty"`         // 视频有效播放次数
	VideoOuterPlayUserCount   int64   `json:"video_outer_play_user_count,omitempty"`    // 视频有效播放人数
	AvgUserPlayCount          float64 `json:"avg_user_play_count,omitempty"`            // 人均播放次数
	VideoOuterPlayTimeCount   float64 `json:"video_outer_play_time_count,omitempty"`    // 平均有效播放时长
	VideoOuterPlayTimeAvgRate float64 `json:"video_outer_play_time_avg_rate,omitempty"` // 平均有效播放进度
	VideoOuterPlayRate        float64 `json:"video_outer_play_rate,omitempty"`          // 有效播放率
	VideoOuterPlayCost        int64   `json:"video_outer_play_cost,omitempty"`          // 有效播放成本
	VideoOuterPlay10Count     int64   `json:"video_outer_play10_count,omitempty"`       // 10%进度播放次数
	VideoOuterPlay25Count     int64   `json:"video_outer_play25_count,omitempty"`       // 25%进度播放次数
	VideoOuterPlay50Count     int64   `json:"video_outer_play50_count,omitempty"`       // 50%进度播放次数
	VideoOuterPlay75Count     int64   `json:"video_outer_play75_count,omitempty"`       // 75%进度播放次数
	VideoOuterPlay90Count     int64   `json:"video_outer_play90_count,omitempty"`       // 90%进度播放次数
	VideoOuterPlay95Count     int64   `json:"video_outer_play95_count,omitempty"`       // 95%进度播放次数
	VideoOuterPlay100Count    int64   `json:"video_outer_play100_count,omitempty"`      // 100%进度播放次数
	VideoOuterPlay3sCount     int64   `json:"video_outer_play3s_count,omitempty"`       // 3s播放完成次数
	VideoOuterPlay3sRate      float64 `json:"video_outer_play3s_rate,omitempty"`        // 3s播放完成率
	VideoOuterPlay5sCount     int64   `json:"video_outer_play5s_count,omitempty"`       // 5s播放完成次数
	VideoOuterPlay5sRate      float64 `json:"video_outer_play5s_rate,omitempty"`        // 5s播放完成率
	VideoOuterPlay7sCount     int64   `json:"video_outer_play7s_count,omitempty"`       // 7s播放完成次数

	// ========== 互动 ==========
	ReadCount             int64 `json:"read_count,omitempty"`                // 阅读次数
	ReadCost              int64 `json:"read_cost,omitempty"`                 // 阅读成本
	CommentCount          int64 `json:"comment_count,omitempty"`             // 评论次数
	CommentCost           int64 `json:"comment_cost,omitempty"`              // 评论成本
	PraiseCount           int64 `json:"praise_count,omitempty"`              // 点赞次数
	PraiseCost            int64 `json:"praise_cost,omitempty"`               // 点赞成本
	ForwardCount          int64 `json:"forward_count,omitempty"`             // 分享次数
	ForwardCost           int64 `json:"forward_cost,omitempty"`              // 分享成本
	NoInterestCount       int64 `json:"no_interest_count,omitempty"`         // 不感兴趣点击次数
	LiveStreamCrtClickCnt int64 `json:"live_stream_crt_click_cnt,omitempty"` // 素材点击次数

	// ========== 外层点击细分 ==========
	ClickImageCount               int64 `json:"click_image_count,omitempty"`                 // 外层图片点击次数
	ClickHeadCount                int64 `json:"click_head_count,omitempty"`                  // 头像点击次数
	ClickDetailCount              int64 `json:"click_detail_count,omitempty"`                // 文字链点击次数
	ClickPoiCount                 int64 `json:"click_poi_count,omitempty"`                   // 本地门店点击次数
	ZoneHeaderClickCount          int64 `json:"zone_header_click_count,omitempty"`           // 顶部运营位点击次数
	BasicInfoClientCount          int64 `json:"basic_info_client_count,omitempty"`           // 官方信息区点击次数
	AccountInfoClickCount         int64 `json:"account_info_click_count,omitempty"`          // 官方账号区点击次数
	ClkAccountLivingStatusPv      int64 `json:"clk_account_living_status_pv,omitempty"`      // 官方账号区直播中头像点击次数
	ClkAccountinfoWeappPv         int64 `json:"clk_accountinfo_weapp_pv,omitempty"`          // 官方账号区小程序组件点击次数
	ClkAccountinfoFinderPv        int64 `json:"clk_accountinfo_finder_pv,omitempty"`         // 官方账号区视频号组件点击次数
	ClkAccountinfoBizPv           int64 `json:"clk_accountinfo_biz_pv,omitempty"`            // 官方账号区公众号组件点击次数
	ClkAccountInfoProducttabPv    int64 `json:"clk_account_info_producttab_pv,omitempty"`    // 官方账号区系列tab总点击次数
	ClkAccountInfoProductdetailPv int64 `json:"clk_account_info_productdetail_pv,omitempty"` // 官方账号区商品卡总点击次数
	ActivityInfoClickCount        int64 `json:"activity_info_click_count,omitempty"`         // 热门活动区点击次数

	// ========== 线索 ==========
	OverallLeadsPurchaseCount int64 `json:"overall_leads_purchase_count,omitempty"` // 综合销售线索人数
	EffectiveLeadsCount       int64 `json:"effective_leads_count,omitempty"`        // 有效线索次数
	EffectiveCost             int64 `json:"effective_cost,omitempty"`               // 有效线索成本
	EffectLeadsPurchaseCount  int64 `json:"effect_leads_purchase_count,omitempty"`  // 有效线索人数
	EffectLeadsPurchaseCost   int64 `json:"effect_leads_purchase_cost,omitempty"`   // 有效线索成本（人数）

	// ========== 落地页 ==========
	PlatformPageViewCount  int64   `json:"platform_page_view_count,omitempty"`   // 落地页曝光次数
	PlatformPageViewRate   float64 `json:"platform_page_view_rate,omitempty"`    // 落地页曝光率
	LanButtonClickCount    int64   `json:"lan_button_click_count,omitempty"`     // 落地页组件点击次数
	LanJumpButtonClickers  int64   `json:"lan_jump_button_clickers,omitempty"`   // 落地页组件点击人数
	LanButtonClickCost     int64   `json:"lan_button_click_cost,omitempty"`      // 落地页点击成本
	LanJumpButtonCtr       float64 `json:"lan_jump_button_ctr,omitempty"`        // 落地页点击率（人数）
	LanJumpButtonClickCost int64   `json:"lan_jump_button_click_cost,omitempty"` // 落地页点击成本（人数）
	LanJumpButtonRate      float64 `json:"lan_jump_button_rate,omitempty"`       // 落地页点击率

	// ========== 关键页面 ==========
	KeyPageViewCount          int64   `json:"key_page_view_count,omitempty"`            // 关键页面访问次数
	KeyPageViewByDisplayCount int64   `json:"key_page_view_by_display_count,omitempty"` // 关键页面访问次数（曝光归因）
	KeyPageViewByClickCount   int64   `json:"key_page_view_by_click_count,omitempty"`   // 关键页面访问次数（点击归因）
	KeyPageUv                 int64   `json:"key_page_uv,omitempty"`                    // 关键页面访问人数
	KeyPageViewCost           int64   `json:"key_page_view_cost,omitempty"`             // 关键页面访问成本
	KeyPageViewRate           float64 `json:"key_page_view_rate,omitempty"`             // 关键页面访问率

	// ========== 商品页 ==========
	LandingCommodityDetailExpPv        int64   `json:"landing_commodity_detail_exp_pv,omitempty"`          // 商品页浏览次数
	AppCommodityPageViewByDisplayCount int64   `json:"app_commodity_page_view_by_display_count,omitempty"` // 商品页浏览次数（曝光归因）
	AppCommodityPageViewByClickCount   int64   `json:"app_commodity_page_view_by_click_count,omitempty"`   // 商品页浏览次数（点击归因）
	ViewCommodityPageUv                int64   `json:"view_commodity_page_uv,omitempty"`                   // 商品页浏览人数
	WebCommodityPageViewRate           float64 `json:"web_commodity_page_view_rate,omitempty"`             // 商品页浏览率
	WebCommodityPageViewCost           int64   `json:"web_commodity_page_view_cost,omitempty"`             // 商品页浏览成本

	// ========== 导航 ==========
	OwnPageNavigationCount      int64 `json:"own_page_navigation_count,omitempty"`      // 自有页导航次数
	OwnPageNaviCost             int64 `json:"own_page_navi_cost,omitempty"`             // 自有页导航成本
	PlatformPageNavigationCount int64 `json:"platform_page_navigation_count,omitempty"` // 平台页导航次数
	PlatformPageNavigationCost  int64 `json:"platform_page_navigation_cost,omitempty"`  // 平台页导航成本
	PlatformShopNavigationCount int64 `json:"platform_shop_navigation_count,omitempty"` // 平台页门店点击次数
	PlatformShopNavigationCost  int64 `json:"platform_shop_navigation_cost,omitempty"`  // 平台页门店页导航成本

	// ========== 活动页与引导页 ==========
	ActivePageViews                    int64 `json:"active_page_views,omitempty"`                       // 活动页面访问次数
	ActivePageViewers                  int64 `json:"active_page_viewers,omitempty"`                     // 活动页面访问人数
	ActivePageInteractionAmount        int64 `json:"active_page_interaction_amount,omitempty"`          // 活动页面互动次数
	ActivePageInteractionUsers         int64 `json:"active_page_interaction_users,omitempty"`           // 活动页面互动人数
	GuideToFollowPageViews             int64 `json:"guide_to_follow_page_views,omitempty"`              // 加粉引导页浏览次数
	GuideToFollowPageViewers           int64 `json:"guide_to_follow_page_viewers,omitempty"`            // 加粉引导页浏览人数
	GuideToFollowPageInteractionAmount int64 `json:"guide_to_follow_page_interaction_amount,omitempty"` // 加粉引导页互动次数
	GuideToFollowPageInteractionUsers  int64 `json:"guide_to_follow_page_interaction_users,omitempty"`  // 加粉引导页互动人数
	PlatformKeyPageViewUserCount       int64 `json:"platform_key_page_view_user_count,omitempty"`       // 落地页曝光人数

	// ========== 咨询 ==========
	PageConsultCount      int64   `json:"page_consult_count,omitempty"`       // 在线咨询次数
	ConsultUvCount        int64   `json:"consult_uv_count,omitempty"`         // 在线咨询人数
	PageConsultRate       float64 `json:"page_consult_rate,omitempty"`        // 在线咨询率
	PageConsultCost       int64   `json:"page_consult_cost,omitempty"`        // 在线咨询成本
	ConsultLeaveInfoUsers int64   `json:"consult_leave_info_users,omitempty"` // 咨询留资人数
	ConsultLeaveInfoCost  int64   `json:"consult_leave_info_cost,omitempty"`  // 咨询留咨成本
	PotentialConsultCount int64   `json:"potential_consult_count,omitempty"`  // 潜在客户线索次数-咨询
	EffectiveConsultCount int64   `json:"effective_consult_count,omitempty"`  // 有效线索次数-咨询
	ToolConsultCount      int64   `json:"tool_consult_count,omitempty"`       // 附加创意智能咨询次数

	// ========== 表单预约 ==========
	PageReservationCount          int64   `json:"page_reservation_count,omitempty"`            // 表单预约次数
	PageReservationByDisplayCount int64   `json:"page_reservation_by_display_count,omitempty"` // 表单预约次数（曝光归因）
	PageReservationByClickCount   int64   `json:"page_reservation_by_click_count,omitempty"`   // 表单预约次数（点击归因）
	ReservationUv                 int64   `json:"reservation_uv,omitempty"`                    // 表单预约人数
	ReservationAmount             int64   `json:"reservation_amount,omitempty"`                // 表单预约金额
	PageReservationCost           int64   `json:"page_reservation_cost,omitempty"`             // 表单预约成本
	PageReservationCostWithPeople int64   `json:"page_reservation_cost_with_people,omitempty"` // 表单预约成本（人数）
	PageReservationRate           float64 `json:"page_reservation_rate,omitempty"`             // 表单预约率
	PageReservationRoi            float64 `json:"page_reservation_roi,omitempty"`              // 表单预约ROI
	BizReservationUv              int64   `json:"biz_reservation_uv,omitempty"`                // 公众号内表单预约人数
	BizReservationFollowRate      float64 `json:"biz_reservation_follow_rate,omitempty"`       // 公众号内表单预约率
	ExternalFormReservationCount  int64   `json:"external_form_reservation_count,omitempty"`   // 附加创意表单预约次数
	PotentialReserveCount         int64   `json:"potential_reserve_count,omitempty"`           // 潜在客户线索次数-表单
	ReservationCheckUv            int64   `json:"reservation_check_uv,omitempty"`              // 意向表单人数
	ReservationCheckUvCost        int64   `json:"reservation_check_uv_cost,omitempty"`         // 意向表单成本（人数）
	ReservationCheckUvRate        float64 `json:"reservation_check_uv_rate,omitempty"`         // 意向表单率（人数）
	EffectiveReserveCount         int64   `json:"effective_reserve_count,omitempty"`           // 有效线索次数-表单
	ValidLeadsUv                  int64   `json:"valid_leads_uv,omitempty"`                    // 有效线索人数-表单
	TryOutIntentionUv             int64   `json:"try_out_intention_uv,omitempty"`              // 试听意向人数
	IneffectiveLeadsUv            int64   `json:"ineffective_leads_uv,omitempty"`              // 无效线索人数

	// ========== 优惠券 ==========
	CouponGetPv                  int64   `json:"coupon_get_pv,omitempty"`                    // 领券次数
	CouponGetCost                int64   `json:"coupon_get_cost,omitempty"`                  // 领券成本
	CouponGetRate                float64 `json:"coupon_get_rate,omitempty"`                  // 领券率
	PlatformCouponClickCount     int64   `json:"platform_coupon_click_count,omitempty"`      // 平台优惠券点击次数
	PurchaseAmountWithCoupon     int64   `json:"purchase_amount_with_coupon,omitempty"`      // 用券金额
	CouponPurchaseRate           float64 `json:"coupon_purchase_rate,omitempty"`             // 用券率
	PurchaseAmountWithCouponCost int64   `json:"purchase_amount_with_coupon_cost,omitempty"` // 用券成本

	// ========== 电话 ==========
	PagePhoneCallDirectCount int64   `json:"page_phone_call_direct_count,omitempty"` // 电话直拨次数
	PagePhoneCallDirectRate  float64 `json:"page_phone_call_direct_rate,omitempty"`  // 电话直拨率
	PagePhoneCallDirectCost  int64   `json:"page_phone_call_direct_cost,omitempty"`  // 电话直拨成本
	PotentialPhoneCount      int64   `json:"potential_phone_count,omitempty"`        // 潜在客户线索次数-电话
	PotentialCustomerPhoneUv int64   `json:"potential_customer_phone_uv,omitempty"`  // 潜在客户电话独立人数
	EffectivePhoneCount      int64   `json:"effective_phone_count,omitempty"`        // 有效线索次数-电话
	ValidPhoneUv             int64   `json:"valid_phone_uv,omitempty"`               // 有效电话独立人数

	// ========== 核销与到店 ==========
	CouponUsageNumber int64   `json:"coupon_usage_number,omitempty"` // 核销次数
	CouponUsageRate   float64 `json:"coupon_usage_rate,omitempty"`   // 核销率
	CouponUsageCost   int64   `json:"coupon_usage_cost,omitempty"`   // 核销成本
	StoreVisitor      int64   `json:"store_visitor,omitempty"`       // 到店客流

	// ========== 微信本地支付 ==========
	WechatLocalPayCount     int64   `json:"wechat_local_pay_count,omitempty"`     // 微信本地支付次数
	WechatLocalPayuserCount int64   `json:"wechat_local_payuser_count,omitempty"` // 微信本地支付人数
	WechatLocalPayAmount    int64   `json:"wechat_local_pay_amount,omitempty"`    // 微信本地支付金额
	WechatLocalPayRoi       float64 `json:"wechat_local_pay_roi,omitempty"`       // 微信本地支付ROI

	// ========== 课程与关注 ==========
	ClassParticipatedFisrtUv     int64   `json:"class_participated_fisrt_uv,omitempty"`      // 首次参课人数
	ClassParticipatedFisrtUvCost int64   `json:"class_participated_fisrt_uv_cost,omitempty"` // 首次参课成本
	ClassParticipatedFisrtUvRate float64 `json:"class_participated_fisrt_uv_rate,omitempty"` // 首次参课率
	ScanFollowCount              int64   `json:"scan_follow_count,omitempty"`                // 扫码关注次数
	ScanFollowUserCount          int64   `json:"scan_follow_user_count,omitempty"`           // 扫码关注人数
	ScanFollowUserCost           int64   `json:"scan_follow_user_cost,omitempty"`            // 扫码关注成本
	ScanFollowUserRate           float64 `json:"scan_follow_user_rate,omitempty"`            // 扫码关注率

	// ========== 企微 ==========
	AfterAddWecomConsultDedupPv       int64 `json:"after_add_wecom_consult_dedup_pv,omitempty"`        // 企微咨询去重次数
	AfterAddWecomConsultDedupPvCost   int64 `json:"after_add_wecom_consult_dedup_pv_cost,omitempty"`   // 企微咨询去重成本
	AfterAddWecomIntentionDedupPv     int64 `json:"after_add_wecom_intention_dedup_pv,omitempty"`      // 企微意向去重次数
	AfterAddWecomIntentionDedupPvCost int64 `json:"after_add_wecom_intention_dedup_pv_cost,omitempty"` // 企微意向去重成本

	// ========== 社群 ==========
	JoinChatGroupAmount         int64   `json:"join_chat_group_amount,omitempty"`           // 入群次数
	JoinChatGroupNumberOfPeople int64   `json:"join_chat_group_number_of_people,omitempty"` // 入群人数
	JoinChatGroupCostByPeople   int64   `json:"join_chat_group_cost_by_people,omitempty"`   // 入群成本（人数）
	QuitChatGroupAmount         int64   `json:"quit_chat_group_amount,omitempty"`           // 退群次数
	QuitChatGroupRate           float64 `json:"quit_chat_group_rate,omitempty"`             // 退群率

	// ========== 加粉 ==========
	ScanCodeAddFansCount     int64 `json:"scan_code_add_fans_count,omitempty"`      // 扫码加粉次数
	ScanCodeAddFansCountCost int64 `json:"scan_code_add_fans_count_cost,omitempty"` // 扫码加粉成本
	ScanCodeAddFansUv        int64 `json:"scan_code_add_fans_uv,omitempty"`         // 扫码加粉人数
	ScanCodeAddFansUvCost    int64 `json:"scan_code_add_fans_uv_cost,omitempty"`    // 扫码加粉人数成本

	// ========== 企微个人 ==========
	WecomAddPersonalDedupPv     int64 `json:"wecom_add_personal_dedup_pv,omitempty"`      // 添加企微个人去重次数
	WecomAddPersonalDedupPvCost int64 `json:"wecom_add_personal_dedup_pv_cost,omitempty"` // 添加企微个人去重成本

	// ========== 其他行为 ==========
	LotteryLeadsCount int64 `json:"lottery_leads_count,omitempty"` // 抽奖线索次数
	LotteryLeadsCost  int64 `json:"lottery_leads_cost,omitempty"`  // 抽奖线索成本
	TryOutUser        int64 `json:"try_out_user,omitempty"`        // 试用用户数
	AddWishlistCount  int64 `json:"add_wishlist_count,omitempty"`  // 加入心愿单次数

	// ========== 购物车 ==========
	AddCartPv      int64 `json:"add_cart_pv,omitempty"`       // 加入购物车次数
	AddCartAmount  int64 `json:"add_cart_amount,omitempty"`   // 加入购物车金额
	AddToCartPrice int64 `json:"add_to_cart_price,omitempty"` // 加入购物车成本

	// ========== 下单 ==========
	OrderPv                      int64   `json:"order_pv,omitempty"`                          // 下单次数
	OrderUv                      int64   `json:"order_uv,omitempty"`                          // 下单人数
	OrderAmount                  int64   `json:"order_amount,omitempty"`                      // 下单金额
	OrderUnitPrice               int64   `json:"order_unit_price,omitempty"`                  // 下单客单价
	OrderRate                    float64 `json:"order_rate,omitempty"`                        // 下单率
	OrderCost                    int64   `json:"order_cost,omitempty"`                        // 下单成本
	OrderRoi                     float64 `json:"order_roi,omitempty"`                         // 下单ROI
	Order24hCount                int64   `json:"order_24h_count,omitempty"`                   // 24小时下单次数
	Order24hAmount               int64   `json:"order_24h_amount,omitempty"`                  // 24小时下单金额
	Order24hRate                 float64 `json:"order_24h_rate,omitempty"`                    // 24小时下单率
	Order24hCost                 int64   `json:"order_24h_cost,omitempty"`                    // 24小时下单成本
	Order24hRoi                  float64 `json:"order_24h_roi,omitempty"`                     // 24小时下单ROI
	FirstDayOrderCount           int64   `json:"first_day_order_count,omitempty"`             // 首日下单次数
	FirstDayOrderAmount          int64   `json:"first_day_order_amount,omitempty"`            // 首日下单金额
	FirstDayOrderRoi             float64 `json:"first_day_order_roi,omitempty"`               // 首日下单ROI
	OrderClk7dPv                 int64   `json:"order_clk_7d_pv,omitempty"`                   // 点击7天下单次数
	OrderClk7dAmount             int64   `json:"order_clk_7d_amount,omitempty"`               // 点击7天下单金额
	OrderClk7dUnitPrice          int64   `json:"order_clk_7d_unit_price,omitempty"`           // 点击7天下单客单价
	OrderClk7dRoi                float64 `json:"order_clk_7d_roi,omitempty"`                  // 点击7天下单ROI
	OrderClk15dPv                int64   `json:"order_clk_15d_pv,omitempty"`                  // 点击15天下单次数
	OrderClk15dAmount            int64   `json:"order_clk_15d_amount,omitempty"`              // 点击15天下单金额
	OrderClk15dUnitPrice         int64   `json:"order_clk_15d_unit_price,omitempty"`          // 点击15天下单客单价
	OrderClk15dRoi               float64 `json:"order_clk_15d_roi,omitempty"`                 // 点击15天下单ROI
	OrderClk30dPv                int64   `json:"order_clk_30d_pv,omitempty"`                  // 点击30天下单次数
	OrderClk30dAmount            int64   `json:"order_clk_30d_amount,omitempty"`              // 点击30天下单金额
	OrderClk30dUnitPrice         int64   `json:"order_clk_30d_unit_price,omitempty"`          // 点击30天下单客单价
	OrderClk30dRoi               float64 `json:"order_clk_30d_roi,omitempty"`                 // 点击30天下单ROI
	BizOrderUv                   int64   `json:"biz_order_uv,omitempty"`                      // 公众号下单人数
	BizOrderRate                 float64 `json:"biz_order_rate,omitempty"`                    // 公众号下单率
	OrderFollow1dPv              int64   `json:"order_follow_1d_pv,omitempty"`                // 关注1天下单次数
	OrderFollow1dAmount          int64   `json:"order_follow_1d_amount,omitempty"`            // 关注1天下单金额
	OrderByDisplayCount          int64   `json:"order_by_display_count,omitempty"`            // 下单次数（曝光归因）
	OrderByDisplayAmount         int64   `json:"order_by_display_amount,omitempty"`           // 下单金额（曝光归因）
	OrderByDisplayRate           float64 `json:"order_by_display_rate,omitempty"`             // 下单率（曝光归因）
	OrderByDisplayCost           int64   `json:"order_by_display_cost,omitempty"`             // 下单成本（曝光归因）
	OrderByDisplayRoi            float64 `json:"order_by_display_roi,omitempty"`              // 下单ROI（曝光归因）
	Order24hByDisplayCount       int64   `json:"order_24h_by_display_count,omitempty"`        // 24小时曝光归因下单次数
	Order24hByDisplayAmount      int64   `json:"order_24h_by_display_amount,omitempty"`       // 24小时曝光归因下单金额
	Order24hByDisplayRoi         float64 `json:"order_24h_by_display_roi,omitempty"`          // 24小时曝光归因下单ROI
	FirstDayOrderByDisplayCount  int64   `json:"first_day_order_by_display_count,omitempty"`  // 首日曝光归因下单次数
	FirstDayOrderByDisplayAmount int64   `json:"first_day_order_by_display_amount,omitempty"` // 首日曝光归因下单金额
	OrderByClickCount            int64   `json:"order_by_click_count,omitempty"`              // 下单次数（点击归因）
	OrderByClickAmount           int64   `json:"order_by_click_amount,omitempty"`             // 下单金额（点击归因）
	OrderByClickRate             float64 `json:"order_by_click_rate,omitempty"`               // 下单率（点击归因）
	OrderByClickCost             int64   `json:"order_by_click_cost,omitempty"`               // 下单成本（点击归因）
	OrderByClickRoi              float64 `json:"order_by_click_roi,omitempty"`                // 下单ROI（点击归因）
	FirstDayOrderByClickCount    int64   `json:"first_day_order_by_click_count,omitempty"`    // 首日点击归因下单次数
	FirstDayOrderByClickAmount   int64   `json:"first_day_order_by_click_amount,omitempty"`   // 首日点击归因下单金额
	Order24hByClickCount         int64   `json:"order_24h_by_click_count,omitempty"`          // 24小时点击归因下单次数
	Order24hByClickAmount        int64   `json:"order_24h_by_click_amount,omitempty"`         // 24小时点击归因下单金额
	Order24hByClickRoi           float64 `json:"order_24h_by_click_roi,omitempty"`            // 24小时点击归因下单ROI

	// ========== 发货与签收 ==========
	DeliverCount int64   `json:"deliver_count,omitempty"`  // 发货次数
	DeliverRate  float64 `json:"deliver_rate,omitempty"`   // 发货率
	DeliverCost  int64   `json:"deliver_cost,omitempty"`   // 发货成本
	SignInCount  int64   `json:"sign_in_count,omitempty"`  // 签收次数
	SignInAmount int64   `json:"sign_in_amount,omitempty"` // 签收金额
	SignInRate   float64 `json:"sign_in_rate,omitempty"`   // 签收率
	SignInCost   int64   `json:"sign_in_cost,omitempty"`   // 签收成本
	SignInRoi    float64 `json:"sign_in_roi,omitempty"`    // 签收ROI

	// ========== 会员卡 ==========
	PurchaseMemberCardPv        int64   `json:"purchase_member_card_pv,omitempty"`         // 购买会员卡次数
	PurchaseMemberCardDedupPv   int64   `json:"purchase_member_card_dedup_pv,omitempty"`   // 购买会员卡去重次数
	PurchaseMemberCardDedupCost int64   `json:"purchase_member_card_dedup_cost,omitempty"` // 购买会员卡去重成本
	PurchaseMemberCardDedupRate float64 `json:"purchase_member_card_dedup_rate,omitempty"` // 购买会员卡去重率

	// ========== 下载与安装 ==========
	DownloadCount      int64   `json:"download_count,omitempty"`       // 下载次数
	ActivatedRate      float64 `json:"activated_rate,omitempty"`       // 激活率
	DownloadRate       float64 `json:"download_rate,omitempty"`        // 下载率
	DownloadCost       int64   `json:"download_cost,omitempty"`        // 下载成本
	AddDesktopPv       int64   `json:"add_desktop_pv,omitempty"`       // 添加桌面次数
	AddDesktopCost     int64   `json:"add_desktop_cost,omitempty"`     // 添加桌面成本
	InstallCount       int64   `json:"install_count,omitempty"`        // 安装次数
	InstallRate        float64 `json:"install_rate,omitempty"`         // 安装率
	InstallCost        int64   `json:"install_cost,omitempty"`         // 安装成本
	ActivatedCount     int64   `json:"activated_count,omitempty"`      // 激活次数
	ActivatedCost      int64   `json:"activated_cost,omitempty"`       // 激活成本
	ClickActivatedRate float64 `json:"click_activated_rate,omitempty"` // 点击激活率

	// ========== 注册 ==========
	RegPv                  int64   `json:"reg_pv,omitempty"`                    // 注册次数
	RegisterByDisplayCount int64   `json:"register_by_display_count,omitempty"` // 注册次数（曝光归因）
	RegisterByClickCount   int64   `json:"register_by_click_count,omitempty"`   // 注册次数（点击归因）
	RegCost                int64   `json:"reg_cost,omitempty"`                  // 注册成本
	RegClkRate             float64 `json:"reg_clk_rate,omitempty"`              // 注册率
	ActivateRegisterRate   float64 `json:"activate_register_rate,omitempty"`    // 激活注册率
	RegPlaPv               int64   `json:"reg_pla_pv,omitempty"`                // 平台注册次数
	RegAllDedupPv          int64   `json:"reg_all_dedup_pv,omitempty"`          // 注册去重次数
	RegCostPla             int64   `json:"reg_cost_pla,omitempty"`              // 平台注册成本
	RegClickRatePla        float64 `json:"reg_click_rate_pla,omitempty"`        // 平台注册率
	RegDedupPv             int64   `json:"reg_dedup_pv,omitempty"`              // 注册去重次数
	MiniGameRegisterUsers  int64   `json:"mini_game_register_users,omitempty"`  // 小游戏注册用户数
	MiniGameRegisterCost   int64   `json:"mini_game_register_cost,omitempty"`   // 小游戏注册成本
	MiniGameRegisterRate   float64 `json:"mini_game_register_rate,omitempty"`   // 小游戏注册率
	BizRegCount            int64   `json:"biz_reg_count,omitempty"`             // 公众号注册次数
	BizRegUv               int64   `json:"biz_reg_uv,omitempty"`                // 公众号注册人数
	BizRegRate             float64 `json:"biz_reg_rate,omitempty"`              // 公众号注册率
	BizRegOrderAmount      int64   `json:"biz_reg_order_amount,omitempty"`      // 公众号注册订单金额
	BizRegCost             int64   `json:"biz_reg_cost,omitempty"`              // 公众号注册成本
	BizRegRoi              float64 `json:"biz_reg_roi,omitempty"`               // 公众号注册ROI

	// ========== 留存 ==========
	RetentionCount          int64   `json:"retention_count,omitempty"`             // 留存次数
	RetentionCost           int64   `json:"retention_cost,omitempty"`              // 留存成本
	RetentionRate           float64 `json:"retention_rate,omitempty"`              // 留存率
	AppRetentionD3Uv        int64   `json:"app_retention_d3_uv,omitempty"`         // 3日留存人数
	AppRetentionD3Cost      int64   `json:"app_retention_d3_cost,omitempty"`       // 3日留存成本
	AppRetentionD3Rate      float64 `json:"app_retention_d3_rate,omitempty"`       // 3日留存率
	AppRetentionD5Uv        int64   `json:"app_retention_d5_uv,omitempty"`         // 5日留存人数
	AppRetentionD5Cost      int64   `json:"app_retention_d5_cost,omitempty"`       // 5日留存成本
	AppRetentionD5Rate      float64 `json:"app_retention_d5_rate,omitempty"`       // 5日留存率
	AppRetentionD7Uv        int64   `json:"app_retention_d7_uv,omitempty"`         // 7日留存人数
	AppRetentionD7Cost      int64   `json:"app_retention_d7_cost,omitempty"`       // 7日留存成本
	AppRetentionD7Rate      float64 `json:"app_retention_d7_rate,omitempty"`       // 7日留存率
	AppRetentionLt7         int64   `json:"app_retention_lt7,omitempty"`           // 7日内留存人数
	AppRetentionLt7Cost     int64   `json:"app_retention_lt7_cost,omitempty"`      // 7日内留存成本
	MiniGameRetentionD1     int64   `json:"mini_game_retention_d1,omitempty"`      // 小游戏次日留存
	MiniGameRetentionD1Cost int64   `json:"mini_game_retention_d1_cost,omitempty"` // 小游戏次日留存成本
	MiniGameRetentionD1Rate float64 `json:"mini_game_retention_d1_rate,omitempty"` // 小游戏次日留存率
	AppKeyPageRetentionRate float64 `json:"app_key_page_retention_rate,omitempty"` // 关键页面留存率

	// ========== 付费 ==========
	PurchasePv       int64   `json:"purchase_pv,omitempty"`        // 付费次数
	PurchaseImpPv    int64   `json:"purchase_imp_pv,omitempty"`    // 付费次数（曝光归因）
	PurchaseClkPv    int64   `json:"purchase_clk_pv,omitempty"`    // 付费次数（点击归因）
	PurchaseAmount   int64   `json:"purchase_amount,omitempty"`    // 付费金额
	PurchaseCost     int64   `json:"purchase_cost,omitempty"`      // 付费成本
	PurchaseClkRate  float64 `json:"purchase_clk_rate,omitempty"`  // 付费率
	PurchaseActRate  float64 `json:"purchase_act_rate,omitempty"`  // 激活付费率
	PurchaseRoi      float64 `json:"purchase_roi,omitempty"`       // 付费ROI
	PurchaseActArpu  int64   `json:"purchase_act_arpu,omitempty"`  // 激活ARPU
	PurchaseRegArpu  int64   `json:"purchase_reg_arpu,omitempty"`  // 注册ARPU
	PurchaseRegArppu int64   `json:"purchase_reg_arppu,omitempty"` // 注册ARPPU

	// ========== 首日付费 ==========
	FirstDayPayCount       int64   `json:"first_day_pay_count,omitempty"`        // 首日付费次数
	FirstDayPayAmount      int64   `json:"first_day_pay_amount,omitempty"`       // 首日付费金额
	FirstDayPayCost        int64   `json:"first_day_pay_cost,omitempty"`         // 首日付费成本
	RoiActivatedD1         float64 `json:"roi_activated_d1,omitempty"`           // 首日付费ROI
	FirstDayPayAmountArpu  int64   `json:"first_day_pay_amount_arpu,omitempty"`  // 首日付费ARPU
	FirstDayPayAmountArppu int64   `json:"first_day_pay_amount_arppu,omitempty"` // 首日付费ARPPU

	// ========== 多日付费 ==========
	ActiveD3PayCount          int64   `json:"active_d3_pay_count,omitempty"`          // 3日付费次数
	PaymentAmountActivatedD3  int64   `json:"payment_amount_activated_d3,omitempty"`  // 3日付费金额
	RoiActivatedD3            float64 `json:"roi_activated_d3,omitempty"`             // 3日付费ROI
	ActiveD7PayCount          int64   `json:"active_d7_pay_count,omitempty"`          // 7日付费次数
	PaymentAmountActivatedD7  int64   `json:"payment_amount_activated_d7,omitempty"`  // 7日付费金额
	ActiveD7ClickPayRate      float64 `json:"active_d7_click_pay_rate,omitempty"`     // 7日点击付费率
	ActiveD7ActivePayRate     float64 `json:"active_d7_active_pay_rate,omitempty"`    // 7日激活付费率
	ActiveD7PayCost           int64   `json:"active_d7_pay_cost,omitempty"`           // 7日付费成本
	RoiActivatedD7            float64 `json:"roi_activated_d7,omitempty"`             // 7日付费ROI
	ActiveD14PayCount         int64   `json:"active_d14_pay_count,omitempty"`         // 14日付费次数
	PaymentAmountActivatedD14 int64   `json:"payment_amount_activated_d14,omitempty"` // 14日付费金额
	RoiActivatedD14           float64 `json:"roi_activated_d14,omitempty"`            // 14日付费ROI
	ActiveD30PayCount         int64   `json:"active_d30_pay_count,omitempty"`         // 30日付费次数
	PaymentAmountActivatedD30 int64   `json:"payment_amount_activated_d30,omitempty"` // 30日付费金额
	RoiActivatedD30           float64 `json:"roi_activated_d30,omitempty"`            // 30日付费ROI

	// ========== 小游戏付费 ==========
	MiniGamePayingArpu        int64   `json:"mini_game_paying_arpu,omitempty"`          // 小游戏付费ARPU
	Minigame24hPayUv          int64   `json:"minigame_24h_pay_uv,omitempty"`            // 小游戏24小时付费人数
	Minigame24hPayRoi         float64 `json:"minigame_24h_pay_roi,omitempty"`           // 小游戏24小时付费ROI
	Minigame24hPayArpu        int64   `json:"minigame_24h_pay_arpu,omitempty"`          // 小游戏24小时付费ARPU
	MiniGameFirstDayPayingRoi float64 `json:"mini_game_first_day_paying_roi,omitempty"` // 小游戏首日付费ROI
	MiniGamePayingArpuD1      int64   `json:"mini_game_paying_arpu_d1,omitempty"`       // 小游戏首日付费ARPU
	MiniGameD3PayCount        int64   `json:"mini_game_d3_pay_count,omitempty"`         // 小游戏3日付费次数
	MiniGamePayD3Uv           int64   `json:"mini_game_pay_d3_uv,omitempty"`            // 小游戏3日付费人数
	MiniGamePayD3Roi          float64 `json:"mini_game_pay_d3_roi,omitempty"`           // 小游戏3日付费ROI
	MiniGameD7PayCount        int64   `json:"mini_game_d7_pay_count,omitempty"`         // 小游戏7日付费次数
	MiniGamePayD7Uv           int64   `json:"mini_game_pay_d7_uv,omitempty"`            // 小游戏7日付费人数
	MiniGamePayD7Roi          float64 `json:"mini_game_pay_d7_roi,omitempty"`           // 小游戏7日付费ROI
	MiniGameD14PayCount       int64   `json:"mini_game_d14_pay_count,omitempty"`        // 小游戏14日付费次数
	MiniGamePayD14Uv          int64   `json:"mini_game_pay_d14_uv,omitempty"`           // 小游戏14日付费人数
	MiniGamePayD14Roi         float64 `json:"mini_game_pay_d14_roi,omitempty"`          // 小游戏14日付费ROI
	MiniGameD30PayCount       int64   `json:"mini_game_d30_pay_count,omitempty"`        // 小游戏30日付费次数
	MiniGamePayD30Uv          int64   `json:"mini_game_pay_d30_uv,omitempty"`           // 小游戏30日付费人数
	MiniGamePayingAmountD30   int64   `json:"mini_game_paying_amount_d30,omitempty"`    // 小游戏30日付费金额
	MiniGamePayD30Roi         float64 `json:"mini_game_pay_d30_roi,omitempty"`          // 小游戏30日付费ROI

	// ========== 首次付费 ==========
	FirstPayCount          int64   `json:"first_pay_count,omitempty"`            // 首次付费次数
	FirstPayCost           int64   `json:"first_pay_cost,omitempty"`             // 首次付费成本
	FirstPayRate           float64 `json:"first_pay_rate,omitempty"`             // 首次付费率
	LeadsPurchaseUv        int64   `json:"leads_purchase_uv,omitempty"`          // 线索购买人数
	MiniGameFirstPayAmount int64   `json:"mini_game_first_pay_amount,omitempty"` // 小游戏首次付费金额
	FirstDayFirstPayCount  int64   `json:"first_day_first_pay_count,omitempty"`  // 首日首次付费次数
	PaymentCostActivatedD1 int64   `json:"payment_cost_activated_d1,omitempty"`  // 首日付费成本
	FirstDayFirstPayRate   float64 `json:"first_day_first_pay_rate,omitempty"`   // 首日首次付费率

	// ========== 关键行为 ==========
	KeyBehaviorConversionsCount int64   `json:"key_behavior_conversions_count,omitempty"` // 关键行为转化量
	KeyBehaviorConversionsCost  int64   `json:"key_behavior_conversions_cost,omitempty"`  // 关键行为转化成本
	KeyBehaviorConversionsRate  float64 `json:"key_behavior_conversions_rate,omitempty"`  // 关键行为转化率

	// ========== 申请与授信 ==========
	ApplyPv                   int64   `json:"apply_pv,omitempty"`                      // 申请次数
	ApplyCost                 int64   `json:"apply_cost,omitempty"`                    // 申请成本
	BizPageApplyUv            int64   `json:"biz_page_apply_uv,omitempty"`             // 公众号申请人数
	BizPageApplyRate          float64 `json:"biz_page_apply_rate,omitempty"`           // 公众号申请率
	BizPageApplyCost          int64   `json:"biz_page_apply_cost,omitempty"`           // 公众号申请成本
	PreCreditPv               int64   `json:"pre_credit_pv,omitempty"`                 // 预授信次数
	PreCreditAmount           int64   `json:"pre_credit_amount,omitempty"`             // 预授信金额
	PreCreditCost             int64   `json:"pre_credit_cost,omitempty"`               // 预授信成本
	BizPreCreditUv            int64   `json:"biz_pre_credit_uv,omitempty"`             // 公众号预授信人数
	BizPreCreditUvCost        int64   `json:"biz_pre_credit_uv_cost,omitempty"`        // 公众号预授信成本
	CreditPv                  int64   `json:"credit_pv,omitempty"`                     // 授信次数
	CreditAmount              int64   `json:"credit_amount,omitempty"`                 // 授信金额
	CreditCost                int64   `json:"credit_cost,omitempty"`                   // 授信成本
	BizCreditUv               int64   `json:"biz_credit_uv,omitempty"`                 // 公众号授信人数
	BizCreditCost             int64   `json:"biz_credit_cost,omitempty"`               // 公众号授信成本
	BizCreditRate             float64 `json:"biz_credit_rate,omitempty"`               // 公众号授信率
	CreApplicationRate        float64 `json:"cre_application_rate,omitempty"`          // 授信申请率
	WithdrawDepositPv         int64   `json:"withdraw_deposit_pv,omitempty"`           // 提现次数
	WithdrawDepositAmount     int64   `json:"withdraw_deposit_amount,omitempty"`       // 提现金额
	BizWithdrawDepositsUv     int64   `json:"biz_withdraw_deposits_uv,omitempty"`      // 公众号提现人数
	BizWithdrawDepositsUvCost int64   `json:"biz_withdraw_deposits_uv_cost,omitempty"` // 公众号提现成本

	// ========== 游戏 ==========
	CouponClickCount        int64   `json:"coupon_click_count,omitempty"`           // 优惠券点击次数
	CouponIssueCount        int64   `json:"coupon_issue_count,omitempty"`           // 优惠券发放次数
	CouponGetCount          int64   `json:"coupon_get_count,omitempty"`             // 优惠券获取次数
	GameAuthorizeCount      int64   `json:"game_authorize_count,omitempty"`         // 游戏授权次数
	GameCreateRoleCount     int64   `json:"game_create_role_count,omitempty"`       // 游戏创角次数
	MiniGameCreateRoleUsers int64   `json:"mini_game_create_role_users,omitempty"`  // 小游戏创角人数
	MiniGameCreateRoleCost  int64   `json:"mini_game_create_role_cost,omitempty"`   // 小游戏创角成本
	MiniGameCreateRoleRate  float64 `json:"mini_game_create_role_rate,omitempty"`   // 小游戏创角率
	GameTutorialFinishCount int64   `json:"game_tutorial_finish_count,omitempty"`   // 游戏教程完成次数
	MiniGameKeyPageViewers  int64   `json:"mini_game_key_page_viewers,omitempty"`   // 小游戏关键页面访问人数
	MiniGameKeyPageViewCost int64   `json:"mini_game_key_page_view_cost,omitempty"` // 小游戏关键页面访问成本

	// ========== 广告变现 ==========
	AdMonetizationAmount  int64 `json:"ad_monetization_amount,omitempty"`   // 广告变现金额
	AdMonetizationActArpu int64 `json:"ad_monetization_act_arpu,omitempty"` // 广告变现激活ARPU
	AdMonetizationEcpm    int64 `json:"ad_monetization_ecpm,omitempty"`     // 广告变现eCPM
	AdMonetizationIpu     int64 `json:"ad_monetization_ipu,omitempty"`      // 广告变现IPU
	AdMonetizationLtv     int64 `json:"ad_monetization_ltv,omitempty"`      // 广告变现LTV

	// ========== 广告变现多日 ==========
	IncomeVal1  int64   `json:"income_val_1,omitempty"`  // 首日广告变现金额
	IncomeRoi1  float64 `json:"income_roi_1,omitempty"`  // 首日广告变现ROI
	IncomeVal3  int64   `json:"income_val_3,omitempty"`  // 3日广告变现金额
	IncomeRoi3  float64 `json:"income_roi_3,omitempty"`  // 3日广告变现ROI
	IncomeVal7  int64   `json:"income_val_7,omitempty"`  // 7日广告变现金额
	IncomeVal30 int64   `json:"income_val_30,omitempty"` // 30日广告变现金额
	IncomeVal60 int64   `json:"income_val_60,omitempty"` // 60日广告变现金额
	IncomeRoi30 float64 `json:"income_roi_30,omitempty"` // 30日广告变现ROI
	IncomeRoi60 float64 `json:"income_roi_60,omitempty"` // 60日广告变现ROI

	// ========== 种草 ==========
	EffectiveSeedingCount            int64   `json:"effective_seeding_count,omitempty"`            // 有效种草次数
	EffectiveSeedingCost             int64   `json:"effective_seeding_cost,omitempty"`             // 有效种草成本
	EffectiveSeedingRate             float64 `json:"effective_seeding_rate,omitempty"`             // 有效种草率
	EffectiveSeedingConverstionsRate float64 `json:"effective_seeding_convertions_rate,omitempty"` // 有效种草转化率

	// ========== 佣金与净订单 ==========
	CommissionAmount int64   `json:"commission_amount,omitempty"` // 佣金金额
	CommissionRoi    float64 `json:"commission_roi,omitempty"`    // 佣金ROI
	OrderNetPv       int64   `json:"order_net_pv,omitempty"`      // 净订单次数
	OrderNetAmount   int64   `json:"order_net_amount,omitempty"`  // 净订单金额
	OrderNetRoi      float64 `json:"order_net_roi,omitempty"`     // 净订单ROI
	OrderNetPvCost   int64   `json:"order_net_pv_cost,omitempty"` // 净订单成本
}
