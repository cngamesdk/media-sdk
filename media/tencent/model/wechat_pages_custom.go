package model

import "errors"

// ========== 基于组件创建微信原生页 ==========
// https://developers.e.qq.com/v3.0/docs/api/wechat_pages_custom/add

// 页面组件类型枚举
const (
	WechatPagesCustomElementTypeTopImage     = "TOP_IMAGE"     // 顶部图片组件
	WechatPagesCustomElementTypeTopSlider    = "TOP_SLIDER"    // 顶部轮播图组件
	WechatPagesCustomElementTypeTopVideo     = "TOP_VIDEO"     // 顶部视频组件
	WechatPagesCustomElementTypeImage        = "IMAGE"         // 基础图片组件
	WechatPagesCustomElementTypeSlider       = "SLIDER"        // 基础轮播图组件
	WechatPagesCustomElementTypeVideo        = "VIDEO"         // 基础视频组件
	WechatPagesCustomElementTypeText         = "TEXT"          // 基础文本组件
	WechatPagesCustomElementTypeAppDownload  = "APP_DOWNLOAD"  // App 下载组件
	WechatPagesCustomElementTypeWeapp        = "WEAPP"         // 进入小程序组件
	WechatPagesCustomElementTypeGh           = "GH"            // 关注公众号组件
	WechatPagesCustomElementTypeEnterpriseWx = "ENTERPRISE_WX" // 添加商家微信组件
	WechatPagesCustomElementTypeImageText    = "IMAGE_TEXT"    // 图文复合组件
)

// 全局组件类型枚举
const (
	WechatPagesCustomGlobalElementTypeFloatButton  = "FLOAT_BUTTON"          // 浮层组件
	WechatPagesCustomGlobalElementTypeSideBarFloat = "SIDE_BAR_FLOAT_BUTTON" // 侧边栏浮层组件
)

// 侧边栏浮层组件转化类型枚举
const (
	WechatPagesCustomSideBarElemTypeTel          = "TEL"           // 一键拨号
	WechatPagesCustomSideBarElemTypeWxService    = "WX_SERVICE"    // 微信客服
	WechatPagesCustomSideBarElemTypeEnterpriseWx = "ENTERPRISE_WX" // 商家微信
)

// 图文复合组件跳转方式枚举
const (
	WechatPagesCustomImageTextJumpModeBtn   = "btn_jump"   // 按钮跳转
	WechatPagesCustomImageTextJumpModeTotal = "total_jump" // 整体跳转
)

// 图文复合组件转化类型枚举
const (
	WechatPagesCustomSubElemTypeGh           = "GH"            // 关注公众号
	WechatPagesCustomSubElemTypeEnterpriseWx = "ENTERPRISE_WX" // 商家微信
)

// 广告位置常量
const (
	WechatPagesCustomAdLocationSns = "sns" // 微信朋友圈
	WechatPagesCustomAdLocationGzh = "gzh" // 公众号
)

