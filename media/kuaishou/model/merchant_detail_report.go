package model

import "errors"

// MerchantDetailReportReq 小店通转化数据报表请求
type MerchantDetailReportReq struct {
	accessTokenReq
	AdvertiserId          int64  `json:"advertiser_id"`                     // 广告主ID，必填
	ViewType              int    `json:"view_type"`                         // 数据类型（必填）：1=账户 2=广告计划 3=广告组 4=广告创意
	GroupType             int    `json:"group_type,omitempty"`              // 汇总方式：1=天（默认） 2=小时（此时start_date只能是当天，end_date只能是下一天）
	StartDate             string `json:"start_date"`                        // 过滤筛选条件，格式 yyyy-MM-dd，必填
	EndDate               string `json:"end_date"`                          // 过滤筛选条件，格式 yyyy-MM-dd，必填
	CampaignIds           string `json:"campaign_ids,omitempty"`            // 计划id列表，逗号分割，viewType=2时必填
	UnitIds               string `json:"unit_ids,omitempty"`                // 单元id列表，逗号分割，viewType=3时必填
	CreativeIds           string `json:"creative_ids,omitempty"`            // 创意id列表，逗号分割，viewType=4时必填
	ProgrammedCreativeIds string `json:"programmed_creative_ids,omitempty"` // 程序化创意id列表，逗号分割，viewType=4时可填
	Page                  int    `json:"page,omitempty"`                    // 请求的页码，默认为1
	PageSize              int    `json:"page_size,omitempty"`               // 每页行数，默认为20，最大支持2000
}

func (receiver *MerchantDetailReportReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *MerchantDetailReportReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.ViewType <= 0 {
		err = errors.New("view_type is empty")
		return
	}
	return
}

