package model

import "errors"

// AuaxSeriesListReq 智投短剧查询请求
type AuaxSeriesListReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"` // 广告主账号ID，必填
	Cursor       int64 `json:"cursor"`        // 查询起始游标（类似页码，首次查询不传，之后传上次返回的游标值）
	Limit        int   `json:"limit"`         // 查询数据量，必填
}

func (receiver *AuaxSeriesListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AuaxSeriesListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.Limit <= 0 {
		err = errors.New("limit is empty")
		return
	}
	return
}

// AuaxSeriesInfo 智投短剧信息
type AuaxSeriesInfo struct {
	Id            int64  `json:"id"`             // 短剧ID
	KolUserId     int64  `json:"kol_user_id"`    // 快手号ID
	KolUserType   int    `json:"kol_user_type"`  // 快手号类型：1-普通快手号，2-蓝V服务号，3-聚星达人
	Title         string `json:"title"`          // 短剧标题
	Description   string `json:"description"`    // 短剧简介
	CoverImg      string `json:"cover_img"`      // 短剧封面URL
	EpisodeAmount int64  `json:"episode_amount"` // 剧集数量
	PayMode       []int  `json:"pay_mode"`       // 付费模式：1-打包，2-虚拟币，3-观看广告解锁
	ValidStatus   bool   `json:"valid_status"`   // 短剧状态：true-可用，false-不可用
	InvalidReason string `json:"invalid_reason"` // 不可用原因
}

// AuaxSeriesListResp 智投短剧查询响应数据（仅data部分）
type AuaxSeriesListResp struct {
	AuaxSeriesInfos []AuaxSeriesInfo `json:"auax_series_infos"` // 智投短剧信息
	Cursor          int64            `json:"cursor"`            // 下次查询起始游标，null表示全部数据已查询完
}
