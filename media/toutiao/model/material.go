package model

import "errors"

type EbpVideoUploadReq struct {
	accessTokenReq
	AccountID      int64    `json:"account_id,omitempty"`      // 账户ID (必填)
	AccountType    string   `json:"account_type,omitempty"`    // 账户类型，默认EBP
	UploadType     string   `json:"upload_type,omitempty"`     // 视频上传方式，默认UPLOAD_BY_FILE
	VideoSignature string   `json:"video_signature,omitempty"` // 视频MD5值 (条件必填)
	VideoFile      []byte   `json:"video_file,omitempty"`      // 视频文件二进制内容 (条件必填)
	FileName       string   `json:"file_name,omitempty"`       // 素材文件名
	Labels         []string `json:"labels,omitempty"`          // 标签
	VideoURL       string   `json:"video_url,omitempty"`       // 视频URL地址
	IsAigc         bool     `json:"is_aigc,omitempty"`         // 是否为AIGC生成
}

func (p *EbpVideoUploadReq) Format() {
	p.accessTokenReq.Format()
}

func (p *EbpVideoUploadReq) Validate() (err error) {
	if validateErr := p.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if p.AccountID <= 0 {
		err = errors.New("account_id is empty")
		return
	}

	// 2. 设置默认值
	p.setDefaults()

	// 3. 验证账户类型
	if validateErr := p.validateAccountType(); validateErr != nil {
		return validateErr
	}

	// 4. 验证上传方式
	if validateErr := p.validateUploadType(); validateErr != nil {
		return validateErr
	}

	// 5. 验证文件名
	if validateErr := p.validateFileName(); validateErr != nil {
		return validateErr
	}

	return
}

// 常量定义
const (

	// 视频上传方式
	UploadTypeFile = "UPLOAD_BY_FILE" // 文件上传（默认值）
	UploadTypeURL  = "UPLOAD_BY_URL"  // URL链接上传

	// 视频格式要求
	MinBitrate  = 516  // 最小码率 (kbps)
	MaxFileSize = 500  // 最大文件大小 (MB)
	MinDuration = 4    // 最小时长 (秒)
	MaxDuration = 300  // 最大时长 (秒)
	MinWidth    = 1280 // 最小宽度
	MaxWidth    = 2560 // 最大宽度
	MinHeight   = 720  // 最小高度
	MaxHeight   = 1440 // 最大高度

	// 文件名长度限制
	MaxFileNameLength = 255
)

// 支持的视频格式
var SupportedVideoFormats = map[string]bool{
	"mp4":  true,
	"mpeg": true,
	"3gp":  true,
	"avi":  true,
}

// setDefaults 设置默认值
func (p *EbpVideoUploadReq) setDefaults() {
	if p.AccountType == "" {
		p.AccountType = AccountTypeEBP
	}
	if p.UploadType == "" {
		p.UploadType = UploadTypeFile
	}
}

// validateAccountType 验证账户类型
func (p *EbpVideoUploadReq) validateAccountType() error {
	validTypes := map[string]bool{
		AccountTypeEBP: true,
	}

	if !validTypes[p.AccountType] {
		return errors.New("account_type值无效，允许值：EBP")
	}

	return nil
}

// validateUploadType 验证上传方式
func (p *EbpVideoUploadReq) validateUploadType() error {
	validTypes := map[string]bool{
		UploadTypeFile: true,
		UploadTypeURL:  true,
	}

	if !validTypes[p.UploadType] {
		return errors.New("upload_type值无效，允许值：UPLOAD_BY_FILE、UPLOAD_BY_URL")
	}

	// 根据上传方式验证相应字段
	switch p.UploadType {
	case UploadTypeFile:
		return p.validateFileUpload()
	case UploadTypeURL:
		return p.validateURLUpload()
	}

	return nil
}

// validateFileUpload 验证文件上传方式
func (p *EbpVideoUploadReq) validateFileUpload() error {
	// video_signature 必填
	if p.VideoSignature == "" {
		return errors.New("当upload_type=UPLOAD_BY_FILE时，video_signature为必填")
	}

	// video_file 必填
	if len(p.VideoFile) == 0 {
		return errors.New("当upload_type=UPLOAD_BY_FILE时，video_file为必填")
	}

	// 视频格式验证需要在业务层进行
	// 1. 文件大小 ≤ 500MB
	// 2. 视频码率 > 516kbps
	// 3. 分辨率 1280*720 ≤ 分辨率 < 2560*1440
	// 4. 时长 4s < 时长 < 300s
	// 5. 格式支持：mp4、mpeg、3gp、avi

	return nil
}

// validateURLUpload 验证URL上传方式
func (p *EbpVideoUploadReq) validateURLUpload() error {
	// video_url 必填
	if p.VideoURL == "" {
		return errors.New("当upload_type=UPLOAD_BY_URL时，video_url为必填")
	}

	// 仅支持开发者购买直播山云素材服务上传生成的tos链接
	// 需在业务层验证URL是否为有效的tos链接

	return nil
}

// validateFileName 验证文件名
func (p *EbpVideoUploadReq) validateFileName() error {
	if p.FileName == "" {
		return nil // 不传默认取文件名
	}

	if len(p.FileName) > MaxFileNameLength {
		return errors.New("文件名字符长度不能超过255")
	}

	return nil
}

func (a *EbpVideoUploadReq) GetHeaders() headersMap {
	headers := a.accessTokenReq.GetHeaders()
	headers.Json()
	return headers
}

