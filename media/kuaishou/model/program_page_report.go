package model

import "errors"

// ProgramPageReportReq 程序化落地页报表请求
type ProgramPageReportReq struct {
	accessTokenReq
	AdvertiserId int64    `json:"advertiser_id"`         // 广告主ID，必填
	StartDate    string   `json:"start_date,omitempty"`  // 开始时间，格式 yyyy-MM-dd
	EndDate      string   `json:"end_date,omitempty"`    // 结束时间，格式 yyyy-MM-dd
	ReportDims   []string `json:"report_dims,omitempty"` // 广告场景：adScene=按广告场景 placementType=按广告范围(快手/联盟)
	Page         int      `json:"page,omitempty"`        // 请求页码，默认为1
	PageSize     int      `json:"page_size,omitempty"`   // 每页行数，默认为20，最大支持1000
}

func (receiver *ProgramPageReportReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *ProgramPageReportReq) Validate() (err error) {
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

// ProgramPageReportDetail 程序化落地页报表数据明细
type ProgramPageReportDetail struct {
	// 日期/时间
	ReportDate     string `json:"report_date"`      // 日期，yyyy-MM-dd
	ReportHour     int64  `json:"report_hour"`      // 小时
	ReportDateHour string `json:"report_date_hour"` // 日期+小时

	// 页面信息
	PageId      string `json:"page_id"`       // 页面ID
	PageUrl     string `json:"page_url"`      // 页面URL
	PageName    string `json:"page_name"`     // 页面名称
	GroupId     string `json:"group_id"`      // 页面组ID
	GroupType   int    `json:"group_type"`    // 组类型
	GroupName   string `json:"group_name"`    // 组名称
	ImgCoverUrl string `json:"img_cover_url"` // 落地页截图URL

	// 状态
	ViewStatus       int    `json:"view_status"`        // 状态值
	ViewStatusReason string `json:"view_status_reason"` // 状态描述

	// 维度字段
	AdScene       string `json:"ad_scene"`       // 投放场景
	PlacementType string `json:"placement_type"` // 投放位置

	// 核心消耗与曝光
	TotalCharge    int64 `json:"total_charge"`    // 总消费
	Impression     int64 `json:"impression"`      // 封面曝光数
	Click          int64 `json:"click"`           // 素材曝光数
	PhotoClick     int64 `json:"photo_click"`     // 封面点击数
	ActionbarClick int64 `json:"actionbar_click"` // 行为数
	AdShow         int64 `json:"ad_show"`         // 曝光数

	// 播放相关
	PlayedThreeSeconds     int64 `json:"played_three_seconds"`     // 有效播放数
	PlayedFiveSeconds      int64 `json:"played_five_seconds"`      // 播放5s
	PlayedEnd              int64 `json:"played_end"`               // 播放完成
	AdPhotoPlayed75percent int64 `json:"ad_photo_played75percent"` // 75%进度播放数
	AdPhotoPlayed10s       int64 `json:"ad_photo_played10s"`       // 10s播放数
	AdPhotoPlayed2s        int64 `json:"ad_photo_played2s"`        // 2s播放数

	// 社交互动
	Likes      int64   `json:"likes"`       // 点赞数
	CancelLike int64   `json:"cancel_like"` // 取消点赞数
	Comment    int64   `json:"comment"`     // 评论数
	Follow     int64   `json:"follow"`      // 新增关注数
	Unfollow   int64   `json:"unfollow"`    // 取消关注数
	Share      int64   `json:"share"`       // 分享数
	Report     int64   `json:"report"`      // 举报数
	Block      int64   `json:"block"`       // 拉黑数
	Negative   float64 `json:"negative"`    // 减少此类作品数

	// 表单/转化
	FormCount                            float64 `json:"form_count"`                               // 表单提交数
	Conversion                           int64   `json:"conversion"`                               // 激活数
	DownloadCompleted                    int64   `json:"download_completed"`                       // 安卓下载完成数
	DownloadStarted                      int64   `json:"download_started"`                         // 安卓开始下载数
	DownloadInstalled                    int64   `json:"download_installed"`                       // 安卓安装完成数
	EventLandingpageStartedDownloadClick int64   `json:"event_landingpage_started_download_click"` // 开始下载

	// 付费
	EventPayFirstDay               int64   `json:"event_pay_first_day"`                 // 首日付费次数
	EventPayPurchaseAmountFirstDay float64 `json:"event_pay_purchase_amount_first_day"` // 激活当日付费金额
	EventPay                       int64   `json:"event_pay"`                           // 付费次数
	EventPayPurchaseAmount         float64 `json:"event_pay_purchase_amount"`           // 付费金额

	// 注册/留存
	EventRegister    int64 `json:"event_register"`      // 注册数
	EventNextDayStay int64 `json:"event_next_day_stay"` // 次日留存数

	// 完件/授信
	EventJinJianApp             int64 `json:"event_jinjian_app"`               // App完件数
	EventCreditGrantApp         int64 `json:"event_credit_grant_app"`          // App授信数
	EventJinJianLandingPage     int64 `json:"event_jinjian_landing_page"`      // 落地页完件数
	EventCreditGrantLandingPage int64 `json:"event_credit_grant_landing_page"` // 落地页授信数

	// 订单
	EventOrderPaid               int64   `json:"event_order_paid"`                 // 订单支付数
	EventOrderPaidPurchaseAmount float64 `json:"event_order_paid_purchase_amount"` // 订单成交金额
	EventOrderSubmit             int64   `json:"event_order_submit"`               // 提交订单数
	EventAddShoppingCart         int64   `json:"event_add_shopping_cart"`          // 添加购物车数

	// 近似购买
	ApproxPayCount int64 `json:"approx_pay_count"` // 近似购买数

	// 咨询/微信/电话
	EventConversionClick int64 `json:"event_conversion_click"` // 有效咨询数
	EventAddWechat       int64 `json:"event_add_wechat"`       // 微信调起数
	EventMakingCalls     int64 `json:"event_making_calls"`     // 电话拨打数

	// 多转化/广告观看
	EventMultiConversion int64 `json:"event_multi_conversion"` // 落地页多转化次数
	EventWatchAppAd      int64 `json:"event_watch_app_ad"`     // 当日APP内广告观看总次数
	EventAdWatchTimes    int64 `json:"event_ad_watch_times"`   // 当日APP内广告观看次数达到5次的人数

	// 电话/意向/微信/成交
	EventPhoneGetThrough    int64 `json:"event_phone_get_through"`   // 电话建联数
	EventIntentionConfirmed int64 `json:"event_intention_confirmed"` // 意向确认数
	EventWechatConnected    int64 `json:"event_wechat_connected"`    // 微信加粉数
	EventOrderSuccessed     int64 `json:"event_order_successed"`     // 有效线索成交数
	EventPhoneCardActivate  int64 `json:"event_phone_card_activate"` // 电话卡激活数
	EventMeasurementHouse   int64 `json:"event_measurement_house"`   // 量房数

	// 率指标
	ClickRatio                  float64 `json:"click_ratio"`                    // 点击率
	PhotoClickRatio             float64 `json:"photo_click_ratio"`              // 封面点击率
	Play3sRatio                 float64 `json:"play3s_ratio"`                   // 3s播放率
	ActionRatio                 float64 `json:"action_ratio"`                   // 素材点击率
	Play3sActionRatio           float64 `json:"play3s_action_ratio"`            // 有效播放行为率
	DownloadStartedRatio        float64 `json:"download_started_ratio"`         // 安卓下载开始率
	DownloadCompletedRatio      float64 `json:"download_completed_ratio"`       // 安卓下载完成率
	DownloadConversionRatio     float64 `json:"download_conversion_ratio"`      // 下载完成激活率
	ClickConversionRatio        float64 `json:"click_conversion_ratio"`         // 点击激活率
	FormActionRatio             float64 `json:"form_action_ratio"`              // 表单提交率
	Play5sRatio                 float64 `json:"play5s_ratio"`                   // 5s播放率
	PlayEndRatio                float64 `json:"play_end_ratio"`                 // 完播率
	ActionNewRatio              float64 `json:"action_new_ratio"`               // 行为率(新)
	AdPhotoPlayed75percentRatio float64 `json:"ad_photo_played75percent_ratio"` // 75%进度播放率
	AdPhotoPlayed10sRatio       float64 `json:"ad_photo_played10s_ratio"`       // 10s播放率
	AdPhotoPlayed2sRatio        float64 `json:"ad_photo_played2s_ratio"`        // 2s播放率

	// 成本指标
	FormCost              float64 `json:"form_cost"`               // 表单提交单价
	Impression1kCost      float64 `json:"impression1k_cost"`       // 平均千次封面曝光花费
	Click1kCost           float64 `json:"click1k_cost"`            // 平均千次素材曝光花费
	Play3sCost            float64 `json:"play3s_cost"`             // 平均有效播放单价
	PhotoClickCost        float64 `json:"photo_click_cost"`        // 平均封面点击单价
	ClickCost             float64 `json:"click_cost"`              // 单次点击成本
	ActionCost            float64 `json:"action_cost"`             // 平均行为单价
	ConversionCost        float64 `json:"conversion_cost"`         // 激活单价
	DownloadStartedCost   float64 `json:"download_started_cost"`   // 安卓下载开始单价
	DownloadCompletedCost float64 `json:"download_completed_cost"` // 安卓下载完成单价
	ThousandShowCost      float64 `json:"thousand_show_cost"`      // 平均千次广告曝光花费(元)

	// 付费ROI/成本
	EventPayFirstDayRoi  float64 `json:"event_pay_first_day_roi"`  // 激活当日ROI
	EventPayRoi          float64 `json:"event_pay_roi"`            // 付费ROI
	EventPayFirstDayCost float64 `json:"event_pay_first_day_cost"` // 首日付费次数成本
	EventPayCost         float64 `json:"event_pay_cost"`           // 付费次数成本

	// 注册/留存 成本率
	EventRegisterCost     float64 `json:"event_register_cost"`       // 注册成本
	EventRegisterRatio    float64 `json:"event_register_ratio"`      // 注册率
	EventNextDayStayCost  float64 `json:"event_next_day_stay_cost"`  // 次日留存成本
	EventNextDayStayRatio float64 `json:"event_next_day_stay_ratio"` // 次日留存率

	// 完件/授信 成本率
	EventJinJianAppCost              float64 `json:"event_jinjian_app_cost"`                // App完件成本
	EventCreditGrantAppCost          float64 `json:"event_credit_grant_app_cost"`           // App授信成本
	EventCreditGrantAppRatio         float64 `json:"event_credit_grant_app_ratio"`          // App授信率
	EventJinJianLandingPageCost      float64 `json:"event_jinjian_landing_page_cost"`       // 落地页完件成本
	EventCreditGrantLandingPageCost  float64 `json:"event_credit_grant_landing_page_cost"`  // 落地页授信成本
	EventCreditGrantLandingPageRatio float64 `json:"event_credit_grant_landing_page_ratio"` // 落地页授信率

	// 订单 成本
	EventOrderPaidCost       float64 `json:"event_order_paid_cost"`        // 订单支付成本
	EventOrderPaidRoi        float64 `json:"event_order_paid_roi"`         // 订单支付率
	EventOrderSubmitCost     float64 `json:"event_order_submit_cost"`      // 提交订单成本
	EventAddShoppingCartCost float64 `json:"event_add_shopping_cart_cost"` // 添加购物车成本

	// 近似购买 成本率
	ApproxPayRatio float64 `json:"approx_pay_ratio"` // 近似购买率
	ApproxPayCost  float64 `json:"approx_pay_cost"`  // 近似购买成本

	// 首日授信
	EventCreditGrantFirstDayApp              int64   `json:"event_credit_grant_first_day_app"`                // app首日授信数
	EventCreditGrantFirstDayLandingPage      int64   `json:"event_credit_grant_first_day_landing_page"`       // 落地页首日授信数
	EventCreditGrantFirstDayAppCost          float64 `json:"event_credit_grant_first_day_app_cost"`           // app首日授信成本
	EventCreditGrantFirstDayAppRatio         float64 `json:"event_credit_grant_first_day_app_ratio"`          // app首日授信率
	EventCreditGrantFirstDayLandingPageCost  float64 `json:"event_credit_grant_first_day_landing_page_cost"`  // 落地页首日授信成本
	EventCreditGrantFirstDayLandingPageRatio float64 `json:"event_credit_grant_first_day_landing_page_ratio"` // 落地页首日授信率

	// 新增付费人数
	EventNewUserPay      int64   `json:"event_new_user_pay"`       // 新增付费人数
	EventNewUserPayCost  float64 `json:"event_new_user_pay_cost"`  // 新增付费人数成本
	EventNewUserPayRatio float64 `json:"event_new_user_pay_ratio"` // 新增付费人数率

	// 有效线索
	EventValidClues    int64   `json:"event_valid_clues"`     // 有效线索数
	EventValidClueCost float64 `json:"event_valid_clue_cost"` // 有效线索成本

	// 电商/商家
	AdProductCnt         int64   `json:"ad_product_cnt"`          // 商品成交数
	EventGoodsView       int64   `json:"event_goods_view"`        // 商品访问数
	MerchantRecoFans     int64   `json:"merchant_reco_fans"`      // 涨粉数
	EventOrderAmountRoi  float64 `json:"event_order_amount_roi"`  // 小店推广roi
	EventGoodsViewCost   float64 `json:"event_goods_view_cost"`   // 商品访问成本
	MerchantRecoFansCost float64 `json:"merchant_reco_fans_cost"` // 涨粉成本

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
	EventNewUserJinjianPage         int64   `json:"event_new_user_jinjian_page"`           // 落地页新增完件人数
	EventNewUserJinjianPageCost     float64 `json:"event_new_user_jinjian_page_cost"`      // 落地页新增完件人数成本
	EventNewUserJinjianPageRoi      float64 `json:"event_new_user_jinjian_page_roi"`       // 落地页新增完件人数率
	EventNewUserCreditGrantPage     int64   `json:"event_new_user_credit_grant_page"`      // 落地页新增授信人数
	EventNewUserCreditGrantPageCost float64 `json:"event_new_user_credit_grant_page_cost"` // 落地页新增授信人数成本
	EventNewUserCreditGrantPageRoi  float64 `json:"event_new_user_credit_grant_page_roi"`  // 落地页新增授信人数率

	// 电话拨打/咨询/微信 成本率
	EventMakingCallsCost      float64 `json:"event_making_calls_cost"`      // 电话拨打成本
	EventMakingCallsRatio     float64 `json:"event_making_calls_ratio"`     // 电话拨打率
	EventConversionClickCost  float64 `json:"event_conversion_click_cost"`  // 有效咨询成本
	EventConversionClickRatio float64 `json:"event_conversion_click_ratio"` // 有效咨询率
	EventAddWechatCost        float64 `json:"event_add_wechat_cost"`        // 微信调起成本
	EventAddWechatRatio       float64 `json:"event_add_wechat_ratio"`       // 微信调起率

	// 直播相关
	LivePlayed3s            int64   `json:"live_played3s"`               // 直播观看3秒数
	LivePlay3sCost          float64 `json:"live_play3s_cost"`            // 直播观看3秒成本
	LiveRecoFansRatio       float64 `json:"live_reco_fans_ratio"`        // 直播涨粉率
	LiveEventGoodsView      int64   `json:"live_event_goods_view"`       // 直播间商品点击数
	LiveEventGoodsViewCost  float64 `json:"live_event_goods_view_cost"`  // 直播间商品点击成本
	LiveEventGoodsViewRatio float64 `json:"live_event_goods_view_ratio"` // 直播间商品点击率
	LiveEventOrderPaidRoi   float64 `json:"live_event_order_paid_roi"`   // 直播订单支付率

	// 预约表单
	EventAppointForm           int64   `json:"event_appoint_form"`             // 预约表单数
	EventAppointFormCost       float64 `json:"event_appoint_form_cost"`        // 预约表单点击成本
	EventAppointFormRatio      float64 `json:"event_appoint_form_ratio"`       // 预约表单点击率
	EventAppointJumpClick      int64   `json:"event_appoint_jump_click"`       // 预约跳转点击数
	EventAppointJumpClickCost  float64 `json:"event_appoint_jump_click_cost"`  // 预约跳转点击成本
	EventAppointJumpClickRatio float64 `json:"event_appoint_jump_click_ratio"` // 预约跳转点击率

	// 联盟广告收入
	UnionEventPayPurchaseAmount7d    float64 `json:"union_event_pay_purchase_amount7d"`     // 联盟广告收入
	UnionEventPayPurchaseAmount7dRoi float64 `json:"union_event_pay_purchase_amount7d_roi"` // 联盟变现ROI

	// 附加组件
	EventDspGiftForm int64 `json:"event_dsp_gift_form"` // 附加组件表单提交

	// 唤醒应用
	EventAppInvoked      int64   `json:"event_app_invoked"`       // 唤醒应用数
	EventAppInvokedCost  float64 `json:"event_app_invoked_cost"`  // 唤醒应用成本
	EventAppInvokedRatio float64 `json:"event_app_invoked_ratio"` // 唤醒应用率

	// 多转化 成本率
	EventMultiConversionRatio float64 `json:"event_multi_conversion_ratio"` // 落地页多转化率
	EventMultiConversionCost  float64 `json:"event_multi_conversion_cost"`  // 落地页多转化成本

	// 广告观看 成本率
	EventAdWatchTimesCost  float64 `json:"event_ad_watch_times_cost"`  // 广告观看次数成本
	EventAdWatchTimesRatio float64 `json:"event_ad_watch_times_ratio"` // 广告观看次数转化率

	// 智能电话
	EventGetThrough      int64   `json:"event_get_through"`       // 智能电话-确认接通数
	EventGetThroughCost  float64 `json:"event_get_through_cost"`  // 智能电话-确认接通成本
	EventGetThroughRatio float64 `json:"event_get_through_ratio"` // 智能电话-确认接通率

	// 广告观看次数
	EventAdWatch5Times       int64   `json:"event_ad_watch5_times"`        // 5次广告观看数
	EventAdWatch5TimesCost   float64 `json:"event_ad_watch5_times_cost"`   // 5次广告观看成本
	EventAdWatch5TimesRatio  float64 `json:"event_ad_watch5_times_ratio"`  // 5次广告观看转化率
	EventAdWatch10Times      int64   `json:"event_ad_watch10_times"`       // 10次广告观看数
	EventAdWatch10TimesCost  float64 `json:"event_ad_watch10_times_cost"`  // 10次广告观看成本
	EventAdWatch10TimesRatio float64 `json:"event_ad_watch10_times_ratio"` // 10次广告观看转化率
	EventAdWatch20Times      int64   `json:"event_ad_watch20_times"`       // 20次广告观看数
	EventAdWatch20TimesCost  float64 `json:"event_ad_watch20_times_cost"`  // 20次广告观看成本
	EventAdWatch20TimesRatio float64 `json:"event_ad_watch20_times_ratio"` // 20次广告观看转化率

	// 留资咨询
	EventConsultationValidRetained             int64   `json:"event_consultation_valid_retained"`               // 留资咨询数
	EventPreComponentConsultationValidRetained int64   `json:"event_pre_component_consultation_valid_retained"` // 附加咨询组件留资咨询数
	EventConsultationValidRetainedCost         float64 `json:"event_consultation_valid_retained_cost"`          // 留咨咨询成本
	EventConsultationValidRetainedRatio        float64 `json:"event_consultation_valid_retained_ratio"`         // 留咨咨询率

	// 转化数/深度转化数
	ConversionNum                     int64   `json:"conversion_num"`                        // 转化数(回传时间)
	DeepConversionNum                 int64   `json:"deep_conversion_num"`                   // 深度转化数(回传时间)
	ConversionNumByImpression7d       int64   `json:"conversion_num_by_impression7d"`        // 转化数(计费时间)
	DeepConversionNumByImpression7d   int64   `json:"deep_conversion_num_by_impression7d"`   // 深度转化数(计费时间)
	ConversionNumCost                 float64 `json:"conversion_num_cost"`                   // 转化成本(回传时间)
	ConversionRatio                   float64 `json:"conversion_ratio"`                      // 转化率(回传时间)
	DeepConversionCost                float64 `json:"deep_conversion_cost"`                  // 深度转化成本(回传时间)
	DeepConversionRatio               float64 `json:"deep_conversion_ratio"`                 // 深度转化率(回传时间)
	ConversionCostByImpression7d      float64 `json:"conversion_cost_by_impression7d"`       // 转化成本(计费时间)
	ConversionRatioByImpression7d     float64 `json:"conversion_ratio_by_impression7d"`      // 转化率(计费时间)
	DeepConversionCostByImpression7d  float64 `json:"deep_conversion_cost_by_impression7d"`  // 深度转化成本(计费时间)
	DeepConversionRatioByImpression7d float64 `json:"deep_conversion_ratio_by_impression7d"` // 深度转化率(计费时间)

	// 微信小程序
	EventWechatQrCodeLinkClick int64 `json:"event_wechat_qr_code_link_click"` // 微信小程序深度加粉数

	// 加权付费
	EventPayWeightedPurchaseAmount         float64 `json:"event_pay_weighted_purchase_amount"`           // 加权付费金额
	EventPayWeightedPurchaseAmountFirstDay float64 `json:"event_pay_weighted_purchase_amount_first_day"` // 首日加权付费金额
	EventPayWeightedRoi                    float64 `json:"event_pay_weighted_roi"`                       // 付费ROI
	EventPayFirstDayWeightedRoi            float64 `json:"event_pay_first_day_weighted_roi"`             // 首日付费ROI

	// 7日付费/多日付费金额
	EventPayWeekByConversion                      int64   `json:"event_pay_week_by_conversion"`                          // 7日付费次数
	EventPayWeekByConversionCost                  float64 `json:"event_pay_week_by_conversion_cost"`                     // 7日付费次数成本
	EventPayPurchaseAmountThreeDayByConversion    float64 `json:"event_pay_purchase_amount_three_day_by_conversion"`     // 激活后三日付费金额
	EventPayPurchaseAmountWeekByConversion        float64 `json:"event_pay_purchase_amount_week_by_conversion"`          // 激活后七日付费金额
	EventPayPurchaseAmountThreeDayByConversionRoi float64 `json:"event_pay_purchase_amount_three_day_by_conversion_roi"` // 激活后3日ROI
	EventPayPurchaseAmountWeekByConversionRoi     float64 `json:"event_pay_purchase_amount_week_by_conversion_roi"`      // 激活后7日ROI

	// 3日/7日留存
	EventThreeDayStayByConversion      int64   `json:"event_three_day_stay_by_conversion"`       // 3日留存数
	EventWeekStayByConversion          int64   `json:"event_week_stay_by_conversion"`            // 7日留存数
	EventThreeDayStayByConversionCost  float64 `json:"event_three_day_stay_by_conversion_cost"`  // 3日留存成本
	EventThreeDayStayByConversionRatio float64 `json:"event_three_day_stay_by_conversion_ratio"` // 3日留存率
	EventWeekStayByConversionCost      float64 `json:"event_week_stay_by_conversion_cost"`       // 7日留存成本
	EventWeekStayByConversionRatio     float64 `json:"event_week_stay_by_conversion_ratio"`      // 7日留存率

	// 外呼电话
	EventOutboundCall      int64   `json:"event_outbound_call"`       // 电话拨打数
	EventOutboundCallCost  float64 `json:"event_outbound_call_cost"`  // 电话拨打成本
	EventOutboundCallRatio float64 `json:"event_outbound_call_ratio"` // 电话拨打率

	// 关键行为
	KeyAction      int64   `json:"key_action"`       // 关键行为数
	KeyActionCost  float64 `json:"key_action_cost"`  // 关键行为成本
	KeyActionRatio float64 `json:"key_action_ratio"` // 关键行为率

	// 激活后24小时付费
	EventPayPurchaseAmountOneDayByConversion    float64 `json:"event_pay_purchase_amount_one_day_by_conversion"`     // 激活后24小时付费金额
	EventPayPurchaseAmountOneDayByConversionRoi float64 `json:"event_pay_purchase_amount_one_day_by_conversion_roi"` // 激活后24小时付费ROI
}

// ProgramPageReportResp 程序化落地页报表响应数据（仅data部分）
type ProgramPageReportResp struct {
	TotalCount int64                     `json:"total_count"` // 数据的总行数
	Details    []ProgramPageReportDetail `json:"details"`     // 数据明细信息
}
