package model

import "errors"

// ========== 官方落地页-基于组件创建 ==========
// https://developers.e.qq.com/v3.0/docs/api/official_landing_page_component/add

// 组件类型枚举
const (
	OfficialLandingPageElementTypeHeadOutsideMaterial  = "HeadOutsideMaterial"  // 顶部外显素材组件
	OfficialLandingPageElementTypeHeadVideo            = "HeadVideo"            // 顶部视频组件
	OfficialLandingPageElementTypeHeadImage            = "HeadImage"            // 顶部图片组件
	OfficialLandingPageElementTypeHead180PanoramaImage = "Head180PanoramaImage" // 顶部 180 度全景图组件
	OfficialLandingPageElementTypeHeadCarousel         = "HeadCarousel"         // 顶部轮播图组件
	OfficialLandingPageElementTypeHeadText             = "HeadText"             // 顶部文本组件
	OfficialLandingPageElementTypeBlockVideo           = "BlockVideo"           // 视频组件
	OfficialLandingPageElementTypeBlockImage           = "BlockImage"           // 图片组件
	OfficialLandingPageElementTypeBlockCarousel        = "BlockCarousel"        // 轮播图组件
	OfficialLandingPageElementTypeBlockText            = "BlockText"            // 简单文本组件
	OfficialLandingPageElementTypeBlockDivider         = "BlockDivider"         // 分割线组件
	OfficialLandingPageElementTypeBlockCustomPage      = "BlockCustomPage"      // 内嵌网页组件
	OfficialLandingPageElementTypeBlockButton          = "BlockButton"          // 按钮组件
	OfficialLandingPageElementTypeFixedButton          = "FixedButton"          // 底部悬浮按钮组件
	OfficialLandingPageElementTypeRightFixedButton     = "RightFixedButton"     // 侧边悬浮按钮组件
	OfficialLandingPageElementTypeBlockShelf           = "BlockShelf"           // 图文复合组件
	OfficialLandingPageElementTypeBlockShelfGroup      = "BlockShelfGroup"      // 双图文复合组件
)

// 转化类型枚举
const (
	OfficialLandingPageConvertTypeUnknown     = "unknown"     // 暂不支持
	OfficialLandingPageConvertTypeDownload    = "download"    // 下载 APP
	OfficialLandingPageConvertTypeLink        = "link"        // 跳转链接
	OfficialLandingPageConvertTypeWeapp       = "weapp"       // 打开小程序
	OfficialLandingPageConvertTypeGh          = "gh"          // 关注公众号
	OfficialLandingPageConvertTypeFollowVideo = "followVideo" // 关注视频号
)

// 字段长度/范围常量
const (
	MinOfficialLandingPagePageNameBytes             = 1
	MaxOfficialLandingPagePageNameBytes             = 20
	MinOfficialLandingPagePageTitleBytes            = 1
	MaxOfficialLandingPagePageTitleBytes            = 20
	MinOfficialLandingPageIosAppIDBytes             = 1
	MaxOfficialLandingPageIosAppIDBytes             = 64
	MinOfficialLandingPageAndroidAppIDBytes         = 1
	MaxOfficialLandingPageAndroidAppIDBytes         = 64
	MinOfficialLandingPageShareTitleBytes           = 1
	MaxOfficialLandingPageShareTitleBytes           = 14
	MinOfficialLandingPageShareDescriptionBytes     = 1
	MaxOfficialLandingPageShareDescriptionBytes     = 20
	MinOfficialLandingPageShareThumbMaterialIDBytes = 1
	MaxOfficialLandingPageShareThumbMaterialIDBytes = 20
	MinOfficialLandingPageBgColorBytes              = 1
	MaxOfficialLandingPageBgColorBytes              = 7
	MaxOfficialLandingPageElementsCount             = 10
	MinOfficialLandingPageMaterialIDBytes           = 1
	MaxOfficialLandingPageMaterialIDBytes           = 64
	MinOfficialLandingPageCarouselCount             = 2
	MaxOfficialLandingPageCarouselCount             = 6
	MaxOfficialLandingPageHotAreaCount              = 3
	MaxOfficialLandingPageMargin                    = 50
	MaxOfficialLandingPageBlockTextContentBytes     = 10000
	MaxOfficialLandingPageDividerThicknessMin       = 1
	MaxOfficialLandingPageDividerThicknessMax       = 500
)

