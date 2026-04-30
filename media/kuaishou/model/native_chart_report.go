package model

import "errors"

// NativeChartReportSearchParam 原生报表概览请求参数
type NativeChartReportSearchParam struct {
	AuthorId       []int64 `json:"author_id"`        // 快手号ID
	CampaignType   []int   `json:"campaign_type"`    // 营销目标：2-提升应用安装，5-收集销售线索，7-提升应用活跃，19-快手小程序推广，30-短剧推广
	OcpcActionType []int   `json:"ocpc_action_type"` // 转化目标：180-激活，53-表单优化，190-付费，191-首日ROI
	KolUserType    []int   `json:"kol_user_type"`    // 原生广告类型：1-普通快手号，2-蓝V，3-聚星达人
	ReportEndDay   int64   `json:"report_end_day"`   // 结束时间，时间戳毫秒，必填
	ReportStartDay int64   `json:"report_start_day"` // 开始时间，时间戳毫秒，必填
}

// NativeChartReportReq 原生效果数据整体概览请求
type NativeChartReportReq struct {
	accessTokenReq
	AdvertiserId int64                        `json:"advertiser_id"` // 广告主账号ID，必填
	SearchParam  NativeChartReportSearchParam `json:"search_param"`  // 原生报表概览请求参数，必填
}

func (receiver *NativeChartReportReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *NativeChartReportReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.SearchParam.ReportStartDay <= 0 {
		err = errors.New("search_param.report_start_day is empty")
		return
	}
	if receiver.SearchParam.ReportEndDay <= 0 {
		err = errors.New("search_param.report_end_day is empty")
		return
	}
	return
}

// NativeReportSummary 原生报表数据概览条目
type NativeReportSummary struct {
	Type                  string  `json:"type"`                     // 指标类型：totalCharge-花费，paiedAmt-整体付费金额，conversion-整体转化数
	Yoy                   float64 `json:"yoy"`                      // 环比
	Direct                float64 `json:"direct"`                   // 直接数据（仅转化数、付费金额区分）
	Indirect              float64 `json:"indirect"`                 // 间接数据（仅转化数、付费金额区分）
	NativeSoftChargeRatio float64 `json:"native_soft_charge_ratio"` // 原生消耗增量收益（仅花费有此指标）
	AttributeWindow       int     `json:"attribute_window"`         // 归因窗口：1=1天，7=7天
	SumValue              float64 `json:"sum_value"`                // 总体值
	NativeConversionRatio float64 `json:"native_conversion_ratio"`  // 原生转化增量收益
}

// NativeReportConversionChart 转化数分日柱状图
type NativeReportConversionChart struct {
	Date            int64   `json:"date"`             // 日期，时间戳毫秒
	AttributeWindow int     `json:"attribute_window"` // 归因窗口：1=1天，7=7天
	SumValue        float64 `json:"sum_value"`        // 总体值
}

// NativeReportFunnelInfo 漏斗数据
type NativeReportFunnelInfo struct {
	AdShow          int64 `json:"ad_show"`          // 曝光数
	ActionbarClick  int64 `json:"actionbar_click"`  // 行为数
	ConversionCnt   int64 `json:"conversion_cnt"`   // 直接转化数
	AttributeWindow int   `json:"attribute_window"` // 归因窗口：1=1天，7=7天
}

// NativeChartReportResp 原生效果数据整体概览响应数据（仅data部分）
type NativeChartReportResp struct {
	SummaryNature   []NativeReportSummary         `json:"summary_nature"`   // 自然流量数据
	Summary         []NativeReportSummary         `json:"summary"`          // 数据概览
	ConversionChart []NativeReportConversionChart `json:"conversion_chart"` // 转化数分日柱状图
	Funnel          []NativeReportFunnelInfo      `json:"funnel"`           // 漏斗数据
}
