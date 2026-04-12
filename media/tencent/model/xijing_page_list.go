package model

import "errors"

// ========== 蹊径获取落地页列表 ==========
// https://developers.e.qq.com/v3.0/docs/api/xijing_page_list/get

// 蹊径落地页类型枚举
const (
	XijingPageTypeDefaultH5        = "XJ_DEFAULT_H5"         // 默认 H5
	XijingPageTypeAndroidAppH5     = "XJ_ANDROID_APP_H5"     // Android App H5
	XijingPageTypeIosAppH5         = "XJ_IOS_APP_H5"         // iOS App H5
	XijingPageTypeWebsiteH5        = "XJ_WEBSITE_H5"         // 网站 H5
	XijingPageTypeAndroidAppNative = "XJ_ANDROID_APP_NATIVE" // Android App 原生
	XijingPageTypeIosAppNative     = "XJ_IOS_APP_NATIVE"     // iOS App 原生
	XijingPageTypeWebsiteNative    = "XJ_WEBSITE_NATIVE"     // 网站原生
	XijingPageTypeFenglingLbs      = "XJ_FENGLING_LBS"       // 蜂灵 LBS
)

// 落地页发布状态枚举
const (
	XijingPagePublishStatusUnpublish = "LANDING_PAGE_STATUS_UNPUBLISH" // 未发布
	XijingPagePublishStatusPublished = "LANDING_PAGE_STATUS_PUBLISHED" // 已发布
	XijingPagePublishStatusOffline   = "LANDING_PAGE_STATUS_OFFLINE"   // 已下线
	XijingPagePublishStatusDeleting  = "LANDING_PAGE_STATUS_DELETING"  // 删除中
	XijingPagePublishStatusDeleted   = "LANDING_PAGE_STATUS_DELETED"   // 已删除
)

// 落地页审核状态枚举
const (
	XijingPageStatusEditing  = "LANDING_PAGE_STATUS_EDITING"  // 编辑中
	XijingPageStatusPending  = "LANDING_PAGE_STATUS_PENDING"  // 审核中
	XijingPageStatusApproved = "LANDING_PAGE_STATUS_APPROVED" // 审核通过
	XijingPageStatusRejected = "LANDING_PAGE_STATUS_REJECTED" // 审核拒绝
	XijingPageStatusDeleted  = "LANDING_PAGE_STATUS_DELETED"  // 已删除
)

// 落地页查询来源枚举
const (
	XijingPageSourceGrant = "GRANT" // 授权落地页
	XijingPageSourceOwner = "OWNER" // 自有落地页（默认）
)

// App 类型枚举
const (
	XijingPageAppTypeAndroid = "ANDROID" // Android
	XijingPageAppTypeIos     = "IOS"     // iOS
)

// 落地页状态查询方式枚举
const (
	XijingPageQueryTypeDefault = "DEFAULT" // 默认（不含已删除）
	XijingPageQueryTypeDeleted = "DELETED" // 仅已删除
	XijingPageQueryTypeAll     = "ALL"     // 全部
)

// 分页常量
const (
	MinXijingPageListPageIndex     = 1     // page_index 最小值
	MaxXijingPageListPageIndex     = 99999 // page_index 最大值
	MinXijingPageListPageSize      = 1     // page_size 最小值
	MaxXijingPageListPageSize      = 100   // page_size 最大值
	DefaultXijingPageListPageIndex = 1     // page_index 默认值
	DefaultXijingPageListPageSize  = 10    // page_size 默认值
)

// 字段长度常量
const (
	MaxXijingPageListServiceIDBytes     = 256 // page_service_id 最大字节数
	MinXijingPageListPageNameBytes      = 1   // page_name 最小字节数
	MaxXijingPageListPageNameBytes      = 20  // page_name 最大字节数
	MaxXijingPageListModifyTimeBytes    = 30  // page_last_modify_start/end_time 最大字节数
	MaxXijingPageListPageTypeCount      = 8   // page_type 最大数量
	MaxXijingPageListPublishStatusCount = 5   // page_publish_status 最大数量
	MaxXijingPageListPageStatusCount    = 5   // page_status 最大数量
)

// XijingPageListGetReq 蹊径获取落地页列表请求（GET）
// https://developers.e.qq.com/v3.0/docs/api/xijing_page_list/get
type XijingPageListGetReq struct {
	GlobalReq
	AccountID               int64    `json:"account_id"`                            // 广告主帐号 id (必填)
	PageID                  int64    `json:"page_id,omitempty"`                     // 蹊径落地页 id
	PageServiceID           string   `json:"page_service_id,omitempty"`             // 落地页服务 id，0-256字节
	PageName                string   `json:"page_name,omitempty"`                   // 落地页名称，1-20字节（page_source=GRANT时不支持）
	PageType                []string `json:"page_type,omitempty"`                   // 落地页类型，最多8个
	PageLastModifyStartTime string   `json:"page_last_modify_start_time,omitempty"` // 最后更新时间起始，0-30字节
	PageLastModifyEndTime   string   `json:"page_last_modify_end_time,omitempty"`   // 最后更新时间结束，0-30字节
	PageSize                int      `json:"page_size,omitempty"`                   // 每页条数，1-100，默认10
	PageIndex               int      `json:"page_index,omitempty"`                  // 搜索页码，1-99999，默认1
	PagePublishStatus       []string `json:"page_publish_status,omitempty"`         // 落地页发布状态，最多5个
	PageStatus              []string `json:"page_status,omitempty"`                 // 落地页审核状态，最多5个
	PageSource              string   `json:"page_source,omitempty"`                 // 查询类型，GRANT/OWNER，默认OWNER
	PageOwnerID             int64    `json:"page_owner_id,omitempty"`               // 授权落地页所属账户 id
	AppID                   int64    `json:"app_id,omitempty"`                      // 绑定的 App id
	AppType                 string   `json:"app_type,omitempty"`                    // 绑定的 App 类型，ANDROID/IOS
	QueryType               string   `json:"query_type,omitempty"`                  // 落地页状态查询方式，DEFAULT/DELETED/ALL
}

