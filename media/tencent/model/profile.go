package model

import "errors"

// ========== 获取朋友圈头像昵称跳转页 ==========
// https://developers.e.qq.com/v3.0/docs/api/profiles/get

// 朋友圈头像昵称跳转页过滤字段常量
const (
	ProfileFilterFieldProfileType           = "profile_type"
	ProfileFilterFieldProfileID             = "profile_id"
	ProfileFilterFieldSupportWechatChannels = "support_wechat_channels"
	ProfileFilterFieldMarketingGoal         = "marketing_goal"
	ProfileFilterFieldMarketingSubGoal      = "marketing_sub_goal"
	ProfileFilterFieldMarketingCarrierType  = "marketing_carrier_type"
	ProfileFilterFieldMarketingTargetType   = "marketing_target_type"
	ProfileFilterFieldMarketingCarrierID    = "marketing_carrier_id"
)

// 朋友圈头像昵称跳转页类型枚举
const (
	ProfileTypeDefinition   = "PROFILE_TYPE_DEFINITION"    // 自定义类型
	ProfileTypeAutoGenerate = "PROFILE_TYPE_AUTO_GENERATE" // 自动填充类型
)

// 分页常量
const (
	MinProfileGetPage         = 1     // page 最小值
	MaxProfileGetPage         = 99999 // page 最大值
	MinProfileGetPageSize     = 1     // page_size 最小值
	MaxProfileGetPageSize     = 100   // page_size 最大值
	DefaultProfileGetPage     = 1     // page 默认值
	DefaultProfileGetPageSize = 10    // page_size 默认值

	MaxProfileGetFilteringCount = 4 // filtering 最大长度
)

// ProfileFilteringItem 朋友圈头像昵称跳转页过滤条件
type ProfileFilteringItem struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)
	Operator string   `json:"operator"` // 操作符 (必填)
	Values   []string `json:"values"`   // 字段取值 (必填)
}

// Validate 验证单个过滤条件
func (f *ProfileFilteringItem) Validate() error {
	if f.Field == "" {
		return errors.New("field为必填")
	}
	if f.Operator == "" {
		return errors.New("operator为必填")
	}
	if len(f.Values) == 0 {
		return errors.New("values为必填，至少包含1个值")
	}
	return nil
}

// ProfileGetReq 获取朋友圈头像昵称跳转页请求
// https://developers.e.qq.com/v3.0/docs/api/profiles/get
type ProfileGetReq struct {
	GlobalReq
	AccountID      int64                   `json:"account_id,omitempty"`      // 广告主帐号 id (必填)
	Filtering      []*ProfileFilteringItem `json:"filtering,omitempty"`       // 过滤条件，最大4条
	Page           int                     `json:"page,omitempty"`            // 搜索页码，1-99999，默认1
	PageSize       int                     `json:"page_size,omitempty"`       // 每页条数，1-100，默认10
	OrganizationID int64                   `json:"organization_id,omitempty"` // 业务单元 id，0-9999999999
}

func (p *ProfileGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page == 0 {
		p.Page = DefaultProfileGetPage
	}
	if p.PageSize == 0 {
		p.PageSize = DefaultProfileGetPageSize
	}
}

// Validate 验证获取朋友圈头像昵称跳转页请求参数
func (p *ProfileGetReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if len(p.Filtering) > MaxProfileGetFilteringCount {
		return errors.New("filtering数组长度不能超过4")
	}
	for i, f := range p.Filtering {
		if f == nil {
			return errors.New("filtering[" + itoa(i) + "]不能为空")
		}
		if err := f.Validate(); err != nil {
			return errors.New("filtering[" + itoa(i) + "]: " + err.Error())
		}
	}
	if p.Page < MinProfileGetPage || p.Page > MaxProfileGetPage {
		return errors.New("page须在1-99999之间")
	}
	if p.PageSize < MinProfileGetPageSize || p.PageSize > MaxProfileGetPageSize {
		return errors.New("page_size须在1-100之间")
	}
	if p.OrganizationID < 0 || p.OrganizationID > 9999999999 {
		return errors.New("organization_id须在0-9999999999之间")
	}
	return p.GlobalReq.Validate()
}

