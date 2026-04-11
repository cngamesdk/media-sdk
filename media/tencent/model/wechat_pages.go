package model

import "errors"

// ========== 获取微信落地页列表 ==========
// https://developers.e.qq.com/v3.0/docs/api/wechat_pages/get

// 微信落地页类型枚举（其余类型见 creative.go）
const (
	PageTypeWechatCanvasMiniProgram = "PAGE_TYPE_WECHAT_CANVAS_MINI_PROGRAM" // 微信原生页小程序
)

// 微信原生页类型枚举
const (
	CanvasTypeCommonPage = "CANVAS_TYPE_COMMON_PAGE" // 普通原生页
)

// 微信落地页素材创建来源类型枚举
const (
	WechatPageSourceTypeEqq = "SOURCE_TYPE_EQQ" // 腾讯广告创建
	WechatPageSourceTypeMp  = "SOURCE_TYPE_MP"  // 微信公众号创建
)

// 购买类型枚举
const (
	BuyingTypeAuction    = "BUYINGTYPE_AUCTION"     // 竞价
	BuyingTypeContract   = "BUYINGTYPE_CONTRACT"    // 合约
	BuyingTypeReserved   = "BUYINGTYPE_RESERVED"    // 预定
	BuyingTypeFixedPrice = "BUYINGTYPE_FIXED_PRICE" // 固定价格
)

// 商品模式枚举
const (
	ProductModeSingle   = "SINGLE"   // 单品
	ProductModeMultiple = "MULTIPLE" // 多品
)

// 直播视频模式枚举
const (
	LiveVideoModeLive = "LIVE_VIDEO_MODE_LIVE" // 直播
)

// 直播视频子模式枚举
const (
	LiveVideoSubModeLiveReservation = "LIVE_VIDEO_SUBMODE_LIVE_RESERVATION" // 直播预约
)

// 营销场景枚举
const (
	MarketingSceneDefault                               = "DEFAULT"
	MarketingSceneGameReservation                       = "GAME_RESERVATION"
	MarketingSceneGamePromotion                         = "GAME_PROMOTION"
	MarketingSceneGameClosedBetaTest                    = "GAME_CLOSED_BETA_TEST"
	MarketingSceneEcommerceGoodsDirectPurchaseDaily     = "ECOMMERCE_GOODS_DIRECT_PURCHASE_DAILY"
	MarketingSceneEcommerceGoodsLivePurchaseDaily       = "ECOMMERCE_GOODS_LIVE_PURCHASE_DAILY"
	MarketingSceneEcommerceConsumerCollectCluesDaily    = "ECOMMERCE_CONSUMER_COLLECT_CLUES_DAILY"
	MarketingSceneEcommerceConsumerAddFollowersDaily    = "ECOMMERCE_CONSUMER_ADD_FOLLOWERS_DAILY"
	MarketingSceneEcommerceConsumerOfficialAccountDaily = "ECOMMERCE_CONSUMER_OFFICIAL_ACCOUNTS_DAILY"
	MarketingSceneEcommerceConsumerChannelsDaily        = "ECOMMERCE_CONSUMER_CHANNELS_DAILY"
	MarketingSceneEcommerceConsumerAndroidNewDaily      = "ECOMMERCE_CONSUMER_ANDROID_NEW_DAILY"
	MarketingSceneEcommerceConsumerIosNewDaily          = "ECOMMERCE_CONSUMER_IOS_NEW_DAILY"
	MarketingSceneEcommerceContentBrandDaily            = "ECOMMERCE_CONTENT_BRAND_DAILY"
)

