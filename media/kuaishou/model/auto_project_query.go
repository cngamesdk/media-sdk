package model

import "errors"

// AutoProjectQueryReq 智投项目查询请求
type AutoProjectQueryReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主账号ID，必填
	Page         int   `json:"page"`          // 当前页码
	PageSize     int   `json:"page_size"`     // 每页大小
}

func (receiver *AutoProjectQueryReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AutoProjectQueryReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	return
}

// BookInfo 小说信息
type BookInfo struct {
	BookName string `json:"book_name"` // 小说标题
	Id       int64  `json:"id"`        // 小说id
}

// PhotoPackageDetail 素材包信息
type PhotoPackageDetail struct {
	Name           string `json:"name"`             // 素材包名称
	PhotoPackageId int64  `json:"photo_package_id"` // 素材包id
}

// SeriesInfo 短剧信息
type SeriesInfo struct {
	CoverImg             string `json:"cover_img"`              // 短剧封面
	Description          string `json:"description"`            // 短剧描述
	EpisodeAmount        int    `json:"episode_amount"`         // 剧集数量
	Id                   string `json:"id"`                     // 短剧id
	SupportAdUnlock      bool   `json:"support_ad_unlock"`      // 是否支持观看剧集解锁
	SupportMultiTemplate bool   `json:"support_multi_template"` // 是否支持短剧付费面板优选
	Title                string `json:"title"`                  // 短剧标题
}

// AutoProjectItem 智投项目信息
type AutoProjectItem struct {
	AccountAutoManage   int                  `json:"account_auto_manage"`   // 智投开关
	AuaxProjectId       int64                `json:"auax_project_id"`       // 项目id
	AutoManageType      int                  `json:"auto_manage_type"`      // 账号智投模式类型
	AutoPhotoScope      int                  `json:"auto_photo_scope"`      // 基建素材选取范围
	BookInfo            BookInfo             `json:"book_info"`             // 小说信息
	CreateType          int                  `json:"create_type"`           // 创建类型
	Description         string               `json:"description"`           // 作品广告语
	KolUserId           int64                `json:"kol_user_id"`           // 广告主快手ID
	KolUserInfo         KolUserInfoResp      `json:"kol_user_info"`         // 快手号信息
	KolUserType         int                  `json:"kol_user_type"`         // 快手号类型
	Name                string               `json:"name"`                  // 智投名称
	OcpxActionType      int                  `json:"ocpx_action_type"`      // 转化目标
	PhotoPackageDetails []PhotoPackageDetail `json:"photo_package_details"` // 素材包信息
	PutStatus           int                  `json:"put_status"`            // 投放状态
	RoiRatio            float64              `json:"roi_ratio"`             // roi系数
	SeriesInfo          SeriesInfo           `json:"series_info"`           // 短剧信息
}

// AutoProjectQueryResp 智投项目查询响应数据（仅data部分）
type AutoProjectQueryResp struct {
	CurrentPage int               `json:"current_page"` // 当前页码
	List        []AutoProjectItem `json:"list"`         // 智投项目列表
	PageSize    int               `json:"page_size"`    // 每页大小
	TotalCount  int64             `json:"total_count"`  // 数据总条数
}
