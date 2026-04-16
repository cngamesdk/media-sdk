package model

import (
	"errors"
)

// ========== 获取定向标签报表 ==========
// https://developers.e.qq.com/v3.0/docs/api/targeting_tag_reports/get

// TargetingTagReportsGetReq 获取定向标签报表请求
type TargetingTagReportsGetReq struct {
	GlobalReq
	AccountID int64                        `json:"account_id"`          // 广告主帐号id，有操作权限的帐号id，不支持代理商id (必填)
	Type      string                       `json:"type"`                // api类型 (必填)
	Level     string                       `json:"level"`               // 定向标签报表类型级别 (必填)
	DateRange *TargetingTagReportDateRange `json:"date_range"`          // 日期范围 (必填)
	Filtering []*TargetingTagReportFilter  `json:"filtering,omitempty"` // 过滤条件，level为ADGROUP或DYNAMIC_CREATIVE时必填
	GroupBy   []string                     `json:"group_by"`            // 聚合参数 (必填)
	OrderBy   []*TargetingTagReportOrderBy `json:"order_by,omitempty"`  // 排序字段
	TimeLine  string                       `json:"time_line,omitempty"` // 时间口径
	Page      int                          `json:"page,omitempty"`      // 搜索页码，默认值：1
	PageSize  int                          `json:"page_size,omitempty"` // 一页显示的数据条数，默认值：10
	Fields    []string                     `json:"fields"`              // 指定返回的字段列表 (必填)
}

// TargetingTagReportDateRange 日期范围
type TargetingTagReportDateRange struct {
	StartDate string `json:"start_date"` // 开始日期，格式：YYYY-MM-DD (必填)
	EndDate   string `json:"end_date"`   // 结束日期，格式：YYYY-MM-DD (必填)
}

// TargetingTagReportFilter 过滤条件
type TargetingTagReportFilter struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)，可选值：adgroup_id、gender
	Operator string   `json:"operator"` // 操作符 (必填)
	Values   []string `json:"values"`   // 字段取值 (必填)
}

// TargetingTagReportOrderBy 排序字段
type TargetingTagReportOrderBy struct {
	SortField string `json:"sort_field"` // 排序字段 (必填)
	SortType  string `json:"sort_type"`  // 排序方式 (必填)
}

// api类型常量
const (
	TargetingTagTypeGender         = "GENDER"
	TargetingTagTypeAge            = "AGE"
	TargetingTagTypeRegion         = "REGION"
	TargetingTagTypeCity           = "CITY"
	TargetingTagTypeCustomAudience = "CUSTOM_AUDIENCE"
	TargetingTagTypeOs             = "OS"
)

// 定向标签报表级别常量
const (
	TargetingTagLevelAdvertiser      = "ADVERTISER"
	TargetingTagLevelAdgroup         = "ADGROUP"
	TargetingTagLevelDynamicCreative = "DYNAMIC_CREATIVE"
)

// 过滤字段常量
const (
	TargetingTagReportFilterAdgroupId = "adgroup_id"
	TargetingTagReportFilterGender    = "gender"
)

// 过滤操作符常量
const (
	TargetingTagReportOperatorEquals = "EQUALS"
	TargetingTagReportOperatorIn     = "IN"
)

// 定向标签报表限制常量
const (
	TargetingTagReportMinFilteringCount = 1
	TargetingTagReportMaxFilteringCount = 2
	TargetingTagReportMinValuesCount    = 1
	TargetingTagReportMaxValuesCount    = 100
	TargetingTagReportMaxValuesLength   = 64
	TargetingTagReportMinGroupByCount   = 1
	TargetingTagReportMaxGroupByCount   = 10
	TargetingTagReportMaxGroupByLength  = 64
	TargetingTagReportMinOrderByCount   = 1
	TargetingTagReportMaxOrderByCount   = 2
	TargetingTagReportMinPage           = 1
	TargetingTagReportMaxPage           = 100
	TargetingTagReportMinPageSize       = 1
	TargetingTagReportMaxPageSize       = 2000
	TargetingTagReportDefaultPage       = 1
	TargetingTagReportDefaultPageSize   = 10
	TargetingTagReportMinFieldsCount    = 1
	TargetingTagReportMaxFieldsCount    = 1024
	TargetingTagReportMinFieldLength    = 1
	TargetingTagReportMaxFieldLength    = 64
	TargetingTagReportDateLength        = 10
	TargetingTagReportMaxDataLimit      = 20000 // page * pageSize 最大值
)

func (p *TargetingTagReportsGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page <= 0 {
		p.Page = TargetingTagReportDefaultPage
	}
	if p.PageSize <= 0 {
		p.PageSize = TargetingTagReportDefaultPageSize
	}
}