// MerchantDetailReportDetail 小店通转化数据报表明细
type MerchantDetailReportDetail struct {
	// 日期/时间
	ReportDate     string `json:"report_date"`      // 日期
	ReportHour     int64  `json:"report_hour"`      // 小时
	ReportDateHour string `json:"report_date_hour"` // 日期+小时

	// 计划/组/创意标识
	CampaignName string `json:"campaign_name"` // 计划名称
	CampaignId   int64  `json:"campaign_id"`   // 计划ID
	UnitName     string `json:"unit_name"`     // 广告组名称
	UnitId       int64  `json:"unit_id"`       // 广告组ID
	CreativeName string `json:"creative_name"` // 创意名称
	CreativeId   int64  `json:"creative_id"`   // 创意ID

	// 状态/时间
	PutStatus          int64 `json:"put_status"`           // 投放状态
	CampaignCreateTime int64 `json:"campaign_create_time"` // 广告计划创建时间
	UnitCreateTime     int64 `json:"unit_create_time"`     // 广告组创建时间
	CreativeCreateTime int64 `json:"creative_create_time"` // 广告创意创建时间

	// 核心消耗与曝光
	CostTotal       int64   `json:"cost_total"`        // 花费
	Impression      int64   `json:"impression"`        // 封面曝光数
	PhotoClick      int64   `json:"photo_click"`       // 封面点击数
	PhotoClickRatio float64 `json:"photo_click_ratio"` // 封面点击率
	Click           int64   `json:"click"`             // 素材曝光数
	ActionbarClick  int64   `json:"actionbar_click"`   // 行为数
	ActionRatio     float64 `json:"action_ratio"`      // 行为率

	// 成本指标
	MerchantPhotoImpression1kCost float64 `json:"merchant_photo_impression_1k_cost"` // 平均千次封面曝光花费(元)
	MerchantPhotoClickCost        float64 `json:"merchant_photo_click_cost"`         // 平均封面点击单价(元)
	MerchantImpression1kCost      float64 `json:"merchant_impression_1k_cost"`       // 平均千次素材曝光花费(元)
	MerchantClickCost             float64 `json:"merchant_click_cost"`               // 平均行为单价(元)

	// 播放相关
	Play3sRatio        float64 `json:"play_3s_ratio"`        // 3s播放率
	Play5sRatio        float64 `json:"play_5s_ratio"`        // 5s播放率
	PlayEndRatio       float64 `json:"play_end_ratio"`       // 完播率
	PlayedThreeSeconds int64   `json:"played_three_seconds"` // 3s播放数
	PlayedFiveSeconds  int64   `json:"played_five_seconds"`  // 5s播放数
	PlayedEnd          int64   `json:"played_end"`           // 完播数

	// 社交互动
	Share    int64   `json:"share"`    // 分享数
	Comment  int64   `json:"comment"`  // 评论数
	Likes    int64   `json:"likes"`    // 点赞数
	Report   int64   `json:"report"`   // 举报数
	Block    int64   `json:"block"`    // 拉黑数
	Negative float64 `json:"negative"` // 减少此类作品数

	// 涨粉
	MerchantRecoFans int64   `json:"merchant_reco_fans"` // 涨粉量
	RecoFansCost     float64 `json:"reco_fans_cost"`     // 涨粉成本（元）

	// 订单/GMV
	PaiedOrder int64   `json:"paied_order"` // 支付订单数
	OrderCost  float64 `json:"order_cost"`  // 下单成本
	Gmv        float64 `json:"gmv"`         // GMV

	// 多日累计GMV
	T0Gmv  int64 `json:"t0_gmv"`  // 当日累计GMV
	T1Gmv  int64 `json:"t1_gmv"`  // 投后1日累计GMV
	T3Gmv  int64 `json:"t3_gmv"`  // 投后3日累计GMV
	T7Gmv  int64 `json:"t7_gmv"`  // 投后7日累计GMV
	T15Gmv int64 `json:"t15_gmv"` // 投后15日累计GMV
	T30Gmv int64 `json:"t30_gmv"` // 投后30日累计GMV

	// ROI
	Roi    float64 `json:"roi"`     // ROI
	T0Roi  float64 `json:"t0_roi"`  // 当日累计ROI
	T1Roi  float64 `json:"t1_roi"`  // 投后1日累计ROI
	T3Roi  float64 `json:"t3_roi"`  // 投后3日累计ROI
	T7Roi  float64 `json:"t7_roi"`  // 投后7日累计ROI
	T15Roi float64 `json:"t15_roi"` // 投后15日累计ROI
	T30Roi float64 `json:"t30_roi"` // 投后30日累计ROI

	// 涨粉留存
	T1Retention       int64   `json:"t1_retention"`        // 投后1日涨粉留存量
	T7Retention       int64   `json:"t7_retention"`        // 投后7日涨粉留存量
	T30Retention      int64   `json:"t30_retention"`       // 投后30日涨粉留存量
	T1RetentionRatio  float64 `json:"t1_retention_ratio"`  // 投后1日涨粉留存率
	T7RetentionRatio  float64 `json:"t7_retention_ratio"`  // 投后7日涨粉留存率
	T30RetentionRatio float64 `json:"t30_retention_ratio"` // 投后30日涨粉留存率

	// 多日累计订单成交量
	T0OrderCnt  int64 `json:"t0_order_cnt"`  // 当日累计订单成交量
	T1OrderCnt  int64 `json:"t1_order_cnt"`  // 投后1日累计订单成交量
	T3OrderCnt  int64 `json:"t3_order_cnt"`  // 投后3日累计订单成交量
	T7OrderCnt  int64 `json:"t7_order_cnt"`  // 投后7日累计订单成交量
	T15OrderCnt int64 `json:"t15_order_cnt"` // 投后15日累计订单成交量
	T30OrderCnt int64 `json:"t30_order_cnt"` // 投后30日累计订单成交量

	// 直播相关
	LivePlayed3s                     int64   `json:"live_played_3s"`                        // 直播观看3s
	LiveEventGoodsView               int64   `json:"live_event_goods_view"`                 // 直播间商品点击数
	LiveReward                       int64   `json:"live_reward"`                           // 直播打赏
	LiveComment                      int64   `json:"live_comment"`                          // 直播评论
	LiveShare                        int64   `json:"live_share"`                            // 直播分享
	ConversionLivePlay3sCost         float64 `json:"conversion_live_play_3_s_cost"`         // 直播观看3s成本
	ConversionLiveEventGoodsViewCost float64 `json:"conversion_live_event_goods_view_cost"` // 直播间商品点击成本
	LiveEventGoodsViewRatio          float64 `json:"live_event_goods_view_ratio"`           // 直播间商品点击率
}

// MerchantDetailReportResp 小店通转化数据报表响应数据（仅data部分）
type MerchantDetailReportResp struct {
	TotalCount int64                        `json:"total_count"` // 数据的总行数
	Details    []MerchantDetailReportDetail `json:"details"`     // 数据明细信息
}
