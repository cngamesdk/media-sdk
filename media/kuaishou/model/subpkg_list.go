package model

import "errors"

// SubpkgListReq 获取分包管理/回收站列表请求
// https://ad.e.kuaishou.com/rest/openapi/gw/dsp/appcenter/subPackage/list
type SubpkgListReq struct {
	accessTokenReq
	AdvertiserId int64    `json:"advertiser_id"`       // 广告主ID，必填
	AppId        int64    `json:"app_id"`              // 应用ID，必填
	KeyWord      string   `json:"key_word,omitempty"`  // 搜索关键词（渠道号或分包备注关键词）
	Page         int      `json:"page,omitempty"`      // 当前页，默认1
	ListType     int      `json:"list_type,omitempty"` // 列表类型：不填为分包管理列表，2-分包回收列表
	PageSize     int      `json:"page_size,omitempty"` // 分页大小，默认10
	Version      []string `json:"version,omitempty"`   // 版本信息，仅分包管理列表生效，多选
	Status       int      `json:"status,omitempty"`    // 分包状态：0-全部，1-审核中，2-审核未通过，4-已发布，6-构建中，7-更新中，8-构建失败
}

func (receiver *SubpkgListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *SubpkgListReq) Validate() (err error) {
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

// SubpkgListItem 应用分包条目
type SubpkgListItem struct {
	CanRecycle       *bool  `json:"can_recycle"`        // 是否可恢复（仅回收站列表时有效）
	CanUpdate        *bool  `json:"can_update"`         // 是否可更新（仅管理列表时有效）
	ChannelId        string `json:"channel_id"`         // 渠道号(分包号)
	DeleteTime       int64  `json:"delete_time"`        // 删除时间（仅回收站列表时有效）
	Description      string `json:"description"`        // 分包描述
	PackageId        int64  `json:"package_id"`         // 应用包ID
	ParentPackageId  int64  `json:"parent_package_id"`  // 母包ID
	RealAppVersion   string `json:"real_app_version"`   // 应用版本信息
	SubPackageStatus int    `json:"sub_package_status"` // 应用分包状态：1-审核中，2-审核失败，3-待发布，4-已发布，5-已下架，6-创建中，7-更新中，8-构建失败
	UpdateTime       int64  `json:"update_time"`        // 更新时间（仅管理列表时有效）
	Url              string `json:"url"`                // 应用下载地址
}

// SubpkgListResp 获取分包管理/回收站列表响应数据（仅data部分）
type SubpkgListResp struct {
	CurrentPage int              `json:"current_page"` // 当前页
	PageSize    int              `json:"page_size"`    // 分页大小
	TotalCount  int64            `json:"total_count"`  // 总数
	List        []SubpkgListItem `json:"list"`         // 应用分包列表
}
