package model

import "errors"

// NativeNatureChartReportReq 原生报表披露自然流量数据整体概览请求
type NativeNatureChartReportReq struct {
	accessTokenReq
	SearchParam  NativeChartReportSearchParam `json:"search_param"`  // 原生报表概览请求参数，必填
	AdvertiserId int64                        `json:"advertiser_id"` // 广告主账号ID，必填
}

func (receiver *NativeNatureChartReportReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *NativeNatureChartReportReq) Validate() (err error) {
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

// NatureReportSummary 原生报表自然流量数据概览条目
type NatureReportSummary struct {
	SumValue float64 `json:"sum_value"` // 总体数据
	Type     string  `json:"type"`      // 指标类型：totalConversion-整体转化数，adConversion-原生广告经营，natureConversion-自然经营
	Yoy      float64 `json:"yoy"`       // 同比
}

// NativeNatureChartReportResp 原生报表披露自然流量数据整体概览响应数据（仅data部分）
type NativeNatureChartReportResp struct {
	SummaryOperation  []NatureReportSummary `json:"summary_operation"`  // 经营数据
	SummaryCost       []NatureReportSummary `json:"summary_cost"`       // 成本数据
	SummaryConversion []NatureReportSummary `json:"summary_conversion"` // 转化数据
}