// ========== 转化配置结构 ==========

// OfficialLandingPageConvertUnknown 转化-暂不支持
type OfficialLandingPageConvertUnknown struct {
	ConvertType string `json:"convert_type"` // 值为 unknown
}

// OfficialLandingPageConvertDownload 转化-下载 APP
type OfficialLandingPageConvertDownload struct {
	ConvertType        string `json:"convert_type"`                   // 值为 download
	DeeplinkUrlIos     string `json:"deeplink_url_ios,omitempty"`     // ios 直达链接，1-2048字节
	DeeplinkUrlAndroid string `json:"deeplink_url_android,omitempty"` // android 直达链接，1-2048字节
}

// OfficialLandingPageConvertLink 转化-跳转链接
type OfficialLandingPageConvertLink struct {
	ConvertType        string `json:"convert_type"`                   // 值为 link
	Src                string `json:"src"`                            // url 链接，https开头，1-2048字节 (必填)
	AppIdIos           string `json:"app_id_ios,omitempty"`           // ios 直达链接跳转的 appId，1-64字节
	DeeplinkUrlIos     string `json:"deeplink_url_ios,omitempty"`     // ios 直达链接，1-2048字节
	AppIdAndroid       string `json:"app_id_android,omitempty"`       // android 直达链接跳转的 appId，1-64字节
	DeeplinkUrlAndroid string `json:"deeplink_url_android,omitempty"` // android 直达链接，1-2048字节
}

// OfficialLandingPageConvertWeapp 转化-打开小程序
type OfficialLandingPageConvertWeapp struct {
	ConvertType string `json:"convert_type"`          // 值为 weapp
	WeappId     string `json:"weapp_id"`              // 原始 id，gh_xxxx，1-128字节 (必填)
	WeappPath   string `json:"weapp_path"`            // 小程序路径，1-1024字节 (必填)
	BackupLink  string `json:"backup_link,omitempty"` // 备用链接，https开头，1-2048字节
}

// OfficialLandingPageConvertGh 转化-关注公众号
type OfficialLandingPageConvertGh struct {
	ConvertType string `json:"convert_type"`           // 值为 gh
	Appid       string `json:"appid"`                  // wxid，以 wx 开头，1-128字节 (必填)
	OneClick    bool   `json:"one_click,omitempty"`    // 是否开启一键关注，默认 false
	DirectFocus bool   `json:"direct_focus,omitempty"` // 是否开启立即关注，仅图文复合系列组件可用，默认 false
}

// OfficialLandingPageConvertFollowVideo 转化-关注视频号
type OfficialLandingPageConvertFollowVideo struct {
	ConvertType string `json:"convert_type"`        // 值为 followVideo
	NickName    string `json:"nick_name"`           // 视频号名称，1-256字节 (必填)
	OneClick    bool   `json:"one_click,omitempty"` // 是否开启一键关注，默认 false
}

// ========== 热区配置 ==========

// OfficialLandingPageHotArea 热区配置
type OfficialLandingPageHotArea struct {
	Left               float64                                `json:"left"`                          // 热区左上顶点距左侧百分比，0-100
	Top                float64                                `json:"top"`                           // 热区左上顶点距顶部百分比，0-100
	Width              float64                                `json:"width"`                         // 热区宽度占图片宽度百分比，0-100
	Height             float64                                `json:"height"`                        // 热区高度占图片高度百分比，0-100
	ConvertUnknown     *OfficialLandingPageConvertUnknown     `json:"convert_unknown,omitempty"`     // 转化-暂不支持
	ConvertDownload    *OfficialLandingPageConvertDownload    `json:"convert_download,omitempty"`    // 转化-下载 APP
	ConvertLink        *OfficialLandingPageConvertLink        `json:"convert_link,omitempty"`        // 转化-跳转链接
	ConvertWeapp       *OfficialLandingPageConvertWeapp       `json:"convert_weapp,omitempty"`       // 转化-打开小程序
	ConvertGh          *OfficialLandingPageConvertGh          `json:"convert_gh,omitempty"`          // 转化-关注公众号
	ConvertFollowVideo *OfficialLandingPageConvertFollowVideo `json:"convert_followVideo,omitempty"` // 转化-关注视频号
}

