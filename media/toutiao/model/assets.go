package model

import "errors"

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
