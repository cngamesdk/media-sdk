package model

import "errors"

// CreativeReportReq 广告创意数据请求（自定义，不包括省心投物料）
type CreativeReportReq struct {
	accessTokenReq
	AdvertiserId        int64    `json:"advertiser_id"`                  // 广告主id，必填
	StartDate           string   `json:"start_date,omitempty"`           // 查询开始日期，格式 yyyy-MM-dd，只能查最近一周，时间跨度不能超一周
	EndDate             string   `json:"end_date,omitempty"`             // 查询结束日期，格式 yyyy-MM-dd，只能查最近一周，时间跨度不能超一周
	Page                int      `json:"page,omitempty"`                 // 请求的页码，默认为1
	PageSize            int      `json:"page_size,omitempty"`            // 每页行数，默认为20，最大支持2000
	StartDateMin        string   `json:"start_date_min,omitempty"`       // 增量拉取开始时间，格式 yyyy-MM-dd HH:mm
	EndDateMin          string   `json:"end_date_min,omitempty"`         // 增量拉取结束时间，格式 yyyy-MM-dd HH:mm
	CampaignType        int      `json:"campaign_type,omitempty"`        // 计划类型：1=作品推广 2=提升应用安装 3=获取电商下单 4=推广品牌活动 5=收集销售线索 6=保量广告 7=提高应用活跃
	ReportDims          []string `json:"report_dims,omitempty"`          // 维度拆分：adScene=按广告场景，placementType=按广告范围(快手/联盟)
	TemporalGranularity string   `json:"temporal_granularity,omitempty"` // 时间粒度：DAILY=天粒度，HOURLY=小时粒度，默认天粒度
	CampaignIds         []int64  `json:"campaign_ids,omitempty"`         // 广告计划ID集合，过滤条件，单次最多5000
	UnitIds             []int64  `json:"unit_ids,omitempty"`             // 广告组ID集合，过滤条件，单次最多5000
	CreativeIds         []int64  `json:"creative_ids,omitempty"`         // 广告创意ID集合，过滤条件，单次最多5000
	ExtendInfo          []string `json:"extend_info,omitempty"`          // 额外信息：photo=获取视频信息(视频ID和视频MD5)
}