func (r *XijingPageListGetReq) Format() {
	r.GlobalReq.Format()
	if r.PageIndex == 0 {
		r.PageIndex = DefaultXijingPageListPageIndex
	}
	if r.PageSize == 0 {
		r.PageSize = DefaultXijingPageListPageSize
	}
}

// Validate 验证蹊径获取落地页列表请求参数
func (r *XijingPageListGetReq) Validate() error {
	if r.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if r.PageServiceID != "" && len(r.PageServiceID) > MaxXijingPageListServiceIDBytes {
		return errors.New("page_service_id长度不能超过256字节")
	}
	if r.PageName != "" && (len(r.PageName) < MinXijingPageListPageNameBytes || len(r.PageName) > MaxXijingPageListPageNameBytes) {
		return errors.New("page_name长度须在1-20字节之间")
	}
	if len(r.PageType) > MaxXijingPageListPageTypeCount {
		return errors.New("page_type数组长度不能超过8")
	}
	if r.PageLastModifyStartTime != "" && len(r.PageLastModifyStartTime) > MaxXijingPageListModifyTimeBytes {
		return errors.New("page_last_modify_start_time长度不能超过30字节")
	}
	if r.PageLastModifyEndTime != "" && len(r.PageLastModifyEndTime) > MaxXijingPageListModifyTimeBytes {
		return errors.New("page_last_modify_end_time长度不能超过30字节")
	}
	if r.PageIndex < MinXijingPageListPageIndex || r.PageIndex > MaxXijingPageListPageIndex {
		return errors.New("page_index须在1-99999之间")
	}
	if r.PageSize < MinXijingPageListPageSize || r.PageSize > MaxXijingPageListPageSize {
		return errors.New("page_size须在1-100之间")
	}
	if len(r.PagePublishStatus) > MaxXijingPageListPublishStatusCount {
		return errors.New("page_publish_status数组长度不能超过5")
	}
	if len(r.PageStatus) > MaxXijingPageListPageStatusCount {
		return errors.New("page_status数组长度不能超过5")
	}
	return r.GlobalReq.Validate()
}

// XijingPageAppInfo 落地页绑定的 App 信息
type XijingPageAppInfo struct {
	AndroidAppID int64 `json:"android_app_id,omitempty"` // Android App id
	IosAppID     int64 `json:"ios_app_id,omitempty"`     // iOS App id
}

// XijingPageListItem 蹊径落地页列表项
type XijingPageListItem struct {
	PageID             int64              `json:"page_id"`               // 蹊径落地页 id
	PageServiceID      string             `json:"page_service_id"`       // 落地页服务 id
	PageName           string             `json:"page_name"`             // 落地页名称（管理用）
	PageType           string             `json:"page_type"`             // 落地页类型
	PagePublishStatus  string             `json:"page_publish_status"`   // 落地页发布状态
	PageStatus         string             `json:"page_status"`           // 落地页审核状态
	PageLastModifyTime string             `json:"page_last_modify_time"` // 落地页最后更新时间
	PageOwnerID        int64              `json:"page_owner_id"`         // 授权落地页所属账户 id
	PublishURL         string             `json:"publish_url"`           // 落地页 URL（未通过审核则不返回）
	RejectReason       string             `json:"reject_reason"`         // 落地页审核拒绝理由
	PlayableType       string             `json:"playable_type"`         // 互动落地页类型
	PublishAppID       *XijingPageAppInfo `json:"publish_app_id"`        // 已发布版本落地页 App 信息
	UnpublishAppID     *XijingPageAppInfo `json:"unpublish_app_id"`      // 未发布版本落地页 App 信息
}

// XijingPageListGetResp 蹊径获取落地页列表响应
// https://developers.e.qq.com/v3.0/docs/api/xijing_page_list/get
type XijingPageListGetResp struct {
	List       []*XijingPageListItem `json:"list"`        // 落地页列表
	PageInfo   *PageInfo             `json:"page_info"`   // 分页配置信息
	TotalPages int                   `json:"total_pages"` // 总页数
	PageSize   int                   `json:"page_size"`   // 每页条数
	PageIndex  int                   `json:"page_index"`  // 当前页码
}
