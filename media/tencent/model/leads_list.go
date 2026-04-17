package model

import "errors"

// ========== 获取线索列表 ==========
// https://developers.e.qq.com/v3.0/docs/api/leads_list/get

// 线索列表分页默认值
const (
	LeadsListDefaultPage     = 1
	LeadsListDefaultPageSize = 10
	LeadsListMaxPage         = 1000
	LeadsListMaxPageSize     = 200
)

// 时间类型枚举
const (
	TimeTypeCreatedTime = "TIME_TYPE_CREATED_TIME" // 线索入库时间
	TimeTypeActionTime  = "TIME_TYPE_ACTION_TIME"  // 线索提交时间
)

// 时间戳范围
const (
	LeadsListMinTimestamp int64 = 0
	LeadsListMaxTimestamp int64 = 9999999999
)

// 深度翻页参数长度
const (
	LeadsListSearchAfterValuesLength = 2
)

// LeadsListGetReq 获取线索列表请求
type LeadsListGetReq struct {
	GlobalReq
	AccountID             int64           `json:"account_id"`                         // 广告主账号id (必填)
	TimeRange             *LeadsTimeRange `json:"time_range"`                         // 时间范围 (必填)
	Page                  int             `json:"page,omitempty"`                     // 搜索页码，默认值：1，最大值1000
	PageSize              int             `json:"page_size,omitempty"`                // 一页显示的数据条数，默认值：10，最大值200
	LastSearchAfterValues []string        `json:"last_search_after_values,omitempty"` // 线索深度翻页参数
}

// LeadsTimeRange 时间范围
type LeadsTimeRange struct {
	StartTime int64  `json:"start_time"` // 开始时间戳，精确到秒 (必填)
	EndTime   int64  `json:"end_time"`   // 结束时间戳，精确到秒 (必填)
	TimeType  string `json:"time_type"`  // 时间类型 (必填)
}

func (p *LeadsListGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page <= 0 {
		p.Page = LeadsListDefaultPage
	}
	if p.PageSize <= 0 {
		p.PageSize = LeadsListDefaultPageSize
	}
}

// Validate 验证获取线索列表请求参数
func (p *LeadsListGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	// 验证account_id
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	// 验证time_range
	if p.TimeRange == nil {
		return errors.New("time_range为必填")
	}
	if p.TimeRange.StartTime < LeadsListMinTimestamp || p.TimeRange.StartTime > LeadsListMaxTimestamp {
		return errors.New("time_range.start_time必须在0-9999999999之间")
	}
	if p.TimeRange.EndTime < LeadsListMinTimestamp || p.TimeRange.EndTime > LeadsListMaxTimestamp {
		return errors.New("time_range.end_time必须在0-9999999999之间")
	}
	if p.TimeRange.StartTime >= p.TimeRange.EndTime {
		return errors.New("time_range.start_time必须小于end_time")
	}
	if p.TimeRange.TimeType == "" {
		return errors.New("time_range.time_type为必填")
	}
	if p.TimeRange.TimeType != TimeTypeCreatedTime && p.TimeRange.TimeType != TimeTypeActionTime {
		return errors.New("time_range.time_type值无效，可选值：TIME_TYPE_CREATED_TIME、TIME_TYPE_ACTION_TIME")
	}

	// 验证page
	if p.Page < 1 || p.Page > LeadsListMaxPage {
		return errors.New("page必须在1-1000之间")
	}

	// 验证page_size
	if p.PageSize < 1 || p.PageSize > LeadsListMaxPageSize {
		return errors.New("page_size必须在1-200之间")
	}

	// 验证last_search_after_values
	if len(p.LastSearchAfterValues) > 0 && len(p.LastSearchAfterValues) != LeadsListSearchAfterValuesLength {
		return errors.New("last_search_after_values长度必须为2")
	}

	return nil
}

// LeadsListGetResp 获取线索列表响应
type LeadsListGetResp struct {
	LeadsInfo *LeadsInfo `json:"leads_info,omitempty"` // 返回结构
	PageInfo  *PageInfo  `json:"page_info,omitempty"`  // 分页配置信息
}

