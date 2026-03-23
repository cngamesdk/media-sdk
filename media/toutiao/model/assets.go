package model

import (
	"errors"
	"time"
)

// 常量定义 - 资产类型
const (
	AssetTypeApp         = "APP"             // 应用类
	AssetTypeTetris      = "TETRIS_EXTERNAL" // 建站落地页
	AssetTypeThird       = "THIRD_EXTERNAL"  // 三方落地页
	AssetTypeMiniProgram = "MINI_PROGRAMME"  // 小程序
	AssetTypeOffline     = "OFFLINE_EVENT"   // 离线
	AssetTypeOther       = "OTHER"           // 其他
	AssetTypeQuickApp    = "QUICK_APP"       // 快应用
	AssetTypeSite        = "SITE"            // 橙子落地页

	AssetTypeAweme      = "AWEME"      // 抖音主页
	AssetTypeEnterprise = "ENTERPRISE" // 企业号落地页
	AssetTypeMicroApp   = "MICRO_APP"  // 字节小程序
	AssetTypeOrange     = "ORANGE"     // 橙子落地页
	AssetTypeThirdparty = "THIRDPARTY" // 自研落地页
	AssetTypeWechatApp  = "WECHAT_APP" // 微信小程序
)

// 常量定义 - 小程序类型
const (
	MiniProgramTypeByteApp  = "BYTE_APP"  // 字节小程序
	MiniProgramTypeByteGame = "BYTE_GAME" // 字节小游戏
)

// 常量定义 - 应用类型
const (
	AppTypeAndroid = "Android" // Android应用
	AppTypeIOS     = "IOS"     // iOS应用
	AppTypeDefault = "DEFAULT" // 默认
	AppTypeHarmony = "HARMONY" // 鸿蒙
)

// 常量定义 - 应用创建类型
const (
	AppCreateTypeNormal = "NORMAL" // 普通创建（默认）
	AppCreateTypeUG     = "UG"     // UG创建
)

// 长度限制常量
const (
	MaxQuickAppNameLength        = 20  // 快应用名称最大长度
	MaxQuickAppPackageNameLength = 140 // 快应用包名最大长度
	MaxAppNameLength             = 125 // 应用名称最大长度
	MaxAppPackageNameLength      = 140 // 应用包名最大长度
)

// 长度限制常量
const (
	MaxLandingPageNameLength = 25  // 落地页名称最大长度
	MaxDescriptionLength     = 150 // 落地页描述最大长度
)

type EventManagerAssetsCreateReq struct {
	accessTokenReq
	AdvertiserID     int64             `json:"advertiser_id"`                // 客户ID (必填)
	AssetType        string            `json:"asset_type"`                   // 资产类型 (必填)
	ThirdPartAsset   *ThirdPartAsset   `json:"third_part_asset,omitempty"`   // 三方落地页资产信息
	QuickAppAsset    *QuickAppAsset    `json:"quick_app_asset,omitempty"`    // 快应用资产信息
	AppAsset         *AppAsset         `json:"app_asset,omitempty"`          // 应用信息
	SiteAsset        *SiteAsset        `json:"site_asset,omitempty"`         // 橙子落地页信息
	MiniProgramAsset *MiniProgramAsset `json:"mini_program_asset,omitempty"` // 字节小程序资产信息
}

// SiteAsset 橙子落地页信息
type SiteAsset struct {
	SiteID   int64  `json:"site_id"`   // 橙子建站站点id (橙子落地页必填)
	SiteName string `json:"site_name"` // 橙子建站站点名称 (橙子落地页必填)
}

// MiniProgramAsset 字节小程序资产信息
type MiniProgramAsset struct {
	MiniProgramID   string `json:"mini_program_id"`   // 字节小程序AppID (必填)
	MiniProgramName string `json:"mini_program_name"` // 字节小程序的名称 (必填)
	InstanceID      int64  `json:"instance_id"`       // 字节小程序资产ID (必填)
	MiniProgramType string `json:"mini_program_type"` // 小程序类型 (必填)
}

// QuickAppAsset 快应用资产信息
type QuickAppAsset struct {
	Name        string `json:"name"`         // 快应用名称 (必填)
	PackageName string `json:"package_name"` // 快应用包名 (必填)
}

// AppAsset 应用信息
type AppAsset struct {
	Name          string `json:"name"`                      // 应用名称 (必填)
	PackageName   string `json:"package_name"`              // 应用包名 (必填)
	DownloadURL   string `json:"download_url"`              // 应用下载链接 (必填)
	AppID         int64  `json:"app_id,omitempty"`          // 应用ID (Android应用必填)
	PackageID     string `json:"package_id,omitempty"`      // 母包ID (Android应用必填)
	AppType       string `json:"app_type"`                  // 应用类型 (必填)
	AppCreateType string `json:"app_create_type,omitempty"` // 应用创建类型，默认NORMAL
}