// ========== 各组件配置 ==========

// OfficialLandingPageHeadVideoConfig 顶部视频组件配置
type OfficialLandingPageHeadVideoConfig struct {
	MaterialId string `json:"material_id"` // 视频素材 id，1-64字节
}

// OfficialLandingPageHeadImageConfig 顶部图片组件配置
type OfficialLandingPageHeadImageConfig struct {
	MaterialId string `json:"material_id"` // 图片素材 id，1-64字节
}

// OfficialLandingPageHead180PanoramaImageConfig 顶部 180 度全景图组件配置
type OfficialLandingPageHead180PanoramaImageConfig struct {
	MaterialId string `json:"material_id"` // 图片素材 id，1-64字节，尺寸 4096x2048px
}

// OfficialLandingPageHeadCarouselConfig 顶部轮播图组件配置
type OfficialLandingPageHeadCarouselConfig struct {
	MaterialIdList []string `json:"material_id_list"` // 图片素材 id 列表，2-6个，每个素材尺寸需相同
	Type           string   `json:"type,omitempty"`   // 轮播样式：full（平滑滚动）/ center（居中轮播），默认 full
}

// OfficialLandingPageTextContent 文本内容（标题/详情）
type OfficialLandingPageTextContent struct {
	Content string `json:"content,omitempty"` // 文案内容，不支持 emoji，1-14字节
	Color   string `json:"color,omitempty"`   // 颜色，1-7字节，默认 #000000
}

// OfficialLandingPageHeadTextConfig 顶部文本组件配置
type OfficialLandingPageHeadTextConfig struct {
	Title  *OfficialLandingPageTextContent `json:"title,omitempty"`  // 标题
	Detail *OfficialLandingPageTextContent `json:"detail,omitempty"` // 内容
}

// OfficialLandingPageBlockVideoConfig 普通视频组件配置
type OfficialLandingPageBlockVideoConfig struct {
	MaterialId string `json:"material_id"` // 视频素材 id，1-64字节
}

// OfficialLandingPageBlockImageConfig 图片组件配置
type OfficialLandingPageBlockImageConfig struct {
	MaterialId   string                        `json:"material_id"`             // 图片素材 id，1-64字节
	Areas        []*OfficialLandingPageHotArea `json:"areas,omitempty"`         // 热区列表，0-3个
	MarginTop    float64                       `json:"margin_top,omitempty"`    // 与上一个组件的距离，0-50
	MarginBottom float64                       `json:"margin_bottom,omitempty"` // 与下一个组件的距离，0-50
}

// OfficialLandingPageBlockCarouselConfig 轮播图组件配置
type OfficialLandingPageBlockCarouselConfig struct {
	MaterialIdList []string                      `json:"material_id_list"`        // 图片素材 id 列表，2-6个
	Type           string                        `json:"type,omitempty"`          // 轮播样式：full / center，默认 full
	Areas          []*OfficialLandingPageHotArea `json:"areas,omitempty"`         // 热区列表，0-3个
	MarginTop      float64                       `json:"margin_top,omitempty"`    // 与上一个组件的距离，0-50
	MarginBottom   float64                       `json:"margin_bottom,omitempty"` // 与下一个组件的距离，0-50
}

// OfficialLandingPageBlockTextConfig 简单文本组件配置
type OfficialLandingPageBlockTextConfig struct {
	Content      string  `json:"content,omitempty"`       // 文本文案，1-10000字节
	Color        string  `json:"color,omitempty"`         // 字体颜色，1-7字节，默认 #000000
	FontSize     string  `json:"font_size,omitempty"`     // 字体大小：12/14/16/18/20/24/32/36，默认 16
	FontWeight   string  `json:"font_weight,omitempty"`   // 是否加粗：bold/normal，默认 normal
	FontStyle    string  `json:"font_style,omitempty"`    // 字体样式：italic/normal，默认 normal
	TextAlign    string  `json:"text_align,omitempty"`    // 对齐方式：left/right/center，默认 left
	MarginTop    float64 `json:"margin_top,omitempty"`    // 与上一个组件的距离，0-50
	MarginBottom float64 `json:"margin_bottom,omitempty"` // 与下一个组件的距离，0-50
}

