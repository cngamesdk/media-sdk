package model

import "errors"

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