// ThirdPartAsset 三方落地页资产信息
type ThirdPartAsset struct {
	Name        string `json:"name"`        // 落地页名称
	Description string `json:"description"` // 落地页描述信息
}

// Validate 验证三方落地页资产信息
func (t *ThirdPartAsset) Validate() error {
	// 验证落地页名称
	if t.Name == "" {
		return errors.New("name为必填")
	}
	if len(t.Name) > MaxLandingPageNameLength {
		return errors.New("落地页名称长度不能超过25个字符")
	}

	// 验证落地页描述
	if t.Description != "" && len(t.Description) > MaxDescriptionLength {
		return errors.New("落地页描述信息长度不能超过150个字符")
	}

	return nil
}

func (a *EventManagerAssetsCreateReq) Format() {
	a.accessTokenReq.Format()
}

// isValidAssetType 验证资产类型是否有效
func isValidAssetType(assetType string) bool {
	validTypes := map[string]bool{
		AssetTypeApp:         true,
		AssetTypeTetris:      true,
		AssetTypeThird:       true,
		AssetTypeMiniProgram: true,
		AssetTypeOffline:     true,
		AssetTypeOther:       true,
		AssetTypeQuickApp:    true,
		AssetTypeSite:        true,
	}
	return validTypes[assetType]
}

func (a *EventManagerAssetsCreateReq) Validate() (err error) {
	if validateErr := a.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	// 1. 验证客户ID
	if a.AdvertiserID == 0 {
		err = errors.New("advertiser_id为必填")
		return
	}

	// 2. 验证资产类型
	if a.AssetType == "" {
		err = errors.New("asset_type为必填")
		return
	}

	if !isValidAssetType(a.AssetType) {
		err = errors.New("asset_type值无效，允许值：APP、TETRIS_EXTERNAL、THIRD_EXTERNAL、MINI_PROGRAMME、OFFLINE_EVENT、OTHER、QUICK_APP")
		return
	}

	// 3. 根据资产类型验证对应资产信息
	switch a.AssetType {
	case AssetTypeSite:
		return a.validateSiteAsset()
	case AssetTypeMiniProgram:
		return a.validateMiniProgramAsset()
	case AssetTypeQuickApp:
		return a.validateQuickAppAsset()
	case AssetTypeApp:
		return a.validateAppAsset()
	case AssetTypeThird:
		return a.validateThirdPartAsset()
	default:
		return nil
	}
}

// validateSiteAsset 验证橙子落地页信息
func (a *EventManagerAssetsCreateReq) validateSiteAsset() error {
	if a.SiteAsset == nil {
		return errors.New("asset_type为SITE时，site_asset为必填")
	}

	if a.SiteAsset.SiteID == 0 {
		return errors.New("橙子落地页site_id为必填")
	}

	if a.SiteAsset.SiteName == "" {
		return errors.New("橙子落地页site_name为必填")
	}

	return nil
}

// validateMiniProgramAsset 验证字节小程序资产信息
func (a *EventManagerAssetsCreateReq) validateMiniProgramAsset() error {
	if a.MiniProgramAsset == nil {
		return errors.New("asset_type为MINI_PROGRAMME时，mini_program_asset为必填")
	}

	// 验证小程序ID
	if a.MiniProgramAsset.MiniProgramID == "" {
		return errors.New("mini_program_id为必填")
	}

	// 验证小程序名称
	if a.MiniProgramAsset.MiniProgramName == "" {
		return errors.New("mini_program_name为必填")
	}

	// 验证资产实例ID
	if a.MiniProgramAsset.InstanceID == 0 {
		return errors.New("instance_id为必填")
	}

	// 验证小程序类型
	if a.MiniProgramAsset.MiniProgramType == "" {
		return errors.New("mini_program_type为必填")
	}

	if a.MiniProgramAsset.MiniProgramType != MiniProgramTypeByteApp &&
		a.MiniProgramAsset.MiniProgramType != MiniProgramTypeByteGame {
		return errors.New("mini_program_type值无效，允许值：BYTE_APP、BYTE_GAME")
	}

	return nil
}

// validateQuickAppAsset 验证快应用资产信息
func (a *EventManagerAssetsCreateReq) validateQuickAppAsset() error {
	if a.QuickAppAsset == nil {
		return errors.New("asset_type为QUICK_APP时，quick_app_asset为必填")
	}

	if a.QuickAppAsset.Name == "" {
		return errors.New("快应用名称name为必填")
	}
	if len(a.QuickAppAsset.Name) > MaxQuickAppNameLength {
		return errors.New("快应用名称长度不能超过20个字符")
	}

	if a.QuickAppAsset.PackageName == "" {
		return errors.New("快应用包名package_name为必填")
	}
	if len(a.QuickAppAsset.PackageName) > MaxQuickAppPackageNameLength {
		return errors.New("快应用包名长度不能超过140个字符")
	}

	return nil
}