// OfficialLandingPageBlockDividerConfig 分割线组件配置
type OfficialLandingPageBlockDividerConfig struct {
	LineColor    string  `json:"line_color,omitempty"`    // 线条颜色，1-7字节，默认 #000000
	Thickness    float64 `json:"thickness,omitempty"`     // 线条粗细，1-500，默认 1
	Type         string  `json:"type,omitempty"`          // 线条类型：solid/dashed，默认 solid
	MarginTop    float64 `json:"margin_top,omitempty"`    // 与上一个组件的距离，0-50
	MarginBottom float64 `json:"margin_bottom,omitempty"` // 与下一个组件的距离，0-50
}

// OfficialLandingPageBlockCustomPageConfig 内嵌网页组件配置
type OfficialLandingPageBlockCustomPageConfig struct {
	Src string `json:"src,omitempty"` // url 链接，https开头，1-2048字节
}

// OfficialLandingPageBlockButtonConfig 普通按钮组件配置
type OfficialLandingPageBlockButtonConfig struct {
	Color              string                                 `json:"color,omitempty"`               // 字体颜色，1-7字节，默认 #FFFFFF
	BackgroundColor    string                                 `json:"background_color,omitempty"`    // 背景颜色，1-7字节，默认 #296BEF
	Content            string                                 `json:"content,omitempty"`             // 按钮文案，1-8字节
	HasIcon            bool                                   `json:"has_icon,omitempty"`            // 是否显示图标，默认 false
	ButtonSize         string                                 `json:"button_size,omitempty"`         // 按钮大小：small/middle/large，默认 large
	MarginTop          float64                                `json:"margin_top,omitempty"`          // 与上一个组件的距离，0-50
	MarginBottom       float64                                `json:"margin_bottom,omitempty"`       // 与下一个组件的距离，0-50
	ConvertUnknown     *OfficialLandingPageConvertUnknown     `json:"convert_unknown,omitempty"`     // 转化-暂不支持
	ConvertDownload    *OfficialLandingPageConvertDownload    `json:"convert_download,omitempty"`    // 转化-下载 APP
	ConvertLink        *OfficialLandingPageConvertLink        `json:"convert_link,omitempty"`        // 转化-跳转链接
	ConvertWeapp       *OfficialLandingPageConvertWeapp       `json:"convert_weapp,omitempty"`       // 转化-打开小程序
	ConvertGh          *OfficialLandingPageConvertGh          `json:"convert_gh,omitempty"`          // 转化-关注公众号
	ConvertFollowVideo *OfficialLandingPageConvertFollowVideo `json:"convert_followVideo,omitempty"` // 转化-关注视频号
}

// OfficialLandingPageFixedButtonConfig 底部悬浮按钮组件配置
type OfficialLandingPageFixedButtonConfig struct {
	IconMaterialId     string                                 `json:"icon_material_id,omitempty"`        // 图片素材 id，750x750px，1-64字节
	IconShape          string                                 `json:"icon_shape,omitempty"`              // 图片形状：rect/round，默认 rect
	Title              string                                 `json:"title,omitempty"`                   // 标题文案，1-10字节
	Desc               string                                 `json:"desc,omitempty"`                    // 描述文案，1-20字节
	ButtonContent      string                                 `json:"button_content,omitempty"`          // 按钮文案，1-5字节
	TitleColor         string                                 `json:"title_color,omitempty"`             // 标题颜色，1-7字节，默认 #000000
	DescColor          string                                 `json:"desc_color,omitempty"`              // 字体颜色，1-7字节，默认 #4C4C4C
	ButtonFontColor    string                                 `json:"button_font_color,omitempty"`       // 按钮文案颜色，1-7字节，默认 #FFFFFF
	ButtonBgColor      string                                 `json:"button_background_color,omitempty"` // 按钮背景颜色，1-7字节，默认 #296BEF
	BackgroundColor    string                                 `json:"background_color,omitempty"`        // 组件颜色，1-7字节，默认 #f0f0f0
	ConvertUnknown     *OfficialLandingPageConvertUnknown     `json:"convert_unknown,omitempty"`         // 转化-暂不支持
	ConvertDownload    *OfficialLandingPageConvertDownload    `json:"convert_download,omitempty"`        // 转化-下载 APP
	ConvertLink        *OfficialLandingPageConvertLink        `json:"convert_link,omitempty"`            // 转化-跳转链接
	ConvertWeapp       *OfficialLandingPageConvertWeapp       `json:"convert_weapp,omitempty"`           // 转化-打开小程序
	ConvertGh          *OfficialLandingPageConvertGh          `json:"convert_gh,omitempty"`              // 转化-关注公众号
	ConvertFollowVideo *OfficialLandingPageConvertFollowVideo `json:"convert_followVideo,omitempty"`     // 转化-关注视频号
}

