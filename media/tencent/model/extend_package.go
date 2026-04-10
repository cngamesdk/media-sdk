package model

import (
	"errors"
	"regexp"
)

// ========== 创建应用分包 ==========
// https://developers.e.qq.com/v3.0/docs/api/extend_package/add

// 字段限制常量
const (
	MinExtendPackageChannelListLen           = 1    // channel_list 最小长度
	MaxExtendPackageChannelListLen           = 200  // channel_list 最大长度
	MinExtendPackageChannelIDBytes           = 1    // channel_id 最小长度
	MaxExtendPackageChannelIDBytes           = 200  // channel_id 最大长度
	MinExtendPackageChannelNameBytes         = 1    // channel_name 最小长度（创建）
	MaxExtendPackageChannelNameBytes         = 255  // channel_name 最大长度（创建）
	MaxExtendPackageUpdateChannelNameBytes   = 1024 // channel_name 最大长度（更新）
	MinExtendPackageCustomizedChannelIDBytes = 1    // customized_channel_id 最小长度
	MaxExtendPackageCustomizedChannelIDBytes = 256  // customized_channel_id 最大长度
)

// channelIDRegex channel_id 合法字符正则：仅允许英文字母、数字和 _ . -
var channelIDRegex = regexp.MustCompile(`^[a-zA-Z0-9_.\-]+$`)

// ExtendPackageChannelItem 渠道号信息
type ExtendPackageChannelItem struct {
	ChannelID           string `json:"channel_id"`                      // 渠道标识 (必填)，仅英文/数字/_.-，1-200 字节
	ChannelName         string `json:"channel_name,omitempty"`          // 渠道包名称，1-255 字节，默认 {package_id}_{channel_id}
	CustomizedChannelID string `json:"customized_channel_id,omitempty"` // 自定义渠道包 id，1-256 字节，须与 channel_id 一致
}

// Validate 验证单条渠道号信息
func (c *ExtendPackageChannelItem) Validate() error {
	if c.ChannelID == "" {
		return errors.New("channel_id为必填")
	}
	if len(c.ChannelID) < MinExtendPackageChannelIDBytes || len(c.ChannelID) > MaxExtendPackageChannelIDBytes {
		return errors.New("channel_id长度须在1-200字节之间")
	}
	if !channelIDRegex.MatchString(c.ChannelID) {
		return errors.New("channel_id只能由英文字母、数字和_.-组成")
	}
	if c.ChannelName != "" {
		if len(c.ChannelName) < MinExtendPackageChannelNameBytes || len(c.ChannelName) > MaxExtendPackageChannelNameBytes {
			return errors.New("channel_name长度须在1-255字节之间")
		}
	}
	if c.CustomizedChannelID != "" {
		if len(c.CustomizedChannelID) < MinExtendPackageCustomizedChannelIDBytes || len(c.CustomizedChannelID) > MaxExtendPackageCustomizedChannelIDBytes {
			return errors.New("customized_channel_id长度须在1-256字节之间")
		}
	}
	return nil
}

// ExtendPackageAddReq 创建应用分包请求（POST JSON）
// https://developers.e.qq.com/v3.0/docs/api/extend_package/add
type ExtendPackageAddReq struct {
	GlobalReq
	AccountID   int64                       `json:"account_id"`   // 广告主帐号 id (必填)
	PackageID   int64                       `json:"package_id"`   // Android 应用 id (必填)，≥0 且 <2^63
	ChannelList []*ExtendPackageChannelItem `json:"channel_list"` // 渠道号信息 (必填)，1-200 条
}