// ProfileItem 朋友圈头像昵称跳转页列表项
type ProfileItem struct {
	OwnerID          int64  `json:"owner_id"`           // 广告主帐号 id
	ProfileType      string `json:"profile_type"`       // 朋友圈头像及昵称跳转页类型
	ProfileID        int64  `json:"profile_id"`         // 朋友圈头像及昵称跳转页 id
	HeadImageID      string `json:"head_image_id"`      // 头像图片 id
	HeadImageURL     string `json:"head_image_url"`     // 头像 url
	ProfileName      string `json:"profile_name"`       // 昵称
	Description      string `json:"description"`        // 朋友圈头像及昵称跳转页简介
	CreatedTime      int64  `json:"created_time"`       // 创建时间，时间戳
	LastModifiedTime int64  `json:"last_modified_time"` // 最后修改时间，时间戳
	ProfileURL       string `json:"profile_url"`        // 朋友圈头像昵称跳转页 url
	SystemStatus     string `json:"system_status"`      // 朋友圈头像及昵称跳转页状态
	MdmID            int64  `json:"mdm_id"`             // 主体 id
}

// ProfileGetResp 获取朋友圈头像昵称跳转页响应
// https://developers.e.qq.com/v3.0/docs/api/profiles/get
type ProfileGetResp struct {
	List     []*ProfileItem `json:"list"`      // 返回信息列表
	PageInfo *PageInfo      `json:"page_info"` // 分页配置信息
}

// ========== 删除朋友圈头像昵称跳转页 ==========
// https://developers.e.qq.com/v3.0/docs/api/profiles/delete

// ProfileDeleteReq 删除朋友圈头像昵称跳转页请求（POST JSON）
// https://developers.e.qq.com/v3.0/docs/api/profiles/delete
// 注意：只能删除 profile_type 为 PROFILE_TYPE_DEFINITION 的跳转页
type ProfileDeleteReq struct {
	GlobalReq
	AccountID      int64 `json:"account_id,omitempty"`      // 广告主帐号 id
	ProfileID      int64 `json:"profile_id"`                // 朋友圈头像及昵称跳转页 id (必填)
	OrganizationID int64 `json:"organization_id,omitempty"` // 业务单元 id，0-9999999999
}

