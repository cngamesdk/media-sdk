package model

import (
	"errors"
	"time"
)

type ReportCustomConfigGetReq struct {
	accessTokenReq
	AdvertiserID int64    `json:"advertiser_id"`         // 客户ID (必填)
	DataTopics   []string `json:"data_topics,omitempty"` // 数据主题查询列表
}

// 常量定义 - 数据主题
const (
	DataTopicBasic         = "BASIC_DATA"          // 营销基础数据
	DataTopicQuery         = "QUERY_DATA"          // 搜索词数据
	DataTopicBidword       = "BIDWORD_DATA"        // 关键词数据
	DataTopicMaterial      = "MATERIAL_DATA"       // 素材数据
	DataTopicProduct       = "PRODUCT_DATA"        // 产品数据
	DataTopicOneKeyBoost   = "ONE_KEY_BOOST_DATA"  // 一键起量（巨量营销升级版）
	DataTopicDMP           = "DMP_DATA"            // 人群包数据
	DataTopicVideoDuration = "VIDEO_DURATION_DATA" // 视频分秒数据
	DataTopicMaterialBoost = "MATERIAL_BOOST_DATA" // 素材一键起量
)

// 所有数据主题列表
var AllDataTopics = []string{
	DataTopicBasic,
	DataTopicQuery,
	DataTopicBidword,
	DataTopicMaterial,
	DataTopicProduct,
	DataTopicOneKeyBoost,
	DataTopicDMP,
	DataTopicVideoDuration,
	DataTopicMaterialBoost,
}

// Validate 验证查询参数
func (p *ReportCustomConfigGetReq) Validate() error {
	if validateErr := p.accessTokenReq.Validate(); validateErr != nil {
		return validateErr
	}
	// 验证客户ID
	if p.AdvertiserID == 0 {
		return errors.New("advertiser_id为必填")
	}

	// 验证数据主题
	if err := p.validateDataTopics(); err != nil {
		return err
	}

	return nil
}

// validateDataTopics 验证数据主题列表
func (p *ReportCustomConfigGetReq) validateDataTopics() error {
	if len(p.DataTopics) == 0 {
		return nil // 不传表示查询所有数据
	}

	// 创建有效数据主题的映射
	validTopics := make(map[string]bool)
	for _, topic := range AllDataTopics {
		validTopics[topic] = true
	}

	// 验证每个数据主题
	for _, topic := range p.DataTopics {
		if !validTopics[topic] {
			return errors.New("data_topics包含无效值，请参考文档中的允许值")
		}
	}

	return nil
}

// DataTopicConfigResponse 数据主题配置响应
type ReportCustomConfigGetResp struct {
	List []*DataTopicConfig `json:"list"` // 数据主题配置列表
}

// DataTopicConfig 数据主题配置
type DataTopicConfig struct {
	DataTopic  string       `json:"data_topic"` // 数据主体
	Dimensions []*Dimension `json:"dimensions"` // 维度列表
	Metrics    []*Metric    `json:"metrics"`    // 指标列表
}

// Metric 指标信息
type Metric struct {
	Field         string   `json:"field"`                    // 字段
	Name          string   `json:"name"`                     // 字段名称
	Description   string   `json:"description"`              // 字段描述
	ExclusionDims []string `json:"exclusion_dims,omitempty"` // 与指标互斥的维度
}

// Dimension 维度信息
type Dimension struct {
	Field            string        `json:"field"`                       // 维度字段
	Name             string        `json:"name"`                        // 维度名称
	Description      string        `json:"description"`                 // 维度描述
	SortAble         bool          `json:"sort_able"`                   // 是否支持排序
	FilterAble       bool          `json:"filter_able"`                 // 是否支持筛选
	FilterConfig     *FilterConfig `json:"filter_config,omitempty"`     // 筛选条件
	ExclusionDims    []string      `json:"exclusion_dims,omitempty"`    // 与维度互斥的维度列表
	ExclusionMetrics []string      `json:"exclusion_metrics,omitempty"` // 与维度互斥的指标列表
}

// FilterConfig 筛选条件配置
type FilterConfig struct {
	Type       int           `json:"type,omitempty"`        // 筛选字段类型
	Operator   int           `json:"operator,omitempty"`    // 筛选方式
	ValueLimit int           `json:"value_limit,omitempty"` // 筛选字段输入数量上限
	RangeValue []*RangeValue `json:"range_value,omitempty"` // 筛选字段枚举列表
}

// RangeValue 筛选字段枚举值
type RangeValue struct {
	Label string `json:"label"` // 筛选字段名称
	Value string `json:"value"` // 筛选字段具体值
}

type ReportCustomGetReq struct {
	accessTokenReq
	AdvertiserID int64              `json:"advertiser_id"`     // 客户ID (必填)
	DataTopic    string             `json:"data_topic"`        // 数据主题，默认BASIC_DATA
	Dimensions   []string           `json:"dimensions"`        // 维度列表 (必填)
	Metrics      []string           `json:"metrics"`           // 指标列表 (必填)
	Filters      []*FilterCondition `json:"filters,omitempty"` // 过滤条件列表
	StartTime    string             `json:"start_time"`        // 开始时间 (必填)
	EndTime      string             `json:"end_time"`          // 结束时间 (必填)
	OrderBy      []OrderBy          `json:"order_by"`          // 排序
	PageInfoReq
}

const (
	OrderByAsc  = "ASC"  // 升序
	OrderByDesc = "DESC" // 降序
)

type OrderBy struct {
	Field string `json:"field,omitempty"` // 排序字段
	Type  string `json:"type,omitempty"`  // 排序类型
}