// 过滤字段常量
const (
	WechatPagesGetFilterFieldPageID                  = "page_id"
	WechatPagesGetFilterFieldPageName                = "page_name"
	WechatPagesGetFilterFieldAdcreativeTemplateID    = "adcreative_template_id"
	WechatPagesGetFilterFieldMarketingGoal           = "marketing_goal"
	WechatPagesGetFilterFieldMarketingSubGoal        = "marketing_sub_goal"
	WechatPagesGetFilterFieldMarketingTargetType     = "marketing_target_type"
	WechatPagesGetFilterFieldMarketingCarrierType    = "marketing_carrier_type"
	WechatPagesGetFilterFieldPageType                = "page_type"
	WechatPagesGetFilterFieldMarketingCarrierID      = "marketing_carrier_id"
	WechatPagesGetFilterFieldCanvasType              = "canvas_type"
	WechatPagesGetFilterFieldPageStatus              = "page_status"
	WechatPagesGetFilterFieldSiteSet                 = "site_set"
	WechatPagesGetFilterFieldMarketingScene          = "marketing_scene"
	WechatPagesGetFilterFieldSourceType              = "source_type"
	WechatPagesGetFilterFieldLiveVideoMode           = "live_video_mode"
	WechatPagesGetFilterFieldLiveVideoSubMode        = "live_video_sub_mode"
	WechatPagesGetFilterFieldLiveNoticeID            = "live_notice_id"
	WechatPagesGetFilterFieldProductCatalogID        = "product_catalog_id"
	WechatPagesGetFilterFieldProductSource           = "product_source"
	WechatPagesGetFilterFieldRawAdcreativeTemplateID = "raw_adcreative_template_id"
	WechatPagesGetFilterFieldBuyingType              = "buying_type"
	WechatPagesGetFilterFieldProductMode             = "product_mode"
)

// 过滤操作符常量
const (
	WechatPagesGetFilterOperatorEquals   = "EQUALS"
	WechatPagesGetFilterOperatorIn       = "IN"
	WechatPagesGetFilterOperatorContains = "CONTAINS"
)

// 分页常量
const (
	MinWechatPagesGetPage         = 1     // page 最小值
	MaxWechatPagesGetPage         = 99999 // page 最大值
	MinWechatPagesGetPageSize     = 1     // page_size 最小值
	MaxWechatPagesGetPageSize     = 100   // page_size 最大值
	DefaultWechatPagesGetPage     = 1     // page 默认值
	DefaultWechatPagesGetPageSize = 10    // page_size 默认值

	MaxWechatPagesGetFilteringCount     = 10   // filtering 最大长度
	MaxWechatPagesGetPageNameBytes      = 120  // page_name 值最大字节数
	MaxWechatPagesGetLiveNoticeIDBytes  = 1024 // live_notice_id 值最大字节数
	MaxWechatPagesGetProductSourceBytes = 128  // product_source 值最大字节数
)

// WechatPagesGetFilteringItem 微信落地页过滤条件
type WechatPagesGetFilteringItem struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)
	Operator string   `json:"operator"` // 操作符 (必填)
	Values   []string `json:"values"`   // 字段取值 (必填)
}

// Validate 验证单个过滤条件
func (f *WechatPagesGetFilteringItem) Validate() error {
	if f.Field == "" {
		return errors.New("field为必填")
	}
	if f.Operator == "" {
		return errors.New("operator为必填")
	}
	if len(f.Values) == 0 {
		return errors.New("values为必填，至少包含1个值")
	}
	if f.Field == WechatPagesGetFilterFieldPageName {
		if len(f.Values[0]) == 0 || len(f.Values[0]) > MaxWechatPagesGetPageNameBytes {
			return errors.New("page_name过滤值长度须在1-120字节之间")
		}
	}
	if f.Field == WechatPagesGetFilterFieldLiveNoticeID {
		if len(f.Values[0]) == 0 || len(f.Values[0]) > MaxWechatPagesGetLiveNoticeIDBytes {
			return errors.New("live_notice_id过滤值长度须在1-1024字节之间")
		}
	}
	if f.Field == WechatPagesGetFilterFieldProductSource {
		if len(f.Values[0]) == 0 || len(f.Values[0]) > MaxWechatPagesGetProductSourceBytes {
			return errors.New("product_source过滤值长度须在1-128字节之间")
		}
	}
	return nil
}

// WechatPagesGetReq 获取微信落地页列表请求（GET）
// https://developers.e.qq.com/v3.0/docs/api/wechat_pages/get
type WechatPagesGetReq struct {
	GlobalReq
	AccountID int64                          `json:"account_id"`          // 广告主帐号 id (必填)
	OwnerUID  int64                          `json:"owner_uid,omitempty"` // 原生页授权方 uid，默认0（当前广告主的原生页）
	Filtering []*WechatPagesGetFilteringItem `json:"filtering,omitempty"` // 过滤条件，最大10条
	Page      int                            `json:"page,omitempty"`      // 搜索页码，1-99999，默认1
	PageSize  int                            `json:"page_size,omitempty"` // 每页条数，1-100，默认10
}

func (p *WechatPagesGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page == 0 {
		p.Page = DefaultWechatPagesGetPage
	}
	if p.PageSize == 0 {
		p.PageSize = DefaultWechatPagesGetPageSize
	}
}

