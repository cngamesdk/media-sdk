package model

import "errors"

// UnitListReq 查询广告组请求
type UnitListReq struct {
	accessTokenReq
	AdvertiserId           int64    `json:"advertiser_id"`                       // 广告主ID，必填
	AppName                string   `json:"app_name,omitempty"`                  // 应用名称
	CampaignId             int64    `json:"campaign_id,omitempty"`               // 广告计划ID，筛选条件，空=不限
	CampaignType           int64    `json:"campaign_type,omitempty"`             // 计划类型
	DeepConversionTypeList []string `json:"deep_conversion_type_list,omitempty"` // 深度转化类型列表
	OcpxActionTypeList     []string `json:"ocpx_action_type_list,omitempty"`     // OCPX优化目标列表
	PutStatusList          []string `json:"put_status_list,omitempty"`           // 投放状态筛选：1=投放 2=暂停 3=删除
	ReviewStatusList       []string `json:"review_status_list,omitempty"`        // 审核状态筛选：1=审核中 2=审核通过 3=审核不通过 7=待提审
	StartDate              string   `json:"start_date,omitempty"`                // 开始日期，格式：yyyy-MM-dd，需与end_date同时传
	EndDate                string   `json:"end_date,omitempty"`                  // 结束日期，格式：yyyy-MM-dd，需与start_date同时传
	Status                 int      `json:"status,omitempty"`                    // 广告组状态：-2=不限(含删除) 10=广告组删除 40=广告创意删除
	TimeFilterType         int      `json:"time_filter_type,omitempty"`          // 时间过滤类型：0/不传=更新时间 1=创建时间
	UnitId                 int64    `json:"unit_id,omitempty"`                   // 广告组ID，筛选条件，空=不限
	UnitIds                []string `json:"unit_ids,omitempty"`                  // 广告组ID集合，最多100个
	UnitName               string   `json:"unit_name,omitempty"`                 // 广告组名称，精确匹配，不支持模糊搜索
	Page                   int      `json:"page,omitempty"`                      // 页码，默认1
	PageSize               int      `json:"page_size,omitempty"`                 // 每页数量，默认20
}

