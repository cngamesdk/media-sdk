package model

import "errors"

// ========== 更新线索基本信息 ==========
// https://developers.e.qq.com/v3.0/docs/api/leads/update

// 更新线索基本信息匹配类型枚举（不含CONTACT）
const (
	LeadsUpdateMatchTypeLeadsId      = "LEADSID"      // 线索id匹配
	LeadsUpdateMatchTypeOuterLeadsId = "OUTERLEADSID" // 外部线索id匹配
	LeadsUpdateMatchTypeClickId      = "CLICKID"      // 点击id匹配
)

// 回传线索列表长度限制
const (
	LeadsContactListMinLength = 1
	LeadsContactListMaxLength = 50
)

// LeadsUpdateReq 更新线索基本信息请求
type LeadsUpdateReq struct {
	GlobalReq
	AccountID        int64               `json:"account_id"`           // 广告主账号id (必填)
	MatchType        string              `json:"match_type,omitempty"` // 线索匹配类型，不填认为是OUTERLEADSID
	LeadsContactList []*LeadsContactItem `json:"leads_contact_list"`   // 回传线索信息的列表 (必填)，1-50条
}

// LeadsContactItem 线索联系信息
type LeadsContactItem struct {
	OuterLeadsId         string `json:"outer_leads_id,omitempty"`          // 外部线索id，1-64字节
	LeadsId              int64  `json:"leads_id,omitempty"`                // 线索id
	ClickId              string `json:"click_id,omitempty"`                // 点击id，1-64字节
	LeadsUserType        string `json:"leads_user_type,omitempty"`         // 线索用户类型
	LeadsUserWechatAppid string `json:"leads_user_wechat_appid,omitempty"` // 线索用户的微信AppId，1-64字节
	LeadsUserId          string `json:"leads_user_id,omitempty"`           // 线索用户id，1-64字节
	LeadsTel             string `json:"leads_tel,omitempty"`               // 手机号，1-32字节
	LeadsQq              int64  `json:"leads_qq,omitempty"`                // QQ号
	LeadsWechat          string `json:"leads_wechat,omitempty"`            // 微信号，1-64字节
	LeadsName            string `json:"leads_name,omitempty"`              // 姓名，0-128字节
	LeadsGender          string `json:"leads_gender,omitempty"`            // 性别
	LeadsEmail           string `json:"leads_email,omitempty"`             // 邮箱，0-64字节
	LeadsArea            string `json:"leads_area,omitempty"`              // 所在地，0-128字节
	Bundle               string `json:"bundle,omitempty"`                  // 其他线索信息，0-1024字节
	Memo                 string `json:"memo,omitempty"`                    // 备注，1-128字节
	ShopName             string `json:"shop_name,omitempty"`               // 门店名称，0-128字节
}

func (p *LeadsUpdateReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证更新线索基本信息请求参数
func (p *LeadsUpdateReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	// 验证account_id
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	// 验证match_type
	if p.MatchType != "" && p.MatchType != LeadsUpdateMatchTypeLeadsId &&
		p.MatchType != LeadsUpdateMatchTypeOuterLeadsId && p.MatchType != LeadsUpdateMatchTypeClickId {
		return errors.New("match_type值无效，可选值：LEADSID、OUTERLEADSID、CLICKID")
	}

	// 验证leads_contact_list
	if len(p.LeadsContactList) == 0 {
		return errors.New("leads_contact_list为必填")
	}
	if len(p.LeadsContactList) < LeadsContactListMinLength || len(p.LeadsContactList) > LeadsContactListMaxLength {
		return errors.New("leads_contact_list数组长度必须在1-50之间")
	}

	return nil
}

// LeadsUpdateResp 更新线索基本信息响应
type LeadsUpdateResp struct {
	FailLeadsList []*LeadsUpdateFailItem `json:"fail_leads_list,omitempty"` // 返回失败的信息列表
}

// LeadsUpdateFailItem 失败的线索信息
type LeadsUpdateFailItem struct {
	Index        int    `json:"index,omitempty"`          // 线索在请求参数中的索引
	OuterLeadsId string `json:"outer_leads_id,omitempty"` // 外部线索id
}