// LeadsInfo 线索信息
type LeadsInfo struct {
	AccountID                int64  `json:"account_id,omitempty"`                  // 广告主账号id
	LeadsId                  int64  `json:"leads_id,omitempty"`                    // 线索id
	OuterLeadsId             string `json:"outer_leads_id,omitempty"`              // 外部线索id
	ClickId                  string `json:"click_id,omitempty"`                    // 点击id
	AdgroupId                int64  `json:"adgroup_id,omitempty"`                  // 广告id
	AdgroupName              string `json:"adgroup_name,omitempty"`                // 广告名称
	DynamicCreativeId        int64  `json:"dynamic_creative_id,omitempty"`         // 动态创意id
	DynamicCreativeName      string `json:"dynamic_creative_name,omitempty"`       // 动态创意名称
	ComponentId              string `json:"component_id,omitempty"`                // 组件id
	ComponentName            string `json:"component_name,omitempty"`              // 组件名称
	PageId                   int64  `json:"page_id,omitempty"`                     // 落地页id
	PageName                 string `json:"page_name,omitempty"`                   // 落地页名称
	PageUrl                  string `json:"page_url,omitempty"`                    // 落地页地址
	LeadsType                string `json:"leads_type,omitempty"`                  // 线索类型
	LeadsSubType             string `json:"leads_sub_type,omitempty"`              // 二级线索类型
	ChatId                   string `json:"chat_id,omitempty"`                     // 会话id
	LeadsSource              string `json:"leads_source,omitempty"`                // 线索来源
	LeadsFollowTag           string `json:"leads_follow_tag,omitempty"`            // 线索状态
	OuterLeadsConvertType    string `json:"outer_leads_convert_type,omitempty"`    // 外部线索状态
	OuterLeadsIneffectReason string `json:"outer_leads_ineffect_reason,omitempty"` // 外部无效原因
	LeadsName                string `json:"leads_name,omitempty"`                  // 姓名
	LeadsTelephone           string `json:"leads_telephone,omitempty"`             // 电话
	TelephoneLocation        string `json:"telephone_location,omitempty"`          // 号码归属地
	LeadsArea                string `json:"leads_area,omitempty"`                  // 所在地
	LeadsEmail               string `json:"leads_email,omitempty"`                 // 邮箱
	LeadsQq                  string `json:"leads_qq,omitempty"`                    // qq号
	LeadsWechat              string `json:"leads_wechat,omitempty"`                // 微信号
	LeadsGender              string `json:"leads_gender,omitempty"`                // 性别
	Nationality              string `json:"nationality,omitempty"`                 // 国籍
	WorkingYears             string `json:"working_years,omitempty"`               // 工作年限
	Age                      string `json:"age,omitempty"`                         // 年龄
	Profession               string `json:"profession,omitempty"`                  // 职业
	IdNumber                 string `json:"id_number,omitempty"`                   // 身份证号
	Address                  string `json:"address,omitempty"`                     // 详细地址
	Bundle                   string `json:"bundle,omitempty"`                      // 其他线索信息
	CustomQa                 string `json:"custom_qa,omitempty"`                   // 互动问答
	LeadsCreateTime          string `json:"leads_create_time,omitempty"`           // 线索入库时间
	LeadsActionTime          string `json:"leads_action_time,omitempty"`           // 线索提交时间
	LeadsTags                string `json:"leads_tags,omitempty"`                  // 线索标签
	ShopName                 string `json:"shop_name,omitempty"`                   // 门店名称
	ShopAddress              string `json:"shop_address,omitempty"`                // 门店地址
	CallMiddleNum            string `json:"call_middle_num,omitempty"`             // 智能电话中间号
	CallConsumerHotline      string `json:"call_consumer_hotline,omitempty"`       // 智能电话-客服电话
	CallTouchTag             string `json:"call_touch_tag,omitempty"`              // 智能电话-接通状态
	CallDuration             string `json:"call_duration,omitempty"`               // 智能电话-通话时长
	CallRecordUrl            string `json:"call_record_url,omitempty"`             // 智能电话-通话录音链接
	LayerFormContent         string `json:"layer_form_content,omitempty"`          // 意向表单内容
	IsBroadCastLeads         string `json:"is_broad_cast_leads,omitempty"`         // 是否为多发表单
	NickName                 string `json:"nick_name,omitempty"`                   // 微信昵称
	OwnerName                string `json:"owner_name,omitempty"`                  // 当前归属人
	OwnerId                  int64  `json:"owner_id,omitempty"`                    // 当前归属人的userId
	AllFollowRecords         string `json:"all_follow_records,omitempty"`          // 全部跟进记录
	CorrelationFactorRank    string `json:"correlation_factor_rank,omitempty"`     // 决策动因相关度
	ClaimAccountId           int64  `json:"claim_account_id,omitempty"`            // 广告归因的广告主id
	LeadsResponseDuration    int64  `json:"leads_response_duration,omitempty"`     // 线索响应时长
}