// 字段长度常量
const (
	MinWechatPagesCustomPageNameBytes      = 1                // page_name 最小字节数
	MaxWechatPagesCustomPageNameBytes      = 120              // page_name 最大字节数
	MinWechatPagesCustomImageIDBytes       = 1                // image_id 最小字节数
	MaxWechatPagesCustomImageIDBytes       = 20               // image_id 最大字节数
	MinWechatPagesCustomVideoIDBytes       = 1                // video_id 最小字节数
	MaxWechatPagesCustomVideoIDBytes       = 20               // video_id 最大字节数
	MinWechatPagesCustomTopSliderCount     = 3                // top_slider image_id_list 最小数量
	MaxWechatPagesCustomTopSliderCount     = 6                // top_slider image_id_list 最大数量
	MinWechatPagesCustomSliderCount        = 2                // slider image_id_list 最小数量
	MaxWechatPagesCustomSliderCount        = 5                // slider image_id_list 最大数量
	MinWechatPagesCustomTextBytes          = 1                // text 最小字节数
	MaxWechatPagesCustomTextBytes          = 30000            // text 最大字节数
	MinWechatPagesCustomAppDownloadTitle   = 1                // app_download title 最小字节数
	MaxWechatPagesCustomAppDownloadTitle   = 30               // app_download title 最大字节数
	MinWechatPagesCustomAppIosIDBytes      = 1                // app_ios_id 最小字节数
	MaxWechatPagesCustomAppIosIDBytes      = 128              // app_ios_id 最大字节数
	MaxWechatPagesCustomDeepLinkURLBytes   = 1024             // deep_link_url 最大字节数
	MinWechatPagesCustomAppAndroidIDBytes  = 1                // app_android_id 最小字节数
	MaxWechatPagesCustomAppAndroidIDBytes  = 128              // app_android_id 最大字节数
	MaxWechatPagesCustomChannelPkgIDBytes  = 128              // app_android_channel_package_id 最大字节数
	MaxWechatPagesCustomAppMarketPkgBytes  = 500              // app_market_package 最大字节数
	MinWechatPagesCustomWeappUsernameBytes = 1                // weapp_username 最小字节数
	MaxWechatPagesCustomWeappUsernameBytes = 50               // weapp_username 最大字节数
	MinWechatPagesCustomWeappPathBytes     = 1                // weapp_path 最小字节数
	MaxWechatPagesCustomWeappPathBytes     = 250              // weapp_path 最大字节数
	MaxWechatPagesCustomBtnTitleBytes      = 10               // 按钮文案最大字节数
	MaxWechatPagesCustomFloatTitleBytes    = 30               // 浮层组件标题最大字节数
	MaxWechatPagesCustomFloatDescBytes     = 42               // 浮层组件描述最大字节数
	MaxWechatPagesCustomWordingBytes       = 4                // 侧边栏浮层 wording 最大字节数
	MaxWechatPagesCustomPhoneNumberBytes   = 11               // 电话号码最大字节数
	MaxWechatPagesCustomWechatServiceURL   = 1024             // 微信客服 URL 最大字节数
	MaxWechatPagesCustomMaterialIDBytes    = 16               // material_id 最大字节数
	MaxWechatPagesCustomItemTitleBytes     = 10               // 图文 title 最大字节数
	MaxWechatPagesCustomItemDescBytes      = 10               // 图文 desc 最大字节数
	MinWechatPagesCustomCorpIDBytes        = 1                // corp_id 最小字节数
	MaxWechatPagesCustomEnterpriseGroupID  = int64(999999999) // enterprise_wx group_id 最大值
)

// ========== 子结构体定义 ==========

// WechatPagesCustomAppIosSpec iOS App 下载元素
type WechatPagesCustomAppIosSpec struct {
	AppIosID         string `json:"app_ios_id"`                   // iOS 应用 id (必填)，1-128字节
	DeepLinkURL      string `json:"deep_link_url,omitempty"`      // App 直达页 URL，1-1024字节
	JumpAppstoreType int    `json:"jump_appstore_type,omitempty"` // AppStore 跳转类型，0=外跳，1=内跳
}

// WechatPagesCustomAppAndroidSpec Android App 下载元素
type WechatPagesCustomAppAndroidSpec struct {
	AppAndroidID               string `json:"app_android_id"`                           // Android 应用 id (必填)，1-128字节
	DeepLinkURL                string `json:"deep_link_url,omitempty"`                  // App 直达页 URL，1-1024字节
	AppAndroidChannelPackageID string `json:"app_android_channel_package_id,omitempty"` // Android 渠道包 id，1-128字节
	AppMarketPackage           string `json:"app_market_package,omitempty"`             // 厂商应用商店包名，1-500字节
}