// OfficialLandingPageRightFixedButtonContent 侧边悬浮按钮组件内容项
type OfficialLandingPageRightFixedButtonContent struct {
	Title              string                                 `json:"title,omitempty"`                   // 标题文案，1-4字节
	TitleColor         string                                 `json:"title_color,omitempty"`             // 标题颜色，1-7字节，默认 #000000
	ButtonContent      string                                 `json:"button_content,omitempty"`          // 按钮文案，1-4字节
	ButtonFontColor    string                                 `json:"button_font_color,omitempty"`       // 按钮文案颜色，1-7字节，默认 #FFFFFF
	ButtonBgColor      string                                 `json:"button_background_color,omitempty"` // 按钮背景颜色，1-7字节，默认 #296BEF
	IconColor          string                                 `json:"icon_color,omitempty"`              // 图标颜色，1-7字节，默认 #000000
	ConvertUnknown     *OfficialLandingPageConvertUnknown     `json:"convert_unknown,omitempty"`         // 转化-暂不支持
	ConvertDownload    *OfficialLandingPageConvertDownload    `json:"convert_download,omitempty"`        // 转化-下载 APP
	ConvertLink        *OfficialLandingPageConvertLink        `json:"convert_link,omitempty"`            // 转化-跳转链接
	ConvertWeapp       *OfficialLandingPageConvertWeapp       `json:"convert_weapp,omitempty"`           // 转化-打开小程序
	ConvertGh          *OfficialLandingPageConvertGh          `json:"convert_gh,omitempty"`              // 转化-关注公众号
	ConvertFollowVideo *OfficialLandingPageConvertFollowVideo `json:"convert_followVideo,omitempty"`     // 转化-关注视频号
}

// OfficialLandingPageRightFixedButtonConfig 侧边悬浮按钮组件配置
type OfficialLandingPageRightFixedButtonConfig struct {
	BackgroundColor string                                        `json:"background_color,omitempty"` // 组件颜色，1-7字节，默认 #FFFFFF
	Type            string                                        `json:"type,omitempty"`             // 组件样式：with-icon/with-button，默认 with-icon
	Content         []*OfficialLandingPageRightFixedButtonContent `json:"content,omitempty"`          // 侧边悬浮按钮组件内容，数组长度为 1
}

// OfficialLandingPageBlockShelfConfig 图文复合组件配置
type OfficialLandingPageBlockShelfConfig struct {
	CardType           string                                 `json:"card_type,omitempty"`               // 样式类型：card（横版图文）/title（竖版图文），默认 card
	StyleType          string                                 `json:"style_type,omitempty"`              // 卡片大小，仅 card_type=card 时生效：0（中卡）/1（小卡）
	Title              string                                 `json:"title,omitempty"`                   // 标题文案，1-10字节
	TitleColor         string                                 `json:"title_color,omitempty"`             // 标题颜色，1-7字节，默认 #000000
	Desc               string                                 `json:"desc,omitempty"`                    // 描述文案，1-20字节
	DescColor          string                                 `json:"desc_color,omitempty"`              // 字体颜色，1-7字节，默认 #7F7F7F
	ButtonContent      string                                 `json:"button_content,omitempty"`          // 按钮文案，1-5字节
	ButtonFontColor    string                                 `json:"button_font_color,omitempty"`       // 按钮文案颜色，1-7字节，默认 #FFFFFF
	ButtonBgColor      string                                 `json:"button_background_color,omitempty"` // 按钮背景颜色，1-7字节，默认 #296BEF
	CardBgColor        string                                 `json:"card_background_color,omitempty"`   // 卡片背景颜色，1-7字节，默认 #FFFFFF
	CardBorderColor    string                                 `json:"card_border_color,omitempty"`       // 卡片边框颜色，1-7字节，默认 #E5E5E5
	IconMaterialId     string                                 `json:"icon_material_id,omitempty"`        // 图片素材 id，1-64字节
	MarginTop          float64                                `json:"margin_top,omitempty"`              // 与上一个组件的距离，0-50
	MarginBottom       float64                                `json:"margin_bottom,omitempty"`           // 与下一个组件的距离，0-50
	ConvertUnknown     *OfficialLandingPageConvertUnknown     `json:"convert_unknown,omitempty"`         // 转化-暂不支持
	ConvertDownload    *OfficialLandingPageConvertDownload    `json:"convert_download,omitempty"`        // 转化-下载 APP
	ConvertLink        *OfficialLandingPageConvertLink        `json:"convert_link,omitempty"`            // 转化-跳转链接
	ConvertWeapp       *OfficialLandingPageConvertWeapp       `json:"convert_weapp,omitempty"`           // 转化-打开小程序
	ConvertGh          *OfficialLandingPageConvertGh          `json:"convert_gh,omitempty"`              // 转化-关注公众号
	ConvertFollowVideo *OfficialLandingPageConvertFollowVideo `json:"convert_followVideo,omitempty"`     // 转化-关注视频号
}