// Validate 验证获取微信落地页列表请求参数
func (p *WechatPagesGetReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if len(p.Filtering) > MaxWechatPagesGetFilteringCount {
		return errors.New("filtering数组长度不能超过10")
	}
	for i, f := range p.Filtering {
		if f == nil {
			return errors.New("filtering[" + itoa(i) + "]不能为空")
		}
		if err := f.Validate(); err != nil {
			return errors.New("filtering[" + itoa(i) + "]: " + err.Error())
		}
	}
	if p.Page < MinWechatPagesGetPage || p.Page > MaxWechatPagesGetPage {
		return errors.New("page须在1-99999之间")
	}
	if p.PageSize < MinWechatPagesGetPageSize || p.PageSize > MaxWechatPagesGetPageSize {
		return errors.New("page_size须在1-100之间")
	}
	return p.GlobalReq.Validate()
}

// ========== 响应结构体 ==========

// WechatPageDeepLinkAndroidSpec Android deep link 信息
type WechatPageDeepLinkAndroidSpec struct {
	DeepLinkURL  string `json:"deep_link_url"`  // App deep link URL
	AppAndroidID string `json:"app_android_id"` // Android 应用 id
}

// WechatPageDeepLinkIosSpec iOS deep link 信息
type WechatPageDeepLinkIosSpec struct {
	DeepLinkURL string `json:"deep_link_url"` // App deep link URL
	AppIosID    string `json:"app_ios_id"`    // iOS 应用 id
}

// WechatPageMiniProgramSpec 微信小程序信息
type WechatPageMiniProgramSpec struct {
	Title            string   `json:"title,omitempty"`              // 按钮文案
	MiniProgramID    string   `json:"mini_program_id,omitempty"`    // 小程序 id
	MiniProgramPath  string   `json:"mini_program_path,omitempty"`  // 小程序路径
	MiniProgramPaths []string `json:"mini_program_paths,omitempty"` // 小程序落地页路径列表
}

// WechatPageLinkSpec 外链信息
type WechatPageLinkSpec struct {
	Title               string                         `json:"title,omitempty"`                  // 按钮文案
	URL                 string                         `json:"url,omitempty"`                    // 跳转链接
	DeepLinkAndroidSpec *WechatPageDeepLinkAndroidSpec `json:"deep_link_android_spec,omitempty"` // Android deep link 信息
	DeepLinkIosSpec     *WechatPageDeepLinkIosSpec     `json:"deep_link_ios_spec,omitempty"`     // iOS deep link 信息
	MiniProgramSpec     *WechatPageMiniProgramSpec     `json:"mini_program_spec,omitempty"`      // 小程序信息
}

// WechatPageAppIosSpec iOS App 信息元素
type WechatPageAppIosSpec struct {
	DeepLinkURL string `json:"deep_link_url,omitempty"` // App 直达页 URL
	AppIosID    string `json:"app_ios_id,omitempty"`    // iOS 应用 id
}

// WechatPageAppAndroidSpec Android App 信息元素
type WechatPageAppAndroidSpec struct {
	DeepLinkURL                string `json:"deep_link_url,omitempty"`                  // App 直达页 URL
	AppAndroidID               string `json:"app_android_id,omitempty"`                 // Android 应用 id
	AppAndroidChannelPackageID string `json:"app_android_channel_package_id,omitempty"` // Android 渠道包 id
	AppMarketPackage           string `json:"app_market_package,omitempty"`             // 厂商应用商店包名
}

// WechatPageMiniGameProgramSpec 微信小游戏信息
type WechatPageMiniGameProgramSpec struct {
	Title               string `json:"title,omitempty"`                  // 按钮文案
	MiniGameProgramID   string `json:"mini_game_program_id,omitempty"`   // 小游戏 id
	MiniGameProgramPath string `json:"mini_game_program_path,omitempty"` // 小游戏监测参数
}

// WechatPageFengyeSpec 蜂鸟落地页信息
type WechatPageFengyeSpec struct {
	Title    string `json:"title,omitempty"`     // 按钮文案
	FengyeID string `json:"fengye_id,omitempty"` // 蜂鸟落地页 id
}

// WechatPageCardSpec 微信卡券信息
type WechatPageCardSpec struct {
	Title  string `json:"title,omitempty"`   // 按钮文案
	CardID string `json:"card_id,omitempty"` // 微信卡券 id
}