// WechatPagesCustomAppDownloadSpec App 下载组件元素
type WechatPagesCustomAppDownloadSpec struct {
	Title          string                           `json:"title"`                      // 按钮文案 (必填)，1-30字节
	FontColor      string                           `json:"font_color,omitempty"`       // 按钮字体颜色，7字节，默认#FFFFFF
	BgColor        string                           `json:"bg_color,omitempty"`         // 按钮填充颜色，7字节，默认#07C160
	PaddingTop     int                              `json:"padding_top,omitempty"`      // 距上组件距离，0-100，默认28
	PaddingBottom  int                              `json:"padding_bottom,omitempty"`   // 距下组件距离，0-100，默认28
	AppIosSpec     *WechatPagesCustomAppIosSpec     `json:"app_ios_spec,omitempty"`     // iOS App 下载元素
	AppAndroidSpec *WechatPagesCustomAppAndroidSpec `json:"app_android_spec,omitempty"` // Android App 下载元素
}

// WechatPagesCustomWeappSpec 进入小程序组件元素
type WechatPagesCustomWeappSpec struct {
	WeappUsername       string `json:"weapp_username"`                   // 小程序原始 id (必填)，1-50字节
	WeappPath           string `json:"weapp_path"`                       // 小程序路径 (必填)，1-250字节
	BtnTitle            string `json:"btn_title,omitempty"`              // 按钮文案，1-10字节，默认"进入小程序"
	BtnBorderColorTheme string `json:"btn_border_color_theme,omitempty"` // 边框颜色，7字节，默认#FFFFFF
	BtnBgColorTheme     string `json:"btn_bg_color_theme,omitempty"`     // 按钮填充颜色，7字节，默认#07C160
	FontColor           string `json:"font_color,omitempty"`             // 按钮字体颜色，7字节，默认#FFFFFF
	BtnFontType         int    `json:"btn_font_type,omitempty"`          // 按钮字体样式，0=常规，1=加粗
	UseIcon             int    `json:"use_icon,omitempty"`               // 是否使用图标，0=不使用，1=使用
	PaddingTop          int    `json:"padding_top,omitempty"`            // 距上组件距离，0-100，默认28
	PaddingBottom       int    `json:"padding_bottom,omitempty"`         // 距下组件距离，0-100，默认28
}

// WechatPagesCustomGhSpec 关注公众号组件元素
type WechatPagesCustomGhSpec struct {
	FastFollow          int    `json:"fast_follow,omitempty"`            // 一键关注，0=关闭，1=开启，默认1
	BtnTitle            string `json:"btn_title,omitempty"`              // 按钮文案，1-10字节，默认"关注公众号"
	FontColor           string `json:"font_color,omitempty"`             // 按钮字体颜色，7字节，默认#FFFFFF
	BtnBgColorTheme     string `json:"btn_bg_color_theme,omitempty"`     // 按钮填充颜色，7字节，默认#07C160
	BtnBorderColorTheme string `json:"btn_border_color_theme,omitempty"` // 边框颜色，7字节，默认#FFFFFF
	BtnFontType         int    `json:"btn_font_type,omitempty"`          // 按钮字体样式，0=常规，1=加粗
	PaddingTop          int    `json:"padding_top,omitempty"`            // 距上组件距离，0-100，默认28
	PaddingBottom       int    `json:"padding_bottom,omitempty"`         // 距下组件距离，0-100，默认28
	UseIcon             int    `json:"use_icon,omitempty"`               // 是否使用图标，0=不使用，1=使用
}

// WechatPagesCustomEnterpriseWxSpec 添加商家微信组件元素
type WechatPagesCustomEnterpriseWxSpec struct {
	CorpID              string `json:"corp_id"`                          // 绑定的企业 id (必填)
	GroupID             int64  `json:"group_id"`                         // 客服分组 id (必填)，0-999999999
	BtnTitle            string `json:"btn_title,omitempty"`              // 按钮文案，1-10字节，默认"联系商家"
	FontColor           string `json:"font_color,omitempty"`             // 按钮字体颜色，7字节，默认#FFFFFF
	BtnBgColorTheme     string `json:"btn_bg_color_theme,omitempty"`     // 按钮填充颜色，7字节，默认#07C160
	BtnBorderColorTheme string `json:"btn_border_color_theme,omitempty"` // 边框颜色，7字节，默认#FFFFFF
	BtnFontType         int    `json:"btn_font_type,omitempty"`          // 按钮字体样式，0=常规，1=加粗
	PaddingTop          int    `json:"padding_top,omitempty"`            // 距上组件距离，0-100，默认28
	PaddingBottom       int    `json:"padding_bottom,omitempty"`         // 距下组件距离，0-100，默认28
	UseIcon             int    `json:"use_icon,omitempty"`               // 是否使用图标，0=不使用，1=使用
}

