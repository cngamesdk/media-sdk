package model

import "errors"

// UnitCreateAgeRange 自定义年龄范围
type UnitCreateAgeRange struct {
	Min int `json:"min"` // 最小年龄，最小18，18<=min<=75
	Max int `json:"max"` // 最大年龄，最大100，max-min>=5
}

// UnitCreateBehaviorKeyword 行为/兴趣关键词
type UnitCreateBehaviorKeyword struct {
	Id   int64  `json:"id"`   // 关键词ID，必填，id与name需匹配
	Name string `json:"name"` // 关键词名称，必填，id与name需匹配
}

// UnitCreateBehavior 行为定向
type UnitCreateBehavior struct {
	Keyword      []UnitCreateBehaviorKeyword `json:"keyword,omitempty"`       // 行为关键词
	Label        []string                    `json:"label,omitempty"`         // 行为类目词，格式："8-802-80202"
	SceneType    []string                    `json:"scene_type,omitempty"`    // 行为场景：1=社区 2=APP 3=电商 4=推广
	StrengthType int                         `json:"strength_type,omitempty"` // 行为强度：0=不限 1=高强度
	TimeType     int                         `json:"time_type,omitempty"`     // 行为天数：0=7天 1=15天 2=30天 3=90天 4=180天
}

// UnitCreateInterest 兴趣定向
type UnitCreateInterest struct {
	Keyword      []UnitCreateBehaviorKeyword `json:"keyword,omitempty"`       // 兴趣关键词
	Label        []string                    `json:"label,omitempty"`         // 兴趣类目词，格式："8-802-80202"
	StrengthType int                         `json:"strength_type,omitempty"` // 兴趣强度：0=不限 1=高强度
}

// UnitCreateBehaviorInterest 行为兴趣定向
type UnitCreateBehaviorInterest struct {
	Behavior *UnitCreateBehavior `json:"behavior,omitempty"` // 行为定向
	Interest *UnitCreateInterest `json:"interest,omitempty"` // 兴趣定向
}

// UnitCreateCelebrityFansStar 达人粉丝内容
type UnitCreateCelebrityFansStar struct {
	Category []string `json:"category,omitempty"` // 达人分类，格式："first_label_id,second_label_id"
	Id       string   `json:"id,omitempty"`       // 达人ID；type=1时格式"33-177"；type=2时为author_id
	Name     string   `json:"name,omitempty"`     // 达人名称
	Type     int      `json:"type,omitempty"`     // 达人类型：1=达人分类 2=快手达人
}

// UnitCreateCelebrity 快手达人定向
type UnitCreateCelebrity struct {
	Behaviors []string                      `json:"behaviors,omitempty"`  // 行为类型：0=关注 1=视频互动 2=直播互动
	FansStars []UnitCreateCelebrityFansStar `json:"fans_stars,omitempty"` // 达人内容列表
}

// UnitCreateBehaviorInterestParamShow 行为兴趣4.0
type UnitCreateBehaviorInterestParamShow struct {
	CategoryIds    []string `json:"category_ids,omitempty"`    // 分类ID列表，格式："8-802-80202"
	KeywordIds     []int    `json:"keyword_ids,omitempty"`     // 关键词ID列表
	CustomBehavior int      `json:"custom_behavior,omitempty"` // 自定义行为意向：0=关闭 1=开启
	SceneTypes     []int    `json:"scene_types,omitempty"`     // 场景类型：1=视频 2=APP 4=广告；custom_behavior=1时必填
	TimeType       int      `json:"time_type,omitempty"`       // 时间范围：0=7天 1=15天 2=30天 3=90天 4=180天 5=60天；custom_behavior=1时必填
}

// UnitCreateDistanceShow 新商圈定向
type UnitCreateDistanceShow struct {
	Address      string `json:"address,omitempty"`       // 地址
	Lng          string `json:"lng,omitempty"`           // 经度
	Lat          string `json:"lat,omitempty"`           // 纬度
	Radius       int64  `json:"radius,omitempty"`        // 半径
	LocationName string `json:"location_name,omitempty"` // 地点名称
	PoiId        string `json:"poi_id,omitempty"`        // POI ID
}

