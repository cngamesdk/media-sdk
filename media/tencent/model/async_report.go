package model

import (
	"errors"
)

// ========== 创建异步报表任务 ==========
// https://developers.e.qq.com/v3.0/docs/api/async_reports/add

// AsyncReportsAddReq 创建异步报表任务请求
type AsyncReportsAddReq struct {
	GlobalReq
	AccountID      int64                `json:"account_id,omitempty"`      // 推广帐号id
	TaskName       string               `json:"task_name"`                 // 任务名称 (必填)
	ReportFields   []string             `json:"report_fields"`             // 异步报表返回字段 (必填)
	Level          string               `json:"level"`                     // 异步报表类型级别 (必填)
	Filtering      []*AsyncReportFilter `json:"filtering,omitempty"`       // 过滤条件
	TimeLine       string               `json:"time_line,omitempty"`       // 时间口径
	GroupBy        []string             `json:"group_by,omitempty"`        // 聚合参数
	Granularity    string               `json:"granularity"`               // 异步报表粒度 (必填)
	Date           string               `json:"date"`                      // 日期，格式：YYYY-MM-DD (必填)
	OrganizationID int64                `json:"organization_id,omitempty"` // 业务单元id
}

// AsyncReportFilter 过滤条件
type AsyncReportFilter struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)
	Operator string   `json:"operator"` // 操作符 (必填)
	Values   []string `json:"values"`   // 字段取值 (必填)
}

// 异步报表粒度常量
const (
	AsyncReportGranularityHourly = "HOURLY"
	AsyncReportGranularityDaily  = "DAILY"
)

// 异步报表过滤字段常量
const (
	AsyncReportFilterAdgroupId         = "adgroup_id"
	AsyncReportFilterDynamicCreativeId = "dynamic_creative_id"
	AsyncReportFilterComponentId       = "component_id"
	AsyncReportFilterComponentType     = "component_type"
	AsyncReportFilterBidwordId         = "bidword_id"
	AsyncReportFilterImageId           = "image_id"
	AsyncReportFilterVideoId           = "video_id"
	AsyncReportFilterChannelId         = "channel_id"
	AsyncReportFilterUnionPositionId   = "union_position_id"
	AsyncReportFilterLandingPageType   = "landing_page_type"
	AsyncReportFilterLandingPageId     = "landing_page_id"
	AsyncReportFilterMd5               = "md5"
	AsyncReportFilterHour              = "hour"
)

// 异步报表过滤操作符常量
const (
	AsyncReportOperatorEquals        = "EQUALS"
	AsyncReportOperatorContains      = "CONTAINS"
	AsyncReportOperatorLessEquals    = "LESS_EQUALS"
	AsyncReportOperatorLess          = "LESS"
	AsyncReportOperatorGreaterEquals = "GREATER_EQUALS"
	AsyncReportOperatorGreater       = "GREATER"
	AsyncReportOperatorIn            = "IN"
	AsyncReportOperatorNotEquals     = "NOT_EQUALS"
)

// 异步报表限制常量
const (
	AsyncReportMinTaskNameLength    = 1
	AsyncReportMaxTaskNameLength    = 120
	AsyncReportMinReportFieldsCount = 1
	AsyncReportMaxReportFieldsCount = 1024
	AsyncReportMinReportFieldLength = 1
	AsyncReportMaxReportFieldLength = 64
	AsyncReportMinFilteringCount    = 1
	AsyncReportMaxFilteringCount    = 5
	AsyncReportMinValuesCount       = 1
	AsyncReportMaxValuesCount       = 100
	AsyncReportMinValuesLength      = 1
	AsyncReportMaxValuesLength      = 64
	AsyncReportMinGroupByCount      = 1
	AsyncReportMaxGroupByCount      = 5
	AsyncReportMaxGroupByLength     = 255
	AsyncReportDateLength           = 10
	AsyncReportMaxOrganizationID    = 9999999999
)

func (p *AsyncReportsAddReq) Format() {
	p.GlobalReq.Format()
}