// WechatPagesCustomImageTextItem 图文复合组件配置项
type WechatPagesCustomImageTextItem struct {
	BorderColor      string                             `json:"border_color,omitempty"`       // 边框颜色，7字节，默认#e5e5e5
	TitleColor       string                             `json:"title_color,omitempty"`        // 标题颜色，7字节，默认#353535
	DescColor        string                             `json:"desc_color,omitempty"`         // 描述颜色，7字节，默认#b2b2b2
	BgColor          string                             `json:"bg_color,omitempty"`           // 背景颜色，7字节，默认#ffffff
	JumpMode         string                             `json:"jump_mode,omitempty"`          // 跳转方式，btn_jump/total_jump，默认btn_jump
	MaterialID       string                             `json:"material_id"`                  // 素材 id (必填)，1-16字节
	Title            string                             `json:"title"`                        // 标题 (必填)，1-10字节
	Desc             string                             `json:"desc"`                         // 描述 (必填)，1-10字节
	SubElemType      string                             `json:"sub_elem_type"`                // 转化类型 (必填)，GH/ENTERPRISE_WX
	GhSpec           *WechatPagesCustomGhSpec           `json:"gh_spec,omitempty"`            // 关注公众号组件元素
	EnterpriseWxSpec *WechatPagesCustomEnterpriseWxSpec `json:"enterprise_wx_spec,omitempty"` // 商家微信组件元素
}

// WechatPagesCustomImageTextSpec 图文复合组件元素
type WechatPagesCustomImageTextSpec struct {
	PaddingTop    int                               `json:"padding_top,omitempty"`    // 距上组件距离，0-100，默认20
	PaddingBottom int                               `json:"padding_bottom,omitempty"` // 距下组件距离，0-100，默认20
	AlignMode     int                               `json:"align_mode,omitempty"`     // 双品对齐方式，0=左对齐，1=居中对齐
	ImageTextItem []*WechatPagesCustomImageTextItem `json:"image_text_item"`          // 图文复合组件配置项 (必填)
}

// WechatPagesCustomTopImageSpec 顶部图片组件元素
type WechatPagesCustomTopImageSpec struct {
	ImageID    string `json:"image_id"`              // 图片 id (必填)，1-20字节
	Width      int    `json:"width"`                 // 图片宽度 (必填)
	Height     int    `json:"height"`                // 图片高度 (必填)
	AdLocation string `json:"ad_location"`           // 广告位置 (必填)，sns=朋友圈，gzh=公众号，3字节
	OuterStyle int    `json:"outer_style,omitempty"` // 顶部图片外层类型（sns时有效），0=普通广告，1=卡片广告
}

// WechatPagesCustomTopSliderSpec 顶部轮播图组件元素
type WechatPagesCustomTopSliderSpec struct {
	ImageIDList []string `json:"image_id_list"` // 图片 id 列表 (必填)，3/4/6张
	Width       int      `json:"width"`         // 图片宽度 (必填)
	Height      int      `json:"height"`        // 图片高度 (必填)
}

// WechatPagesCustomTopVideoSpec 顶部视频组件元素
type WechatPagesCustomTopVideoSpec struct {
	VideoID      string `json:"video_id"`                 // 视频 id (必填)，1-20字节
	Width        int    `json:"width"`                    // 视频宽度 (必填)
	Height       int    `json:"height"`                   // 视频高度 (必填)
	AdLocation   string `json:"ad_location"`              // 广告位置 (必填)，sns=朋友圈，gzh=公众号，3字节
	OuterStyle   int    `json:"outer_style,omitempty"`    // 顶部视频外层类型（sns时有效），0=普通，1=基础卡片，2=全幅卡片
	UsedForOuter int    `json:"used_for_outer,omitempty"` // 是否用顶部素材作广告外层，0=是，1=否
}