// UnitCreateTarget 定向数据
type UnitCreateTarget struct {
	// 设备系统
	AndroidOsv int `json:"android_osv,omitempty"` // Android版本：3=不限 4=4.x+ 5=5.x+ ... 10=10.x+
	IosOsv     int `json:"ios_osv,omitempty"`     // iOS版本：6=不限 7=7.x+ ... 16=16.x+
	HarmonyOsv int `json:"harmony_osv,omitempty"` // 鸿蒙版本：1=不限 4=1.x-4.x 5=5.x+
	PlatformOs int `json:"platform_os,omitempty"` // 操作系统：0=不限 1=Android 2=iOS 3=Android&iOS 4=鸿蒙

	// 设备品牌与价格
	DeviceBrand    []string `json:"device_brand,omitempty"`     // 设备品牌
	DeviceBrandIds []string `json:"device_brand_ids,omitempty"` // 设备品牌ID列表，[]为不限
	DevicePrice    []string `json:"device_price,omitempty"`     // 设备价格，[]为不限

	// 网络与地域
	Network     int      `json:"network,omitempty"`      // 网络：1=Wi-Fi 2=移动网络 0=不限
	Operators   []int    `json:"operators,omitempty"`    // 运营商：1=移动 2=电信 3=联通；[]为不限
	Region      []string `json:"region,omitempty"`       // 地域，[]为不限；传父级ID不传子级ID
	DistrictIds []string `json:"district_ids,omitempty"` // 商圈定向，不能与region同用；最多100个
	UserType    int      `json:"user_type,omitempty"`    // 用户类型：0=实时位置 1=居住地 2=全部 4=游客 5=家乡
	IpType      int      `json:"ip_type,omitempty"`      // 地域IP：0=默认IP 1=广告协会IP

	// 人口属性
	Gender      int                 `json:"gender,omitempty"`        // 性别：1=女 2=男 0=不限
	AgeV2       *UnitCreateAgeRange `json:"age_v2,omitempty"`        // 自定义年龄范围，不能与ages_range_v2同用
	AgesRangeV2 []string            `json:"ages_range_v2,omitempty"` // 固定年龄段，不能与age_v2同用

	// 行为兴趣
	BehaviorInterest          *UnitCreateBehaviorInterest          `json:"behavior_interest,omitempty"`            // 行为兴趣定向
	BehaviorType              int                                  `json:"behavior_type,omitempty"`                // 行为兴趣类型：0=不限 1=自定义 2=系统推荐
	Celebrity                 *UnitCreateCelebrity                 `json:"celebrity,omitempty"`                    // 快手达人定向，仅快手信息流
	BehaviorInterestParamShow *UnitCreateBehaviorInterestParamShow `json:"behavior_interest_param_show,omitempty"` // 行为兴趣4.0，需白名单

	// 人群过滤
	DisableInstalledAppSwitch int      `json:"disable_installed_app_switch,omitempty"` // 过滤已安装APP：0=过滤（默认）1=不限
	ExcludePopulation         []string `json:"exclude_population,omitempty"`           // 排除人群ID列表
	FilterConvertedLevel      int      `json:"filter_converted_level,omitempty"`       // 过滤已转化用户：0=不限 1=单元 2=计划 3=账户 4=公司 5=APP 6=自定义商品 7=企业微信
	FilterTimeRange           int      `json:"filter_time_range,omitempty"`            // 用户转化时间范围：0=30天 1=60天 2=90天
	FilterConvertedWechatId   []string `json:"filter_converted_wechat_id,omitempty"`   // 过滤已转化企业微信ID
	PaidAudience              []string `json:"paid_audience,omitempty"`                // 付费人群ID，不能与population/exclude_population同用
	Population                []string `json:"population,omitempty"`                   // 定向人群ID，不能与paid_audience同用

	// 智能定向
	IntelliExtendOption int      `json:"intelli_extend_option,omitempty"` // 智能定向开关：0=关闭 1=开启 2=Pro开启；与auto_population互斥
	AutoPopulation      int      `json:"auto_population,omitempty"`       // 智能人群：0=默认 1=开启；与intelli_extend_option互斥
	SeedPopulation      []string `json:"seed_population,omitempty"`       // 种子人群，需开启种子人群功能

	// APP与媒体定向
	AppIds          []string `json:"app_ids,omitempty"`           // APP行为-按APP名称，仅Android，不能与app_interest_ids同用
	AppInterestIds  []string `json:"app_interest_ids,omitempty"`  // APP行为-按分类，仅Android，不能与app_ids同用
	AppNames        []string `json:"app_names,omitempty"`         // APP名称
	Media           []string `json:"media,omitempty"`             // 媒体定向包，仅联盟
	ExcludeMedia    []string `json:"exclude_media,omitempty"`     // 媒体排除包，不能与media同用；仅联盟
	MediaSourceType int      `json:"media_source_type,omitempty"` // 媒体包来源：0=不限 1=行业优质流量包 2=广告主自定义包

	// 其他定向
	DistanceShow                []UnitCreateDistanceShow `json:"distance_show,omitempty"`                 // 新商圈定向
	TargetSource                int                      `json:"target_source,omitempty"`                 // 定向来源
	SharedUser                  int                      `json:"shared_user,omitempty"`                   // 过滤共享手机用户：0=默认 1=过滤
	TemplateId                  int64                    `json:"template_id,omitempty"`                   // 定向模板ID，关联后target字段失效
	SearchPopulationRetargeting int                      `json:"search_population_retargeting,omitempty"` // 搜索人群重定向：0=关闭 1=开启
}