// validateAppAsset 验证应用信息
func (a *EventManagerAssetsCreateReq) validateAppAsset() error {
	if a.AppAsset == nil {
		return errors.New("asset_type为APP时，app_asset为必填")
	}

	// 验证应用名称
	if a.AppAsset.Name == "" {
		return errors.New("应用名称name为必填")
	}
	if len(a.AppAsset.Name) > MaxAppNameLength {
		return errors.New("应用名称长度不能超过125个字符")
	}

	// 验证应用包名
	if a.AppAsset.PackageName == "" {
		return errors.New("应用包名package_name为必填")
	}
	if len(a.AppAsset.PackageName) > MaxAppPackageNameLength {
		return errors.New("应用包名长度不能超过140个字符")
	}

	// 验证下载链接
	if a.AppAsset.DownloadURL == "" {
		return errors.New("应用下载链接download_url为必填")
	}

	// 验证应用类型
	if a.AppAsset.AppType == "" {
		return errors.New("应用类型app_type为必填")
	}
	if a.AppAsset.AppType != AppTypeAndroid && a.AppAsset.AppType != AppTypeIOS {
		return errors.New("app_type值无效，允许值：Android、IOS")
	}

	// Android应用特有验证
	if a.AppAsset.AppType == AppTypeAndroid {
		if a.AppAsset.AppID == 0 {
			return errors.New("Android应用app_id为必填")
		}
		if a.AppAsset.PackageID == "" {
			return errors.New("Android应用package_id为必填")
		}
	}

	// 设置默认值
	if a.AppAsset.AppCreateType == "" {
		a.AppAsset.AppCreateType = AppCreateTypeNormal
	}
	if a.AppAsset.AppCreateType != AppCreateTypeNormal && a.AppAsset.AppCreateType != AppCreateTypeUG {
		return errors.New("app_create_type值无效，允许值：NORMAL、UG")
	}

	return nil
}

// validateThirdPartAsset 验证三方落地页资产信息
func (a *EventManagerAssetsCreateReq) validateThirdPartAsset() error {
	if a.ThirdPartAsset == nil {
		return errors.New("asset_type为THIRD_EXTERNAL时，third_part_asset为必填")
	}
	return a.ThirdPartAsset.Validate()
}

func (a *EventManagerAssetsCreateReq) GetHeaders() headersMap {
	headers := a.accessTokenReq.GetHeaders()
	headers.Json()
	return headers
}

type EventManagerAssetsCreateResp struct {
	AssetId int64 `json:"asset_id"` // 资产ID
}

type EventAssetsListReq struct {
	accessTokenReq
	AdvertiserID int64        `json:"advertiser_id"`       // 投放账户id (必填)
	Filtering    *AssetFilter `json:"filtering,omitempty"` // 过滤条件
	PageInfoReq
}

func (p *EventAssetsListReq) Format() {
	p.accessTokenReq.Format()
}

// AssetFilter 资产过滤条件
type AssetFilter struct {
	AssetIDs        []int64 `json:"asset_ids,omitempty"`         // 资产id列表，最大100
	AssetType       string  `json:"asset_type,omitempty"`        // 资产类型
	ModifyStartTime string  `json:"modify_start_time,omitempty"` // 资产修改开始时间 YYYY-MM-DD
	ModifyEndTime   string  `json:"modify_end_time,omitempty"`   // 资产修改结束时间 YYYY-MM-DD
}

// 分页限制常量
const (
	MaxAssetIDsCount = 100    // 资产ID列表最大数量
	DefaultPage      = 1      // 默认页数
	DefaultPageSize  = 10     // 默认页面大小
	MaxPageSize      = 100    // 最大页面大小
	MaxPage          = 999999 // 最大页数
)

// Validate 验证资产查询参数
func (p *EventAssetsListReq) Validate() error {
	// 1. 验证客户ID
	if p.AdvertiserID == 0 {
		return errors.New("advertiser_id为必填")
	}

	// 2. 设置分页默认值
	p.setPageDefaults()

	// 3. 验证过滤条件
	if err := p.validateFiltering(); err != nil {
		return err
	}

	// 5. 验证分页参数
	if err := p.validatePageParams(); err != nil {
		return err
	}

	return nil
}

// setPageDefaults 设置分页默认值
func (p *EventAssetsListReq) setPageDefaults() {
	if p.Page <= 0 {
		p.Page = DefaultPage
	}
	if p.PageSize <= 0 {
		p.PageSize = DefaultPageSize
	}
}

