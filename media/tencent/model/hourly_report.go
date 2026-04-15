package model

import (
	"errors"
)

// ========== 获取小时报表 ==========
// https://developers.e.qq.com/v3.0/docs/api/hourly_reports/get

// HourlyReportsGetReq 获取小时报表请求
type HourlyReportsGetReq struct {
	GlobalReq
	AccountID int64                  `json:"account_id"`          // 广告主帐号id (必填)
	Level     string                 `json:"level"`               // 获取报表类型级别 (必填)
	DateRange *HourlyReportDateRange `json:"date_range"`          // 日期范围 (必填)
	Filtering []*HourlyReportFilter  `json:"filtering,omitempty"` // 过滤条件
	GroupBy   []string               `json:"group_by"`            // 聚合参数 (必填)
	OrderBy   []*HourlyReportOrderBy `json:"order_by,omitempty"`  // 排序字段
	TimeLine  string                 `json:"time_line,omitempty"` // 时间口径
	Page      int                    `json:"page,omitempty"`      // 搜索页码，默认值：1
	PageSize  int                    `json:"page_size,omitempty"` // 一页显示的数据条数，默认值：10
	Fields    []string               `json:"fields"`              // 指定返回的字段列表 (必填)
}

// HourlyReportDateRange 日期范围
type HourlyReportDateRange struct {
	StartDate string `json:"start_date"` // 开始日期，格式：YYYY-MM-DD，且等于end_date (必填)
	EndDate   string `json:"end_date"`   // 结束日期，格式：YYYY-MM-DD，且等于start_date (必填)
}

// HourlyReportFilter 过滤条件
type HourlyReportFilter struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)
	Operator string   `json:"operator"` // 操作符 (必填)
	Values   []string `json:"values"`   // 字段取值 (必填)
}

// HourlyReportOrderBy 排序字段
type HourlyReportOrderBy struct {
	SortField string `json:"sort_field"` // 排序字段 (必填)
	SortType  string `json:"sort_type"`  // 排序方式 (必填)
}

// 小时报表级别常量（比日报表少）
var validHourlyReportLevels = map[string]bool{
	ReportLevelAdvertiser:      true,
	ReportLevelAdgroup:         true,
	ReportLevelDynamicCreative: true,
	ReportLevelChannel:         true,
	ReportLevelBidword:         true,
	ReportLevelProject:         true,
	ReportLevelProjectCreative: true,
	ReportLevelVideoHighlight:  true,
}

// 小时报表过滤字段常量
const (
	HourlyReportFilterAdgroupId         = "adgroup_id"
	HourlyReportFilterDynamicCreativeId = "dynamic_creative_id"
	HourlyReportFilterComponentId       = "component_id"
	HourlyReportFilterBidwordId         = "bidword_id"
	HourlyReportFilterChannelId         = "channel_id"
	HourlyReportFilterComponentType     = "component_type"
	HourlyReportFilterImageId           = "image_id"
	HourlyReportFilterVideoId           = "video_id"
)

// 小时报表限制常量
const (
	HourlyReportMinFilteringCount = 1
	HourlyReportMaxFilteringCount = 40
	HourlyReportMinValuesCount    = 1
	HourlyReportMaxValuesCount    = 100
	HourlyReportMaxValuesLength   = 64
	HourlyReportMinGroupByCount   = 1
	HourlyReportMaxGroupByCount   = 10
	HourlyReportMaxGroupByLength  = 64
	HourlyReportMinOrderByCount   = 1
	HourlyReportMaxOrderByCount   = 2
	HourlyReportMinPage           = 1
	HourlyReportMaxPage           = 100
	HourlyReportMinPageSize       = 1
	HourlyReportMaxPageSize       = 2000
	HourlyReportDefaultPage       = 1
	HourlyReportDefaultPageSize   = 10
	HourlyReportMinFieldsCount    = 1
	HourlyReportMaxFieldsCount    = 1024
	HourlyReportMinFieldLength    = 1
	HourlyReportMaxFieldLength    = 64
	HourlyReportDateLength        = 10
	HourlyReportMaxDataLimit      = 20000 // page * pageSize 最大值
)

func (p *HourlyReportsGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page <= 0 {
		p.Page = HourlyReportDefaultPage
	}
	if p.PageSize <= 0 {
		p.PageSize = HourlyReportDefaultPageSize
	}
}