// WechatPagesCustomImageSpec 基础图片组件元素
type WechatPagesCustomImageSpec struct {
	ImageID       string `json:"image_id"`                 // 图片 id (必填)，1-20字节
	Width         int    `json:"width"`                    // 图片宽度 (必填)
	Height        int    `json:"height"`                   // 图片高度 (必填)
	PaddingTop    int    `json:"padding_top,omitempty"`    // 距上组件距离，0-100
	PaddingBottom int    `json:"padding_bottom,omitempty"` // 距下组件距离，0-100
}

// WechatPagesCustomSliderSpec 基础轮播图组件元素
type WechatPagesCustomSliderSpec struct {
	ImageIDList []string `json:"image_id_list"`      // 图片 id 列表 (必填)，2-5张
	Width       int      `json:"width"`              // 图片宽度 (必填)
	Height      int      `json:"height"`             // 图片高度 (必填)
	SliderStyle int      `json:"slider_style"`       // 轮播样式 (必填)，0=普通，1=卡片
	BgColor     string   `json:"bg_color,omitempty"` // 轮播背景颜色（slider_style=1时有效），7字节，默认#FFFFFF
}

// WechatPagesCustomVideoSpec 基础视频组件元素
type WechatPagesCustomVideoSpec struct {
	VideoID  string `json:"video_id"`            // 视频 id (必填)，1-20字节
	Width    int    `json:"width"`               // 视频宽度 (必填)
	Height   int    `json:"height"`              // 视频高度 (必填)
	InMiddle int    `json:"in_middle,omitempty"` // 居中展示，0=不居中，1=居中
}

// WechatPagesCustomTextSpec 基础文本组件元素
type WechatPagesCustomTextSpec struct {
	Text          string `json:"text"`                     // 文本内容 (必填)，1-30000字节
	PaddingTop    int    `json:"padding_top,omitempty"`    // 距上组件距离，0-100，默认22
	PaddingBottom int    `json:"padding_bottom,omitempty"` // 距下组件距离，0-100，默认22
	FontSize      int    `json:"font_size,omitempty"`      // 字体大小，可选值：14/15/16/18/20/24/36，默认15
	FontColor     string `json:"font_color,omitempty"`     // 字体颜色，7字节，默认#595959
	TextAlignment int    `json:"text_alignment,omitempty"` // 文字对齐，0=左对齐，1=居中，2=右对齐
	FontStyle     int    `json:"font_style,omitempty"`     // 字体样式，0=常规，1=加粗
}

// WechatPagesCustomElementSpec 页面组件素材内容
type WechatPagesCustomElementSpec struct {
	ElementType      string                             `json:"element_type"`                 // 组件类型 (必填)
	TopImageSpec     *WechatPagesCustomTopImageSpec     `json:"top_image_spec,omitempty"`     // 顶部图片组件元素
	TopSliderSpec    *WechatPagesCustomTopSliderSpec    `json:"top_slider_spec,omitempty"`    // 顶部轮播图组件元素
	TopVideoSpec     *WechatPagesCustomTopVideoSpec     `json:"top_video_spec,omitempty"`     // 顶部视频组件元素
	ImageSpec        *WechatPagesCustomImageSpec        `json:"image_spec,omitempty"`         // 基础图片组件元素
	SliderSpec       *WechatPagesCustomSliderSpec       `json:"slider_spec,omitempty"`        // 基础轮播图组件元素
	VideoSpec        *WechatPagesCustomVideoSpec        `json:"video_spec,omitempty"`         // 基础视频组件元素
	TextSpec         *WechatPagesCustomTextSpec         `json:"text_spec,omitempty"`          // 基础文本组件元素
	AppDownloadSpec  *WechatPagesCustomAppDownloadSpec  `json:"app_download_spec,omitempty"`  // App 下载组件元素
	WeappSpec        *WechatPagesCustomWeappSpec        `json:"weapp_spec,omitempty"`         // 进入小程序组件元素
	GhSpec           *WechatPagesCustomGhSpec           `json:"gh_spec,omitempty"`            // 关注公众号组件元素
	EnterpriseWxSpec *WechatPagesCustomEnterpriseWxSpec `json:"enterprise_wx_spec,omitempty"` // 添加商家微信组件元素
	ImageTextSpec    *WechatPagesCustomImageTextSpec    `json:"image_text_spec,omitempty"`    // 图文复合组件元素
}

