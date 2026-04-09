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
