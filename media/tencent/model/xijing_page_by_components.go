package model

import "errors"

// ========== 蹊径基于组件创建落地页 ==========
// https://developers.e.qq.com/v3.0/docs/api/xijing_page_by_components/add

// 字段长度常量
const (
	MaxXijingPageByCompBgImageIDBytes   = 256 // bg_image_id 最大字节数
	MaxXijingPageByCompClipboardBytes   = 300 // clipboard 最大字节数
	MinXijingPageByCompDeepLinkBytes    = 1   // page_deeplink 最小字节数
	MaxXijingPageByCompDeepLinkBytes    = 200 // page_deeplink 最大字节数
	MinXijingPageByCompPageNameBytes    = 1   // page_name 最小字节数
	MaxXijingPageByCompPageNameBytes    = 20  // page_name 最大字节数
	MinXijingPageByCompPageTitleBytes   = 1   // page_title 最小字节数
	MaxXijingPageByCompPageTitleBytes   = 20  // page_title 最大字节数
	MinXijingPageByCompMobileAppIDBytes = 1   // mobile_app_id 最小字节数
	MaxXijingPageByCompMobileAppIDBytes = 20  // mobile_app_id 最大字节数
	MaxXijingPageByCompPagesCount       = 10  // pages 最大数量
	MaxXijingPageByCompComponentsCount  = 10  // component_spec_list 最大数量
)

// XijingPageByCompCommonSetting 组件通用设置
type XijingPageByCompCommonSetting struct {
	Position           string `json:"position"`              // 按钮位置
	WhiteSpace         int    `json:"white_space"`           // 白边间距
	DistanceToViewPort int    `json:"distance_to_view_port"` // 距视口距离
}

// XijingPageByCompBtnSetting 按钮特定设置
type XijingPageByCompBtnSetting struct {
	Desc            string `json:"desc"`             // 按钮描述/文案
	BackgroundColor string `json:"background_color"` // 按钮背景颜色，如 "#1890ff"
	Color           string `json:"color"`            // 按钮字体颜色，如 "rgb(255, 255, 255)"
}

// XijingPageByCompFixedButtonSpec 固定按钮组件规格
type XijingPageByCompFixedButtonSpec struct {
	ButtonStyle   string                         `json:"button_style"`   // 按钮样式，如 "fixedBtn-1"
	CommonSetting *XijingPageByCompCommonSetting `json:"common_setting"` // 通用设置
	BtnSetting    *XijingPageByCompBtnSetting    `json:"btn_setting"`    // 按钮设置
}

// XijingPageByCompHotArea 图片热区
type XijingPageByCompHotArea struct {
	Width  int `json:"width"`  // 热区宽度
	Height int `json:"height"` // 热区高度
	Top    int `json:"top"`    // 距顶部距离
	Left   int `json:"left"`   // 距左侧距离
}

// XijingPageByCompPadding 图片内边距
type XijingPageByCompPadding struct {
	Top    int `json:"top"`    // 上内边距
	Right  int `json:"right"`  // 右内边距
	Bottom int `json:"bottom"` // 下内边距
	Left   int `json:"left"`   // 左内边距
}

// XijingPageByCompProgressBar 进度条样式
type XijingPageByCompProgressBar struct {
	Color           string `json:"color"`           // 进度条颜色，如 "rgb(255, 255, 255)"
	BackgroundColor string `json:"backgroundColor"` // 进度条背景颜色，如 "#1890ff"
}

// XijingPageByCompImageItem 图片列表中的单张图片
type XijingPageByCompImageItem struct {
	ImageID     string                       `json:"image_id"`     // 图片 id（从图片模块上传获取）
	HotArea     []*XijingPageByCompHotArea   `json:"hot_area"`     // 图片热区列表
	Padding     *XijingPageByCompPadding     `json:"padding"`      // 图片内边距
	ProgressBar *XijingPageByCompProgressBar `json:"progress_bar"` // 进度条样式
}

// XijingPageByCompImageListSpec 图片列表组件规格
type XijingPageByCompImageListSpec struct {
	ImageList []*XijingPageByCompImageItem `json:"image_list"` // 图片列表
}

// XijingPageByCompComponentSpec 组件规格
// 组件类型枚举复用 xijing_template.go 中的 XijingTemplateComponentType* 常量
type XijingPageByCompComponentSpec struct {
	Type            string                           `json:"type"`                        // 组件类型 (必填)
	FixedButtonSpec *XijingPageByCompFixedButtonSpec `json:"fixed_button_spec,omitempty"` // 固定按钮组件规格
	ImageListSpec   *XijingPageByCompImageListSpec   `json:"image_list_spec,omitempty"`   // 图片列表组件规格
}