// WechatPagesCustomPageSpec 页面配置
type WechatPagesCustomPageSpec struct {
	BgColor              string                          `json:"bg_color,omitempty"`      // 背景颜色，7字节
	PageElementsSpecList []*WechatPagesCustomElementSpec `json:"page_elements_spec_list"` // 页面组件列表 (必填)
}

// WechatPagesCustomTelSpec 一键拨号参数结构
type WechatPagesCustomTelSpec struct {
	PhoneNumber string `json:"phone_number"` // 电话号码 (必填)，1-11字节
	PhoneType   string `json:"phone_type"`   // 电话类型 (必填)，1字节，0=座机，1=手机
}

// WechatPagesCustomWechatServiceSpec 微信客服元素
type WechatPagesCustomWechatServiceSpec struct {
	WechatServiceURL string `json:"wechat_service_url"` // 微信客服 URL (必填)，1-1024字节
}

// WechatPagesCustomSideBarFloatSpec 侧边栏浮层组件元素
type WechatPagesCustomSideBarFloatSpec struct {
	Wording           string                              `json:"wording,omitempty"`             // 按钮文案，1-4字节
	TitleColor        string                              `json:"title_color,omitempty"`         // 字体颜色，7字节，默认#000000
	ElemType          string                              `json:"elem_type"`                     // 转化类型 (必填)，TEL/WX_SERVICE/ENTERPRISE_WX
	TelSpec           *WechatPagesCustomTelSpec           `json:"tel_spec,omitempty"`            // 一键拨号参数（elem_type=TEL时必填）
	WechatServiceSpec *WechatPagesCustomWechatServiceSpec `json:"wechat_service_spec,omitempty"` // 微信客服元素（elem_type=WX_SERVICE时必填）
	EnterpriseWxSpec  *WechatPagesCustomEnterpriseWxSpec  `json:"enterprise_wx_spec,omitempty"`  // 商家微信元素（elem_type=ENTERPRISE_WX时必填）
}

// WechatPagesCustomFloatButtonSpec 浮层组件元素
type WechatPagesCustomFloatButtonSpec struct {
	StyleType        int                                `json:"style_type"`                   // 浮层组件样式 (必填)，0=图片+标题+描述，1=标题+描述，2=标题
	ImageID          string                             `json:"image_id,omitempty"`           // 浮层组件图片 id，1-20字节（style_type=0时必填）
	Title            string                             `json:"title"`                        // 浮层组件标题 (必填)，1-30字节
	TitleColor       string                             `json:"title_color,omitempty"`        // 标题颜色，7字节，默认#171717
	Desc             string                             `json:"desc,omitempty"`               // 浮层组件描述，1-42字节（style_type=0/1时必填）
	DescColor        string                             `json:"desc_color,omitempty"`         // 描述颜色，7字节，默认#4C4C4C
	AppearType       int                                `json:"appear_type,omitempty"`        // 出现方式，0=进入页面时，1=滑动时，默认0
	DisappearType    int                                `json:"disappear_type,omitempty"`     // 消失方式，0=不消失，1=页面底部消失，默认0
	ForbidPageList   []int                              `json:"forbid_page_list,omitempty"`   // 不展示浮层按钮的页码列表，值>=1
	ElementType      string                             `json:"element_type"`                 // 浮层按钮内部组件类型 (必填)，APP_DOWNLOAD/WEAPP/GH/ENTERPRISE_WX
	AppDownloadSpec  *WechatPagesCustomAppDownloadSpec  `json:"app_download_spec,omitempty"`  // App 下载组件元素
	WeappSpec        *WechatPagesCustomWeappSpec        `json:"weapp_spec,omitempty"`         // 进入小程序组件元素
	GhSpec           *WechatPagesCustomGhSpec           `json:"gh_spec,omitempty"`            // 关注公众号组件元素
	EnterpriseWxSpec *WechatPagesCustomEnterpriseWxSpec `json:"enterprise_wx_spec,omitempty"` // 商家微信组件元素
}

