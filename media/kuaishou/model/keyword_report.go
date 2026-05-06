package model

import "errors"

// KeywordReportReq 关键词报表请求
type KeywordReportReq struct {
	accessTokenReq
	AdvertiserId        int64    `json:"advertiser_id"`                  // 广告主ID，必填
	StartDate           string   `json:"start_date,omitempty"`           // 过滤筛选条件，格式 yyyy-MM-dd
	EndDate             string   `json:"end_date,omitempty"`             // 过滤筛选条件，格式 yyyy-MM-dd
	Page                int      `json:"page,omitempty"`                 // 页码，默认为1
	PageSize            int      `json:"page_size,omitempty"`            // 每页行数，默认为20，最大支持2000
	StartDateMin        string   `json:"start_date_min,omitempty"`       // 增量拉取开始时间，格式 yyyy-MM-dd HH:mm
	EndDateMin          string   `json:"end_date_min,omitempty"`         // 增量拉取结束时间，格式 yyyy-MM-dd HH:mm
	CampaignType        int      `json:"campaign_type,omitempty"`        // 计划类型
	ReportDims          []string `json:"report_dims,omitempty"`          // 报表维度
	TemporalGranularity string   `json:"temporal_granularity,omitempty"` // 时间粒度：DAILY=天粒度 HOURLY=小时粒度
	CampaignIds         []int64  `json:"campaign_ids,omitempty"`         // 计划ID集合
	UnitIds             []int64  `json:"unit_ids,omitempty"`             // 广告组ID集合，每次查询数量不超过5000
	Detailed            bool     `json:"detailed,omitempty"`             // 是否按unit_id分组
	WordInfoIds         []int64  `json:"word_info_ids,omitempty"`        // 推广关键词ID集合，每次查询数量不超过5000
	ExtendInfo          []string `json:"extend_info,omitempty"`          // 扩展信息
	SelectedColumns     []string `json:"selected_columns,omitempty"`     // 选择列
}

