package model

import "errors"

// ========== 获取视频文件 ==========
// https://developers.e.qq.com/v3.0/docs/api/videos/get

// 视频过滤字段常量
const (
	VideoFilterFieldProductCatalogID         = "product_catalog_id"
	VideoFilterFieldProductOuterID           = "product_outer_id"
	VideoFilterFieldMediaSignature           = "media_signature"
	VideoFilterFieldMediaID                  = "media_id"
	VideoFilterFieldMediaWidth               = "media_width"
	VideoFilterFieldMediaHeight              = "media_height"
	VideoFilterFieldCreatedTime              = "created_time"
	VideoFilterFieldLastModifiedTime         = "last_modified_time"
	VideoFilterFieldSourceType               = "source_type"
	VideoFilterFieldOwnerAccountID           = "owner_account_id"
	VideoFilterFieldStatus                   = "status"
	VideoFilterFieldMediaDescription         = "media_description"
	VideoFilterFieldSampleAspectRatio        = "sample_aspect_ratio"
	VideoFilterFieldFirstPublicationStatus   = "first_publication_status"
	VideoFilterFieldVideoID                  = "video_id"
	VideoFilterFieldFileSize                 = "file_size"
	VideoFilterFieldHeight                   = "height"
	VideoFilterFieldWidth                    = "width"
	VideoFilterFieldRatio                    = "ratio"
	VideoFilterFieldVideoDurationMillisecond = "video_duration_millisecond"
	VideoFilterFieldAigcFlag                 = "aigc_flag"
)

// 视频来源类型枚举
const (
	VideoSourceTypeLocal         = "SOURCE_TYPE_LOCAL"
	VideoSourceTypeAPI           = "SOURCE_TYPE_API"
	VideoSourceTypeVideoMakerXsj = "SOURCE_TYPE_VIDEO_MAKER_XSJ"
	VideoSourceTypeTCC           = "SOURCE_TYPE_TCC"
	VideoSourceTypeDerive        = "SOURCE_TYPE_DERIVE"
	VideoSourceTypeDerivation    = "SOURCE_TYPE_DERIVATION"
	VideoSourceTypeHuxuan        = "SOURCE_TYPE_HUXUAN"
	VideoSourceTypeHuxuanDerive  = "SOURCE_TYPE_HUXUAN_DERIVE"
)

// 视频状态枚举
const (
	VideoStatusNormal  = "ADSTATUS_NORMAL"
	VideoStatusDeleted = "ADSTATUS_DELETED"
)

// 首发状态枚举
const (
	FirstPublicationStatusDefault          = "FIRST_PUBLICATION_STATUS_DEFAULT"
	FirstPublicationStatusFirstPublication = "FIRST_PUBLICATION_STATUS_FIRST_PUBLICATION"
)

// AIGC 标记枚举
const (
	AigcFlagUnknown     = "AIGC_FLAG_UNKNOWN"
	AigcFlagNotAI       = "AIGC_FLAG_NOT_AI"
	AigcFlagUseMuseAI   = "AIGC_FLAG_USE_MUSE_AI"
	AigcFlagUseOthersAI = "AIGC_FLAG_USE_OTHERS_AI"
)

// 字段限制常量
const (
	MaxVideoGetFilteringCount = 4     // filtering 最大长度
	MinVideoGetPage           = 1     // page 最小值
	MaxVideoGetPage           = 99999 // page 最大值
	MinVideoGetPageSize       = 1     // page_size 最小值
	MaxVideoGetPageSize       = 100   // page_size 最大值
	DefaultVideoGetPage       = 1     // page 默认值
	DefaultVideoGetPageSize   = 10    // page_size 默认值
)

// VideoFilteringItem 视频过滤条件
type VideoFilteringItem struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)
	Operator string   `json:"operator"` // 操作符 (必填)
	Values   []string `json:"values"`   // 字段取值 (必填)
}

