package model

import "errors"

// ========== 获取图片信息 ==========
// https://developers.e.qq.com/v3.0/docs/api/images/get

// 图片过滤字段常量
const (
	ImageFilterFieldImageSignature         = "image_signature"
	ImageFilterFieldImageID                = "image_id"
	ImageFilterFieldImageWidth             = "image_width"
	ImageFilterFieldImageHeight            = "image_height"
	ImageFilterFieldCreatedTime            = "created_time"
	ImageFilterFieldLastModifiedTime       = "last_modified_time"
	ImageFilterFieldSourceType             = "source_type"
	ImageFilterFieldProductCatalogID       = "product_catalog_id"
	ImageFilterFieldProductOuterID         = "product_outer_id"
	ImageFilterFieldOwnerAccountID         = "owner_account_id"
	ImageFilterFieldStatus                 = "status"
	ImageFilterFieldImageDescription       = "image_description"
	ImageFilterFieldSampleAspectRatio      = "sample_aspect_ratio"
	ImageFilterFieldFirstPublicationStatus = "first_publication_status"
	ImageFilterFieldQualityStatus          = "quality_status"
	ImageFilterFieldSimilarityStatus       = "similarity_status"
	ImageFilterFieldAigcFlag               = "aigc_flag"
	ImageFilterFieldFileSize               = "file_size"
	ImageFilterFieldHeight                 = "height"
	ImageFilterFieldWidth                  = "width"
	ImageFilterFieldRatio                  = "ratio"
)

// 图片来源类型枚举
const (
	ImageSourceTypeLocal          = "SOURCE_TYPE_LOCAL"
	ImageSourceTypeMuse           = "SOURCE_TYPE_MUSE"
	ImageSourceTypeAPI            = "SOURCE_TYPE_API"
	ImageSourceTypeQuickDraw      = "SOURCE_TYPE_QUICK_DRAW"
	ImageSourceTypeVideoSnapshots = "SOURCE_TYPE_VIDEO_SNAPSHOTS"
	ImageSourceTypeTCC            = "SOURCE_TYPE_TCC"
)

// 图片状态枚举
const (
	ImageStatusNormal  = "ADSTATUS_NORMAL"
	ImageStatusDeleted = "ADSTATUS_DELETED"
)

// 图片首发状态枚举
const (
	ImageFirstPublicationStatusDefault          = "FIRST_PUBLICATION_STATUS_DEFAULT"
	ImageFirstPublicationStatusFirstPublication = "FIRST_PUBLICATION_STATUS_FIRST_PUBLICATION"
)

// 图片质量状态枚举
const (
	ImageQualityStatusDefault    = "QUALITY_STATUS_DEFAULT"
	ImageQualityStatusLowQuality = "QUALITY_STATUS_LOW_QUALITY"
)

// 图片 AIGC 标记枚举
const (
	ImageAigcFlagUnknown     = "AIGC_FLAG_UNKNOWN"
	ImageAigcFlagNotAI       = "AIGC_FLAG_NOT_AI"
	ImageAigcFlagUseMuseAI   = "AIGC_FLAG_USE_MUSE_AI"
	ImageAigcFlagUseOthersAI = "AIGC_FLAG_USE_OTHERS_AI"
)

// 字段限制常量
const (
	MaxImageGetFilteringCount = 4     // filtering 最大长度
	MinImageGetPage           = 1     // page 最小值
	MaxImageGetPage           = 99999 // page 最大值
	MinImageGetPageSize       = 1     // page_size 最小值
	MaxImageGetPageSize       = 100   // page_size 最大值
	DefaultImageGetPage       = 1     // page 默认值
	DefaultImageGetPageSize   = 10    // page_size 默认值
)

// ImageFilteringItem 图片过滤条件
type ImageFilteringItem struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)
	Operator string   `json:"operator"` // 操作符 (必填)
	Values   []string `json:"values"`   // 字段取值 (必填)
}

// Validate 验证单个过滤条件
func (f *ImageFilteringItem) Validate() error {
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

// ImageGetReq 获取图片信息请求
// https://developers.e.qq.com/v3.0/docs/api/images/get
type ImageGetReq struct {
	GlobalReq
	AccountID        int64                 `json:"account_id,omitempty"`        // 广告主账户 id，与 organization_id 必填其一
	OrganizationID   int64                 `json:"organization_id,omitempty"`   // 业务单元 id，与 account_id 必填其一
	Filtering        []*ImageFilteringItem `json:"filtering,omitempty"`         // 过滤条件，最大4条
	Page             int                   `json:"page,omitempty"`              // 搜索页码，1-99999，默认1
	PageSize         int                   `json:"page_size,omitempty"`         // 每页条数，1-100，默认10
	LabelID          int64                 `json:"label_id,omitempty"`          // 标签 id
	BusinessScenario int                   `json:"business_scenario,omitempty"` // 业务场景：1=内容素材包，2=投放素材包
}

func (p *ImageGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page == 0 {
		p.Page = DefaultImageGetPage
	}
	if p.PageSize == 0 {
		p.PageSize = DefaultImageGetPageSize
	}
}

// Validate 验证获取图片信息请求参数
func (p *ImageGetReq) Validate() error {
	if p.AccountID == 0 && p.OrganizationID == 0 {
		return errors.New("account_id 和 organization_id 需必填其一")
	}
	if len(p.Filtering) > MaxImageGetFilteringCount {
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
	if p.Page < MinImageGetPage || p.Page > MaxImageGetPage {
		return errors.New("page须在1-99999之间")
	}
	if p.PageSize < MinImageGetPageSize || p.PageSize > MaxImageGetPageSize {
		return errors.New("page_size须在1-100之间")
	}
	return p.GlobalReq.Validate()
}

// ImageAssetItem 图片信息列表项
type ImageAssetItem struct {
	ImageID           string `json:"image_id"`            // 图片 id
	Width             int    `json:"width"`               // 图片宽度，单位 px
	Height            int    `json:"height"`              // 图片高度，单位 px
	FileSize          int64  `json:"file_size"`           // 图片大小，单位 B(byte)
	Type              string `json:"type"`                // 图片类型
	Signature         string `json:"signature"`           // 图片文件签名（md5值）
	Description       string `json:"description"`         // 图片文件描述
	SourceSignature   string `json:"source_signature"`    // 图片源文件签名，裁剪前源文件的 md5 值
	PreviewURL        string `json:"preview_url"`         // 预览地址
	SourceType        string `json:"source_type"`         // 图片来源
	ImageUsage        string `json:"image_usage"`         // 图片用途
	CreatedTime       int64  `json:"created_time"`        // 创建时间，时间戳
	LastModifiedTime  int64  `json:"last_modified_time"`  // 最后修改时间，时间戳
	ProductCatalogID  int64  `json:"product_catalog_id"`  // 商品库 id
	ProductOuterID    string `json:"product_outer_id"`    // 商品 id
	SourceReferenceID string `json:"source_reference_id"` // 素材来源关联 id
	OwnerAccountID    string `json:"owner_account_id"`    // 素材拥有者 id
	Status            string `json:"status"`              // 状态
	SampleAspectRatio string `json:"sample_aspect_ratio"` // 图片宽高比
	SimilarityStatus  string `json:"similarity_status"`   // 相似度检测状态
}

// ImageGetResp 获取图片信息响应
// https://developers.e.qq.com/v3.0/docs/api/images/get
type ImageGetResp struct {
	List     []*ImageAssetItem `json:"list"`      // 返回信息列表
	PageInfo *PageInfo         `json:"page_info"` // 分页配置信息
}
