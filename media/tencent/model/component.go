package model

import "errors"

// ========== 获取创意组件 ==========
// https://developers.e.qq.com/v3.0/docs/api/components/get

// 常量定义 - 组件过滤字段
const (
	ComponentFieldComponentID            = "component_id"
	ComponentFieldComponentType          = "component_type"
	ComponentFieldComponentSubType       = "component_sub_type"
	ComponentFieldCreatedTime            = "created_time"
	ComponentFieldLastModifiedTime       = "last_modified_time"
	ComponentFieldGenerationType         = "generation_type"
	ComponentFieldPotentialStatus        = "potential_status"
	ComponentFieldVideoID                = "video_id"
	ComponentFieldImageID                = "image_id"
	ComponentFieldVideoSignature         = "video_signature"
	ComponentFieldImageSignature         = "image_signature"
	ComponentFieldFirstPublicationStatus = "first_publication_status"
	ComponentFieldSimilarityStatus       = "similarity_status"
	ComponentFieldScene                  = "scene"
)

// 常量定义 - 共享组件读取方式
const (
	ComponentIDFilteringModeSharingByAgencyBusinessUnit   = "SHARING_BY_AGENCY_BUSINESS_UNIT"   // 代理商业务单元共享
	ComponentIDFilteringModeSharingByCustomerBusinessUnit = "SHARING_BY_CUSTOMER_BUSINESS_UNIT" // 广告主业务单元共享
)

// 常量定义 - 组件生成类型
const (
	ComponentGenerationTypeManual = "GENERATION_TYPE_MANUAL" // 手动
	ComponentGenerationTypeAuto   = "GENERATION_TYPE_AUTO"   // 自动
)

// 常量定义 - 组件潜力状态
const (
	ComponentPotentialStatusHigh       = "POTENTIAL_STATUS_HIGH"        // 高潜力
	ComponentPotentialStatusLow        = "POTENTIAL_STATUS_LOW"         // 低潜力
	ComponentPotentialStatusNoJudgment = "POTENTIAL_STATUS_NO_JUDGMENT" // 无判断
)

// 常量定义 - 组件首次发布状态
const (
	ComponentFirstPublicationStatusYes = "FIRST_PUBLICATION_YES" // 首次发布
	ComponentFirstPublicationStatusNo  = "FIRST_PUBLICATION_NO"  // 非首次发布
)

// 常量定义 - 相似度检测状态
const (
	ComponentSimilarityStatusPass    = "SIMILARITY_STATUS_PASS"    // 通过
	ComponentSimilarityStatusReject  = "SIMILARITY_STATUS_REJECT"  // 拒绝
	ComponentSimilarityStatusPending = "SIMILARITY_STATUS_PENDING" // 待检测
)

// 字段限制常量
const (
	MaxComponentFilteringCount = 10
	DefaultComponentPage       = 1
	DefaultComponentPageSize   = 10
)

// ComponentsGetReq 获取创意组件请求
// https://developers.e.qq.com/v3.0/docs/api/components/get
type ComponentsGetReq struct {
	GlobalReq
	AccountID                int64                   `json:"account_id"`                            // 广告主帐号id (必填)
	OrganizationID           int64                   `json:"organization_id,omitempty"`             // 业务单元id
	Filtering                []*ComponentQueryFilter `json:"filtering,omitempty"`                   // 过滤条件，数组长度0-10
	Page                     int                     `json:"page,omitempty"`                        // 搜索页码，默认1，最大100
	PageSize                 int                     `json:"page_size,omitempty"`                   // 每页条数，默认10，最大100000
	IsDeleted                bool                    `json:"is_deleted,omitempty"`                  // 是否已删除
	Fields                   []string                `json:"fields,omitempty"`                      // 指定返回的字段列表
	ComponentIDFilteringMode string                  `json:"component_id_filtering_mode,omitempty"` // 共享组件读取方式
}

// ComponentQueryFilter 组件查询过滤条件
type ComponentQueryFilter struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)
	Operator string   `json:"operator"` // 操作符 (必填)
	Values   []string `json:"values"`   // 字段取值 (必填)
}

func (p *ComponentsGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page <= 0 {
		p.Page = DefaultComponentPage
	}
	if p.PageSize <= 0 {
		p.PageSize = DefaultComponentPageSize
	}
}

