package model

import "errors"

// LiveUserReportReq 直播间报表请求
type LiveUserReportReq struct {
	accessTokenReq
	AdvertiserId        int64  `json:"advertiser_id"`                  // 广告主ID，必填
	StartDate           string `json:"start_date,omitempty"`           // 开始时间，格式 yyyy-MM-dd
	EndDate             string `json:"end_date,omitempty"`             // 结束时间，格式 yyyy-MM-dd
	Page                int    `json:"page,omitempty"`                 // 页码，默认为1
	PageSize            int    `json:"page_size,omitempty"`            // 页面大小，默认为20，最大支持2000
	ViewType            int    `json:"view_type,omitempty"`            // 固定值1
	StartDateMin        string `json:"start_date_min,omitempty"`       // 增量拉取开始时间
	EndDateMin          string `json:"end_date_min,omitempty"`         // 增量拉取结束时间
	UserId              int64  `json:"user_id,omitempty"`              // 主播id
	LiveStreamId        int64  `json:"live_stream_id,omitempty"`       // 直播间id
	TemporalGranularity string `json:"temporal_granularity,omitempty"` // 时间粒度：DAILY=天粒度 HOURLY=小时粒度
}

func (receiver *LiveUserReportReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *LiveUserReportReq) Validate() (err error) {
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

// LiveUserReportDetail 直播间报表数据明细
type LiveUserReportDetail struct {
	// 主播/直播间标识
	UserId        int64 `json:"user_id"`         // 主播id
	LiveStreamId  int64 `json:"live_stream_id"`  // 直播间id
	LiveStartTime int64 `json:"live_start_time"` // 直播开始时间
	LiveEndTime   int64 `json:"live_end_time"`   // 直播结束时间

	// 计划/组/创意标识
	CampaignId int64 `json:"campaign_id"` // 计划ID
	UnitId     int64 `json:"unit_id"`     // 广告组ID
	CreativeId int64 `json:"creative_id"` // 创意ID

	// 核心消耗与曝光
	Charge          float64 `json:"charge"`            // 花费（元）
	Impression      int64   `json:"impression"`        // 封面曝光数
	PhotoClick      int64   `json:"photo_click"`       // 封面点击数
	Aclick          int64   `json:"aclick"`            // 素材曝光数
	Bclick          int64   `json:"bclick"`            // 行为数
	PhotoClickRatio float64 `json:"photo_click_ratio"` // 封面点击率
	ActionRatio     float64 `json:"action_ratio"`      // 素材点击率

	// 成本指标
	Impression1kCost float64 `json:"impression_1k_cost"` // 平均千次封面曝光花费(元)
	Click1kCost      float64 `json:"click_1k_cost"`      // 平均千次素材曝光花费(元)
	PhotoClickCost   float64 `json:"photo_click_cost"`   // 平均封面点击单价(元)
	ActionCost       float64 `json:"action_cost"`        // 平均行为单价(元)

	// 播放相关
	Play3sRatio  float64 `json:"play_3s_ratio"`  // 3秒播放率
	Play5sRatio  float64 `json:"play_5s_ratio"`  // 5秒播放率
	PlayEndRatio float64 `json:"play_end_ratio"` // 完播率

	// 社交互动
	Share    int64 `json:"share"`    // 分享数
	Comment  int64 `json:"comment"`  // 评论数
	Likes    int64 `json:"likes"`    // 点赞数
	Report   int64 `json:"report"`   // 举报数
	Block    int64 `json:"block"`    // 拉黑数
	Negative int64 `json:"negative"` // 不感兴趣数

	// 转化
	Conversion   int64 `json:"conversion"`     // 激活数
	LivePlayed3s int64 `json:"live_played_3s"` // 直播观看3秒数

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
	SimpleLivePlayedStarted       int64 `json:"simple_live_played_started"`        // 简易直播间开始播放数
	StandardLivePlayedStarted     int64 `json:"standard_live_played_started"`      // 标准直播间开始播放数
	SimpleLiveRoomPlayedSeconds   int64 `json:"simple_live_room_played_seconds"`   // 简易直播间观看时长
	StandardLiveRoomPlayedSeconds int64 `json:"standard_live_room_played_seconds"` // 标准直播间观看时长

	// 组件数据
	ConversionComponentImpression int64   `json:"conversion_component_impression"` // 组件展示量
	ConversionComponentClick      int64   `json:"conversion_component_click"`      // 组件点击量
	ConversionComponentRate       float64 `json:"conversion_component_rate"`       // 组件点击率

	// 落地页曝光
	AdLandingPageImpression     int64 `json:"ad_landing_page_impression"`      // 广告主第三方半屏落地页曝光
	AdAppDownloadHalfImpression int64 `json:"ad_app_download_half_impression"` // 安卓APP下载类半屏落地页曝光

	// 铃铛组件
	JingleBellId   int64  `json:"jingle_bell_id"`   // 小铃铛id
	JingleBellName string `json:"jingle_bell_name"` // 小铃铛名称
}

// LiveUserReportResp 直播间报表响应数据（仅data部分）
type LiveUserReportResp struct {
	TotalCount int64                  `json:"total_count"` // 数据的总行数
	Details    []LiveUserReportDetail `json:"details"`     // 数据明细信息
}