func (p *HourlyReportsGetReq) Validate() error {
	// 验证全局参数
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	// 验证account_id (必填)
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	// 验证level (必填)
	if p.Level == "" {
		return errors.New("level为必填")
	}
	if !validHourlyReportLevels[p.Level] {
		return errors.New("level值无效，小时报表仅支持：REPORT_LEVEL_ADVERTISER、REPORT_LEVEL_ADGROUP、REPORT_LEVEL_DYNAMIC_CREATIVE、REPORT_LEVEL_CHANNEL、REPORT_LEVEL_BIDWORD、REPORT_LEVEL_PROJECT、REPORT_LEVEL_PROJECT_CREATIVE、REPORT_LEVEL_VIDEO_HIGHLIGHT")
	}

	// 验证date_range (必填)
	if p.DateRange == nil {
		return errors.New("date_range为必填")
	}
	if p.DateRange.StartDate == "" {
		return errors.New("date_range.start_date为必填")
	}
	if len(p.DateRange.StartDate) != HourlyReportDateLength {
		return errors.New("date_range.start_date长度必须为10字节")
	}
	if p.DateRange.EndDate == "" {
		return errors.New("date_range.end_date为必填")
	}
	if len(p.DateRange.EndDate) != HourlyReportDateLength {
		return errors.New("date_range.end_date长度必须为10字节")
	}
	if p.DateRange.StartDate != p.DateRange.EndDate {
		return errors.New("小时报表date_range.start_date必须等于end_date")
	}

	// 验证filtering
	if len(p.Filtering) > 0 {
		if len(p.Filtering) < HourlyReportMinFilteringCount || len(p.Filtering) > HourlyReportMaxFilteringCount {
			return errors.New("filtering数组长度必须在1-40之间")
		}
		for _, f := range p.Filtering {
			if f.Field == "" {
				return errors.New("filtering.field为必填")
			}
			if f.Operator == "" {
				return errors.New("filtering.operator为必填")
			}
			if len(f.Values) < HourlyReportMinValuesCount || len(f.Values) > HourlyReportMaxValuesCount {
				return errors.New("filtering.values数组长度必须在1-100之间")
			}
			for _, v := range f.Values {
				if len(v) > HourlyReportMaxValuesLength {
					return errors.New("filtering.values字段长度不能超过64字节")
				}
			}
		}
	}

	// 验证group_by (必填)
	if len(p.GroupBy) < HourlyReportMinGroupByCount || len(p.GroupBy) > HourlyReportMaxGroupByCount {
		return errors.New("group_by数组长度必须在1-10之间")
	}
	for _, g := range p.GroupBy {
		if len(g) > HourlyReportMaxGroupByLength {
			return errors.New("group_by字段长度不能超过64字节")
		}
	}

	// 验证order_by
	if len(p.OrderBy) > 0 {
		if len(p.OrderBy) < HourlyReportMinOrderByCount || len(p.OrderBy) > HourlyReportMaxOrderByCount {
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
	if p.Page < HourlyReportMinPage || p.Page > HourlyReportMaxPage {
		return errors.New("page必须在1-100之间")
	}

	// 验证page_size
	if p.PageSize < HourlyReportMinPageSize || p.PageSize > HourlyReportMaxPageSize {
		return errors.New("page_size必须在1-2000之间")
	}

	// 验证page * pageSize <= 20000
	if p.Page*p.PageSize > HourlyReportMaxDataLimit {
		return errors.New("page * page_size必须小于等于20000")
	}

	// 验证fields (必填)
	if len(p.Fields) < HourlyReportMinFieldsCount || len(p.Fields) > HourlyReportMaxFieldsCount {
		return errors.New("fields数组长度必须在1-1024之间")
	}
	for _, f := range p.Fields {
		if len(f) < HourlyReportMinFieldLength || len(f) > HourlyReportMaxFieldLength {
			return errors.New("fields字段长度必须在1-64字节之间")
		}
	}

	return nil
}

// HourlyReportsGetResp 获取小时报表响应
type HourlyReportsGetResp struct {
	List     []*HourlyReportListItem `json:"list,omitempty"`
	PageInfo *PageInfo               `json:"page_info,omitempty"`
}

// HourlyReportListItem 小时报表列表项
type HourlyReportListItem struct {
	// 基础维度字段
	Hour      string `json:"hour,omitempty"`       // 小时
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

	// ========== 付费（含首日/多日/小游戏等，同日报表） ==========
	PurchasePv             int64   `json:"purchase_pv,omitempty"`                // 付费次数
	PurchaseImpPv          int64   `json:"purchase_imp_pv,omitempty"`            // 付费次数（曝光归因）
	PurchaseClkPv          int64   `json:"purchase_clk_pv,omitempty"`            // 付费次数（点击归因）
	PurchaseAmount         int64   `json:"purchase_amount,omitempty"`            // 付费金额
	PurchaseCost           int64   `json:"purchase_cost,omitempty"`              // 付费成本
	PurchaseClkRate        float64 `json:"purchase_clk_rate,omitempty"`          // 付费率
	PurchaseActRate        float64 `json:"purchase_act_rate,omitempty"`          // 激活付费率
	PurchaseRoi            float64 `json:"purchase_roi,omitempty"`               // 付费ROI
	PurchaseActArpu        int64   `json:"purchase_act_arpu,omitempty"`          // 激活ARPU
	PurchaseRegArpu        int64   `json:"purchase_reg_arpu,omitempty"`          // 注册ARPU
	PurchaseRegArppu       int64   `json:"purchase_reg_arppu,omitempty"`         // 注册ARPPU
	CheoutPv1d             int64   `json:"cheout_pv_1d,omitempty"`               // 1日充值次数
	CheoutFd               int64   `json:"cheout_fd,omitempty"`                  // 首日充值金额
	Cheout1dCost           int64   `json:"cheout_1d_cost,omitempty"`             // 1日充值成本
	Cheout1dRate           float64 `json:"cheout_1d_rate,omitempty"`             // 1日充值率
	CheoutFdReward         int64   `json:"cheout_fd_reward,omitempty"`           // 首日充值ROI
	CheoutPv3d             int64   `json:"cheout_pv_3d,omitempty"`               // 3日充值次数
	CheoutTd               int64   `json:"cheout_td,omitempty"`                  // 3日充值金额
	Cheout3dCost           int64   `json:"cheout_3d_cost,omitempty"`             // 3日充值成本
	Cheout3dRate           float64 `json:"cheout_3d_rate,omitempty"`             // 3日充值率
	CheoutTdReward         int64   `json:"cheout_td_reward,omitempty"`           // 3日充值ROI
	CheoutPv5d             int64   `json:"cheout_pv_5d,omitempty"`               // 5日充值次数
	Cheout5dRate           float64 `json:"cheout_5d_rate,omitempty"`             // 5日充值率
	Cheout5dCost           int64   `json:"cheout_5d_cost,omitempty"`             // 5日充值成本
	CheoutPv7d             int64   `json:"cheout_pv_7d,omitempty"`               // 7日充值次数
	CheoutOw               int64   `json:"cheout_ow,omitempty"`                  // 7日充值金额
	Cheout7dCost           int64   `json:"cheout_7d_cost,omitempty"`             // 7日充值成本
	Cheout7dRate           float64 `json:"cheout_7d_rate,omitempty"`             // 7日充值率
	CheoutOwReward         int64   `json:"cheout_ow_reward,omitempty"`           // 7日充值ROI
	CheoutTw               int64   `json:"cheout_tw,omitempty"`                  // 14日充值金额
	CheoutTwReward         int64   `json:"cheout_tw_reward,omitempty"`           // 14日充值ROI
	PurchaseClk15dPv       int64   `json:"purchase_clk_15d_pv,omitempty"`        // 点击15日付费次数
	Cheout15d              int64   `json:"cheout_15d,omitempty"`                 // 15日充值金额
	Cheout15dReward        int64   `json:"cheout_15d_reward,omitempty"`          // 15日充值ROI
	PurchaseClk30dPv       int64   `json:"purchase_clk_30d_pv,omitempty"`        // 点击30日付费次数
	CheoutOm               int64   `json:"cheout_om,omitempty"`                  // 30日充值金额
	CheoutOmReward         int64   `json:"cheout_om_reward,omitempty"`           // 30日充值ROI
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
	AdMonetizationAmount  int64   `json:"ad_monetization_amount,omitempty"`   // 广告变现金额
	AdMonetizationActArpu int64   `json:"ad_monetization_act_arpu,omitempty"` // 广告变现激活ARPU
	AdMonetizationEcpm    int64   `json:"ad_monetization_ecpm,omitempty"`     // 广告变现eCPM
	AdMonetizationIpu     int64   `json:"ad_monetization_ipu,omitempty"`      // 广告变现IPU
	AdMonetizationLtv     int64   `json:"ad_monetization_ltv,omitempty"`      // 广告变现LTV
	IncomeVal1            int64   `json:"income_val_1,omitempty"`             // 首日广告变现金额
	IncomeRoi1            float64 `json:"income_roi_1,omitempty"`             // 首日广告变现ROI
	IncomeVal3            int64   `json:"income_val_3,omitempty"`             // 3日广告变现金额
	IncomeRoi3            float64 `json:"income_roi_3,omitempty"`             // 3日广告变现ROI
	IncomeVal7            int64   `json:"income_val_7,omitempty"`             // 7日广告变现金额
	IncomeRoi7            float64 `json:"income_roi_7,omitempty"`             // 7日广告变现ROI
	IncomeVal14           int64   `json:"income_val_14,omitempty"`            // 14日广告变现金额
	IncomeRoi14           float64 `json:"income_roi_14,omitempty"`            // 14日广告变现ROI
	IncomeVal30           int64   `json:"income_val_30,omitempty"`            // 30日广告变现金额
	IncomeVal60           int64   `json:"income_val_60,omitempty"`            // 60日广告变现金额
	IncomeRoi30           float64 `json:"income_roi_30,omitempty"`            // 30日广告变现ROI
	IncomeRoi60           float64 `json:"income_roi_60,omitempty"`            // 60日广告变现ROI

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

	// ========== 视频号（小时报表独有） ==========
	VideoFollowCount         int64 `json:"video_follow_count,omitempty"`          // 视频号关注次数
	VideoPlayCount           int64 `json:"video_play_count,omitempty"`            // 视频号播放次数
	VideoHeartCount          int64 `json:"video_heart_count,omitempty"`           // 视频号点赞次数
	VideoCommentCount        int64 `json:"video_comment_count,omitempty"`         // 视频号评论次数
	ChannelsSharePlaPv       int64 `json:"channels_share_pla_pv,omitempty"`       // 视频号分享次数
	ChannelsReadOfflinePv    int64 `json:"channels_read_offline_pv,omitempty"`    // 视频号阅读次数
	ChannelsHeartOfflinePv   int64 `json:"channels_heart_offline_pv,omitempty"`   // 视频号离线点赞次数
	ChannelsCommentOfflinePv int64 `json:"channels_comment_offline_pv,omitempty"` // 视频号离线评论次数
	ChannelsShareOfflinePv   int64 `json:"channels_share_offline_pv,omitempty"`   // 视频号离线分享次数
	ChannelsFavOfflinePv     int64 `json:"channels_fav_offline_pv,omitempty"`     // 视频号离线收藏次数

	// ========== 直播（小时报表独有） ==========
	VideoLiveSubscribeCount          int64   `json:"video_live_subscribe_count,omitempty"`             // 直播预约次数
	VideoLiveExpCount                int64   `json:"video_live_exp_count,omitempty"`                   // 直播曝光次数
	LiveStreamExpUv                  int64   `json:"live_stream_exp_uv,omitempty"`                     // 直播曝光人数
	ChannelsLiveExitPlaDuration      float64 `json:"channels_live_exit_pla_duration,omitempty"`        // 直播观看时长
	VideoLiveHeartCount              int64   `json:"video_live_heart_count,omitempty"`                 // 直播点赞次数
	VideoLiveHeartUserCount          int64   `json:"video_live_heart_user_count,omitempty"`            // 直播点赞人数
	VideoLiveCommentCount            int64   `json:"video_live_comment_count,omitempty"`               // 直播评论次数
	VideoLiveCommentUserCount        int64   `json:"video_live_comment_user_count,omitempty"`          // 直播评论人数
	VideoLiveShareCount              int64   `json:"video_live_share_count,omitempty"`                 // 直播分享次数
	VideoLiveShareUserCount          int64   `json:"video_live_share_user_count,omitempty"`            // 直播分享人数
	VideoLiveCickCommodityCount      int64   `json:"video_live_cick_commodity_count,omitempty"`        // 直播商品点击次数
	VideoLiveClickCommodityUserCount int64   `json:"video_live_click_commodity_user_count,omitempty"`  // 直播商品点击人数
	VideoLiveCommodityBubbleExpCount int64   `json:"video_live_commodity_bubble_exp_count,omitempty"`  // 直播商品气泡曝光次数
	LiveStreamCommodityBubbleClkPv   int64   `json:"live_stream_commodity_bubble_clk_pv,omitempty"`    // 直播商品气泡点击次数
	LiveStreamCommodityShopBagClkPv  int64   `json:"live_stream_commodity_shop_bag_clk_pv,omitempty"`  // 直播购物袋点击次数
	LiveStreamCommodityShopListExpPv int64   `json:"live_stream_commodity_shop_list_exp_pv,omitempty"` // 直播购物清单曝光次数
	LiveStreamAvgTime                float64 `json:"live_stream_avg_time,omitempty"`                   // 直播平均观看时长

	// ========== 关注（小时报表独有） ==========
	FromFollowUv            int64   `json:"from_follow_uv,omitempty"`              // 关注人数
	FromFollowCost          int64   `json:"from_follow_cost,omitempty"`            // 关注成本
	FromFollowByDisplayUv   int64   `json:"from_follow_by_display_uv,omitempty"`   // 关注人数（曝光归因）
	FromFollowByDisplayCost int64   `json:"from_follow_by_display_cost,omitempty"` // 关注成本（曝光归因）
	FromFollowByClickUv     int64   `json:"from_follow_by_click_uv,omitempty"`     // 关注人数（点击归因）
	FromFollowByClickCost   int64   `json:"from_follow_by_click_cost,omitempty"`   // 关注成本（点击归因）
	BizFollowRate           float64 `json:"biz_follow_rate,omitempty"`             // 公众号关注率
	BizFollowCost           int64   `json:"biz_follow_cost,omitempty"`             // 公众号关注成本
	BizFollowUv             int64   `json:"biz_follow_uv,omitempty"`               // 公众号关注人数
	BizConsultCount         int64   `json:"biz_consult_count,omitempty"`           // 公众号咨询次数
	BizReadingCount         int64   `json:"biz_reading_count,omitempty"`           // 公众号阅读次数

	// ========== 破框播放（小时报表独有） ==========
	BreakFramePlayPv         int64   `json:"break_frame_play_pv,omitempty"`           // 破框播放次数
	BreakFramePlayUv         int64   `json:"break_frame_play_uv,omitempty"`           // 破框播放人数
	AvgBreakFramePlayPerUser float64 `json:"avg_break_frame_play_per_user,omitempty"` // 人均破框播放次数
	BreakFrameIpClkPv        int64   `json:"break_frame_ip_clk_pv,omitempty"`         // 破框IP点击次数
	BreakFrameIpClkUv        int64   `json:"break_frame_ip_clk_uv,omitempty"`         // 破框IP点击人数

	// ========== 素材互动（小时报表独有） ==========
	ClkMaterialUv      int64   `json:"clk_material_uv,omitempty"`        // 素材点击人数
	ClkMaterialRate    float64 `json:"clk_material_rate,omitempty"`      // 素材点击率
	ClkNickPv          int64   `json:"clk_nick_pv,omitempty"`            // 昵称点击次数
	ClkNickUv          int64   `json:"clk_nick_uv,omitempty"`            // 昵称点击人数
	ClkHeadUv          int64   `json:"clk_head_uv,omitempty"`            // 头像点击人数
	ClkActionBtnPv     int64   `json:"clk_action_btn_pv,omitempty"`      // 行动按钮点击次数
	ClkActionBtnUv     int64   `json:"clk_action_btn_uv,omitempty"`      // 行动按钮点击人数
	ClkTagCommentPv    int64   `json:"clk_tag_comment_pv,omitempty"`     // 评论标签点击次数
	ClkTagContentPv    int64   `json:"clk_tag_content_pv,omitempty"`     // 内容标签点击次数
	ClkPoiPv           int64   `json:"clk_poi_pv,omitempty"`             // POI点击次数
	ClkDetailUv        int64   `json:"clk_detail_uv,omitempty"`          // 详情点击人数
	ClkDetailRate      float64 `json:"clk_detail_rate,omitempty"`        // 详情点击率
	ClkSliderCardBtnPv int64   `json:"clk_slider_card_btn_pv,omitempty"` // 滑动卡片按钮点击次数

	// ========== 分享互动（小时报表独有） ==========
	CvsBubbleShareClkPv       int64   `json:"cvs_bubble_share_clk_pv,omitempty"`       // 气泡分享点击次数
	CvsBubbleShareClkUv       int64   `json:"cvs_bubble_share_clk_uv,omitempty"`       // 气泡分享点击人数
	LpStarPageExpPv           int64   `json:"lp_star_page_exp_pv,omitempty"`           // 落地页星标曝光次数
	FinderTopicSliderPv       int64   `json:"finder_topic_slider_pv,omitempty"`        // 话题滑动次数
	FinderTopicSliderPerUser  float64 `json:"finder_topic_slider_per_user,omitempty"`  // 人均话题滑动次数
	FinderTopicSliderManualPv int64   `json:"finder_topic_slider_manual_pv,omitempty"` // 手动话题滑动次数
	ShareFriendPv             int64   `json:"share_friend_pv,omitempty"`               // 分享好友次数
	ShareFeedPv               int64   `json:"share_feed_pv,omitempty"`                 // 分享朋友圈次数
	PraiseUv                  int64   `json:"praise_uv,omitempty"`                     // 点赞人数
	CommentUv                 int64   `json:"comment_uv,omitempty"`                    // 评论人数

	// ========== 参与和停留（小时报表独有） ==========
	EngagePv                  int64   `json:"engage_pv,omitempty"`                      // 参与次数
	InteractSuccPv            int64   `json:"interact_succ_pv,omitempty"`               // 互动成功次数
	DurationPerUser           float64 `json:"duration_per_user,omitempty"`              // 人均停留时长
	DurationOuterPerUser      float64 `json:"duration_outer_per_user,omitempty"`        // 外层人均停留时长
	DurationKeyPagePerUser    float64 `json:"duration_key_page_per_user,omitempty"`     // 关键页面人均停留时长
	ClkAdElementPv            int64   `json:"clk_ad_element_pv,omitempty"`              // 广告元素点击次数
	ChannsPraisePlaPv         int64   `json:"channels_praise_pla_pv,omitempty"`         // 视频号点赞次数（平台）
	ChannelsLiveOutEnterPlaUv int64   `json:"channels_live_out_enter_pla_uv,omitempty"` // 直播外层进入人数

	// ========== 品专区域点击（小时报表独有） ==========
	ClkBheaderPv     int64 `json:"clk_bheader_pv,omitempty"`      // 品专头部点击次数
	ClkBhNamePv      int64 `json:"clk_bh_name_pv,omitempty"`      // 品专头部名称点击次数
	ClkBhStorePv     int64 `json:"clk_bh_store_pv,omitempty"`     // 品专头部商店点击次数
	ClkBhServicePv   int64 `json:"clk_bh_service_pv,omitempty"`   // 品专头部服务点击次数
	ClkBhPhonePv     int64 `json:"clk_bh_phone_pv,omitempty"`     // 品专头部电话点击次数
	ClkBhAnimPv      int64 `json:"clk_bh_anim_pv,omitempty"`      // 品专头部动画点击次数
	ClkBaccountPv    int64 `json:"clk_baccount_pv,omitempty"`     // 品专账号区点击次数
	ClkBaLivetagPv   int64 `json:"clk_ba_livetag_pv,omitempty"`   // 品专账号直播标签点击次数
	ClkBaBizPv       int64 `json:"clk_ba_biz_pv,omitempty"`       // 品专账号公众号点击次数
	ClkBaFinderPv    int64 `json:"clk_ba_finder_pv,omitempty"`    // 品专账号视频号点击次数
	ClkBaWeappPv     int64 `json:"clk_ba_weapp_pv,omitempty"`     // 品专账号小程序点击次数
	ClkBaWegamePv    int64 `json:"clk_ba_wegame_pv,omitempty"`    // 品专账号小游戏点击次数
	ClkBaMorePv      int64 `json:"clk_ba_more_pv,omitempty"`      // 品专账号更多点击次数
	ClkBacountPv     int64 `json:"clk_bacount_pv,omitempty"`      // 品专账号数量点击次数
	ClkBmarketingPv  int64 `json:"clk_bmarketing_pv,omitempty"`   // 品专营销区点击次数
	ClkBmTabPv       int64 `json:"clk_bm_tab_pv,omitempty"`       // 品专营销标签点击次数
	ClkBmProductPv   int64 `json:"clk_bm_product_pv,omitempty"`   // 品专营销商品点击次数
	ClkBmActivityPv  int64 `json:"clk_bm_activity_pv,omitempty"`  // 品专营销活动点击次数
	ClkBmVerticalPv  int64 `json:"clk_bm_vertical_pv,omitempty"`  // 品专营销垂直点击次数
	ClkBmPrivilegePv int64 `json:"clk_bm_privilege_pv,omitempty"` // 品专营销特权点击次数
	ClkBmSeriesPv    int64 `json:"clk_bm_series_pv,omitempty"`    // 品专营销系列点击次数
	ClkBmDetailPv    int64 `json:"clk_bm_detail_pv,omitempty"`    // 品专营销详情点击次数
	ClkBmLivePv      int64 `json:"clk_bm_live_pv,omitempty"`      // 品专营销直播点击次数
	ClkBredpocketPv  int64 `json:"clk_bredpocket_pv,omitempty"`   // 品专红包点击次数
	ClkBrSubPv       int64 `json:"clk_br_sub_pv,omitempty"`       // 品专红包订阅点击次数
	ClkBrSharePv     int64 `json:"clk_br_share_pv,omitempty"`     // 品专红包分享点击次数
	ClkBrBtnPv       int64 `json:"clk_br_btn_pv,omitempty"`       // 品专红包按钮点击次数
	ClkBrDrivePv     int64 `json:"clk_br_drive_pv,omitempty"`     // 品专红包驱动点击次数
	ClkBquickPv      int64 `json:"clk_bquick_pv,omitempty"`       // 品专快捷点击次数
	ClkBappPv        int64 `json:"clk_bapp_pv,omitempty"`         // 品专应用点击次数
	ClkRpsPv         int64 `json:"clk_rps_pv,omitempty"`          // RPS点击次数
	SliderPv         int64 `json:"slider_pv,omitempty"`           // 滑动次数

	// ========== 其他小时报表独有字段 ==========
	ClkRedpocketBtnGetPv       int64   `json:"clk_redpocket_btn_get_pv,omitempty"`       // 红包获取按钮点击次数
	ClkRedpocketBtnSharePv     int64   `json:"clk_redpocket_btn_share_pv,omitempty"`     // 红包分享按钮点击次数
	ClkRedpocketBtnJumpPv      int64   `json:"clk_redpocket_btn_jump_pv,omitempty"`      // 红包跳转按钮点击次数
	ClkRedpocketBtnSubscribePv int64   `json:"clk_redpocket_btn_subscribe_pv,omitempty"` // 红包订阅按钮点击次数
	ClkBlessingCardPv          int64   `json:"clk_blessing_card_pv,omitempty"`           // 祝福卡片点击次数
	ClkShortcutMenusPv         int64   `json:"clk_shortcut_menus_pv,omitempty"`          // 快捷菜单点击次数
	ChannelsDetailBtnPv        int64   `json:"channels_detail_btn_pv,omitempty"`         // 视频号详情按钮点击次数
	ZoneHeaderLiveClickCnt     int64   `json:"zone_header_live_click_cnt,omitempty"`     // 顶部运营位直播点击次数
	ClkSliderCardProductPv     int64   `json:"clk_slider_card_product_pv,omitempty"`     // 滑动卡片商品点击次数
	ClkSliderCardProductUv     int64   `json:"clk_slider_card_product_uv,omitempty"`     // 滑动卡片商品点击人数
	InsuranceDedupPv           int64   `json:"insurance_dedup_pv,omitempty"`             // 保险去重次数
	ClkGoodsHeaderPv           int64   `json:"clk_goods_header_pv,omitempty"`            // 商品头部点击次数
	ClkGoodsInfoPv             int64   `json:"clk_goods_info_pv,omitempty"`              // 商品信息点击次数
	ClkGoodsRecommendPv        int64   `json:"clk_goods_recommend_pv,omitempty"`         // 商品推荐点击次数
	ClkMiddleShowwindowPv      int64   `json:"clk_middle_showwindow_pv,omitempty"`       // 中部橱窗点击次数
	ClkFooterPv                int64   `json:"clk_footer_pv,omitempty"`                  // 底部点击次数
	ClkMiddleGoodsPv           int64   `json:"clk_middle_goods_pv,omitempty"`            // 中部商品点击次数
	ClkMiddleBtnPv             int64   `json:"clk_middle_btn_pv,omitempty"`              // 中部按钮点击次数
	ClkMiddleSectionPv         int64   `json:"clk_middle_section_pv,omitempty"`          // 中部区域点击次数
	ClkMiddleGridviewPv        int64   `json:"clk_middle_gridview_pv,omitempty"`         // 中部网格点击次数
	ClkBreakPv                 int64   `json:"clk_break_pv,omitempty"`                   // 破框点击次数
	ClkRedpocketShakePv        int64   `json:"clk_redpocket_shake_pv,omitempty"`         // 红包摇一摇点击次数
	ClkRedpocketShakeUv        int64   `json:"clk_redpocket_shake_uv,omitempty"`         // 红包摇一摇点击人数
	ClkRelatedVideoPv          int64   `json:"clk_related_video_pv,omitempty"`           // 相关视频点击次数
	ClkBrandPediaPv            int64   `json:"clk_brand_pedia_pv,omitempty"`             // 品牌百科点击次数
	ClkActivityNewsPv          int64   `json:"clk_activity_news_pv,omitempty"`           // 活动资讯点击次数
	ClkLeftGridInfoPv          int64   `json:"clk_left_grid_info_pv,omitempty"`          // 左侧网格信息点击次数
	ClkLeftGridMiddlePv        int64   `json:"clk_left_grid_middle_pv,omitempty"`        // 左侧网格中部点击次数
	ClkRightGridPv             int64   `json:"clk_right_grid_pv,omitempty"`              // 右侧网格点击次数
	CvsCpnVideoPlayPv          int64   `json:"cvs_cpn_video_play_pv,omitempty"`          // 组件视频播放次数
	CvsCpnVideoPlayUv          int64   `json:"cvs_cpn_video_play_uv,omitempty"`          // 组件视频播放人数
	CvsCpnVideoPlayDuration    int64   `json:"cvs_cpn_video_play_duration,omitempty"`    // 组件视频播放时长
	AvgCpnplayVideoTime        float64 `json:"avg_cpnplay_video_time,omitempty"`         // 人均组件视频播放时长
	ClkAdFlipCardPv            int64   `json:"clk_ad_flip_card_pv,omitempty"`            // 翻牌点击次数
	AfterAddWecomNegativePv    int64   `json:"after_add_wecom_negative_pv,omitempty"`    // 企微负反馈次数
	ClkAppservicePv            int64   `json:"clk_appservice_pv,omitempty"`              // 小程序服务点击次数
	ClkTopicElementPv          int64   `json:"clk_topic_element_pv,omitempty"`           // 话题元素点击次数
	ClkHotElementPv            int64   `json:"clk_hot_element_pv,omitempty"`             // 热门元素点击次数
	ClkHotElementUv            int64   `json:"clk_hot_element_uv,omitempty"`             // 热门元素点击人数
	HotElementCtr              float64 `json:"hot_element_ctr,omitempty"`                // 热门元素点击率
	ClkInterpageBtnPv          int64   `json:"clk_interpage_btn_pv,omitempty"`           // 插屏按钮点击次数
	ClkWechatShopPv            int64   `json:"clk_wechat_shop_pv,omitempty"`             // 微信小店点击次数
	ClkTopicinfoPv             int64   `json:"clk_topicinfo_pv,omitempty"`               // 话题信息点击次数
	ClkLearnMorePv             int64   `json:"clk_learn_more_pv,omitempty"`              // 了解更多点击次数
	ClkSellingPointsElementPv  int64   `json:"clk_selling_points_element_pv,omitempty"`  // 卖点元素点击次数
	ClkSellingPointsElementUv  int64   `json:"clk_selling_points_element_uv,omitempty"`  // 卖点元素点击人数
	ClkFullwidthBackgroundPv   int64   `json:"clk_fullwidth_background_pv,omitempty"`    // 全宽背景点击次数
	ClkHeaderAreaPv            int64   `json:"clk_header_area_pv,omitempty"`             // 头部区域点击次数
}
