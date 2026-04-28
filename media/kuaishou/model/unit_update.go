package model

import "errors"

// UnitUpdateReq 修改广告组请求
type UnitUpdateReq struct {
	accessTokenReq
	// 必填字段
	AdvertiserId int64  `json:"advertiser_id"` // 广告主ID，必填
	UnitId       int64  `json:"unit_id"`       // 广告组ID，必填
	UnitName     string `json:"unit_name"`     // 广告组名称，必填，1-100字符，同计划内唯一
	BidType      int    `json:"bid_type"`      // 出价类型，必填：2=CPC 10=OCPM 12=MCB最大转化量

	// 计划与时间
	CampaignId int64  `json:"campaign_id,omitempty"` // 广告计划ID
	BeginTime  string `json:"begin_time,omitempty"`  // 投放开始时间，格式：yyyy-MM-dd，须>=当前日期；周期稳投不支持修改
	EndTime    string `json:"end_time,omitempty"`    // 投放结束时间，格式：yyyy-MM-dd，传空字符串=长期投放

	// 预算
	DayBudget         int64    `json:"day_budget,omitempty"`          // 日预算，单位：分；0=不限；最低100元，最高1亿元；不能与day_budget_schedule同用；周期稳投不支持修改
	DayBudgetSchedule []string `json:"day_budget_schedule,omitempty"` // 分日预算，单位：分；不能与day_budget同用；优先级高于day_budget

	// 出价
	Bid                int64   `json:"bid,omitempty"`                  // 出价，bid_type=CPC时必填，单位：分；0.2-100元
	CpaBid             int64   `json:"cpa_bid,omitempty"`              // 出价，bid_type=OCPM时必填，单位：分
	DeepConversionBid  int64   `json:"deep_conversion_bid,omitempty"`  // 深度转化目标出价，单位：分；须>cpa_bid；不可从0变非0或反向
	OcpxActionType     int     `json:"ocpx_action_type,omitempty"`     // 优化目标，bid_type=OCPM时必填
	DeepConversionType int64   `json:"deep_conversion_type,omitempty"` // 深度转化目标：3=付费 7=次留 10=完成 11=授信 13=加购 14=提单 15=购买 44=有效线索 92=付费ROI 181=24H次留
	RoiRatio           float64 `json:"roi_ratio,omitempty"`            // 付费ROI系数，优化目标为ROI时必填，范围(0,100]，最多3位小数

	// 版位
	SceneId []string `json:"scene_id,omitempty"` // 资源位，有创意后不可修改；联盟不可修改

	// 定向
	Target *UnitCreateTarget `json:"target,omitempty"` // 定向数据

	// 应用相关
	AppId           int64    `json:"app_id,omitempty"`            // 应用ID，campaign_type=2/7/9时必填
	AppDownloadType int      `json:"app_download_type,omitempty"` // 应用下载方式：0=直接下载 1=落地页下载
	AppIconUrl      string   `json:"app_icon_url,omitempty"`      // 应用图标URL
	AppStore        []string `json:"app_store,omitempty"`         // 应用商店：huawei/oppo/vivo/xiaomi/meizu/smartisan/honor
	UseAppMarket    int      `json:"use_app_market,omitempty"`    // 优先系统应用商店：0=否 1=是；仅Android应用
	PackageId       int64    `json:"package_id,omitempty"`        // 应用包ID

	// URL与落地页
	Url             string `json:"url,omitempty"`               // 推广链接，最多1000字符
	UrlType         int    `json:"url_type,omitempty"`          // 链接类型，campaign_type=3时必填
	WebUriType      int    `json:"web_uri_type,omitempty"`      // 落地页类型：1=自有链接 2=建站 3=程序化落地页 4=微信小程序
	SiteId          int64  `json:"site_id,omitempty"`           // 建站ID，web_uri_type=2时必填
	SiteType        int    `json:"site_type,omitempty"`         // 预约广告类型：1=iOS预约 2=Android预约
	GroupId         int64  `json:"group_id,omitempty"`          // 程序化落地页ID，web_uri_type=3时必填
	SchemaId        string `json:"schema_id,omitempty"`         // 微信小程序ID，web_uri_type=4时使用
	SchemaUri       string `json:"schema_uri,omitempty"`        // 唤起链接，应用推广类计划必填
	DownloadPageUrl string `json:"download_page_url,omitempty"` // 自定义落地页URL，site_id优先级更高
	ULink           string `json:"u_link,omitempty"`            // iOS ulink，campaign_type=7/35，最多2000字符

	// 小程序
	CustomMiniAppData *UnitCreateCustomMiniApp `json:"custom_mini_app_data,omitempty"` // 小程序信息，campaign_type=19时必填

	// DPA商品广告
	DpaUnitParam *UnitCreateDpaUnit `json:"dpa_unit_param,omitempty"` // DPA商品信息，商品推广类计划必填

	// 搜索广告
	NegativeWordParam           *UnitCreateNegativeWord `json:"negative_word_param,omitempty"`           // 搜索广告否定词
	ExtendSearch                bool                    `json:"extend_search,omitempty"`                 // 智能词扩展
	QuickSearch                 int                     `json:"quick_search,omitempty"`                  // 搜索极速启动：0=关闭 1=开启
	TargetExplore               int                     `json:"target_explore,omitempty"`                // 搜索受众探索：0=关闭 1=开启；需quick_search=1
	SearchPopulationRetargeting int                     `json:"search_population_retargeting,omitempty"` // 搜索人群重定向：0=关闭 1=开启

	// 版式与样式
	AdType          int `json:"ad_type,omitempty"`           // 广告类型
	DetailUnitType  int `json:"detail_unit_type,omitempty"`  // 详情单元类型
	CardType        int `json:"card_type,omitempty"`         // 卡片类型
	OuterLoopNative int `json:"outer_loop_native,omitempty"` // 开启原生展示：0=关闭 1=开启
	ShowMode        int `json:"show_mode,omitempty"`         // 创意展示方式：1=轮播 2=精选（默认）
	UnitType        int `json:"unit_type,omitempty"`         // 广告组类型

	// 直播
	JingleBellId      int64 `json:"jingle_bell_id,omitempty"`      // 铃铛组件ID，campaign_type=16时必填
	LiveUserId        int64 `json:"live_user_id,omitempty"`        // 主播ID，campaign_type=16时必填
	LiveComponentType int   `json:"live_component_type,omitempty"` // 直播组件类型：0=铃铛 1=房产 2=团购 3=服务键 4=无组件 5=小程序 6=手柄 14=KS招聘
	ConversionType    int   `json:"conversion_type,omitempty"`     // 转化路径，campaign_type=16时必填，值=6

	// 高级创意
	AdvCardList     []int64 `json:"adv_card_list,omitempty"`     // 绑定高级创意卡片ID列表
	AdvCardOption   int     `json:"adv_card_option,omitempty"`   // 高级创意开关：0=关闭 1=开启
	AssetMining     bool    `json:"asset_mining,omitempty"`      // 程序化2.0素材挖掘，unit_type=7时可选
	AutoCreatePhoto bool    `json:"auto_create_photo,omitempty"` // 自动生成图片
	SmartCover      bool    `json:"smart_cover,omitempty"`       // 程序化2.0智能选帧，unit_type=7时可选
	SmartBid        int     `json:"smart_bid,omitempty"`         // 智能出价

	// 试玩广告
	PlayableSwitch      int    `json:"playable_switch,omitempty"`      // 试玩开关
	PlayableId          int64  `json:"playable_id,omitempty"`          // 试玩ID
	PlayableUrl         string `json:"playable_url,omitempty"`         // 试玩URL
	PlayableOrientation int    `json:"playable_orientation,omitempty"` // 试玩方向
	PlayButton          string `json:"play_button,omitempty"`          // 试玩按钮文案：1=立即试玩 2=试玩一下 3=立即体验 4=免装试玩 5=免装体验

	// 短剧
	SeriesPayTemplateId      int64   `json:"series_pay_template_id,omitempty"`       // 付费模板ID，campaign_type=30时使用
	SeriesPayMode            int     `json:"series_pay_mode,omitempty"`              // 付费模式：1=套餐 2=虚拟货币
	SeriesPayTemplateIdMulti []int64 `json:"series_pay_template_id_multi,omitempty"` // 短剧付费模板列表，最多5个

	// 微信小程序/电商
	UnitMaterialType int `json:"unit_material_type,omitempty"` // 广告素材类型，campaign_type=32/35时必填

	// 其他
	ImMessageMount        bool   `json:"im_message_mount,omitempty"`        // 挂载私信留资组件
	ComponentId           int64  `json:"component_id,omitempty"`            // 组件ID
	ConsultId             int64  `json:"consult_id,omitempty"`              // 咨询组件ID，仅线索收集计划使用建站/程序化落地页时可用
	ConvertId             int64  `json:"convert_id,omitempty"`              // 转化追踪工具ID
	EnhanceConversionType int    `json:"enhance_conversion_type,omitempty"` // 增强目标：8=7日留存
	PutStatus             int    `json:"put_status,omitempty"`              // 广告组状态：1=投放 2=暂停；周期稳投不支持修改
	ScheduleTime          string `json:"schedule_time,omitempty"`           // 投放时段，24*7格式，0/1字符串
	NightScheduledTag     int    `json:"night_scheduled_tag,omitempty"`     // 夜间投放时段类型：0=非夜间 1=全夜间 2=指定夜间
	SplashAdSwitch        bool   `json:"splash_ad_switch,omitempty"`        // 开屏广告开关
	Speed                 int    `json:"speed,omitempty"`                   // 投放速度
	VideoLandingPage      bool   `json:"video_landing_page,omitempty"`      // 视频落地页
}

func (receiver *UnitUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *UnitUpdateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.UnitId <= 0 {
		err = errors.New("unit_id is empty")
		return
	}
	if len(receiver.UnitName) == 0 {
		err = errors.New("unit_name is empty")
		return
	}
	if receiver.BidType <= 0 {
		err = errors.New("bid_type is empty")
		return
	}
	return
}

// UnitUpdateResp 修改广告组响应数据（仅data部分）
type UnitUpdateResp struct {
	UnitId int64 `json:"unit_id"` // 广告组ID
}
