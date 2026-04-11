package model

import "errors"

// ========== 获取蹊径落地页模板列表 ==========
// https://developers.e.qq.com/v3.0/docs/api/xijing_template_list/get

// 模板查询类型枚举
const (
	XijingTemplateSourceGrant = "GRANT" // 授权模板
	XijingTemplateSourceOwner = "OWNER" // 自有模板（默认）
)

// 分页常量
const (
	MinXijingTemplateListPage         = 1     // page 最小值
	MaxXijingTemplateListPage         = 99999 // page 最大值
	MinXijingTemplateListPageSize     = 1     // page_size 最小值
	MaxXijingTemplateListPageSize     = 100   // page_size 最大值
	DefaultXijingTemplateListPage     = 1     // page 默认值
	DefaultXijingTemplateListPageSize = 10    // page_size 默认值
)

// 字段长度常量
const (
	MinXijingTemplateListTemplateIDBytes = 1  // page_template_id 最小字节数
	MaxXijingTemplateListTemplateIDBytes = 32 // page_template_id 最大字节数
)

// XijingTemplateListGetReq 获取蹊径落地页模板列表请求（GET）
// https://developers.e.qq.com/v3.0/docs/api/xijing_template_list/get
type XijingTemplateListGetReq struct {
	GlobalReq
	AccountID      int64  `json:"account_id"`                 // 广告主帐号 id (必填)
	PageTemplateID string `json:"page_template_id,omitempty"` // 蹊径落地页模板 id，1-32字节
	IsInteraction  *bool  `json:"is_interaction,omitempty"`   // 是否为互动模板，默认false
	IsPublic       *bool  `json:"is_public,omitempty"`        // 是否为公共模板，默认true
	TemplateSource string `json:"template_source,omitempty"`  // 查询类型，GRANT/OWNER，默认OWNER
	PageSize       int    `json:"page_size,omitempty"`        // 每页条数，1-100，默认10
	Page           int    `json:"page,omitempty"`             // 搜索页码，1-99999，默认1
}

func (r *XijingTemplateListGetReq) Format() {
	r.GlobalReq.Format()
	if r.Page == 0 {
		r.Page = DefaultXijingTemplateListPage
	}
	if r.PageSize == 0 {
		r.PageSize = DefaultXijingTemplateListPageSize
	}
}

// Validate 验证获取蹊径落地页模板列表请求参数
func (r *XijingTemplateListGetReq) Validate() error {
	if r.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if r.PageTemplateID != "" {
		if len(r.PageTemplateID) < MinXijingTemplateListTemplateIDBytes || len(r.PageTemplateID) > MaxXijingTemplateListTemplateIDBytes {
			return errors.New("page_template_id长度须在1-32字节之间")
		}
	}
	if r.Page < MinXijingTemplateListPage || r.Page > MaxXijingTemplateListPage {
		return errors.New("page须在1-99999之间")
	}
	if r.PageSize < MinXijingTemplateListPageSize || r.PageSize > MaxXijingTemplateListPageSize {
		return errors.New("page_size须在1-100之间")
	}
	return r.GlobalReq.Validate()
}

// XijingTemplateListItem 蹊径落地页模板列表项
type XijingTemplateListItem struct {
	PageTemplateID       string   `json:"page_template_id"`        // 蹊径落地页模板 id
	TemplateOwnerID      int64    `json:"template_owner_id"`       // 授权模板所属账户 id
	PageTemplateName     string   `json:"page_template_name"`      // 模板名称
	PageTemplateCoverURL string   `json:"page_template_cover_url"` // 模板封面图 URL
	PlayableType         string   `json:"playable_type"`           // 互动落地页类型
	Labels               []string `json:"labels"`                  // 模板支持的落地页类型标签
	IsComplex            bool     `json:"is_complex"`              // 是否为复杂模板
}

// XijingTemplateListGetResp 获取蹊径落地页模板列表响应
// https://developers.e.qq.com/v3.0/docs/api/xijing_template_list/get
type XijingTemplateListGetResp struct {
	List     []*XijingTemplateListItem `json:"list"`      // 返回信息列表
	PageInfo *PageInfo                 `json:"page_info"` // 分页配置信息
}
