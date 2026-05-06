package model

import "errors"

// LiveComponentReportReq 直播间组件报表请求
type LiveComponentReportReq struct {
	accessTokenReq
	AdvertiserId        int64  `json:"advertiser_id"`                  // 广告主ID，必填
	StartDate           string `json:"start_date,omitempty"`           // 过滤筛选条件，格式 yyyy-MM-dd
	EndDate             string `json:"end_date,omitempty"`             // 过滤筛选条件，格式 yyyy-MM-dd
	Page                int    `json:"page,omitempty"`                 // 请求的页码，默认为1
	PageSize            int    `json:"page_size,omitempty"`            // 每页行数，默认为20，最大支持2000
	ViewType            int    `json:"view_type,omitempty"`            // 固定值1
	StartDateMin        string `json:"start_date_min,omitempty"`       // 增量拉取开始时间，格式 yyyy-MM-dd HH:mm
	EndDateMin          string `json:"end_date_min,omitempty"`         // 增量拉取结束时间，格式 yyyy-MM-dd HH:mm
	JingleBellId        int64  `json:"jingle_bell_id,omitempty"`       // 铃铛组件id
	TemporalGranularity string `json:"temporal_granularity,omitempty"` // 时间粒度：DAILY=天粒度 HOURLY=小时粒度，默认天粒度
}

func (receiver *LiveComponentReportReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *LiveComponentReportReq) Validate() (err error) {
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

// LiveComponentReportDetail 直播间组件报表数据明细
type LiveComponentReportDetail struct {
	// 主播/组件标识
	UserId         int64  `json:"user_id"`          // 主播id
	AuthorId       int64  `json:"author_id"`        // 主播id
	JingleBellId   int64  `json:"jingle_bell_id"`   // 小铃铛id
	JingleBellName string `json:"jingle_bell_name"` // 小铃铛名称

	// 直播间观看
	LiveRoomAvgPlayedSeconds float64 `json:"live_room_avg_played_seconds"` // 直播间平均观看时长
	LivePlayedStarted        float64 `json:"live_played_started"`          // 直播间进人数
	LivePlayedStartedCost    float64 `json:"live_played_started_cost"`     // 直播间进入成本

	// 直播间互动
	AdLiveShare   int64 `json:"ad_live_share"`   // 直播间分享数
	AdLiveComment int64 `json:"ad_live_comment"` // 直播间评论数

	// 粉丝关注
	AdLiveFollow     int64   `json:"ad_live_follow"`      // 粉丝关注数
	AdLiveFollowCost float64 `json:"ad_live_follow_cost"` // 粉丝关注成本

	// 直播间播放
	SimpleLivePlayedStarted   int64 `json:"simple_live_played_started"`   // 简易直播间开始播放数
	StandardLivePlayedStarted int64 `json:"standard_live_played_started"` // 标准直播间开始播放数

	// 组件数据
	ConversionComponentImpression int64   `json:"conversion_component_impression"` // 组件展示量
	ConversionComponentClick      int64   `json:"conversion_component_click"`      // 组件点击量
	ConversionComponentRate       float64 `json:"conversion_component_rate"`       // 组件点击率

	// 落地页曝光
	AdLandingPageImpression     int64 `json:"ad_landing_page_impression"`      // 广告主第三方半屏落地页曝光
	AdAppDownloadHalfImpression int64 `json:"ad_app_download_half_impression"` // 安卓APP下载类半屏落地页曝光
}

// LiveComponentReportResp 直播间组件报表响应数据（仅data部分）
type LiveComponentReportResp struct {
	TotalCount int64                       `json:"total_count"` // 数据的总行数
	Details    []LiveComponentReportDetail `json:"details"`     // 数据明细信息
}
