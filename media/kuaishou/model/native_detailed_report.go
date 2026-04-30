package model

import "errors"

// NativeDetailedReportReq 原生效果数据报表明细请求
type NativeDetailedReportReq struct {
	accessTokenReq
	AdvertiserId int64                  `json:"advertiser_id"` // 广告主账号ID，必填
	PageInfo     NatureDetailedPageInfo `json:"page_info"`     // 分页信息
}

func (receiver *NativeDetailedReportReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *NativeDetailedReportReq) Validate() (err error) {
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

// NativeDetailReportItem 原生报表明细条目
type NativeDetailReportItem struct {
	CampaignId                                  int64   `json:"campaign_id"`                                         // 计划ID
	CampaignName                                string  `json:"campaign_name"`                                       // 计划名称
	UnitId                                      int64   `json:"unit_id"`                                             // 广告组ID
	UnitName                                    string  `json:"unit_name"`                                           // 广告组名称
	CreativeId                                  int64   `json:"creative_id"`                                         // 创意ID
	CreativeName                                string  `json:"creative_name"`                                       // 创意名称
	CreateTime                                  string  `json:"create_time"`                                         // 创建时间 毫秒
	Roi                                         bool    `json:"roi"`                                                 // ROI
	IndirectSubmit0DCnt                         float64 `json:"indirect_submit_0_d_cnt"`                             // 当日间接表单提交
	IndirectSubmit7DCnt                         float64 `json:"indirect_submit_7_d_cnt"`                             // 7日间接表单提交
	T0IndirectConversionCnt                     float64 `json:"t_0_indirect_conversion_cnt"`                         // 当日间接激活数
	T7IndirectConversionCnt                     float64 `json:"t_7_indirect_conversion_cnt"`                         // 7日间接激活数
	T0IndirectPaiedCnt                          float64 `json:"t_0_indirect_paied_cnt"`                              // 当日间接付费次数
	T7IndirectPaiedCnt                          float64 `json:"t_7_indirect_paied_cnt"`                              // 7日间接付费次数
	T0IndirectPaiedAmt                          float64 `json:"t_0_indirect_paied_amt"`                              // 当日间接付费金额
	T7IndirectPaiedAmt                          float64 `json:"t_7_indirect_paied_amt"`                              // 7日间接付费金额
	T7IndirectedPaiedRoi                        float64 `json:"t_7_indirected_paied_roi"`                            // 7日间接ROI
	TotalCharge                                 int64   `json:"total_charge"`                                        // 花费(厘)
	Impression                                  int64   `json:"impression"`                                          // 封面曝光数
	PhotoClick                                  int64   `json:"photo_click"`                                         // 封面点击数
	Click                                       int64   `json:"click"`                                               // 素材曝光数
	ActionbarClick                              int64   `json:"actionbar_click"`                                     // 行为数
	PhotoClickRatio                             float64 `json:"photo_click_ratio"`                                   // 封面点击率
	Impression1kCost                            float64 `json:"impression_1_k_cost"`                                 // 平均千次封面曝光花费(币)
	Click1kCost                                 float64 `json:"click_1_k_cost"`                                      // 平均千次素材曝光花费(币)
	PhotoClickCost                              float64 `json:"photo_click_cost"`                                    // 平均封面点击单价(币)
	ActionCost                                  float64 `json:"action_cost"`                                         // 平均行为单价(币)
	AdShow                                      int64   `json:"ad_show"`                                             // 曝光数
	ActionNewRatio                              float64 `json:"action_new_ratio"`                                    // 行为率
	AdPhotoPlayed2SRatio                        float64 `json:"ad_photo_played_2_s_ratio"`                           // 2s播放率
	Play3SRatio                                 float64 `json:"play_3_s_ratio"`                                      // 3s播放率
	Play5SRatio                                 float64 `json:"play_5_s_ratio"`                                      // 5s播放率
	AdPhotoPlayed10SRatio                       float64 `json:"ad_photo_played_10_s_ratio"`                          // 10s播放率
	AdPhotoPlayed75PercentRatio                 float64 `json:"ad_photo_played_75_percent_ratio"`                    // 75%进度播放率
	PlayEndRatio                                float64 `json:"play_end_ratio"`                                      // 完播率
	Share                                       int64   `json:"share"`                                               // 分享数
	Comment                                     int64   `json:"comment"`                                             // 评论数
	Likes                                       int64   `json:"likes"`                                               // 点赞数
	Follow                                      int64   `json:"follow"`                                              // 新增关注数
	Report                                      int64   `json:"report"`                                              // 举报数
	Block                                       int64   `json:"block"`                                               // 拉黑数
	Negative                                    int64   `json:"negative"`                                            // 不感兴趣数
	PlayedNum                                   int64   `json:"played_num"`                                          // 播放数
	PlayedThreeSeconds                          int64   `json:"played_three_seconds"`                                // 3s播放数(有效播放数)
	AdPhotoPlayed10S                            int64   `json:"ad_photo_played_10_s"`                                // 10s播放数
	AdPhotoPlayed75Percent                      int64   `json:"ad_photo_played_75_percent"`                          // 75%播放进度数
	PlayedEnd                                   int64   `json:"played_end"`                                          // 完播数
	Conversion                                  int64   `json:"conversion"`                                          // 激活数
	ConversionCost                              float64 `json:"conversion_cost"`                                     // 激活单价
	KeyAction                                   int64   `json:"key_action"`                                          // 关键行为数
	KeyActionCost                               float64 `json:"key_action_cost"`                                     // 关键行为成本
	KeyActionRatio                              float64 `json:"key_action_ratio"`                                    // 关键行为率
	EventNewUserPay                             int64   `json:"event_new_user_pay"`                                  // 新增付费人数
	EventNewUserPayCost                         float64 `json:"event_new_user_pay_cost"`                             // 新增付费人数成本
	EventNewUserPayRatio                        float64 `json:"event_new_user_pay_ratio"`                            // 新增付费人数率
	EventPayFirstDay                            int64   `json:"event_pay_first_day"`                                 // 首日付费次数
	EventPayFirstDayCost                        float64 `json:"event_pay_first_day_cost"`                            // 首日付费次数成本
	EventPayPurchaseAmountFirstDay              float64 `json:"event_pay_purchase_amount_first_day"`                 // 激活当日付费金额
	EventPayFirstDayRoi                         float64 `json:"event_pay_first_day_roi"`                             // 激活当日ROI
	EventPayPurchaseAmountOneDay                float64 `json:"event_pay_purchase_amount_one_day"`                   // 激活后24h付费金额(回传时间)
	EventPayPurchaseAmountOneDayRoi             float64 `json:"event_pay_purchase_amount_one_day_roi"`               // 激活后24h-ROI(回传时间)
	EventPayPurchaseAmountOneDayByConversion    float64 `json:"event_pay_purchase_amount_one_day_by_conversion"`     // 激活后24h付费金额(激活时间)
	EventPayPurchaseAmountOneDayByConversionRoi float64 `json:"event_pay_purchase_amount_one_day_by_conversion_roi"` // 激活后24h-ROI(激活时间)
	EventPay                                    int64   `json:"event_pay"`                                           // 付费次数
	EventPayCost                                float64 `json:"event_pay_cost"`                                      // 付费次数成本
	EventPayPurchaseAmount                      float64 `json:"event_pay_purchase_amount"`                           // 付费金额
	T0DirectConversionCnt                       int64   `json:"t_0_direct_conversion_cnt"`                           // 激活数(计费时间)
	T0ConversionCnt                             int64   `json:"t_0_conversion_cnt"`                                  // 当日累计激活数
	T7ConversionCnt                             int64   `json:"t_7_conversion_cnt"`                                  // 7日累计激活数
	T0DirectPaiedCnt                            int64   `json:"t_0_direct_paied_cnt"`                                // 付费次数(计费时间)
	T0PaiedCnt                                  int64   `json:"t_0_paied_cnt"`                                       // 当日累计付费次数
	T7PaiedCnt                                  int64   `json:"t_7_paied_cnt"`                                       // 7日累计付费次数
	AccumulatedPaiedCost                        float64 `json:"accumulated_paied_cost"`                              // 累计付费次数成本
	T0DirectPaiedAmt                            float64 `json:"t_0_direct_paied_amt"`                                // 付费金额(计费时间)
	T0PaiedAmt                                  float64 `json:"t_0_paied_amt"`                                       // 当日累计付费金额
	T7PaiedAmt                                  float64 `json:"t_7_paied_amt"`                                       // 7日累计付费金额
	T0PaiedRoi                                  float64 `json:"t_0_paied_roi"`                                       // 当日累计ROI
	T7PaiedRoi                                  float64 `json:"t_7_paied_roi"`                                       // 7日累计ROI
	ConversionNum                               int64   `json:"conversion_num"`                                      // 转化数(回传时间)
	ConversionNumCost                           float64 `json:"conversion_num_cost"`                                 // 转化成本(回传时间)
	ConversionRatio                             float64 `json:"conversion_ratio"`                                    // 转化率(回传时间)
	DeepConversionNum                           int64   `json:"deep_conversion_num"`                                 // 深度转化数(回传时间)
	DeepConversionCost                          float64 `json:"deep_conversion_cost"`                                // 深度转化成本(回传时间)
	DeepConversionRatio                         float64 `json:"deep_conversion_ratio"`                               // 深度转化率(回传时间)
	DirectSubmit1DCnt                           int64   `json:"direct_submit_1_d_cnt"`                               // 表单提交数(计费时间)
	Submit1DCnt                                 int64   `json:"submit_1_d_cnt"`                                      // 当日累计表单提交
	Submit7DCnt                                 int64   `json:"submit_7_d_cnt"`                                      // 7日累计表单提交
	SubmitUnitPriceCost                         float64 `json:"submit_unit_price_cost"`                              // 累计表单提交单价
	EventValidClues                             int64   `json:"event_valid_clues"`                                   // 有效线索数
	EventValidClueCost                          float64 `json:"event_valid_clue_cost"`                               // 有效线索成本
	MerchantRecoFans                            int64   `json:"merchant_reco_fans"`                                  // 涨粉数
	MerchantRecoFansCost                        float64 `json:"merchant_reco_fans_cost"`                             // 涨粉成本
	IndirectEventPayCnt                         int64   `json:"indirect_event_pay_cnt"`                              // 间接转化数(回传时间)
	KolUserTypeDesc                             string  `json:"kol_user_type_desc"`                                  // 原生广告类型
	OcpcActionType                              string  `json:"ocpc_action_type"`                                    // 优化目标
	CampaignType                                string  `json:"campaign_type"`                                       // 营销目标
	AuthorId                                    string  `json:"author_id"`                                           // 快手号ID
}

// NativeDetailedReportResp 原生效果数据报表明细响应数据（仅data部分）
type NativeDetailedReportResp struct {
	ResultList []NativeDetailReportItem `json:"result_list"` // 报表明细
	Sum        []NativeDetailReportItem `json:"sum"`         // 全局汇总
	PageInfo   NatureDetailedPageInfo   `json:"page_info"`   // 分页信息
}
