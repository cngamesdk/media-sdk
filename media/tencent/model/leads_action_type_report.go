package model

import "errors"

// ========== 线索上报 DMP 平台 ==========
// https://developers.e.qq.com/v3.0/docs/api/leads_action_type_report/add

// 行为上报匹配类型枚举
const (
	LeadsActionReportMatchTypeLeadsId      = "LEADSID"      // 线索id匹配
	LeadsActionReportMatchTypeOuterLeadsId = "OUTERLEADSID" // 外部线索id匹配
	LeadsActionReportMatchTypeContact      = "CONTACT"      // 联系方式匹配
	LeadsActionReportMatchTypeClickId      = "CLICKID"      // 点击id匹配
)

// 用户行为类型常量
const (
	ActionTypePhoneConnected = "PHONE_CONNECTED" // 电话接通，需填写call_duration
)

// 行为上报列表长度限制
const (
	LeadsActionTypeReportListMinLength = 1
	LeadsActionTypeReportListMaxLength = 50
)

// LeadsActionTypeReportAddReq 线索上报DMP平台请求
type LeadsActionTypeReportAddReq struct {
	GlobalReq
	AccountID                 int64                        `json:"account_id"`                    // 广告主账号id (必填)
	MatchType                 string                       `json:"match_type,omitempty"`          // 线索匹配类型，不填认为是OUTERLEADSID
	LeadsActionTypeReportList []*LeadsActionTypeReportItem `json:"leads_action_type_report_list"` // 回传线索信息的列表 (必填)，1-50条
}

// LeadsActionTypeReportItem 行为上报线索信息
type LeadsActionTypeReportItem struct {
	OuterLeadsId string `json:"outer_leads_id,omitempty"` // 外部线索id，1-64字节
	LeadsId      int64  `json:"leads_id,omitempty"`       // 线索id
	LeadsTel     string `json:"leads_tel,omitempty"`      // 手机号，1-32字节
	LeadsQq      int64  `json:"leads_qq,omitempty"`       // QQ号
	LeadsWechat  string `json:"leads_wechat,omitempty"`   // 微信号，1-64字节
	ClickId      string `json:"click_id,omitempty"`       // 点击id，1-64字节
	ActionType   string `json:"action_type"`              // 用户行为类型 (必填)，1-64字节
	CallDuration int64  `json:"call_duration,omitempty"`  // 通话时长，单位(s)
}

func (p *LeadsActionTypeReportAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证线索上报DMP平台请求参数
func (p *LeadsActionTypeReportAddReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	// 验证account_id
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	// 验证match_type
	if p.MatchType != "" && p.MatchType != LeadsActionReportMatchTypeLeadsId &&
		p.MatchType != LeadsActionReportMatchTypeOuterLeadsId && p.MatchType != LeadsActionReportMatchTypeContact &&
		p.MatchType != LeadsActionReportMatchTypeClickId {
		return errors.New("match_type值无效，可选值：LEADSID、OUTERLEADSID、CONTACT、CLICKID")
	}

	// 验证leads_action_type_report_list
	if len(p.LeadsActionTypeReportList) == 0 {
		return errors.New("leads_action_type_report_list为必填")
	}
	if len(p.LeadsActionTypeReportList) < LeadsActionTypeReportListMinLength || len(p.LeadsActionTypeReportList) > LeadsActionTypeReportListMaxLength {
		return errors.New("leads_action_type_report_list数组长度必须在1-50之间")
	}

	for _, item := range p.LeadsActionTypeReportList {
		if item.ActionType == "" {
			return errors.New("leads_action_type_report_list.action_type为必填")
		}
		// action_type为PHONE_CONNECTED时call_duration必填
		if item.ActionType == ActionTypePhoneConnected && item.CallDuration <= 0 {
			return errors.New("action_type为PHONE_CONNECTED时call_duration必填且大于0")
		}
	}

	return nil
}

// LeadsActionTypeReportAddResp 线索上报DMP平台响应
type LeadsActionTypeReportAddResp struct {
	FailLeadsList []*LeadsActionTypeReportFailItem `json:"fail_leads_list,omitempty"` // 返回失败的信息列表
}

// LeadsActionTypeReportFailItem 失败的线索信息
type LeadsActionTypeReportFailItem struct {
	Index int `json:"index,omitempty"` // 线索在请求参数中的索引
}