func (p *AsyncReportsAddReq) Validate() error {
	// 验证全局参数
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	// 验证task_name (必填)
	if p.TaskName == "" {
		return errors.New("task_name为必填")
	}
	if len(p.TaskName) < AsyncReportMinTaskNameLength || len(p.TaskName) > AsyncReportMaxTaskNameLength {
		return errors.New("task_name长度必须在1-120字节之间")
	}

	// 验证report_fields (必填)
	if len(p.ReportFields) < AsyncReportMinReportFieldsCount || len(p.ReportFields) > AsyncReportMaxReportFieldsCount {
		return errors.New("report_fields数组长度必须在1-1024之间")
	}
	for _, f := range p.ReportFields {
		if len(f) < AsyncReportMinReportFieldLength || len(f) > AsyncReportMaxReportFieldLength {
			return errors.New("report_fields字段长度必须在1-64字节之间")
		}
	}

	// 验证level (必填)
	if p.Level == "" {
		return errors.New("level为必填")
	}
	if !isValidAsyncReportLevel(p.Level) {
		return errors.New("level值无效")
	}

	// 验证filtering
	if len(p.Filtering) > 0 {
		if len(p.Filtering) < AsyncReportMinFilteringCount || len(p.Filtering) > AsyncReportMaxFilteringCount {
			return errors.New("filtering数组长度必须在1-5之间")
		}
		for _, f := range p.Filtering {
			if f.Field == "" {
				return errors.New("filtering.field为必填")
			}
			if f.Operator == "" {
				return errors.New("filtering.operator为必填")
			}
			if len(f.Values) < AsyncReportMinValuesCount || len(f.Values) > AsyncReportMaxValuesCount {
				return errors.New("filtering.values数组长度必须在1-100之间")
			}
			for _, v := range f.Values {
				if len(v) < AsyncReportMinValuesLength || len(v) > AsyncReportMaxValuesLength {
					return errors.New("filtering.values字段长度必须在1-64字节之间")
				}
			}
		}
	}

	// 验证time_line
	if p.TimeLine != "" {
		if p.TimeLine != TimeLineRequestTime && p.TimeLine != TimeLineReportingTime && p.TimeLine != TimeLineActiveTime {
			return errors.New("time_line值无效，允许值：REQUEST_TIME、REPORTING_TIME、ACTIVE_TIME")
		}
	}

	// 验证group_by
	if len(p.GroupBy) > 0 {
		if len(p.GroupBy) < AsyncReportMinGroupByCount || len(p.GroupBy) > AsyncReportMaxGroupByCount {
			return errors.New("group_by数组长度必须在1-5之间")
		}
		for _, g := range p.GroupBy {
			if len(g) > AsyncReportMaxGroupByLength {
				return errors.New("group_by字段长度不能超过255字节")
			}
		}
	}

	// 验证granularity (必填)
	if p.Granularity == "" {
		return errors.New("granularity为必填")
	}
	if p.Granularity != AsyncReportGranularityHourly && p.Granularity != AsyncReportGranularityDaily {
		return errors.New("granularity值无效，允许值：HOURLY、DAILY")
	}

	// 验证date (必填)
	if p.Date == "" {
		return errors.New("date为必填")
	}
	if len(p.Date) != AsyncReportDateLength {
		return errors.New("date长度必须为10字节")
	}

	// 验证organization_id
	if p.OrganizationID < 0 || p.OrganizationID > AsyncReportMaxOrganizationID {
		return errors.New("organization_id必须在0-9999999999之间")
	}

	return nil
}

// isValidAsyncReportLevel 验证异步报表级别
func isValidAsyncReportLevel(level string) bool {
	validLevels := map[string]bool{
		ReportLevelAdvertiser:              true,
		ReportLevelAdgroup:                 true,
		ReportLevelBidword:                 true,
		ReportLevelQueryword:               true,
		"REPORT_LEVEL_AGE":                 true,
		"REPORT_LEVEL_GENDER":              true,
		"REPORT_LEVEL_REGION":              true,
		"REPORT_LEVEL_CITY":                true,
		ReportLevelDynamicCreative:         true,
		ReportLevelComponent:               true,
		ReportLevelMaterialImage:           true,
		ReportLevelMaterialVideo:           true,
		ReportLevelChannel:                 true,
		"REPORT_LEVEL_LANDING_PAGE":        true,
		ReportLevelMarketingAsset:          true,
		"REPORT_LEVEL_AD_UNION":            true,
		ReportLevelProductCatalog:          true,
		ReportLevelProject:                 true,
		ReportLevelProjectCreative:         true,
		"REPORT_LEVEL_OS":                  true,
		ReportLevelProductCreativeTemplate: true,
	}
	return validLevels[level]
}

// AsyncReportsAddResp 创建异步报表任务响应
type AsyncReportsAddResp struct {
	TaskID int64 `json:"task_id,omitempty"` // 任务id
}
