package model

import "errors"

// ========== 蹊径创建互动落地页 ==========
// https://developers.e.qq.com/v3.0/docs/api/xijing_page_interactive/add

// 互动落地页类型枚举
const (
	XijingInteractivePageTypeCompressedPackage = "COMPRESSED_PACKAGE" // 压缩包类型
)

// 转化类型枚举
const (
	XijingInteractiveTransformTypeAppDownload = "TRANSFORM_APP_DOWNLOAD" // App 下载
	XijingInteractiveTransformTypeWebsiteLink = "TRANSFORM_WEBSITE_LINK" // 网页链接
)

// 字段长度常量
const (
	MaxXijingInteractiveFileSize         = 7340032 // ZIP 文件最大字节数（7MB）
	MaxXijingInteractiveFileNameBytes    = 32      // 文件名最大字节数
	MinXijingInteractivePageTitleBytes   = 1       // page_title 最小字节数
	MaxXijingInteractivePageTitleBytes   = 20      // page_title 最大字节数
	MinXijingInteractivePageNameBytes    = 1       // page_name 最小字节数
	MaxXijingInteractivePageNameBytes    = 20      // page_name 最大字节数
	MinXijingInteractiveMobileAppIDBytes = 1       // mobile_app_id 最小字节数
	MaxXijingInteractiveMobileAppIDBytes = 20      // mobile_app_id 最大字节数
	MinXijingInteractivePageConfigBytes  = 1       // page_config 最小字节数
	MaxXijingInteractivePageConfigBytes  = 8000    // page_config 最大字节数
)

// XijingPageInteractiveAddReq 蹊径创建互动落地页请求（POST multipart）
// https://developers.e.qq.com/v3.0/docs/api/xijing_page_interactive/add
type XijingPageInteractiveAddReq struct {
	GlobalReq
	AccountID           int64  `json:"account_id"`            // 广告主帐号 id (必填)
	IsAutoSubmit        int    `json:"is_auto_submit"`        // 是否自动提审 (必填)，0=否，1=是
	PageType            string `json:"page_type"`             // 落地页类型 (必填)，PAGE_TYPE_XIJING_ANDROID/IOS
	InteractivePageType string `json:"interactive_page_type"` // 互动落地页类型 (必填)，COMPRESSED_PACKAGE
	PageTitle           string `json:"page_title"`            // 落地页标题（展示用）(必填)，1-20字节
	PageName            string `json:"page_name"`             // 落地页名称（管理用）(必填)，1-20字节
	MobileAppID         string `json:"mobile_app_id"`         // App ID (必填)，1-20字节
	TransformType       string `json:"transform_type"`        // 转化类型，可选
	PageConfig          string `json:"page_config"`           // 页面配置 JSON，1-8000字节，可选
	FileName            string `json:"-"`                     // ZIP 文件名（上传时使用）
	FileData            []byte `json:"-"`                     // ZIP 文件内容（上传时使用）
}

func (r *XijingPageInteractiveAddReq) Format() {
	r.GlobalReq.Format()
}

// Validate 验证蹊径创建互动落地页请求参数
func (r *XijingPageInteractiveAddReq) Validate() error {
	if r.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if r.IsAutoSubmit < 0 || r.IsAutoSubmit > 1 {
		return errors.New("is_auto_submit须为0或1")
	}
	if r.PageType == "" {
		return errors.New("page_type为必填")
	}
	if r.InteractivePageType == "" {
		return errors.New("interactive_page_type为必填")
	}
	if len(r.PageTitle) < MinXijingInteractivePageTitleBytes || len(r.PageTitle) > MaxXijingInteractivePageTitleBytes {
		return errors.New("page_title长度须在1-20字节之间")
	}
	if len(r.PageName) < MinXijingInteractivePageNameBytes || len(r.PageName) > MaxXijingInteractivePageNameBytes {
		return errors.New("page_name长度须在1-20字节之间")
	}
	if len(r.MobileAppID) < MinXijingInteractiveMobileAppIDBytes || len(r.MobileAppID) > MaxXijingInteractiveMobileAppIDBytes {
		return errors.New("mobile_app_id长度须在1-20字节之间")
	}
	if len(r.FileData) > MaxXijingInteractiveFileSize {
		return errors.New("file文件大小不能超过7MB")
	}
	if len(r.FileName) > MaxXijingInteractiveFileNameBytes {
		return errors.New("文件名长度不能超过32字节")
	}
	if r.PageConfig != "" && (len(r.PageConfig) < MinXijingInteractivePageConfigBytes || len(r.PageConfig) > MaxXijingInteractivePageConfigBytes) {
		return errors.New("page_config长度须在1-8000字节之间")
	}
	return r.GlobalReq.Validate()
}

// XijingPageInteractiveAddResp 蹊径创建互动落地页响应
// 响应结构与 XijingPageAddResp 相同，复用 xijing_page.go 中的定义
// https://developers.e.qq.com/v3.0/docs/api/xijing_page_interactive/add
type XijingPageInteractiveAddResp = XijingPageAddResp