// validateFiltering 验证过滤条件
func (p *EventAssetsListReq) validateFiltering() error {
	if p.Filtering == nil {
		return nil
	}

	// 验证资产ID列表长度
	if len(p.Filtering.AssetIDs) > MaxAssetIDsCount {
		return errors.New("asset_ids列表长度不能超过100")
	}

	// 验证资产类型
	if p.Filtering.AssetType != "" && !isValidAssetQueryType(p.Filtering.AssetType) {
		return errors.New("asset_type值无效，允许值：THIRD_EXTERNAL、TETRIS_EXTERNAL、APP、QUICK_APP、MINI_PROGRAMME")
	}
	// 检查是否只传了其中一个时间
	if (p.Filtering.ModifyStartTime != "" && p.Filtering.ModifyEndTime == "") ||
		(p.Filtering.ModifyStartTime == "" && p.Filtering.ModifyEndTime != "") {
		return errors.New("开始时间和结束时间必须同时传入")
	}

	// 如果两个时间都为空，跳过验证
	if p.Filtering.ModifyStartTime == "" && p.Filtering.ModifyEndTime == "" {
		return nil
	}

	// 验证时间格式
	start, err := time.Parse(DateFormat, p.Filtering.ModifyStartTime)
	if err != nil {
		return errors.New("modify_start_time格式错误，应为YYYY-MM-DD")
	}

	end, err := time.Parse(DateFormat, p.Filtering.ModifyEndTime)
	if err != nil {
		return errors.New("modify_end_time格式错误，应为YYYY-MM-DD")
	}

	// 验证开始时间 <= 结束时间
	if start.After(end) {
		return errors.New("开始时间不能大于结束时间")
	}

	return nil
}

// validatePageParams 验证分页参数
func (p *EventAssetsListReq) validatePageParams() error {
	if p.Page > MaxPage {
		return errors.New("page不能超过999999")
	}

	if p.PageSize > MaxPageSize {
		return errors.New("page_size不能超过100")
	}

	return nil
}

// isValidAssetQueryType 验证资产类型是否有效
func isValidAssetQueryType(assetType string) bool {
	validTypes := map[string]bool{
		AssetTypeThird:       true,
		AssetTypeTetris:      true,
		AssetTypeApp:         true,
		AssetTypeQuickApp:    true,
		AssetTypeMiniProgram: true,
	}
	return validTypes[assetType]
}

// AssetInfo 资产信息
type AssetInfo struct {
	AssetType    string `json:"asset_type"`               // 资产类型
	AssetID      int64  `json:"asset_id"`                 // 资产id
	AssetName    string `json:"asset_name"`               // 资产名称
	ShareType    string `json:"share_type,omitempty"`     // 资产来源
	CreateTime   string `json:"create_time"`              // 资产创建时间，格式 YYYY-MM-DD HH:MM:SS
	ModifyTime   string `json:"modify_time"`              // 资产修改时间，格式 YYYY-MM-DD HH:MM:SS
	AppForceInfo string `json:"app_force_info,omitempty"` // 异常资产限制倒计时
}

// EventAssetsListResp 资产列表响应
type EventAssetsListResp struct {
	AssetList []*AssetInfo `json:"asset_list"` // 账户下的资产列表，不支持查询已删除的资产
	PageInfoContainerResp
}

// 常量定义 - 资产来源
const (
	ShareTypeMyCreations = "MY_CREATIONS"  // 我创建的
	ShareTypeSharing     = "SHARING"       // 共享中
	ShareTypeExpired     = "SHATE_EXPIRED" // 共享失效
)

// EventAssetsDetailReq 获取已创建资产详情（新）
type EventAssetsDetailReq struct {
	accessTokenReq
	AdvertiserID int64   `json:"advertiser_id"` // 客户id (必填)
	AssetIDs     []int64 `json:"asset_ids"`     // 资产id列表 (必填)
}

// 常量定义
const (
	MaxAssetIDsLength = 50 // 资产id列表最大长度
)

func (p *EventAssetsDetailReq) Format() {
	p.accessTokenReq.Format()
}

// Validate 验证资产详情查询参数
func (p *EventAssetsDetailReq) Validate() error {
	// 1. 验证客户ID
	if p.AdvertiserID == 0 {
		return errors.New("advertiser_id为必填")
	}

	// 2. 验证资产id列表
	if len(p.AssetIDs) == 0 {
		return errors.New("asset_ids为必填")
	}

	if len(p.AssetIDs) > MaxAssetIDsLength {
		return errors.New("asset_ids列表长度不能超过50")
	}

	if validateErr := p.accessTokenReq.Validate(); validateErr != nil {
		return validateErr
	}

	return nil
}

// EventAssetsDetailResp 响应
type EventAssetsDetailResp struct {
	AssetList []*AssetDetailInfo `json:"asset_list"` // 资产列表
}

