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

// ========== 添加图片文件 ==========
// https://developers.e.qq.com/v3.0/docs/api/images/add

// 上传方式枚举
const (
	ImageUploadTypeFile  = "UPLOAD_TYPE_FILE"  // 文件二进制流上传
	ImageUploadTypeBytes = "UPLOAD_TYPE_BYTES" // base64 编码上传
)

// 图片用途枚举
const (
	ImageUsageDefault          = "IMAGE_USAGE_DEFAULT"           // 默认
	ImageUsageMarketingPendant = "IMAGE_USAGE_MARKETING_PENDANT" // 营销挂件
	ImageUsageShopImg          = "IMAGE_USAGE_SHOP_IMG"          // 卖点图片
)

// 字段限制常量
const (
	ImageSignatureBytes      = 32       // signature 固定长度（字节）
	MaxImageDescriptionBytes = 255      // description 最大字节数
	MaxImageBytesSize        = 10485760 // bytes 字段最大长度（10M）
	MinImageResizeWidth      = 1        // resize_width 最小值
	MaxImageResizeWidth      = 4000     // resize_width 最大值
	MinImageResizeHeight     = 1        // resize_height 最小值
	MaxImageResizeHeight     = 4000     // resize_height 最大值
)

// ImageAddReq 添加图片文件请求（multipart/form-data）
// https://developers.e.qq.com/v3.0/docs/api/images/add
type ImageAddReq struct {
	GlobalReq
	AccountID      int64  // 推广账户 id，与 organization_id 必填其一
	OrganizationID int64  // 业务单元 id，与 account_id 必填其一
	UploadType     string // 上传方式 (必填)：UPLOAD_TYPE_FILE 或 UPLOAD_TYPE_BYTES
	Signature      string // 图片文件签名，md5 值 (必填)，固定 32 字节
	ImageFile      []byte // 图片二进制流，upload_type=UPLOAD_TYPE_FILE 时必填，支持 jpg/png/gif，≤10M
	ImageFileName  string // 图片文件名（含扩展名），用于 multipart 表单
	Bytes          string // 图片 base64 编码，upload_type=UPLOAD_TYPE_BYTES 时必填，1~10485760 字节
	ImageUsage     string // 图片用途，可选值见枚举
	Description    string // 图片文件描述，0-255 字节，不支持@等特殊符号
	ResizeWidth    int    // 图片宽度，单位 px，1-4000
	ResizeHeight   int    // 图片高度，单位 px，1-4000
	ResizeFileSize int    // 图片大小，单位 B(byte)
}

func (p *ImageAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证添加图片文件请求参数
func (p *ImageAddReq) Validate() error {
	if p.AccountID == 0 && p.OrganizationID == 0 {
		return errors.New("account_id 和 organization_id 需必填其一")
	}
	if p.UploadType == "" {
		return errors.New("upload_type为必填")
	}
	if p.UploadType != ImageUploadTypeFile && p.UploadType != ImageUploadTypeBytes {
		return errors.New("upload_type可选值：UPLOAD_TYPE_FILE 或 UPLOAD_TYPE_BYTES")
	}
	if p.Signature == "" {
		return errors.New("signature为必填")
	}
	if len(p.Signature) != ImageSignatureBytes {
		return errors.New("signature长度必须为32字节")
	}
	if p.UploadType == ImageUploadTypeFile {
		if len(p.ImageFile) == 0 {
			return errors.New("upload_type=UPLOAD_TYPE_FILE 时，file为必填")
		}
		if p.ImageFileName == "" {
			return errors.New("upload_type=UPLOAD_TYPE_FILE 时，image_file_name为必填，需包含文件扩展名")
		}
	}
	if p.UploadType == ImageUploadTypeBytes {
		if p.Bytes == "" {
			return errors.New("upload_type=UPLOAD_TYPE_BYTES 时，bytes为必填")
		}
		if len(p.Bytes) > MaxImageBytesSize {
			return errors.New("bytes长度不能超过10485760字节")
		}
	}
	if len(p.Description) > MaxImageDescriptionBytes {
		return errors.New("description长度不能超过255字节")
	}
	if p.ResizeWidth != 0 && (p.ResizeWidth < MinImageResizeWidth || p.ResizeWidth > MaxImageResizeWidth) {
		return errors.New("resize_width须在1-4000之间")
	}
	if p.ResizeHeight != 0 && (p.ResizeHeight < MinImageResizeHeight || p.ResizeHeight > MaxImageResizeHeight) {
		return errors.New("resize_height须在1-4000之间")
	}
	return p.GlobalReq.Validate()
}

// ImageAddResp 添加图片文件响应
// https://developers.e.qq.com/v3.0/docs/api/images/add
type ImageAddResp struct {
	ImageID        string `json:"image_id"`        // 图片 id
	ImageWidth     int    `json:"image_width"`     // 图片宽度，单位 px
	ImageHeight    int    `json:"image_height"`    // 图片高度，单位 px
	ImageFileSize  int64  `json:"image_file_size"` // 图片大小，单位 B(byte)
	ImageType      string `json:"image_type"`      // 图片类型
	ImageSignature string `json:"image_signature"` // 图片文件签名（md5值）
	OuterImageID   string `json:"outer_image_id"`  // 调用方图片 id
	PreviewURL     string `json:"preview_url"`     // 预览地址
	Description    string `json:"description"`     // 图片文件描述
}

// ========== 修改图片信息 ==========
// https://developers.e.qq.com/v3.0/docs/api/images/update

// 字段限制常量
const (
	MinImageIDBytes = 1  // image_id 最小长度（字节）
	MaxImageIDBytes = 64 // image_id 最大长度（字节）
)

// ImageUpdateReq 修改图片信息请求
// https://developers.e.qq.com/v3.0/docs/api/images/update
type ImageUpdateReq struct {
	GlobalReq
	AccountID      int64  `json:"account_id,omitempty"`      // 推广账户 id，与 organization_id 必填其一
	OrganizationID int64  `json:"organization_id,omitempty"` // 业务单元 id，与 account_id 必填其一
	ImageID        string `json:"image_id"`                  // 图片 id (必填)，1-64 字节
	Description    string `json:"description"`               // 图片文件描述 (必填)，0-255 字节，不支持@等特殊符号
}

func (p *ImageUpdateReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证修改图片信息请求参数
func (p *ImageUpdateReq) Validate() error {
	if p.AccountID == 0 && p.OrganizationID == 0 {
		return errors.New("account_id 和 organization_id 需必填其一")
	}
	if p.ImageID == "" {
		return errors.New("image_id为必填")
	}
	if len(p.ImageID) < MinImageIDBytes || len(p.ImageID) > MaxImageIDBytes {
		return errors.New("image_id长度须在1-64字节之间")
	}
	if len(p.Description) > MaxImageDescriptionBytes {
		return errors.New("description长度不能超过255字节")
	}
	return p.GlobalReq.Validate()
}

// ImageUpdateResp 修改图片信息响应
// https://developers.e.qq.com/v3.0/docs/api/images/update
type ImageUpdateResp struct {
	ImageID string `json:"image_id"` // 图片 id
}
