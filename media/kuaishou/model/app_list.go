package model

import "errors"

// AppListReq 获取应用列表请求
// https://developers.e.kuaishou.com/docs?docType=DSP&documentId=2774&menuId=3510
type AppListReq struct {
	accessTokenReq
	AdvertiserId int64    `json:"advertiser_id"`        // 广告主ID，必填
	ListType     int      `json:"list_type"`            // 列表类型，必填：1=我创建的 2=共享给我的，默认1
	AppIds       []string `json:"app_ids,omitempty"`    // 批量应用ID查询
	KeyWord      string   `json:"key_word,omitempty"`   // 搜索关键词，支持应用ID或应用名称
	AppStatus    *int     `json:"app_status,omitempty"` // 应用状态：不传=全部 1=审核中 2=审核失败 3=待发布 4=已发布 5=已下架
	Page         int      `json:"page,omitempty"`       // 当前页码，默认1
	PageSize     int      `json:"page_size,omitempty"`  // 分页大小，默认20
	Platform     string   `json:"platform,omitempty"`   // 平台：android 或 ios
	StartDate    string   `json:"start_date,omitempty"` // 发布时间范围-起始(需同时填写end_date)
	EndDate      string   `json:"end_date,omitempty"`   // 发布时间范围-截止
}

func (receiver *AppListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.ListType != 1 && receiver.ListType != 2 {
		err = errors.New("list_type must be 1(我创建的) or 2(共享给我的)")
		return
	}
	return
}

// AppListAppSource 应用创建者信息（source_type=2时返回）
type AppListAppSource struct {
	AccountId   int64  `json:"account_id"`   // 创建者账号ID
	AccountName string `json:"account_name"` // 创建者账号名称
}

// AppListItem 应用列表条目
type AppListItem struct {
	AccountId         int64            `json:"account_id"`          // 应用创建者账号ID
	AppDetailImg      string           `json:"app_detail_img"`      // 应用详情图
	AppIconUrl        string           `json:"app_icon_url"`        // 应用图标URL
	AppId             int64            `json:"app_id"`              // 应用ID
	AppPrivacyUrl     string           `json:"app_privacy_url"`     // 应用隐私政策链接
	AppSource         AppListAppSource `json:"app_source"`          // 应用创建者信息(source_type=2时返回)
	AppStatus         int              `json:"app_status"`          // 应用状态：1=审核中 2=审核失败 3=待发布 4=已发布 5=已下架
	IosAppId          string           `json:"ios_app_id"`          // 解析出的iOS App ID
	OfflineAppStores  string           `json:"offline_app_stores"`  // 下线应用商店(huawei/oppo/vivo/xiaomi/meizu/smartisan/honor)
	PackageId         int64            `json:"package_id"`          // 应用包ID
	PackageName       string           `json:"package_name"`        // 应用包名
	PackageSize       int64            `json:"package_size"`        // 应用包大小
	Platform          string           `json:"platform"`            // 平台：android 或 ios
	RealAppName       string           `json:"real_app_name"`       // 应用名称
	RealAppVersion    string           `json:"real_app_version"`    // 应用版本信息
	ReviewDetail      string           `json:"review_detail"`       // 审核详情
	SourceType        int              `json:"source_type"`         // 应用来源：1=我创建的 2=共享给我的
	TraceActivation   int              `json:"trace_activation"`    // 转化追踪状态
	UpdateTime        int64            `json:"update_time"`         // 更新时间(毫秒)
	Url               string           `json:"url"`                 // 应用下载地址
	UseSdk            int              `json:"use_sdk"`             // 是否接入快手广告监测SDK
	VersionCode       int64            `json:"version_code"`        // 应用版本code
	ShareType         int              `json:"share_type"`          // 共享类型
	CreateSource      int              `json:"create_source"`       // 应用创建来源：1=文件上传 2=链接下载
	ApkDownloadUrl    string           `json:"apk_download_url"`    // 应用链接下载地址
	TaskId            int64            `json:"task_id"`             // 应用链接下载任务ID
	DownloadStatus    int              `json:"download_status"`     // 应用链接下载状态
	CreateFailedTypes []int            `json:"create_failed_types"` // 链接应用创建失败原因
}

// AppListResp 获取应用列表响应数据（仅data部分）
type AppListResp struct {
	CurrentPage int           `json:"current_page"` // 当前页码
	PageSize    int           `json:"page_size"`    // 分页大小
	TotalCount  int64         `json:"total_count"`  // 总数量
	List        []AppListItem `json:"list"`         // 应用列表
}