func (p *ExtendPackageAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证创建应用分包请求参数
func (p *ExtendPackageAddReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.PackageID < 0 {
		return errors.New("package_id须大于等于0")
	}
	if len(p.ChannelList) < MinExtendPackageChannelListLen || len(p.ChannelList) > MaxExtendPackageChannelListLen {
		return errors.New("channel_list数组长度须在1-200之间")
	}
	for i, ch := range p.ChannelList {
		if ch == nil {
			return errors.New("channel_list[" + itoa(i) + "]不能为空")
		}
		if err := ch.Validate(); err != nil {
			return errors.New("channel_list[" + itoa(i) + "]: " + err.Error())
		}
	}
	return p.GlobalReq.Validate()
}

// ExtendPackageResultItem 渠道包操作结果（成功/失败列表共用）
type ExtendPackageResultItem struct {
	ChannelName      string `json:"channel_name"`       // 安卓应用渠道包名称
	ChannelPackageID string `json:"channel_package_id"` // 安卓应用渠道包 id
	Message          string `json:"message"`            // 渠道包操作结果描述
}

// ExtendPackageAddResp 创建应用分包响应
// https://developers.e.qq.com/v3.0/docs/api/extend_package/add
type ExtendPackageAddResp struct {
	PackageID      int64                      `json:"package_id"`      // Android 应用 id
	SuccessResults []*ExtendPackageResultItem `json:"success_results"` // 渠道包操作成功信息列表
	FailedResults  []*ExtendPackageResultItem `json:"failed_results"`  // 应用分包操作失败列表
}

// ========== 更新应用子包版本 ==========
// https://developers.e.qq.com/v3.0/docs/api/extend_package/update

// ExtendPackageUpdateChannelItem 更新渠道号信息
// 与创建的差异：channel_name 最大长度 1024 字节，无 customized_channel_id 字段
type ExtendPackageUpdateChannelItem struct {
	ChannelID   string `json:"channel_id"`             // 渠道标识 (必填)，仅英文/数字/_.-，1-200 字节
	ChannelName string `json:"channel_name,omitempty"` // 安卓应用渠道包名称，1-1024 字节
}

// Validate 验证单条更新渠道号信息
func (c *ExtendPackageUpdateChannelItem) Validate() error {
	if c.ChannelID == "" {
		return errors.New("channel_id为必填")
	}
	if len(c.ChannelID) < MinExtendPackageChannelIDBytes || len(c.ChannelID) > MaxExtendPackageChannelIDBytes {
		return errors.New("channel_id长度须在1-200字节之间")
	}
	if !channelIDRegex.MatchString(c.ChannelID) {
		return errors.New("channel_id只能由英文字母、数字和_.-组成")
	}
	if c.ChannelName != "" {
		if len(c.ChannelName) < MinExtendPackageChannelNameBytes || len(c.ChannelName) > MaxExtendPackageUpdateChannelNameBytes {
			return errors.New("channel_name长度须在1-1024字节之间")
		}
	}
	return nil
}

// ExtendPackageUpdateReq 更新应用子包版本请求（POST JSON）
// https://developers.e.qq.com/v3.0/docs/api/extend_package/update
type ExtendPackageUpdateReq struct {
	GlobalReq
	AccountID   int64                             `json:"account_id"`   // 广告主帐号 id (必填)
	PackageID   int64                             `json:"package_id"`   // Android 应用 id (必填)，≥0 且 <2^63
	ChannelList []*ExtendPackageUpdateChannelItem `json:"channel_list"` // 渠道号信息 (必填)，1-200 条
}

func (p *ExtendPackageUpdateReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证更新应用子包版本请求参数
func (p *ExtendPackageUpdateReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.PackageID < 0 {
		return errors.New("package_id须大于等于0")
	}
	if len(p.ChannelList) < MinExtendPackageChannelListLen || len(p.ChannelList) > MaxExtendPackageChannelListLen {
		return errors.New("channel_list数组长度须在1-200之间")
	}
	for i, ch := range p.ChannelList {
		if ch == nil {
			return errors.New("channel_list[" + itoa(i) + "]不能为空")
		}
		if err := ch.Validate(); err != nil {
			return errors.New("channel_list[" + itoa(i) + "]: " + err.Error())
		}
	}
	return p.GlobalReq.Validate()
}

// ExtendPackageUpdateResp 更新应用子包版本响应
// https://developers.e.qq.com/v3.0/docs/api/extend_package/update
type ExtendPackageUpdateResp struct {
	PackageID      int64                      `json:"package_id"`      // Android 应用 id
	SuccessResults []*ExtendPackageResultItem `json:"success_results"` // 渠道包操作成功信息列表
	FailedResults  []*ExtendPackageResultItem `json:"failed_results"`  // 应用分包操作失败列表
}

// ========== 查询应用分包列表 ==========
// https://developers.e.qq.com/v3.0/docs/api/extend_package/get

// 查询过滤字段常量
const (
	ExtendPackageGetFilterFieldChannelPackageID = "channel_package_id" // 渠道包 id，操作符 EQUALS
	ExtendPackageGetFilterFieldChannelName      = "channel_name"       // 渠道包名称，操作符 CONTAINS
)

// 查询操作符常量
const (
	ExtendPackageGetFilterOperatorEquals   = "EQUALS"   // 精确匹配
	ExtendPackageGetFilterOperatorContains = "CONTAINS" // 模糊匹配
)

// 分页常量
const (
	MinExtendPackageGetPage         = 1     // page 最小值
	MaxExtendPackageGetPage         = 99999 // page 最大值
	MinExtendPackageGetPageSize     = 1     // page_size 最小值
	MaxExtendPackageGetPageSize     = 100   // page_size 最大值
	DefaultExtendPackageGetPage     = 1     // page 默认值
	DefaultExtendPackageGetPageSize = 10    // page_size 默认值

	MinExtendPackageGetFilteringCount = 1 // filtering 最小长度
	MaxExtendPackageGetFilteringCount = 2 // filtering 最大长度

	MaxExtendPackageGetFilterValueBytes = 1024 // filtering.values 单个值最大字节数
)

// ExtendPackageGetFilteringItem 查询过滤条件
type ExtendPackageGetFilteringItem struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)
	Operator string   `json:"operator"` // 操作符 (必填)
	Values   []string `json:"values"`   // 字段取值 (必填)，数组长度为1
}