// OfficialLandingPageShelfGroupItem 双图文复合组件单个卡片
type OfficialLandingPageShelfGroupItem struct {
	Title              string                                 `json:"title,omitempty"`                   // 标题文案，1-10字节
	TitleColor         string                                 `json:"title_color,omitempty"`             // 标题颜色，1-7字节，默认 #000000
	Desc               string                                 `json:"desc,omitempty"`                    // 描述文案，1-20字节
	DescColor          string                                 `json:"desc_color,omitempty"`              // 字体颜色，1-7字节，默认 #7F7F7F
	ButtonContent      string                                 `json:"button_content,omitempty"`          // 按钮文案，1-5字节
	ButtonFontColor    string                                 `json:"button_font_color,omitempty"`       // 按钮文案颜色，1-7字节，默认 #FFFFFF
	ButtonBgColor      string                                 `json:"button_background_color,omitempty"` // 按钮背景颜色，1-7字节，默认 #296BEF
	CardBgColor        string                                 `json:"card_background_color,omitempty"`   // 卡片背景颜色，1-7字节，默认 #FFFFFF
	CardBorderColor    string                                 `json:"card_border_color,omitempty"`       // 卡片边框颜色，1-7字节，默认 #E5E5E5
	IconMaterialId     string                                 `json:"icon_material_id,omitempty"`        // 图片素材 id，1-64字节
	ConvertUnknown     *OfficialLandingPageConvertUnknown     `json:"convert_unknown,omitempty"`         // 转化-暂不支持
	ConvertDownload    *OfficialLandingPageConvertDownload    `json:"convert_download,omitempty"`        // 转化-下载 APP
	ConvertLink        *OfficialLandingPageConvertLink        `json:"convert_link,omitempty"`            // 转化-跳转链接
	ConvertWeapp       *OfficialLandingPageConvertWeapp       `json:"convert_weapp,omitempty"`           // 转化-打开小程序
	ConvertGh          *OfficialLandingPageConvertGh          `json:"convert_gh,omitempty"`              // 转化-关注公众号
	ConvertFollowVideo *OfficialLandingPageConvertFollowVideo `json:"convert_followVideo,omitempty"`     // 转化-关注视频号
}

// OfficialLandingPageBlockShelfGroupConfig 双图文复合组件配置
type OfficialLandingPageBlockShelfGroupConfig struct {
	TextAlign    string                               `json:"text_align,omitempty"`    // 对齐方式：left/right/center，默认 left
	Content      []*OfficialLandingPageShelfGroupItem `json:"content,omitempty"`       // 卡片列表，数组长度为 2
	MarginTop    float64                              `json:"margin_top,omitempty"`    // 与上一个组件的距离，0-50
	MarginBottom float64                              `json:"margin_bottom,omitempty"` // 与下一个组件的距离，0-50
}

// ========== 组件列表项 ==========

