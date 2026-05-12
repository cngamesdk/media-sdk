package model

import "errors"

// AppReleaseDetailReq 获取新版应用发布详情请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/release/detail
type AppReleaseDetailReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID，必填
	PackageId    int64 `json:"package_id"`    // 应用包ID，支持母包ID或分包ID，必填
}

func (receiver *AppReleaseDetailReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppReleaseDetailReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.PackageId <= 0 {
		err = errors.New("package_id is empty")
		return
	}
	return
}

// PermissionDetail 权限详情
type PermissionDetail struct {
	Description string `json:"description"` // 权限描述
	Id          int    `json:"id"`          // 权限ID
	Name        string `json:"name"`        // 权限名称
}

// SensitivePermission 敏感权限
type SensitivePermission struct {
	Group             string             `json:"group"`              // 权限分组
	PermissionDetails []PermissionDetail `json:"permission_details"` // 权限详情列表
}

// AndroidPermission 安卓权限信息
type AndroidPermission struct {
	OrdinaryPermissions  []PermissionDetail    `json:"ordinary_permissions"`  // 普通权限列表
	PermissionIds        []string              `json:"permission_ids"`        // 权限ID列表
	SensitivePermissions []SensitivePermission `json:"sensitive_permissions"` // 敏感权限列表
}

// AppReleaseDetailResp 获取新版应用发布详情响应数据（仅data部分）
type AppReleaseDetailResp struct {
	AccountAppId          int               `json:"account_app_id"`         // 账号业务线
	AccountId             int64             `json:"account_id"`             // 创建应用的账号ID
	AppDetailImg          string            `json:"app_detail_img"`         // 应用详情图片
	AppIconUrl            string            `json:"app_icon_url"`           // 应用图标链接
	AppId                 int64             `json:"app_id"`                 // 应用ID
	AppPrivacyUrl         string            `json:"app_privacy_url"`        // 应用隐私政策链接
	CertMd5               string            `json:"cert_md_5"`              // 签名MD5
	ChannelId             string            `json:"channel_id"`             // 渠道号(分包号)
	Compatibility         string            `json:"compatibility"`          // 兼容性
	CreateTime            int64             `json:"create_time"`            // 创建时间
	Developer             string            `json:"developer"`              // 开发者
	DocumentNumber        string            `json:"document_number"`        // 证件号码
	EffectStatus          int               `json:"effect_status"`          // 生效状态
	FunctionIntroduction  string            `json:"function_introduction"`  // 应用功能介绍
	GlobalAppId           int64             `json:"global_app_id"`          // 全局应用ID
	IosAppId              string            `json:"ios_app_id"`             // 解析出的iosAppID
	IosPageId             string            `json:"ios_page_id"`            // iOS15自定义产品页
	Md5                   string            `json:"md_5"`                   // MD5
	OfflineAppStores      string            `json:"offline_app_stores"`     // 下架的应用商店
	PackageId             int64             `json:"package_id"`             // 应用包ID
	PackageName           string            `json:"package_name"`           // 应用包名
	PackageSize           int64             `json:"package_size"`           // 应用包大小
	ParentPackageId       int64             `json:"parent_package_id"`      // 分包的母包ID
	PermissionInformation []int             `json:"permission_information"` // 权限信息ID列表
	Permissions           AndroidPermission `json:"permissions"`            // 安卓权限信息
	Platform              string            `json:"platform"`               // 平台：android或ios
	PutStatus             int               `json:"put_status"`             // 投放状态
	RealAppName           string            `json:"real_app_name"`          // 应用名称
	RealAppVersion        string            `json:"real_app_version"`       // 应用版本信息
	RecordNumber          string            `json:"record_number"`          // 备案号
	ReviewDetail          string            `json:"review_detail"`          // 审核详情
	ReviewStatus          int               `json:"review_status"`          // 审核状态
	UpdateTime            int64             `json:"update_time"`            // 更新时间
	Url                   string            `json:"url"`                    // 应用下载地址
	UseSdk                int               `json:"use_sdk"`                // 是否接入快手广告监测SDK
	VersionCode           int64             `json:"version_code"`           // 应用版本号
}