func (receiver *UnitListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *UnitListReq) Validate() (err error) {
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

// UnitListAdvCard 高级创意卡片
type UnitListAdvCard struct {
	AdvCardId       int64  `json:"adv_card_id"`       // 卡片ID
	CardType        int    `json:"card_type"`         // 卡片类型
	Price           int    `json:"price"`             // 价格
	SalePrice       int    `json:"sale_price"`        // 促销价格
	SdpaCardContent string `json:"sdpa_card_content"` // SDPA卡片内容
	Subtitle        string `json:"subtitle"`          // 副标题
	Title           string `json:"title"`             // 标题
	Url             string `json:"url"`               // 链接
}

// UnitListBackflowForecast 回流预估信息
type UnitListBackflowForecast struct {
	BackflowCvLower   int     `json:"backflow_cv_lower"`  // 回流估算下限
	BackflowCvUpper   int     `json:"backflow_cv_upper"`  // 回流估算上限
	BackflowPayment   float64 `json:"backflow_payment"`   // 回流估算总金额
	BackflowRoi       float64 `json:"backflow_roi"`       // 首日回流ROI估算
	BackflowTimestamp int64   `json:"backflow_timestamp"` // 数据时间戳，13位毫秒时间戳
}

// UnitListDiverseData 应用信息
type UnitListDiverseData struct {
	AppName        string `json:"app_name"`         // 应用名称
	AppPackageName string `json:"app_package_name"` // 应用包名
	DeviceOsType   int    `json:"device_os_type"`   // 应用系统类型：0=未知 1=Android 2=iOS
}

// UnitListGroupPageReview 程序化落地页审核信息
type UnitListGroupPageReview struct {
	PageId       int64  `json:"page_id"`       // 页面ID
	ReviewDetail string `json:"review_detail"` // 审核拒绝原因
	ReviewStatus int    `json:"review_status"` // 审核状态：1=审核中 2=审核通过 3=审核不通过
	Url          string `json:"url"`           // 页面URL
}

// UnitListGroupPage 程序化落地页组信息
type UnitListGroupPage struct {
	GroupId          string                    `json:"group_id"`           // 落地页组ID
	GroupName        string                    `json:"group_name"`         // 落地页组名称
	PageReviewDetail []UnitListGroupPageReview `json:"page_review_detail"` // 组内页面列表
}

// UnitListSchedule 投放时段（历史字段，即将废弃）
type UnitListSchedule struct {
	Mon  []string `json:"mon"`  // 周一投放小时，范围0-23
	Tues []string `json:"tues"` // 周二投放小时
	Wed  []string `json:"wed"`  // 周三投放小时
	Thur []string `json:"thur"` // 周四投放小时
	Fri  []string `json:"fri"`  // 周五投放小时
	Sat  []string `json:"sat"`  // 周六投放小时
	Sun  []string `json:"sun"`  // 周日投放小时
}

// UnitListSeriesCardInfo 剧集卡片信息
type UnitListSeriesCardInfo struct {
	PicId       int64    `json:"pic_id"`      // 封面图片ID
	Label       []string `json:"label"`       // 剧集卡标签
	CoverImage  string   `json:"cover_image"` // 封面图片URL
	Title       string   `json:"title"`       // 卡片标题
	Description string   `json:"description"` // 卡片描述
}

// UnitDetail 广告组详情
type UnitDetail struct {
	UnitId                      int64                     `json:"unit_id"`                       // 广告组ID
	UnitName                    string                    `json:"unit_name"`                     // 广告组名称
	UnitType                    int                       `json:"unit_type"`                     // 创意制作方式：3=DPA自定义 4=常规自定义 7=程序化2.0 10=智能创意
	UnitSource                  int                       `json:"unit_source"`                   // 广告组来源：0=常规 1=托管
	CampaignId                  int64                     `json:"campaign_id"`                   // 广告计划ID
	AdType                      int                       `json:"ad_type"`                       // 计划类型：0=信息流 1=搜索
	Status                      int                       `json:"status"`                        // 广告组状态：-1=不限 1=计划暂停 3=计划超预算 6=余额不足 10=广告组删除 11=审核中 12=审核未通过 14=已结束 15=暂停 17=单元超预算 19=未到投放时间 20=有效 22=未到投放期
	PutStatus                   int64                     `json:"put_status"`                    // 投放状态：1=投放 2=暂停 3=删除
	ReviewDetail                string                    `json:"review_detail"`                 // 审核拒绝原因
	BidType                     int64                     `json:"bid_type"`                      // 出价类型：1=CPM 2=CPC 6=OCPC 10=OCPM 20=eCPC
	Bid                         int64                     `json:"bid"`                           // 出价，单位：分
	CpaBid                      int64                     `json:"cpa_bid"`                       // OCPC出价，单位：分
	DeepConversionBid           int64                     `json:"deep_conversion_bid"`           // 深度转化目标出价，单位：分
	OcpxActionType              int64                     `json:"ocpx_action_type"`              // 优化目标
	DeepConversionType          int64                     `json:"deep_conversion_type"`          // 深度转化目标：3=付费 7=次留 10=完成 11=授信 13=加购 14=提单 15=购买 44=有效线索 92=付费ROI 181=24H次留 0=无
	RoiRatio                    float64                   `json:"roi_ratio"`                     // 付费ROI系数，范围(0,100]
	DayBudget                   int64                     `json:"day_budget"`                    // 日预算，单位：分
	DayBudgetSchedule           []string                  `json:"day_budget_schedule"`           // 分日预算，单位：分；优先级高于day_budget
	BeginTime                   string                    `json:"begin_time"`                    // 投放开始时间，格式：yyyy-MM-dd
	EndTime                     string                    `json:"end_time"`                      // 投放结束时间，格式：yyyy-MM-dd；null=长期投放
	CreateTime                  string                    `json:"create_time"`                   // 创建时间
	UpdateTime                  string                    `json:"update_time"`                   // 最后修改时间，格式：yyyy-MM-dd HH:mm:ss
	SceneId                     []string                  `json:"scene_id"`                      // 资源位：1=精选 2=场景精选-信息流 6=竖版信息流 7=双列 24=激励视频 11=快看场景
	ShowMode                    int                       `json:"show_mode"`                     // 创意展示方式：0=未知 1=轮播 2=精选
	ScheduleTime                string                    `json:"schedule_time"`                 // 投放时段，24*7字符串
	Schedule                    *UnitListSchedule         `json:"schedule"`                      // 投放时段（历史字段，即将废弃）
	NightScheduledTag           int                       `json:"night_scheduled_tag"`           // 夜间投放类型：0=非夜间 1=全夜间 2=指定夜间
	Target                      *UnitCreateTarget         `json:"target"`                        // 定向数据
	AppId                       int64                     `json:"app_id"`                        // 应用ID
	AppDownloadType             int                       `json:"app_download_type"`             // 应用下载方式：0=直接下载 1=落地页下载
	AppIconUrl                  string                    `json:"app_icon_url"`                  // 应用图标URL
	AppStore                    []string                  `json:"app_store"`                     // 应用商店列表
	UseAppMarket                int                       `json:"use_app_market"`                // 优先系统应用商店：0=否 1=是
	PackageId                   int64                     `json:"package_id"`                    // 应用包ID
	Url                         string                    `json:"url"`                           // 推广链接
	UrlType                     int                       `json:"url_type"`                      // 链接类型
	WebUriType                  int                       `json:"web_uri_type"`                  // 落地页类型：默认1 2=建站落地页
	SiteId                      int64                     `json:"site_id"`                       // 建站ID
	SiteType                    int                       `json:"site_type"`                     // 预约广告类型：1=iOS预约
	GroupId                     string                    `json:"group_id"`                      // 程序化落地页ID
	SchemaId                    string                    `json:"schema_id"`                     // 微信小程序外链
	SchemaUri                   string                    `json:"schema_uri"`                    // 唤起链接
	ULink                       string                    `json:"u_link"`                        // iOS ulink
	CustomMiniAppData           *UnitCreateCustomMiniApp  `json:"custom_mini_app_data"`          // 小程序参数
	NegativeWordParam           *UnitCreateNegativeWord   `json:"negative_word_param"`           // 搜索广告否定词
	ExtendSearch                bool                      `json:"extend_search"`                 // 智能词扩展
	QuickSearch                 int                       `json:"quick_search"`                  // 搜索极速启动：0=关闭 1=开启
	TargetExplore               int                       `json:"target_explore"`                // 搜索受众探索：0=关闭 1=开启
	SearchPopulationRetargeting int                       `json:"search_population_retargeting"` // 搜索人群重定向：0=关闭 1=开启
	PageGroupDetail             *UnitListGroupPage        `json:"page_group_detail"`             // 程序化落地页组信息
	AdvCardList                 []UnitListAdvCard         `json:"adv_card_list"`                 // 高级创意列表
	AdvCardOption               int                       `json:"adv_card_option"`               // 高级创意开关：0=关闭 1=开启
	AssetMining                 bool                      `json:"asset_mining"`                  // 程序化2.0素材挖掘
	SmartCover                  bool                      `json:"smart_cover"`                   // 程序化2.0智能选帧
	SmartBid                    int                       `json:"smart_bid"`                     // 智能出价
	AutoCreatePhoto             bool                      `json:"auto_create_photo"`             // 自动生成图片
	PlayableSwitch              int                       `json:"playable_switch"`               // 试玩开关
	PlayableId                  int64                     `json:"playable_id"`                   // 试玩ID
	PlayableUrl                 string                    `json:"playable_url"`                  // 试玩URL
	PlayableOrientation         int                       `json:"playable_orientation"`          // 试玩方向：-1=默认 0=竖横 1=竖版 2=横版
	PlayableFileName            string                    `json:"playable_file_name"`            // 试玩文件名
	PlayButton                  string                    `json:"play_button"`                   // 试玩按钮文案
	DspVersion                  int                       `json:"dsp_version"`                   // DSP版本
	StudyStatus                 int                       `json:"study_status"`                  // 学习期状态：1=学习中 2=学习成功 3=学习失败
	CompensateStatus            int                       `json:"compensate_status"`             // 保价状态：0=无需保价 1=保价中 2=保价待确认 3=保价完成 4=已过期
	DpaUnitSubType              int                       `json:"dpa_unit_sub_type"`             // 商品广告类型：1=DPA 2=SDPA 3=动态商品卡
	DpaCategories               []string                  `json:"dpa_categories"`                // DPA类目
	DpaOuterIds                 []string                  `json:"dpa_outer_ids"`                 // DPA外部商品ID列表
	DpaDynamicParams            int                       `json:"dpa_dynamic_params"`            // 动态参数：0=关闭 -1=自动参数
	DpaDynamicParamsForDp       string                    `json:"dpa_dynamic_params_for_dp"`     // 应用直链动态参数，最多100字符
	DpaDynamicParamsForUri      string                    `json:"dpa_dynamic_params_for_uri"`    // 落地页链接动态参数，最多100字符
	LibraryId                   int64                     `json:"library_id"`                    // 商品库ID
	OuterId                     string                    `json:"outer_id"`                      // 第三方商品ID
	ProductId                   string                    `json:"product_id"`                    // 商品ID
	ProductImage                string                    `json:"product_image"`                 // 商品主图
	ProductName                 string                    `json:"product_name"`                  // 商品名称
	ProductPrice                float64                   `json:"product_price"`                 // 商品价格，单位：元
	JingleBellId                int64                     `json:"jingle_bell_id"`                // 铃铛组件ID
	LiveUserId                  int64                     `json:"live_user_id"`                  // 主播ID/短剧作者ID
	LiveComponentType           int                       `json:"live_component_type"`           // 直播组件类型：0=铃铛 1=房产 2=团购 3=服务键 4=无组件 5=小程序 6=手柄 14=KS招聘
	KwaiBookId                  int64                     `json:"kwai_book_id"`                  // 快手小说ID
	SeriesId                    int64                     `json:"series_id"`                     // 短剧ID
	EpisodeId                   int64                     `json:"episode_id"`                    // 剧集ID
	SeriesCardType              int                       `json:"series_card_type"`              // 剧集卡开关：0=关闭 1=开启
	SeriesCardInfo              *UnitListSeriesCardInfo   `json:"series_card_info"`              // 剧集卡信息
	SeriesPayTemplateId         int64                     `json:"series_pay_template_id"`        // 付费模板ID
	SeriesPayMode               int                       `json:"series_pay_mode"`               // 付费模式：1=套餐 2=虚拟货币
	SeriesPayTemplateIdMulti    []int64                   `json:"series_pay_template_id_multi"`  // 短剧付费模板列表
	UnitMaterialType            int                       `json:"unit_material_type"`            // 素材类型
	OuterLoopNative             int                       `json:"outer_loop_native"`             // 原生展示：0=关闭 1=开启
	ImMessageMount              bool                      `json:"im_message_mount"`              // 挂载私信留资组件
	ComponentId                 int64                     `json:"component_id"`                  // 组件ID
	ConsultId                   int64                     `json:"consult_id"`                    // 咨询组件：0=未使用 1=使用
	ConvertId                   int64                     `json:"convert_id"`                    // 转化追踪工具ID
	EnhanceConversionType       int                       `json:"enhance_conversion_type"`       // 增强转化目标
	PageAuditStatus             int                       `json:"page_audit_status"`             // 落地页审核状态
	SplashAdSwitch              bool                      `json:"splash_ad_switch"`              // 开屏广告开关
	Speed                       int                       `json:"speed"`                         // 投放速度
	VideoLandingPage            bool                      `json:"video_landing_page"`            // 视频落地页
	LinkIntegrationType         int                       `json:"link_integration_type"`         // 投放链路优化
	CreateChannel               int64                     `json:"create_channel"`                // 创建渠道
	DiverseData                 *UnitListDiverseData      `json:"diverse_data"`                  // 应用信息
	BackflowForecast            *UnitListBackflowForecast `json:"backflow_forecast"`             // 回流预估信息
	SupportUnitIds              []string                  `json:"support_unit_ids"`              // 支持的广告组ID列表
}

// UnitListResp 查询广告组响应数据（仅data部分）
type UnitListResp struct {
	TotalCount int64        `json:"total_count"` // 广告组总数量
	Details    []UnitDetail `json:"details"`     // 广告组详情列表
}