type EbpVideoUploadResp struct {
	VideoID        string  `json:"video_id"`                  // 视频ID
	Size           int64   `json:"size,omitempty"`            // 视频大小（字节）
	Width          int     `json:"width,omitempty"`           // 视频宽度
	Height         int     `json:"height,omitempty"`          // 视频高度
	VideoURL       string  `json:"video_url,omitempty"`       // 视频地址（url格式）
	Duration       float64 `json:"duration,omitempty"`        // 视频时长（秒）
	CoverImageURL  string  `json:"cover_image_url,omitempty"` // 封面图url（用于预览）
	CoverImageURI  string  `json:"cover_image_uri,omitempty"` // 封面图uri（用于创编）
	MaterialID     int64   `json:"material_id,omitempty"`     // 素材ID
	VideoSignature string  `json:"video_signature,omitempty"` // 视频的md5值
}

type FileImageAdReq struct {
	accessTokenReq
	AdvertiserID   int64  `json:"advertiser_id"`             // 广告主ID (必填)
	UploadType     string `json:"upload_type,omitempty"`     // 图片上传方式，默认UPLOAD_BY_FILE
	ImageSignature string `json:"image_signature,omitempty"` // 图片MD5值 (条件必填)
	ImageFile      []byte `json:"image_file,omitempty"`      // 图片文件二进制内容 (条件必填)
	ImageURL       string `json:"image_url,omitempty"`       // 图片URL地址 (条件必填)
	Filename       string `json:"filename,omitempty"`        // 素材文件名
	IsAigc         bool   `json:"is_aigc,omitempty"`         // 是否为AIGC生成
}

func (p *FileImageAdReq) Format() {
	p.accessTokenReq.Format()
}

// 支持的图片格式
var SupportedImageFormats = map[string]bool{
	"jpg":  true,
	"jpeg": true,
	"png":  true,
	"bmp":  true,
	"gif":  true,
}

// Validate 验证图片上传参数
func (p *FileImageAdReq) Validate() error {
	if validateErr := p.accessTokenReq.Validate(); validateErr != nil {
		return validateErr
	}
	// 1. 验证必填字段
	if p.AdvertiserID <= 0 {
		return errors.New("advertiser_id为必填")
	}

	// 2. 设置默认值
	p.setDefaults()

	// 3. 验证上传方式
	if err := p.validateUploadType(); err != nil {
		return err
	}

	// 4. 验证文件名
	if err := p.validateFilename(); err != nil {
		return err
	}

	return nil
}

// setDefaults 设置默认值
func (p *FileImageAdReq) setDefaults() {
	if p.UploadType == "" {
		p.UploadType = UploadTypeFile
	}
}

// validateUploadType 验证上传方式
func (p *FileImageAdReq) validateUploadType() error {
	validTypes := map[string]bool{
		UploadTypeFile: true,
		UploadTypeURL:  true,
	}

	if !validTypes[p.UploadType] {
		return errors.New("upload_type值无效，允许值：UPLOAD_BY_FILE、UPLOAD_BY_URL")
	}

	// 根据上传方式验证相应字段
	switch p.UploadType {
	case UploadTypeFile:
		return p.validateFileUpload()
	case UploadTypeURL:
		return p.validateURLUpload()
	}

	return nil
}

// validateFileUpload 验证文件上传方式
func (p *FileImageAdReq) validateFileUpload() error {
	// image_signature 必填
	if p.ImageSignature == "" {
		return errors.New("当upload_type=UPLOAD_BY_FILE时，image_signature为必填")
	}

	// image_file 必填
	if len(p.ImageFile) == 0 {
		return errors.New("当upload_type=UPLOAD_BY_FILE时，image_file为必填")
	}

	// 图片格式和大小验证需要在业务层进行
	// 1. 文件大小 ≤ 5MB
	// 2. 格式支持：jpg、jpeg、png、bmp、gif

	return nil
}

// validateURLUpload 验证URL上传方式
func (p *FileImageAdReq) validateURLUpload() error {
	// image_url 必填
	if p.ImageURL == "" {
		return errors.New("当upload_type=UPLOAD_BY_URL时，image_url为必填")
	}

	// URL格式验证
	// 需在业务层验证URL是否有效

	return nil
}

// validateFilename 验证文件名
func (p *FileImageAdReq) validateFilename() error {
	if p.Filename == "" {
		return nil // 不传则默认取文件名
	}

	if len(p.Filename) > MaxFileNameLength {
		return errors.New("文件名字符长度不能超过255")
	}

	return nil
}

func (p *FileImageAdReq) GetHeaders() headersMap {
	headers := p.accessTokenReq.GetHeaders()
	headers.Json()
	return headers
}

type FileImageAdResp struct {
	ID         string `json:"id"`                    // 图片ID
	Size       int64  `json:"size,omitempty"`        // 图片大小（字节）
	Width      int    `json:"width,omitempty"`       // 图片宽度
	Height     int    `json:"height,omitempty"`      // 图片高度
	URL        string `json:"url,omitempty"`         // 图片预览地址
	Format     string `json:"format,omitempty"`      // 图片格式
	Signature  string `json:"signature,omitempty"`   // 图片MD5值
	MaterialID int64  `json:"material_id,omitempty"` // 素材ID，多合一报表中的素材id，一个素材唯一对应一个素材id
}