// Validate 验证单个过滤条件
func (f *ExtendPackageGetFilteringItem) Validate() error {
	if f.Field == "" {
		return errors.New("field为必填")
	}
	if f.Operator == "" {
		return errors.New("operator为必填")
	}
	if len(f.Values) == 0 {
		return errors.New("values为必填，至少包含1个值")
	}
	if len(f.Values) > 1 {
		return errors.New("values数组长度须为1")
	}
	if len(f.Values[0]) == 0 || len(f.Values[0]) > MaxExtendPackageGetFilterValueBytes {
		return errors.New("values[0]长度须在1-1024字节之间")
	}
	return nil
}

// ExtendPackageGetReq 查询应用分包列表请求（GET）
// https://developers.e.qq.com/v3.0/docs/api/extend_package/get
type ExtendPackageGetReq struct {
	GlobalReq
	AccountID int64                            `json:"account_id"`          // 推广帐号 id (必填)
	PackageID int64                            `json:"package_id"`          // Android 应用 id (必填)，≥0 且 <2^63
	Filtering []*ExtendPackageGetFilteringItem `json:"filtering"`           // 过滤条件，1-2条
	Page      int                              `json:"page,omitempty"`      // 搜索页码，1-99999，默认1
	PageSize  int                              `json:"page_size,omitempty"` // 每页条数，1-100，默认10
}

func (p *ExtendPackageGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page == 0 {
		p.Page = DefaultExtendPackageGetPage
	}
	if p.PageSize == 0 {
		p.PageSize = DefaultExtendPackageGetPageSize
	}
}

// Validate 验证查询应用分包列表请求参数
func (p *ExtendPackageGetReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.PackageID < 0 {
		return errors.New("package_id须大于等于0")
	}
	if len(p.Filtering) > MaxExtendPackageGetFilteringCount {
		return errors.New("filtering数组长度不能超过2")
	}
	for i, f := range p.Filtering {
		if f == nil {
			return errors.New("filtering[" + itoa(i) + "]不能为空")
		}
		if err := f.Validate(); err != nil {
			return errors.New("filtering[" + itoa(i) + "]: " + err.Error())
		}
	}
	if p.Page < MinExtendPackageGetPage || p.Page > MaxExtendPackageGetPage {
		return errors.New("page须在1-99999之间")
	}
	if p.PageSize < MinExtendPackageGetPageSize || p.PageSize > MaxExtendPackageGetPageSize {
		return errors.New("page_size须在1-100之间")
	}
	return p.GlobalReq.Validate()
}

// ExtendPackageGetItem 应用分包列表项
type ExtendPackageGetItem struct {
	PackageID           int64  `json:"package_id"`            // Android 应用 id
	ChannelName         string `json:"channel_name"`          // 安卓应用渠道包名称
	ChannelPackageID    string `json:"channel_package_id"`    // 安卓应用渠道包 id
	ChannelID           string `json:"channel_id"`            // 渠道标识
	SystemStatus        string `json:"system_status"`         // 渠道包状态
	CreatedTime         int64  `json:"created_time"`          // 创建时间，时间戳
	LastModifiedTime    int64  `json:"last_modified_time"`    // 最后修改时间，时间戳
	CustomizedChannelID string `json:"customized_channel_id"` // 自定义渠道包 id
}

// ExtendPackageGetResp 查询应用分包列表响应
// https://developers.e.qq.com/v3.0/docs/api/extend_package/get
type ExtendPackageGetResp struct {
	List     []*ExtendPackageGetItem `json:"list"`      // 返回信息列表
	PageInfo *PageInfo               `json:"page_info"` // 分页配置信息
}
