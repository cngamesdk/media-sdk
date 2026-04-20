package model

import "errors"

// ========== 更新线索归因信息 ==========
// https://developers.e.qq.com/v3.0/docs/api/leads_claim/update

// 归因线索列表长度限制
const (
	LeadsClaimListMinLength = 1
	LeadsClaimListMaxLength = 10
)

// LeadsClaimUpdateReq 更新线索归因信息请求
type LeadsClaimUpdateReq struct {
	GlobalReq
	AccountID      int64             `json:"account_id"`       // 广告主账号id (必填)
	LeadsClaimList []*LeadsClaimItem `json:"leads_claim_list"` // 回传线索归因信息的列表 (必填)，1-10条
}

// LeadsClaimItem 线索归因信息
type LeadsClaimItem struct {
	OuterLeadsId         string `json:"outer_leads_id"`                    // 外部线索id (必填)
	LeadsUserType        string `json:"leads_user_type,omitempty"`         // 线索用户类型
	LeadsUserWechatAppid string `json:"leads_user_wechat_appid,omitempty"` // 线索用户的微信AppId，1-64字节
	LeadsUserId          string `json:"leads_user_id,omitempty"`           // 线索用户id，1-64字节
	CampaignId           int64  `json:"campaign_id,omitempty"`             // 推广计划id
	AdgroupId            int64  `json:"adgroup_id,omitempty"`              // 广告id
	WechatAgencyId       string `json:"wechat_agency_id,omitempty"`        // 微信服务商id，1-32字节
}

func (p *LeadsClaimUpdateReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证更新线索归因信息请求参数
func (p *LeadsClaimUpdateReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	// 验证account_id
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	// 验证leads_claim_list
	if len(p.LeadsClaimList) == 0 {
		return errors.New("leads_claim_list为必填")
	}
	if len(p.LeadsClaimList) < LeadsClaimListMinLength || len(p.LeadsClaimList) > LeadsClaimListMaxLength {
		return errors.New("leads_claim_list数组长度必须在1-10之间")
	}

	for _, item := range p.LeadsClaimList {
		if item.OuterLeadsId == "" {
			return errors.New("leads_claim_list.outer_leads_id为必填")
		}
	}

	return nil
}

// LeadsClaimUpdateResp 更新线索归因信息响应
type LeadsClaimUpdateResp struct {
	FailLeadsList []*LeadsClaimFailItem `json:"fail_leads_list,omitempty"` // 返回失败的信息列表
}

// LeadsClaimFailItem 失败的线索信息
type LeadsClaimFailItem struct {
	OuterLeadsId string `json:"outer_leads_id,omitempty"` // 外部线索id
}
