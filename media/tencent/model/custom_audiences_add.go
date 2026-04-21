package model

import "errors"

// ========== 创建客户人群 ==========
// https://developers.e.qq.com/v3.0/docs/api/custom_audiences/add

// CustomAudiencesAddReq 创建客户人群请求
type CustomAudiencesAddReq struct {
	GlobalReq
	AccountID       int64         `json:"account_id"`                  // 推广帐号 id (必填)
	Name            string        `json:"name"`                        // 人群名称，同一帐号下不许重名，1-32字节 (必填)
	Type            string        `json:"type"`                        // 人群类型 (必填)，可选值：CUSTOMER_FILE, LOOKALIKE, USER_ACTION, KEYWORD, AD, COMBINE
	OuterAudienceID string        `json:"outer_audience_id,omitempty"` // 广告主对人群在自己系统里的编码，1-128字节
	Description     string        `json:"description,omitempty"`       // 定向描述，0-250字节
	AudienceSpec    *AudienceSpec `json:"audience_spec,omitempty"`     // 人群信息，和 type 相关
}

func (p *CustomAudiencesAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证创建客户人群请求参数
func (p *CustomAudiencesAddReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.Name == "" {
		return errors.New("name为必填")
	}
	if len(p.Name) > 32 {
		return errors.New("name长度不能超过32字节")
	}
	if p.Type == "" {
		return errors.New("type为必填")
	}
	validTypes := map[string]bool{
		"CUSTOMER_FILE": true,
		"LOOKALIKE":     true,
		"USER_ACTION":   true,
		"KEYWORD":       true,
		"AD":            true,
		"COMBINE":       true,
	}
	if !validTypes[p.Type] {
		return errors.New("type值无效，可选值：CUSTOMER_FILE, LOOKALIKE, USER_ACTION, KEYWORD, AD, COMBINE")
	}
	return nil
}

// AudienceSpec 人群信息
type AudienceSpec struct {
	LookalikeSpec  *LookalikeSpec  `json:"lookalike_spec,omitempty"`   // Lookalike 人群信息，type=LOOKALIKE 时必填
	UserActionSpec *UserActionSpec `json:"user_action_spec,omitempty"` // UserAction 人群信息，type=USER_ACTION 时必填
	KeywordSpec    *KeywordSpec    `json:"keyword_spec,omitempty"`     // Keyword 人群信息，type=KEYWORD 时必填
	AdRuleSpec     *AdRuleSpec     `json:"ad_rule_spec,omitempty"`     // 广告人群信息，type=AD 时必填
	CombineSpec    *CombineSpec    `json:"combine_spec,omitempty"`     // 组合人群信息，type=COMBINE 时必填
}

// LookalikeSpec Lookalike 人群信息
type LookalikeSpec struct {
	SeedAudienceID  int64 `json:"seed_audience_id"`  // 种子人群 id (必填)
	ExpandUserCount int64 `json:"expand_user_count"` // lookalike 目标人数，500000的整数倍，500000-100000000 (必填)
}

// UserActionSpec UserAction 人群信息
type UserActionSpec struct {
	UserActionSetID       int64                  `json:"user_action_set_id"`                // 用户行为源 id (必填)
	MatchRuleType         string                 `json:"match_rule_type"`                   // 匹配规则类型，可选值：URL, ACTION (必填)
	ExtractType           string                 `json:"extract_type,omitempty"`            // 行为人群提取类型，可选值：FILTER, AGGREGATION
	TimeWindow            int                    `json:"time_window"`                       // 时间窗，0-180 (必填)
	URLMatchRule          *URLMatchRule          `json:"url_match_rule,omitempty"`          // url 匹配规则，match_rule_type=URL 时必填
	ActionMatchRule       *ActionMatchRule       `json:"action_match_rule,omitempty"`       // 行为匹配规则，match_rule_type=ACTION 且 extract_type 为空或 FILTER 时必填
	ActionAggregationRule *ActionAggregationRule `json:"action_aggregation_rule,omitempty"` // 行为聚合规则，match_rule_type=ACTION 且 extract_type=AGGREGATION 时必填
}

// URLMatchRule url 匹配规则
type URLMatchRule struct {
	URLMatcherGroup []*URLMatcherGroup `json:"url_matcher_group"` // 匹配规则组，AND 关系，数组最大长度 16 (必填)
}

// URLMatcherGroup 匹配规则组
type URLMatcherGroup struct {
	URLMatcher []*URLMatcher `json:"url_matcher"` // 匹配规则，OR 关系，数组最大长度 16 (必填)
}

// URLMatcher 匹配规则
type URLMatcher struct {
	ParamValue string `json:"param_value"` // 参数值，1-128字节 (必填)
	Operator   string `json:"operator"`    // 运算符，可选值：LT, GT, EQ, NE, CONTAIN, NOT_CONTAIN (必填)
}

// ActionMatchRule 行为匹配规则
type ActionMatchRule struct {
	ActionType        string               `json:"action_type"`                   // 标准行为类型 (必填)
	CustomAction      string               `json:"custom_action,omitempty"`       // 自定义行为类型，action_type=CUSTOM 时必填
	ParamMatcherGroup []*ParamMatcherGroup `json:"param_matcher_group,omitempty"` // 匹配规则组，AND 关系，数组最大长度 16
}

// ParamMatcherGroup 参数匹配规则组
type ParamMatcherGroup struct {
	ParamMatcher []*ParamMatcher `json:"param_matcher"` // 匹配规则，OR 关系，数组最大长度 16 (必填)
}

// ParamMatcher 参数匹配规则
type ParamMatcher struct {
	ParamName  string `json:"param_name"`  // 参数名称，1-128字节 (必填)
	ParamValue string `json:"param_value"` // 参数值，1-128字节 (必填)
	Operator   string `json:"operator"`    // 运算符，可选值：LT, GT, EQ, NE, CONTAIN, NOT_CONTAIN (必填)
}

// ActionAggregationRule 行为聚合规则
type ActionAggregationRule struct {
	ActionType       string              `json:"action_type"`             // 标准行为类型 (必填)
	CustomAction     string              `json:"custom_action,omitempty"` // 自定义行为类型，action_type=CUSTOM 时必填
	AggregationGroup []*AggregationGroup `json:"aggregation_group"`       // 聚合规则数组，AND 关系，1-4个 (必填)
	FilterGroup      []*FilterGroup      `json:"filter_group,omitempty"`  // 匹配规则组，AND 关系，1-16个
}

// AggregationGroup 聚合规则组
type AggregationGroup struct {
	AggregationMatcher []*AggregationMatcher `json:"aggregation_matcher"` // 匹配规则组，OR 关系，长度为1 (必填)
}

// AggregationMatcher 聚合匹配规则
type AggregationMatcher struct {
	AggregationType    string `json:"aggregation_type"`               // 聚合类型，可选值：SUM, MAX, MIN, COUNT (必填)
	CountType          string `json:"count_type,omitempty"`           // 频次类型，aggregation_type=COUNT 时必填，可选值：BY_TIMES, BY_DAY
	ParamName          string `json:"param_name,omitempty"`           // 参数名称，aggregation_type!=COUNT 时必填，1-128字节
	Comparator         string `json:"comparator"`                     // 比较符，可选值：COMPARATOR_GE, COMPARATOR_LE, COMPARATOR_BETWEEN, COMPARATOR_EQ (必填)
	ComparisonValue    int64  `json:"comparison_value,omitempty"`     // 参数值，comparator!=COMPARATOR_BETWEEN 时必填
	ComparisonMinValue int64  `json:"comparison_min_value,omitempty"` // 参数值，comparator=COMPARATOR_BETWEEN 时必填
	ComparisonMaxValue int64  `json:"comparison_max_value,omitempty"` // 参数值，comparator=COMPARATOR_BETWEEN 时必填
}

// FilterGroup 过滤规则组
type FilterGroup struct {
	ParamMatcher []*ParamMatcher `json:"param_matcher"` // 匹配规则，OR 关系，长度为1 (必填)
}

// KeywordSpec Keyword 人群信息
type KeywordSpec struct {
	IncludeKeyword []string `json:"include_keyword"`           // 包含关键词，最多 100 个，单个不超过 10 字节 (必填)
	ExcludeKeyword []string `json:"exclude_keyword,omitempty"` // 排除关键词，最多 20 个，单个不超过 10 字节
}

// AdRuleSpec 广告人群信息
type AdRuleSpec struct {
	RuleType       string   `json:"rule_type"`                 // 广告行为类型，可选值：EXPOSURE, CLICK, CONVERSION (必填)
	ConversionType []string `json:"conversion_type,omitempty"` // 广告转化类型，rule_type=CONVERSION 时必填
	StartDate      string   `json:"start_date"`                // 数据起始日期，格式 yyyy-MM-dd (必填)
	EndDate        string   `json:"end_date,omitempty"`        // 数据结束日期，格式 yyyy-MM-dd
	AdgroupIDList  []int64  `json:"adgroup_id_list,omitempty"` // 需要提取人群的 adgroup id 列表，最多 75 个
}

// CombineSpec 组合人群信息
type CombineSpec struct {
	Include [][]CombineAudienceItem `json:"include"`           // 包含的人群，二维数组，第一层 AND，第二层 OR (必填)
	Exclude [][]CombineAudienceItem `json:"exclude,omitempty"` // 排除的人群，二维数组
}

// CombineAudienceItem 组合人群元素
type CombineAudienceItem struct {
	AudienceID int64 `json:"audience_id"`           // 人群 id 或标签 id (必填)
	TimeWindow int   `json:"time_window,omitempty"` // 时间窗，仅对客户标签有效，1-365
}

// CustomAudiencesAddResp 创建客户人群响应
type CustomAudiencesAddResp struct {
	AudienceID int64 `json:"audience_id"` // 人群 id
}