// FilterCondition 过滤条件
type FilterCondition struct {
	Field    string   `json:"field"`    // 过滤的消耗指标字段 (条件必填)
	Type     int      `json:"type"`     // 字段类型 (条件必填)
	Operator int      `json:"operator"` // 处理方式 (条件必填)
	Values   []string `json:"values"`   // 过滤字段具体值 (条件必填)
}

// 常量定义 - 字段类型
const (
	FieldTypeEnum   = 1 // 固定枚举值
	FieldTypeInput  = 2 // 固定输入值
	FieldTypeNumber = 3 // 数值类型
)

// 常量定义 - 处理方式
const (
	OperatorEqual              = 1  // 等于
	OperatorLessThan           = 2  // 小于
	OperatorLessThanOrEqual    = 3  // 小于等于
	OperatorGreaterThan        = 4  // 大于
	OperatorGreaterThanOrEqual = 5  // 大于等于
	OperatorNotEqual           = 6  // 不等于
	OperatorContains           = 7  // 包含
	OperatorNotContains        = 8  // 不包含
	OperatorRange              = 9  // 范围查询
	OperatorMultiValueInclude  = 10 // 多个值匹配包含
	OperatorMultiValueAll      = 11 // 多个值匹配都要包含
)

func (p *ReportCustomGetReq) Format() {
	p.accessTokenReq.Format()
	p.PageInfoReq.Format()
}

// 时间格式
const TimeFormat = "2006-01-02 15:04:05"

// Validate 验证查询参数
func (p *ReportCustomGetReq) Validate() error {
	if validateErr := p.accessTokenReq.Validate(); validateErr != nil {
		return validateErr
	}
	// 1. 验证客户ID
	if p.AdvertiserID == 0 {
		return errors.New("advertiser_id为必填")
	}

	// 2. 设置默认值
	p.setDefaults()

	// 3. 验证数据主题
	if err := p.validateDataTopic(); err != nil {
		return err
	}

	// 4. 验证维度列表
	if len(p.Dimensions) == 0 {
		return errors.New("dimensions为必填")
	}

	// 5. 验证指标列表
	if len(p.Metrics) == 0 {
		return errors.New("metrics为必填")
	}

	// 6. 验证过滤条件
	if err := p.validateFilters(); err != nil {
		return err
	}

	// 7. 验证时间
	if err := p.validateTimeRange(); err != nil {
		return err
	}

	return nil
}

// setDefaults 设置默认值
func (p *ReportCustomGetReq) setDefaults() {
	if p.DataTopic == "" {
		p.DataTopic = DataTopicBasic
	}
}

// validateDataTopic 验证数据主题
func (p *ReportCustomGetReq) validateDataTopic() error {
	validTopics := make(map[string]bool)
	for _, topic := range AllDataTopics {
		validTopics[topic] = true
	}

	if !validTopics[p.DataTopic] {
		return errors.New("data_topic值无效，请参考文档中的枚举值")
	}

	return nil
}

// validateFilters 验证过滤条件
func (p *ReportCustomGetReq) validateFilters() error {
	for i, filter := range p.Filters {
		if err := filter.validate(); err != nil {
			return errors.New("filters[" + string(rune(i)) + "]验证失败: " + err.Error())
		}
	}
	return nil
}

// validate 验证单个过滤条件
func (f *FilterCondition) validate() error {
	// 验证字段
	if f.Field == "" {
		return errors.New("field为条件必填")
	}

	// 验证字段类型
	if f.Type < FieldTypeEnum || f.Type > FieldTypeNumber {
		return errors.New("type值无效，允许值：1(固定枚举值)、2(固定输入值)、3(数值类型)")
	}

	// 验证处理方式
	if f.Operator < OperatorEqual || f.Operator > OperatorMultiValueAll {
		return errors.New("operator值无效，允许值：1-11")
	}

	// 验证过滤值
	if len(f.Values) == 0 {
		return errors.New("values为条件必填")
	}

	// 根据字段类型和处理方式进一步验证
	switch f.Type {
	case FieldTypeNumber:
		// 数值类型验证
		if f.Operator == OperatorContains || f.Operator == OperatorNotContains {
			return errors.New("数值类型不支持包含/不包含操作")
		}
	}

	return nil
}

// validateTimeRange 验证时间范围
func (p *ReportCustomGetReq) validateTimeRange() error {
	// 验证开始时间
	if p.StartTime == "" {
		return errors.New("start_time为必填")
	}

	start, err := time.Parse(TimeFormat, p.StartTime)
	if err != nil {
		return errors.New("start_time格式错误，应为yyyy-MM-dd hh:mm:ss")
	}

	// 验证结束时间
	if p.EndTime == "" {
		return errors.New("end_time为必填")
	}

	end, err := time.Parse(TimeFormat, p.EndTime)
	if err != nil {
		return errors.New("end_time格式错误，应为yyyy-MM-dd hh:mm:ss")
	}

	// 验证开始时间不能晚于结束时间
	if start.After(end) {
		return errors.New("开始时间不能晚于结束时间")
	}

	return nil
}

// ReportDataRow 报表数据行
type ReportDataRow struct {
	Dimensions map[string]string `json:"dimensions"` // 维度数据
	Metrics    map[string]string `json:"metrics"`    // 指标数据
}

// ReportCustomGetResp 报表数据响应
type ReportCustomGetResp struct {
	Rows         []*ReportDataRow  `json:"rows"`          // 指标数据列表
	TotalMetrics map[string]string `json:"total_metrics"` // 指标汇总数据
	PageInfoContainerResp
}