// AssetDetailInfo 资产信息
type AssetDetailInfo struct {
	AssetID             int64  `json:"asset_id"`                         // 资产id，不会返回已删除资产
	AssetName           string `json:"asset_name"`                       // 资产名称，不会返回已删除资产
	AssetType           string `json:"asset_type"`                       // 资产类型
	SiteID              int64  `json:"site_id,omitempty"`                // 橙子落地页站点id
	AppType             string `json:"app_type,omitempty"`               // 应用类型
	AppID               int64  `json:"app_id,omitempty"`                 // 应用ID
	AppName             string `json:"app_name,omitempty"`               // 应用名称
	PackageID           string `json:"package_id,omitempty"`             // 应用包id
	PackageName         string `json:"package_name,omitempty"`           // 应用包名
	DownloadURL         string `json:"download_url,omitempty"`           // 应用下载链接
	QuickAppID          int64  `json:"quick_app_id,omitempty"`           // 快应用id
	QuickAppPackageName string `json:"quick_app_package_name,omitempty"` // 快应用包名
	MicroAppID          string `json:"micro_app_id,omitempty"`           // 小程序appid
	MicroAppInstanceID  int64  `json:"micro_app_instance_id,omitempty"`  // 小程序资产id
}

// EventManagerAvailableEventsGetReq 获取可创建事件列表
type EventManagerAvailableEventsGetReq struct {
	accessTokenReq
	AdvertiserID int64 `json:"advertiser_id"` // 客户ID (必填)
	AssetID      int64 `json:"asset_id"`      // 资产ID (必填)
}

func (p *EventManagerAvailableEventsGetReq) Format() {
	p.accessTokenReq.Format()
}

// Validate 验证资产删除参数
func (p *EventManagerAvailableEventsGetReq) Validate() error {
	if p.AdvertiserID == 0 {
		return errors.New("advertiser_id为必填")
	}
	if p.AssetID == 0 {
		return errors.New("asset_id为必填")
	}
	if validateErr := p.accessTokenReq.Validate(); validateErr != nil {
		return validateErr
	}
	return nil
}

// EventManagerAvailableEventsGetResp 返回
type EventManagerAvailableEventsGetResp struct {
	EventConfigs []*EventConfig `json:"event_configs"` // 可创建事件列表
}

// EventConfig
type EventConfig struct {
	EventID     int64            `json:"event_id"`             // 事件ID
	EventType   string           `json:"event_type"`           // 事件类型
	EventCnName string           `json:"event_cn_name"`        // 事件中文名称
	Description string           `json:"description"`          // 事件描述信息
	TrackTypes  []string         `json:"track_types"`          // 事件回传方式列表
	Properties  []*EventProperty `json:"properties,omitempty"` // 事件的附加属性
}

// EventProperty 事件附加属性
type EventProperty struct {
	Field        string            `json:"field"`                 // 附加属性英文名称
	FieldName    string            `json:"field_name"`            // 附加属性中文名称
	VariableType string            `json:"variable_type"`         // 附加属性值类型
	EnumValue    map[string]string `json:"enum_value,omitempty"`  // 附加属性枚举值
	Unit         string            `json:"unit,omitempty"`        // 附加属性单位
	Description  string            `json:"description,omitempty"` // 附加属性描述
}

// 常量定义 - 事件回传方式
const (
	// 落地页支持的回传方式
	TrackTypeJSSDK       = "JSSDK"        // JS埋码
	TrackTypeExternalAPI = "EXTERNAL_API" // API回传
	TrackTypeXPath       = "XPATH"        // XPath圈选

	// 应用支持的回传方式
	TrackTypeApplicationAPI = "APPLICATION_API" // 应用API
	TrackTypeApplicationSDK = "APPLICATION_SDK" // 应用SDK

	// 快应用支持的回传方式
	TrackTypeQuickAppAPI = "QUICK_APP_API" // 快应用API

	// 小程序支持的回传方式
	TrackTypeMiniProgramSDK = "MINI_PROGRAMME_SDK" // 小程序SDK
	TrackTypeMiniProgramAPI = "MINI_PROGRAMME_API" // 小程序API
)

type EventManagerEventsCreateReq struct {
	accessTokenReq
	AdvertiserID int64    `json:"advertiser_id"` // 客户ID (必填)
	AssetID      int64    `json:"asset_id"`      // 资产ID (必填)
	EventID      int64    `json:"event_id"`      // 事件ID (必填)
	TrackTypes   []string `json:"track_types"`   // 事件回传方式列表 (必填)
}

func (p *EventManagerEventsCreateReq) Format() {
	p.accessTokenReq.Format()
}

func (p *EventManagerEventsCreateReq) GetHeaders() headersMap {
	headers := p.accessTokenReq.GetHeaders()
	headers.Json()
	return headers
}

// 所有有效回传方式列表
var ValidTrackTypes = []string{
	TrackTypeJSSDK,
	TrackTypeExternalAPI,
	TrackTypeXPath,
	TrackTypeApplicationAPI,
	TrackTypeApplicationSDK,
	TrackTypeQuickAppAPI,
}

