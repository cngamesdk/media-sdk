package model

import "errors"

// AppReleaseListReq 获取新版应用发布列表请求【单元创编】
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/release/list
type AppReleaseListReq struct {
	accessTokenReq
	AdvertiserId int64   `json:"advertiser_id"`       // 广告主ID，必填
	KeyWord      string  `json:"key_word,omitempty"`  // 搜索关键词，支持应用ID或应用名称搜索
	Page         int     `json:"page,omitempty"`      // 当前页码，默认1
	PageSize     int     `json:"page_size,omitempty"` // 分页大小，默认10
	ListType     *int    `json:"list_type,omitempty"` // 列表类型：不传=全部 1=我创建的 2=共享给我的
	Platform     string  `json:"platform,omitempty"`  // 平台：android 或 ios
	AppIds       []int64 `json:"app_ids,omitempty"`   // 应用ID列表，最多支持100个
}

func (receiver *AppReleaseListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppReleaseListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.AppIds) > 100 {
		err = errors.New("app_ids must not exceed 100 items")
		return
	}
	return
}

// AppReleaseListItem 应用发布列表条目
type AppReleaseListItem struct {
	AccountId             int64    `json:"account_id"`             // 应用创建者账号ID
	AppDetailImg          string   `json:"app_detail_img"`         // 应用详情图
	AppIconUrl            string   `json:"app_icon_url"`           // 应用图标URL
	AppId                 int64    `json:"app_id"`                 // 应用ID
	AppPrivacyUrl         string   `json:"app_privacy_url"`        // 应用隐私政策URL
	IosAppId              string   `json:"ios_app_id"`             // 解析出的iOS App ID
	OfflineAppStores      string   `json:"offline_app_stores"`     // 下线应用商店(huawei/oppo/vivo/xiaomi/meizu/smartisan/honor)
	PackageId             int64    `json:"package_id"`             // 应用包ID
	PackageName           string   `json:"package_name"`           // 应用包名
	PackageSize           int64    `json:"package_size"`           // 应用包大小
	PermissionInformation []string `json:"permission_information"` // 权限信息ID列表
	Platform              string   `json:"platform"`               // 平台：android 或 ios
	RealAppName           string   `json:"real_app_name"`          // 应用名称
	RealAppVersion        string   `json:"real_app_version"`       // 应用版本信息
	SourceType            int      `json:"source_type"`            // 应用来源：1=我创建的 2=共享给我的
	UpdateTime            int64    `json:"update_time"`            // 更新时间(毫秒)
	Url                   string   `json:"url"`                    // 应用下载地址
	UseSdk                int      `json:"use_sdk"`                // 是否接入快手广告监测SDK：0=未接入 1=已接入
	VersionCode           int64    `json:"version_code"`           // 应用版本code
}

// AppReleaseListResp 获取新版应用发布列表响应数据（仅data部分）
type AppReleaseListResp struct {
	CurrentPage int                  `json:"current_page"` // 当前页码
	PageSize    int                  `json:"page_size"`    // 分页大小
	TotalCount  int64                `json:"total_count"`  // 总数量
	List        []AppReleaseListItem `json:"list"`         // 应用列表
}