func (p *ProfileDeleteReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证删除朋友圈头像昵称跳转页请求参数
func (p *ProfileDeleteReq) Validate() error {
	if p.ProfileID == 0 {
		return errors.New("profile_id为必填")
	}
	if p.OrganizationID < 0 || p.OrganizationID > 9999999999 {
		return errors.New("organization_id须在0-9999999999之间")
	}
	return p.GlobalReq.Validate()
}

// ProfileDeleteResp 删除朋友圈头像昵称跳转页响应
// https://developers.e.qq.com/v3.0/docs/api/profiles/delete
type ProfileDeleteResp struct {
	ProfileID int64 `json:"profile_id"` // 朋友圈头像及昵称跳转页 id
}

// ========== 创建朋友圈头像昵称跳转页 ==========
// https://developers.e.qq.com/v3.0/docs/api/profiles/add

// 营销目的类型枚举
const (
	ProfileMarketingGoalUnknown                 = "MARKETING_GOAL_UNKNOWN"
	ProfileMarketingGoalUserGrowth              = "MARKETING_GOAL_USER_GROWTH"
	ProfileMarketingGoalProductSales            = "MARKETING_GOAL_PRODUCT_SALES"
	ProfileMarketingGoalLeadRetention           = "MARKETING_GOAL_LEAD_RETENTION"
	ProfileMarketingGoalBrandPromotion          = "MARKETING_GOAL_BRAND_PROMOTION"
	ProfileMarketingGoalIncreaseFansInteraction = "MARKETING_GOAL_INCREASE_FANS_INTERACTION"
)

// 营销载体类型枚举（仅 PROFILE_TYPE_AUTO_GENERATE 时有效）
const (
	ProfileMarketingCarrierTypeAppAndroid = "MARKETING_CARRIER_TYPE_APP_ANDROID"
	ProfileMarketingCarrierTypeAppIOS     = "MARKETING_CARRIER_TYPE_APP_IOS"
)

// 字段长度常量
const (
	MinProfileHeadImageIDBytes        = 1    // head_image_id 最小长度
	MaxProfileHeadImageIDBytes        = 64   // head_image_id 最大长度
	MinProfileNameBytes               = 1    // profile_name 最小长度
	MaxProfileNameBytes               = 30   // profile_name 最大长度（字节）
	MinProfileDescriptionBytes        = 1    // description 最小长度
	MaxProfileDescriptionBytes        = 240  // description 最大长度（字节）
	MaxProfileMarketingCarrierIDBytes = 2048 // marketing_carrier_id 最大长度
)

// ProfileAddReq 创建朋友圈头像昵称跳转页请求（POST JSON）
// https://developers.e.qq.com/v3.0/docs/api/profiles/add
// 使用说明：
//   - profile_type 为 PROFILE_TYPE_DEFINITION 时，必须传入 head_image_id、profile_name、description
//   - profile_type 为 PROFILE_TYPE_AUTO_GENERATE 时，marketing_carrier_type 仅支持 APP_ANDROID / APP_IOS
//   - 每个账户最多创建 5 个自定义类型跳转页
type ProfileAddReq struct {
	GlobalReq
	AccountID            int64  `json:"account_id,omitempty"`             // 广告主帐号 id
	MarketingGoal        string `json:"marketing_goal,omitempty"`         // 营销目的类型
	MarketingSubGoal     string `json:"marketing_sub_goal,omitempty"`     // 二级营销目的类型
	MarketingCarrierType string `json:"marketing_carrier_type,omitempty"` // 营销载体类型（仅 AUTO_GENERATE 时有效）
	MarketingTargetType  string `json:"marketing_target_type,omitempty"`  // 推广产品类型
	MarketingCarrierID   string `json:"marketing_carrier_id,omitempty"`   // 营销载体 id，0-2048 字节
	ProfileType          string `json:"profile_type"`                     // 朋友圈头像及昵称跳转页类型 (必填)
	HeadImageID          string `json:"head_image_id,omitempty"`          // 头像图片 id，DEFINITION 时必填，1-64 字节
	ProfileName          string `json:"profile_name,omitempty"`           // 昵称，DEFINITION 时必填，1-30 字节
	Description          string `json:"description,omitempty"`            // 简介，DEFINITION 时必填，1-240 字节
	OrganizationID       int64  `json:"organization_id,omitempty"`        // 业务单元 id，0-9999999999
	MdmID                int64  `json:"mdm_id,omitempty"`                 // 主体 id
}

func (p *ProfileAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证创建朋友圈头像昵称跳转页请求参数
func (p *ProfileAddReq) Validate() error {
	if p.ProfileType == "" {
		return errors.New("profile_type为必填")
	}
	if p.ProfileType != ProfileTypeDefinition && p.ProfileType != ProfileTypeAutoGenerate {
		return errors.New("profile_type只能为PROFILE_TYPE_DEFINITION或PROFILE_TYPE_AUTO_GENERATE")
	}
	if p.ProfileType == ProfileTypeDefinition {
		if p.HeadImageID == "" {
			return errors.New("profile_type为PROFILE_TYPE_DEFINITION时，head_image_id为必填")
		}
		if len(p.HeadImageID) < MinProfileHeadImageIDBytes || len(p.HeadImageID) > MaxProfileHeadImageIDBytes {
			return errors.New("head_image_id长度须在1-64字节之间")
		}
		if p.ProfileName == "" {
			return errors.New("profile_type为PROFILE_TYPE_DEFINITION时，profile_name为必填")
		}
		if len(p.ProfileName) < MinProfileNameBytes || len(p.ProfileName) > MaxProfileNameBytes {
			return errors.New("profile_name长度须在1-30字节之间")
		}
		if p.Description == "" {
			return errors.New("profile_type为PROFILE_TYPE_DEFINITION时，description为必填")
		}
		if len(p.Description) < MinProfileDescriptionBytes || len(p.Description) > MaxProfileDescriptionBytes {
			return errors.New("description长度须在1-240字节之间")
		}
	}
	if len(p.MarketingCarrierID) > MaxProfileMarketingCarrierIDBytes {
		return errors.New("marketing_carrier_id长度不能超过2048字节")
	}
	if p.OrganizationID < 0 || p.OrganizationID > 9999999999 {
		return errors.New("organization_id须在0-9999999999之间")
	}
	return p.GlobalReq.Validate()
}

// ProfileAddResp 创建朋友圈头像昵称跳转页响应
// https://developers.e.qq.com/v3.0/docs/api/profiles/add
type ProfileAddResp struct {
	ProfileID int64 `json:"profile_id"` // 朋友圈头像及昵称跳转页 id
}
