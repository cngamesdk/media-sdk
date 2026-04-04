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

// ========== 删除创意组件 ==========
// https://developers.e.qq.com/v3.0/docs/api/components/delete

// 常量定义 - 组件删除策略
const (
	DeleteStrategyForce      = "DELETE_STRATEGY_FORCE"      // 强制删除
	DeleteStrategyRestricted = "DELETE_STRATEGY_RESTRICTED" // 受限删除（使用中不可删除）
)

// ComponentsDeleteReq 删除创意组件请求
// https://developers.e.qq.com/v3.0/docs/api/components/delete
type ComponentsDeleteReq struct {
	GlobalReq
	AccountID      int64  `json:"account_id,omitempty"`      // 广告主帐号id
	OrganizationID int64  `json:"organization_id,omitempty"` // 业务单元id
	ComponentID    int64  `json:"component_id"`              // 创意组件id (必填)
	DeleteStrategy string `json:"delete_strategy,omitempty"` // 删除策略
}

func (p *ComponentsDeleteReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证删除创意组件请求参数
func (p *ComponentsDeleteReq) Validate() error {
	if p.ComponentID == 0 {
		return errors.New("component_id为必填")
	}
	if p.DeleteStrategy != "" &&
		p.DeleteStrategy != DeleteStrategyForce &&
		p.DeleteStrategy != DeleteStrategyRestricted {
		return errors.New("delete_strategy值无效，允许值：DELETE_STRATEGY_FORCE、DELETE_STRATEGY_RESTRICTED")
	}
	return p.GlobalReq.Validate()
}

// ComponentsDeleteResp 删除创意组件响应
type ComponentsDeleteResp struct {
	ComponentID int64 `json:"component_id"` // 删除的创意组件id
}

// ========== 修改创意组件共享 ==========
// https://developers.e.qq.com/v3.0/docs/api/component_sharing/update

// 常量定义 - 共享账号类型
const (
	SharedAccountTypeInvalid      = "INVALID"      // 无效
	SharedAccountTypeAdvertiser   = "ADVERTISER"   // 广告主账号
	SharedAccountTypeOrganization = "ORGANIZATION" // 业务单元
)

// SharedAccount 共享账号信息
type SharedAccount struct {
	SharedAccountID   int64  `json:"shared_account_id"`   // 共享账号id (必填)
	SharedAccountType string `json:"shared_account_type"` // 共享账号类型 (必填)
}

// ComponentSharingUpdateReq 修改创意组件共享请求
// https://developers.e.qq.com/v3.0/docs/api/component_sharing/update
type ComponentSharingUpdateReq struct {
	GlobalReq
	OrganizationID    int64            `json:"organization_id"`     // 业务单元id (必填)
	ComponentID       int64            `json:"component_id"`        // 创意组件id (必填)
	SharedAccountList []*SharedAccount `json:"shared_account_list"` // 共享账号列表 (必填)，1-100条
}

func (p *ComponentSharingUpdateReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证修改创意组件共享请求参数
func (p *ComponentSharingUpdateReq) Validate() error {
	if p.OrganizationID == 0 {
		return errors.New("organization_id为必填")
	}
	if p.ComponentID == 0 {
		return errors.New("component_id为必填")
	}
	if len(p.SharedAccountList) == 0 {
		return errors.New("shared_account_list为必填，至少包含1条记录")
	}
	if len(p.SharedAccountList) > 100 {
		return errors.New("shared_account_list数组长度不能超过100")
	}
	for _, a := range p.SharedAccountList {
		if a.SharedAccountID == 0 {
			return errors.New("shared_account_id为必填")
		}
		if a.SharedAccountType == "" {
			return errors.New("shared_account_type为必填")
		}
		if a.SharedAccountType != SharedAccountTypeInvalid &&
			a.SharedAccountType != SharedAccountTypeAdvertiser &&
			a.SharedAccountType != SharedAccountTypeOrganization {
			return errors.New("shared_account_type值无效，允许值：INVALID、ADVERTISER、ORGANIZATION")
		}
	}
	return p.GlobalReq.Validate()
}

// ComponentSharingUpdateResp 修改创意组件共享响应
type ComponentSharingUpdateResp struct {
	ComponentID int64 `json:"component_id"` // 创意组件id
}

// ========== 查询创意组件共享信息 ==========
// https://developers.e.qq.com/v3.0/docs/api/component_sharing/get

// 分页限制常量（共享查询接口特有）
const (
	MaxComponentSharingPage = 99999 // page 最大值
)

// ComponentSharingGetReq 查询创意组件共享信息请求
// https://developers.e.qq.com/v3.0/docs/api/component_sharing/get
type ComponentSharingGetReq struct {
	GlobalReq
	OrganizationID int64 `json:"organization_id"`        // 业务单元id (必填)
	ComponentID    int64 `json:"component_id,omitempty"` // 创意组件id
	Page           int   `json:"page,omitempty"`         // 搜索页码，默认1，最大99999
	PageSize       int   `json:"page_size,omitempty"`    // 每页条数，默认10，最大100
	IsDeleted      bool  `json:"is_deleted,omitempty"`   // 是否已删除
}

func (p *ComponentSharingGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page <= 0 {
		p.Page = DefaultComponentPage
	}
	if p.PageSize <= 0 {
		p.PageSize = DefaultComponentPageSize
	}
}

// Validate 验证查询创意组件共享信息请求参数
func (p *ComponentSharingGetReq) Validate() error {
	if p.OrganizationID == 0 {
		return errors.New("organization_id为必填")
	}
	if p.Page < MinPage || p.Page > MaxComponentSharingPage {
		return errors.New("page必须在1-99999之间")
	}
	if p.PageSize < MinPageSize || p.PageSize > MaxPageSize {
		return errors.New("page_size必须在1-100之间")
	}
	return p.GlobalReq.Validate()
}

// ComponentSharingListItem 创意组件共享列表项
type ComponentSharingListItem struct {
	SharedAccountID   int64  `json:"shared_account_id"`   // 共享账号id
	SharedAccountType string `json:"shared_account_type"` // 共享账号类型
}

// ComponentSharingGetResp 查询创意组件共享信息响应
type ComponentSharingGetResp struct {
	List []*ComponentSharingListItem `json:"list,omitempty"` // 共享信息列表
	PageInfoContainer
}

// ========== 获取创意组件详情 ==========
// https://developers.e.qq.com/v3.0/docs/api/component_detail/get

// 分页限制常量（组件详情查询接口特有）
const (
	MaxComponentDetailPage     = 99999 // page 最大值
	MaxComponentDetailPageSize = 200   // page_size 最大值
)

// 常量定义 - 组件详情过滤字段
const (
	ComponentDetailFilterFieldComponentID       = "component_id"
	ComponentDetailFilterFieldComponentType     = "component_type"
	ComponentDetailFilterFieldComponentSubType  = "component_sub_type"
	ComponentDetailFilterFieldElementSpecID     = "element_spec_id"
	ComponentDetailFilterFieldDynamicCreativeID = "dynamic_creative_id"
	ComponentDetailFilterFieldDataModelVersion  = "data_model_version"
	ComponentDetailFilterFieldCustomName        = "component_custom_name"
	ComponentDetailFilterFieldGenerationType    = "generation_type"
)

// 常量定义 - 营销目标
const (
	MarketingGoalAppPromotion   = "MARKETING_GOAL_APP_PROMOTION"   // 推广应用
	MarketingGoalWebsiteVisits  = "MARKETING_GOAL_WEBSITE_VISITS"  // 推广网站
	MarketingGoalSalesLead      = "MARKETING_GOAL_SALES_LEAD"      // 销售线索收集
	MarketingGoalLocalAds       = "MARKETING_GOAL_LOCAL_ADS"       // 推广本地店铺
	MarketingGoalEcommerce      = "MARKETING_GOAL_ECOMMERCE"       // 商品销售
	MarketingGoalBrandAwareness = "MARKETING_GOAL_BRAND_AWARENESS" // 品牌曝光
	MarketingGoalVideoViews     = "MARKETING_GOAL_VIDEO_VIEWS"     // 视频推广
	MarketingGoalLiveStreaming  = "MARKETING_GOAL_LIVE_STREAMING"  // 直播推广
	MarketingGoalChannelFans    = "MARKETING_GOAL_CHANNEL_FANS"    // 视频号涨粉
	MarketingGoalBrandInteract  = "MARKETING_GOAL_BRAND_INTERACT"  // 品牌互动
	MarketingGoalSearchBrand    = "MARKETING_GOAL_SEARCH_BRAND"    // 搜索品牌专区
	MarketingGoalShakeAd        = "MARKETING_GOAL_SHAKE_AD"        // 摇一摇广告
)

// 常量定义 - 营销载体类型（部分已在adgroup.go中定义）
const (
	MarketingCarrierTypeApp           = "MARKETING_CARRIER_TYPE_APP"             // 移动应用
	MarketingCarrierTypeMinProgram    = "MARKETING_CARRIER_TYPE_MINI_PROGRAM"    // 小程序
	MarketingCarrierTypeWechatCanvas  = "MARKETING_CARRIER_TYPE_WECHAT_CANVAS"   // 微信原生页
	MarketingCarrierTypeH5            = "MARKETING_CARRIER_TYPE_H5"              // H5
	MarketingCarrierTypeQqApp         = "MARKETING_CARRIER_TYPE_QQ_APP"          // QQ应用
	MarketingCarrierTypeQqMiniProgram = "MARKETING_CARRIER_TYPE_QQ_MINI_PROGRAM" // QQ小程序
	MarketingCarrierTypeQqMiniGame    = "MARKETING_CARRIER_TYPE_QQ_MINI_GAME"    // QQ小游戏
)

// 常量定义 - 营销诉求类型
const (
	MarketingTargetTypeAppDownload        = "MARKETING_TARGET_TYPE_APP_DOWNLOAD"         // 应用下载
	MarketingTargetTypeVisitWechatCanvas  = "MARKETING_TARGET_TYPE_VISIT_WECHAT_CANVAS"  // 访问微信原生页
	MarketingTargetTypeVisitMiniProgram   = "MARKETING_TARGET_TYPE_VISIT_MINI_PROGRAM"   // 访问小程序
	MarketingTargetTypeOnlineConsult      = "MARKETING_TARGET_TYPE_ONLINE_CONSULT"       // 在线咨询
	MarketingTargetTypeFormConsult        = "MARKETING_TARGET_TYPE_FORM_CONSULT"         // 表单咨询
	MarketingTargetTypeWechatMiniGame     = "MARKETING_TARGET_TYPE_WECHAT_MINI_GAME"     // 微信小游戏
	MarketingTargetTypeVisitH5            = "MARKETING_TARGET_TYPE_VISIT_H5"             // 访问H5
	MarketingTargetTypeEcommerce          = "MARKETING_TARGET_TYPE_ECOMMERCE"            // 电商
	MarketingTargetTypeVisitStore         = "MARKETING_TARGET_TYPE_VISIT_STORE"          // 访问门店
	MarketingTargetTypeWechatChannelsLive = "MARKETING_TARGET_TYPE_WECHAT_CHANNELS_LIVE" // 视频号直播
)

// 常量定义 - 站点集合（部分已在adgroup.go中定义）
const (
	SiteSetQQ          = "SITE_SET_QQ"           // QQ
	SiteSetQZone       = "SITE_SET_QZONE"        // QQ空间
	SiteSetOceanEngine = "SITE_SET_OCEAN_ENGINE" // 优量汇
	SiteSetOther       = "SITE_SET_OTHER"        // 其他
)

// OptimizationGoalStruct 优化目标结构
type OptimizationGoalStruct struct {
	OptimizationGoal   string `json:"optimization_goal,omitempty"`    // 优化目标
	BidObjective       string `json:"bid_objective,omitempty"`        // 出价目标
	DeepConversionType string `json:"deep_conversion_type,omitempty"` // 深度优化类型
	DeepConversionSpec string `json:"deep_conversion_spec,omitempty"` // 深度优化规格
}

// AdContext 广告上下文
type AdContext struct {
	MarketingGoal           string                   `json:"marketing_goal,omitempty"`             // 营销目标
	MarketingCarrierType    string                   `json:"marketing_carrier_type,omitempty"`     // 营销载体类型
	MarketingTargetType     string                   `json:"marketing_target_type,omitempty"`      // 营销诉求类型
	SiteSet                 []string                 `json:"site_set,omitempty"`                   // 站点集合
	CreativeTemplateID      int64                    `json:"creative_template_id,omitempty"`       // 创意模版id
	MarketingCarrierDetail  *MarketingCarrierDetail  `json:"marketing_carrier_detail,omitempty"`   // 营销载体详情
	OptimizationGoalStruct  *OptimizationGoalStruct  `json:"optimization_goal_struct,omitempty"`   // 优化目标结构
	MpaSpec                 *MpaSpec                 `json:"mpa_spec,omitempty"`                   // MPA规格
	MarketingAssetOuterSpec *MarketingAssetOuterSpec `json:"marketing_asset_outer_spec,omitempty"` // 营销资产外层规格
}

// ComponentDetailQueryFilter 组件详情过滤条件
type ComponentDetailQueryFilter struct {
	Field    string      `json:"field"`    // 过滤字段
	Operator string      `json:"operator"` // 操作符
	Values   interface{} `json:"values"`   // 过滤值
}

// OfficialDetail 官方页面详情
type OfficialDetail struct {
	PageID                      int64  `json:"page_id,omitempty"`                         // 页面id
	PageName                    string `json:"page_name,omitempty"`                       // 页面名称
	PlayableType                string `json:"playable_type,omitempty"`                   // 可玩类型
	PreviewURL                  string `json:"preview_url,omitempty"`                     // 预览链接
	PageStatus                  string `json:"page_status,omitempty"`                     // 页面状态
	WechatChannelsLiveReserveID int64  `json:"wechat_channels_live_reserve_id,omitempty"` // 视频号预约直播id
	DisableCode                 string `json:"disable_code,omitempty"`                    // 禁用原因码
	DisableMessage              string `json:"disable_message,omitempty"`                 // 禁用原因
	QuoteCreativeMaterial       bool   `json:"quote_creative_material,omitempty"`         // 是否引用创意素材
}

// WechatMiniProgramPageDetail 微信小程序页面详情
type WechatMiniProgramPageDetail struct {
	MiniProgramID       int64    `json:"mini_program_id,omitempty"`        // 小程序id
	MiniProgramNickName string   `json:"mini_program_nick_name,omitempty"` // 小程序名称
	MiniProgramIconURL  string   `json:"mini_program_icon_url,omitempty"`  // 小程序图标url
	MiniProgramPath     string   `json:"mini_program_path,omitempty"`      // 小程序路径
	MiniProgramPaths    []string `json:"mini_program_paths,omitempty"`     // 小程序路径列表
}

// WechatMiniGamePageDetail 微信小游戏页面详情
type WechatMiniGamePageDetail struct {
	MiniGameID                int64  `json:"mini_game_id,omitempty"`                 // 小游戏id
	MiniGameNickName          string `json:"mini_game_nick_name,omitempty"`          // 小游戏名称
	MiniGameIconURL           string `json:"mini_game_icon_url,omitempty"`           // 小游戏图标url
	MiniGameTrackingParameter string `json:"mini_game_tracking_parameter,omitempty"` // 小游戏追踪参数
}

// H5Detail H5页面详情
type H5Detail struct {
	PageURL string `json:"page_url,omitempty"` // 页面url
}

// AndroidAppDetail Android应用详情
type AndroidAppDetail struct {
	AndroidAppID              int64  `json:"android_app_id,omitempty"`               // Android应用id
	AndroidChannelPackageID   int64  `json:"android_channel_package_id,omitempty"`   // Android渠道包id
	AndroidAppName            string `json:"android_app_name,omitempty"`             // Android应用名称
	AndroidChannelPackageName string `json:"android_channel_package_name,omitempty"` // Android渠道包名称
}

// IosAppDetail iOS应用详情
type IosAppDetail struct {
	IosAppID   int64  `json:"ios_app_id,omitempty"`   // iOS应用id
	IosAppName string `json:"ios_app_name,omitempty"` // iOS应用名称
}

// QqMiniProgramPageDetail QQ小程序页面详情
type QqMiniProgramPageDetail struct {
	MiniProgramID       int64  `json:"mini_program_id,omitempty"`        // 小程序id
	MiniProgramNickName string `json:"mini_program_nick_name,omitempty"` // 小程序名称
	MiniProgramIconURL  string `json:"mini_program_icon_url,omitempty"`  // 小程序图标url
	MiniProgramPath     string `json:"mini_program_path,omitempty"`      // 小程序路径
}

// QqMiniGamePageDetail QQ小游戏页面详情
type QqMiniGamePageDetail struct {
	MiniGameID       int64  `json:"mini_game_id,omitempty"`        // 小游戏id
	MiniGameNickName string `json:"mini_game_nick_name,omitempty"` // 小游戏名称
	MiniGameIconURL  string `json:"mini_game_icon_url,omitempty"`  // 小游戏图标url
}

// WechatOfficialAccountDetailDetail 微信公众号详情
type WechatOfficialAccountDetailDetail struct {
	OfficialAccountID   int64  `json:"official_account_id,omitempty"`   // 公众号id
	OfficialAccountName string `json:"official_account_name,omitempty"` // 公众号名称
}

// ChannelsShopProductDetail 视频号商品详情
type ChannelsShopProductDetail struct {
	ProductID   int64  `json:"product_id,omitempty"`   // 商品id
	ProductName string `json:"product_name,omitempty"` // 商品名称
}

// ChannelsReserveLiveDetail 视频号预约直播详情
type ChannelsReserveLiveDetail struct {
	ReserveLiveID   int64  `json:"reserve_live_id,omitempty"`   // 预约直播id
	ReserveLiveName string `json:"reserve_live_name,omitempty"` // 预约直播名称
}

// ChannelsWatchLiveDetail 视频号看直播详情
type ChannelsWatchLiveDetail struct {
	LiveID   int64  `json:"live_id,omitempty"`   // 直播id
	LiveName string `json:"live_name,omitempty"` // 直播名称
}

// WecomConsultPageDetail 企业微信咨询页面详情
type WecomConsultPageDetail struct {
	WecomConsultID   int64  `json:"wecom_consult_id,omitempty"`   // 企业微信咨询id
	WecomConsultName string `json:"wecom_consult_name,omitempty"` // 企业微信咨询名称
}

// WechatChannelsFeedPageDetail 微信视频号信息流页面详情
type WechatChannelsFeedPageDetail struct {
	ChannelsID   int64  `json:"channels_id,omitempty"`   // 视频号id
	ChannelsName string `json:"channels_name,omitempty"` // 视频号名称
}

// H5ProfilePageDetail H5个人主页详情
type H5ProfilePageDetail struct {
	ProfileID   int64  `json:"profile_id,omitempty"`   // 主页id
	ProfileName string `json:"profile_name,omitempty"` // 主页名称
}

// ComponentDetailPageDetail 组件详情页面详情（包含15种子类型）
type ComponentDetailPageDetail struct {
	OfficialDetail                    *OfficialDetail                    `json:"official_detail,omitempty"`                       // 官方页面详情
	WechatMiniProgramPageDetail       *WechatMiniProgramPageDetail       `json:"wechat_mini_program_page_detail,omitempty"`       // 微信小程序页面详情
	WechatMiniGamePageDetail          *WechatMiniGamePageDetail          `json:"wechat_mini_game_page_detail,omitempty"`          // 微信小游戏页面详情
	H5Detail                          *H5Detail                          `json:"h5_detail,omitempty"`                             // H5页面详情
	AndroidAppDetail                  *AndroidAppDetail                  `json:"android_app_detail,omitempty"`                    // Android应用详情
	IosAppDetail                      *IosAppDetail                      `json:"ios_app_detail,omitempty"`                        // iOS应用详情
	QqMiniProgramPageDetail           *QqMiniProgramPageDetail           `json:"qq_mini_program_page_detail,omitempty"`           // QQ小程序页面详情
	QqMiniGamePageDetail              *QqMiniGamePageDetail              `json:"qq_mini_game_page_detail,omitempty"`              // QQ小游戏页面详情
	WechatOfficialAccountDetailDetail *WechatOfficialAccountDetailDetail `json:"wechat_official_account_detail_detail,omitempty"` // 微信公众号详情
	ChannelsShopProductDetail         *ChannelsShopProductDetail         `json:"channels_shop_product_detail,omitempty"`          // 视频号商品详情
	ChannelsReserveLiveDetail         *ChannelsReserveLiveDetail         `json:"channels_reserve_live_detail,omitempty"`          // 视频号预约直播详情
	ChannelsWatchLiveDetail           *ChannelsWatchLiveDetail           `json:"channels_watch_live_detail,omitempty"`            // 视频号看直播详情
	WecomConsultPageDetail            *WecomConsultPageDetail            `json:"wecom_consult_page_detail,omitempty"`             // 企业微信咨询详情
	WechatChannelsFeedPageDetail      *WechatChannelsFeedPageDetail      `json:"wechat_channels_feed_page_detail,omitempty"`      // 微信视频号信息流详情
	H5ProfilePageDetail               *H5ProfilePageDetail               `json:"h5_profile_page_detail,omitempty"`                // H5个人主页详情
	IsBackup                          bool                               `json:"is_backup,omitempty"`                             // 是否备用
	BackupIndex                       int                                `json:"backup_index,omitempty"`                          // 备用索引
	DisableCode                       string                             `json:"disable_code,omitempty"`                          // 禁用原因码
	DisableMessage                    string                             `json:"disable_message,omitempty"`                       // 禁用原因
}

// ComponentDetailJumpInfoItem 组件详情跳转信息项
type ComponentDetailJumpInfoItem struct {
	PageType   string                     `json:"page_type,omitempty"`   // 页面类型
	PageDetail *ComponentDetailPageDetail `json:"page_detail,omitempty"` // 页面详情
}

// ComponentDetailImageItem 组件详情图片项
type ComponentDetailImageItem struct {
	ImageID  int64       `json:"image_id,omitempty"`  // 图片id
	ImageURL string      `json:"image_url,omitempty"` // 图片url
	JumpInfo interface{} `json:"jump_info,omitempty"` // 跳转信息
}

// ComponentDetailVideoItem 组件详情视频项
type ComponentDetailVideoItem struct {
	VideoID  int64  `json:"video_id,omitempty"`  // 视频id
	VideoURL string `json:"video_url,omitempty"` // 视频url
	CoverID  int64  `json:"cover_id,omitempty"`  // 封面id
	CoverURL string `json:"cover_url,omitempty"` // 封面url
}

// ComponentDetailBrand 组件详情品牌信息
type ComponentDetailBrand struct {
	BrandName    string      `json:"brand_name,omitempty"`     // 品牌名称
	BrandImageID int64       `json:"brand_image_id,omitempty"` // 品牌图片id
	JumpInfo     interface{} `json:"jump_info,omitempty"`      // 跳转信息
}

// ComponentDetail 组件详情
type ComponentDetail struct {
	JumpInfo  *ComponentDetailJumpInfo    `json:"jump_info,omitempty"`  // 跳转信息
	ImageList []*ComponentDetailImageItem `json:"image_list,omitempty"` // 图片列表
	VideoList []*ComponentDetailVideoItem `json:"video_list,omitempty"` // 视频列表
	Brand     *ComponentDetailBrand       `json:"brand,omitempty"`      // 品牌信息
}

// ComponentDetailJumpInfo 组件详情跳转信息
type ComponentDetailJumpInfo struct {
	JumpInfoList []*ComponentDetailJumpInfoItem `json:"jump_info_list,omitempty"` // 跳转信息列表
}

// ComponentDetailListItem 创意组件详情列表项
type ComponentDetailListItem struct {
	AccountID       int64            `json:"account_id,omitempty"`       // 广告主id
	OrganizationID  int64            `json:"organization_id,omitempty"`  // 业务单元id
	ComponentID     int64            `json:"component_id,omitempty"`     // 组件id
	ComponentDetail *ComponentDetail `json:"component_detail,omitempty"` // 组件详情
}

// ComponentDetailGetReq 获取创意组件详情请求
// https://developers.e.qq.com/v3.0/docs/api/component_detail/get
type ComponentDetailGetReq struct {
	GlobalReq
	AccountID      int64                         `json:"account_id,omitempty"`      // 广告主id
	OrganizationID int64                         `json:"organization_id,omitempty"` // 业务单元id
	Filtering      []*ComponentDetailQueryFilter `json:"filtering,omitempty"`       // 过滤条件，最多4个
	Page           int                           `json:"page,omitempty"`            // 搜索页码，默认1，最大99999
	PageSize       int                           `json:"page_size,omitempty"`       // 每页条数，默认10，最大200
	AdContext      *AdContext                    `json:"ad_context,omitempty"`      // 广告上下文
}

func (p *ComponentDetailGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page <= 0 {
		p.Page = DefaultComponentPage
	}
	if p.PageSize <= 0 {
		p.PageSize = DefaultComponentPageSize
	}
}

// Validate 验证获取创意组件详情请求参数
func (p *ComponentDetailGetReq) Validate() error {
	if p.AccountID == 0 && p.OrganizationID == 0 {
		return errors.New("account_id和organization_id至少填写一个")
	}
	if len(p.Filtering) > 4 {
		return errors.New("filtering最多4个过滤条件")
	}
	if p.Page < MinPage || p.Page > MaxComponentDetailPage {
		return errors.New("page必须在1-99999之间")
	}
	if p.PageSize < MinPageSize || p.PageSize > MaxComponentDetailPageSize {
		return errors.New("page_size必须在1-200之间")
	}
	return p.GlobalReq.Validate()
}

// ComponentDetailGetResp 获取创意组件详情响应
type ComponentDetailGetResp struct {
	List []*ComponentDetailListItem `json:"list,omitempty"` // 组件详情列表
	PageInfoContainer
}