// WechatPageFollowSpec 关注公众号信息
type WechatPageFollowSpec struct {
	Title string `json:"title,omitempty"` // 按钮文案
}

// WechatPageServiceSpec 客服组件信息
type WechatPageServiceSpec struct {
	Title string `json:"title,omitempty"` // 按钮文案
}

// WechatPageWecomSpec 企业微信组件信息
type WechatPageWecomSpec struct {
	Title   string `json:"title,omitempty"`   // 按钮文案，1-10个全角字符
	GroupID int64  `json:"groupid,omitempty"` // 企业微信组件客服分组 id
	SetID   int64  `json:"setid,omitempty"`   // 地理位置集合 id
}

// WechatPageTelSpec 电话组件信息
type WechatPageTelSpec struct {
	Title    string `json:"title,omitempty"`     // 电话组件标题，1-10个全角字符
	PhoneNum string `json:"phone_num,omitempty"` // 电话号码，1-12个全角字符
}

// WechatPageVideoChannelSpec 视频号关注元素
type WechatPageVideoChannelSpec struct {
	Title          string `json:"title,omitempty"`           // 按钮文案
	FinderNickname string `json:"finder_nickname,omitempty"` // 视频号昵称
	FastFollow     int    `json:"fast_follow,omitempty"`     // 一键关注（0:关闭 1:开启）
}

// WechatPageButtonSpec 按钮组件元素
type WechatPageButtonSpec struct {
	Title               string                         `json:"title,omitempty"`                  // 按钮文案
	URL                 string                         `json:"url,omitempty"`                    // 跳转链接
	LinkSpec            *WechatPageLinkSpec            `json:"link_spec,omitempty"`              // 外链信息
	AppIosSpec          *WechatPageAppIosSpec          `json:"app_ios_spec,omitempty"`           // iOS App 信息元素
	AppAndroidSpec      *WechatPageAppAndroidSpec      `json:"app_android_spec,omitempty"`       // Android App 信息元素
	MiniProgramSpec     *WechatPageMiniProgramSpec     `json:"mini_program_spec,omitempty"`      // 微信小程序信息
	MiniGameProgramSpec *WechatPageMiniGameProgramSpec `json:"mini_game_program_spec,omitempty"` // 微信小游戏信息
	FengyeSpec          *WechatPageFengyeSpec          `json:"fengye_spec,omitempty"`            // 蜂鸟落地页信息
	CardSpec            *WechatPageCardSpec            `json:"card_spec,omitempty"`              // 微信卡券信息
	FollowSpec          *WechatPageFollowSpec          `json:"follow_spec,omitempty"`            // 关注公众号信息
	ServiceSpec         *WechatPageServiceSpec         `json:"service_spec,omitempty"`           // 客服组件信息
	WecomSpec           *WechatPageWecomSpec           `json:"wecom_spec,omitempty"`             // 企业微信组件信息
	UseIcon             int                            `json:"use_icon,omitempty"`               // 是否启用图标（0:关闭 1:开启）
	TelSpec             *WechatPageTelSpec             `json:"tel_spec,omitempty"`               // 电话组件信息
	VideoChannelSpec    *WechatPageVideoChannelSpec    `json:"video_channel_spec,omitempty"`     // 视频号关注元素
}

// WechatPageImageSpec 图片组件元素
type WechatPageImageSpec struct {
	ImageIDList []string `json:"image_id_list,omitempty"` // 图片 id 列表
}

// WechatPageVideoSpec 视频组件元素
type WechatPageVideoSpec struct {
	VideoID int64 `json:"video_id,omitempty"` // 视频 id
}

// WechatPageTextSpec 文字组件元素
type WechatPageTextSpec struct {
	Text string `json:"text,omitempty"` // 文字内容
}

// WechatPageFormSpec 表单组件
type WechatPageFormSpec struct {
	Title string `json:"title,omitempty"` // 按钮文案
}

// WechatPageFloatButtonSpec 浮层组件按钮
type WechatPageFloatButtonSpec struct {
	Title string `json:"title,omitempty"` // 按钮文案
}

// WechatPageElementFloat 浮层组件
type WechatPageElementFloat struct {
	ImageIDList     string                     `json:"image_id_list,omitempty"`     // 图片 id
	Title           string                     `json:"title,omitempty"`             // 图文标题
	Desc            string                     `json:"desc,omitempty"`              // 图文描述
	FloatButtonSpec *WechatPageFloatButtonSpec `json:"float_button_spec,omitempty"` // 浮层组件按钮
}

