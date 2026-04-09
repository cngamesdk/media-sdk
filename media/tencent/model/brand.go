package model

import "errors"

// ========== 创建品牌形象 ==========
// https://developers.e.qq.com/v3.0/docs/api/brand/add

// 字段限制常量
const (
	MinBrandNameBytes      = 1          // name 最小长度（字节）
	MaxBrandNameBytes      = 100        // name 最大长度（字节）
	MaxBrandImageFileBytes = 400 * 1024 // brand_image_file 最大大小（400KB）
)

// BrandAddReq 创建品牌形象请求（multipart/form-data）
// https://developers.e.qq.com/v3.0/docs/api/brand/add
type BrandAddReq struct {
	GlobalReq
	AccountID          int64  // 推广账户 id (必填)
	Name               string // 品牌形象名字 (必填)，1-100 字节
	BrandImageFile     []byte // 品牌形象图片文件 (必填)，仅支持 512x512，≤400KB，格式：jpg/jpeg/png
	BrandImageFileName string // 图片文件名（含扩展名），用于 multipart 表单
}

func (p *BrandAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证创建品牌形象请求参数
func (p *BrandAddReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.Name == "" {
		return errors.New("name为必填")
	}
	if len(p.Name) < MinBrandNameBytes || len(p.Name) > MaxBrandNameBytes {
		return errors.New("name长度须在1-100字节之间")
	}
	if len(p.BrandImageFile) == 0 {
		return errors.New("brand_image_file为必填")
	}
	if len(p.BrandImageFile) > MaxBrandImageFileBytes {
		return errors.New("brand_image_file大小不能超过400KB")
	}
	if p.BrandImageFileName == "" {
		return errors.New("brand_image_file_name为必填，需包含文件扩展名")
	}
	return p.GlobalReq.Validate()
}

// BrandAddResp 创建品牌形象响应
// https://developers.e.qq.com/v3.0/docs/api/brand/add
type BrandAddResp struct {
	AccountID   int64  `json:"account_id"`   // 推广账户 id
	Name        string `json:"name"`         // 品牌形象名字
	ImageID     string `json:"image_id"`     // 品牌形象图片 id
	Width       int    `json:"width"`        // 宽
	Height      int    `json:"height"`       // 高
	ImageURL    string `json:"image_url"`    // 品牌形象 URL
	CreatedTime int64  `json:"created_time"` // 创建时间，时间戳
}

// ========== 获取品牌形象列表 ==========
// https://developers.e.qq.com/v3.0/docs/api/brand/get

// 分页常量
const (
	MinBrandGetPage         = 1     // page 最小值
	MaxBrandGetPage         = 99999 // page 最大值
	MinBrandGetPageSize     = 1     // page_size 最小值
	MaxBrandGetPageSize     = 100   // page_size 最大值
	DefaultBrandGetPage     = 1     // page 默认值
	DefaultBrandGetPageSize = 10    // page_size 默认值
)

// BrandGetReq 获取品牌形象列表请求
// https://developers.e.qq.com/v3.0/docs/api/brand/get
type BrandGetReq struct {
	GlobalReq
	AccountID int64 `json:"account_id"`          // 推广账户 id (必填)
	Page      int   `json:"page,omitempty"`      // 搜索页码，1-99999，默认 1
	PageSize  int   `json:"page_size,omitempty"` // 每页条数，1-100，默认 10
}

func (p *BrandGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page == 0 {
		p.Page = DefaultBrandGetPage
	}
	if p.PageSize == 0 {
		p.PageSize = DefaultBrandGetPageSize
	}
}

// Validate 验证获取品牌形象列表请求参数
func (p *BrandGetReq) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.Page < MinBrandGetPage || p.Page > MaxBrandGetPage {
		return errors.New("page须在1-99999之间")
	}
	if p.PageSize < MinBrandGetPageSize || p.PageSize > MaxBrandGetPageSize {
		return errors.New("page_size须在1-100之间")
	}
	return p.GlobalReq.Validate()
}

// BrandItem 品牌形象列表项
type BrandItem struct {
	AccountID   int64  `json:"account_id"`   // 推广账户 id
	Name        string `json:"name"`         // 品牌形象名字
	ImageID     string `json:"image_id"`     // 品牌形象图片 id
	Width       int    `json:"width"`        // 宽
	Height      int    `json:"height"`       // 高
	ImageURL    string `json:"image_url"`    // 品牌形象 URL
	CreatedTime int64  `json:"created_time"` // 创建时间，时间戳
}

// BrandGetResp 获取品牌形象列表响应
// https://developers.e.qq.com/v3.0/docs/api/brand/get
type BrandGetResp struct {
	List     []*BrandItem `json:"list"`      // 返回信息列表
	PageInfo *PageInfo    `json:"page_info"` // 分页配置信息
}
