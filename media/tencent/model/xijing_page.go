package model

import "errors"

// ========== 蹊径基于模板创建落地页 ==========
// https://developers.e.qq.com/v3.0/docs/api/xijing_page/add

// 蹊径创建落地页类型枚举（补充 xijing_template.go 中已有的 Android/iOS 类型）
const (
	XijingPageTypeXijingWebsite = "PAGE_TYPE_XIJING_WEBSITE" // 蹊径网站落地页
)

// 字段长度常量
const (
	MinXijingPageAddPageNameBytes    = 1   // page_name 最小字节数
	MaxXijingPageAddPageNameBytes    = 20  // page_name 最大字节数
	MinXijingPageAddPageTitleBytes   = 1   // page_title 最小字节数
	MaxXijingPageAddPageTitleBytes   = 20  // page_title 最大字节数
	MaxXijingPageAddClipboardBytes   = 300 // clipboard 最大字节数
	MinXijingPageAddDeepLinkBytes    = 1   // page_deeplink 最小字节数
	MaxXijingPageAddDeepLinkBytes    = 200 // page_deeplink 最大字节数
	MinXijingPageAddMobileAppIDBytes = 1   // mobile_app_id 最小字节数
	MaxXijingPageAddMobileAppIDBytes = 20  // mobile_app_id 最大字节数
	MinXijingPageAddFormIDBytes      = 1   // form_id 最小字节数
	MaxXijingPageAddFormIDBytes      = 20  // form_id 最大字节数
	MinXijingPageAddTemplateIDBytes  = 1   // page_template_id 最小字节数
	MaxXijingPageAddTemplateIDBytes  = 32  // page_template_id 最大字节数
	MaxXijingPageAddPagesCount       = 10  // pages 最大数量
	MaxXijingPageAddComponentsCount  = 10  // component_spec_list 最大数量
)

// XijingPageComponentSpec 创建落地页组件规格（复用 XijingTemplateComponentSpec 结构体）
// 组件类型枚举见 xijing_template.go 中的 XijingTemplateComponentType* 常量
type XijingPageComponentSpec = XijingTemplateComponentSpec

// XijingPageAddPage 单个落地页配置
type XijingPageAddPage struct {
	PageType          string                     `json:"page_type"`               // 落地页类型 (必填)
	PageName          string                     `json:"page_name"`               // 落地页名称（管理用）(必填)，1-20字节
	PageTitle         string                     `json:"page_title"`              // 落地页标题（展示用）(必填)，1-20字节
	Clipboard         string                     `json:"clipboard,omitempty"`     // 剪贴板内容，0-300字节
	PageDeeplink      string                     `json:"page_deeplink,omitempty"` // 页面级 deeplink，1-200字节
	MobileAppID       string                     `json:"mobile_app_id,omitempty"` // AppId，1-20字节（Android/iOS落地页按模板要求填写）
	FormID            string                     `json:"form_id,omitempty"`       // FormId，1-20字节（网站落地页按模板要求填写）
	PageTemplateID    string                     `json:"page_template_id"`        // 蹊径落地页模板 id (必填)，1-32字节
	ComponentSpecList []*XijingPageComponentSpec `json:"component_spec_list"`     // 组件列表 (必填)，最多10个
}

// XijingPageAddReq 蹊径基于模板创建落地页请求（POST JSON）
// https://developers.e.qq.com/v3.0/docs/api/xijing_page/add
type XijingPageAddReq struct {
	GlobalReq
	AccountID    int64                `json:"account_id"`     // 广告主帐号 id (必填)
	IsAutoSubmit int                  `json:"is_auto_submit"` // 是否自动提审 (必填)，0=否，1=是
	Pages        []*XijingPageAddPage `json:"pages"`          // 落地页配置列表 (必填)，最多10个
}

func (r *XijingPageAddReq) Format() {
	r.GlobalReq.Format()
}

// Validate 验证蹊径基于模板创建落地页请求参数
func (r *XijingPageAddReq) Validate() error {
	if r.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if r.IsAutoSubmit < 0 || r.IsAutoSubmit > 1 {
		return errors.New("is_auto_submit须为0或1")
	}
	if len(r.Pages) == 0 {
		return errors.New("pages为必填，至少包含1个落地页配置")
	}
	if len(r.Pages) > MaxXijingPageAddPagesCount {
		return errors.New("pages数组长度不能超过10")
	}
	for i, page := range r.Pages {
		if page == nil {
			return errors.New("pages[" + itoa(i) + "]不能为空")
		}
		if page.PageType == "" {
			return errors.New("pages[" + itoa(i) + "].page_type为必填")
		}
		if len(page.PageName) < MinXijingPageAddPageNameBytes || len(page.PageName) > MaxXijingPageAddPageNameBytes {
			return errors.New("pages[" + itoa(i) + "].page_name长度须在1-20字节之间")
		}
		if len(page.PageTitle) < MinXijingPageAddPageTitleBytes || len(page.PageTitle) > MaxXijingPageAddPageTitleBytes {
			return errors.New("pages[" + itoa(i) + "].page_title长度须在1-20字节之间")
		}
		if len(page.Clipboard) > MaxXijingPageAddClipboardBytes {
			return errors.New("pages[" + itoa(i) + "].clipboard长度不能超过300字节")
		}
		if page.PageDeeplink != "" && (len(page.PageDeeplink) < MinXijingPageAddDeepLinkBytes || len(page.PageDeeplink) > MaxXijingPageAddDeepLinkBytes) {
			return errors.New("pages[" + itoa(i) + "].page_deeplink长度须在1-200字节之间")
		}
		if page.MobileAppID != "" && (len(page.MobileAppID) < MinXijingPageAddMobileAppIDBytes || len(page.MobileAppID) > MaxXijingPageAddMobileAppIDBytes) {
			return errors.New("pages[" + itoa(i) + "].mobile_app_id长度须在1-20字节之间")
		}
		if page.FormID != "" && (len(page.FormID) < MinXijingPageAddFormIDBytes || len(page.FormID) > MaxXijingPageAddFormIDBytes) {
			return errors.New("pages[" + itoa(i) + "].form_id长度须在1-20字节之间")
		}
		if len(page.PageTemplateID) < MinXijingPageAddTemplateIDBytes || len(page.PageTemplateID) > MaxXijingPageAddTemplateIDBytes {
			return errors.New("pages[" + itoa(i) + "].page_template_id长度须在1-32字节之间")
		}
		if len(page.ComponentSpecList) > MaxXijingPageAddComponentsCount {
			return errors.New("pages[" + itoa(i) + "].component_spec_list数组长度不能超过10")
		}
		for j, comp := range page.ComponentSpecList {
			if comp == nil {
				return errors.New("pages[" + itoa(i) + "].component_spec_list[" + itoa(j) + "]不能为空")
			}
			if comp.Type == "" {
				return errors.New("pages[" + itoa(i) + "].component_spec_list[" + itoa(j) + "].type为必填")
			}
		}
	}
	return r.GlobalReq.Validate()
}

// XijingPageAddResultItem 单个落地页创建结果
type XijingPageAddResultItem struct {
	LandingPageID int64  `json:"landingPageId"` // 蹊径平台生成的落地页 id
	ID            string `json:"id"`            // 落地页服务 id
	Message       string `json:"message"`       // 操作结果信息
	Code          int    `json:"code"`          // 单条返回码
}

// XijingPageAddResp 蹊径基于模板创建落地页响应
// https://developers.e.qq.com/v3.0/docs/api/xijing_page/add
type XijingPageAddResp struct {
	List []*XijingPageAddResultItem `json:"list"` // 创建结果列表
}