// WechatPageGoodsButtonSpec 商品按钮
type WechatPageGoodsButtonSpec struct {
	Title string `json:"title,omitempty"` // 按钮文案
}

// WechatPageElementGoods 商品组件
type WechatPageElementGoods struct {
	GoodsButtonSpec *WechatPageGoodsButtonSpec `json:"goods_button_spec,omitempty"` // 商品按钮
}

// WechatPageElementSwipe 滑动组件
type WechatPageElementSwipe struct {
	JumpURL   string `json:"jump_url,omitempty"`   // 滑动组件跳转链接
	SwipeText string `json:"swipe_text,omitempty"` // 滑动组件文案，1-10个全角字符
}

// WechatPageElementWebview Webview 组件
type WechatPageElementWebview struct {
	URL string `json:"url,omitempty"` // Webview 链接
}

// WechatPageAnimateFloatButtonSpec 动效浮层按钮
type WechatPageAnimateFloatButtonSpec struct {
	Title string `json:"title,omitempty"` // 按钮文案
}

// WechatPageElementAnimateFloat 动效浮层组件（灰度中）
type WechatPageElementAnimateFloat struct {
	ImageIDList            string                            `json:"image_id_list,omitempty"`             // 图片 id
	Title                  string                            `json:"title,omitempty"`                     // 图文标题
	Desc                   string                            `json:"desc,omitempty"`                      // 图文描述
	AnimateFloatButtonSpec *WechatPageAnimateFloatButtonSpec `json:"animate_float_button_spec,omitempty"` // 动效浮层按钮
}

// WechatPageElementSpec 原生页组件素材内容
type WechatPageElementSpec struct {
	ElementType         string                         `json:"element_type,omitempty"`          // 原生页中的组件类型
	ImageSpec           *WechatPageImageSpec           `json:"image_spec,omitempty"`            // 图片组件元素
	VideoSpec           *WechatPageVideoSpec           `json:"video_spec,omitempty"`            // 视频组件元素
	TextSpec            *WechatPageTextSpec            `json:"text_spec,omitempty"`             // 文字组件元素
	ButtonSpec          *WechatPageButtonSpec          `json:"button_spec,omitempty"`           // 按钮组件元素
	FormSpec            *WechatPageFormSpec            `json:"form_spec,omitempty"`             // 表单组件
	ElementShelf        *WechatPageElementShelf        `json:"element_shelf,omitempty"`         // 图文复合组件
	ElementFloat        *WechatPageElementFloat        `json:"element_float,omitempty"`         // 浮层组件
	ElementGoods        *WechatPageElementGoods        `json:"element_goods,omitempty"`         // 商品组件
	ElementSwipe        *WechatPageElementSwipe        `json:"element_swipe,omitempty"`         // 滑动组件元素
	ElementWebview      *WechatPageElementWebview      `json:"element_webview,omitempty"`       // Webview 组件元素
	ElementAnimateFloat *WechatPageElementAnimateFloat `json:"element_animate_float,omitempty"` // 动效浮层组件（灰度中）
}

// WechatPageShareContentSpec 微信原生页分享信息
type WechatPageShareContentSpec struct {
	ShareTitle       string `json:"share_title,omitempty"`       // 分享标题
	ShareDescription string `json:"share_description,omitempty"` // 分享描述
}

// WechatPagesGetItem 微信落地页列表项
type WechatPagesGetItem struct {
	PageID               int64                       `json:"page_id"`                           // 落地页 id
	PageName             string                      `json:"page_name"`                         // 落地页名称
	CreatedTime          int64                       `json:"created_time"`                      // 创建时间，时间戳
	LastModifiedTime     int64                       `json:"last_modified_time"`                // 最后修改时间，时间戳
	PageTemplateID       int64                       `json:"page_template_id"`                  // 落地页模板 id
	PageElementsSpecList []*WechatPageElementSpec    `json:"page_elements_spec_list,omitempty"` // 组件素材内容列表
	ShareContentSpec     *WechatPageShareContentSpec `json:"share_content_spec,omitempty"`      // 微信原生页分享信息
	PreviewURL           string                      `json:"preview_url"`                       // 微信原生页预览 URL（1小时有效）
	PageType             string                      `json:"page_type"`                         // 落地页类型
	SourceType           string                      `json:"source_type"`                       // 素材创建来源类型
	VideoResourceStatus  string                      `json:"video_resource_status"`             // 微信原生页视频资源状态
	CanvasType           string                      `json:"canvas_type"`                       // 原生页类型
	OwnerUID             int64                       `json:"owner_uid"`                         // 原生页授权方 uid
	PageStatus           string                      `json:"page_status"`                       // 状态
}

