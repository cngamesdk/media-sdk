package model

import (
	"errors"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/random"
	"time"
)

const (
	DevelopersUrl = "https://developers.e.qq.com"
	ApiUrl        = "https://api.e.qq.com"
	ApiUrl3       = ApiUrl + "/v3.0"
)

// GlobalReq 全局参数
type GlobalReq struct {
	AccessToken string `json:"access_token"` // 授权令牌 (必填)
	Timestamp   int64  `json:"timestamp"`    // 当前时间戳，单位为秒 (必填)
	Nonce       string `json:"nonce"`        // 随机字符串标识，不超过32个字符 (必填)
}

// 常量定义
const (
	MaxTimestampDiff = 300         // 最大时间误差（秒）
	MaxNonceLength   = 32          // 随机字符串最大长度
	TimezoneOffset   = 8 * 60 * 60 // GMT+8时区偏移（秒）
)

func (p *GlobalReq) Format() {
	if p.Nonce == "" {
		uuid, _ := random.UUIdV4()
		p.Nonce = cryptor.Md5String(uuid)
	}
	if p.Timestamp <= 0 {
		p.Timestamp = GetCurrentTimestamp()
	}
}

// Validate 验证API请求公共参数
func (p *GlobalReq) Validate() error {
	// 1. 验证access_token
	if p.AccessToken == "" {
		return errors.New("access_token为必填")
	}

	// 2. 验证timestamp
	if p.Timestamp == 0 {
		return errors.New("timestamp为必填")
	}
	if err := validateTimestamp(p.Timestamp); err != nil {
		return err
	}

	// 3. 验证nonce
	if p.Nonce == "" {
		return errors.New("nonce为必填")
	}
	if len(p.Nonce) > MaxNonceLength {
		return errors.New("nonce长度不能超过32个字符")
	}

	return nil
}

// validateTimestamp 验证时间戳
func validateTimestamp(timestamp int64) error {
	now := time.Now().Unix()
	diff := now - timestamp
	if diff < 0 {
		diff = -diff
	}
	if diff > MaxTimestampDiff {
		return errors.New("timestamp与服务器时间误差超过300秒")
	}
	return nil
}

// TimestampToTime 将时间戳转换为GMT+8时区的时间
func TimestampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0).In(time.FixedZone("GMT+8", TimezoneOffset))
}

// GetCurrentTimestamp 获取当前GMT+8时区的秒级时间戳
func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

// 常量定义 - 分页方式
const (
	PaginationModeNormal = "PAGINATION_MODE_NORMAL" // 普通分页模式
	PaginationModeCursor = "PAGINATION_MODE_CURSOR" // 游标分页模式
)

// 分页限制常量
const (
	MinPage              = 1
	MaxPage              = 100
	MinPageSize          = 1
	MaxPageSize          = 100
	MaxNormalPageProduct = 1000 // page_size * page 的最大值
	MinCursor            = 1
)

type CursorPageV2Req struct {
	PaginationMode string `json:"pagination_mode"`     // 分页方式 (必填)
	Page           int    `json:"page,omitempty"`      // 搜索页码
	PageSize       int    `json:"page_size,omitempty"` // 一页显示的数据条数
	Cursor         string `json:"cursor,omitempty"`    // 游标翻页模式下的游标值
}

func (p *CursorPageV2Req) Format() {
	p.setPageDefaults()
}

func (p *CursorPageV2Req) Validate() error {
	if validateErr := p.validatePageParams(); validateErr != nil {
		return validateErr
	}
	if validateErr := p.validatePaginationMode(); validateErr != nil {
		return validateErr
	}
	if validateErr := p.validateCursor(); validateErr != nil {
		return validateErr
	}
	return nil
}

// validatePageParams 验证分页参数
func (p *CursorPageV2Req) validatePageParams() error {
	if p.Page < MinPage || p.Page > MaxPage {
		return errors.New("page必须在1-100之间")
	}
	if p.PageSize < MinPageSize || p.PageSize > MaxPageSize {
		return errors.New("page_size必须在1-100之间")
	}
	return nil
}

// validatePaginationMode 验证分页方式
func (p *CursorPageV2Req) validatePaginationMode() error {
	if p.PaginationMode == "" {
		p.PaginationMode = PaginationModeNormal
		return nil
	}
	if p.PaginationMode != PaginationModeNormal && p.PaginationMode != PaginationModeCursor {
		return errors.New("pagination_mode值无效，允许值：PAGINATION_MODE_NORMAL、PAGINATION_MODE_CURSOR")
	}
	return nil
}