func (receiver *KeywordReportReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *KeywordReportReq) Validate() (err error) {
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

// KeywordReportDetail 关键词报表数据明细
type KeywordReportDetail struct {
	// 维度标识
	StatDate      string `json:"stat_date"`      // 数据日期，格式 YYYY-MM-DD
	StatHour      int64  `json:"stat_hour"`      // 数据小时
	WordInfoId    int64  `json:"word_info_id"`   // 推广关键词ID
	Word          string `json:"word"`           // 关键词文本
	MatchType     int    `json:"match_type"`     // 匹配类型：1=精确 2=短语 3=广泛
	CampaignId    int64  `json:"campaign_id"`    // 计划ID
	CampaignName  string `json:"campaign_name"`  // 计划名称
	UnitId        int64  `json:"unit_id"`        // 广告组ID
	UnitName      string `json:"unit_name"`      // 广告组名称
	AdScene       string `json:"ad_scene"`       // 广告场景
	PlacementType string `json:"placement_type"` // 投放类型

	// 核心消耗与曝光
	Charge               float64 `json:"charge"`                 // 花费（元）
	Show                 int64   `json:"show"`                   // 封面曝光数
	Aclick               int64   `json:"aclick"`                 // 素材曝光数
	Bclick               int64   `json:"bclick"`                 // 行为数
	PhotoClick           int64   `json:"photo_click"`            // 封面点击数
	PhotoClickRatio      float64 `json:"photo_click_ratio"`      // 封面点击率
	ActionRatio          float64 `json:"action_ratio"`           // 素材点击率
	ActionNewRatio       float64 `json:"action_new_ratio"`       // 新行为率
	ImpressionK1Cost     float64 `json:"impression_1k_cost"`     // 平均千次封面曝光花费（元）
	PhotoClickCost       float64 `json:"photo_click_cost"`       // 平均封面点击单价（元）
	Click1kCost          float64 `json:"click_1k_cost"`          // 平均千次素材曝光花费（元）
	ActionCost           float64 `json:"action_cost"`            // 平均行为单价（元）
	AdShow               float64 `json:"ad_show"`                // 广告展示数
	ClickConversionRatio float64 `json:"click_conversion_ratio"` // 点击转化率
	ConversionCost       float64 `json:"conversion_cost"`        // 转化成本
	RelevanceScoreAvg    float64 `json:"relevance_score_avg"`    // 素材相关性得分

	// 社交互动
	Share            int64 `json:"share"`              // 分享数
	Comment          int64 `json:"comment"`            // 评论数
	Like             int64 `json:"like"`               // 点赞数
	Follow           int64 `json:"follow"`             // 新增关注数
	CancelLike       int64 `json:"cancel_like"`        // 取消点赞数
	CancelFollow     int64 `json:"cancel_follow"`      // 取消关注数
	Report           int64 `json:"report"`             // 举报数
	Block            int64 `json:"block"`              // 拉黑数
	Negative         int64 `json:"negative"`           // 不感兴趣数
	EventNoIntention int64 `json:"event_no_intention"` // 用户无意向数

	// 视频播放
	AdPhotoPlayed2s             int64   `json:"ad_photo_played_2s"`              // 2秒播放数
	AdPhotoPlayed10s            int64   `json:"ad_photo_played_10s"`             // 10秒播放数
	AdPhotoPlayed75percent      int64   `json:"ad_photo_played_75percent"`       // 75%进度播放数
	PlayedThreeSeconds          int64   `json:"played_three_seconds"`            // 有效播放数
	PlayedFiveSeconds           int64   `json:"played_five_seconds"`             // 5秒播放数
	PlayedEnd                   int64   `json:"played_end"`                      // 完播数
	LivePlayed3s                int64   `json:"live_played_3s"`                  // 直播3秒播放数
	Play3sRatio                 float64 `json:"play_3s_ratio"`                   // 3秒播放率
	Play5sRatio                 float64 `json:"play_5s_ratio"`                   // 5秒播放率
	PlayEndRatio                float64 `json:"play_end_ratio"`                  // 完播率
	AdPhotoPlayed2sRatio        float64 `json:"ad_photo_played_2s_ratio"`        // 2秒播放率
	AdPhotoPlayed10sRatio       float64 `json:"ad_photo_played_10s_ratio"`       // 10秒播放率
	AdPhotoPlayed75percentRatio float64 `json:"ad_photo_played_75percent_ratio"` // 75%进度播放率

	// 应用下载与激活
	DownloadStarted         int64   `json:"download_started"`          // 应用下载开始数
	DownloadCompleted       int64   `json:"download_completed"`        // 应用下载完成数
	DownloadInstalled       int64   `json:"download_installed"`        // 应用安装数
	Activation              int64   `json:"activation"`                // 激活数
	Submit                  int64   `json:"submit"`                    // 提交数
	DownloadStartedCost     float64 `json:"download_started_cost"`     // 安卓下载开始成本（元）
	DownloadCompletedCost   float64 `json:"download_completed_cost"`   // 安卓下载完成成本（元）
	DownloadStartedRatio    float64 `json:"download_started_ratio"`    // 安卓下载开始率
	DownloadCompletedRatio  float64 `json:"download_completed_ratio"`  // 安卓下载完成率
	DownloadConversionRatio float64 `json:"download_conversion_ratio"` // 下载完成激活率

	// 注册与次留
	EventRegister         int64   `json:"event_register"`            // 注册数
	EventRegisterCost     float64 `json:"event_register_cost"`       // 注册成本
	EventRegisterRatio    float64 `json:"event_register_ratio"`      // 注册率
	EventNextDayStay      int64   `json:"event_next_day_stay"`       // 次日留存数
	EventNextDayStayCost  float64 `json:"event_next_day_stay_cost"`  // 次日留存成本
	EventNextDayStayRatio float64 `json:"event_next_day_stay_ratio"` // 次日留存率
	EventAppInvoked       int64   `json:"event_app_invoked"`         // 应用唤起数
	EventAppInvokedCost   float64 `json:"event_app_invoked_cost"`    // 应用唤起成本
	EventAppInvokedRatio  float64 `json:"event_app_invoked_ratio"`   // 应用唤起率
	EventWatchAppAd       int64   `json:"event_watch_app_ad"`        // 观看应用广告数

	// 付款与订单
	EventPay                                    int64   `json:"event_pay"`                                           // 付款笔数
	EventPayPurchaseAmount                      float64 `json:"event_pay_purchase_amount"`                           // 付款金额（元）
	EventPayRoi                                 float64 `json:"event_pay_roi"`                                       // 付款ROI
	EventPayFirstDay                            int64   `json:"event_pay_first_day"`                                 // 首日付款笔数
	EventPayPurchaseAmountFirstDay              float64 `json:"event_pay_purchase_amount_first_day"`                 // 首日付款金额（元）
	EventPayFirstDayRoi                         float64 `json:"event_pay_first_day_roi"`                             // 首日ROI
	EventPayWeightedPurchaseAmount              float64 `json:"event_pay_weighted_purchase_amount"`                  // 加权付款金额（元）
	EventPayWeightedPurchaseAmountFirstDay      float64 `json:"event_pay_weighted_purchase_amount_first_day"`        // 首日加权付款金额（元）
	EventPayPurchaseAmountOneDay                float64 `json:"event_pay_purchase_amount_one_day"`                   // 24小时付款金额（元）
	EventPayPurchaseAmountOneDayRoi             float64 `json:"event_pay_purchase_amount_one_day_roi"`               // 24小时付款ROI
	EventPayPurchaseAmountOneDayByConversion    float64 `json:"event_pay_purchase_amount_one_day_by_conversion"`     // 24小时转化口径付款金额（元）
	EventPayPurchaseAmountOneDayByConversionRoi float64 `json:"event_pay_purchase_amount_one_day_by_conversion_roi"` // 24小时转化口径付款ROI
	EventNewUserPay                             int64   `json:"event_new_user_pay"`                                  // 新用户付款笔数
	EventNewUserPayCost                         float64 `json:"event_new_user_pay_cost"`                             // 新用户付款成本
	EventNewUserPayRatio                        float64 `json:"event_new_user_pay_ratio"`                            // 新用户付款率
	EventOrderSubmit                            int64   `json:"event_order_submit"`                                  // 提交订单数
	EventOrderPaid                              int64   `json:"event_order_paid"`                                    // 订单付款数
	EventOrderPaidPurchaseAmount                float64 `json:"event_order_paid_purchase_amount"`                    // 订单付款金额（元）
	EventOrderPaidCost                          float64 `json:"event_order_paid_cost"`                               // 订单付款成本
	EventOrderPaidRoi                           float64 `json:"event_order_paid_roi"`                                // 订单付款ROI
	EventOrderSuccessed                         int64   `json:"event_order_successed"`                               // 订单完成数
	UnionEventPayPurchaseAmount7d               float64 `json:"union_event_pay_purchase_amount_7d"`                  // 联盟7日付款金额（元）
	UnionEventPayPurchaseAmount7dRoi            float64 `json:"union_event_pay_purchase_amount_7d_roi"`              // 联盟7日付款ROI

	// 电商
	AdProductCnt             int64   `json:"ad_product_cnt"`               // 商品成交数
	EventGoodsView           int64   `json:"event_goods_view"`             // 商品浏览数
	EventGoodsViewCost       float64 `json:"event_goods_view_cost"`        // 商品浏览成本
	LiveEventGoodsView       int64   `json:"live_event_goods_view"`        // 直播商品浏览数
	MerchantRecoFans         int64   `json:"merchant_reco_fans"`           // 涨粉数
	MerchantRecoFansCost     float64 `json:"merchant_reco_fans_cost"`      // 涨粉成本
	EventOrderAmountRoi      float64 `json:"event_order_amount_roi"`       // 小店通ROI
	EventAddShoppingCart     int64   `json:"event_add_shopping_cart"`      // 加购数
	EventAddShoppingCartCost float64 `json:"event_add_shopping_cart_cost"` // 加购成本

	// 表单/线索
	FormCount                                  int64   `json:"form_count"`                                      // 表单提交数（线索数）
	FormCost                                   float64 `json:"form_cost"`                                       // 线索成本
	FormActionRatio                            float64 `json:"form_action_ratio"`                               // 表单提交点击率
	EventValidClues                            int64   `json:"event_valid_clues"`                               // 有效线索数
	EventValidCluesCost                        float64 `json:"event_valid_clues_cost"`                          // 有效线索成本
	EventFormSubmit                            int64   `json:"event_form_submit"`                               // 表单提交数（追踪时间口径）
	EventFormSubmitRatio                       float64 `json:"event_form_submit_ratio"`                         // 表单提交率
	EventFormSubmitCost                        float64 `json:"event_form_submit_cost"`                          // 表单提交成本
	DirectSubmit1dCnt                          int64   `json:"direct_submit_1d_cnt"`                            // 1日直接表单提交数
	DirectSubmit1dCost                         float64 `json:"direct_submit_1d_cost"`                           // 1日直接表单提交成本
	EventAppointForm                           int64   `json:"event_appoint_form"`                              // 预约表单提交数
	EventAppointFormCost                       float64 `json:"event_appoint_form_cost"`                         // 预约表单成本
	EventAppointFormRatio                      float64 `json:"event_appoint_form_ratio"`                        // 预约表单提交率
	EventAppointJumpClick                      int64   `json:"event_appoint_jump_click"`                        // 预约跳转点击数
	EventAppointJumpClickCost                  float64 `json:"event_appoint_jump_click_cost"`                   // 预约跳转点击成本
	EventAppointJumpClickRatio                 float64 `json:"event_appoint_jump_click_ratio"`                  // 预约跳转点击率
	EventButtonClick                           int64   `json:"event_button_click"`                              // 按钮点击数
	EventButtonClickCost                       float64 `json:"event_button_click_cost"`                         // 按钮点击成本
	EventButtonClickRatio                      float64 `json:"event_button_click_ratio"`                        // 按钮点击率
	EventConsultationValidRetained             int64   `json:"event_consultation_valid_retained"`               // 有效咨询留资数
	EventConsultationValidRetainedCost         float64 `json:"event_consultation_valid_retained_cost"`          // 有效咨询留资成本
	EventConsultationValidRetainedRatio        float64 `json:"event_consultation_valid_retained_ratio"`         // 有效咨询留资率
	EventPreComponentConsultationValidRetained int64   `json:"event_pre_component_consultation_valid_retained"` // 前置组件有效咨询留资数
	EventConversionClickCost                   float64 `json:"event_conversion_click_cost"`                     // 转化点击成本
	EventConversionClickRatio                  float64 `json:"event_conversion_click_ratio"`                    // 转化点击率
	EventMultiConversion                       int64   `json:"event_multi_conversion"`                          // 多目标转化数
	EventMultiConversionRatio                  float64 `json:"event_multi_conversion_ratio"`                    // 多目标转化率
	EventMultiConversionCost                   float64 `json:"event_multi_conversion_cost"`                     // 多目标转化成本
	EventDspGiftForm                           int64   `json:"event_dsp_gift_form"`                             // DSP礼品表单提交数

	// 关键行为
	KeyAction      int64   `json:"key_action"`       // 关键行为数
	KeyActionCost  float64 `json:"key_action_cost"`  // 关键行为成本
	KeyActionRatio float64 `json:"key_action_ratio"` // 关键行为率

	// 电话/通话
	EventMakingCalls        int64   `json:"event_making_calls"`        // 拨打电话数
	EventMakingCallsCost    float64 `json:"event_making_calls_cost"`   // 拨打电话成本
	EventMakingCallsRatio   float64 `json:"event_making_calls_ratio"`  // 拨打电话率
	EventOutboundCall       int64   `json:"event_outbound_call"`       // 外呼电话数
	EventOutboundCallCost   float64 `json:"event_outbound_call_cost"`  // 外呼电话成本
	EventOutboundCallRatio  float64 `json:"event_outbound_call_ratio"` // 外呼电话率
	EventPhoneGetThrough    int64   `json:"event_phone_get_through"`   // 电话接通数
	EventGetThrough         int64   `json:"event_get_through"`         // 智能电话接通数
	EventGetThroughCost     float64 `json:"event_get_through_cost"`    // 智能电话接通成本
	EventGetThroughRatio    float64 `json:"event_get_through_ratio"`   // 智能电话接通率
	EventIntentionConfirmed int64   `json:"event_intention_confirmed"` // 意向确认数

	// 微信
	EventAddWechat             int64   `json:"event_add_wechat"`                // 微信复制数
	EventAddWechatCost         float64 `json:"event_add_wechat_cost"`           // 微信复制成本
	EventAddWechatRatio        float64 `json:"event_add_wechat_ratio"`          // 微信复制率
	EventWechatQrCodeLinkClick int64   `json:"event_wechat_qr_code_link_click"` // 微信小程序深度加粉数
	EventWechatConnected       int64   `json:"event_wechat_connected"`          // 微信加粉数

	// 金融（金融贷款/信用卡）
	EventJinJianApp                          int64   `json:"event_jin_jian_app"`                              // 进件完成数
	EventJinJianAppCost                      float64 `json:"event_jin_jian_app_cost"`                         // 进件完成成本
	EventJinJianLandingPage                  int64   `json:"event_jin_jian_landing_page"`                     // 落地页进件数
	EventJinJianLandingPageCost              float64 `json:"event_jin_jian_landing_page_cost"`                // 落地页进件成本
	EventCreditGrantApp                      int64   `json:"event_credit_grant_app"`                          // 授信完成数
	EventCreditGrantAppCost                  float64 `json:"event_credit_grant_app_cost"`                     // 授信完成成本
	EventCreditGrantAppRatio                 float64 `json:"event_credit_grant_app_ratio"`                    // 授信完成率
	EventCreditGrantLandingPage              int64   `json:"event_credit_grant_landing_page"`                 // 落地页授信数
	EventCreditGrantLandingPageCost          float64 `json:"event_credit_grant_landing_page_cost"`            // 落地页授信成本
	EventCreditGrantLandingRatio             float64 `json:"event_credit_grant_landing_ratio"`                // 落地页授信率
	EventCreditGrantFirstDayApp              int64   `json:"event_credit_grant_first_day_app"`                // 首日APP授信数
	EventCreditGrantFirstDayAppCost          float64 `json:"event_credit_grant_first_day_app_cost"`           // 首日APP授信成本
	EventCreditGrantFirstDayAppRatio         float64 `json:"event_credit_grant_first_day_app_ratio"`          // 首日APP授信率
	EventCreditGrantFirstDayLandingPage      int64   `json:"event_credit_grant_first_day_landing_page"`       // 首日落地页授信数
	EventCreditGrantFirstDayLandingPageCost  float64 `json:"event_credit_grant_first_day_landing_page_cost"`  // 首日落地页授信成本
	EventCreditGrantFirstDayLandingPageRatio float64 `json:"event_credit_grant_first_day_landing_page_ratio"` // 首日落地页授信率
	EventCreditCardRecheck                   int64   `json:"event_credit_card_recheck"`                       // 信用卡核卡数
	EventCreditCardRecheckFirstDay           int64   `json:"event_credit_card_recheck_first_day"`             // 信用卡首日核卡数
	EventNewUserJinjianApp                   int64   `json:"event_new_user_jinjian_app"`                      // 新用户APP进件数
	EventNewUserJinjianAppCost               float64 `json:"event_new_user_jinjian_app_cost"`                 // 新用户APP进件成本
	EventNewUserJinjianAppRoi                float64 `json:"event_new_user_jinjian_app_roi"`                  // 新用户APP进件ROI
	EventNewUserCreditGrantApp               int64   `json:"event_new_user_credit_grant_app"`                 // 新用户APP授信数
	EventNewUserCreditGrantAppCost           float64 `json:"event_new_user_credit_grant_app_cost"`            // 新用户APP授信成本
	EventNewUserCreditGrantAppRoi            float64 `json:"event_new_user_credit_grant_app_roi"`             // 新用户APP授信ROI
	EventNewUserJinjianPage                  int64   `json:"event_new_user_jinjian_page"`                     // 新用户落地页进件数
	EventNewUserJinjianPageCost              float64 `json:"event_new_user_jinjian_page_cost"`                // 新用户落地页进件成本
	EventNewUserJinjianPageRoi               float64 `json:"event_new_user_jinjian_page_roi"`                 // 新用户落地页进件ROI
	EventNewUserCreditGrantPage              int64   `json:"event_new_user_credit_grant_page"`                // 新用户落地页授信数
	EventNewUserCreditGrantPageCost          float64 `json:"event_new_user_credit_grant_page_cost"`           // 新用户落地页授信成本
	EventNewUserCreditGrantPageRoi           float64 `json:"event_new_user_credit_grant_page_roi"`            // 新用户落地页授信ROI
	EventPhoneCardActivate                   int64   `json:"event_phone_card_activate"`                       // 电话卡激活数

	// 广告观看次数
	EventAdWatchTimes        int64   `json:"event_ad_watch_times"`          // 广告观看次数
	EventAdWatchTimesRatio   float64 `json:"event_ad_watch_times_ratio"`    // 广告观看率
	EventAdWatchTimesCost    float64 `json:"event_ad_watch_times_cost"`     // 广告观看成本
	EventAdWatch5Times       int64   `json:"event_ad_watch_5_times"`        // 广告观看5次数
	EventAdWatch5TimesCost   float64 `json:"event_ad_watch_5_times_cost"`   // 广告观看5次成本
	EventAdWatch5TimesRatio  float64 `json:"event_ad_watch_5_times_ratio"`  // 广告观看5次率
	EventAdWatch10Times      int64   `json:"event_ad_watch_10_times"`       // 广告观看10次数
	EventAdWatch10TimesCost  float64 `json:"event_ad_watch_10_times_cost"`  // 广告观看10次成本
	EventAdWatch10TimesRatio float64 `json:"event_ad_watch_10_times_ratio"` // 广告观看10次率
	EventAdWatch20Times      int64   `json:"event_ad_watch_20_times"`       // 广告观看20次数
	EventAdWatch20TimesCost  float64 `json:"event_ad_watch_20_times_cost"`  // 广告观看20次成本
	EventAdWatch20TimesRatio float64 `json:"event_ad_watch_20_times_ratio"` // 广告观看20次率

	// 教育/试听
	EventAudition         int64   `json:"event_audition"`            // 试听数
	EventAudition30dCnt   int64   `json:"event_audition_30d_cnt"`    // 30日试听数
	EventAuditionCost     float64 `json:"event_audition_cost"`       // 试听成本
	AllLessonFinishCnt    int64   `json:"all_lesson_finish_cnt"`     // 全部课程完成数
	AllLessonFinish30dCnt int64   `json:"all_lesson_finish_30d_cnt"` // 30日全部课程完成数
	HighPriceClassPayCnt  int64   `json:"high_price_class_pay_cnt"`  // 高价课付款数

	// 量房/其他
	EventMeasurementHouse int64 `json:"event_measurement_house"` // 量房数

	// 约看/其他
	ApproxPayCost  float64 `json:"approx_pay_cost"`  // 阿里近似购买成本
	ApproxPayCount int64   `json:"approx_pay_count"` // 阿里近似购买数
	ApproxPayRatio float64 `json:"approx_pay_ratio"` // 阿里近似购买率
}

// KeywordReportResp 关键词报表响应数据（仅data部分）
type KeywordReportResp struct {
	TotalCount int64                 `json:"total_count"` // 数据的总行数
	Details    []KeywordReportDetail `json:"details"`     // 数据明细信息
}