// UnitCreateCustomMiniApp 小程序信息
type UnitCreateCustomMiniApp struct {
	MiniAppIdPlatform string `json:"mini_app_id_platform"`       // 小程序APPID，必填，最多30字符
	MiniAppType       int    `json:"mini_app_type,omitempty"`    // 小程序类型：1=小程序（默认）2=小游戏
	BootstrapPage     string `json:"bootstrap_page"`             // 小程序启动页，mini_app_type=1时必填
	BootstrapParams   string `json:"bootstrap_params,omitempty"` // 小程序启动参数
}

// UnitCreateDpaUnit DPA商品信息
type UnitCreateDpaUnit struct {
	DpaUnitSubType           int      `json:"dpa_unit_sub_type,omitempty"`            // 商品广告类型：1=DPA 2=SDPA 3=动态商品卡
	LibraryId                int64    `json:"library_id,omitempty"`                   // 商品库ID
	DpaCategoryIds           []string `json:"dpa_category_ids,omitempty"`             // DPA类目ID集合，格式："level1-level2-level3"
	DpaOuterIds              []string `json:"dpa_outer_ids,omitempty"`                // DPA外部商品ID集合，优先级高于dpa_category_ids
	DpaDynamicParams         int      `json:"dpa_dynamic_params,omitempty"`           // 启用动态参数：0=关闭（默认）1=开启
	DpaDynamicParamsForDp    string   `json:"dpa_dynamic_params_for_dp,omitempty"`    // DPA应用直链动态参数，最多100字符
	DpaDynamicParamsForUri   string   `json:"dpa_dynamic_params_for_uri,omitempty"`   // DPA落地页链接动态参数，最多100字符
	OuterId                  string   `json:"outer_id,omitempty"`                     // 外部商品ID，SDPA必填
	ProductId                string   `json:"product_id,omitempty"`                   // 快手商品ID，SDPA必填
	DpaUnitClickUrl          string   `json:"dpa_unit_click_url,omitempty"`           // DPA单元点击URL
	DpaUnitActionbarClickUrl string   `json:"dpa_unit_actionbar_click_url,omitempty"` // DPA单元行动号召点击URL
	UnitImpressionUrl        string   `json:"unit_impression_url,omitempty"`          // 单元曝光URL
}

// UnitCreateNegativeWord 搜索广告否定词
type UnitCreateNegativeWord struct {
	ExactWords  []string `json:"exact_words,omitempty"`  // 精确否定词，最多200个，单词最多20字符
	PhraseWords []string `json:"phrase_words,omitempty"` // 短语否定词，最多200个，单词最多20字符
}

// UnitCreateSeriesCard 剧集卡片信息
type UnitCreateSeriesCard struct {
	PicId       int64    `json:"pic_id,omitempty"`      // 剧集卡封面图片ID，3:4比例，PNG/JPEG/JPG，最大2MB
	Label       []string `json:"label,omitempty"`       // 剧集卡标签，从指定选项中选2个
	Description string   `json:"description,omitempty"` // 剧集卡描述，最多100字符
}