// Validate 验证事件启用参数
func (p *EventManagerEventsCreateReq) Validate() error {
	// 1. 验证客户ID
	if p.AdvertiserID == 0 {
		return errors.New("advertiser_id为必填")
	}

	// 2. 验证资产ID
	if p.AssetID == 0 {
		return errors.New("asset_id为必填")
	}

	// 3. 验证事件ID
	if p.EventID == 0 {
		return errors.New("event_id为必填")
	}

	// 4. 验证回传方式列表
	if len(p.TrackTypes) == 0 {
		return errors.New("track_types为必填")
	}

	// 5. 验证每个回传方式是否有效
	validTrackTypeMap := make(map[string]bool)
	for _, tt := range ValidTrackTypes {
		validTrackTypeMap[tt] = true
	}

	for _, tt := range p.TrackTypes {
		if !validTrackTypeMap[tt] {
			return errors.New("track_types包含无效值，允许值：JSSDK、EXTERNAL_API、XPATH、APPLICATION_API、APPLICATION_SDK、QUICK_APP_API")
		}
	}
	if validateErr := p.accessTokenReq.Validate(); validateErr != nil {
		return validateErr
	}
	return nil
}

// 资产下创建事件 返回
type EventManagerEventsCreateResp struct {
}

// 获取资产下已创建事件列表
// https://open.oceanengine.com/labels/7/docs/1709793086075972?origin=left_nav
type EventManagerEventConfigsGetReq struct {
	accessTokenReq
	AdvertiserID int64  `json:"advertiser_id"`       // 客户ID (必填)
	AssetID      int64  `json:"asset_id"`            // 资产ID (必填)
	SortType     string `json:"sort_type,omitempty"` // 创建时间排序方式，默认ASC
}

func (p *EventManagerEventConfigsGetReq) Format() {
	p.accessTokenReq.Format()
}

// 常量定义 - 排序方式
const (
	SortTypeAsc  = "ASC"  // 升序（默认）
	SortTypeDesc = "DESC" // 降序
)

// Validate 验证事件列表查询参数
func (p *EventManagerEventConfigsGetReq) Validate() error {
	// 1. 验证客户ID
	if p.AdvertiserID == 0 {
		return errors.New("advertiser_id为必填")
	}

	// 2. 验证资产ID
	if p.AssetID == 0 {
		return errors.New("asset_id为必填")
	}

	// 3. 设置默认排序方式
	if p.SortType == "" {
		p.SortType = SortTypeAsc
	}

	// 4. 验证排序方式
	if p.SortType != SortTypeAsc && p.SortType != SortTypeDesc {
		return errors.New("sort_type值无效，允许值：ASC、DESC")
	}

	if validateErr := p.accessTokenReq.Validate(); validateErr != nil {
		return validateErr
	}
	return nil
}

// EventManagerEventConfigsGetResp 事件配置列表响应
type EventManagerEventConfigsGetResp struct {
	EventConfigs []*EventManagerEventConfig `json:"event_configs"` // 已创建事件列表
}

// EventManagerEventConfig 事件配置
type EventManagerEventConfig struct {
	EventID                  int64              `json:"event_id"`                            // 事件ID
	EventType                string             `json:"event_type"`                          // 事件类型
	EventCnName              string             `json:"event_cn_name"`                       // 事件中文名称
	AttributionConfiguration *AttributionConfig `json:"attribution_configuration,omitempty"` // 属性配置
	DebuggingStatus          string             `json:"debugging_status,omitempty"`          // 激活免联调状态
	TrackTypes               []string           `json:"track_types"`                         // 事件回传方式列表
	CreateTime               string             `json:"create_time"`                         // 事件创建时间
	Properties               []*EventProperty   `json:"properties,omitempty"`                // 事件的附加属性
}

// AttributionConfig 属性配置
type AttributionConfig struct {
	AttributionWindow int64 `json:"attribution_window"` // 归因窗口
}

// 常量定义 - 调试状态
const (
	DebuggingStatusActive   = "Active"   // 已激活
	DebuggingStatusInactive = "Inactive" // 未激活
)