// WechatPagesCustomGlobalElementSpec 全局组件素材内容
type WechatPagesCustomGlobalElementSpec struct {
	ElementType      string                             `json:"element_type"`                  // 全局组件类型 (必填)，FLOAT_BUTTON/SIDE_BAR_FLOAT_BUTTON
	FloatButtonSpec  *WechatPagesCustomFloatButtonSpec  `json:"float_button_spec,omitempty"`   // 浮层组件元素
	SideBarFloatSpec *WechatPagesCustomSideBarFloatSpec `json:"side_bar_float_spec,omitempty"` // 侧边栏浮层组件元素
}

// WechatPagesCustomGlobalSpec 全局元素配置
type WechatPagesCustomGlobalSpec struct {
	GlobalElementsSpecList []*WechatPagesCustomGlobalElementSpec `json:"global_elements_spec_list"` // 全局组件列表 (必填)
}

// WechatPagesCustomAddReq 基于组件创建微信原生页请求（POST JSON）
// https://developers.e.qq.com/v3.0/docs/api/wechat_pages_custom/add
type WechatPagesCustomAddReq struct {
	GlobalReq
	AccountID        int64                        `json:"account_id"`            // 广告主帐号 id (必填)
	PageName         string                       `json:"page_name"`             // 落地页名称 (必填)，1-120字节
	PageSpecsList    []*WechatPagesCustomPageSpec `json:"page_specs_list"`       // 页面列表 (必填)
	GlobalSpec       *WechatPagesCustomGlobalSpec `json:"global_spec,omitempty"` // 全局元素配置
	ShareContentSpec *WechatPageShareContentSpec  `json:"share_content_spec"`    // 微信原生页分享信息 (必填)
}

func (p *WechatPagesCustomAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证基于组件创建微信原生页请求参数
func (p *WechatPagesCustomAddReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if len(p.PageName) < MinWechatPagesCustomPageNameBytes || len(p.PageName) > MaxWechatPagesCustomPageNameBytes {
		return errors.New("page_name长度须在1-120字节之间")
	}
	if len(p.PageSpecsList) == 0 {
		return errors.New("page_specs_list为必填，至少包含1个页面")
	}
	for i, page := range p.PageSpecsList {
		if page == nil {
			return errors.New("page_specs_list[" + itoa(i) + "]不能为空")
		}
		if len(page.PageElementsSpecList) == 0 {
			return errors.New("page_specs_list[" + itoa(i) + "].page_elements_spec_list为必填，至少包含1个组件")
		}
		for j, elem := range page.PageElementsSpecList {
			if elem == nil {
				return errors.New("page_specs_list[" + itoa(i) + "].page_elements_spec_list[" + itoa(j) + "]不能为空")
			}
			if elem.ElementType == "" {
				return errors.New("page_specs_list[" + itoa(i) + "].page_elements_spec_list[" + itoa(j) + "].element_type为必填")
			}
		}
	}
	if p.ShareContentSpec == nil {
		return errors.New("share_content_spec为必填")
	}
	if p.ShareContentSpec.ShareTitle == "" {
		return errors.New("share_content_spec.share_title为必填")
	}
	if p.ShareContentSpec.ShareDescription == "" {
		return errors.New("share_content_spec.share_description为必填")
	}
	return p.GlobalReq.Validate()
}

// WechatPagesCustomAddResp 基于组件创建微信原生页响应
// https://developers.e.qq.com/v3.0/docs/api/wechat_pages_custom/add
type WechatPagesCustomAddResp struct {
	PageID int64 `json:"page_id"` // 落地页 id
}
