package model

import (
	"errors"
	"strings"
)

// AssetQueryParams 资产查询参数
type EbpAppListReq struct {
	accessTokenReq
	AccountID            int64             `json:"account_id"`                       // 账户ID (必填)
	AccountType          string            `json:"account_type"`                     // 账户类型，允许值：EBP 升级版巨量引擎工作台 (必填)
	AssetManagementScope string            `json:"asset_management_scope,omitempty"` // 资产范围：DIRECT仅查询入参组织创建或被共享的资产，TRAVERSE查询入参组织及下属组织
	Filtering            *EbpAppListFilter `json:"filtering,omitempty"`              // 过滤条件
	PublishTime          *PublishTimeRange `json:"publish_time,omitempty"`           // 按发布时间查询的时间范围
	PageInfoReq
}

// EbpAppListFilter 资产过滤条件
type EbpAppListFilter struct {
	SearchKey      string `json:"search_key,omitempty"`      // 搜索关键字（应用名称/APPID）
	AssetOwnership string `json:"asset_ownership,omitempty"` // 资产来源：CREATE仅查询组织及下级组织创建的资产，SHARE仅查询组织及下级组织被共享的资产
	Status         string `json:"status,omitempty"`          // 筛选应用状态
}

// PublishTimeRange 发布时间范围
type PublishTimeRange struct {
	Start string `json:"start,omitempty"` // 发布起始时间，格式：%Y-%m-%d
	End   string `json:"end,omitempty"`   // 发布结束时间，格式：%Y-%m-%d
}

func (receiver *EbpAppListReq) Format() {
	receiver.accessTokenReq.Format()
	receiver.AccountType = strings.TrimSpace(receiver.AccountType)
	receiver.PageInfoReq.Format()
}

func (receiver *EbpAppListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if len(receiver.AccountType) <= 0 {
		err = errors.New("account_type is empty")
		return
	}
	if receiver.AccountID <= 0 {
		err = errors.New("account_id is invalid")
		return
	}
	return
}

// EbpAppListReq 获取安卓应用列表数据
type EbpAppListResp struct {
	BasicAppList []BasicAppInfo `json:"basic_app_list,omitempty"` // 应用包列表信息
	PageInfo     PageInfoResp   `json:"page_info,omitempty"`      // 分页信息
}

// BasicAppInfo 应用包信息
type BasicAppInfo struct {
	BasicPackageID     string `json:"basic_package_id"`               // 应用包ID（创建分包需要使用此id入参）
	AppName            string `json:"app_name"`                       // 应用名称
	AppNameEn          string `json:"app_name_en"`                    // 英文应用名称
	PackageName        string `json:"package_name"`                   // 包名
	VersionCode        string `json:"version_code"`                   // 版本号
	VersionName        string `json:"version_name"`                   // 版本名称
	AppLogo            string `json:"app_logo"`                       // 应用logo
	PublishTime        string `json:"publish_time"`                   // 发布时间
	Reason             string `json:"reason,omitempty"`               // 拒审原因
	SuccessReason      string `json:"success_reason,omitempty"`       // 审核成功信息
	HistoryAccountID   int64  `json:"history_account_id,omitempty"`   // 历史来源账户id
	HistoryAccountType string `json:"history_account_type,omitempty"` // 历史来源账户类型
	HistoryAccountName string `json:"history_account_name,omitempty"` // 历史来源账户名称
	IsEBPAsset         bool   `json:"is_ebp_asset"`                   // 是否为EBP资产
	HasExtendPackage   bool   `json:"has_extend_package"`             // 是否有分包
	DownloadURL        string `json:"download_url,omitempty"`         // 下载链接
	CreateTime         string `json:"create_time"`                    // 创建时间
}

// EbpAppExtendCreateReq 创建安卓分包请求参数
type EbpAppExtendCreateReq struct {
	accessTokenReq
	AccountID       int64          `json:"account_id"`        // 账户ID (必填)
	AccountType     string         `json:"account_type"`      // 账户类型，允许值：EBP 升级版巨量引擎工作台 (必填)
	PackageID       string         `json:"package_id"`        // 母包ID (必填)
	Mode            string         `json:"mode"`              // 分包模式：AUTO 自动，CUSTOMIZE 自定义，MANUAL 手动 (必填)
	ChannelCount    int            `json:"channel_count"`     // 自动生成模式下要指定数量 (必填，AUTO模式使用)
	ChannelInfoList []*ChannelInfo `json:"channel_info_list"` // 自定义模式下必须传入渠道号信息 (必填，CUSTOMIZE模式使用)
}

func (receiver *EbpAppExtendCreateReq) Format() {
	receiver.accessTokenReq.Format()
	receiver.AccountType = strings.TrimSpace(receiver.AccountType)
	receiver.Mode = strings.TrimSpace(receiver.Mode)
	receiver.PackageID = strings.TrimSpace(receiver.PackageID)
}

func (receiver *EbpAppExtendCreateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AccountID <= 0 {
		err = errors.New("account_id is empty")
		return
	}
	if len(receiver.AccountType) <= 0 {
		err = errors.New("account_type is empty")
		return
	}
	if len(receiver.PackageID) <= 0 {
		err = errors.New("package_id is empty")
		return
	}
	if len(receiver.Mode) <= 0 {
		err = errors.New("mode is empty")
		return
	}
	return
}

func (receiver *EbpAppExtendCreateReq) GetHeaders() headersMap {
	headers := receiver.accessTokenReq.GetHeaders()
	headers.Json()
	return headers
}

// ChannelInfo 渠道信息
type ChannelInfo struct {
	ChannelID string `json:"channel_id"`       // 渠道号 (必填)
	Remark    string `json:"remark,omitempty"` // 备注，渠道包含信息，至多20个字符
}

// 常量定义
const (
	// SubpackageMode 分包模式
	SubpackageModeAuto      = "AUTO"      // 自动
	SubpackageModeCustomize = "CUSTOMIZE" // 自定义
	SubpackageModeManual    = "MANUAL"    // 手动
)

// 验证常量
const (
	MaxChannelIDLength = 50 // 渠道号最大长度
	MaxRemarkLength    = 20 // 备注最大长度
)

type EbpAppExtendCreateResp struct {
	PackageId string `json:"package_id"`
}
