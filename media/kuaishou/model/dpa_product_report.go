package model

import "errors"

// DpaProductReportReq 商品库报表请求
type DpaProductReportReq struct {
	accessTokenReq
	AdvertiserId int64  `json:"advertiser_id"`       // 投放账号ID，必填
	LibraryId    int64  `json:"library_id"`          // 商品库ID，必填
	StartDate    string `json:"start_date"`          // 起始时间，格式 YYYY-MM-DD，与end_date相差不能超过半年，必填
	EndDate      string `json:"end_date"`            // 截止时间，格式 YYYY-MM-DD，与start_date相差不能超过半年，必填
	Page         int    `json:"page,omitempty"`      // 页码，默认为1
	PageSize     int    `json:"page_size,omitempty"` // 页大小，默认为20
}

func (receiver *DpaProductReportReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *DpaProductReportReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.LibraryId <= 0 {
		err = errors.New("library_id is empty")
		return
	}
	return
}

// DpaProductReportDetail 商品库报表单商品数据明细
type DpaProductReportDetail struct {
	// 商品信息
	AdvertiserId    int64  `json:"advertiser_id"`     // 投放账号ID
	LibraryId       int64  `json:"library_id"`        // 商品库ID
	OuterId         string `json:"outer_id"`          // 商品ID
	ProductName     string `json:"product_name"`      // 商品名称
	Title           string `json:"title"`             // 商品标题
	Category        string `json:"category"`          // 商品类目
	Theme           string `json:"theme"`             // 商品题材
	SubIndustryName string `json:"sub_industry_name"` // 商品类型
	PutStatus       int    `json:"put_status"`        // 商品投放状态：0=关 1=开
	CreateTime      int64  `json:"create_time"`       // 商品创建时间
	UpdateTime      int64  `json:"update_time"`       // 商品最后更新时间

	// DPA专属指标
	DpaCostTotal        int64   `json:"dpa_cost_total"`        // 花费
	DpaImpressionNum    int64   `json:"dpa_impression_num"`    // 商品曝光数
	DpaImpression1kCost float64 `json:"dpa_impression1k_cost"` // 商品千次曝光成本
	AdItemClick         int64   `json:"ad_item_click"`         // 行为数
	AdItemActionCost    float64 `json:"ad_item_action_cost"`   // 行为成本
	AdItemActionRatio   float64 `json:"ad_item_action_ratio"`  // 行为率

	// 转化汇总
	TotalConversionNum   float64 `json:"total_conversion_num"`   // 总的转化数
	TotalConversionCost  float64 `json:"total_conversion_cost"`  // 转化成本
	TotalConversionRatio float64 `json:"total_conversion_ratio"` // 转化率

	// 曝光/点击
	AdShow          int64   `json:"ad_show"`           // 广告曝光
	Impression      int64   `json:"impression"`        // 封面曝光数
	PhotoClick      int64   `json:"photo_click"`       // 封面点击数
	PhotoClickRatio float64 `json:"photo_click_ratio"` // 封面点击率
	Click           int64   `json:"click"`             // 素材曝光数
	ActionRatio     float64 `json:"action_ratio"`      // 素材点击率

	// 成本指标
	Impression1kCost float64 `json:"impression_1k_cost"` // 平均千次封面曝光花费(元)
	Click1kCost      float64 `json:"click_1k_cost"`      // 平均千次素材曝光花费
	PhotoClickCost   float64 `json:"photo_click_cost"`   // 平均封面点击单价

	// 下载相关
	DownloadStarted         int64   `json:"download_started"`          // 安卓下载开始数
	DownloadStartedRatio    float64 `json:"download_started_ratio"`    // 安卓开始下载率
	DownloadCompleted       int64   `json:"download_completed"`        // 安卓下载完成数
	DownloadCompletedCost   float64 `json:"download_completed_cost"`   // 安卓下载完成单价
	DownloadCompletedRatio  float64 `json:"download_completed_ratio"`  // 安卓下载完成率
	DownloadConversionRatio float64 `json:"download_conversion_ratio"` // 下载完成激活率

	// 激活
	Conversion     int64   `json:"conversion"`      // 激活数
	ConversionCost float64 `json:"conversion_cost"` // 激活单价

	// 新增付费人数
	EventNewUserPay      int64   `json:"event_new_user_pay"`       // 新增付费人数
	EventNewUserPayCost  float64 `json:"event_new_user_pay_cost"`  // 新增付费人数成本
	EventNewUserPayRatio float64 `json:"event_new_user_pay_ratio"` // 新增付费人数率

	// 首日付费
	EventPayFirstDay               int64   `json:"event_pay_first_day"`                 // 首日付费次数
	EventPayFirstDayCost           float64 `json:"event_pay_first_day_cost"`            // 首日付费次数成本
	EventPayPurchaseAmountFirstDay float64 `json:"event_pay_purchase_amount_first_day"` // 激活当日付费金额
	EventPayFirstDayRoi            float64 `json:"event_pay_first_day_roi"`             // 激活当日ROI

	// 激活后24h次日留存
	Event24hStay                 int64   `json:"event_24h_stay"`                    // 激活后24h次日留存数（激活时间）
	Event24hStayRatio            float64 `json:"event_24h_stay_ratio"`              // 激活后24h次日留存率（激活时间）
	Event24hStayByConversionCost float64 `json:"event_24h_stay_by_conversion_cost"` // 激活后24h次日留存成本（激活时间）

	// 激活后24h付费
	EventPayPurchaseAmountOneDayByConversion    float64 `json:"event_pay_purchase_amount_one_day_by_conversion"`     // 激活后24h付费金额(激活时间)
	EventPayPurchaseAmountOneDayByConversionRoi float64 `json:"event_pay_purchase_amount_one_day_by_conversion_roi"` // 激活后24h-ROI(激活时间)

	// 自然日次日留存
	EventNextDayStay      float64 `json:"event_next_day_stay"`       // 自然日次日留存数（回传时间）
	EventNextDayStayCost  float64 `json:"event_next_day_stay_cost"`  // 自然日次日留存成本（回传时间）
	EventNextDayStayRatio float64 `json:"event_next_day_stay_ratio"` // 自然日次日留存率（回传时间）
}

// DpaProductReportPageInfo 商品库报表分页信息
type DpaProductReportPageInfo struct {
	TotalCount  int64 `json:"total_count"`  // 总行数
	CurrentPage int   `json:"current_page"` // 当前页码
	PageSize    int   `json:"page_size"`    // 每页大小
}

// DpaProductReportResp 商品库报表响应数据（仅data部分）
type DpaProductReportResp struct {
	PageInfo   DpaProductReportPageInfo `json:"page_info"`   // 分页信息
	ResultList []DpaProductReportDetail `json:"result_list"` // 单商品报表数据集
	Sum        []DpaProductReportDetail `json:"sum"`         // 全局汇总
}