// UnitCreateReq 创建广告组请求
type UnitCreateReq struct {
	accessTokenReq
	// 必填字段
	AdvertiserId int64             `json:"advertiser_id"` // 广告主ID，必填
	CampaignId   int64             `json:"campaign_id"`   // 广告计划ID，必填
	UnitName     string            `json:"unit_name"`     // 广告组名称，必填，1-100字符，同计划内唯一
	BidType      int               `json:"bid_type"`      // 出价类型，必填：2=CPC 10=OCPM 12=MCB最大转化量
	SceneId      []string          `json:"scene_id"`      // 资源位，必填：1=精选 5=联盟 6=全屏上下 7=双列信息流 10=联盟场景 24=激励视频 27=开屏 39=搜索 53=一天 56=内容消费
	Target       *UnitCreateTarget `json:"target"`        // 定向数据，必填
	UnitType     int               `json:"unit_type"`     // 创意制作方式，必填：3=DPA自定义 4=常规自定义 7=程序化2.0 10=智能创意
	BeginTime    string            `json:"begin_time"`    // 投放开始时间，必填，格式：yyyy-MM-dd，须>=当前日期

	// 时间与预算
	EndTime           string   `json:"end_time,omitempty"`            // 投放结束时间，格式：yyyy-MM-dd，须>=begin_time
	DayBudget         int64    `json:"day_budget,omitempty"`          // 日预算，单位：分；0=不限；最低100元，最高1亿元；不能与day_budget_schedule同用
	DayBudgetSchedule []string `json:"day_budget_schedule,omitempty"` // 分日预算，单位：分；不能与day_budget同用；优先级高于day_budget

	// 出价相关
	Bid                int64   `json:"bid,omitempty"`                  // 出价，bid_type=CPC时必填，单位：分；0.2-100元
	CpaBid             int64   `json:"cpa_bid,omitempty"`              // 出价，bid_type=OCPM时必填，单位：分
	DeepConversionBid  int64   `json:"deep_conversion_bid,omitempty"`  // 深度转化目标出价，单位：分；须>cpa_bid
	OcpxActionType     int     `json:"ocpx_action_type,omitempty"`     // 优化目标，bid_type=OCPM时必填
	DeepConversionType int64   `json:"deep_conversion_type,omitempty"` // 深度转化目标
	RoiRatio           float64 `json:"roi_ratio,omitempty"`            // 付费ROI系数，优化目标为ROI时必填，范围(0,100]

	// 应用相关
	AppId           int64    `json:"app_id,omitempty"`            // 应用ID，campaign_type=2/7/9时必填
	AppDownloadType int      `json:"app_download_type,omitempty"` // 应用下载方式：0=直接下载 1=落地页下载
	AppIconUrl      string   `json:"app_icon_url,omitempty"`      // 应用图标URL
	AppStore        []string `json:"app_store,omitempty"`         // 应用商店列表：huawei/oppo/vivo/xiaomi/meizu/smartisan/honor
	UseAppMarket    int      `json:"use_app_market,omitempty"`    // 从系统应用商店下载：0=否 1=优先系统应用商店
	PackageId       int64    `json:"package_id,omitempty"`        // 应用包ID

	// URL与落地页
	Url             string `json:"url,omitempty"`               // 推广链接，最多1000字符
	UrlType         int    `json:"url_type,omitempty"`          // 链接类型，campaign_type=3时必填
	WebUriType      int    `json:"web_uri_type,omitempty"`      // 落地页类型：1=自有链接 2=建站 3=程序化落地页 4=微信小程序
	SiteId          int64  `json:"site_id,omitempty"`           // 建站ID，web_uri_type=2时必填
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
	NegativeWordParam *UnitCreateNegativeWord `json:"negative_word_param,omitempty"` // 搜索广告否定词
	ExtendSearch      bool                    `json:"extend_search,omitempty"`       // 智能词扩展：false=关闭 true=开启，默认false
	QuickSearch       int                     `json:"quick_search,omitempty"`        // 搜索极速启动：0=关闭（默认）1=开启
	TargetExplore     int                     `json:"target_explore,omitempty"`      // 搜索受众探索：0=关闭 1=开启；需quick_search=1

	// 版位与样式
	AdType          int `json:"ad_type,omitempty"`           // 广告类型
	DetailUnitType  int `json:"detail_unit_type,omitempty"`  // 详情单元类型
	CardType        int `json:"card_type,omitempty"`         // 卡片类型
	OuterLoopNative int `json:"outer_loop_native,omitempty"` // 开启原生展示：0=关闭 1=开启
	ShowMode        int `json:"show_mode,omitempty"`         // 创意展示方式：1=轮播 2=精选（默认）

	// 短剧
	SeriesId                 int64                 `json:"series_id,omitempty"`                    // 短剧ID，campaign_type=30时必填
	EpisodeId                int64                 `json:"episode_id,omitempty"`                   // 剧集ID，campaign_type=30时必填
	SeriesCardType           int                   `json:"series_card_type,omitempty"`             // 剧集卡开关：0=关闭 1=开启
	SeriesCardInfo           *UnitCreateSeriesCard `json:"series_card_info,omitempty"`             // 剧集卡信息，开启剧集卡时必填
	SeriesPayTemplateId      int64                 `json:"series_pay_template_id,omitempty"`       // 付费模板ID
	SeriesPayMode            int                   `json:"series_pay_mode,omitempty"`              // 付费模式：1=套餐 2=虚拟货币 3=看广告解锁
	SeriesPayTemplateIdMulti []int64               `json:"series_pay_template_id_multi,omitempty"` // 短剧付费模板列表，最多5个

	// 小说
	KwaiBookId int64 `json:"kwai_book_id,omitempty"` // 快手小说ID，campaign_type=34时必填

	// 直播
	JingleBellId      int64 `json:"jingle_bell_id,omitempty"`      // 铃铛组件ID，campaign_type=16时必填
	LiveUserId        int64 `json:"live_user_id,omitempty"`        // 主播ID，campaign_type=16/30/34时必填
	LiveComponentType int   `json:"live_component_type,omitempty"` // 直播组件类型：0=铃铛 1=房产 2=团购 3=服务键 4=无组件 5=小程序 6=手柄 14=KS招聘

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

	// 微信小程序/电商
	UnitMaterialType int `json:"unit_material_type,omitempty"` // 广告素材类型，campaign_type=32/35时必填

	// 其他
	ImMessageMount        bool     `json:"im_message_mount,omitempty"`        // 挂载私信留资组件
	ComponentId           int64    `json:"component_id,omitempty"`            // 咨询组件ID
	ConsultId             int64    `json:"consult_id,omitempty"`              // 咨询组件ID
	ConversionType        int      `json:"conversion_type,omitempty"`         // 转化类型
	ConvertId             int64    `json:"convert_id,omitempty"`              // 转化追踪工具ID
	EnhanceConversionType int      `json:"enhance_conversion_type,omitempty"` // 增强目标：8=7日留存
	PutStatus             int      `json:"put_status,omitempty"`              // 广告组状态：1=投放 2=暂停，默认1
	ScheduleTime          string   `json:"schedule_time,omitempty"`           // 投放时段，24*7格式，0/1字符串
	NightScheduledTag     int      `json:"night_scheduled_tag,omitempty"`     // 夜间投放时段类型：0=非夜间 1=全夜间 2=指定夜间
	SplashAdSwitch        bool     `json:"splash_ad_switch,omitempty"`        // 开屏广告开关
	Speed                 int      `json:"speed,omitempty"`                   // 投放速度
	VideoLandingPage      bool     `json:"video_landing_page,omitempty"`      // 视频落地页
	BusinessInterest      []string `json:"business_interest,omitempty"`       // 商业兴趣
}

func (receiver *UnitCreateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *UnitCreateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.CampaignId <= 0 {
		err = errors.New("campaign_id is empty")
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
	if len(receiver.SceneId) == 0 {
		err = errors.New("scene_id is empty")
		return
	}
	if receiver.Target == nil {
		err = errors.New("target is empty")
		return
	}
	if receiver.UnitType <= 0 {
		err = errors.New("unit_type is empty")
		return
	}
	if len(receiver.BeginTime) == 0 {
		err = errors.New("begin_time is empty")
		return
	}
	return
}

// UnitCreateResp 创建广告组响应数据（仅data部分）
type UnitCreateResp struct {
	UnitId int64 `json:"unit_id"` // 广告组ID
}