// Validate 验证获取创意组件请求参数
func (p *ComponentsGetReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if err := p.GlobalReq.Validate(); err != nil {
		return err
	}
	if p.Page < MinPage || p.Page > MaxPage {
		return errors.New("page必须在1-100之间")
	}
	if p.PageSize < MinPageSize || p.PageSize > 100000 {
		return errors.New("page_size必须在1-100000之间")
	}
	if p.ComponentIDFilteringMode != "" &&
		p.ComponentIDFilteringMode != ComponentIDFilteringModeSharingByAgencyBusinessUnit &&
		p.ComponentIDFilteringMode != ComponentIDFilteringModeSharingByCustomerBusinessUnit {
		return errors.New("component_id_filtering_mode值无效")
	}
	return validateComponentFiltering(p.Filtering)
}

// validateComponentFiltering 验证组件过滤条件
func validateComponentFiltering(filtering []*ComponentQueryFilter) error {
	if len(filtering) == 0 {
		return nil
	}
	if len(filtering) > MaxComponentFilteringCount {
		return errors.New("filtering数组长度不能超过10")
	}
	for _, f := range filtering {
		if err := f.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Validate 验证单个组件过滤条件
func (f *ComponentQueryFilter) Validate() error {
	if f.Field == "" {
		return errors.New("field为必填")
	}
	if !isValidComponentField(f.Field) {
		return errors.New("field值无效，请参考文档中的允许值")
	}
	if f.Operator == "" {
		return errors.New("operator为必填")
	}
	if !isValidComponentOperatorForField(f.Field, f.Operator) {
		return errors.New("operator值无效，当前字段不支持该操作符")
	}
	if len(f.Values) == 0 {
		return errors.New("values为必填")
	}
	return validateComponentValuesForField(f)
}

// isValidComponentField 验证组件过滤字段是否有效
func isValidComponentField(field string) bool {
	validFields := map[string]bool{
		ComponentFieldComponentID:            true,
		ComponentFieldComponentType:          true,
		ComponentFieldComponentSubType:       true,
		ComponentFieldCreatedTime:            true,
		ComponentFieldLastModifiedTime:       true,
		ComponentFieldGenerationType:         true,
		ComponentFieldPotentialStatus:        true,
		ComponentFieldVideoID:                true,
		ComponentFieldImageID:                true,
		ComponentFieldVideoSignature:         true,
		ComponentFieldImageSignature:         true,
		ComponentFieldFirstPublicationStatus: true,
		ComponentFieldSimilarityStatus:       true,
		ComponentFieldScene:                  true,
	}
	return validFields[field]
}

// isValidComponentOperatorForField 验证组件字段支持的操作符
func isValidComponentOperatorForField(field, operator string) bool {
	switch field {
	case ComponentFieldComponentID:
		return operator == OperatorEquals || operator == OperatorIn
	case ComponentFieldComponentType, ComponentFieldComponentSubType,
		ComponentFieldGenerationType, ComponentFieldPotentialStatus,
		ComponentFieldFirstPublicationStatus, ComponentFieldSimilarityStatus,
		ComponentFieldScene:
		return operator == OperatorEquals || operator == OperatorIn
	case ComponentFieldCreatedTime, ComponentFieldLastModifiedTime:
		return operator == OperatorEquals || operator == OperatorLess ||
			operator == OperatorLessEquals || operator == OperatorGreater ||
			operator == OperatorGreaterEquals
	case ComponentFieldVideoID, ComponentFieldImageID,
		ComponentFieldVideoSignature, ComponentFieldImageSignature:
		return operator == OperatorEquals || operator == OperatorIn
	default:
		return false
	}
}

// validateComponentValuesForField 验证组件字段取值
func validateComponentValuesForField(f *ComponentQueryFilter) error {
	switch f.Field {
	case ComponentFieldComponentID:
		if f.Operator == OperatorEquals && len(f.Values) != 1 {
			return errors.New("operator为EQUALS时，values数组长度必须为1")
		}
		if f.Operator == OperatorIn && (len(f.Values) < MinValuesCount || len(f.Values) > 100) {
			return errors.New("component_id IN操作时，values数组长度必须在1-100之间")
		}
	case ComponentFieldCreatedTime, ComponentFieldLastModifiedTime:
		if len(f.Values) != 1 {
			return errors.New("时间字段values数组长度必须为1")
		}
		if len(f.Values[0]) != CreatedTimeLength {
			return errors.New("时间字段长度必须为10字节")
		}
	case ComponentFieldVideoID, ComponentFieldImageID,
		ComponentFieldVideoSignature, ComponentFieldImageSignature:
		if f.Operator == OperatorIn && (len(f.Values) < MinValuesCount || len(f.Values) > MaxValuesCount) {
			return errors.New("operator为IN时，values数组长度必须在1-100之间")
		}
		if f.Operator == OperatorEquals && len(f.Values) != 1 {
			return errors.New("operator为EQUALS时，values数组长度必须为1")
		}
	}
	return nil
}

// ComponentsGetListItem 创意组件列表项
type ComponentsGetListItem struct {
	AccountID              int64               `json:"account_id"`                         // 广告主帐号id
	OrganizationID         int64               `json:"organization_id,omitempty"`          // 业务单元id
	ComponentID            int64               `json:"component_id"`                       // 创意组件id
	ComponentValue         *CreativeComponents `json:"component_value,omitempty"`          // 创意组件内容
	CreatedTime            int64               `json:"created_time,omitempty"`             // 创建时间，时间戳
	LastModifiedTime       int64               `json:"last_modified_time,omitempty"`       // 最后修改时间，时间戳
	ComponentSubType       string              `json:"component_sub_type,omitempty"`       // 创意组件子类型
	ComponentCustomName    string              `json:"component_custom_name,omitempty"`    // 创意组件自定义名称
	GenerationType         string              `json:"generation_type,omitempty"`          // 创意组件生成类型
	IsDeleted              bool                `json:"is_deleted,omitempty"`               // 是否已删除
	SimilarityStatus       string              `json:"similarity_status,omitempty"`        // 相似度检测状态
	PotentialStatus        string              `json:"potential_status,omitempty"`         // 组件潜力
	DisableMessage         string              `json:"disable_message,omitempty"`          // 不可用错误信息
	FirstPublicationStatus string              `json:"first_publication_status,omitempty"` // 组件首次发布状态
	Scene                  string              `json:"scene,omitempty"`                    // 创意组件适用场景
}

// ComponentsGetResp 获取创意组件响应
type ComponentsGetResp struct {
	List []*ComponentsGetListItem `json:"list,omitempty"` // 创意组件列表
	PageInfoContainer
}

// ========== 创建创意组件 ==========
// https://developers.e.qq.com/v3.0/docs/api/components/add

// 常量定义 - 组件子类型 (component_sub_type)
const (
	// 视频类
	ComponentSubTypeVideo16X9     = "VIDEO_16X9"      // 视频 16:9
	ComponentSubTypeVideo9X16     = "VIDEO_9X16"      // 视频 9:16
	ComponentSubTypeVideo4X3      = "VIDEO_4X3"       // 视频 4:3
	ComponentSubTypeVideoShowcase = "VIDEO_SHOWCASE"  // 橱窗视频
	ComponentSubTypeShortVideo4X3 = "SHORT_VIDEO_4X3" // 短视频 4:3

	// 图片类
	ComponentSubTypeImage16X9     = "IMAGE_16X9"     // 图片 16:9
	ComponentSubTypeImage9X16     = "IMAGE_9X16"     // 图片 9:16
	ComponentSubTypeImage1X1      = "IMAGE_1X1"      // 图片 1:1
	ComponentSubTypeImage3X2      = "IMAGE_3X2"      // 图片 3:2
	ComponentSubTypeImage3X4      = "IMAGE_3X4"      // 图片 3:4
	ComponentSubTypeImage4X3      = "IMAGE_4X3"      // 图片 4:3
	ComponentSubTypeImage5X4      = "IMAGE_5X4"      // 图片 5:4
	ComponentSubTypeImage4X5      = "IMAGE_4X5"      // 图片 4:5
	ComponentSubTypeImage20X7     = "IMAGE_20X7"     // 图片 20:7
	ComponentSubTypeImage7X2      = "IMAGE_7X2"      // 图片 7:2
	ComponentSubTypeImageShowcase = "IMAGE_SHOWCASE" // 橱窗图片
	ComponentSubTypeImage100X9    = "IMAGE_100X9"    // 图片 100:9

	// 图集类
	ComponentSubTypeImageList9X16_4 = "IMAGE_LIST_9X16_4" // 图集 9:16×4
	ComponentSubTypeImageList1X1_3  = "IMAGE_LIST_1X1_3"  // 图集 1:1×3
	ComponentSubTypeImageList1X1_4  = "IMAGE_LIST_1X1_4"  // 图集 1:1×4
	ComponentSubTypeImageList1X1_6  = "IMAGE_LIST_1X1_6"  // 图集 1:1×6
	ComponentSubTypeImageList3X2_3  = "IMAGE_LIST_3X2_3"  // 图集 3:2×3
	ComponentSubTypeImageList1X1_1  = "IMAGE_LIST_1X1_1"  // 图集 1:1×1
	ComponentSubTypeImageList16X9_1 = "IMAGE_LIST_16X9_1" // 图集 16:9×1
	ComponentSubTypeImageList1X1_9  = "IMAGE_LIST_1X1_9"  // 图集 1:1×9

	// 文案类
	ComponentSubTypeElementStory = "ELEMENT_STORY" // 集装箱创意组合
	ComponentSubTypeDescription  = "DESCRIPTION"   // 描述
	ComponentSubTypeTitle        = "TITLE"         // 标题

	// 交互类
	ComponentSubTypeActionButton              = "ACTION_BUTTON"                 // 行动按钮
	ComponentSubTypeLabel                     = "LABEL"                         // 标签
	ComponentSubTypeShowData                  = "SHOW_DATA"                     // 数据外显
	ComponentSubTypeFloatingZoneImageText     = "FLOATING_ZONE_IMAGE_TEXT"      // 浮层卡片（图文）
	ComponentSubTypeFloatingZoneImage         = "FLOATING_ZONE_IMAGE"           // 浮层卡片（单图）
	ComponentSubTypeBarrage                   = "BARRAGE"                       // 弹幕
	ComponentSubTypeAppGiftPackCode           = "APP_GIFT_PACK_CODE"            // 礼包码
	ComponentSubTypeShopImage                 = "SHOP_IMAGE"                    // 卖点图
	ComponentSubTypeMarketingPendant          = "MARKETING_PENDANT"             // 营销挂件
	ComponentSubTypeChosenButton              = "CHOSEN_BUTTON"                 // 选择按钮
	ComponentSubTypeCountDown                 = "COUNT_DOWN"                    // 倒计时
	ComponentSubTypeLivingDesc                = "LIVING_DESC"                   // 轮播文案
	ComponentSubTypeTextLink                  = "TEXT_LINK"                     // 文字链
	ComponentSubTypeEndPage                   = "END_PAGE"                      // 视频结束页
	ComponentSubTypeWxgamePlayablePage        = "WXGAME_PLAYABLE_PAGE"          // 小游戏试玩页
	ComponentSubTypeSocialSkill               = "SOCIAL_SKILL"                  // 首评回复
	ComponentSubTypeMiniCardLink              = "MINI_CARD_LINK"                // 图文链接
	ComponentSubTypeFloatingZoneImageTextList = "FLOATING_ZONE_IMAGE_TEXT_LIST" // 多卡轮播
	ComponentSubTypeConsultLink               = "CONSULT_LINK"                  // 咨询链接
	ComponentSubTypeAudio                     = "AUDIO"                         // 音频
	ComponentSubTypeWechatShopActivityBulkBuy = "WECHAT_SHOP_ACTIVITY_BULK_BUY" // 微信小店活动团购
	ComponentSubTypeWxgameDirectPage          = "WXGAME_DIRECT_PAGE"            // 小游戏直玩

	// 品牌类
	ComponentSubTypeBrand              = "BRAND"                // 品牌形象
	ComponentSubTypeBrandPage          = "BRAND_PAGE"           // 品牌落地页
	ComponentSubTypeBrandSearch        = "BRAND_SEARCH"         // 品牌搜索
	ComponentSubTypeBrandWechatChannel = "BRAND_WECHAT_CHANNEL" // 品牌视频号
	ComponentSubTypeBrandWechat        = "BRAND_WECHAT"         // 品牌微信
	ComponentSubTypeBrandWecom         = "BRAND_WECOM"          // 品牌企业微信
	ComponentSubTypeBrandWechatShop    = "BRAND_WECHAT_SHOP"    // 品牌微信小店
	ComponentSubTypeBrandCustomLink    = "BRAND_CUSTOM_LINK"    // 品牌自定义链接

	// 跳转信息类
	ComponentSubTypeJumpInfoOfficial                    = "JUMP_INFO_OFFICIAL"                       // 官方落地页跳转
	ComponentSubTypeJumpInfoH5                          = "JUMP_INFO_H5"                             // H5跳转
	ComponentSubTypeJumpInfoWechatMiniProgram           = "JUMP_INFO_WECHAT_MINI_PROGRAM"            // 微信小程序跳转
	ComponentSubTypeJumpInfoWechatConsult               = "JUMP_INFO_WECHAT_CONSULT"                 // 微信客服跳转
	ComponentSubTypeJumpInfoWecomConsult                = "JUMP_INFO_WECOM_CONSULT"                  // 企业微信客服跳转
	ComponentSubTypeJumpInfoWechatChannelsWatchLive     = "JUMP_INFO_WECHAT_CHANNELS_WATCH_LIVE"     // 视频号观看直播跳转
	ComponentSubTypeJumpInfoWechatChannelsFeed          = "JUMP_INFO_WECHAT_CHANNELS_FEED"           // 视频号动态跳转
	ComponentSubTypeJumpInfoWechatOfficialAccountDetail = "JUMP_INFO_WECHAT_OFFICIAL_ACCOUNT_DETAIL" // 微信公众号详情跳转
	ComponentSubTypeJumpInfoWechatMiniGame              = "JUMP_INFO_WECHAT_MINI_GAME"               // 微信小游戏跳转
	ComponentSubTypeJumpInfoAndroidApp                  = "JUMP_INFO_ANDROID_APP"                    // Android应用跳转
	ComponentSubTypeJumpInfoIosApp                      = "JUMP_INFO_IOS_APP"                        // iOS应用跳转
	ComponentSubTypeJumpInfoAndroidDirectDownload       = "JUMP_INFO_ANDROID_DIRECT_DOWNLOAD"        // Android一键下载跳转
	ComponentSubTypeJumpInfoAppMarket                   = "JUMP_INFO_APP_MARKET"                     // 厂商下载跳转
	ComponentSubTypeJumpInfoAppDeepLink                 = "JUMP_INFO_APP_DEEP_LINK"                  // 应用直达跳转
	ComponentSubTypeJumpInfoWechatChannelsShopProduct   = "JUMP_INFO_WECHAT_CHANNELS_SHOP_PRODUCT"   // 微信小店商品跳转
	ComponentSubTypeJumpInfoQqMiniGame                  = "JUMP_INFO_QQ_MINI_GAME"                   // QQ小游戏跳转
)

// 创意组件字段限制常量
const (
	MaxComponentCustomNameBytes = 512 // component_custom_name 最大字节数
)

// ComponentsAddReq 创建创意组件请求
// https://developers.e.qq.com/v3.0/docs/api/components/add
type ComponentsAddReq struct {
	GlobalReq
	AccountID           int64               `json:"account_id,omitempty"`            // 广告主帐号id
	OrganizationID      int64               `json:"organization_id,omitempty"`       // 业务单元id
	ComponentSubType    string              `json:"component_sub_type"`              // 创意组件子类型 (必填)
	ComponentValue      *CreativeComponents `json:"component_value"`                 // 创意组件内容 (必填)
	ComponentCustomName string              `json:"component_custom_name,omitempty"` // 创意组件自定义名称，1-512字节
}

func (p *ComponentsAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证创建创意组件请求参数
func (p *ComponentsAddReq) Validate() error {
	if p.ComponentSubType == "" {
		return errors.New("component_sub_type为必填")
	}
	if p.ComponentValue == nil {
		return errors.New("component_value为必填")
	}
	if len(p.ComponentCustomName) > MaxComponentCustomNameBytes {
		return errors.New("component_custom_name长度不能超过512字节")
	}
	return p.GlobalReq.Validate()
}

// ComponentsAddResp 创建创意组件响应
type ComponentsAddResp struct {
	ComponentID int64 `json:"component_id"` // 创建的创意组件id
}
