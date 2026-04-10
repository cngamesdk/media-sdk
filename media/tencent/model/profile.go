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
