package model

import "errors"

// DynamicCreativesGetReq 获取创意请求
// https://developers.e.qq.com/v3.0/docs/api/dynamic_creatives/get
type DynamicCreativesGetReq struct {
	GlobalReq
	AccountID      int64                         `json:"account_id"`                // 广告主帐号id (必填)
	Filtering      []*DynamicCreativeQueryFilter `json:"filtering,omitempty"`       // 过滤条件，数组长度1-10
	Page           int                           `json:"page,omitempty"`            // 搜索页码，默认1
	PageSize       int                           `json:"page_size,omitempty"`       // 每页条数，默认10
	Fields         []string                      `json:"fields,omitempty"`          // 指定返回的字段列表
	IsDeleted      bool                          `json:"is_deleted,omitempty"`      // 是否已删除
	PaginationMode string                        `json:"pagination_mode,omitempty"` // 分页方式
	Cursor         string                        `json:"cursor,omitempty"`          // 游标值（游标分页模式）
}

// DynamicCreativeQueryFilter 创意查询过滤条件
type DynamicCreativeQueryFilter struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)
	Operator string   `json:"operator"` // 操作符 (必填)
	Values   []string `json:"values"`   // 字段取值 (必填)
}

// 常量定义 - 创意过滤字段
const (
	CreativeFieldDynamicCreativeID       = "dynamic_creative_id"
	CreativeFieldDynamicCreativeName     = "dynamic_creative_name"
	CreativeFieldCreatedTime             = "created_time"
	CreativeFieldLastModifiedTime        = "last_modified_time"
	CreativeFieldAdgroupID               = "adgroup_id"
	CreativeFieldConfiguredStatus        = "configured_status"
	CreativeFieldAdgroupCampaignType     = "adgroup.campaign_type"
	CreativeFieldSource                  = "source"
	CreativeFieldComponentID             = "component_id"
	CreativeFieldSmartDeliveryTemplateDC = "smart_delivery_template_dc_id"
	CreativeFieldDataModelVersion        = "data_model_version"
)

// 常量定义 - 创意配置状态
const (
	CreativeConfiguredStatusNormal  = "AD_STATUS_NORMAL"  // 正常
	CreativeConfiguredStatusSuspend = "AD_STATUS_SUSPEND" // 暂停
)

// 常量定义 - 分页方式
const (
	CreativePaginationModeNormal = "PAGINATION_MODE_NORMAL" // 普通分页
	CreativePaginationModeCursor = "PAGINATION_MODE_CURSOR" // 游标分页
)

// 字段限制常量
const (
	MaxCreativeFilteringCount = 10
	MaxCreativeNameLength     = 180
	MaxCreativeCursorLength   = 10
	DefaultCreativePage       = 1
	DefaultCreativePageSize   = 10
)

func (p *DynamicCreativesGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page <= 0 {
		p.Page = DefaultCreativePage
	}
	if p.PageSize <= 0 {
		p.PageSize = DefaultCreativePageSize
	}
	if p.PaginationMode == "" {
		p.PaginationMode = CreativePaginationModeNormal
	}
}

// Validate 验证获取创意请求参数
func (p *DynamicCreativesGetReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	if p.PaginationMode != CreativePaginationModeNormal && p.PaginationMode != CreativePaginationModeCursor {
		return errors.New("pagination_mode值无效，允许值：PAGINATION_MODE_NORMAL、PAGINATION_MODE_CURSOR")
	}

	if p.Page < MinPage || p.Page > MaxPage {
		return errors.New("page必须在1-100之间")
	}

	if p.PageSize < MinPageSize || p.PageSize > MaxPageSize {
		return errors.New("page_size必须在1-100之间")
	}

	if p.Cursor != "" && len(p.Cursor) > MaxCreativeCursorLength {
		return errors.New("cursor长度不能超过10字节")
	}

	return validateDynamicCreativeFiltering(p.Filtering)
}

