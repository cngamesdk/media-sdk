package model

import "errors"

// ========== 获取客户人群 ==========
// https://developers.e.qq.com/v3.0/docs/api/custom_audiences/get

// CustomAudiencesGetReq 获取客户人群请求
type CustomAudiencesGetReq struct {
	GlobalReq
	AccountID  int64 `json:"account_id"`            // 推广帐号 id (必填)
	AudienceID int64 `json:"audience_id,omitempty"` // 人群 id，不传则返回全部
	Page       int   `json:"page,omitempty"`        // 当前页码，最小值 1，默认值 1
	PageSize   int   `json:"page_size,omitempty"`   // 分页大小，最小值 1，最大值 100，默认值 10
}

func (p *CustomAudiencesGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
}

// Validate 验证获取客户人群请求参数
func (p *CustomAudiencesGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.Page < 1 {
		return errors.New("page最小值为1")
	}
	if p.PageSize < 1 || p.PageSize > 100 {
		return errors.New("page_size必须在1-100之间")
	}
	return nil
}

// CustomAudiencesGetResp 获取客户人群响应
type CustomAudiencesGetResp struct {
	List     []*CustomAudienceItem `json:"list,omitempty"`      // 返回人群列表
	PageInfo *PageInfo             `json:"page_info,omitempty"` // 分页信息
}

// CustomAudienceItem 人群信息
type CustomAudienceItem struct {
	AudienceID       int64                   `json:"audience_id,omitempty"`        // 人群 id
	AccountID        int64                   `json:"account_id,omitempty"`         // 人群归属的推广帐号 id
	Name             string                  `json:"name,omitempty"`               // 人群名称
	OuterAudienceID  string                  `json:"outer_audience_id,omitempty"`  // 广告主对人群在自己系统里的编码
	Description      string                  `json:"description,omitempty"`        // 人群描述
	Cooperated       bool                    `json:"cooperated,omitempty"`         // 深度数据合作
	Type             string                  `json:"type,omitempty"`               // 人群类型
	Source           string                  `json:"source,omitempty"`             // 人群来源
	Status           string                  `json:"status,omitempty"`             // 处理状态
	OnlineStatus     string                  `json:"online_status,omitempty"`      // 人群包在线状态，仅处理状态为成功时返回
	IsOwn            bool                    `json:"is_own,omitempty"`             // 是否是人群包 owner
	ErrorCode        int                     `json:"error_code,omitempty"`         // 人群错误码
	UserCount        int64                   `json:"user_count,omitempty"`         // 用户覆盖数
	CreatedTime      string                  `json:"created_time,omitempty"`       // 创建时间，格式 yyyy-MM-dd HH:mm:ss
	LastModifiedTime string                  `json:"last_modified_time,omitempty"` // 最后更新时间，格式 yyyy-MM-dd HH:mm:ss
	AudienceSpec     *CustomAudienceItemSpec `json:"audience_spec,omitempty"`      // 人群信息，和 type 相关
}

// CustomAudienceItemSpec 人群详细信息（响应结构，复用 add 中定义的子结构）
type CustomAudienceItemSpec struct {
	LookalikeSpec  *LookalikeSpec  `json:"lookalike_spec,omitempty"`   // Lookalike 人群信息
	UserActionSpec *UserActionSpec `json:"user_action_spec,omitempty"` // UserAction 人群信息
	KeywordSpec    *KeywordSpec    `json:"keyword_spec,omitempty"`     // Keyword 人群信息
	AdRuleSpec     *AdRuleSpec     `json:"ad_rule_spec,omitempty"`     // 广告人群信息
	CombineSpec    *CombineSpec    `json:"combine_spec,omitempty"`     // 组合人群信息
}
