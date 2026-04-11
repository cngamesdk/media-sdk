package model

import "errors"

// ========== 获取蹊径落地页模板 ==========
// https://developers.e.qq.com/v3.0/docs/api/xijing_template/get

// 蹊径落地页模板类型枚举
const (
	XijingTemplatePageTypeAndroid       = "TEMPLATE_TYPE_ANDROID"    // Android 模板
	XijingTemplatePageTypeXijingAndroid = "PAGE_TYPE_XIJING_ANDROID" // 蹊径 Android 落地页
	XijingTemplatePageTypeXijingIos     = "PAGE_TYPE_XIJING_IOS"     // 蹊径 iOS 落地页
)

// 组件类型枚举
const (
	XijingTemplateComponentTypeVideo         = "COMPONENT_TYPE_VIDEO"           // 视频组件
	XijingTemplateComponentTypeText          = "COMPONENT_TYPE_TEXT"            // 文本组件
	XijingTemplateComponentTypeImages        = "COMPONENT_TYPE_IMAGES"          // 图片列表组件
	XijingTemplateComponentTypeButton        = "COMPONENT_TYPE_BUTTON"          // 按钮组件
	XijingTemplateComponentTypeAppInfoButton = "COMPONENT_TYPE_APP_INFO_BUTTON" // App 信息按钮组件
	XijingTemplateComponentTypeFixedButton   = "COMPONENT_TYPE_FIXED_BUTTON"    // 固定按钮组件
)

// 字段长度常量
const (
	MinXijingTemplateGetTemplateIDBytes = 1  // template_id 最小字节数
	MaxXijingTemplateGetTemplateIDBytes = 32 // template_id 最大字节数
)

// XijingTemplateGetReq 获取蹊径落地页模板请求（GET）
// https://developers.e.qq.com/v3.0/docs/api/xijing_template/get
type XijingTemplateGetReq struct {
	GlobalReq
	AccountID  int64  `json:"account_id"`  // 广告主帐号 id (必填)
	TemplateID string `json:"template_id"` // 蹊径落地页模板 id (必填)，1-32字节
}

func (r *XijingTemplateGetReq) Format() {
	r.GlobalReq.Format()
}

// Validate 验证获取蹊径落地页模板请求参数
func (r *XijingTemplateGetReq) Validate() error {
	if r.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if len(r.TemplateID) < MinXijingTemplateGetTemplateIDBytes || len(r.TemplateID) > MaxXijingTemplateGetTemplateIDBytes {
		return errors.New("template_id长度须在1-32字节之间")
	}
	return r.GlobalReq.Validate()
}

// XijingTemplateVideoSpec 视频组件规格
type XijingTemplateVideoSpec struct {
	VideoID string `json:"video_id"` // 视频 id
}

// XijingTemplateTextSpec 文本组件规格
type XijingTemplateTextSpec struct {
	Text string `json:"text"` // 文本内容
}

// XijingTemplateImageItem 图片列表中的单张图片
type XijingTemplateImageItem struct {
	ImageID string `json:"image_id"` // 图片 id
}

// XijingTemplateImageListSpec 图片列表组件规格
type XijingTemplateImageListSpec struct {
	ImageList []*XijingTemplateImageItem `json:"image_list"` // 图片列表
}

// XijingTemplateButtonSpec 按钮组件规格
type XijingTemplateButtonSpec struct {
	Text string `json:"text"` // 按钮文案
}

// XijingTemplateAppInfoButtonSpec App 信息按钮组件规格
type XijingTemplateAppInfoButtonSpec struct {
	Text string `json:"text"` // 按钮文案
}

// XijingTemplateFixedButtonSpec 固定按钮组件规格
type XijingTemplateFixedButtonSpec struct {
	Desc string `json:"desc"` // 按钮描述
}

// XijingTemplateComponentSpec 组件规格
type XijingTemplateComponentSpec struct {
	Type              string                           `json:"type"`                           // 组件类型
	VideoSpec         *XijingTemplateVideoSpec         `json:"video_spec,omitempty"`           // 视频组件规格
	TextSpec          *XijingTemplateTextSpec          `json:"text_spec,omitempty"`            // 文本组件规格
	ImageListSpec     *XijingTemplateImageListSpec     `json:"image_list_spec,omitempty"`      // 图片列表组件规格
	ButtonSpec        *XijingTemplateButtonSpec        `json:"button_spec,omitempty"`          // 按钮组件规格
	AppInfoButtonSpec *XijingTemplateAppInfoButtonSpec `json:"app_info_button_spec,omitempty"` // App 信息按钮组件规格
	FixedButtonSpec   *XijingTemplateFixedButtonSpec   `json:"fixed_button_spec,omitempty"`    // 固定按钮组件规格
}

// XijingTemplateGetResp 获取蹊径落地页模板响应
// https://developers.e.qq.com/v3.0/docs/api/xijing_template/get
type XijingTemplateGetResp struct {
	PageTemplateID    string                         `json:"page_template_id"`              // 蹊径落地页模板 id
	PageType          string                         `json:"page_type"`                     // 蹊径落地页类型
	PageName          string                         `json:"page_name"`                     // 落地页名称（管理用）
	PageTitle         string                         `json:"page_title"`                    // 落地页标题（展示用）
	Clipboard         string                         `json:"clipboard,omitempty"`           // 剪贴板内容
	PageDeeplink      string                         `json:"page_deeplink,omitempty"`       // 页面级 deeplink
	MobileAppID       string                         `json:"mobile_app_id,omitempty"`       // AppId（落地页类型为 Android/iOS 时必填）
	ComponentSpecList []*XijingTemplateComponentSpec `json:"component_spec_list,omitempty"` // 组件列表
}
