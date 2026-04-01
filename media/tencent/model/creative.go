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
