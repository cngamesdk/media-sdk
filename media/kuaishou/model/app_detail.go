package model

import "errors"

// AppDetailReq 获取应用详情请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/detail
type AppDetailReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID，必填
	PackageId    int64 `json:"package_id"`    // 应用包ID，必填
}

func (receiver *AppDetailReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppDetailReq) Validate() (err error) {
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

// AppDetailAppSource 应用创建者信息
type AppDetailAppSource struct {
	AccountId   int64  `json:"account_id"`   // 应用创建账号ID
	AccountName string `json:"account_name"` // 应用创建账号名称
}

// AppDetailPermissionDetail 权限详情
type AppDetailPermissionDetail struct {
	Description   string   `json:"description"`    // 权限描述
	Id            int      `json:"id"`             // 权限ID
	Name          string   `json:"name"`           // 权限名称
	PermissionIds []string `json:"permission_ids"` // 权限ID列表
}

// AppDetailSensitivePermission 敏感权限
type AppDetailSensitivePermission struct {
	Group             string                      `json:"group"`              // 权限分组
	PermissionDetails []AppDetailPermissionDetail `json:"permission_details"` // 权限详情列表
}

// AppDetailPermissions 应用权限
type AppDetailPermissions struct {
	OrdinaryPermissions  []AppDetailPermissionDetail    `json:"ordinary_permissions"`  // 普通权限列表
	SensitivePermissions []AppDetailSensitivePermission `json:"sensitive_permissions"` // 敏感权限列表
}

// AppDetailResp 获取应用详情响应数据（仅data部分）
type AppDetailResp struct {
	AccountId               int64                `json:"account_id"`                // 创建应用的账号ID
	AppDetailImg            string               `json:"app_detail_img"`            // 应用详情图片
	AppIconUrl              string               `json:"app_icon_url"`              // 应用图标链接
	AppId                   int64                `json:"app_id"`                    // 应用ID
	AppName                 string               `json:"app_name"`                  // 应用名称
	AppPrivacyUrl           string               `json:"app_privacy_url"`           // 应用隐私政策链接
	AppSource               AppDetailAppSource   `json:"app_source"`                // 应用创建者信息
	AppStatus               int                  `json:"app_status"`                // 应用状态：1=审核中 2=审核失败 3=待发布 4=已发布 5=已下架
	ApplyAge                int                  `json:"apply_age"`                 // 使用年龄
	AuditSerialNumber       int64                `json:"audit_serial_number"`       // 审核序列号
	Category                int                  `json:"category"`                  // 应用类型：1=软件 2=游戏
	Compatibility           string               `json:"compatibility"`             // 兼容性
	ContactEmail            string               `json:"contact_email"`             // 联系人邮箱
	ContactName             string               `json:"contact_name"`              // 联系人姓名
	ContactTel              string               `json:"contact_tel"`               // 联系人电话
	CreateFailedTypes       []int                `json:"create_failed_types"`       // 下载失败原因：1=包名不一致 2=敏感权限未填说明
	CreateSource            int                  `json:"create_source"`             // 应用创建来源：1=文件上传 2=链接下载
	Description             string               `json:"description"`               // 应用包备注
	Developer               string               `json:"developer"`                 // 开发者
	DocumentNumber          string               `json:"document_number"`           // 证件号码
	DownloadStatus          int                  `json:"download_status"`           // 应用链接下载状态
	FunctionIntroduction    string               `json:"function_introduction"`     // 功能介绍
	IosAppId                string               `json:"ios_app_id"`                // 解析出的iOS App ID
	Location                string               `json:"location"`                  // 开发者地区
	OfflineAppLetterUrl     string               `json:"offline_app_letter_url"`    // 单机承诺函
	OfflineAppStores        string               `json:"offline_app_stores"`        // 下架的应用商店
	OnlineEarnType          int                  `json:"online_earn_type"`          // 网赚类型：1=是 2=否
	PackageId               int64                `json:"package_id"`                // 应用包ID
	PackageName             string               `json:"package_name"`              // 应用包名
	PackageSize             int64                `json:"package_size"`              // 应用包大小
	Permissions             AppDetailPermissions `json:"permissions"`               // 应用权限
	Platform                string               `json:"platform"`                  // 应用平台：android 或 ios
	PrivacyId               int64                `json:"privacy_id"`                // 隐私ID
	PrivacyType             int                  `json:"privacy_type"`              // 隐私链接类型
	PrivacyUrl              string               `json:"privacy_url"`               // 隐私链接
	PutStatus               int                  `json:"put_status"`                // 投放状态
	RealAppName             string               `json:"real_app_name"`             // 应用名称
	RealAppVersion          string               `json:"real_app_version"`          // 应用版本信息
	RecordCorpLicenseUrl    string               `json:"record_corp_license_url"`   // 备案主体营业执照
	RecordCorpName          string               `json:"record_corp_name"`          // 备案主体名称
	RecordNumber            string               `json:"record_number"`             // 备案号
	ReleaseType             int                  `json:"release_type"`              // 发布类型：1=手动 2=自动
	ReviewDetail            string               `json:"review_detail"`             // 审核详情
	ReviewStatus            int                  `json:"review_status"`             // 审核状态
	SensitivePermissionDesc string               `json:"sensitive_permission_desc"` // 敏感权限用途说明
	ServiceCategory         string               `json:"service_category"`          // 服务类目
	ShareAccountCount       int                  `json:"share_account_count"`       // 应用共享账号个数
	ShareType               int                  `json:"share_type"`                // 共享类型：0=不共享 1=账号 2=主体
	SourceType              int                  `json:"source_type"`               // 应用来源：1=我创建的 2=共享给我的
	TaskId                  int64                `json:"task_id"`                   // 应用链接下载任务ID
	TraceActivation         int                  `json:"trace_activation"`          // 转化追踪状态：0=未联调 1=已联调
	UpdateTime              int64                `json:"update_time"`               // 更新时间(毫秒)
	Url                     string               `json:"url"`                       // 应用下载地址
	UseSdk                  int                  `json:"use_sdk"`                   // 是否接入快手广告监测SDK：0=未接入 1=已接入
	VersionCode             int64                `json:"version_code"`              // 应用版本号
	VersionName             string               `json:"version_name"`              // 应用版本名称
	AppRecordScreenshotUrl  string               `json:"app_record_screenshot_url"` // APP备案截图
	NetworkType             int                  `json:"network_type"`              // 网络类型：1=联网 2=单机
	ApkDownloadUrl          string               `json:"apk_download_url"`          // 应用链接下载地址
}