// validateCursor 验证游标值
func (p *CursorPageV2Req) validateCursor() error {
	if p.Cursor == "" {
		return nil
	}
	if len(p.Cursor) > MaxCursorLength {
		return errors.New("cursor长度不能超过10字节")
	}
	return nil
}

// setPageDefaults 设置分页默认值
func (p *CursorPageV2Req) setPageDefaults() {
	if p.Page <= 0 {
		p.Page = DefaultPage
	}
	if p.PageSize <= 0 {
		p.PageSize = DefaultPageSize
	}
}

type CursorPageReq struct {
	PaginationMode string `json:"pagination_mode"`     // 分页方式 (必填)
	Page           int    `json:"page,omitempty"`      // 搜索页码
	PageSize       int    `json:"page_size,omitempty"` // 一页显示的数据条数
	Cursor         int64  `json:"cursor,omitempty"`    // 游标翻页模式下的游标值
}

func (p *CursorPageReq) Format() {
	p.setPageDefaults()
}

func (p *CursorPageReq) Validate() error {
	// 验证分页方式
	if p.PaginationMode == "" {
		return errors.New("pagination_mode为必填")
	}
	if p.PaginationMode != PaginationModeNormal && p.PaginationMode != PaginationModeCursor {
		return errors.New("pagination_mode值无效，允许值：PAGINATION_MODE_NORMAL、PAGINATION_MODE_CURSOR")
	}

	// 根据分页方式验证
	switch p.PaginationMode {
	case PaginationModeNormal:
		return p.validateNormalPagination()
	case PaginationModeCursor:
		return p.validateCursorPagination()
	}

	return nil
}

// validateNormalPagination 验证普通分页模式
func (p *CursorPageReq) validateNormalPagination() error {
	// 验证page范围
	if p.Page < MinPage || p.Page > MaxPage {
		return errors.New("page必须在1-100之间")
	}

	// 验证page_size范围
	if p.PageSize < MinPageSize || p.PageSize > MaxPageSize {
		return errors.New("page_size必须在1-100之间")
	}

	// 验证page_size * page <= 1000
	if p.PageSize*p.Page > MaxNormalPageProduct {
		return errors.New("page_size * page 必须小于等于1000")
	}

	return nil
}

// validateCursorPagination 验证游标分页模式
func (p *CursorPageReq) validateCursorPagination() error {
	// cursor可选，如果传了则验证
	if p.Cursor != 0 && p.Cursor < MinCursor {
		return errors.New("cursor必须大于等于1")
	}

	// 游标模式下，page无效（忽略）
	// page_size可选，如果传了则验证范围
	if p.PageSize != 0 {
		if p.PageSize < MinPageSize || p.PageSize > MaxPageSize {
			return errors.New("page_size必须在1-100之间")
		}
	} else {
		p.PageSize = MinPageSize // 设置默认值
	}

	return nil
}

// setPageDefaults 设置分页默认值
func (p *CursorPageReq) setPageDefaults() {
	if p.Page <= 0 {
		p.Page = MinPage
	}
	if p.PageSize <= 0 {
		p.PageSize = MinPageSize
	}
}

type BaseResp struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	MessageCn string      `json:"message_cn"`
	Data      interface{} `json:"data"`
	RequestId string      `json:"request_id"`
}

type PageInfoContainer struct {
	PageInfo *PageInfo `json:"page_info,omitempty"` // 分页配置信息（普通翻页模式）
}

// PageInfo 分页配置信息（普通翻页模式）
type PageInfo struct {
	Page        int `json:"page"`         // 搜索页码，默认值：1
	PageSize    int `json:"page_size"`    // 一页显示的数据条数
	TotalNumber int `json:"total_number"` // 总条数
	TotalPage   int `json:"total_page"`   // 总页数
}

type CursorPageInfoContainer struct {
	CursorPageInfo *CursorPageInfo `json:"cursor_page_info,omitempty"` // 分页配置信息（游标翻页模式）
}

// CursorPageInfo 分页配置信息（游标翻页模式）
type CursorPageInfo struct {
	PageSize    int   `json:"page_size"`    // 一页显示的数据条数
	TotalNumber int   `json:"total_number"` // 总条数
	HasMore     bool  `json:"has_more"`     // 是否有下一页，返回false表示已无下一页，此时务必停止拉取
	Cursor      int64 `json:"cursor"`       // 下一次拉取的游标值
}
