package model

import "errors"

// ========== 获取蹊径落地页互动模板配置 ==========
// https://developers.e.qq.com/v3.0/docs/api/xijing_complex_template/get

// 字段长度常量
const (
	MinXijingComplexTemplateIDBytes = 1  // page_template_id 最小字节数
	MaxXijingComplexTemplateIDBytes = 32 // page_template_id 最大字节数
)

// XijingComplexTemplateGetReq 获取蹊径落地页互动模板配置请求（GET）
// https://developers.e.qq.com/v3.0/docs/api/xijing_complex_template/get
type XijingComplexTemplateGetReq struct {
	GlobalReq
	AccountID      int64  `json:"account_id"`       // 广告主帐号 id (必填)
	PageTemplateID string `json:"page_template_id"` // 蹊径落地页模板 id (必填)，1-32字节
}

func (r *XijingComplexTemplateGetReq) Format() {
	r.GlobalReq.Format()
}

// Validate 验证获取蹊径落地页互动模板配置请求参数
func (r *XijingComplexTemplateGetReq) Validate() error {
	if r.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if len(r.PageTemplateID) < MinXijingComplexTemplateIDBytes || len(r.PageTemplateID) > MaxXijingComplexTemplateIDBytes {
		return errors.New("page_template_id长度须在1-32字节之间")
	}
	return r.GlobalReq.Validate()
}

// XijingComplexTemplateConfigInfo 模板配置信息说明
type XijingComplexTemplateConfigInfo struct {
	Desc           string `json:"desc,omitempty"`             // 配置描述
	Width          int    `json:"width,omitempty"`            // 建议宽度
	Height         int    `json:"height,omitempty"`           // 建议高度
	MaxLength      int    `json:"max_length,omitempty"`       // 最大长度
	InfoNumberOnly bool   `json:"info_number_only,omitempty"` // 是否仅支持数字（文本类型时有效）
	Max            int    `json:"max,omitempty"`              // 最大数量
	Min            int    `json:"min,omitempty"`              // 最小数量
}

// XijingComplexTemplatePageConfig 模板配置项
type XijingComplexTemplatePageConfig struct {
	ID   string                           `json:"id"`             // 配置 id
	Type string                           `json:"type"`           // 配置数据类型
	Info *XijingComplexTemplateConfigInfo `json:"info,omitempty"` // 配置信息说明
}

// XijingComplexTemplateGetResp 获取蹊径落地页互动模板配置响应
// https://developers.e.qq.com/v3.0/docs/api/xijing_complex_template/get
type XijingComplexTemplateGetResp struct {
	PageTemplateID       string                             `json:"page_template_id"`                  // 蹊径落地页模板 id
	PageTemplateName     string                             `json:"page_template_name"`                // 模板名称
	PageTemplateCoverURL string                             `json:"page_template_cover_url,omitempty"` // 模板封面图 URL
	PageName             string                             `json:"page_name,omitempty"`               // 落地页名称（管理用）
	PageTitle            string                             `json:"page_title,omitempty"`              // 落地页标题（展示用）
	PageConfig           []*XijingComplexTemplatePageConfig `json:"page_config,omitempty"`             // 模板配置数组
}
