package model

import "errors"

// AppVersionListReq 获取应用版本记录请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/app/version/list
type AppVersionListReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"`       // 广告主ID，必填
	AppId        int64 `json:"app_id"`              // 应用ID，必填
	Page         int   `json:"page,omitempty"`      // 当前页，默认1
	PageSize     int   `json:"page_size,omitempty"` // 分页大小，默认20
}

func (receiver *AppVersionListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AppVersionListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.AppId <= 0 {
		err = errors.New("app_id is empty")
		return
	}
	return
}

// AppVersionItem 应用版本条目
type AppVersionItem struct {
	RealAppVersion string `json:"real_app_version"` // 应用版本信息
	UpdateTime     int64  `json:"update_time"`      // 更新时间
	VersionCode    int64  `json:"version_code"`     // 应用版本号
	VersionStatus  int    `json:"version_status"`   // 版本状态：1-审核中，2-审核失败，3-待发布，4-已发布，5-已下架
}

// AppVersionListResp 获取应用版本记录响应数据（仅data部分）
type AppVersionListResp struct {
	CurrentPage int              `json:"current_page"` // 当前页
	PageSize    int              `json:"page_size"`    // 分页大小
	TotalCount  int64            `json:"total_count"`  // 总数
	List        []AppVersionItem `json:"list"`         // 应用版本列表
}