type EventManagerOptimizedGoalGetReq struct {
	accessTokenReq
	AdvertiserID       int64  `json:"advertiser_id"`                   // 客户id (必填)
	LandingType        string `json:"landing_type"`                    // 营销目的 (必填)
	AdType             string `json:"ad_type"`                         // 营销类型 (必填)
	AssetType          string `json:"asset_type"`                      // 资产类型 (必填)
	MultiAssetType     string `json:"multi_asset_type,omitempty"`      // 多投放载体
	AssetID            int64  `json:"asset_id,omitempty"`              // 事件管理资产id
	PackageName        string `json:"package_name,omitempty"`          // 应用包名称
	AppType            string `json:"app_type,omitempty"`              // 应用类型
	AppPromotionType   string `json:"app_promotion_type,omitempty"`    // 子目标
	MarketingGoal      string `json:"marketing_goal,omitempty"`        // 营销场景
	QuickAppID         int64  `json:"quick_app_id,omitempty"`          // 快应用资产id
	DeliveryMode       string `json:"delivery_mode,omitempty"`         // 投放模式
	MiniProgramID      string `json:"mini_program_id,omitempty"`       // 字节小程序资产id
	DpaAdtype          string `json:"dpa_adtype,omitempty"`            // dpa营销类型
	MicroPromotionType string `json:"micro_promotion_type,omitempty"`  // 小程序类型
	MicroAppInstanceID int64  `json:"micro_app_instance_id,omitempty"` // 小程序资产id
	DeliveryType       string `json:"delivery_type,omitempty"`         // 投放类型
}

func (p *EventManagerOptimizedGoalGetReq) Format() {
	p.accessTokenReq.Format()
}

// 常量定义 - 多投放载体
const (
	MultiAssetTypeOrangeAndAweme = "ORANGE_AND_AWEME" // 建站落地页和抖音主页
)

// 常量定义 - DPA营销类型
const (
	DpaAdtypeApp  = "DPA_APP"  // 应用下载
	DpaAdtypeLink = "DPA_LINK" // 落地页
)

// 常量定义 - 小程序类型
const (
	MicroPromotionTypeByteApp    = "BYTE_APP"    // 字节小程序
	MicroPromotionTypeByteGame   = "BYTE_GAME"   // 字节小游戏
	MicroPromotionTypeWechatApp  = "WECHAT_APP"  // 微信小程序
	MicroPromotionTypeWechatGame = "WECHAT_GAME" // 微信小游戏
	MicroPromotionTypeAweme      = "AWEME"       // 抖音号
)

// Validate 验证事件资产查询参数
func (p *EventManagerOptimizedGoalGetReq) Validate() error {
	// 1. 验证必填字段
	if p.AdvertiserID == 0 {
		return errors.New("advertiser_id为必填")
	}
	if p.LandingType == "" {
		return errors.New("landing_type为必填")
	}
	if p.AdType == "" {
		return errors.New("ad_type为必填")
	}
	if p.AssetType == "" {
		return errors.New("asset_type为必填")
	}

	// 2. 验证营销目的
	if !isValidLandingType(p.LandingType) {
		return errors.New("landing_type值无效，允许值：APP、DPA、LINK、MICRO_GAME、NATIVE_ACTION、QUICK_APP、SHOP")
	}

	// 3. 验证营销类型
	if !isValidAdType(p.AdType) {
		return errors.New("ad_type值无效，允许值：ALL、SEARCH")
	}

	// 4. 验证资产类型
	if !isValidAssetTypeOptimizedGoal(p.AssetType) {
		return errors.New("asset_type值无效，允许值：APP、AWEME、ENTERPRISE、MICRO_APP、ORANGE、QUICK_APP、THIRDPARTY、WECHAT_APP")
	}

	// 5. 验证多投放载体
	if p.MultiAssetType != "" && p.MultiAssetType != MultiAssetTypeOrangeAndAweme {
		return errors.New("multi_asset_type值无效，允许值：ORANGE_AND_AWEME")
	}

	// 6. 应用类型验证
	if p.AssetType == AssetTypeApp && p.AppType == "" {
		return errors.New("asset_type为APP时，app_type为必填")
	}
	if p.AppType != "" && !isValidAppType(p.AppType) {
		return errors.New("app_type值无效，允许值：ANDROID、DEFAULT、IOS、HARMONY")
	}

	// 7. 子目标验证
	if p.AppPromotionType != "" && !isValidAppPromotionType(p.AppPromotionType) {
		return errors.New("app_promotion_type值无效，允许值：DOWNLOAD、LAUNCH、RESERVE")
	}

	// 8. 营销场景验证
	if p.MarketingGoal != "" && !isValidMarketingGoal(p.MarketingGoal) {
		return errors.New("marketing_goal值无效，允许值：LIVE、VIDEO_AND_IMAGE")
	}

	// 9. 投放模式验证
	if p.DeliveryMode != "" && !isValidDeliveryMode(p.DeliveryMode) {
		return errors.New("delivery_mode值无效，允许值：MANUAL、PROCEDURAL")
	}

	// 10. DPA营销类型验证
	if p.DpaAdtype != "" && !isValidDpaAdtype(p.DpaAdtype) {
		return errors.New("dpa_adtype值无效，允许值：DPA_APP、DPA_LINK")
	}

	// 11. 小程序类型验证（landing_type = MICRO_GAME时有效且必填）
	if p.LandingType == LandingTypeMicroGame && p.MicroPromotionType == "" {
		return errors.New("landing_type为MICRO_GAME时，micro_promotion_type为必填")
	}
	if p.MicroPromotionType != "" && !isValidMicroPromotionType(p.MicroPromotionType) {
		return errors.New("micro_promotion_type值无效，允许值：BYTE_APP、BYTE_GAME、WECHAT_APP、WECHAT_GAME、AWEME")
	}

	// 12. 验证字节小程序资产id
	if p.AssetType == AssetTypeMicroApp && p.MiniProgramID == "" {
		return errors.New("asset_type为MICRO_APP时，mini_program_id为必填")
	}

	// 13. 验证投放类型
	if p.DeliveryType != "" && !isValidDeliveryType(p.DeliveryType) {
		return errors.New("delivery_type值无效，允许值：DURATION、NORMAL")
	}

	if validateErr := p.accessTokenReq.Validate(); validateErr != nil {
		return validateErr
	}

	return nil
}