// validateDynamicCreativeFiltering 验证过滤条件
func validateDynamicCreativeFiltering(filtering []*DynamicCreativeQueryFilter) error {
	if len(filtering) == 0 {
		return nil
	}
	if len(filtering) > MaxCreativeFilteringCount {
		return errors.New("filtering数组长度不能超过10")
	}
	for _, filter := range filtering {
		if err := filter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Validate 验证单个创意过滤条件
func (f *DynamicCreativeQueryFilter) Validate() error {
	if f.Field == "" {
		return errors.New("field为必填")
	}
	if !isValidCreativeField(f.Field) {
		return errors.New("field值无效，请参考文档中的允许值")
	}
	if f.Operator == "" {
		return errors.New("operator为必填")
	}
	if !isValidCreativeOperatorForField(f.Field, f.Operator) {
		return errors.New("operator值无效，当前字段不支持该操作符")
	}
	if len(f.Values) == 0 {
		return errors.New("values为必填")
	}
	return validateCreativeValuesForField(f)
}

// isValidCreativeField 验证创意过滤字段是否有效
func isValidCreativeField(field string) bool {
	validFields := map[string]bool{
		CreativeFieldDynamicCreativeID:       true,
		CreativeFieldDynamicCreativeName:     true,
		CreativeFieldCreatedTime:             true,
		CreativeFieldLastModifiedTime:        true,
		CreativeFieldAdgroupID:               true,
		CreativeFieldConfiguredStatus:        true,
		CreativeFieldAdgroupCampaignType:     true,
		CreativeFieldSource:                  true,
		CreativeFieldComponentID:             true,
		CreativeFieldSmartDeliveryTemplateDC: true,
		CreativeFieldDataModelVersion:        true,
	}
	return validFields[field]
}

// isValidCreativeOperatorForField 验证创意字段支持的操作符
func isValidCreativeOperatorForField(field, operator string) bool {
	switch field {
	case CreativeFieldDynamicCreativeID, CreativeFieldAdgroupID:
		return operator == OperatorEquals || operator == OperatorIn
	case CreativeFieldDynamicCreativeName, CreativeFieldAdgroupCampaignType:
		return operator == OperatorEquals
	case CreativeFieldCreatedTime, CreativeFieldLastModifiedTime:
		return operator == OperatorEquals || operator == OperatorLess ||
			operator == OperatorLessEquals || operator == OperatorGreater ||
			operator == OperatorGreaterEquals
	case CreativeFieldConfiguredStatus:
		return operator == OperatorEquals
	case CreativeFieldSource:
		return operator == OperatorEquals || operator == OperatorIn
	case CreativeFieldComponentID:
		return operator == OperatorEquals
	case CreativeFieldSmartDeliveryTemplateDC, CreativeFieldDataModelVersion:
		return operator == OperatorEquals
	default:
		return false
	}
}

// validateCreativeValuesForField 验证创意字段取值
func validateCreativeValuesForField(f *DynamicCreativeQueryFilter) error {
	switch f.Field {
	case CreativeFieldDynamicCreativeID, CreativeFieldAdgroupID:
		if f.Operator == OperatorEquals {
			if len(f.Values) != 1 {
				return errors.New("operator为EQUALS时，values数组长度必须为1")
			}
		} else if f.Operator == OperatorIn {
			if len(f.Values) < MinValuesCount || len(f.Values) > MaxValuesCount {
				return errors.New("operator为IN时，values数组长度必须在1-100之间")
			}
		}
	case CreativeFieldDynamicCreativeName, CreativeFieldAdgroupCampaignType:
		if len(f.Values) != 1 {
			return errors.New("values数组长度必须为1")
		}
		if len(f.Values[0]) < 1 || len(f.Values[0]) > MaxCreativeNameLength {
			return errors.New("字段长度必须在1-180字节之间")
		}
	case CreativeFieldConfiguredStatus:
		if len(f.Values) != 1 {
			return errors.New("values数组长度必须为1")
		}
		if f.Values[0] != CreativeConfiguredStatusNormal && f.Values[0] != CreativeConfiguredStatusSuspend {
			return errors.New("configured_status值无效，允许值：AD_STATUS_NORMAL、AD_STATUS_SUSPEND")
		}
	case CreativeFieldCreatedTime, CreativeFieldLastModifiedTime:
		if len(f.Values) != 1 {
			return errors.New("values数组长度必须为1")
		}
		if len(f.Values[0]) != CreatedTimeLength {
			return errors.New("时间字段长度必须为10字节")
		}
	case CreativeFieldSource:
		if f.Operator == OperatorEquals {
			if len(f.Values) != 1 {
				return errors.New("operator为EQUALS时，values数组长度必须为1")
			}
		} else if f.Operator == OperatorIn {
			if len(f.Values) < MinValuesCount || len(f.Values) > MaxValuesCount {
				return errors.New("operator为IN时，values数组长度必须在1-100之间")
			}
		}
	case CreativeFieldComponentID:
		if len(f.Values) != 1 {
			return errors.New("values数组长度必须为1")
		}
	}
	return nil
}

// DynamicCreativesGetResp 获取创意响应
type DynamicCreativesGetResp struct {
	List []*DynamicCreativeListItem `json:"list,omitempty"` // 创意列表
	CursorPageInfoV2Container
	PageInfoContainer
}

// DynamicCreativeListItem 创意列表项
type DynamicCreativeListItem struct {
	AdgroupID                  int64                       `json:"adgroup_id"`                             // 广告id
	DynamicCreativeID          int64                       `json:"dynamic_creative_id"`                    // 广告创意id
	DynamicCreativeName        string                      `json:"dynamic_creative_name,omitempty"`        // 广告创意名称
	CreativeTemplateID         int64                       `json:"creative_template_id,omitempty"`         // 创意形式id
	DeliveryMode               string                      `json:"delivery_mode,omitempty"`                // 投放模式
	DynamicCreativeType        string                      `json:"dynamic_creative_type,omitempty"`        // 动态创意类型
	CreativeComponents         *CreativeComponents         `json:"creative_components,omitempty"`          // 创意组件
	ImpressionTrackingURL      string                      `json:"impression_tracking_url,omitempty"`      // 曝光监控地址
	ClickTrackingURL           string                      `json:"click_tracking_url,omitempty"`           // 点击监控链接
	PageTrackURL               string                      `json:"page_track_url,omitempty"`               // 页面级转化跟踪URL
	ProgramCreativeInfo        *ProgramCreativeInfo        `json:"program_creative_info,omitempty"`        // 程序化创意信息
	ConfiguredStatus           string                      `json:"configured_status,omitempty"`            // 客户设置的状态
	CreativeSetApprovalStatus  string                      `json:"creative_set_approval_status,omitempty"` // 动态创意审核状态
	Source                     string                      `json:"source,omitempty"`                       // 创意来源
	AssetInconsistentStatus    string                      `json:"asset_inconsistent_status,omitempty"`    // 推广内容资产与落地页一致性状态
	SourceDynamicCreativeID    int64                       `json:"source_dynamic_creative_id,omitempty"`   // 源广告创意id
	MarketingAssetVerification *MarketingAssetVerification `json:"marketing_asset_verification,omitempty"` // 资产验真结果
	CreativeInsight            *CreativeInsight            `json:"creative_insight,omitempty"`             // 创意洞察数据
	CreatedTime                int64                       `json:"created_time,omitempty"`                 // 创建时间，时间戳
	LastModifiedTime           int64                       `json:"last_modified_time,omitempty"`           // 最后修改时间，时间戳
	IsDeleted                  bool                        `json:"is_deleted,omitempty"`                   // 是否已删除
}

// ProgramCreativeInfo 程序化创意信息
type ProgramCreativeInfo struct {
	MaterialDeriveID   int64                 `json:"material_derive_id,omitempty"`   // 衍生id
	MaterialDeriveInfo []*MaterialDeriveInfo `json:"material_derive_info,omitempty"` // 素材和衍生信息列表
	BidMode            string                `json:"bid_mode,omitempty"`             // 出价方式
	DeriveVersion      string                `json:"derive_version,omitempty"`       // 动态创意类型
}

// MaterialDeriveInfo 素材和衍生信息
type MaterialDeriveInfo struct {
	OriginalMaterialIDList           []int64 `json:"original_material_id_list,omitempty"`            // 原始素材id列表
	OriginalAdcreativeTemplateIDList []int64 `json:"original_adcreative_template_id_list,omitempty"` // 原始创意模板id列表
	OriginalCoverImageID             int64   `json:"original_cover_image_id,omitempty"`              // 原始封面图id
}

// MarketingAssetVerification 资产验真结果
type MarketingAssetVerification struct {
	VerificationStatus string `json:"verification_status,omitempty"` // 验真状态
	VerificationMsg    string `json:"verification_msg,omitempty"`    // 验真说明
}

// CreativeInsight 创意洞察数据
type CreativeInsight struct {
	InsightType  string `json:"insight_type,omitempty"`  // 洞察类型
	InsightValue string `json:"insight_value,omitempty"` // 洞察值
}

// CreativeComponents 创意组件集合
type CreativeComponents struct {
	Title                []*CreativeComponent `json:"title,omitempty"`                  // 标题组件
	Description          []*CreativeComponent `json:"description,omitempty"`            // 文本描述组件
	Image                []*CreativeComponent `json:"image,omitempty"`                  // 单图组件
	ImageList            []*CreativeComponent `json:"image_list,omitempty"`             // 图集组件
	Video                []*CreativeComponent `json:"video,omitempty"`                  // 视频组件
	Brand                []*CreativeComponent `json:"brand,omitempty"`                  // 品牌形象组件
	Consult              []*CreativeComponent `json:"consult,omitempty"`                // 咨询组件
	MainJumpInfo         []*CreativeComponent `json:"main_jump_info,omitempty"`         // 主跳转组件
	Phone                []*CreativeComponent `json:"phone,omitempty"`                  // 电话组件
	Form                 []*CreativeComponent `json:"form,omitempty"`                   // 表单组件
	ActionButton         []*CreativeComponent `json:"action_button,omitempty"`          // 行动按钮组件
	ChosenButton         []*CreativeComponent `json:"chosen_button,omitempty"`          // 选择按钮组件
	Label                []*CreativeComponent `json:"label,omitempty"`                  // 标签组件
	ShowData             []*CreativeComponent `json:"show_data,omitempty"`              // 数据外显组件
	MarketingPendant     []*CreativeComponent `json:"marketing_pendant,omitempty"`      // 营销挂件组件
	AppGiftPackCode      []*CreativeComponent `json:"app_gift_pack_code,omitempty"`     // 礼包码组件
	ShopImage            []*CreativeComponent `json:"shop_image,omitempty"`             // 卖点图组件
	CountDown            []*CreativeComponent `json:"count_down,omitempty"`             // 倒计时组件
	Barrage              []*CreativeComponent `json:"barrage,omitempty"`                // 弹幕组件
	FloatingZone         []*CreativeComponent `json:"floating_zone,omitempty"`          // 浮层卡片组件
	FloatingZoneList     []*CreativeComponent `json:"floating_zone_list,omitempty"`     // 多卡轮播组件
	TextLink             []*CreativeComponent `json:"text_link,omitempty"`              // 文字链组件
	EndPage              []*CreativeComponent `json:"end_page,omitempty"`               // 视频结束页
	LivingDesc           []*CreativeComponent `json:"living_desc,omitempty"`            // 轮播文案
	WechatChannels       []*CreativeComponent `json:"wechat_channels,omitempty"`        // 视频号信息组件
	ShortVideo           []*CreativeComponent `json:"short_video,omitempty"`            // 短视频组件
	ElementStory         []*CreativeComponent `json:"element_story,omitempty"`          // 集装箱创意组合组件
	WxgamePlayablePage   []*CreativeComponent `json:"wxgame_playable_page,omitempty"`   // 小游戏试玩页组件
	WxgameDirectPage     []*CreativeComponent `json:"wxgame_direct_page,omitempty"`     // 小游戏直玩组件
	AppPromotionVideo    []*CreativeComponent `json:"app_promotion_video,omitempty"`    // OTT视频组件
	VideoShowcase        []*CreativeComponent `json:"video_showcase,omitempty"`         // 橱窗视频组件
	ImageShowcase        []*CreativeComponent `json:"image_showcase,omitempty"`         // 橱窗图片组件
	SocialSkill          []*CreativeComponent `json:"social_skill,omitempty"`           // 首评回复组件
	MiniCardLink         []*CreativeComponent `json:"mini_card_link,omitempty"`         // 图文链接组件
	VideoChannelsContent []*CreativeComponent `json:"video_channels_content,omitempty"` // 视频号主页视频组件
	Audio                []*CreativeComponent `json:"audio,omitempty"`                  // 音频组件
}

// CreativeComponent 创意组件通用结构
type CreativeComponent struct {
	ComponentID int64       `json:"component_id,omitempty"` // 创意组件id
	Value       interface{} `json:"value,omitempty"`        // 组件值内容
	IsDeleted   bool        `json:"is_deleted,omitempty"`   // 是否已删除
}

// ========== 创建创意 ==========

// DynamicCreativesAddReq 创建创意请求
// https://developers.e.qq.com/v3.0/docs/api/dynamic_creatives/add
type DynamicCreativesAddReq struct {
	GlobalReq
	AccountID                        int64                `json:"account_id"`                                     // 广告主帐号id (必填)
	AdgroupID                        int64                `json:"adgroup_id"`                                     // 广告id (必填)
	DynamicCreativeName              string               `json:"dynamic_creative_name"`                          // 广告创意名称，同账号下不重复，1-60等宽字符 (必填)
	CreativeComponents               *CreativeComponents  `json:"creative_components"`                            // 创意组件 (必填)
	CreativeTemplateID               int64                `json:"creative_template_id,omitempty"`                 // 创意形式id
	DeliveryMode                     string               `json:"delivery_mode,omitempty"`                        // 投放模式
	DynamicCreativeType              string               `json:"dynamic_creative_type,omitempty"`                // 动态创意类型
	ImpressionTrackingURL            string               `json:"impression_tracking_url,omitempty"`              // 曝光监控地址
	ClickTrackingURL                 string               `json:"click_tracking_url,omitempty"`                   // 点击监控链接
	ProgramCreativeInfo              *ProgramCreativeInfo `json:"program_creative_info,omitempty"`                // 程序化创意信息
	PageTrackURL                     string               `json:"page_track_url,omitempty"`                       // 页面级转化跟踪URL
	AutoDerivedProgramCreativeSwitch bool                 `json:"auto_derived_program_creative_switch,omitempty"` // 自动衍生程序化创意开关
	ConfiguredStatus                 string               `json:"configured_status,omitempty"`                    // 客户设置的状态
	SiteSetValidateModel             string               `json:"site_set_validate_model,omitempty"`              // 版位校验模式
}

// 常量定义 - 投放模式
const (
	DeliveryModeComponent = "DELIVERY_MODE_COMPONENT" // 组件化投放
	DeliveryModeCustomize = "DELIVERY_MODE_CUSTOMIZE" // 自定义投放
)

// 常量定义 - 动态创意类型
const (
	DynamicCreativeTypeCommon  = "DYNAMIC_CREATIVE_TYPE_COMMON"  // 普通动态创意
	DynamicCreativeTypeProgram = "DYNAMIC_CREATIVE_TYPE_PROGRAM" // 程序化动态创意
)

// 常量定义 - 落地页类型
const (
	PageTypeAndroidApp              = "PAGE_TYPE_ANDROID_APP"                // Android应用
	PageTypeIosApp                  = "PAGE_TYPE_IOS_APP"                    // iOS应用
	PageTypeXjAndroidAppH5          = "PAGE_TYPE_XJ_ANDROID_APP_H5"          // 蹊径Android H5
	PageTypeXjIosAppH5              = "PAGE_TYPE_XJ_IOS_APP_H5"              // 蹊径iOS H5
	PageTypeXjWebH5                 = "PAGE_TYPE_XJ_WEB_H5"                  // 蹊径网页H5
	PageTypeXjQuick                 = "PAGE_TYPE_XJ_QUICK"                   // 蹊径快应用
	PageTypeFengyeEcommerce         = "PAGE_TYPE_FENGYE_ECOMMERCE"           // 枫叶电商
	PageTypeAppDeepLink             = "PAGE_TYPE_APP_DEEP_LINK"              // 应用直达
	PageTypeAppMarket               = "PAGE_TYPE_APP_MARKET"                 // 厂商下载
	PageTypeAndroidQuickApp         = "PAGE_TYPE_ANDROID_QUICK_APP"          // Android快应用
	PageTypeWechatCanvas            = "PAGE_TYPE_WECHAT_CANVAS"              // 微信原生页
	PageTypeWechatMiniProgram       = "PAGE_TYPE_WECHAT_MINI_PROGRAM"        // 微信小程序
	PageTypeWechatMiniGame          = "PAGE_TYPE_WECHAT_MINI_GAME"           // 微信小游戏
	PageTypeWechatChannelsFeed      = "PAGE_TYPE_WECHAT_CHANNELS_FEED"       // 视频号动态
	PageTypeWechatChannelsWatchLive = "PAGE_TYPE_WECHAT_CHANNELS_WATCH_LIVE" // 视频号观看直播
	PageTypeWechatShop              = "PAGE_TYPE_WECHAT_SHOP"                // 微信小店
	PageTypeH5                      = "PAGE_TYPE_H5"                         // 自定义H5
	PageTypeQqAppMiniProgram        = "PAGE_TYPE_QQ_APP_MINI_PROGRAM"        // QQ小程序
	PageTypeQqMiniGame              = "PAGE_TYPE_QQ_MINI_GAME"               // QQ小游戏
)

// 常量定义 - 落地页媒体平台类型
const (
	PlatformTypeDefault     = "DEFAULT"
	PlatformTypeAll         = "ALL"
	PlatformTypeScreenPc    = "SCREEN_PC"
	PlatformTypeScreenPhone = "SCREEN_PHONE"
	PlatformTypeOsAndroid   = "OS_ANDROID"
	PlatformTypeOsIos       = "OS_IOS"
)

// 常量定义 - 微信原生页顶部素材替换关系
const (
	OptionCanvasOverrideCreative        = "OPTION_CANVAS_OVERRIDE_CREATIVE"
	OptionCreativeOverrideCanvas        = "OPTION_CREATIVE_OVERRIDE_CANVAS"
	OptionKeepDifferent                 = "OPTION_KEEP_DIFFERENT"
	OptionCreativeOverrideCanvasDynamic = "OPTION_CREATIVE_OVERRIDE_CANVAS_DYNAMIC"
)

// 创意名称长度限制
const (
	MaxDynamicCreativeNameBytes = 180
)

func (p *DynamicCreativesAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证创建创意请求参数
func (p *DynamicCreativesAddReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.AdgroupID == 0 {
		return errors.New("adgroup_id为必填")
	}
	if len(p.DynamicCreativeName) == 0 {
		return errors.New("dynamic_creative_name为必填")
	}
	if len(p.DynamicCreativeName) > MaxDynamicCreativeNameBytes {
		return errors.New("dynamic_creative_name长度不能超过180字节")
	}
	if p.CreativeComponents == nil {
		return errors.New("creative_components为必填")
	}
	return p.GlobalReq.Validate()
}

// DynamicCreativesAddResp 创建创意响应
type DynamicCreativesAddResp struct {
	DynamicCreativeID int64 `json:"dynamic_creative_id"` // 创建的动态创意id
}

// ========== 组件值类型 ==========

// TextComponentValue 文本组件值（标题/描述）
type TextComponentValue struct {
	Content string `json:"content"` // 文本内容
}

// ImageComponentValue 单图组件值
type ImageComponentValue struct {
	ImageID  string    `json:"image_id"`            // 素材图片id (必填)
	ImageURL string    `json:"image_url,omitempty"` // 素材图片url
	JumpInfo *JumpInfo `json:"jump_info,omitempty"` // 落地页内容
}

// ImageListComponentValue 图集组件值
type ImageListComponentValue struct {
	JumpInfo *JumpInfo        `json:"jump_info,omitempty"` // 落地页内容
	List     []*ImageListItem `json:"list,omitempty"`      // 图集列表
}

// ImageListItem 图集列表项
type ImageListItem struct {
	ImageID         string    `json:"image_id"`                    // 素材图片id (必填)
	ImageURL        string    `json:"image_url,omitempty"`         // 素材图片url
	JumpInfo        *JumpInfo `json:"jump_info,omitempty"`         // 落地页内容
	OriginalImageID string    `json:"original_image_id,omitempty"` // 原始素材图片id
}

// VideoComponentValue 视频组件值
type VideoComponentValue struct {
	VideoID  string    `json:"video_id"`            // 视频id (必填)
	CoverID  string    `json:"cover_id,omitempty"`  // 视频封面图片id
	JumpInfo *JumpInfo `json:"jump_info,omitempty"` // 落地页内容
}

// BrandComponentValue 品牌形象组件值
type BrandComponentValue struct {
	BrandName    string    `json:"brand_name,omitempty"`     // 品牌名称
	BrandImageID string    `json:"brand_image_id,omitempty"` // 品牌形象图片id
	JumpInfo     *JumpInfo `json:"jump_info,omitempty"`      // 落地页内容
}

// MainJumpInfoComponentValue 主跳转组件值
type MainJumpInfoComponentValue struct {
	JumpInfo *JumpInfo `json:"jump_info,omitempty"` // 落地页内容
}

// ========== 落地页结构 ==========

// JumpInfo 落地页内容
type JumpInfo struct {
	PageType          string      `json:"page_type"`                     // 落地页类型 (必填)
	PageSpec          *PageSpec   `json:"page_spec"`                     // 落地页内容 (必填)
	JumpinfoAccountID int64       `json:"jumpinfo_account_id,omitempty"` // 推广帐号id
	PlatformType      string      `json:"platform_type,omitempty"`       // 落地页媒体平台类型
	Backups           []*JumpInfo `json:"backups,omitempty"`             // 兜底落地页
}

// PageSpec 落地页内容规格（根据 page_type 选填对应字段）
type PageSpec struct {
	AndroidAppSpec              *AndroidAppSpec              `json:"android_app_spec,omitempty"`                // Android应用
	IosAppSpec                  *IosAppSpec                  `json:"ios_app_spec,omitempty"`                    // iOS应用
	XjAndroidAppH5Spec          *XjPageSpec                  `json:"xj_android_app_h5_spec,omitempty"`          // 蹊径Android H5
	XjIosAppH5Spec              *XjPageSpec                  `json:"xj_ios_app_h5_spec,omitempty"`              // 蹊径iOS H5
	XjWebH5Spec                 *XjPageSpec                  `json:"xj_web_h5_spec,omitempty"`                  // 蹊径网页H5
	XjQuickSpec                 *XjPageSpec                  `json:"xj_quick_spec,omitempty"`                   // 蹊径快应用
	WechatCanvasSpec            *WechatCanvasSpec            `json:"wechat_canvas_spec,omitempty"`              // 微信原生页
	WechatMiniProgramSpec       *WechatMiniProgramSpec       `json:"wechat_mini_program_spec,omitempty"`        // 微信小程序
	WechatMiniGameSpec          *WechatMiniGameSpec          `json:"wechat_mini_game_spec,omitempty"`           // 微信小游戏
	WechatChannelsFeedSpec      *WechatChannelsFeedSpec      `json:"wechat_channels_feed_spec,omitempty"`       // 视频号动态
	WechatChannelsWatchLiveSpec *WechatChannelsWatchLiveSpec `json:"wechat_channels_watch_live_spec,omitempty"` // 视频号观看直播
	WechatShopSpec              *WechatShopSpec              `json:"wechat_shop_spec,omitempty"`                // 微信小店
	H5Spec                      *H5Spec                      `json:"h5_spec,omitempty"`                         // 自定义H5
}

// AndroidAppSpec Android应用落地页
type AndroidAppSpec struct {
	AndroidAppID       string `json:"android_app_id,omitempty"`        // 安卓应用AppId
	WechatCanvasPageID int64  `json:"wechat_canvas_page_id,omitempty"` // 落地页id
	AndroidChannelID   string `json:"android_channel_id,omitempty"`    // 安卓应用渠道包id
}

// IosAppSpec iOS应用落地页
type IosAppSpec struct {
	IosAppID string `json:"ios_app_id,omitempty"` // iOS应用AppId
}

// XjPageSpec 蹊径落地页（Android H5 / iOS H5 / 网页H5 / 快应用通用）
type XjPageSpec struct {
	PageID                      int64  `json:"page_id"`                                   // 落地页id (必填)
	WechatChannelsLiveReserveID string `json:"wechat_channels_live_reserve_id,omitempty"` // 视频号直播预约id
	CustomParam                 string `json:"custom_param,omitempty"`                    // 官方落地页自定义参数
}

// WechatCanvasSpec 微信原生页落地页
type WechatCanvasSpec struct {
	PageID                      int64  `json:"page_id"`                                   // 落地页id (必填)
	OverrideCanvasHeadOption    string `json:"override_canvas_head_option,omitempty"`     // 顶部素材替换关系
	WechatChannelsLiveReserveID string `json:"wechat_channels_live_reserve_id,omitempty"` // 视频号直播预约id
}

// WechatMiniProgramSpec 微信小程序落地页
type WechatMiniProgramSpec struct {
	MiniProgramID             string   `json:"mini_program_id,omitempty"`               // 小程序id
	MiniProgramPath           string   `json:"mini_program_path,omitempty"`             // 小程序路径
	MiniProgramPaths          []string `json:"mini_program_paths,omitempty"`            // 小程序path列表
	MpaMiniProgramWildcardURL string   `json:"mpa_mini_program_wildcard_url,omitempty"` // 通配符
}

// WechatMiniGameSpec 微信小游戏落地页
type WechatMiniGameSpec struct {
	MiniProgramID   string `json:"mini_program_id,omitempty"`   // 小游戏id
	MiniProgramPath string `json:"mini_program_path,omitempty"` // 小游戏路径
}

// WechatChannelsFeedSpec 视频号动态落地页
type WechatChannelsFeedSpec struct {
	FeedID                  string                    `json:"feed_id"`                              // 视频号动态id (必填)
	ActionButton            *ChannelsFeedActionButton `json:"action_button,omitempty"`              // 行动按钮
	WechatChannelsAccountID string                    `json:"wechat_channels_account_id,omitempty"` // 视频号账号id
}

// ChannelsFeedActionButton 视频号动态行动按钮
type ChannelsFeedActionButton struct {
	ButtonText string    `json:"button_text,omitempty"` // 按钮文案
	JumpInfo   *JumpInfo `json:"jump_info,omitempty"`   // 落地页内容
}

// WechatChannelsWatchLiveSpec 视频号观看直播落地页
type WechatChannelsWatchLiveSpec struct {
	WechatChannelsAccountID     string `json:"wechat_channels_account_id,omitempty"`      // 视频号账号id
	WechatChannelsLiveReserveID string `json:"wechat_channels_live_reserve_id,omitempty"` // 视频号直播预约id
}

// WechatShopSpec 微信小店落地页
type WechatShopSpec struct {
	ShopID string `json:"shop_id"` // 微信小店id (必填)
}

// H5Spec 自定义H5落地页
type H5Spec struct {
	PageURL          string `json:"page_url,omitempty"`            // 落地页url
	MpaH5WildcardURL string `json:"mpa_h5_wildcard_url,omitempty"` // 通配符
}

// ========== 组件值类型（补全）==========

// ConsultComponentValue 咨询组件值
type ConsultComponentValue struct {
	ID           int64       `json:"id,omitempty"`             // 咨询组件值
	JumpInfoList []*JumpInfo `json:"jump_info_list,omitempty"` // 兜底落地页内容列表
}

// PhoneComponentValue 电话组件值
type PhoneComponentValue struct {
	ID int64 `json:"id"` // 电话组件值 (必填)
}

// FormComponentValue 表单组件值
type FormComponentValue struct {
	ID int64 `json:"id"` // 表单组件值 (必填)
}

// ActionButtonComponentValue 行动按钮组件值
type ActionButtonComponentValue struct {
	ButtonText string    `json:"button_text,omitempty"` // 按钮文案
	JumpInfo   *JumpInfo `json:"jump_info,omitempty"`   // 落地页内容
}

// ChosenButtonComponentValue 选择按钮组件值
type ChosenButtonComponentValue struct {
	LeftButton  *ButtonItem `json:"left_button,omitempty"`  // 左按钮
	RightButton *ButtonItem `json:"right_button,omitempty"` // 右按钮
}

// ButtonItem 按钮项
type ButtonItem struct {
	Text     string    `json:"text,omitempty"`      // 按钮文案
	JumpInfo *JumpInfo `json:"jump_info,omitempty"` // 落地页内容
}

// 常量定义 - 标签类型
const (
	LabelTypeUnknown       = "LABEL_TYPE_UNKNOWN"       // 未知
	LabelTypeCommon        = "LABEL_TYPE_COMMON"        // 通用
	LabelTypePromotional   = "LABEL_TYPE_PROMOTIONAL"   // 促销
	LabelTypeCustomizeText = "LABEL_TYPE_CUSTOMIZETEXT" // 自定义文字
	LabelTypeIcon          = "LABEL_TYPE_ICON"          // 图标
)

// LabelComponentValue 标签组件值
type LabelComponentValue struct {
	List []*LabelItem `json:"list,omitempty"` // 标签列表
}

// LabelItem 标签列表项
type LabelItem struct {
	Content        string `json:"content,omitempty"`         // 标签内容
	Type           string `json:"type,omitempty"`            // 标签类型
	DisplayContent string `json:"display_content,omitempty"` // 标签显示内容，最大100字节
}

// ShowDataComponentValue 数据外显组件值
type ShowDataComponentValue struct {
	ConversionDataType   string `json:"conversion_data_type,omitempty"`   // 数据外显转换数据类型
	ConversionTargetType string `json:"conversion_target_type,omitempty"` // 数据外显转化目标量类型
}

// MarketingPendantComponentValue 营销挂件组件值
type MarketingPendantComponentValue struct {
	ImageID string `json:"image_id"` // 挂件图id (必填)
}

// AppGiftPackCodeComponentValue 礼包码组件值
type AppGiftPackCodeComponentValue struct {
	Code            string `json:"code,omitempty"`               // 礼包码，最大450字节
	Tips            string `json:"tips,omitempty"`               // 礼包码提示
	Description     string `json:"description,omitempty"`        // 礼包描述，最大36字节
	GameGiftID      string `json:"game_gift_id,omitempty"`       // 游戏圈礼包id
	GameActID       string `json:"game_act_id,omitempty"`        // 游戏圈活动id
	GameGiftImageID string `json:"game_gift_image_id,omitempty"` // 游戏圈礼包图片
}

// ShopImageComponentValue 卖点图组件值
type ShopImageComponentValue struct {
	ShopImageSwitch        bool   `json:"shop_image_switch,omitempty"`         // 卖点图开关
	DynamicShopImageSwitch bool   `json:"dynamic_shop_image_switch,omitempty"` // 卖点图轮播开关
	ShopImageID            string `json:"shop_image_id,omitempty"`             // 卖点图片id
	ShopImageTitle         string `json:"shop_image_title,omitempty"`          // 卖点图标题
	ShopImageDescription   string `json:"shop_image_description,omitempty"`    // 卖点图文案
}

// 常量定义 - 倒计时时间类型
const (
	CountdownTimeStart = "COUNTDOWN_TIME_START" // 开始时间
	CountdownTimeEnd   = "COUNTDOWN_TIME_END"   // 结束时间
)

// CountDownComponentValue 倒计时组件值
type CountDownComponentValue struct {
	Price             string `json:"price,omitempty"`              // 倒计时价格，单位：分
	TimeType          string `json:"time_type,omitempty"`          // 倒计时时间类型
	ExpiringTimestamp int64  `json:"expiring_timestamp,omitempty"` // 倒计时时间锚点，unix时间戳
}

// BarrageComponentValue 弹幕组件值
type BarrageComponentValue struct {
	List []*BarrageItem `json:"list,omitempty"` // 弹幕列表
}

// BarrageItem 弹幕列表项
type BarrageItem struct {
	ID   int64  `json:"id,omitempty"`   // 弹幕id
	Text string `json:"text,omitempty"` // 弹幕文案，1-12个字
}

// 常量定义 - 浮层卡片类型
const (
	FloatingZoneTypeUnknown     = "FLOATING_ZONE_TYPE_UNKNOWN"      // 未知
	FloatingZoneTypeImageText   = "FLOATING_ZONE_TYPE_IMAGE_TEXT"   // 图文
	FloatingZoneTypeSingleImage = "FLOATING_ZONE_TYPE_SINGLE_IMAGE" // 单图
	FloatingZoneTypeMultiButton = "FLOATING_ZONE_TYPE_MULTI_BUTTON" // 多按钮
	FloatingZoneTypeSliderCard  = "FLOATING_ZONE_TYPE_SLIDER_CARD"  // 滑动卡片
)

// FloatingZoneComponentValue 浮层卡片组件值
type FloatingZoneComponentValue struct {
	FloatingZoneSwitch                bool      `json:"floating_zone_switch,omitempty"`                   // 浮层卡片开关
	FloatingZoneImageID               string    `json:"floating_zone_image_id,omitempty"`                 // 浮层卡片图片id（512*512，不超过50KB）
	FloatingZoneName                  string    `json:"floating_zone_name,omitempty"`                     // 文案一，1-10等宽字符
	FloatingZoneDesc                  string    `json:"floating_zone_desc,omitempty"`                     // 文案二，1-14等宽字符
	FloatingZoneButtonText            string    `json:"floating_zone_button_text,omitempty"`              // 按钮文案，1-10等宽字符
	FloatingZoneShowAppPropertySwitch bool      `json:"floating_zone_show_app_property_switch,omitempty"` // 显示已下载人数及评分开关
	FloatingZoneType                  string    `json:"floating_zone_type,omitempty"`                     // 浮层卡片类型
	FloatingZoneSingleImageID         string    `json:"floating_zone_single_image_id,omitempty"`          // 单图片id（482*270，不超过50KB）
	ButtonBaseText                    string    `json:"button_base_text,omitempty"`                       // 视频号基础态文案内容，最大10字节
	JumpInfo                          *JumpInfo `json:"jump_info,omitempty"`                              // 落地页内容
	FloatingZoneInfoType              string    `json:"floating_zone_info_type,omitempty"`                // 浮层卡片外显类型
}

// TextLinkComponentValue 文字链组件值
type TextLinkComponentValue struct {
	LinkNameType string    `json:"link_name_type,omitempty"` // 链接名称类型
	LinkNameText string    `json:"link_name_text,omitempty"` // 文字链文案
	JumpInfo     *JumpInfo `json:"jump_info,omitempty"`      // 落地页内容
}

// EndPageComponentValue 视频结束页组件值
type EndPageComponentValue struct {
	EndPageType string `json:"end_page_type,omitempty"` // 结束页类型
	EndPageDesc string `json:"end_page_desc,omitempty"` // 结束页文案，最大192字节
}

// LivingDescComponentValue 轮播文案组件值
type LivingDescComponentValue struct {
	LivingDescSwitch bool     `json:"living_desc_switch,omitempty"` // 轮播组件开关
	DescList         []string `json:"desc_list,omitempty"`          // 轮播文案，2-5条，每条最大1024字节
}

// 常量定义 - 视频号直播推广形式
const (
	LivePromotedTypeNativeVideo = "LIVE_PROMOTED_TYPE_NATIVE_VIDEO" // 原生视频
	LivePromotedTypeShortVideo  = "LIVE_PROMOTED_TYPE_SHORT_VIDEO"  // 短视频
)

// WechatChannelsComponentValue 视频号信息组件值
type WechatChannelsComponentValue struct {
	LivePromotedType       string `json:"live_promoted_type,omitempty"`       // 视频号直播推广形式
	Username               string `json:"username,omitempty"`                 // 视频号username，最大1024字节
	FinderObjectVisibility bool   `json:"finder_object_visibility,omitempty"` // 是否保存至视频号
}

// ShortVideoComponentValue 短视频组件值
type ShortVideoComponentValue struct {
	ShortVideo1 string `json:"short_video1,omitempty"` // 视频id，最大64字节
	ShortVideo2 string `json:"short_video2,omitempty"` // 视频id，最大64字节
}

// ElementStoryComponentValue 集装箱创意组合组件值
type ElementStoryComponentValue struct {
	List []*ElementStoryItem `json:"list,omitempty"` // 集装箱创意组合，1-14条
}

// ElementStoryItem 集装箱创意组合项
type ElementStoryItem struct {
	Image       string `json:"image,omitempty"`       // 图片id
	Image2      string `json:"image2,omitempty"`      // 素材图片2的id
	Description string `json:"description,omitempty"` // 广告描述
	URL         string `json:"url,omitempty"`         // 跳转链接
	Title       string `json:"title,omitempty"`       // 广告文案
}

// WxgamePlayablePageComponentValue 小游戏试玩页组件值
type WxgamePlayablePageComponentValue struct {
	WxgamePlayablePageSwitch        bool     `json:"wxgame_playable_page_switch,omitempty"`                // 小游戏试玩页开关
	WxgamePlayablePagePath          string   `json:"wxgame_playable_page_path,omitempty"`                  // 小游戏试玩页，最大1024字节
	WxgamePlayablePageEndCoverImg   string   `json:"wxgame_playable_page_end_cover_img,omitempty"`         // 结束页图片id，最大64字节
	WxgamePlayablePageEndDesc       string   `json:"wxgame_playable_page_end_desc,omitempty"`              // 结束页文案，最大200字节
	WxgamePlayablePageTriggerTypes  []string `json:"wxgame_playable_page_trigger_types,omitempty"`         // 开启方式，1-100条
	WxgamePlayablePageTriggerText   string   `json:"wxgame_playable_page_trigger_text,omitempty"`          // 开启文案
	WxgamePlayablePageCardLinkImage string   `json:"wxgame_playable_page_card_link_image,omitempty"`       // 图文链接图片id，最大64字节
	WxgamePlayablePageCardLinkDesc  string   `json:"wxgame_playable_page_card_link_description,omitempty"` // 图文链接描述文案
	WxgamePlayablePageEndTimeType   string   `json:"wxgame_playable_page_end_time_type,omitempty"`         // 结束时间类型
}

// QrcodePosition 二维码坐标
type QrcodePosition struct {
	PositionX string `json:"position_x,omitempty"` // x坐标
	PositionY string `json:"position_y,omitempty"` // y坐标
}

// AppPromotionVideoComponentValue OTT视频组件值
type AppPromotionVideoComponentValue struct {
	Video          string          `json:"video,omitempty"`           // 视频id
	Video2         string          `json:"video2,omitempty"`          // 视频id
	Video3         string          `json:"video3,omitempty"`          // 视频id
	AllowTvQrcode  bool            `json:"allow_tv_qrcode,omitempty"` // 支持TV二维码
	QrcodePosition *QrcodePosition `json:"qrcode_position,omitempty"` // 二维码坐标
	QrcodeWidth    int64           `json:"qrcode_width,omitempty"`    // 二维码边长
}

// VideoShowcaseVideo 橱窗视频
type VideoShowcaseVideo struct {
	VideoID  string    `json:"video_id"`            // 视频id (必填)
	CoverID  string    `json:"cover_id,omitempty"`  // 封面图片id
	JumpInfo *JumpInfo `json:"jump_info,omitempty"` // 落地页内容
}

// VideoShowcaseComponentValue 橱窗视频组件值
type VideoShowcaseComponentValue struct {
	Video     *VideoShowcaseVideo      `json:"video"`                // 视频 (必填)
	ImageList *ImageListComponentValue `json:"image_list,omitempty"` // 图集
}

// ImageShowcaseComponentValue 橱窗图片组件值
type ImageShowcaseComponentValue struct {
	Image     *ImageComponentValue     `json:"image"`                // 图片 (必填)
	ImageList *ImageListComponentValue `json:"image_list,omitempty"` // 图集
}

// SocialSkillComponentValue 首评回复组件值
type SocialSkillComponentValue struct {
	SocialSkillFirstCommentSwitch bool   `json:"social_skill_first_comment_switch,omitempty"` // 首条评论开关
	SocialSkillFirstComment       string `json:"social_skill_first_comment,omitempty"`        // 首条评论内容，最大200字节
}

// MiniCardLinkComponentValue 图文链接组件值
type MiniCardLinkComponentValue struct {
	MiniCardLinkDescription string `json:"mini_card_link_description,omitempty"` // 图文链接描述文案
	MiniCardLinkImage       string `json:"mini_card_link_image,omitempty"`       // 图文链接图片，最大384字节
}

// ========== 更新创意 ==========

// DynamicCreativesUpdateReq 更新创意请求
// https://developers.e.qq.com/v3.0/docs/api/dynamic_creatives/update
type DynamicCreativesUpdateReq struct {
	GlobalReq
	AccountID                        int64               `json:"account_id"`                                     // 广告主帐号id (必填)
	DynamicCreativeID                int64               `json:"dynamic_creative_id"`                            // 广告创意id (必填)
	DynamicCreativeName              string              `json:"dynamic_creative_name,omitempty"`                // 广告创意名称，同账号下不重复，1-60等宽字符
	CreativeComponents               *CreativeComponents `json:"creative_components,omitempty"`                  // 创意组件
	ImpressionTrackingURL            string              `json:"impression_tracking_url,omitempty"`              // 曝光监控地址
	ClickTrackingURL                 string              `json:"click_tracking_url,omitempty"`                   // 点击监控链接
	AutoDerivedProgramCreativeSwitch bool                `json:"auto_derived_program_creative_switch,omitempty"` // 自动衍生程序化创意开关
	ConfiguredStatus                 string              `json:"configured_status,omitempty"`                    // 配置状态
	IsRetryBatchUpdate               bool                `json:"is_retry_batch_update,omitempty"`                // 是否重试批量更新
	SiteSetValidateModel             string              `json:"site_set_validate_model,omitempty"`              // 站点集验证模型
}

func (p *DynamicCreativesUpdateReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证更新创意请求参数
func (p *DynamicCreativesUpdateReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.DynamicCreativeID == 0 {
		return errors.New("dynamic_creative_id为必填")
	}
	if len(p.DynamicCreativeName) > MaxDynamicCreativeNameBytes {
		return errors.New("dynamic_creative_name长度不能超过180字节")
	}
	return p.GlobalReq.Validate()
}

// DynamicCreativesUpdateResp 更新创意响应
type DynamicCreativesUpdateResp struct {
	DynamicCreativeID int64 `json:"dynamic_creative_id"` // 更新的动态创意id
}

// ========== 删除创意 ==========

// DynamicCreativesDeleteReq 删除创意请求
// https://developers.e.qq.com/v3.0/docs/api/dynamic_creatives/delete
type DynamicCreativesDeleteReq struct {
	GlobalReq
	AccountID         int64 `json:"account_id"`          // 广告主帐号id (必填)
	DynamicCreativeID int64 `json:"dynamic_creative_id"` // 广告创意id (必填)
}

func (p *DynamicCreativesDeleteReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证删除创意请求参数
func (p *DynamicCreativesDeleteReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.DynamicCreativeID == 0 {
		return errors.New("dynamic_creative_id为必填")
	}
	return p.GlobalReq.Validate()
}

// DynamicCreativesDeleteResp 删除创意响应
type DynamicCreativesDeleteResp struct {
	DynamicCreativeID int64 `json:"dynamic_creative_id"` // 删除的动态创意id
}
