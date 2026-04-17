package model

import "errors"

// ========== 更新线索状态 ==========
// https://developers.e.qq.com/v3.0/docs/api/leads_status/update

// 线索状态匹配类型枚举（更新线索状态使用，与新增线索的match_type不同）
const (
	LeadsStatusMatchTypeLeadsId      = "LEADSID"      // 线索id匹配
	LeadsStatusMatchTypeOuterLeadsId = "OUTERLEADSID" // 外部线索id匹配
	LeadsStatusMatchTypeContact      = "CONTACT"      // 联系方式匹配
	LeadsStatusMatchTypeClickId      = "CLICKID"      // 点击id匹配
)

// 线索转化状态枚举
const (
	LeadsConvertStatusDeprecated            = "LEADS_CONVERT_STATUS_DEPRECATED"              // 无效
	LeadsConvertStatusPotentialCustomer     = "LEADS_CONVERT_STATUS_POTENTIAL_CUSTOMER"      // 潜在客户
	LeadsConvertStatusHighIntentionCustomer = "LEADS_CONVERT_STATUS_HIGH_INTENTION_CUSTOMER" // 高意向客户
	LeadsConvertStatusTransCompleted        = "LEADS_CONVERT_STATUS_TRANS_COMPLETED"         // 已成单
)

// 线索无效原因枚举
const (
	LeadsIneffectReasonEmpty              = "LEADS_INEFFECT_REASON_EMPTY"               // 空号
	LeadsIneffectReasonIdentityMismatched = "LEADS_INEFFECT_REASON_IDENTITY_MISMATCHED" // 非本人
	LeadsIneffectReasonRegionMismatched   = "LEADS_INEFFECT_REASON_REGION_MISMATCHED"   // 定向外
	LeadsIneffectReasonDataDuplication    = "LEADS_INEFFECT_REASON_DATA_DUPLICATION"    // 重复
	LeadsIneffectReasonTelNotConnected    = "LEADS_INEFFECT_REASON_TEL_NOT_CONNECTED"   // 未接通
	LeadsIneffectReasonNoIntention        = "LEADS_INEFFECT_REASON_NO_INTENTION"        // 无意向
	LeadsIneffectReasonUnknown            = "LEADS_INEFFECT_REASON_UNKNOWN"             // 未知
)

// 回传线索列表长度限制
const (
	LeadsConversionStatusListMinLength = 1
	LeadsConversionStatusListMaxLength = 50
)

// LeadsStatusUpdateReq 更新线索状态请求
type LeadsStatusUpdateReq struct {
	GlobalReq
	AccountID                 int64                        `json:"account_id"`                   // 广告主账号id (必填)
	MatchType                 string                       `json:"match_type,omitempty"`         // 线索匹配类型，不填认为是OUTERLEADSID
	LeadsConversionStatusList []*LeadsConversionStatusItem `json:"leads_conversion_status_list"` // 回传线索信息的列表 (必填)，1-50条
}

// LeadsConversionStatusItem 回传线索信息
type LeadsConversionStatusItem struct {
	OuterLeadsId             string                `json:"outer_leads_id,omitempty"`              // 外部线索id，1-64字节
	LeadsId                  int64                 `json:"leads_id,omitempty"`                    // 线索id
	LeadsTel                 string                `json:"leads_tel,omitempty"`                   // 手机号，1-32字节
	LeadsQq                  int64                 `json:"leads_qq,omitempty"`                    // QQ号
	LeadsWechat              string                `json:"leads_wechat,omitempty"`                // 微信号，1-64字节
	ClickId                  string                `json:"click_id,omitempty"`                    // 点击id，1-64字节
	LeadsConvertType         string                `json:"leads_convert_type,omitempty"`          // 线索状态
	LeadsIneffectReason      string                `json:"leads_ineffect_reason,omitempty"`       // 无效原因
	OuterLeadsConvertType    string                `json:"outer_leads_convert_type,omitempty"`    // 外部线索状态，1-64字节
	OuterLeadsIneffectReason string                `json:"outer_leads_ineffect_reason,omitempty"` // 外部无效原因，1-32字节
	CustomizedTags           []*LeadsCustomizedTag `json:"customized_tags,omitempty"`             // 自定义标签集合，1-50条
}

func (p *LeadsStatusUpdateReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证更新线索状态请求参数
func (p *LeadsStatusUpdateReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	// 验证account_id
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	// 验证match_type
	if p.MatchType != "" && p.MatchType != LeadsStatusMatchTypeLeadsId &&
		p.MatchType != LeadsStatusMatchTypeOuterLeadsId && p.MatchType != LeadsStatusMatchTypeContact &&
		p.MatchType != LeadsStatusMatchTypeClickId {
		return errors.New("match_type值无效，可选值：LEADSID、OUTERLEADSID、CONTACT、CLICKID")
	}

	// 验证leads_conversion_status_list
	if len(p.LeadsConversionStatusList) == 0 {
		return errors.New("leads_conversion_status_list为必填")
	}
	if len(p.LeadsConversionStatusList) < LeadsConversionStatusListMinLength || len(p.LeadsConversionStatusList) > LeadsConversionStatusListMaxLength {
		return errors.New("leads_conversion_status_list数组长度必须在1-50之间")
	}

	// 验证customized_tags
	for _, item := range p.LeadsConversionStatusList {
		if len(item.CustomizedTags) > LeadsCustomizedTagsMaxLength {
			return errors.New("leads_conversion_status_list.customized_tags数组长度不能超过50")
		}
		for _, tag := range item.CustomizedTags {
			if len(tag.TagNameList) > LeadsTagNameListMaxLength {
				return errors.New("leads_conversion_status_list.customized_tags.tag_name_list数组长度不能超过100")
			}
		}
	}

	return nil
}

// LeadsStatusUpdateResp 更新线索状态响应
type LeadsStatusUpdateResp struct {
	FailLeadsList []*LeadsStatusFailItem `json:"fail_leads_list,omitempty"` // 返回失败的信息列表
}

// LeadsStatusFailItem 失败的线索信息
type LeadsStatusFailItem struct {
	Index int `json:"index,omitempty"` // 线索在请求参数中的索引
}