func (p *TargetingTagReportsGetReq) Validate() error {
	// 验证全局参数
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	// 验证account_id (必填)
	if p.AccountID <= 0 {
		return errors.New("account_id为必填")
	}

	// 验证type (必填)
	if p.Type == "" {
		return errors.New("type为必填")
	}
	if !isValidTargetingTagType(p.Type) {
		return errors.New("type值无效，允许值：GENDER、AGE、REGION、CITY、CUSTOM_AUDIENCE、OS")
	}

	// 验证level (必填)
	if p.Level == "" {
		return errors.New("level为必填")
	}
	if !isValidTargetingTagLevel(p.Type, p.Level) {
		return errors.New("level值无效，GENDER/AGE/REGION/CITY/OS允许值：ADVERTISER、ADGROUP；CUSTOM_AUDIENCE允许值：ADVERTISER、ADGROUP、DYNAMIC_CREATIVE")
	}

	// 验证date_range (必填)
	if p.DateRange == nil {
		return errors.New("date_range为必填")
	}
	if p.DateRange.StartDate == "" {
		return errors.New("date_range.start_date为必填")
	}
	if len(p.DateRange.StartDate) != TargetingTagReportDateLength {
		return errors.New("date_range.start_date长度必须为10字节")
	}
	if p.DateRange.EndDate == "" {
		return errors.New("date_range.end_date为必填")
	}
	if len(p.DateRange.EndDate) != TargetingTagReportDateLength {
		return errors.New("date_range.end_date长度必须为10字节")
	}
	if p.DateRange.StartDate > p.DateRange.EndDate {
		return errors.New("date_range.start_date必须小于等于end_date")
	}

	// level为ADGROUP或DYNAMIC_CREATIVE时filtering为必填
	if p.Level == TargetingTagLevelAdgroup || p.Level == TargetingTagLevelDynamicCreative {
		if len(p.Filtering) == 0 {
			return errors.New("level为ADGROUP或DYNAMIC_CREATIVE时filtering为必填")
		}
	}

	// 验证filtering
	if len(p.Filtering) > 0 {
		if len(p.Filtering) < TargetingTagReportMinFilteringCount || len(p.Filtering) > TargetingTagReportMaxFilteringCount {
			return errors.New("filtering数组长度必须在1-2之间")
		}
		for _, f := range p.Filtering {
			if f.Field == "" {
				return errors.New("filtering.field为必填")
			}
			if f.Field != TargetingTagReportFilterAdgroupId && f.Field != TargetingTagReportFilterGender {
				return errors.New("filtering.field值无效，允许值：adgroup_id、gender")
			}
			if f.Operator == "" {
				return errors.New("filtering.operator为必填")
			}
			if f.Field == TargetingTagReportFilterGender && f.Operator != TargetingTagReportOperatorEquals {
				return errors.New("filtering.field为gender时operator只允许EQUALS")
			}
			if f.Field == TargetingTagReportFilterAdgroupId &&
				f.Operator != TargetingTagReportOperatorEquals && f.Operator != TargetingTagReportOperatorIn {
				return errors.New("filtering.field为adgroup_id时operator允许值：EQUALS、IN")
			}
			if len(f.Values) < TargetingTagReportMinValuesCount || len(f.Values) > TargetingTagReportMaxValuesCount {
				return errors.New("filtering.values数组长度必须在1-100之间")
			}
			for _, v := range f.Values {
				if len(v) > TargetingTagReportMaxValuesLength {
					return errors.New("filtering.values字段长度不能超过64字节")
				}
			}
		}
	}

	// 验证group_by (必填)
	if len(p.GroupBy) < TargetingTagReportMinGroupByCount || len(p.GroupBy) > TargetingTagReportMaxGroupByCount {
		return errors.New("group_by数组长度必须在1-10之间")
	}
	for _, g := range p.GroupBy {
		if len(g) > TargetingTagReportMaxGroupByLength {
			return errors.New("group_by字段长度不能超过64字节")
		}
	}

	// 验证order_by
	if len(p.OrderBy) > 0 {
		if len(p.OrderBy) < TargetingTagReportMinOrderByCount || len(p.OrderBy) > TargetingTagReportMaxOrderByCount {
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
	if p.Page < TargetingTagReportMinPage || p.Page > TargetingTagReportMaxPage {
		return errors.New("page必须在1-100之间")
	}

	// 验证page_size
	if p.PageSize < TargetingTagReportMinPageSize || p.PageSize > TargetingTagReportMaxPageSize {
		return errors.New("page_size必须在1-2000之间")
	}

	// 验证page * pageSize <= 20000
	if p.Page*p.PageSize > TargetingTagReportMaxDataLimit {
		return errors.New("page * page_size必须小于等于20000")
	}

	// 验证fields (必填)
	if len(p.Fields) < TargetingTagReportMinFieldsCount || len(p.Fields) > TargetingTagReportMaxFieldsCount {
		return errors.New("fields数组长度必须在1-1024之间")
	}
	for _, f := range p.Fields {
		if len(f) < TargetingTagReportMinFieldLength || len(f) > TargetingTagReportMaxFieldLength {
			return errors.New("fields字段长度必须在1-64字节之间")
		}
	}

	return nil
}

// isValidTargetingTagType 验证api类型
func isValidTargetingTagType(t string) bool {
	validTypes := map[string]bool{
		TargetingTagTypeGender:         true,
		TargetingTagTypeAge:            true,
		TargetingTagTypeRegion:         true,
		TargetingTagTypeCity:           true,
		TargetingTagTypeCustomAudience: true,
		TargetingTagTypeOs:             true,
	}
	return validTypes[t]
}

// isValidTargetingTagLevel 验证报表级别（根据type不同允许的level不同）
func isValidTargetingTagLevel(tagType, level string) bool {
	if tagType == TargetingTagTypeCustomAudience {
		// CUSTOM_AUDIENCE支持：ADVERTISER, ADGROUP, DYNAMIC_CREATIVE
		return level == TargetingTagLevelAdvertiser ||
			level == TargetingTagLevelAdgroup ||
			level == TargetingTagLevelDynamicCreative
	}
	// 其他类型支持：ADVERTISER, ADGROUP
	return level == TargetingTagLevelAdvertiser || level == TargetingTagLevelAdgroup
}

// TargetingTagReportsGetResp 获取定向标签报表响应
type TargetingTagReportsGetResp struct {
	List     []*TargetingTagReportListItem `json:"list,omitempty"`
	PageInfo *PageInfo                     `json:"page_info,omitempty"`
}

// TargetingTagReportListItem 定向标签报表列表项
type TargetingTagReportListItem struct {
	AccountID                                  int64   `json:"account_id,omitempty"`
	AccountInfoClickCount                      int64   `json:"account_info_click_count,omitempty"`
	AcquisitionCost                            int64   `json:"acquisition_cost,omitempty"`
	ActivateRegisterRate                       float64 `json:"activate_register_rate,omitempty"`
	ActivatedCost                              int64   `json:"activated_cost,omitempty"`
	ActivatedCount                             int64   `json:"activated_count,omitempty"`
	ActivatedRate                              float64 `json:"activated_rate,omitempty"`
	ActiveD14PayCount                          int64   `json:"active_d14_pay_count,omitempty"`
	ActiveD30PayCount                          int64   `json:"active_d30_pay_count,omitempty"`
	ActiveD3PayCount                           int64   `json:"active_d3_pay_count,omitempty"`
	ActiveD5ClickFirstPayRate                  float64 `json:"active_d5_click_first_pay_rate,omitempty"`
	ActiveD5FirstPayCost                       int64   `json:"active_d5_first_pay_cost,omitempty"`
	ActiveD5FirstPayUV                         int64   `json:"active_d5_first_pay_uv,omitempty"`
	ActiveD7ActivePayRate                      float64 `json:"active_d7_active_pay_rate,omitempty"`
	ActiveD7ClickPayRate                       float64 `json:"active_d7_click_pay_rate,omitempty"`
	ActiveD7PayCost                            int64   `json:"active_d7_pay_cost,omitempty"`
	ActiveD7PayCount                           int64   `json:"active_d7_pay_count,omitempty"`
	ActivePageInteractionAmount                int64   `json:"active_page_interaction_amount,omitempty"`
	ActivePageInteractionUsers                 int64   `json:"active_page_interaction_users,omitempty"`
	ActivePageViewers                          int64   `json:"active_page_viewers,omitempty"`
	ActivePageViews                            int64   `json:"active_page_views,omitempty"`
	ActivityInfoClickCount                     int64   `json:"activity_info_click_count,omitempty"`
	AdMonetizationActArpu                      int64   `json:"ad_monetization_act_arpu,omitempty"`
	AdMonetizationActArpuReg                   int64   `json:"ad_monetization_act_arpu_reg,omitempty"`
	AdMonetizationActive14dPV                  int64   `json:"ad_monetization_active_14d_pv,omitempty"`
	AdMonetizationActive30dPV                  int64   `json:"ad_monetization_active_30d_pv,omitempty"`
	AdMonetizationActive3dPV                   int64   `json:"ad_monetization_active_3d_pv,omitempty"`
	AdMonetizationActive60dPV                  int64   `json:"ad_monetization_active_60d_pv,omitempty"`
	AdMonetizationActive7dPV                   int64   `json:"ad_monetization_active_7d_pv,omitempty"`
	AdMonetizationAmount                       int64   `json:"ad_monetization_amount,omitempty"`
	AdMonetizationArppu                        int64   `json:"ad_monetization_arppu,omitempty"`
	AdMonetizationBkPlaActive14dAmount         int64   `json:"ad_monetization_bk_pla_active_14d_amount,omitempty"`
	AdMonetizationBkPlaActive14dROI            float64 `json:"ad_monetization_bk_pla_active_14d_roi,omitempty"`
	AdMonetizationBkPlaActive24hAmount         int64   `json:"ad_monetization_bk_pla_active_24h_amount,omitempty"`
	AdMonetizationBkPlaActive3dAmount          int64   `json:"ad_monetization_bk_pla_active_3d_amount,omitempty"`
	AdMonetizationBkPlaActive3dROI             float64 `json:"ad_monetization_bk_pla_active_3d_roi,omitempty"`
	AdMonetizationBkPlaActive7dAmount          int64   `json:"ad_monetization_bk_pla_active_7d_amount,omitempty"`
	AdMonetizationBkPlaActive7dROI             float64 `json:"ad_monetization_bk_pla_active_7d_roi,omitempty"`
	AdMonetizationBkPlaCost                    int64   `json:"ad_monetization_bk_pla_cost,omitempty"`
	AdMonetizationBkPlaDedupActive1dPV         int64   `json:"ad_monetization_bk_pla_dedup_active_1d_pv,omitempty"`
	AdMonetizationBkPlaDedupActive24hArpu      int64   `json:"ad_monetization_bk_pla_dedup_active_24h_arpu,omitempty"`
	AdMonetizationBkPlaDedupActive24hCost      int64   `json:"ad_monetization_bk_pla_dedup_active_24h_cost,omitempty"`
	AdMonetizationBkPlaDedupActive24hPV        int64   `json:"ad_monetization_bk_pla_dedup_active_24h_pv,omitempty"`
	AdMonetizationBkPlaDedupActive24hROI       float64 `json:"ad_monetization_bk_pla_dedup_active_24h_roi,omitempty"`
	AdMonetizationBkPlaDedupPV                 int64   `json:"ad_monetization_bk_pla_dedup_pv,omitempty"`
	AdMonetizationCost                         int64   `json:"ad_monetization_cost,omitempty"`
	AdMonetizationDedupActive14dPV             int64   `json:"ad_monetization_dedup_active_14d_pv,omitempty"`
	AdMonetizationDedupActive30dPV             int64   `json:"ad_monetization_dedup_active_30d_pv,omitempty"`
	AdMonetizationDedupActive3dPV              int64   `json:"ad_monetization_dedup_active_3d_pv,omitempty"`
	AdMonetizationDedupActive60dPV             int64   `json:"ad_monetization_dedup_active_60d_pv,omitempty"`
	AdMonetizationDedupActive7dPV              int64   `json:"ad_monetization_dedup_active_7d_pv,omitempty"`
	AdMonetizationDedupRegActive14dPV          int64   `json:"ad_monetization_dedup_reg_active_14d_pv,omitempty"`
	AdMonetizationDedupRegActive30dPV          int64   `json:"ad_monetization_dedup_reg_active_30d_pv,omitempty"`
	AdMonetizationDedupRegActive60dPV          int64   `json:"ad_monetization_dedup_reg_active_60d_pv,omitempty"`
	AdMonetizationEcpm                         int64   `json:"ad_monetization_ecpm,omitempty"`
	AdMonetizationIpu                          int64   `json:"ad_monetization_ipu,omitempty"`
	AdMonetizationLtv                          int64   `json:"ad_monetization_ltv,omitempty"`
	AdMonetizationPenetrationRatD1             float64 `json:"ad_monetization_penetration_rat_d1,omitempty"`
	AdMonetizationPlaDedupActive1dPV           int64   `json:"ad_monetization_pla_dedup_active_1d_pv,omitempty"`
	AdMonetizationPlaDedupActive3dPV           int64   `json:"ad_monetization_pla_dedup_active_3d_pv,omitempty"`
	AdMonetizationPlaDedupActive7dPV           int64   `json:"ad_monetization_pla_dedup_active_7d_pv,omitempty"`
	AdMonetizationPlaDedupPV                   int64   `json:"ad_monetization_pla_dedup_pv,omitempty"`
	AdMonetizationPlaRegActive1dMixROI         float64 `json:"ad_monetization_pla_reg_active_1d_mix_roi,omitempty"`
	AdMonetizationPlaRegActive24hMixROI        float64 `json:"ad_monetization_pla_reg_active_24h_mix_roi,omitempty"`
	AdMonetizationRegActive14dAmount           int64   `json:"ad_monetization_reg_active_14d_amount,omitempty"`
	AdMonetizationRegActive14dPV               int64   `json:"ad_monetization_reg_active_14d_pv,omitempty"`
	AdMonetizationRegActive14dROI              float64 `json:"ad_monetization_reg_active_14d_roi,omitempty"`
	AdMonetizationRegActive30dAmount           int64   `json:"ad_monetization_reg_active_30d_amount,omitempty"`
	AdMonetizationRegActive30dPV               int64   `json:"ad_monetization_reg_active_30d_pv,omitempty"`
	AdMonetizationRegActive30dROI              float64 `json:"ad_monetization_reg_active_30d_roi,omitempty"`
	AdMonetizationRegActive60dAmount           int64   `json:"ad_monetization_reg_active_60d_amount,omitempty"`
	AdMonetizationRegActive60dPV               int64   `json:"ad_monetization_reg_active_60d_pv,omitempty"`
	AdMonetizationRegActive60dROI              float64 `json:"ad_monetization_reg_active_60d_roi,omitempty"`
	AdMonetizationROI                          float64 `json:"ad_monetization_roi,omitempty"`
	AdPayingCostD1                             int64   `json:"ad_paying_cost_d1,omitempty"`
	AdPayingUsers24h                           int64   `json:"ad_paying_users_24h,omitempty"`
	AdPayingUsers24hPla                        int64   `json:"ad_paying_users_24h_pla,omitempty"`
	AdPayingUsersD1                            int64   `json:"ad_paying_users_d1,omitempty"`
	AdPurArpuCostD1                            int64   `json:"ad_pur_arpu_cost_d1,omitempty"`
	AdPurArpuCostD124h                         int64   `json:"ad_pur_arpu_cost_d1_24h,omitempty"`
	AdPurArpuCostD124hPla                      int64   `json:"ad_pur_arpu_cost_d1_24h_pla,omitempty"`
	AddCartAmount                              int64   `json:"add_cart_amount,omitempty"`
	AddCartPV                                  int64   `json:"add_cart_pv,omitempty"`
	AddDesktopCost                             int64   `json:"add_desktop_cost,omitempty"`
	AddDesktopPV                               int64   `json:"add_desktop_pv,omitempty"`
	AddToCartPrice                             int64   `json:"add_to_cart_price,omitempty"`
	AddWishlistCount                           int64   `json:"add_wishlist_count,omitempty"`
	AdgroupID                                  int64   `json:"adgroup_id,omitempty"`
	AfterAddWecomConsultDedupPV                int64   `json:"after_add_wecom_consult_dedup_pv,omitempty"`
	AfterAddWecomConsultDedupPVCost            int64   `json:"after_add_wecom_consult_dedup_pv_cost,omitempty"`
	AfterAddWecomIntentionDedupPV              int64   `json:"after_add_wecom_intention_dedup_pv,omitempty"`
	AfterAddWecomIntentionDedupPVCost          int64   `json:"after_add_wecom_intention_dedup_pv_cost,omitempty"`
	AfterAddWecomNegativePV                    int64   `json:"after_add_wecom_negative_pv,omitempty"`
	AgeRange                                   string  `json:"age_range,omitempty"`
	AppAdPayingUsers                           int64   `json:"app_ad_paying_users,omitempty"`
	AppCommodityPageViewByClickCount           int64   `json:"app_commodity_page_view_by_click_count,omitempty"`
	AppCommodityPageViewByDisplayCount         int64   `json:"app_commodity_page_view_by_display_count,omitempty"`
	AppKeyPageRetentionRate                    float64 `json:"app_key_page_retention_rate,omitempty"`
	AppRetentionD3Cost                         int64   `json:"app_retention_d3_cost,omitempty"`
	AppRetentionD3Rate                         float64 `json:"app_retention_d3_rate,omitempty"`
	AppRetentionD3UV                           int64   `json:"app_retention_d3_uv,omitempty"`
	AppRetentionD5Cost                         int64   `json:"app_retention_d5_cost,omitempty"`
	AppRetentionD5Rate                         float64 `json:"app_retention_d5_rate,omitempty"`
	AppRetentionD5UV                           int64   `json:"app_retention_d5_uv,omitempty"`
	AppRetentionD7Cost                         int64   `json:"app_retention_d7_cost,omitempty"`
	AppRetentionD7Rate                         float64 `json:"app_retention_d7_rate,omitempty"`
	AppRetentionD7UV                           int64   `json:"app_retention_d7_uv,omitempty"`
	AppRetentionLt7                            int64   `json:"app_retention_lt7,omitempty"`
	AppRetentionLt7Cost                        int64   `json:"app_retention_lt7_cost,omitempty"`
	ApplyCost                                  int64   `json:"apply_cost,omitempty"`
	ApplyDedupPV                               int64   `json:"apply_dedup_pv,omitempty"`
	ApplyPV                                    int64   `json:"apply_pv,omitempty"`
	AuthorizePV                                int64   `json:"authorize_pv,omitempty"`
	AvgViewPerUser                             float64 `json:"avg_view_per_user,omitempty"`
	BasicInfoClientCount                       int64   `json:"basic_info_client_count,omitempty"`
	BfAllDedupPV                               int64   `json:"bf_all_dedup_pv,omitempty"`
	BfDedupPV                                  int64   `json:"bf_dedup_pv,omitempty"`
	BizConsultCount                            int64   `json:"biz_consult_count,omitempty"`
	BizCreditCost                              int64   `json:"biz_credit_cost,omitempty"`
	BizCreditRate                              float64 `json:"biz_credit_rate,omitempty"`
	BizCreditUV                                int64   `json:"biz_credit_uv,omitempty"`
	BizFollowCost                              int64   `json:"biz_follow_cost,omitempty"`
	BizFollowRate                              float64 `json:"biz_follow_rate,omitempty"`
	BizFollowUV                                int64   `json:"biz_follow_uv,omitempty"`
	BizOrderRate                               float64 `json:"biz_order_rate,omitempty"`
	BizOrderUV                                 int64   `json:"biz_order_uv,omitempty"`
	BizPageApplyCost                           int64   `json:"biz_page_apply_cost,omitempty"`
	BizPageApplyRate                           float64 `json:"biz_page_apply_rate,omitempty"`
	BizPageApplyUV                             int64   `json:"biz_page_apply_uv,omitempty"`
	BizPreCreditUV                             int64   `json:"biz_pre_credit_uv,omitempty"`
	BizPreCreditUVCost                         int64   `json:"biz_pre_credit_uv_cost,omitempty"`
	BizReadingCount                            int64   `json:"biz_reading_count,omitempty"`
	BizRegCost                                 int64   `json:"biz_reg_cost,omitempty"`
	BizRegCount                                int64   `json:"biz_reg_count,omitempty"`
	BizRegOrderAmount                          int64   `json:"biz_reg_order_amount,omitempty"`
	BizRegRate                                 float64 `json:"biz_reg_rate,omitempty"`
	BizRegROI                                  float64 `json:"biz_reg_roi,omitempty"`
	BizRegUV                                   int64   `json:"biz_reg_uv,omitempty"`
	BizReservationFollowRate                   float64 `json:"biz_reservation_follow_rate,omitempty"`
	BizReservationUV                           int64   `json:"biz_reservation_uv,omitempty"`
	BizWithdrawDepositsUV                      int64   `json:"biz_withdraw_deposits_uv,omitempty"`
	BizWithdrawDepositsUVCost                  int64   `json:"biz_withdraw_deposits_uv_cost,omitempty"`
	BreakFrameIPClkPV                          int64   `json:"break_frame_ip_clk_pv,omitempty"`
	BreakFramePlayPV                           int64   `json:"break_frame_play_pv,omitempty"`
	ChannelsCommentOfflinePV                   int64   `json:"channels_comment_offline_pv,omitempty"`
	ChannelsDetailBtnPV                        int64   `json:"channels_detail_btn_pv,omitempty"`
	ChannelsFavOfflinePV                       int64   `json:"channels_fav_offline_pv,omitempty"`
	ChannelsHeartOfflinePV                     int64   `json:"channels_heart_offline_pv,omitempty"`
	ChannelsLiveEleCommodityClkPlaDedupPV      int64   `json:"channels_live_ele_commodity_clk_pla_dedup_pv,omitempty"`
	ChannelsLiveEleCommodityClkPlaPV           int64   `json:"channels_live_ele_commodity_clk_pla_pv,omitempty"`
	ChannelsLiveEleConvClkPlaPV                int64   `json:"channels_live_ele_conv_clk_pla_pv,omitempty"`
	ChannelsLiveEleConvExpPlaPV                int64   `json:"channels_live_ele_conv_exp_pla_pv,omitempty"`
	ChannelsLiveExitPlaDuration                int64   `json:"channels_live_exit_pla_duration,omitempty"`
	ChannelsPraisePlaPV                        int64   `json:"channels_praise_pla_pv,omitempty"`
	ChannelsReadOfflinePV                      int64   `json:"channels_read_offline_pv,omitempty"`
	ChannelsShareOfflinePV                     int64   `json:"channels_share_offline_pv,omitempty"`
	ChannelsSharePlaPV                         int64   `json:"channels_share_pla_pv,omitempty"`
	Cheout15d                                  int64   `json:"cheout_15d,omitempty"`
	Cheout15dReward                            int64   `json:"cheout_15d_reward,omitempty"`
	Cheout1dCost                               int64   `json:"cheout_1d_cost,omitempty"`
	Cheout1dRate                               float64 `json:"cheout_1d_rate,omitempty"`
	Cheout3dCost                               int64   `json:"cheout_3d_cost,omitempty"`
	Cheout3dRate                               float64 `json:"cheout_3d_rate,omitempty"`
	Cheout5dCost                               int64   `json:"cheout_5d_cost,omitempty"`
	Cheout5dRate                               float64 `json:"cheout_5d_rate,omitempty"`
	Cheout7dCost                               int64   `json:"cheout_7d_cost,omitempty"`
	Cheout7dRate                               float64 `json:"cheout_7d_rate,omitempty"`
	CheoutFd                                   int64   `json:"cheout_fd,omitempty"`
	CheoutFdReward                             int64   `json:"cheout_fd_reward,omitempty"`
	CheoutOm                                   int64   `json:"cheout_om,omitempty"`
	CheoutOmReward                             int64   `json:"cheout_om_reward,omitempty"`
	CheoutOw                                   int64   `json:"cheout_ow,omitempty"`
	CheoutOwReward                             int64   `json:"cheout_ow_reward,omitempty"`
	CheoutPV1d                                 int64   `json:"cheout_pv_1d,omitempty"`
	CheoutPV3d                                 int64   `json:"cheout_pv_3d,omitempty"`
	CheoutPV5d                                 int64   `json:"cheout_pv_5d,omitempty"`
	CheoutPV7d                                 int64   `json:"cheout_pv_7d,omitempty"`
	CheoutTd                                   int64   `json:"cheout_td,omitempty"`
	CheoutTdReward                             int64   `json:"cheout_td_reward,omitempty"`
	CheoutTw                                   int64   `json:"cheout_tw,omitempty"`
	CheoutTwReward                             int64   `json:"cheout_tw_reward,omitempty"`
	CityID                                     int64   `json:"city_id,omitempty"`
	ClassParticipatedFisrtUV                   int64   `json:"class_participated_fisrt_uv,omitempty"`
	ClassParticipatedFisrtUVCost               int64   `json:"class_participated_fisrt_uv_cost,omitempty"`
	ClassParticipatedFisrtUVRate               float64 `json:"class_participated_fisrt_uv_rate,omitempty"`
	ClickActivatedRate                         float64 `json:"click_activated_rate,omitempty"`
	ClickDetailCount                           int64   `json:"click_detail_count,omitempty"`
	ClickHeadCount                             int64   `json:"click_head_count,omitempty"`
	ClickImageCount                            int64   `json:"click_image_count,omitempty"`
	ClickPoiCount                              int64   `json:"click_poi_count,omitempty"`
	ClickUserCount                             int64   `json:"click_user_count,omitempty"`
	ClkAccountInfoProductdetailPV              int64   `json:"clk_account_info_productdetail_pv,omitempty"`
	ClkAccountInfoProducttabPV                 int64   `json:"clk_account_info_producttab_pv,omitempty"`
	ClkAccountLivingStatusPV                   int64   `json:"clk_account_living_status_pv,omitempty"`
	ClkAccountinfoBizPV                        int64   `json:"clk_accountinfo_biz_pv,omitempty"`
	ClkAccountinfoFinderPV                     int64   `json:"clk_accountinfo_finder_pv,omitempty"`
	ClkAccountinfoWeappPV                      int64   `json:"clk_accountinfo_weapp_pv,omitempty"`
	ClkActionBtnPV                             int64   `json:"clk_action_btn_pv,omitempty"`
	ClkActivityNewsPV                          int64   `json:"clk_activity_news_pv,omitempty"`
	ClkAdElementPV                             int64   `json:"clk_ad_element_pv,omitempty"`
	ClkAdFlipCardPV                            int64   `json:"clk_ad_flip_card_pv,omitempty"`
	ClkAppservicePV                            int64   `json:"clk_appservice_pv,omitempty"`
	ClkBaBizPV                                 int64   `json:"clk_ba_biz_pv,omitempty"`
	ClkBaFinderPV                              int64   `json:"clk_ba_finder_pv,omitempty"`
	ClkBaLivetagPV                             int64   `json:"clk_ba_livetag_pv,omitempty"`
	ClkBaMorePV                                int64   `json:"clk_ba_more_pv,omitempty"`
	ClkBaWeappPV                               int64   `json:"clk_ba_weapp_pv,omitempty"`
	ClkBaWegamePV                              int64   `json:"clk_ba_wegame_pv,omitempty"`
	ClkBaccountPV                              int64   `json:"clk_baccount_pv,omitempty"`
	ClkBacountPV                               int64   `json:"clk_bacount_pv,omitempty"`
	ClkBappPV                                  int64   `json:"clk_bapp_pv,omitempty"`
	ClkBhAnimPV                                int64   `json:"clk_bh_anim_pv,omitempty"`
	ClkBhNamePV                                int64   `json:"clk_bh_name_pv,omitempty"`
	ClkBhPhonePV                               int64   `json:"clk_bh_phone_pv,omitempty"`
	ClkBhServicePV                             int64   `json:"clk_bh_service_pv,omitempty"`
	ClkBhStorePV                               int64   `json:"clk_bh_store_pv,omitempty"`
	ClkBheaderPV                               int64   `json:"clk_bheader_pv,omitempty"`
	ClkBlessingCardPV                          int64   `json:"clk_blessing_card_pv,omitempty"`
	ClkBmActivityPV                            int64   `json:"clk_bm_activity_pv,omitempty"`
	ClkBmDetailPV                              int64   `json:"clk_bm_detail_pv,omitempty"`
	ClkBmLivePV                                int64   `json:"clk_bm_live_pv,omitempty"`
	ClkBmPrivilegePV                           int64   `json:"clk_bm_privilege_pv,omitempty"`
	ClkBmProductPV                             int64   `json:"clk_bm_product_pv,omitempty"`
	ClkBmSeriesPV                              int64   `json:"clk_bm_series_pv,omitempty"`
	ClkBmTabPV                                 int64   `json:"clk_bm_tab_pv,omitempty"`
	ClkBmVerticalPV                            int64   `json:"clk_bm_vertical_pv,omitempty"`
	ClkBmarketingPV                            int64   `json:"clk_bmarketing_pv,omitempty"`
	ClkBquickPV                                int64   `json:"clk_bquick_pv,omitempty"`
	ClkBrBtnPV                                 int64   `json:"clk_br_btn_pv,omitempty"`
	ClkBrDrivePV                               int64   `json:"clk_br_drive_pv,omitempty"`
	ClkBrSharePV                               int64   `json:"clk_br_share_pv,omitempty"`
	ClkBrSubPV                                 int64   `json:"clk_br_sub_pv,omitempty"`
	ClkBrandPediaPV                            int64   `json:"clk_brand_pedia_pv,omitempty"`
	ClkBreakPV                                 int64   `json:"clk_break_pv,omitempty"`
	ClkBredpocketPV                            int64   `json:"clk_bredpocket_pv,omitempty"`
	ClkDetailRate                              float64 `json:"clk_detail_rate,omitempty"`
	ClkDetailUV                                int64   `json:"clk_detail_uv,omitempty"`
	ClkFooterPV                                int64   `json:"clk_footer_pv,omitempty"`
	ClkFullwidthBackgroundPV                   int64   `json:"clk_fullwidth_background_pv,omitempty"`
	ClkGoodsHeaderPV                           int64   `json:"clk_goods_header_pv,omitempty"`
	ClkGoodsInfoPV                             int64   `json:"clk_goods_info_pv,omitempty"`
	ClkGoodsRecommendPV                        int64   `json:"clk_goods_recommend_pv,omitempty"`
	ClkHeaderAreaPV                            int64   `json:"clk_header_area_pv,omitempty"`
	ClkHotElementPV                            int64   `json:"clk_hot_element_pv,omitempty"`
	ClkInterpageBtnPV                          int64   `json:"clk_interpage_btn_pv,omitempty"`
	ClkLearnMorePV                             int64   `json:"clk_learn_more_pv,omitempty"`
	ClkLeftGridInfoPV                          int64   `json:"clk_left_grid_info_pv,omitempty"`
	ClkLeftGridMiddlePV                        int64   `json:"clk_left_grid_middle_pv,omitempty"`
	ClkMaterialRate                            float64 `json:"clk_material_rate,omitempty"`
	ClkMaterialUV                              int64   `json:"clk_material_uv,omitempty"`
	ClkMiddleBtnPV                             int64   `json:"clk_middle_btn_pv,omitempty"`
	ClkMiddleGoodsPV                           int64   `json:"clk_middle_goods_pv,omitempty"`
	ClkMiddleGridviewPV                        int64   `json:"clk_middle_gridview_pv,omitempty"`
	ClkMiddleSectionPV                         int64   `json:"clk_middle_section_pv,omitempty"`
	ClkMiddleShowwindowPV                      int64   `json:"clk_middle_showwindow_pv,omitempty"`
	ClkNickPV                                  int64   `json:"clk_nick_pv,omitempty"`
	ClkPoiPV                                   int64   `json:"clk_poi_pv,omitempty"`
	ClkRedpocketBtnGetPV                       int64   `json:"clk_redpocket_btn_get_pv,omitempty"`
	ClkRedpocketBtnJumpPV                      int64   `json:"clk_redpocket_btn_jump_pv,omitempty"`
	ClkRedpocketBtnSharePV                     int64   `json:"clk_redpocket_btn_share_pv,omitempty"`
	ClkRedpocketBtnSubscribePV                 int64   `json:"clk_redpocket_btn_subscribe_pv,omitempty"`
	ClkRedpocketShakePV                        int64   `json:"clk_redpocket_shake_pv,omitempty"`
	ClkRelatedVideoPV                          int64   `json:"clk_related_video_pv,omitempty"`
	ClkRightGridPV                             int64   `json:"clk_right_grid_pv,omitempty"`
	ClkRpsPV                                   int64   `json:"clk_rps_pv,omitempty"`
	ClkSellingPointsElementPV                  int64   `json:"clk_selling_points_element_pv,omitempty"`
	ClkSellingPointsElementUV                  int64   `json:"clk_selling_points_element_uv,omitempty"`
	ClkShortcutMenusPV                         int64   `json:"clk_shortcut_menus_pv,omitempty"`
	ClkSliderCardBtnPV                         int64   `json:"clk_slider_card_btn_pv,omitempty"`
	ClkSliderCardProductPV                     int64   `json:"clk_slider_card_product_pv,omitempty"`
	ClkTagCommentPV                            int64   `json:"clk_tag_comment_pv,omitempty"`
	ClkTagContentPV                            int64   `json:"clk_tag_content_pv,omitempty"`
	ClkTopicElementPV                          int64   `json:"clk_topic_element_pv,omitempty"`
	ClkTopicinfoPV                             int64   `json:"clk_topicinfo_pv,omitempty"`
	ClkWechatShopPV                            int64   `json:"clk_wechat_shop_pv,omitempty"`
	CommentCost                                int64   `json:"comment_cost,omitempty"`
	CommentCount                               int64   `json:"comment_count,omitempty"`
	CommentUV                                  int64   `json:"comment_uv,omitempty"`
	CommissionAmount                           int64   `json:"commission_amount,omitempty"`
	CommissionROI                              float64 `json:"commission_roi,omitempty"`
	ConsultLeaveInfoCost                       int64   `json:"consult_leave_info_cost,omitempty"`
	ConsultLeaveInfoUsers                      int64   `json:"consult_leave_info_users,omitempty"`
	ConsultUVCount                             int64   `json:"consult_uv_count,omitempty"`
	ConversionsByClickCost                     int64   `json:"conversions_by_click_cost,omitempty"`
	ConversionsByClickCount                    int64   `json:"conversions_by_click_count,omitempty"`
	ConversionsByClickRate                     float64 `json:"conversions_by_click_rate,omitempty"`
	ConversionsByDisplayCost                   int64   `json:"conversions_by_display_cost,omitempty"`
	ConversionsByDisplayCount                  int64   `json:"conversions_by_display_count,omitempty"`
	ConversionsByDisplayRate                   float64 `json:"conversions_by_display_rate,omitempty"`
	ConversionsCost                            int64   `json:"conversions_cost,omitempty"`
	ConversionsCount                           int64   `json:"conversions_count,omitempty"`
	ConversionsRate                            float64 `json:"conversions_rate,omitempty"`
	Cost                                       int64   `json:"cost,omitempty"`
	CouponClickCount                           int64   `json:"coupon_click_count,omitempty"`
	CouponGetCost                              int64   `json:"coupon_get_cost,omitempty"`
	CouponGetCount                             int64   `json:"coupon_get_count,omitempty"`
	CouponGetPV                                int64   `json:"coupon_get_pv,omitempty"`
	CouponGetRate                              float64 `json:"coupon_get_rate,omitempty"`
	CouponIssueCount                           int64   `json:"coupon_issue_count,omitempty"`
	CouponPurchaseRate                         float64 `json:"coupon_purchase_rate,omitempty"`
	CouponUsageCost                            int64   `json:"coupon_usage_cost,omitempty"`
	CouponUsageNumber                          int64   `json:"coupon_usage_number,omitempty"`
	CouponUsageRate                            float64 `json:"coupon_usage_rate,omitempty"`
	CPC                                        int64   `json:"cpc,omitempty"`
	CreApplicationRate                         float64 `json:"cre_application_rate,omitempty"`
	CreditAmount                               int64   `json:"credit_amount,omitempty"`
	CreditCost                                 int64   `json:"credit_cost,omitempty"`
	CreditDedupPV                              int64   `json:"credit_dedup_pv,omitempty"`
	CreditPV                                   int64   `json:"credit_pv,omitempty"`
	CTR                                        float64 `json:"ctr,omitempty"`
	CustomAudienceID                           int64   `json:"custom_audience_id,omitempty"`
	CvsBubbleShareClkPV                        int64   `json:"cvs_bubble_share_clk_pv,omitempty"`
	CvsCpnVideoPlayDuration                    int64   `json:"cvs_cpn_video_play_duration,omitempty"`
	CvsCpnVideoPlayPV                          int64   `json:"cvs_cpn_video_play_pv,omitempty"`
	Date                                       string  `json:"date,omitempty"`
	DeepConversionsCost                        int64   `json:"deep_conversions_cost,omitempty"`
	DeepConversionsCount                       int64   `json:"deep_conversions_count,omitempty"`
	DeepConversionsRate                        float64 `json:"deep_conversions_rate,omitempty"`
	DeliverCost                                int64   `json:"deliver_cost,omitempty"`
	DeliverCount                               int64   `json:"deliver_count,omitempty"`
	DeliverRate                                float64 `json:"deliver_rate,omitempty"`
	DownloadCost                               int64   `json:"download_cost,omitempty"`
	DownloadCount                              int64   `json:"download_count,omitempty"`
	DownloadRate                               float64 `json:"download_rate,omitempty"`
	DurationKeyPagePerUser                     int64   `json:"duration_key_page_per_user,omitempty"`
	DurationOuterPerUser                       int64   `json:"duration_outer_per_user,omitempty"`
	DurationPerUser                            int64   `json:"duration_per_user,omitempty"`
	EffectLeadsPurchaseCost                    int64   `json:"effect_leads_purchase_cost,omitempty"`
	EffectLeadsPurchaseCount                   int64   `json:"effect_leads_purchase_count,omitempty"`
	EffectiveConsultCount                      int64   `json:"effective_consult_count,omitempty"`
	EffectiveCost                              int64   `json:"effective_cost,omitempty"`
	EffectiveLeadsCount                        int64   `json:"effective_leads_count,omitempty"`
	EffectivePhoneCount                        int64   `json:"effective_phone_count,omitempty"`
	EffectiveReserveCount                      int64   `json:"effective_reserve_count,omitempty"`
	EffectiveSeedingCost                       int64   `json:"effective_seeding_cost,omitempty"`
	EffectiveSeedingConverstionsRate           float64 `json:"effective_seeding_convertions_rate,omitempty"`
	EffectiveSeedingCount                      int64   `json:"effective_seeding_count,omitempty"`
	EffectiveSeedingRate                       float64 `json:"effective_seeding_rate,omitempty"`
	EngagePV                                   int64   `json:"engage_pv,omitempty"`
	ExternalFormReservationCount               int64   `json:"external_form_reservation_count,omitempty"`
	FinderTopicSliderManualPV                  int64   `json:"finder_topic_slider_manual_pv,omitempty"`
	FinderTopicSliderPV                        int64   `json:"finder_topic_slider_pv,omitempty"`
	FirOgConvAutoAcquisitionPV                 int64   `json:"fir_og_conv_auto_acquisition_pv,omitempty"`
	FirstDayAdPurArppuCost                     int64   `json:"first_day_ad_pur_arppu_cost,omitempty"`
	FirstDayAdPurArppuCost24h                  int64   `json:"first_day_ad_pur_arppu_cost_24h,omitempty"`
	FirstDayAdPurArppuCost24hPla               int64   `json:"first_day_ad_pur_arppu_cost_24h_pla,omitempty"`
	FirstDayFirstPayCount                      int64   `json:"first_day_first_pay_count,omitempty"`
	FirstDayFirstPayRate                       float64 `json:"first_day_first_pay_rate,omitempty"`
	FirstDayOrderAmount                        int64   `json:"first_day_order_amount,omitempty"`
	FirstDayOrderByClickAmount                 int64   `json:"first_day_order_by_click_amount,omitempty"`
	FirstDayOrderByClickCount                  int64   `json:"first_day_order_by_click_count,omitempty"`
	FirstDayOrderByDisplayAmount               int64   `json:"first_day_order_by_display_amount,omitempty"`
	FirstDayOrderByDisplayCount                int64   `json:"first_day_order_by_display_count,omitempty"`
	FirstDayOrderCount                         int64   `json:"first_day_order_count,omitempty"`
	FirstDayOrderROI                           float64 `json:"first_day_order_roi,omitempty"`
	FirstDayPayAmount                          int64   `json:"first_day_pay_amount,omitempty"`
	FirstDayPayAmountArppu                     int64   `json:"first_day_pay_amount_arppu,omitempty"`
	FirstDayPayAmountArpu                      int64   `json:"first_day_pay_amount_arpu,omitempty"`
	FirstDayPayCost                            int64   `json:"first_day_pay_cost,omitempty"`
	FirstDayPayCount                           int64   `json:"first_day_pay_count,omitempty"`
	FirstPayCost                               int64   `json:"first_pay_cost,omitempty"`
	FirstPayCount                              int64   `json:"first_pay_count,omitempty"`
	FirstPayRate                               float64 `json:"first_pay_rate,omitempty"`
	FollowBizAllDedupPV                        int64   `json:"follow_biz_all_dedup_pv,omitempty"`
	ForwardCost                                int64   `json:"forward_count,omitempty"`
	ForwardCount                               int64   `json:"forward_cost,omitempty"`
	FromFollowByClickCost                      int64   `json:"from_follow_by_click_cost,omitempty"`
	FromFollowByClickUV                        int64   `json:"from_follow_by_click_uv,omitempty"`
	FromFollowByDisplayCost                    int64   `json:"from_follow_by_display_cost,omitempty"`
	FromFollowByDisplayUV                      int64   `json:"from_follow_by_display_uv,omitempty"`
	FromFollowCost                             int64   `json:"from_follow_cost,omitempty"`
	FromFollowUV                               int64   `json:"from_follow_uv,omitempty"`
	GameAuthorizeCount                         int64   `json:"game_authorize_count,omitempty"`
	GameCreateRoleCount                        int64   `json:"game_create_role_count,omitempty"`
	GameTutorialFinishCount                    int64   `json:"game_tutorial_finish_count,omitempty"`
	GenderID                                   int64   `json:"gender_id,omitempty"`
	GuideToFollowPageInteractionAmount         int64   `json:"guide_to_follow_page_interaction_amount,omitempty"`
	GuideToFollowPageInteractionUsers          int64   `json:"guide_to_follow_page_interaction_users,omitempty"`
	GuideToFollowPageViewers                   int64   `json:"guide_to_follow_page_viewers,omitempty"`
	GuideToFollowPageViews                     int64   `json:"guide_to_follow_page_views,omitempty"`
	HotElementCTR                              float64 `json:"hot_element_ctr,omitempty"`
	IncomePV1dPla                              int64   `json:"income_pv_1d_pla,omitempty"`
	IncomePV24hPla                             int64   `json:"income_pv_24h_pla,omitempty"`
	IncomePVPla                                int64   `json:"income_pv_pla,omitempty"`
	IncomeROI1                                 float64 `json:"income_roi_1,omitempty"`
	IncomeROI14                                float64 `json:"income_roi_14,omitempty"`
	IncomeROI124h                              float64 `json:"income_roi_1_24h,omitempty"`
	IncomeROI124hPla                           float64 `json:"income_roi_1_24h_pla,omitempty"`
	IncomeROI3                                 float64 `json:"income_roi_3,omitempty"`
	IncomeROI30                                float64 `json:"income_roi_30,omitempty"`
	IncomeROI60                                float64 `json:"income_roi_60,omitempty"`
	IncomeROI7                                 float64 `json:"income_roi_7,omitempty"`
	IncomeVal1                                 int64   `json:"income_val_1,omitempty"`
	IncomeVal14                                int64   `json:"income_val_14,omitempty"`
	IncomeVal24h                               int64   `json:"income_val_24h,omitempty"`
	IncomeVal24hPla                            int64   `json:"income_val_24h_pla,omitempty"`
	IncomeVal24hPlaROI                         float64 `json:"income_val_24h_pla_roi,omitempty"`
	IncomeVal3                                 int64   `json:"income_val_3,omitempty"`
	IncomeVal30                                int64   `json:"income_val_30,omitempty"`
	IncomeVal60                                int64   `json:"income_val_60,omitempty"`
	IncomeVal7                                 int64   `json:"income_val_7,omitempty"`
	IneffectiveLeadsUV                         int64   `json:"ineffective_leads_uv,omitempty"`
	InstallCost                                int64   `json:"install_cost,omitempty"`
	InstallCount                               int64   `json:"install_count,omitempty"`
	InstallRate                                float64 `json:"install_rate,omitempty"`
	InsuranceDedupPV                           int64   `json:"insurance_dedup_pv,omitempty"`
	IntentionAfterPaymentDedupPV               int64   `json:"intention_after_payment_dedup_pv,omitempty"`
	InteractSuccPV                             int64   `json:"interact_succ_pv,omitempty"`
	JoinChatGroupAmount                        int64   `json:"join_chat_group_amount,omitempty"`
	JoinChatGroupCostByPeople                  int64   `json:"join_chat_group_cost_by_people,omitempty"`
	JoinChatGroupNumberOfPeople                int64   `json:"join_chat_group_number_of_people,omitempty"`
	KeyBehaviorConversionsCost                 int64   `json:"key_behavior_conversions_cost,omitempty"`
	KeyBehaviorConversionsCount                int64   `json:"key_behavior_conversions_count,omitempty"`
	KeyBehaviorConversionsRate                 float64 `json:"key_behavior_conversions_rate,omitempty"`
	KeyPageUV                                  int64   `json:"key_page_uv,omitempty"`
	KeyPageViewByClickCount                    int64   `json:"key_page_view_by_click_count,omitempty"`
	KeyPageViewByDisplayCount                  int64   `json:"key_page_view_by_display_count,omitempty"`
	KeyPageViewCost                            int64   `json:"key_page_view_cost,omitempty"`
	KeyPageViewCount                           int64   `json:"key_page_view_count,omitempty"`
	KeyPageViewRate                            float64 `json:"key_page_view_rate,omitempty"`
	LanButtonClickCost                         int64   `json:"lan_button_click_cost,omitempty"`
	LanButtonClickCount                        int64   `json:"lan_button_click_count,omitempty"`
	LanJumpButtonClickCost                     int64   `json:"lan_jump_button_click_cost,omitempty"`
	LanJumpButtonClickers                      int64   `json:"lan_jump_button_clickers,omitempty"`
	LanJumpButtonCTR                           float64 `json:"lan_jump_button_ctr,omitempty"`
	LanJumpButtonRate                          float64 `json:"lan_jump_button_rate,omitempty"`
	LandingCommodityDetailExpPV                int64   `json:"landing_commodity_detail_exp_pv,omitempty"`
	LeadsPurchaseUV                            int64   `json:"leads_purchase_uv,omitempty"`
	LiveStreamCommodityBubbleClkPV             int64   `json:"live_stream_commodity_bubble_clk_pv,omitempty"`
	LiveStreamCommodityShopBagClkPV            int64   `json:"live_stream_commodity_shop_bag_clk_pv,omitempty"`
	LiveStreamCommodityShopListExpPV           int64   `json:"live_stream_commodity_shop_list_exp_pv,omitempty"`
	LiveStreamCrtClickCnt                      int64   `json:"live_stream_crt_click_cnt,omitempty"`
	LoanDedupCost                              int64   `json:"loan_dedup_cost,omitempty"`
	LoanDedupPV                                int64   `json:"loan_dedup_pv,omitempty"`
	LoanQuotaOpenDedupCost                     int64   `json:"loan_quota_open_dedup_cost,omitempty"`
	LoanQuotaOpenDedupPV                       int64   `json:"loan_quota_open_dedup_pv,omitempty"`
	LotteryLeadsCost                           int64   `json:"lottery_leads_cost,omitempty"`
	LotteryLeadsCount                          int64   `json:"lottery_leads_count,omitempty"`
	LpStarPageExpPV                            int64   `json:"lp_star_page_exp_pv,omitempty"`
	MiniGameAdMonetizationAmount               int64   `json:"mini_game_ad_monetization_amount,omitempty"`
	MiniGameAdMonetizationAmountD14            int64   `json:"mini_game_ad_monetization_amount_d14,omitempty"`
	MiniGameAdMonetizationAmountD3             int64   `json:"mini_game_ad_monetization_amount_d3,omitempty"`
	MiniGameAdMonetizationAmountD7             int64   `json:"mini_game_ad_monetization_amount_d7,omitempty"`
	MiniGameAdMonetizationArpu                 int64   `json:"mini_game_ad_monetization_arpu,omitempty"`
	MiniGameAdMonetizationCost                 int64   `json:"mini_game_ad_monetization_cost,omitempty"`
	MiniGameAdMonetizationROI                  float64 `json:"mini_game_ad_monetization_roi,omitempty"`
	MiniGameBfCost                             int64   `json:"mini_game_bf_cost,omitempty"`
	MiniGameBfIncomeAmount                     int64   `json:"mini_game_bf_income_amount,omitempty"`
	MiniGameBfIncomeD1Amount                   int64   `json:"mini_game_bf_income_d1_amount,omitempty"`
	MiniGameBfIncomeD1Arpu                     int64   `json:"mini_game_bf_income_d1_arpu,omitempty"`
	MiniGameBfIncomeD1Cost                     int64   `json:"mini_game_bf_income_d1_cost,omitempty"`
	MiniGameBfIncomeD1ROI                      float64 `json:"mini_game_bf_income_d1_roi,omitempty"`
	MiniGameBfIncomePlaArpu                    int64   `json:"mini_game_bf_income_pla_arpu,omitempty"`
	MiniGameBfIncomePlaROI                     float64 `json:"mini_game_bf_income_pla_roi,omitempty"`
	MiniGameBfPurchaseArpu                     int64   `json:"mini_game_bf_purchase_arpu,omitempty"`
	MiniGameBfPurchaseCost                     int64   `json:"mini_game_bf_purchase_cost,omitempty"`
	MiniGameBfPurchaseD1Arpu                   int64   `json:"mini_game_bf_purchase_d1_arpu,omitempty"`
	MiniGameBfPurchaseD1Cost                   int64   `json:"mini_game_bf_purchase_d1_cost,omitempty"`
	MiniGameBfPurchaseD1ROI                    float64 `json:"mini_game_bf_purchase_d1_roi,omitempty"`
	MiniGameBfPurchaseROI                      float64 `json:"mini_game_bf_purchase_roi,omitempty"`
	MiniGameBfUV                               int64   `json:"mini_game_bf_uv,omitempty"`
	MiniGameCreateRoleCost                     int64   `json:"mini_game_create_role_cost,omitempty"`
	MiniGameCreateRoleRate                     float64 `json:"mini_game_create_role_rate,omitempty"`
	MiniGameCreateRoleUsers                    int64   `json:"mini_game_create_role_users,omitempty"`
	MiniGameD14PayCount                        int64   `json:"mini_game_d14_pay_count,omitempty"`
	MiniGameD30PayCount                        int64   `json:"mini_game_d30_pay_count,omitempty"`
	MiniGameD3PayCount                         int64   `json:"mini_game_d3_pay_count,omitempty"`
	MiniGameD7PayCount                         int64   `json:"mini_game_d7_pay_count,omitempty"`
	MiniGameFirstDayAdMonetizationAmount       int64   `json:"mini_game_first_day_ad_monetization_amount,omitempty"`
	MiniGameFirstDayAdPayingArpu               int64   `json:"mini_game_first_day_ad_paying_arpu,omitempty"`
	MiniGameFirstDayAdPayingCost               int64   `json:"mini_game_first_day_ad_paying_cost,omitempty"`
	MiniGameFirstDayPayingROI                  float64 `json:"mini_game_first_day_paying_roi,omitempty"`
	MiniGameFirstPayAmount                     int64   `json:"mini_game_first_pay_amount,omitempty"`
	MiniGameFirstPayPlaCost                    int64   `json:"mini_game_first_pay_pla_cost,omitempty"`
	MiniGameFirstPayingUsers                   int64   `json:"mini_game_first_paying_users,omitempty"`
	MiniGameIncomeROI1                         float64 `json:"mini_game_income_roi_1,omitempty"`
	MiniGameKeyPageViewCost                    int64   `json:"mini_game_key_page_view_cost,omitempty"`
	MiniGameKeyPageViewers                     int64   `json:"mini_game_key_page_viewers,omitempty"`
	MiniGameMixedMonetizationROID1             float64 `json:"mini_game_mixed_monetization_roi_d1,omitempty"`
	MiniGameMixedMonetizationROID14            float64 `json:"mini_game_mixed_monetization_roi_d14,omitempty"`
	MiniGameMixedMonetizationROID14ByReporting float64 `json:"mini_game_mixed_monetization_roi_d14_by_reporting,omitempty"`
	MiniGameMixedMonetizationROID1ByReporting  float64 `json:"mini_game_mixed_monetization_roi_d1_by_reporting,omitempty"`
	MiniGameMixedMonetizationROID3             float64 `json:"mini_game_mixed_monetization_roi_d3,omitempty"`
	MiniGameMixedMonetizationROID3ByReporting  float64 `json:"mini_game_mixed_monetization_roi_d3_by_reporting,omitempty"`
	MiniGameMixedMonetizationROID7             float64 `json:"mini_game_mixed_monetization_roi_d7,omitempty"`
	MiniGameMixedMonetizationROID7ByReporting  float64 `json:"mini_game_mixed_monetization_roi_d7_by_reporting,omitempty"`
	MiniGamePayD14PlaUV                        int64   `json:"mini_game_pay_d14_pla_uv,omitempty"`
	MiniGamePayD14ROI                          float64 `json:"mini_game_pay_d14_roi,omitempty"`
	MiniGamePayD14UV                           int64   `json:"mini_game_pay_d14_uv,omitempty"`
	MiniGamePayD1PlaRate                       float64 `json:"mini_game_pay_d1_pla_rate,omitempty"`
	MiniGamePayD30PlaUV                        int64   `json:"mini_game_pay_d30_pla_uv,omitempty"`
	MiniGamePayD30ROI                          float64 `json:"mini_game_pay_d30_roi,omitempty"`
	MiniGamePayD30UV                           int64   `json:"mini_game_pay_d30_uv,omitempty"`
	MiniGamePayD3PlaUV                         int64   `json:"mini_game_pay_d3_pla_uv,omitempty"`
	MiniGamePayD3ROI                           float64 `json:"mini_game_pay_d3_roi,omitempty"`
	MiniGamePayD3UV                            int64   `json:"mini_game_pay_d3_uv,omitempty"`
	MiniGamePayD7PlaUV                         int64   `json:"mini_game_pay_d7_pla_uv,omitempty"`
	MiniGamePayD7ROI                           float64 `json:"mini_game_pay_d7_roi,omitempty"`
	MiniGamePayD7UV                            int64   `json:"mini_game_pay_d7_uv,omitempty"`
	MiniGamePayingAmountD30                    int64   `json:"mini_game_paying_amount_d30,omitempty"`
	MiniGamePayingArpu                         int64   `json:"mini_game_paying_arpu,omitempty"`
	MiniGamePayingArpuD1                       int64   `json:"mini_game_paying_arpu_d1,omitempty"`
	MiniGamePayingUsersPlaD1                   int64   `json:"mini_game_paying_users_pla_d1,omitempty"`
	MiniGameRate                               float64 `json:"mini_game_rate,omitempty"`
	MiniGameRegisterCost                       int64   `json:"mini_game_register_cost,omitempty"`
	MiniGameRegisterRate                       float64 `json:"mini_game_register_rate,omitempty"`
	MiniGameRegisterUsers                      int64   `json:"mini_game_register_users,omitempty"`
	MiniGameRetentionD1                        int64   `json:"mini_game_retention_d1,omitempty"`
	MiniGameRetentionD1Cost                    int64   `json:"mini_game_retention_d1_cost,omitempty"`
	MiniGameRetentionD1Rate                    float64 `json:"mini_game_retention_d1_rate,omitempty"`
	Minigame24hPayArpu                         int64   `json:"minigame_24h_pay_arpu,omitempty"`
	Minigame24hPayROI                          float64 `json:"minigame_24h_pay_roi,omitempty"`
	Minigame24hPayUV                           int64   `json:"minigame_24h_pay_uv,omitempty"`
	Minigame3dIncomeCount                      int64   `json:"minigame_3d_income_count,omitempty"`
	Minigame3dIncomeROI                        float64 `json:"minigame_3d_income_roi,omitempty"`
	Minigame7dIncomeCount                      int64   `json:"minigame_7d_income_count,omitempty"`
	Minigame7dIncomeROI                        float64 `json:"minigame_7d_income_roi,omitempty"`
	MinigamePurchasePlaClk14dAmount            int64   `json:"minigame_purchase_pla_clk_14d_amount,omitempty"`
	MinigamePurchasePlaClk30dAmount            int64   `json:"minigame_purchase_pla_clk_30d_amount,omitempty"`
	MinigamePurchasePlaClk3dAmount             int64   `json:"minigame_purchase_pla_clk_3d_amount,omitempty"`
	MinigamePurchasePlaClk7dAmount             int64   `json:"minigame_purchase_pla_clk_7d_amount,omitempty"`
	MixPayActive14dROI                         float64 `json:"mix_pay_active_14d_roi,omitempty"`
	MixPayActive24hROI                         float64 `json:"mix_pay_active_24h_roi,omitempty"`
	MixPayActive3dROI                          float64 `json:"mix_pay_active_3d_roi,omitempty"`
	MixPayActive7dROI                          float64 `json:"mix_pay_active_7d_roi,omitempty"`
	MixPayArpu                                 int64   `json:"mix_pay_arpu,omitempty"`
	MixPayCost                                 int64   `json:"mix_pay_cost,omitempty"`
	MixPayROI                                  float64 `json:"mix_pay_roi,omitempty"`
	MixPayROIFirstDay                          float64 `json:"mix_pay_roi_first_day,omitempty"`
	MixPurchaseActive14dAmount                 int64   `json:"mix_purchase_active_14d_amount,omitempty"`
	MixPurchaseActive24hAmount                 int64   `json:"mix_purchase_active_24h_amount,omitempty"`
	MixPurchaseActive3dAmount                  int64   `json:"mix_purchase_active_3d_amount,omitempty"`
	MixPurchaseActive7dAmount                  int64   `json:"mix_purchase_active_7d_amount,omitempty"`
	MixPurchaseAmount                          int64   `json:"mix_purchase_amount,omitempty"`
	MixPurchaseAmountFirstDay                  int64   `json:"mix_purchase_amount_first_day,omitempty"`
	MixPurchaseDedupPV                         int64   `json:"mix_purchase_dedup_pv,omitempty"`
	MixedMonetizationROID1                     float64 `json:"mixed_monetization_roi_d1,omitempty"`
	MixedMonetizationROID14                    float64 `json:"mixed_monetization_roi_d14,omitempty"`
	MixedMonetizationROID3                     float64 `json:"mixed_monetization_roi_d3,omitempty"`
	MixedMonetizationROID7                     float64 `json:"mixed_monetization_roi_d7,omitempty"`
	NoInterestCount                            int64   `json:"no_interest_count,omitempty"`
	OpenAccountPV                              int64   `json:"open_account_pv,omitempty"`
	OpenAccountPVCost                          int64   `json:"open_account_pv_cost,omitempty"`
	Order24hAmount                             int64   `json:"order_24h_amount,omitempty"`
	Order24hByClickAmount                      int64   `json:"order_24h_by_click_amount,omitempty"`
	Order24hByClickCount                       int64   `json:"order_24h_by_click_count,omitempty"`
	Order24hByClickROI                         float64 `json:"order_24h_by_click_roi,omitempty"`
	Order24hByDisplayAmount                    int64   `json:"order_24h_by_display_amount,omitempty"`
	Order24hByDisplayCount                     int64   `json:"order_24h_by_display_count,omitempty"`
	Order24hByDisplayROI                       float64 `json:"order_24h_by_display_roi,omitempty"`
	Order24hCost                               int64   `json:"order_24h_cost,omitempty"`
	Order24hCount                              int64   `json:"order_24h_count,omitempty"`
	Order24hRate                               float64 `json:"order_24h_rate,omitempty"`
	Order24hROI                                float64 `json:"order_24h_roi,omitempty"`
	OrderAmount                                int64   `json:"order_amount,omitempty"`
	OrderByClickAmount                         int64   `json:"order_by_click_amount,omitempty"`
	OrderByClickCost                           int64   `json:"order_by_click_cost,omitempty"`
	OrderByClickCount                          int64   `json:"order_by_click_count,omitempty"`
	OrderByClickRate                           float64 `json:"order_by_click_rate,omitempty"`
	OrderByClickROI                            float64 `json:"order_by_click_roi,omitempty"`
	OrderByDisplayAmount                       int64   `json:"order_by_display_amount,omitempty"`
	OrderByDisplayCost                         int64   `json:"order_by_display_cost,omitempty"`
	OrderByDisplayCount                        int64   `json:"order_by_display_count,omitempty"`
	OrderByDisplayRate                         float64 `json:"order_by_display_rate,omitempty"`
	OrderByDisplayROI                          float64 `json:"order_by_display_roi,omitempty"`
	OrderClk23dAmount                          int64   `json:"order_clk2_3d_amount,omitempty"`
	OrderClk23dCost                            int64   `json:"order_clk2_3d_cost,omitempty"`
	OrderClk23dPV                              int64   `json:"order_clk2_3d_pv,omitempty"`
	OrderClk23dROI                             float64 `json:"order_clk2_3d_roi,omitempty"`
	OrderClk15dAmount                          int64   `json:"order_clk_15d_amount,omitempty"`
	OrderClk15dPV                              int64   `json:"order_clk_15d_pv,omitempty"`
	OrderClk15dROI                             float64 `json:"order_clk_15d_roi,omitempty"`
	OrderClk15dUnitPrice                       int64   `json:"order_clk_15d_unit_price,omitempty"`
	OrderClk30dAmount                          int64   `json:"order_clk_30d_amount,omitempty"`
	OrderClk30dPV                              int64   `json:"order_clk_30d_pv,omitempty"`
	OrderClk30dROI                             float64 `json:"order_clk_30d_roi,omitempty"`
	OrderClk30dUnitPrice                       int64   `json:"order_clk_30d_unit_price,omitempty"`
	OrderClk3dAmount                           int64   `json:"order_clk_3d_amount,omitempty"`
	OrderClk3dCost                             int64   `json:"order_clk_3d_cost,omitempty"`
	OrderClk3dPV                               int64   `json:"order_clk_3d_pv,omitempty"`
	OrderClk3dROI                              float64 `json:"order_clk_3d_roi,omitempty"`
	OrderClk7dAmount                           int64   `json:"order_clk_7d_amount,omitempty"`
	OrderClk7dPV                               int64   `json:"order_clk_7d_pv,omitempty"`
	OrderClk7dROI                              float64 `json:"order_clk_7d_roi,omitempty"`
	OrderClk7dUnitPrice                        int64   `json:"order_clk_7d_unit_price,omitempty"`
	OrderCost                                  int64   `json:"order_cost,omitempty"`
	OrderDedupClkPV                            int64   `json:"order_dedup_clk_pv,omitempty"`
	OrderDedupImpPV                            int64   `json:"order_dedup_imp_pv,omitempty"`
	OrderFirstDedupPV                          int64   `json:"order_first_dedup_pv,omitempty"`
	OrderFirstDedupPVCost                      int64   `json:"order_first_dedup_pv_cost,omitempty"`
	OrderFirstDedupPVRate                      float64 `json:"order_first_dedup_pv_rate,omitempty"`
	OrderFollow1dAmount                        int64   `json:"order_follow_1d_amount,omitempty"`
	OrderFollow1dPV                            int64   `json:"order_follow_1d_pv,omitempty"`
	OrderImp3dAmount                           int64   `json:"order_imp_3d_amount,omitempty"`
	OrderImp3dCost                             int64   `json:"order_imp_3d_cost,omitempty"`
	OrderImp3dPV                               int64   `json:"order_imp_3d_pv,omitempty"`
	OrderImp3dROI                              float64 `json:"order_imp_3d_roi,omitempty"`
	OrderNetAmount                             int64   `json:"order_net_amount,omitempty"`
	OrderNetPV                                 int64   `json:"order_net_pv,omitempty"`
	OrderNetPVCost                             int64   `json:"order_net_pv_cost,omitempty"`
	OrderNetROI                                float64 `json:"order_net_roi,omitempty"`
	OrderPV                                    int64   `json:"order_pv,omitempty"`
	OrderRate                                  float64 `json:"order_rate,omitempty"`
	OrderRefundActive24hPV                     int64   `json:"order_refund_active_24h_pv,omitempty"`
	OrderRefundActive24hRate                   float64 `json:"order_refund_active_24h_rate,omitempty"`
	OrderROI                                   float64 `json:"order_roi,omitempty"`
	OrderSettle24hCost                         int64   `json:"order_settle_24h_cost,omitempty"`
	OrderSettle24hPV                           int64   `json:"order_settle_24h_pv,omitempty"`
	OrderSettle24hRate                         float64 `json:"order_settle_24h_rate,omitempty"`
	OrderUnitPrice                             int64   `json:"order_unit_price,omitempty"`
	OrderUV                                    int64   `json:"order_uv,omitempty"`
	OsPlatform                                 string  `json:"os_platform,omitempty"`
	OverallLeadsPurchaseCount                  int64   `json:"overall_leads_purchase_count,omitempty"`
	OwnPageNaviCost                            int64   `json:"own_page_navi_cost,omitempty"`
	OwnPageNavigationCount                     int64   `json:"own_page_navigation_count,omitempty"`
	PageConsultCost                            int64   `json:"page_consult_cost,omitempty"`
	PageConsultCount                           int64   `json:"page_consult_count,omitempty"`
	PageConsultRate                            float64 `json:"page_consult_rate,omitempty"`
	PagePhoneCallDirectCost                    int64   `json:"page_phone_call_direct_cost,omitempty"`
	PagePhoneCallDirectCount                   int64   `json:"page_phone_call_direct_count,omitempty"`
	PagePhoneCallDirectRate                    float64 `json:"page_phone_call_direct_rate,omitempty"`
	PageReservationByClickCount                int64   `json:"page_reservation_by_click_count,omitempty"`
	PageReservationByDisplayCount              int64   `json:"page_reservation_by_display_count,omitempty"`
	PageReservationCost                        int64   `json:"page_reservation_cost,omitempty"`
	PageReservationCostWithPeople              int64   `json:"page_reservation_cost_with_people,omitempty"`
	PageReservationCount                       int64   `json:"page_reservation_count,omitempty"`
	PageReservationRate                        float64 `json:"page_reservation_rate,omitempty"`
	PageReservationROI                         float64 `json:"page_reservation_roi,omitempty"`
	PayBkActive14dROI                          float64 `json:"pay_bk_active_14d_roi,omitempty"`
	PayBkActive1dROI                           float64 `json:"pay_bk_active_1d_roi,omitempty"`
	PayBkActive24hROI                          float64 `json:"pay_bk_active_24h_roi,omitempty"`
	PayBkActive3dROI                           float64 `json:"pay_bk_active_3d_roi,omitempty"`
	PayBkActive7dROI                           float64 `json:"pay_bk_active_7d_roi,omitempty"`
	PayingUsersD1Cost                          int64   `json:"paying_users_d1_cost,omitempty"`
	PaymentAmountActivatedD14                  int64   `json:"payment_amount_activated_d14,omitempty"`
	PaymentAmountActivatedD3                   int64   `json:"payment_amount_activated_d3,omitempty"`
	PaymentAmountActivatedD30                  int64   `json:"payment_amount_activated_d30,omitempty"`
	PaymentAmountActivatedD7                   int64   `json:"payment_amount_activated_d7,omitempty"`
	PaymentCostActivatedD1                     int64   `json:"payment_cost_activated_d1,omitempty"`
	PlatformCouponClickCount                   int64   `json:"platform_coupon_click_count,omitempty"`
	PlatformKeyPageViewUserCount               int64   `json:"platform_key_page_view_user_count,omitempty"`
	PlatformPageNavigationCost                 int64   `json:"platform_page_navigation_cost,omitempty"`
	PlatformPageNavigationCount                int64   `json:"platform_page_navigation_count,omitempty"`
	PlatformPageViewCount                      int64   `json:"platform_page_view_count,omitempty"`
	PlatformPageViewRate                       float64 `json:"platform_page_view_rate,omitempty"`
	PlatformShopNavigationCost                 int64   `json:"platform_shop_navigation_cost,omitempty"`
	PlatformShopNavigationCount                int64   `json:"platform_shop_navigation_count,omitempty"`
	PotentialConsultCount                      int64   `json:"potential_consult_count,omitempty"`
	PotentialCustomerPhoneUV                   int64   `json:"potential_customer_phone_uv,omitempty"`
	PotentialPhoneCount                        int64   `json:"potential_phone_count,omitempty"`
	PotentialReserveCount                      int64   `json:"potential_reserve_count,omitempty"`
	PraiseCost                                 int64   `json:"praise_cost,omitempty"`
	PraiseCount                                int64   `json:"praise_count,omitempty"`
	PraiseUV                                   int64   `json:"praise_uv,omitempty"`
	PreCreditAmount                            int64   `json:"pre_credit_amount,omitempty"`
	PreCreditCost                              int64   `json:"pre_credit_cost,omitempty"`
	PreCreditDedupPV                           int64   `json:"pre_credit_dedup_pv,omitempty"`
	PreCreditPV                                int64   `json:"pre_credit_pv,omitempty"`
	PreviewConversionsCount                    int64   `json:"preview_conversions_count,omitempty"`
	PreviewDeepConversionsCount                int64   `json:"preview_deep_conversions_count,omitempty"`
	PurchaseActArpu                            int64   `json:"purchase_act_arpu,omitempty"`
	PurchaseActRate                            float64 `json:"purchase_act_rate,omitempty"`
	PurchaseAmount                             int64   `json:"purchase_amount,omitempty"`
	PurchaseAmountWithCoupon                   int64   `json:"purchase_amount_with_coupon,omitempty"`
	PurchaseAmountWithCouponCost               int64   `json:"purchase_amount_with_coupon_cost,omitempty"`
	PurchaseBkActive14dAmount                  int64   `json:"purchase_bk_active_14d_amount,omitempty"`
	PurchaseBkActive1dAmount                   int64   `json:"purchase_bk_active_1d_amount,omitempty"`
	PurchaseBkActive24hAmount                  int64   `json:"purchase_bk_active_24h_amount,omitempty"`
	PurchaseBkActive24hPV                      int64   `json:"purchase_bk_active_24h_pv,omitempty"`
	PurchaseBkActive3dAmount                   int64   `json:"purchase_bk_active_3d_amount,omitempty"`
	PurchaseBkActive7dAmount                   int64   `json:"purchase_bk_active_7d_amount,omitempty"`
	PurchaseBkAmount                           int64   `json:"purchase_bk_amount,omitempty"`
	PurchaseClk2Rate                           float64 `json:"purchase_clk2_rate,omitempty"`
	PurchaseClk15dPV                           int64   `json:"purchase_clk_15d_pv,omitempty"`
	PurchaseClk30dPV                           int64   `json:"purchase_clk_30d_pv,omitempty"`
	PurchaseClkAmount                          int64   `json:"purchase_clk_amount,omitempty"`
	PurchaseClkCost                            int64   `json:"purchase_clk_cost,omitempty"`
	PurchaseClkPV                              int64   `json:"purchase_clk_pv,omitempty"`
	PurchaseClkRate                            float64 `json:"purchase_clk_rate,omitempty"`
	PurchaseClkROI                             float64 `json:"purchase_clk_roi,omitempty"`
	PurchaseCost                               int64   `json:"purchase_cost,omitempty"`
	PurchaseDedupBkActive1dPV                  int64   `json:"purchase_dedup_bk_active_1d_pv,omitempty"`
	PurchaseDedupBkPV                          int64   `json:"purchase_dedup_bk_pv,omitempty"`
	PurchaseDedupClkPV                         int64   `json:"purchase_dedup_clk_pv,omitempty"`
	PurchaseDedupImpPV                         int64   `json:"purchase_dedup_imp_pv,omitempty"`
	PurchaseDedupPV                            int64   `json:"purchase_dedup_pv,omitempty"`
	PurchaseDedupRegActive1dPV                 int64   `json:"purchase_dedup_reg_active_1d_pv,omitempty"`
	PurchaseDedupRegPV                         int64   `json:"purchase_dedup_reg_pv,omitempty"`
	PurchaseFirstAllDedupTouch24hAmount        int64   `json:"purchase_first_all_dedup_touch_24h_amount,omitempty"`
	PurchaseFirstAllDedupTouch24hPV            int64   `json:"purchase_first_all_dedup_touch_24h_pv,omitempty"`
	PurchaseFirstAllDedupTouch24hUnitPrice     int64   `json:"purchase_first_all_dedup_touch_24h_unit_price,omitempty"`
	PurchaseImpAmount                          int64   `json:"purchase_imp_amount,omitempty"`
	PurchaseImpCost                            int64   `json:"purchase_imp_cost,omitempty"`
	PurchaseImpPV                              int64   `json:"purchase_imp_pv,omitempty"`
	PurchaseImpRate                            float64 `json:"purchase_imp_rate,omitempty"`
	PurchaseImpROI                             float64 `json:"purchase_imp_roi,omitempty"`
	PurchaseMemberCardDedupCost                int64   `json:"purchase_member_card_dedup_cost,omitempty"`
	PurchaseMemberCardDedupPV                  int64   `json:"purchase_member_card_dedup_pv,omitempty"`
	PurchaseMemberCardDedupRate                float64 `json:"purchase_member_card_dedup_rate,omitempty"`
	PurchaseMemberCardPV                       int64   `json:"purchase_member_card_pv,omitempty"`
	PurchasePlaActive14dAmount                 int64   `json:"purchase_pla_active_14d_amount,omitempty"`
	PurchasePlaActive14dPV                     int64   `json:"purchase_pla_active_14d_pv,omitempty"`
	PurchasePlaActive14dROI                    float64 `json:"purchase_pla_active_14d_roi,omitempty"`
	PurchasePlaActive1dAmount                  int64   `json:"purchase_pla_active_1d_amount,omitempty"`
	PurchasePlaActive1dROI                     float64 `json:"purchase_pla_active_1d_roi,omitempty"`
	PurchasePlaActive30dAmount                 int64   `json:"purchase_pla_active_30d_amount,omitempty"`
	PurchasePlaActive30dPV                     int64   `json:"purchase_pla_active_30d_pv,omitempty"`
	PurchasePlaActive30dROI                    float64 `json:"purchase_pla_active_30d_roi,omitempty"`
	PurchasePlaActive3dAmount                  int64   `json:"purchase_pla_active_3d_amount,omitempty"`
	PurchasePlaActive3dPV                      int64   `json:"purchase_pla_active_3d_pv,omitempty"`
	PurchasePlaActive3dROI                     float64 `json:"purchase_pla_active_3d_roi,omitempty"`
	PurchasePlaActive7dAmount                  int64   `json:"purchase_pla_active_7d_amount,omitempty"`
	PurchasePlaActive7dPV                      int64   `json:"purchase_pla_active_7d_pv,omitempty"`
	PurchasePlaActive7dROI                     float64 `json:"purchase_pla_active_7d_roi,omitempty"`
	PurchasePlaAmount                          int64   `json:"purchase_pla_amount,omitempty"`
	PurchasePlaBkActive1dAmount                int64   `json:"purchase_pla_bk_active_1d_amount,omitempty"`
	PurchasePlaClk1dAmount                     int64   `json:"purchase_pla_clk_1d_amount,omitempty"`
	PurchasePlaPV                              int64   `json:"purchase_pla_pv,omitempty"`
	PurchasePV                                 int64   `json:"purchase_pv,omitempty"`
	PurchaseRegActive14dAmount                 int64   `json:"purchase_reg_active_14d_amount,omitempty"`
	PurchaseRegActive1dAmount                  int64   `json:"purchase_reg_active_1d_amount,omitempty"`
	PurchaseRegActive1dPV                      int64   `json:"purchase_reg_active_1d_pv,omitempty"`
	PurchaseRegActive24hAmount                 int64   `json:"purchase_reg_active_24h_amount,omitempty"`
	PurchaseRegActive3dAmount                  int64   `json:"purchase_reg_active_3d_amount,omitempty"`
	PurchaseRegActive7dAmount                  int64   `json:"purchase_reg_active_7d_amount,omitempty"`
	PurchaseRegAmount                          int64   `json:"purchase_reg_amount,omitempty"`
	PurchaseRegArppu                           int64   `json:"purchase_reg_arppu,omitempty"`
	PurchaseRegArpu                            int64   `json:"purchase_reg_arpu,omitempty"`
	PurchaseROI                                float64 `json:"purchase_roi,omitempty"`
	QuitChatGroupAmount                        int64   `json:"quit_chat_group_amount,omitempty"`
	QuitChatGroupRate                          float64 `json:"quit_chat_group_rate,omitempty"`
	ReadCost                                   int64   `json:"read_cost,omitempty"`
	ReadCount                                  int64   `json:"read_count,omitempty"`
	RealCostAppAutoDownload                    int64   `json:"real_cost_app_auto_download,omitempty"`
	RegAllDedupPV                              int64   `json:"reg_all_dedup_pv,omitempty"`
	RegClickRatePla                            float64 `json:"reg_click_rate_pla,omitempty"`
	RegClkRate                                 float64 `json:"reg_clk_rate,omitempty"`
	RegCost                                    int64   `json:"reg_cost,omitempty"`
	RegCostPla                                 int64   `json:"reg_cost_pla,omitempty"`
	RegDedup1dPV                               int64   `json:"reg_dedup1d_pv,omitempty"`
	RegDedupCost                               int64   `json:"reg_dedup_cost,omitempty"`
	RegDedupPV                                 int64   `json:"reg_dedup_pv,omitempty"`
	RegPlaPV                                   int64   `json:"reg_pla_pv,omitempty"`
	RegPV                                      int64   `json:"reg_pv,omitempty"`
	RegionID                                   int64   `json:"region_id,omitempty"`
	RegisterByClickCount                       int64   `json:"register_by_click_count,omitempty"`
	RegisterByDisplayCount                     int64   `json:"register_by_display_count,omitempty"`
	RequestConversionsCost                     int64   `json:"request_conversions_cost,omitempty"`
	RequestConversionsCount                    int64   `json:"request_conversions_count,omitempty"`
	ReservationAmount                          int64   `json:"reservation_amount,omitempty"`
	ReservationCheckUV                         int64   `json:"reservation_check_uv,omitempty"`
	ReservationCheckUVCost                     int64   `json:"reservation_check_uv_cost,omitempty"`
	ReservationCheckUVRate                     float64 `json:"reservation_check_uv_rate,omitempty"`
	ReservationUV                              int64   `json:"reservation_uv,omitempty"`
	RetentionAllDedupPV                        int64   `json:"retention_all_dedup_pv,omitempty"`
	RetentionCost                              int64   `json:"retention_cost,omitempty"`
	RetentionCount                             int64   `json:"retention_count,omitempty"`
	RetentionD1UVRate                          float64 `json:"retention_d1_uv_rate,omitempty"`
	RetentionDedupPV                           int64   `json:"retention_dedup_pv,omitempty"`
	RetentionPlaDedupPV                        int64   `json:"retention_pla_dedup_pv,omitempty"`
	RetentionPlaDedupRate                      float64 `json:"retention_pla_dedup_rate,omitempty"`
	RetentionPlaPV                             int64   `json:"retention_pla_pv,omitempty"`
	RetentionRate                              float64 `json:"retention_rate,omitempty"`
	ROIActivatedD1                             float64 `json:"roi_activated_d1,omitempty"`
	ROIActivatedD14                            float64 `json:"roi_activated_d14,omitempty"`
	ROIActivatedD3                             float64 `json:"roi_activated_d3,omitempty"`
	ROIActivatedD30                            float64 `json:"roi_activated_d30,omitempty"`
	ROIActivatedD7                             float64 `json:"roi_activated_d7,omitempty"`
	ScanCodeAddFansCount                       int64   `json:"scan_code_add_fans_count,omitempty"`
	ScanCodeAddFansCountCost                   int64   `json:"scan_code_add_fans_count_cost,omitempty"`
	ScanCodeAddFansUV                          int64   `json:"scan_code_add_fans_uv,omitempty"`
	ScanCodeAddFansUVCost                      int64   `json:"scan_code_add_fans_uv_cost,omitempty"`
	ScanCodePlaPV                              int64   `json:"scan_code_pla_pv,omitempty"`
	ScanFollowCount                            int64   `json:"scan_follow_count,omitempty"`
	ScanFollowUserCost                         int64   `json:"scan_follow_user_cost,omitempty"`
	ScanFollowUserCount                        int64   `json:"scan_follow_user_count,omitempty"`
	ScanFollowUserRate                         float64 `json:"scan_follow_user_rate,omitempty"`
	SecOgConvAutoAcquisitionPV                 int64   `json:"sec_og_conv_auto_acquisition_pv,omitempty"`
	SecurityHighPriceOrderPV                   int64   `json:"security_high_price_order_pv,omitempty"`
	SecurityLowPriceOrderPV                    int64   `json:"security_low_price_order_pv,omitempty"`
	SecurityNegativeDedupPV                    int64   `json:"security_negative_dedup_pv,omitempty"`
	ShareFeedPV                                int64   `json:"share_feed_pv,omitempty"`
	ShareFriendPV                              int64   `json:"share_friend_pv,omitempty"`
	SignInAmount                               int64   `json:"sign_in_amount,omitempty"`
	SignInCost                                 int64   `json:"sign_in_cost,omitempty"`
	SignInCount                                int64   `json:"sign_in_count,omitempty"`
	SignInRate                                 float64 `json:"sign_in_rate,omitempty"`
	SignInROI                                  float64 `json:"sign_in_roi,omitempty"`
	SiteSet                                    string  `json:"site_set,omitempty"`
	SliderPV                                   int64   `json:"slider_pv,omitempty"`
	StayPay15dPV                               int64   `json:"stay_pay_15d_pv,omitempty"`
	StayPay30dPV                               int64   `json:"stay_pay_30d_pv,omitempty"`
	StayPay7dPV                                int64   `json:"stay_pay_7d_pv,omitempty"`
	StorePayAmountOff                          int64   `json:"store_pay_amount_off,omitempty"`
	StorePayPVOff                              int64   `json:"store_pay_pv_off,omitempty"`
	StorePayUVOff                              int64   `json:"store_pay_uv_off,omitempty"`
	StoreVisitor                               int64   `json:"store_visitor,omitempty"`
	ThousandDisplayPrice                       int64   `json:"thousand_display_price,omitempty"`
	ToolConsultCount                           int64   `json:"tool_consult_count,omitempty"`
	TryOutIntentionUV                          int64   `json:"try_out_intention_uv,omitempty"`
	TryOutUser                                 int64   `json:"try_out_user,omitempty"`
	ValidClickCount                            int64   `json:"valid_click_count,omitempty"`
	ValidLeadsUV                               int64   `json:"valid_leads_uv,omitempty"`
	ValidPhoneUV                               int64   `json:"valid_phone_uv,omitempty"`
	ValuableClickCost                          int64   `json:"valuable_click_cost,omitempty"`
	ValuableClickCount                         int64   `json:"valuable_click_count,omitempty"`
	ValuableClickRate                          float64 `json:"valuable_click_rate,omitempty"`
	VideoCommentCount                          int64   `json:"video_comment_count,omitempty"`
	VideoFollowCount                           int64   `json:"video_follow_count,omitempty"`
	VideoHeartCount                            int64   `json:"video_heart_count,omitempty"`
	VideoLiveCickCommodityCount                int64   `json:"video_live_cick_commodity_count,omitempty"`
	VideoLiveClickCommodityUserCount           int64   `json:"video_live_click_commodity_user_count,omitempty"`
	VideoLiveCommentCount                      int64   `json:"video_live_comment_count,omitempty"`
	VideoLiveCommentUserCount                  int64   `json:"video_live_comment_user_count,omitempty"`
	VideoLiveCommodityBubbleExpCount           int64   `json:"video_live_commodity_bubble_exp_count,omitempty"`
	VideoLiveExpCount                          int64   `json:"video_live_exp_count,omitempty"`
	VideoLiveHeartCount                        int64   `json:"video_live_heart_count,omitempty"`
	VideoLiveHeartUserCount                    int64   `json:"video_live_heart_user_count,omitempty"`
	VideoLiveShareCount                        int64   `json:"video_live_share_count,omitempty"`
	VideoLiveShareUserCount                    int64   `json:"video_live_share_user_count,omitempty"`
	VideoLiveSubscribeCount                    int64   `json:"video_live_subscribe_count,omitempty"`
	VideoPlayCount                             int64   `json:"video_play_count,omitempty"`
	ViewCommodityPageUV                        int64   `json:"view_commodity_page_uv,omitempty"`
	ViewCount                                  int64   `json:"view_count,omitempty"`
	ViewDramaDedupPV                           int64   `json:"view_drama_dedup_pv,omitempty"`
	ViewDramaDedupPVCost                       int64   `json:"view_drama_dedup_pv_cost,omitempty"`
	ViewDramaDedupPVRate                       float64 `json:"view_drama_dedup_pv_rate,omitempty"`
	ViewUserCount                              int64   `json:"view_user_count,omitempty"`
	VisitStoreDedupPVCost                      int64   `json:"visit_store_dedup_pv_cost,omitempty"`
	VisitStoreFirstDedupPV                     int64   `json:"visit_store_first_dedup_pv,omitempty"`
	VisitStorePV                               int64   `json:"visit_store_pv,omitempty"`
	VisitStorePVCost                           int64   `json:"visit_store_pv_cost,omitempty"`
	WebCommodityPageViewCost                   int64   `json:"web_commodity_page_view_cost,omitempty"`
	WebCommodityPageViewRate                   float64 `json:"web_commodity_page_view_rate,omitempty"`
	WechatAddFansAfterV30sDedupPV              int64   `json:"wechat_add_fans_after_v30s_dedup_pv,omitempty"`
	WechatLocalPayAmount                       int64   `json:"wechat_local_pay_amount,omitempty"`
	WechatLocalPayCount                        int64   `json:"wechat_local_pay_count,omitempty"`
	WechatLocalPayROI                          float64 `json:"wechat_local_pay_roi,omitempty"`
	WechatLocalPayuserCount                    int64   `json:"wechat_local_payuser_count,omitempty"`
	WecomAddPersonalDedupPV                    int64   `json:"wecom_add_personal_dedup_pv,omitempty"`
	WecomAddPersonalDedupPVCost                int64   `json:"wecom_add_personal_dedup_pv_cost,omitempty"`
	WithdrawDedupPV                            int64   `json:"withdraw_dedup_pv,omitempty"`
	WithdrawDepositAmount                      int64   `json:"withdraw_deposit_amount,omitempty"`
	WithdrawDepositClk7dAmount                 int64   `json:"withdraw_deposit_clk_7d_amount,omitempty"`
	WithdrawDepositClk7dROI                    float64 `json:"withdraw_deposit_clk_7d_roi,omitempty"`
	WithdrawDepositPV                          int64   `json:"withdraw_deposit_pv,omitempty"`
	ZoneHeaderClickCount                       int64   `json:"zone_header_click_count,omitempty"`
	ZoneHeaderLiveClickCnt                     int64   `json:"zone_header_live_click_cnt,omitempty"`
}