// WechatPagesGetResp 获取微信落地页列表响应
// https://developers.e.qq.com/v3.0/docs/api/wechat_pages/get
type WechatPagesGetResp struct {
	List     []*WechatPagesGetItem `json:"list"`      // 返回信息列表
	PageInfo *PageInfo             `json:"page_info"` // 分页配置信息
}

// ========== 基于模板创建微信原生页 ==========
// https://developers.e.qq.com/v3.0/docs/api/wechat_pages/add

// 创建接口字段长度常量
const (
	MinWechatPagesAddPageNameBytes         = 1     // page_name 最小长度
	MaxWechatPagesAddPageNameBytes         = 120   // page_name 最大长度
	MaxWechatPagesAddElementsCount         = 40    // page_elements_spec_list 最大长度
	MaxWechatPagesAddImageIDListCount      = 6     // image_id_list 最大数量
	MaxWechatPagesAddImageIDBytes          = 64    // 单个 image_id 最大长度
	MaxWechatPagesAddTextBytes             = 30720 // text_spec.text 最大长度
	MaxWechatPagesAddButtonTitleBytes      = 120   // button title 最大长度
	MaxWechatPagesAddButtonURLBytes        = 1023  // button url 最大长度
	MaxWechatPagesAddDeepLinkURLBytes      = 2048  // deep_link_url 最大长度
	MaxWechatPagesAddAppIDBytes            = 128   // app_android_id / app_ios_id 最大长度
	MaxWechatPagesAddChannelPackageIDBytes = 128   // app_android_channel_package_id 最大长度
	MaxWechatPagesAddAppMarketPackageBytes = 512   // app_market_package 最大长度
	MaxWechatPagesAddMiniProgramIDBytes    = 384   // mini_program_id 最大长度
	MaxWechatPagesAddMiniProgramPathBytes  = 2048  // mini_program_path 最大长度
	MaxWechatPagesAddMiniProgramPathsCount = 255   // mini_program_paths 最大数量
	MaxWechatPagesAddMiniGameIDBytes       = 384   // mini_game_program_id 最大长度
	MaxWechatPagesAddMiniGamePathBytes     = 750   // mini_game_program_path 最大长度
	MaxWechatPagesAddFengyeIDBytes         = 384   // fengye_id 最大长度
	MaxWechatPagesAddCardIDBytes           = 384   // card_id 最大长度
	MaxWechatPagesAddWecomTitleBytes       = 64    // wecom_spec.title 最大长度
	MaxWechatPagesAddWecomGroupID          = int64(9999999999)
	MaxWechatPagesAddTelTitleBytes         = 64   // tel_spec.title 最大长度
	MaxWechatPagesAddPhoneNumBytes         = 64   // tel_spec.phone_num 最大长度
	MaxWechatPagesAddFinderNicknameBytes   = 120  // finder_nickname 最大长度
	MaxWechatPagesAddFormTitleBytes        = 30   // form_spec.title 最大长度
	MaxWechatPagesAddShelfSpecCount        = 32   // shelf_spec 最大长度
	MaxWechatPagesAddSwipeURLBytes         = 1023 // element_swipe.jump_url 最大长度
	MaxWechatPagesAddSwipeTextBytes        = 30   // element_swipe.swipe_text 最大长度
	MaxWechatPagesAddWebviewURLBytes       = 1023 // element_webview.url 最大长度
)

// WechatPageAppDownloadSpec 应用下载信息（用于图文复合组件按钮）
type WechatPageAppDownloadSpec struct {
	Title           string                     `json:"title"`                       // 按钮文案 (必填)，须含"安装"或"下载"，1-120字节
	AppIosSpec      *WechatPageAppIosSpec      `json:"app_ios_spec,omitempty"`      // iOS App 信息元素
	AppAndroidSpec  *WechatPageAppAndroidSpec  `json:"app_android_spec,omitempty"`  // Android App 信息元素
	MiniProgramSpec *WechatPageMiniProgramSpec `json:"mini_program_spec,omitempty"` // 小程序信息
	WecomSpec       *WechatPageWecomSpec       `json:"wecom_spec,omitempty"`        // 企业微信组件信息
}

