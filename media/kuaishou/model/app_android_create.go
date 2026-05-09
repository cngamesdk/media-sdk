package model

import "errors"

// AppAndroidCreateReq 创建Android应用请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/create/android
type AppAndroidCreateReq struct {
	accessTokenReq
	AdvertiserId   int64                 `json:"advertiser_id"`    // 广告主ID，必填
	AppInfo        AppAndroidAppInfo     `json:"app_info"`         // 应用数据，必填
	AppPrivacyInfo AppAndroidPrivacyInfo `json:"app_privacy_info"` // 隐私声明数据，必填
	PackageInfo    AppAndroidPackageInfo `json:"package_info"`     // 应用包数据，必填
}

// AppAndroidAppInfo 应用数据
type AppAndroidAppInfo struct {
	AppId                  int64  `json:"app_id"`                              // 应用ID，必填
	ReleaseType            int    `json:"release_type"`                        // 发版类型，必填：1=手动发版 2=自动发版
	AppDetailImg           string `json:"app_detail_img,omitempty"`            // 应用详情图(从图片上传接口获取)
	AppIconUrl             string `json:"app_icon_url,omitempty"`              // 应用图标(从图片上传接口获取)
	ApplyAge               *int   `json:"apply_age,omitempty"`                 // 适用年龄段：1=全年龄 2=未成年 3=成年
	Category               *int   `json:"category,omitempty"`                  // 应用类别：1=软件 2=游戏
	ContactEmail           string `json:"contact_email,omitempty"`             // 联系邮箱
	ContactName            string `json:"contact_name,omitempty"`              // 联系人姓名
	ContactTel             string `json:"contact_tel,omitempty"`               // 联系电话
	Description            string `json:"description,omitempty"`               // 应用备注
	Developer              string `json:"developer,omitempty"`                 // 开发者名称(需与软件著作权一致)
	Location               string `json:"location,omitempty"`                  // 开发者所在地区，格式：["北京市","北京市","海淀区"]
	OfflineAppStores       string `json:"offline_app_stores,omitempty"`        // 下线应用商店(可选值：huawei/oppo/xiaomi/meizu/vivo/smartisan/honor)
	OnlineEarnType         *int   `json:"online_earn_type,omitempty"`          // 是否盈利类型：1=盈利 2=非盈利
	UseSdk                 *int   `json:"use_sdk,omitempty"`                   // 是否接入快手SDK：0=未接入 1=已接入
	CreateSource           *int   `json:"create_source,omitempty"`             // 创建渠道：1=文件创编 2=链接创编
	FunctionIntroduction   string `json:"function_introduction,omitempty"`     // 应用功能介绍(100-1000字)
	RecordNumber           string `json:"record_number,omitempty"`             // ICP备案号
	DocumentNumber         string `json:"document_number,omitempty"`           // 证件号码
	ServiceCategory        string `json:"service_category,omitempty"`          // 服务类目
	NetworkType            *int   `json:"network_type,omitempty"`              // 网络类型：1=线上 2=线下
	OfflineAppLetterUrl    string `json:"offline_app_letter_url,omitempty"`    // 线下应用承诺函(network_type=2时必填)
	RecordCorpName         string `json:"record_corp_name,omitempty"`          // 备案主体名称
	AppRecordScreenshotUrl string `json:"app_record_screenshot_url,omitempty"` // APP备案截图(从图片上传接口获取)
	RecordCorpLicenseUrl   string `json:"record_corp_license_url,omitempty"`   // 备案主体营业执照图片(从图片上传接口获取)
}

// AppAndroidPrivacyInfo 隐私声明数据
type AppAndroidPrivacyInfo struct {
	PrivacyId int64  `json:"privacy_id"`    // 隐私协议ID，必填
	Url       string `json:"url,omitempty"` // 隐私协议URL
}

// AppAndroidPackageInfo 应用包数据
type AppAndroidPackageInfo struct {
	AppName                 string `json:"app_name,omitempty"`                  // 应用名称(不填则取APK解析名称)
	BlobStoreKey            string `json:"blob_store_key,omitempty"`            // 应用存储key(从APK上传接口获取)
	SensitivePermissionDesc string `json:"sensitive_permission_desc,omitempty"` // 敏感权限使用说明(APK涉及敏感权限时必填)
	Url                     string `json:"url,omitempty"`                       // URL
	TaskId                  int64  `json:"task_id,omitempty"`                   // 任务ID
}

var validAndroidReleaseTypes = map[int]bool{1: true, 2: true}

func (receiver *AppAndroidCreateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppAndroidCreateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.AppInfo.AppId <= 0 {
		err = errors.New("app_info.app_id is empty")
		return
	}
	if !validAndroidReleaseTypes[receiver.AppInfo.ReleaseType] {
		err = errors.New("app_info.release_type must be 1(手动发版) or 2(自动发版)")
		return
	}
	if receiver.AppPrivacyInfo.PrivacyId <= 0 {
		err = errors.New("app_privacy_info.privacy_id is empty")
		return
	}
	return
}

// AppAndroidCreateResp 创建Android应用响应数据（仅data部分）
type AppAndroidCreateResp struct {
	AppId       int64 `json:"app_id"`        // 应用ID
	GlobalAppId int64 `json:"global_app_id"` // 全局应用ID
	PackageId   int64 `json:"package_id"`    // 应用包ID(母包ID)
	PrivacyId   int64 `json:"privacy_id"`    // 隐私协议ID
	TaskId      int64 `json:"task_id"`       // 任务ID
}