// XijingPageByCompPage 单个落地页配置
type XijingPageByCompPage struct {
	PageType          string                           `json:"page_type"`               // 落地页类型 (必填)，PAGE_TYPE_XIJING_ANDROID/IOS
	PageName          string                           `json:"page_name"`               // 落地页名称（管理用）(必填)，1-20字节
	PageTitle         string                           `json:"page_title"`              // 落地页标题（展示用）(必填)，1-20字节
	MobileAppID       string                           `json:"mobile_app_id"`           // App ID (必填)，1-20字节
	BgColor           string                           `json:"bg_color"`                // 页面背景颜色 (必填)，如 "rgba(189, 16, 224, 1)"
	BgImageID         string                           `json:"bg_image_id"`             // 页面背景图片 id (必填)，0-256字节
	Clipboard         string                           `json:"clipboard,omitempty"`     // 剪贴板内容，0-300字节
	PageDeeplink      string                           `json:"page_deeplink,omitempty"` // 页面级 deeplink，1-200字节
	ComponentSpecList []*XijingPageByCompComponentSpec `json:"component_spec_list"`     // 组件列表 (必填)，最多10个
}

// XijingPageByCompAddReq 蹊径基于组件创建落地页请求（POST JSON）
// https://developers.e.qq.com/v3.0/docs/api/xijing_page_by_components/add
type XijingPageByCompAddReq struct {
	GlobalReq
	AccountID    int64                   `json:"account_id"`     // 广告主帐号 id (必填)
	IsAutoSubmit int                     `json:"is_auto_submit"` // 是否自动提审 (必填)，0=否，1=是
	Pages        []*XijingPageByCompPage `json:"pages"`          // 落地页配置列表 (必填)，最多10个
}

func (r *XijingPageByCompAddReq) Format() {
	r.GlobalReq.Format()
}

// Validate 验证蹊径基于组件创建落地页请求参数
func (r *XijingPageByCompAddReq) Validate() error {
	if r.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if r.IsAutoSubmit < 0 || r.IsAutoSubmit > 1 {
		return errors.New("is_auto_submit须为0或1")
	}
	if len(r.Pages) == 0 {
		return errors.New("pages为必填，至少包含1个落地页配置")
	}
	if len(r.Pages) > MaxXijingPageByCompPagesCount {
		return errors.New("pages数组长度不能超过10")
	}
	for i, page := range r.Pages {
		if page == nil {
			return errors.New("pages[" + itoa(i) + "]不能为空")
		}
		if page.PageType == "" {
			return errors.New("pages[" + itoa(i) + "].page_type为必填")
		}
		if len(page.PageName) < MinXijingPageByCompPageNameBytes || len(page.PageName) > MaxXijingPageByCompPageNameBytes {
			return errors.New("pages[" + itoa(i) + "].page_name长度须在1-20字节之间")
		}
		if len(page.PageTitle) < MinXijingPageByCompPageTitleBytes || len(page.PageTitle) > MaxXijingPageByCompPageTitleBytes {
			return errors.New("pages[" + itoa(i) + "].page_title长度须在1-20字节之间")
		}
		if len(page.MobileAppID) < MinXijingPageByCompMobileAppIDBytes || len(page.MobileAppID) > MaxXijingPageByCompMobileAppIDBytes {
			return errors.New("pages[" + itoa(i) + "].mobile_app_id长度须在1-20字节之间")
		}
		if page.BgColor == "" {
			return errors.New("pages[" + itoa(i) + "].bg_color为必填")
		}
		if len(page.BgImageID) > MaxXijingPageByCompBgImageIDBytes {
			return errors.New("pages[" + itoa(i) + "].bg_image_id长度不能超过256字节")
		}
		if len(page.Clipboard) > MaxXijingPageByCompClipboardBytes {
			return errors.New("pages[" + itoa(i) + "].clipboard长度不能超过300字节")
		}
		if page.PageDeeplink != "" && (len(page.PageDeeplink) < MinXijingPageByCompDeepLinkBytes || len(page.PageDeeplink) > MaxXijingPageByCompDeepLinkBytes) {
			return errors.New("pages[" + itoa(i) + "].page_deeplink长度须在1-200字节之间")
		}
		if len(page.ComponentSpecList) > MaxXijingPageByCompComponentsCount {
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

// XijingPageByCompAddResp 蹊径基于组件创建落地页响应
// 响应结构与 XijingPageAddResp 相同，复用 xijing_page.go 中的定义
// https://developers.e.qq.com/v3.0/docs/api/xijing_page_by_components/add
type XijingPageByCompAddResp = XijingPageAddResp