func (receiver *CreativeReportReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *CreativeReportReq) Validate() (err error) {
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

// CreativeReportDetail 广告创意数据明细
type CreativeReportDetail struct {
	// 日期/时间
	StatDate string `json:"stat_date"` // 数据日期，格式：YYYY-MM-DD
	StatHour int64  `json:"stat_hour"` // 数据小时

	// 创意/组/计划标识
	CreativeId   int64  `json:"creative_id"`   // 广告创意ID
	CreativeName string `json:"creative_name"` // 广告创意名称
	Status       int64  `json:"status"`        // 状态：1=投放中 2=已暂停 3=已删除
	UnitId       int64  `json:"unit_id"`       // 广告组ID
	UnitName     string `json:"unit_name"`     // 广告组名称
	CampaignId   int64  `json:"campaign_id"`   // 广告计划ID
	CampaignName string `json:"campaign_name"` // 广告计划名称
	PhotoId      int64  `json:"photo_id"`      // 视频ID（需extend_info=["photo"]）
	PhotoMd5     string `json:"photo_md5"`     // 视频MD5（需extend_info=["photo"]）

	// 维度字段
	AdScene       string `json:"adScene"`        // 广告场景
	AdSceneSnake  string `json:"ad_scene"`       // 广告场景
	PlacementType string `json:"placement_type"` // 广告范围

	// 核心消耗与互动
	Charge   float64 `json:"charge"`   // 花费（元）
	Show     int64   `json:"show"`     // 封面曝光数
	Aclick   int64   `json:"aclick"`   // 素材曝光数
	Bclick   int64   `json:"bclick"`   // 行为数
	Share    int64   `json:"share"`    // 分享数
	Comment  int64   `json:"comment"`  // 评论数
	Like     int64   `json:"like"`     // 点赞数
	Follow   int64   `json:"follow"`   // 新增粉丝数
	Report   int64   `json:"report"`   // 举报数
	Block    int64   `json:"block"`    // 拉黑数
	Negative int64   `json:"negative"` // 减少此类作品数

	// 点击/率指标
	PhotoClick           int64   `json:"photo_click"`            // 封面点击数
	PhotoClickRatio      float64 `json:"photo_click_ratio"`      // 封面点击率
	ActionRatio          float64 `json:"action_ratio"`           // 行为率
	Impression1kCost     float64 `json:"impression_1k_cost"`     // 平均千次曝光花费（元）
	PhotoClickCost       float64 `json:"photo_click_cost"`       // 平均点击单价（元）
	Click1kCost          float64 `json:"click_1k_cost"`          // 平均千次素材曝光花费（元）
	ActionCost           float64 `json:"action_cost"`            // 平均行为单价（元）
	AdShow               float64 `json:"ad_show"`                // 广告曝光
	ActionNewRatio       float64 `json:"action_new_ratio"`       // 行为率 新
	ClickConversionRatio float64 `json:"click_conversion_ratio"` // 点击激活成本

	// 播放相关
	PlayedNum                   int64   `json:"played_num"`                      // 播放数
	AdPhotoPlayed10s            int64   `json:"ad_photo_played_10s"`             // 10s播放数
	AdPhotoPlayed2s             int64   `json:"ad_photo_played_2s"`              // 2s播放数
	AdPhotoPlayed75percent      int64   `json:"ad_photo_played_75percent"`       // 75%进度播放数
	PlayedEnd                   int64   `json:"played_end"`                      // 播放完成
	PlayedFiveSeconds           int64   `json:"played_five_seconds"`             // 播放5s
	PlayedThreeSeconds          int64   `json:"played_three_seconds"`            // 有效播放数
	Play3sRatio                 float64 `json:"play_3s_ratio"`                   // 3s播放率
	Play5sRatio                 float64 `json:"play_5s_ratio"`                   // 5s播放率
	PlayEndRatio                float64 `json:"play_end_ratio"`                  // 完播率
	AdPhotoPlayed75percentRatio float64 `json:"ad_photo_played_75percent_ratio"` // 75%进度播放率
	AdPhotoPlayed10sRatio       float64 `json:"ad_photo_played_10s_ratio"`       // 10s播放率
	AdPhotoPlayed2sRatio        float64 `json:"ad_photo_played_2s_ratio"`        // 2s播放率
	LivePlayed3s                int64   `json:"live_played_3s"`                  // 直播观看数

	// 社交互动
	CancelLike   int64 `json:"cancel_like"`   // 取消点赞数
	CancelFollow int64 `json:"cancel_follow"` // 取消关注数

	// 转化指标
	ConversionCost float64 `json:"conversion_cost"` // 激活单价

	// 私信相关
	PrivateMessageSentCost  float64 `json:"private_message_sent_cost"`  // 私信消息转化成本
	PrivateMessageSentRatio float64 `json:"private_message_sent_ratio"` // 私信消息转化率
	PrivateMessageSentCnt   int64   `json:"private_message_sent_cnt"`   // 私信消息数

	// 直接私信留资
	LeadsSubmitCost     float64 `json:"leads_submit_cost"`      // 直接私信留资成本
	LeadsSubmitCntRatio float64 `json:"leads_submit_cnt_ratio"` // 直接私信留资率
	LeadsSubmitCnt      int64   `json:"leads_submit_cnt"`       // 直接私信留资数

	// 应用下载
	Activation              int64   `json:"activation"`                // 激活数
	Submit                  int64   `json:"submit"`                    // 提交按钮点击数（历史字段，同form_count）
	DownloadStarted         int64   `json:"download_started"`          // 安卓下载开始数
	DownloadCompleted       int64   `json:"download_completed"`        // 安卓下载完成数
	DownloadInstalled       int64   `json:"download_installed"`        // 安卓安装完成数
	DownloadCompletedCost   float64 `json:"download_completed_cost"`   // 安卓下载完成单价（元）
	DownloadCompletedRatio  float64 `json:"download_completed_ratio"`  // 安卓下载完成率
	DownloadConversionRatio float64 `json:"download_conversion_ratio"` // 下载完成激活率
	DownloadStartedCost     float64 `json:"download_started_cost"`     // 安卓下载开始单价（元）
	DownloadStartedRatio    float64 `json:"download_started_ratio"`    // 安卓下载开始率

	// 唤醒应用
	EventAppInvoked      int64   `json:"event_app_invoked"`       // 唤醒应用数
	EventAppInvokedCost  float64 `json:"event_app_invoked_cost"`  // 唤醒应用成本
	EventAppInvokedRatio float64 `json:"event_app_invoked_ratio"` // 唤醒应用率

	// 应用下载-付费
	EventPayFirstDay               float64 `json:"event_pay_first_day"`                 // 首日付费次数
	EventPayPurchaseAmountFirstDay float64 `json:"event_pay_purchase_amount_first_day"` // 首日付费金额
	EventPayFirstDayRoi            float64 `json:"event_pay_first_day_roi"`             // 首日ROI
	EventPay                       int64   `json:"event_pay"`                           // 付费次数
	EventPayPurchaseAmount         float64 `json:"event_pay_purchase_amount"`           // 付费金额
	EventPayRoi                    float64 `json:"event_pay_roi"`                       // ROI

	// 应用下载-注册
	EventRegister      int64   `json:"event_register"`       // 注册数
	EventRegisterCost  float64 `json:"event_register_cost"`  // 注册成本
	EventRegisterRatio float64 `json:"event_register_ratio"` // 注册率

	// 应用下载-完件/授信
	EventJinJianApp          int64   `json:"event_jin_jian_app"`           // 完件数
	EventJinJianAppCost      float64 `json:"event_jin_jian_app_cost"`      // 完件成本
	EventCreditGrantApp      int64   `json:"event_credit_grant_app"`       // 授信数
	EventCreditGrantAppCost  float64 `json:"event_credit_grant_app_cost"`  // 授信成本
	EventCreditGrantAppRatio float64 `json:"event_credit_grant_app_ratio"` // 授信率

	// 应用下载-付款成功
	EventOrderPaid               int64   `json:"event_order_paid"`                 // 付款成功数
	EventOrderPaidPurchaseAmount float64 `json:"event_order_paid_purchase_amount"` // 付款成功金额
	EventOrderPaidCost           float64 `json:"event_order_paid_cost"`            // 单次付款成本
	EventOrderPaidRoi            float64 `json:"event_order_paid_roi"`             // 订单支付率

	// 应用下载-次留
	EventNextDayStayCost  float64 `json:"event_next_day_stay_cost"`  // 次留成本（仅支持分日查询）
	EventNextDayStayRatio float64 `json:"event_next_day_stay_ratio"` // 次留率（仅支持分日查询）
	EventNextDayStay      int64   `json:"event_next_day_stay"`       // 次留数（仅支持分日查询）

	// 新增付费人数
	EventNewUserPay      int64   `json:"event_new_user_pay"`       // 新增付费人数
	EventNewUserPayCost  float64 `json:"event_new_user_pay_cost"`  // 新增付费人数成本
	EventNewUserPayRatio float64 `json:"event_new_user_pay_ratio"` // 新增付费人数率

	// 落地页数据
	FormCount                       int64   `json:"form_count"`                           // 线索提交个数
	FormCost                        float64 `json:"form_cost"`                            // 单个线索成本
	FormActionRatio                 float64 `json:"form_action_ratio"`                    // 表单提交点击率
	EventJinJianLandingPage         int64   `json:"event_jin_jian_landing_page"`          // 落地页完件数
	EventJinJianLandingPageCost     float64 `json:"event_jin_jian_landing_page_cost"`     // 落地页完件成本
	EventCreditGrantLandingPage     int64   `json:"event_credit_grant_landing_page"`      // 落地页授信数
	EventCreditGrantLandingPageCost float64 `json:"event_credit_grant_landing_page_cost"` // 落地页授信成本
	EventCreditGrantLandingRatio    float64 `json:"event_credit_grant_landing_ratio"`     // 落地页授信率

	// 落地页-有效线索
	EventValidClues     int64   `json:"event_valid_clues"`      // 有效线索数
	EventValidCluesCost float64 `json:"event_valid_clues_cost"` // 有效线索成本

	// 落地页-多转化
	EventMultiConversion      int64   `json:"event_multi_conversion"`       // 落地页多转化次数
	EventMultiConversionRatio float64 `json:"event_multi_conversion_ratio"` // 落地页多转化率
	EventMultiConversionCost  float64 `json:"event_multi_conversion_cost"`  // 落地页多转化成本

	// 淘系近似购买
	ApproxPayCost  float64 `json:"approx_pay_cost"`  // 淘系近似购买成本
	ApproxPayCount int64   `json:"approx_pay_count"` // 近似购买数
	ApproxPayRatio float64 `json:"approx_pay_ratio"` // 淘系近似购买率

	// 电商/商家
	AdProductCnt         int64   `json:"ad_product_cnt"`          // 商品成交数
	EventGoodsView       int64   `json:"event_goods_view"`        // 商品访问数
	EventGoodsViewCost   float64 `json:"event_goods_view_cost"`   // 商品访问成本
	MerchantRecoFans     int64   `json:"merchant_reco_fans"`      // 涨粉量
	MerchantRecoFansCost float64 `json:"merchant_reco_fans_cost"` // 涨粉成本
	LiveEventGoodsView   int64   `json:"live_event_goods_view"`   // 直播间商品点击数
	EventOrderAmountRoi  float64 `json:"event_order_amount_roi"`  // 小店推广roi

	// 添加购物车
	EventAddShoppingCart     int64   `json:"event_add_shopping_cart"`      // 添加购物车数
	EventAddShoppingCartCost float64 `json:"event_add_shopping_cart_cost"` // 添加购物车成本

	// 提交订单
	EventOrderSubmit int64 `json:"event_order_submit"` // 提交订单数

	// 附加组件表单
	EventDspGiftForm int64 `json:"event_dsp_gift_form"` // 附加组件表单提交

	// 广告观看（次数）
	EventAdWatch10Times      int64   `json:"event_ad_watch_10_times"`       // 10次广告观看数
	EventAdWatch10TimesCost  float64 `json:"event_ad_watch_10_times_cost"`  // 10次广告观看成本
	EventAdWatch10TimesRatio float64 `json:"event_ad_watch_10_times_ratio"` // 10次广告观看转化率
	EventAdWatch20Times      int64   `json:"event_ad_watch_20_times"`       // 20次广告观看数
	EventAdWatch20TimesCost  float64 `json:"event_ad_watch_20_times_cost"`  // 20次广告观看成本
	EventAdWatch20TimesRatio float64 `json:"event_ad_watch_20_times_ratio"` // 20次广告观看转化率
	EventAdWatch5Times       int64   `json:"event_ad_watch_5_times"`        // 5次广告观看数
	EventAdWatch5TimesCost   float64 `json:"event_ad_watch_5_times_cost"`   // 5次广告观看成本
	EventAdWatch5TimesRatio  float64 `json:"event_ad_watch_5_times_ratio"`  // 5次广告观看转化率

	// 广告观看
	EventWatchAppAd        int64   `json:"event_watch_app_ad"`         // 广告观看
	EventAdWatchTimes      int64   `json:"event_ad_watch_times"`       // 广告观看次数
	EventAdWatchTimesRatio float64 `json:"event_ad_watch_times_ratio"` // 广告观看次数转化率
	EventAdWatchTimesCost  float64 `json:"event_ad_watch_times_cost"`  // 广告观看次数成本

	// 教育/试听
	EventAudition                       int64   `json:"event_audition"`                          // 首次试听到课数
	EventConsultationValidRetained      int64   `json:"event_consultation_valid_retained"`       // 留咨咨询数
	EventConsultationValidRetainedCost  float64 `json:"event_consultation_valid_retained_cost"`  // 留咨咨询成本
	EventConsultationValidRetainedRatio float64 `json:"event_consultation_valid_retained_ratio"` // 留咨咨询率
	EventConversionClickCost            float64 `json:"event_conversion_click_cost"`             // 有效咨询成本
	EventConversionClickRatio           float64 `json:"event_conversion_click_ratio"`            // 有效咨询率
	EventAudition30dCnt                 int64   `json:"event_audition_30d_cnt"`                  // 首次试听到课（归因）
	EventAuditionCost                   float64 `json:"event_audition_cost"`                     // 首次试听到课成本
	AllLessonFinishCnt                  int64   `json:"all_lesson_finish_cnt"`                   // 全部试听完课（回传）
	AllLessonFinish30dCnt               int64   `json:"all_lesson_finish_30d_cnt"`               // 全部试听完课（归因）
	HighPriceClassPayCnt                int64   `json:"high_price_class_pay_cnt"`                // 成交付费（回传）
	HighPriceClassPay30dCnt             int64   `json:"high_price_class_pay_30d_cnt"`            // 成交付费（归因）

	// 首日授信(App)
	EventCreditGrantFirstDayApp      int64   `json:"event_credit_grant_first_day_app"`       // app首日授信数
	EventCreditGrantFirstDayAppCost  float64 `json:"event_credit_grant_first_day_app_cost"`  // 首日授信成本
	EventCreditGrantFirstDayAppRatio float64 `json:"event_credit_grant_first_day_app_ratio"` // 首日授信率

	// 首日授信(落地页)
	EventCreditGrantFirstDayLandingPage      int64   `json:"event_credit_grant_first_day_landing_page"`       // 落地页首日授信数
	EventCreditGrantFirstDayLandingPageCost  float64 `json:"event_credit_grant_first_day_landing_page_cost"`  // 落地页首日授信成本
	EventCreditGrantFirstDayLandingPageRatio float64 `json:"event_credit_grant_first_day_landing_page_ratio"` // 落地页首日授信率

	// 电话拨打
	EventMakingCalls      int64   `json:"event_making_calls"`       // 电话拨打数
	EventMakingCallsCost  float64 `json:"event_making_calls_cost"`  // 电话拨打成本
	EventMakingCallsRatio float64 `json:"event_making_calls_ratio"` // 电话拨打率

	// 激活后24h付费
	EventPayPurchaseAmountOneDay                float64 `json:"event_pay_purchase_amount_one_day"`                   // 激活后24h付费金额(回传时间)
	EventPayPurchaseAmountOneDayByConversion    float64 `json:"event_pay_purchase_amount_one_day_by_conversion"`     // 激活后24h付费金额(激活时间)
	EventPayPurchaseAmountOneDayByConversionRoi float64 `json:"event_pay_purchase_amount_one_day_by_conversion_roi"` // 激活后24小时付费ROI
	EventPayPurchaseAmountOneDayRoi             float64 `json:"event_pay_purchase_amount_one_day_roi"`               // 激活后24h-ROI(回传时间)

	// 加权付费（保险行业）
	EventPayWeightedPurchaseAmount         float64 `json:"event_pay_weighted_purchase_amount"`           // 加权付费金额
	EventPayWeightedPurchaseAmountFirstDay float64 `json:"event_pay_weighted_purchase_amount_first_day"` // 首日加权付费金额

	// 激活后15日付费
	EventPayPurchaseAmount15DayByConversion float64 `json:"event_pay_purchase_amount_15_day_by_conversion"` // 激活后15日付费金额

	// 整体ROI（内购&广告）（新增）
	EventPayWeekOverallRoi     float64 `json:"event_pay_week_overall_roi"`      // 激活后七日整体ROI（内购&广告）
	EventPayThreeDayOverallRoi float64 `json:"event_pay_three_day_overall_roi"` // 激活后三日整体ROI（内购&广告）
	EventPayFirstDayOverallRoi float64 `json:"event_pay_first_day_overall_roi"` // 激活当日整体首日ROI（内购&广告）
	EventPay30DayOverallRoi    float64 `json:"event_pay_30_day_overall_roi"`    // 激活后30日整体ROI
	EventPay15DayOverallRoi    float64 `json:"event_pay_15_day_overall_roi"`    // 激活后15日整体ROI

	// 小游戏IAA广告变现
	MinigameIaaPurchaseAmountWeekByConversionRoi     float64 `json:"minigame_iaa_purchase_amount_week_by_conversion_roi"`      // 激活后七日广告变现ROI
	MinigameIaaPurchaseAmountThreeDayByConversionRoi float64 `json:"minigame_iaa_purchase_amount_three_day_by_conversion_roi"` // 激活后三日广告变现ROI
	MinigameIaaPurchaseAmountFirstDayRoi             float64 `json:"minigame_iaa_purchase_amount_first_day_roi"`               // 当日广告变现ROI
	MinigameIaaPurchaseAmountWeekByConversion        float64 `json:"minigame_iaa_purchase_amount_week_by_conversion"`          // 激活后七日广告LTV
	MinigameIaaPurchaseAmountThreeDayByConversion    float64 `json:"minigame_iaa_purchase_amount_three_day_by_conversion"`     // 激活后三日广告LTV
	MinigameIaaPurchaseAmountFirstDay                float64 `json:"minigame_iaa_purchase_amount_first_day"`                   // 当日广告LTV
	MinigameIaaPurchaseRoi                           float64 `json:"minigame_iaa_purchase_roi"`                                // IAA广告变现ROI
	MinigameIaaPurchaseAmount                        float64 `json:"minigame_iaa_purchase_amount"`                             // IAA广告变现LTV
	MinigameIaaPurchaseAmount30DayByConversionRoi    float64 `json:"minigame_iaa_purchase_amount_30_day_by_conversion_roi"`    // 激活后30日广告变现ROI
	MinigameIaaPurchaseAmount15DayByConversionRoi    float64 `json:"minigame_iaa_purchase_amount_15_day_by_conversion_roi"`    // 激活后15日广告变现ROI
	MinigameIaaPurchaseAmount30DayByConversion       float64 `json:"minigame_iaa_purchase_amount_30_day_by_conversion"`        // 激活后30日广告LTV
	MinigameIaaPurchaseAmount15DayByConversion       float64 `json:"minigame_iaa_purchase_amount_15_day_by_conversion"`        // 激活后15日广告LTV

	// 有效获客
	MmuEffectiveCustomerAcquisition7dCnt int64   `json:"mmu_effective_customer_acquisition_7d_cnt"` // MMU识别产生的有效获客数（计费）
	MmuEffectiveCustomerAcquisitionCnt   int64   `json:"mmu_effective_customer_acquisition_cnt"`    // MMU识别产生的有效获客数（回传）
	EffectiveCustomerAcquisition7dRatio  float64 `json:"effective_customer_acquisition_7d_ratio"`   // 有效获客率（计费）
	EffectiveCustomerAcquisition7dCost   float64 `json:"effective_customer_acquisition_7d_cost"`    // 有效获客成本（计费）
	EffectiveCustomerAcquisition7dCnt    int64   `json:"effective_customer_acquisition_7d_cnt"`     // 有效获客数（计费）

	// 附加组件
	EventPreComponentConsultationValidRetained int64 `json:"event_pre_component_consultation_valid_retained"` // 附加咨询组件留资咨询数
	EventWechatQrCodeLinkClick                 int64 `json:"event_wechat_qr_code_link_click"`                 // 微信小程序深度加粉数

	// 按钮点击
	EventButtonClick      int64   `json:"event_button_click"`       // 按钮点击数
	EventButtonClickCost  float64 `json:"event_button_click_cost"`  // 按钮点击成本
	EventButtonClickRatio float64 `json:"event_button_click_ratio"` // 按钮点击率

	// 新增完件/授信人数(App)
	EventNewUserJinjianApp         int64   `json:"event_new_user_jinjian_app"`           // 新增完件人数
	EventNewUserJinjianAppCost     float64 `json:"event_new_user_jinjian_app_cost"`      // 新增完件人数成本
	EventNewUserJinjianAppRoi      float64 `json:"event_new_user_jinjian_app_roi"`       // 新增完件人数率
	EventNewUserCreditGrantApp     int64   `json:"event_new_user_credit_grant_app"`      // 新增授信人数
	EventNewUserCreditGrantAppCost float64 `json:"event_new_user_credit_grant_app_cost"` // 新增授信人数成本
	EventNewUserCreditGrantAppRoi  float64 `json:"event_new_user_credit_grant_app_roi"`  // 新增授信人数率

	// 新增完件/授信人数(落地页)
	EventNewUserJinjianPage         int64   `json:"event_new_user_jinjian_page"`           // 新增完件人数(落地页)
	EventNewUserJinjianPageCost     float64 `json:"event_new_user_jinjian_page_cost"`      // 新增完件人数成本(落地页)
	EventNewUserJinjianPageRoi      float64 `json:"event_new_user_jinjian_page_roi"`       // 新增完件人数率(落地页)
	EventNewUserCreditGrantPage     int64   `json:"event_new_user_credit_grant_page"`      // 新增授信人数(落地页)
	EventNewUserCreditGrantPageCost float64 `json:"event_new_user_credit_grant_page_cost"` // 新增授信人数成本(落地页)
	EventNewUserCreditGrantPageRoi  float64 `json:"event_new_user_credit_grant_page_roi"`  // 新增授信人数率(落地页)

	// 预约表单
	EventAppointForm           int64   `json:"event_appoint_form"`             // 预约表单数
	EventAppointFormCost       float64 `json:"event_appoint_form_cost"`        // 预约表单点击成本
	EventAppointFormRatio      float64 `json:"event_appoint_form_ratio"`       // 预约表单点击率
	EventAppointJumpClick      int64   `json:"event_appoint_jump_click"`       // 预约跳转点击数
	EventAppointJumpClickCost  float64 `json:"event_appoint_jump_click_cost"`  // 预约跳转点击成本
	EventAppointJumpClickRatio float64 `json:"event_appoint_jump_click_ratio"` // 预约跳转点击率

	// 联盟广告收入
	UnionEventPayPurchaseAmount7d    float64 `json:"union_event_pay_purchase_amount_7d"`     // 联盟广告收入
	UnionEventPayPurchaseAmount7dRoi float64 `json:"union_event_pay_purchase_amount_7d_roi"` // 联盟变现ROI

	// 微信复制
	EventAddWechat      int64   `json:"event_add_wechat"`       // 微信复制数
	EventAddWechatCost  float64 `json:"event_add_wechat_cost"`  // 微信复制成本
	EventAddWechatRatio float64 `json:"event_add_wechat_ratio"` // 微信复制率

	// 智能电话
	EventGetThrough      int64   `json:"event_get_through"`       // 智能电话-确认接通数
	EventGetThroughCost  float64 `json:"event_get_through_cost"`  // 智能电话-确认接通成本
	EventGetThroughRatio float64 `json:"event_get_through_ratio"` // 智能电话-确认接通率

	// 电话/意向/微信/成交
	EventPhoneGetThrough    int64 `json:"event_phone_get_through"`   // 电话建联数
	EventIntentionConfirmed int64 `json:"event_intention_confirmed"` // 意向确认数
	EventWechatConnected    int64 `json:"event_wechat_connected"`    // 微信加粉数
	EventOrderSuccessed     int64 `json:"event_order_successed"`     // 有效线索成交数
	EventPhoneCardActivate  int64 `json:"event_phone_card_activate"` // 电话卡激活数
	EventMeasurementHouse   int64 `json:"event_measurement_house"`   // 量房数

	// 外呼电话
	EventOutboundCall      int64   `json:"event_outbound_call"`       // 电话拨打数
	EventOutboundCallCost  float64 `json:"event_outbound_call_cost"`  // 电话拨打成本
	EventOutboundCallRatio float64 `json:"event_outbound_call_ratio"` // 电话拨打率

	// 关键行为
	KeyAction      int64   `json:"key_action"`       // 关键行为数
	KeyActionCost  float64 `json:"key_action_cost"`  // 关键行为成本
	KeyActionRatio float64 `json:"key_action_ratio"` // 关键行为率

	// 信用卡核卡
	EventCreditCardRecheck         int64 `json:"event_credit_card_recheck"`           // 信用卡核卡数
	EventCreditCardRecheckFirstDay int64 `json:"event_credit_card_recheck_first_day"` // 信用卡首日核卡数

	// 用户无意向
	EventNoIntention int64 `json:"event_no_intention"` // 用户无意向数

	// T0/T3完件
	Jinjian0dCnt     int64   `json:"jinjian_0d_cnt"`      // T0完件数
	Jinjian3dCnt     int64   `json:"jinjian_3d_cnt"`      // T3完件数
	Jinjian0dCntCost float64 `json:"jinjian_0d_cnt_cost"` // T0完件成本
	Jinjian3dCntCost float64 `json:"jinjian_3d_cnt_cost"` // T3完件成本

	// T0/T3授信
	CreditGrant0dCnt      int64   `json:"credit_grant_0d_cnt"`       // T0授信数
	CreditGrant3dCnt      int64   `json:"credit_grant_3d_cnt"`       // T3授信数
	CreditGrant0dCntCost  float64 `json:"credit_grant_0d_cnt_cost"`  // T0授信成本
	CreditGrant3dCntCost  float64 `json:"credit_grant_3d_cnt_cost"`  // T3授信成本
	CreditGrant0dCntRatio float64 `json:"credit_grant_0d_cnt_ratio"` // T0完件授信率
	CreditGrant3dCntRatio float64 `json:"credit_grant_3d_cnt_ratio"` // T3完件授信通过率

	// T0/T3全量授信
	KeyInappAction0dCnt      int64   `json:"key_inapp_action_0d_cnt"`       // T0全量授信数
	KeyInappAction3dCnt      int64   `json:"key_inapp_action_3d_cnt"`       // T3全量授信数
	KeyInappAction0dCntCost  float64 `json:"key_inapp_action_0d_cnt_cost"`  // T0全量授信成本
	KeyInappAction3dCntCost  float64 `json:"key_inapp_action_3d_cnt_cost"`  // T3全量授信成本
	KeyInappAction0dCntRatio float64 `json:"key_inapp_action_0d_cnt_ratio"` // T0全量授信通过率
	KeyInappAction3dCntRatio float64 `json:"key_inapp_action_3d_cnt_ratio"` // T3全量授信通过率

	// T0用信
	DrawCreditLine0dCnt      int64   `json:"draw_credit_line_0d_cnt"`       // T0用信数
	DrawCreditLine0dCntCost  float64 `json:"draw_credit_line_0d_cnt_cost"`  // T0用信成本
	DrawCreditLine0dCntRatio float64 `json:"draw_credit_line_0d_cnt_ratio"` // T0授信用信率

	// 表单提交
	EventFormSubmit      int64   `json:"event_form_submit"`       // 表单提交数（回传时间）
	DirectSubmit1dCnt    int64   `json:"direct_submit_1d_cnt"`    // 表单提交数(计费时间)
	EventFormSubmitRatio float64 `json:"event_form_submit_ratio"` // 表单提交率（回传时间）
	EventFormSubmitCost  float64 `json:"event_form_submit_cost"`  // 表单提交成本（回传时间）
	DirectSubmit1dCost   float64 `json:"direct_submit_1d_cost"`   // 表单提交成本
}

// CreativeReportResp 广告创意数据响应数据（仅data部分）
type CreativeReportResp struct {
	TotalCount int64                  `json:"total_count"` // 数据的总行数
	Details    []CreativeReportDetail `json:"details"`     // 数据明细信息
}