// WechatPageShelfButtonSpec 图文复合组件按钮信息
type WechatPageShelfButtonSpec struct {
	LinkSpec        *WechatPageLinkSpec        `json:"link_spec,omitempty"`         // 外链信息
	AppDownloadSpec *WechatPageAppDownloadSpec `json:"app_download_spec,omitempty"` // 应用下载信息
	ImageIDList     string                     `json:"image_id_list,omitempty"`     // 图片 id，1-64字节
	Title           string                     `json:"title,omitempty"`             // 图文复合标题文案，1-120字节
	Desc            string                     `json:"desc,omitempty"`              // 图文复合描述文案，1-120字节
}

// WechatPageShelfSpec 图文复合组件行
type WechatPageShelfSpec struct {
	ShelfButtonSpec *WechatPageShelfButtonSpec `json:"shelf_button_spec,omitempty"` // 图文复合按钮信息 (必填)
}

// WechatPageElementShelf 图文复合组件
type WechatPageElementShelf struct {
	ShelfSpec []*WechatPageShelfSpec `json:"shelf_spec"` // 按钮信息列表 (必填)，最大32条
}

// WechatPagesAddReq 基于模板创建微信原生页请求（POST JSON）
// https://developers.e.qq.com/v3.0/docs/api/wechat_pages/add
type WechatPagesAddReq struct {
	GlobalReq
	AccountID            int64                       `json:"account_id"`              // 广告主帐号 id (必填)
	PageName             string                      `json:"page_name"`               // 落地页名称 (必填)，1-120字节
	PageTemplateID       int64                       `json:"page_template_id"`        // 落地页模板 id (必填)
	PageElementsSpecList []*WechatPageElementSpec    `json:"page_elements_spec_list"` // 组件素材内容列表 (必填)，最大40条
	ShareContentSpec     *WechatPageShareContentSpec `json:"share_content_spec"`      // 微信原生页分享信息 (必填)
}

func (p *WechatPagesAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证基于模板创建微信原生页请求参数
func (p *WechatPagesAddReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if len(p.PageName) < MinWechatPagesAddPageNameBytes || len(p.PageName) > MaxWechatPagesAddPageNameBytes {
		return errors.New("page_name长度须在1-120字节之间")
	}
	if p.PageTemplateID == 0 {
		return errors.New("page_template_id为必填")
	}
	if len(p.PageElementsSpecList) == 0 {
		return errors.New("page_elements_spec_list为必填，至少包含1个组件")
	}
	if len(p.PageElementsSpecList) > MaxWechatPagesAddElementsCount {
		return errors.New("page_elements_spec_list数组长度不能超过40")
	}
	for i, el := range p.PageElementsSpecList {
		if el == nil {
			return errors.New("page_elements_spec_list[" + itoa(i) + "]不能为空")
		}
		if el.ElementType == "" {
			return errors.New("page_elements_spec_list[" + itoa(i) + "].element_type为必填")
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

// WechatPagesAddResp 基于模板创建微信原生页响应
// https://developers.e.qq.com/v3.0/docs/api/wechat_pages/add
type WechatPagesAddResp struct {
	PageID int64 `json:"page_id"` // 落地页 id
}

// ========== 删除微信落地页 ==========
// https://developers.e.qq.com/v3.0/docs/api/wechat_pages/delete

// WechatPagesDeleteReq 删除微信落地页请求（POST JSON）
// https://developers.e.qq.com/v3.0/docs/api/wechat_pages/delete
type WechatPagesDeleteReq struct {
	GlobalReq
	AccountID int64 `json:"account_id"` // 广告主帐号 id (必填)
	PageID    int64 `json:"page_id"`    // 落地页 id (必填)
}

func (p *WechatPagesDeleteReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证删除微信落地页请求参数
func (p *WechatPagesDeleteReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.PageID == 0 {
		return errors.New("page_id为必填")
	}
	return p.GlobalReq.Validate()
}

// WechatPagesDeleteResp 删除微信落地页响应
// https://developers.e.qq.com/v3.0/docs/api/wechat_pages/delete
type WechatPagesDeleteResp struct {
	PageID int64 `json:"page_id"` // 落地页 id
}