// OfficialLandingPageElement 官方落地页组件列表项
type OfficialLandingPageElement struct {
	ElementType                string                                         `json:"element_type"`                             // 组件类型 (必填)
	HeadOutsideMaterialConfig  interface{}                                    `json:"head_outside_material_config,omitempty"`   // 顶部外显组件配置
	HeadVideoConfig            *OfficialLandingPageHeadVideoConfig            `json:"head_video_config,omitempty"`              // 顶部视频组件配置
	HeadImageConfig            *OfficialLandingPageHeadImageConfig            `json:"head_image_config,omitempty"`              // 顶部图片组件配置
	Head180PanoramaImageConfig *OfficialLandingPageHead180PanoramaImageConfig `json:"head_180_panorama_image_config,omitempty"` // 顶部 180 度全景图组件配置
	HeadCarouselConfig         *OfficialLandingPageHeadCarouselConfig         `json:"head_carousel_config,omitempty"`           // 顶部轮播图组件配置
	HeadTextConfig             *OfficialLandingPageHeadTextConfig             `json:"head_text_config,omitempty"`               // 顶部文本组件配置
	BlockVideoConfig           *OfficialLandingPageBlockVideoConfig           `json:"block_video_config,omitempty"`             // 普通视频组件配置
	BlockImageConfig           *OfficialLandingPageBlockImageConfig           `json:"block_image_config,omitempty"`             // 图片组件配置
	BlockCarouselConfig        *OfficialLandingPageBlockCarouselConfig        `json:"block_carousel_config,omitempty"`          // 轮播图组件配置
	BlockTextConfig            *OfficialLandingPageBlockTextConfig            `json:"block_text_config,omitempty"`              // 简单文本组件配置
	BlockDividerConfig         *OfficialLandingPageBlockDividerConfig         `json:"block_divider_config,omitempty"`           // 分割线组件配置
	BlockCustomPageConfig      *OfficialLandingPageBlockCustomPageConfig      `json:"block_custom_page_config,omitempty"`       // 内嵌网页组件配置
	BlockButtonConfig          *OfficialLandingPageBlockButtonConfig          `json:"block_button_config,omitempty"`            // 普通按钮组件配置
	FixedButtonConfig          *OfficialLandingPageFixedButtonConfig          `json:"fixed_button_config,omitempty"`            // 底部悬浮按钮组件配置
	RightFixedButtonConfig     *OfficialLandingPageRightFixedButtonConfig     `json:"right_fixed_button_config,omitempty"`      // 侧边悬浮按钮组件配置
	BlockShelfConfig           *OfficialLandingPageBlockShelfConfig           `json:"block_shelf_config,omitempty"`             // 图文复合组件配置
	BlockShelfGroupConfig      *OfficialLandingPageBlockShelfGroupConfig      `json:"block_shelf_group_config,omitempty"`       // 双图文复合组件配置
}

// ========== 落地页配置 ==========

// OfficialLandingPageConfig 落地页配置结构
type OfficialLandingPageConfig struct {
	PageName                string `json:"page_name"`                            // 落地页名称（管理用），不支持 emoji，1-20字节 (必填)
	PageTitle               string `json:"page_title"`                           // 落地页标题（展示用），不支持 emoji，1-20字节 (必填)
	IosAppId                string `json:"ios_app_id,omitempty"`                 // ios 应用 id，1-64字节
	AndroidAppId            string `json:"android_app_id,omitempty"`             // android 应用 id，1-64字节
	EnableAndroidMarket     bool   `json:"enable_android_market,omitempty"`      // 打开安卓 APP 跳转厂商功能，需与 android_app_id 同时使用，默认 false
	ShareTitle              string `json:"share_title,omitempty"`                // 分享标题，不支持 emoji，1-14字节
	ShareDescription        string `json:"share_description,omitempty"`          // 分享描述，不支持 emoji，1-20字节
	ShareThumbUrlMaterialId string `json:"share_thumburl_material_id,omitempty"` // 分享缩略图素材 id，250x250px，1-20字节
	BgColor                 string `json:"bg_color,omitempty"`                   // 颜色，1-7字节，默认 #FFFFFF
}

// ========== 请求/响应结构 ==========

// OfficialLandingPageCompAddReq 官方落地页基于组件创建请求（POST JSON）
// https://developers.e.qq.com/v3.0/docs/api/official_landing_page_component/add
type OfficialLandingPageCompAddReq struct {
	GlobalReq
	AccountId    int64                         `json:"account_id"`              // 推广帐号 id (必填)
	PageConfig   *OfficialLandingPageConfig    `json:"page_config"`             // 落地页配置结构 (必填)
	PageElements []*OfficialLandingPageElement `json:"page_elements"`           // 官方落地页组件列表 (必填)，0-10个
	ProtoVersion int                           `json:"proto_version,omitempty"` // 参数协议版本：0（API专有）/ 1（API+SDK通用），默认 0
}