// Validate 验证单个过滤条件
func (f *VideoFilteringItem) Validate() error {
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

// VideoGetReq 获取视频文件请求
// https://developers.e.qq.com/v3.0/docs/api/videos/get
type VideoGetReq struct {
	GlobalReq
	AccountID        int64                 `json:"account_id,omitempty"`        // 广告主帐号 id，与 organization_id 必填其一
	OrganizationID   int64                 `json:"organization_id,omitempty"`   // 业务单元 id，与 account_id 必填其一
	Filtering        []*VideoFilteringItem `json:"filtering,omitempty"`         // 过滤条件，最大4
	Page             int                   `json:"page,omitempty"`              // 搜索页码，1-99999，默认1
	PageSize         int                   `json:"page_size,omitempty"`         // 每页条数，1-100，默认10
	LabelID          int64                 `json:"label_id,omitempty"`          // 标签 id
	BusinessScenario int                   `json:"business_scenario,omitempty"` // 业务场景：1=内容素材包，2=投放素材包
}

func (p *VideoGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page == 0 {
		p.Page = DefaultVideoGetPage
	}
	if p.PageSize == 0 {
		p.PageSize = DefaultVideoGetPageSize
	}
}

// Validate 验证获取视频文件请求参数
func (p *VideoGetReq) Validate() error {
	if p.AccountID == 0 && p.OrganizationID == 0 {
		return errors.New("account_id 和 organization_id 需必填其一")
	}
	if len(p.Filtering) > MaxVideoGetFilteringCount {
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
	if p.Page < MinVideoGetPage || p.Page > MaxVideoGetPage {
		return errors.New("page须在1-99999之间")
	}
	if p.PageSize < MinVideoGetPageSize || p.PageSize > MaxVideoGetPageSize {
		return errors.New("page_size须在1-100之间")
	}
	return p.GlobalReq.Validate()
}

// VideoListItem 视频文件列表项
type VideoListItem struct {
	VideoID                  int64   `json:"video_id"`                   // 视频 id
	Width                    int     `json:"width"`                      // 视频宽度
	Height                   int     `json:"height"`                     // 视频高度
	VideoFrames              int     `json:"video_frames"`               // 视频帧数
	VideoFPS                 float64 `json:"video_fps"`                  // 视频帧率
	VideoCodec               string  `json:"video_codec"`                // 视频格式
	VideoBitRate             int     `json:"video_bit_rate"`             // 视频码率，单位：b/s
	AudioCodec               string  `json:"audio_codec"`                // 音频格式
	AudioBitRate             int     `json:"audio_bit_rate"`             // 音频码率，单位：b/s
	FileSize                 int64   `json:"file_size"`                  // 视频文件大小，单位 B
	Type                     string  `json:"type"`                       // 视频类型
	Signature                string  `json:"signature"`                  // 视频文件签名
	SystemStatus             string  `json:"system_status"`              // 转码状态
	Description              string  `json:"description"`                // 视频文件描述
	PreviewURL               string  `json:"preview_url"`                // 视频文件预览地址
	KeyFrameImageURL         string  `json:"key_frame_image_url"`        // 视频首帧缩略图地址
	CreatedTime              int64   `json:"created_time"`               // 创建时间，时间戳
	LastModifiedTime         int64   `json:"last_modified_time"`         // 最后修改时间，时间戳
	VideoProfileName         string  `json:"video_profile_name"`         // 视频格式类型
	AudioSampleRate          int     `json:"audio_sample_rate"`          // 音频采样率，单位：hz
	MaxKeyframeInterval      int     `json:"max_keyframe_interval"`      // 关键帧的最大间隔帧数
	MinKeyframeInterval      int     `json:"min_keyframe_interval"`      // 关键帧的最小间隔帧数
	SampleAspectRatio        string  `json:"sample_aspect_ratio"`        // 采样纵横比
	AudioProfileName         string  `json:"audio_profile_name"`         // 音频格式类型
	ScanType                 string  `json:"scan_type"`                  // 扫描类型
	ImageDurationMillisecond int     `json:"image_duration_millisecond"` // 画面时长，单位：ms
	AudioDurationMillisecond int     `json:"audio_duration_millisecond"` // 音频时长，单位：ms
	SourceType               string  `json:"source_type"`                // 视频来源
	ProductOuterID           string  `json:"product_outer_id"`           // 商品 id
	SourceReferenceID        string  `json:"source_reference_id"`        // 素材来源关联 id
	OwnerAccountID           string  `json:"owner_account_id"`           // 素材拥有 id
	Status                   string  `json:"status"`                     // 视频状态
	SimilarityStatus         string  `json:"similarity_status"`          // 相似度检测状态
}

// VideoGetResp 获取视频文件响应
// https://developers.e.qq.com/v3.0/docs/api/videos/get
type VideoGetResp struct {
	List     []*VideoListItem `json:"list"`      // 返回信息列表
	PageInfo *PageInfo        `json:"page_info"` // 分页配置信息
}

// ========== 添加视频文件 ==========
// https://developers.e.qq.com/v3.0/docs/api/videos/add

// 字段限制常量
const (
	VideoSignatureBytes      = 32  // signature 固定长度（字节）
	MaxVideoDescriptionBytes = 255 // description 最大字节数
)

// VideoAddReq 添加视频文件请求（multipart/form-data）
// https://developers.e.qq.com/v3.0/docs/api/videos/add
type VideoAddReq struct {
	GlobalReq
	AccountID            int64  // 广告主账户 id，与 organization_id 必填其一
	OrganizationID       int64  // 业务单元 id，与 account_id 必填其一
	VideoFile            []byte // 被上传的视频文件二进制流 (必填)，支持 mp4/mov/avi，最大100M
	VideoFileName        string // 视频文件名（含扩展名，如 video.mp4），用于 multipart 表单
	Signature            string // 视频文件签名 (必填)，固定32字节
	Description          string // 视频文件描述，0-255字节
	AdcreativeTemplateID int64  // 创意形式 id，仅可上传微信规格
}

func (p *VideoAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证添加视频文件请求参数
func (p *VideoAddReq) Validate() error {
	if p.AccountID == 0 && p.OrganizationID == 0 {
		return errors.New("account_id 和 organization_id 需必填其一")
	}
	if len(p.VideoFile) == 0 {
		return errors.New("video_file为必填")
	}
	if p.VideoFileName == "" {
		return errors.New("video_file_name为必填，需包含文件扩展名如 video.mp4")
	}
	if p.Signature == "" {
		return errors.New("signature为必填")
	}
	if len(p.Signature) != VideoSignatureBytes {
		return errors.New("signature长度必须为32字节")
	}
	if len(p.Description) > MaxVideoDescriptionBytes {
		return errors.New("description长度不能超过255字节")
	}
	return p.GlobalReq.Validate()
}

// VideoAddResp 添加视频文件响应
// https://developers.e.qq.com/v3.0/docs/api/videos/add
type VideoAddResp struct {
	VideoID      int64 `json:"video_id"`       // 视频 id
	CoverImageID int64 `json:"cover_image_id"` // 视频封面图 id
}

// ========== 修改视频信息 ==========
// https://developers.e.qq.com/v3.0/docs/api/videos/update

// VideoUpdateReq 修改视频信息请求
// https://developers.e.qq.com/v3.0/docs/api/videos/update
type VideoUpdateReq struct {
	GlobalReq
	AccountID      int64  `json:"account_id,omitempty"`      // 广告主账户 id，与 organization_id 必填其一
	OrganizationID int64  `json:"organization_id,omitempty"` // 业务单元 id，与 account_id 必填其一
	VideoID        int64  `json:"video_id"`                  // 视频 id (必填)
	Description    string `json:"description"`               // 视频文件描述 (必填)，0-255字节
}

func (p *VideoUpdateReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证修改视频信息请求参数
func (p *VideoUpdateReq) Validate() error {
	if p.AccountID == 0 && p.OrganizationID == 0 {
		return errors.New("account_id 和 organization_id 需必填其一")
	}
	if p.VideoID == 0 {
		return errors.New("video_id为必填")
	}
	if len(p.Description) > MaxVideoDescriptionBytes {
		return errors.New("description长度不能超过255字节")
	}
	return p.GlobalReq.Validate()
}

// VideoUpdateResp 修改视频信息响应
// https://developers.e.qq.com/v3.0/docs/api/videos/update
type VideoUpdateResp struct {
	VideoID int64 `json:"video_id"` // 视频 id
}