// 辅助验证函数
func isValidLandingType(landingType string) bool {
	validTypes := map[string]bool{
		LandingTypeApp: true, LandingTypeDPA: true, LandingTypeLink: true,
		LandingTypeMicroGame: true, LandingTypeNativeAction: true,
		LandingTypeQuickApp: true, LandingTypeShop: true,
	}
	return validTypes[landingType]
}

func isValidAdType(adType string) bool {
	return adType == AdTypeAll || adType == AdTypeSearch
}

func isValidAssetTypeOptimizedGoal(assetType string) bool {
	validTypes := map[string]bool{
		AssetTypeApp: true, AssetTypeAweme: true, AssetTypeEnterprise: true,
		AssetTypeMicroApp: true, AssetTypeOrange: true, AssetTypeQuickApp: true,
		AssetTypeThirdparty: true, AssetTypeWechatApp: true,
	}
	return validTypes[assetType]
}

func isValidAppType(appType string) bool {
	validTypes := map[string]bool{
		AppTypeAndroid: true, AppTypeIOS: true, AppTypeDefault: true, AppTypeHarmony: true,
	}
	return validTypes[appType]
}

func isValidAppPromotionType(promotionType string) bool {
	return promotionType == AppPromotionTypeDownload ||
		promotionType == AppPromotionTypeLaunch ||
		promotionType == AppPromotionTypeReserve
}

func isValidMarketingGoal(goal string) bool {
	return goal == MarketingGoalLive || goal == MarketingGoalVideoImage
}

func isValidDeliveryMode(mode string) bool {
	return mode == DeliveryModeManual || mode == DeliveryModeProcedural
}

func isValidDpaAdtype(dpaAdtype string) bool {
	return dpaAdtype == DpaAdtypeApp || dpaAdtype == DpaAdtypeLink
}

func isValidMicroPromotionType(microType string) bool {
	validTypes := map[string]bool{
		MicroPromotionTypeByteApp: true, MicroPromotionTypeByteGame: true,
		MicroPromotionTypeWechatApp: true, MicroPromotionTypeWechatGame: true,
		MicroPromotionTypeAweme: true,
	}
	return validTypes[microType]
}

func isValidDeliveryType(deliveryType string) bool {
	return deliveryType == DeliveryTypeDuration || deliveryType == DeliveryTypeNormal
}

type EventManagerOptimizedGoalGetResp struct {
	Goals    []*OptimizationGoal `json:"goals"`               // 优化目标数据
	AssetIDs []int64             `json:"asset_ids,omitempty"` // 事件管理的资产id
}

// OptimizationGoal 优化目标
type OptimizationGoal struct {
	OptimizationName   string                  `json:"optimization_name"`         // 事件名称 (必填)
	AssetTypes         []string                `json:"asset_types"`               // 资产类型
	HistoryBack        bool                    `json:"history_back"`              // 历史有无回传 (必填)
	TwentyFourHourBack bool                    `json:"twenty_four_hour_back"`     // 24小时历史有无回传
	ExternalAction     string                  `json:"external_action,omitempty"` // 优化目标
	ValueType          string                  `json:"value_type,omitempty"`      // 是否设置差异价值
	DeepGoals          []*DeepOptimizationGoal `json:"deep_goals,omitempty"`      // 深度优化目标列表
	NeedAppForce       bool                    `json:"need_app_force,omitempty"`  // 仅针对 landing_type = APP 的筛选下进行返回
}

// DeepOptimizationGoal 深度优化目标
type DeepOptimizationGoal struct {
	OptimizationName   string   `json:"optimization_name"`              // 事件名称 (必填)
	HistoryBack        bool     `json:"history_back"`                   // 历史有无回传
	TwentyFourHourBack bool     `json:"twenty_four_hour_back"`          // 24小时历史有无回传
	DeepExternalAction string   `json:"deep_external_action,omitempty"` // 深度优化目标
	AssetTypes         []string `json:"asset_types"`                    // 资产类型
}