func (r *OfficialLandingPageCompAddReq) Format() {
	r.GlobalReq.Format()
}

// Validate 验证官方落地页基于组件创建请求参数
func (r *OfficialLandingPageCompAddReq) Validate() error {
	if r.AccountId == 0 {
		return errors.New("account_id为必填")
	}
	if r.PageConfig == nil {
		return errors.New("page_config为必填")
	}
	if len(r.PageConfig.PageName) < MinOfficialLandingPagePageNameBytes || len(r.PageConfig.PageName) > MaxOfficialLandingPagePageNameBytes {
		return errors.New("page_config.page_name长度须在1-20字节之间")
	}
	if len(r.PageConfig.PageTitle) < MinOfficialLandingPagePageTitleBytes || len(r.PageConfig.PageTitle) > MaxOfficialLandingPagePageTitleBytes {
		return errors.New("page_config.page_title长度须在1-20字节之间")
	}
	if r.PageConfig.IosAppId != "" && (len(r.PageConfig.IosAppId) < MinOfficialLandingPageIosAppIDBytes || len(r.PageConfig.IosAppId) > MaxOfficialLandingPageIosAppIDBytes) {
		return errors.New("page_config.ios_app_id长度须在1-64字节之间")
	}
	if r.PageConfig.AndroidAppId != "" && (len(r.PageConfig.AndroidAppId) < MinOfficialLandingPageAndroidAppIDBytes || len(r.PageConfig.AndroidAppId) > MaxOfficialLandingPageAndroidAppIDBytes) {
		return errors.New("page_config.android_app_id长度须在1-64字节之间")
	}
	if r.PageConfig.ShareTitle != "" && (len(r.PageConfig.ShareTitle) < MinOfficialLandingPageShareTitleBytes || len(r.PageConfig.ShareTitle) > MaxOfficialLandingPageShareTitleBytes) {
		return errors.New("page_config.share_title长度须在1-14字节之间")
	}
	if r.PageConfig.ShareDescription != "" && (len(r.PageConfig.ShareDescription) < MinOfficialLandingPageShareDescriptionBytes || len(r.PageConfig.ShareDescription) > MaxOfficialLandingPageShareDescriptionBytes) {
		return errors.New("page_config.share_description长度须在1-20字节之间")
	}
	if r.PageConfig.ShareThumbUrlMaterialId != "" && (len(r.PageConfig.ShareThumbUrlMaterialId) < MinOfficialLandingPageShareThumbMaterialIDBytes || len(r.PageConfig.ShareThumbUrlMaterialId) > MaxOfficialLandingPageShareThumbMaterialIDBytes) {
		return errors.New("page_config.share_thumburl_material_id长度须在1-20字节之间")
	}
	if r.PageConfig.BgColor != "" && (len(r.PageConfig.BgColor) < MinOfficialLandingPageBgColorBytes || len(r.PageConfig.BgColor) > MaxOfficialLandingPageBgColorBytes) {
		return errors.New("page_config.bg_color长度须在1-7字节之间")
	}
	if r.PageElements == nil {
		return errors.New("page_elements为必填")
	}
	if len(r.PageElements) > MaxOfficialLandingPageElementsCount {
		return errors.New("page_elements数组长度不能超过10")
	}
	for i, elem := range r.PageElements {
		if elem == nil {
			return errors.New("page_elements[" + itoa(i) + "]不能为空")
		}
		if elem.ElementType == "" {
			return errors.New("page_elements[" + itoa(i) + "].element_type为必填")
		}
	}
	if r.ProtoVersion < 0 || r.ProtoVersion > 1 {
		return errors.New("proto_version须为0或1")
	}
	return r.GlobalReq.Validate()
}

// OfficialLandingPageCompAddResp 官方落地页基于组件创建响应
type OfficialLandingPageCompAddResp struct {
	PageId        int64 `json:"page_id"`         // 落地页服务 id，用于广告投放端搭建广告创意选择落地页时使用
	LandingPageId int   `json:"landing_page_id"` // 官方落地页 id，仅用于官方落地页模块时使用
}
