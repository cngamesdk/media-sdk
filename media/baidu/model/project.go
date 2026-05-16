package model

const (
	// ProjectFeedServiceURL 查询项目API端点
	ProjectFeedServiceURL    = "/json/sms/service/ProjectFeedService/getProjectFeed"
	ProjectFeedAddServiceURL = "/json/sms/service/ProjectFeedService/addProjectFeed"
)

// 营销目标枚举
const (
	SubjectAppDownloadIOS     = 2  // 应用下载（iOS）
	SubjectAppDownloadAndroid = 3  // 应用下载（Android）
	SubjectSalesLeads         = 8  // 销售线索
	SubjectAppLaunch          = 9  // 应用调起
	SubjectBaijiaHaoPromotion = 10 // 百家号推广（小流量）
)

// 项目状态枚举
const (
	ProjectStatusActive             = 0  // 项目生效
	ProjectStatusPaused             = 2  // 项目暂停
	ProjectStatusPendingActive      = 4  // 账户待激活
	ProjectStatusBudgetInsufficient = 11 // 账户预算不足
	ProjectStatusZeroBalance        = 20 // 账户余额为0
	ProjectStatusTrafficBanned      = 23 // 流量禁推
)

// 出价模式枚举
const (
	BidModeTargetCPA = 1 // 目标转化成本
	BidModeTargetROI = 3 // 目标ROI出价模式
)

// 项目学习状态枚举
const (
	OcpxStatusLearning = 1 // 正在学习
	OcpxStatusSuccess  = 2 // 学习成功
	OcpxStatusFailed   = 3 // 学习失败
)

// 产品目录来源枚举
const (
	CatalogSourceBMC  = 1 // 商品中心(BMC)
	CatalogSourceJimu = 2 // 基木鱼内容中心
)

// 推广场景枚举
const (
	AppSubTypeDownload = 0 // 应用下载
	AppSubTypePreorder = 1 // 新游预约
	AppSubTypeLaunch   = 2 // 应用调起
)

// 产品库类型枚举
const (
	ProductTypeNovel         = 1     // 小说库
	ProductTypeShortDrama    = 2     // 短剧库
	ProductTypeLongVideo     = 3     // 长视频库
	ProductTypeCourse        = 4     // 课程中心
	ProductTypeMedicalBeauty = 5     // 医美项目
	ProductTypeRetail        = 6     // 零售库
	ProductTypeNews          = 7     // 资讯库
	ProductTypeContentCenter = 10001 // 内容中心
)

// 一键起量枚举
const (
	UseLiftBudgetOn  = 1 // 开启起量预算
	UseLiftBudgetOff = 2 // 关闭起量预算
)

// 起量生效时间枚举
const (
	ScheduleModelImmediate = 1 // 立即生效
	ScheduleModelSpecified = 2 // 指定时间
	ScheduleModelWeekly    = 3 // 每周重复
)

// 转化接入方式枚举
const (
	TransFromAppAPI           = 1  // 应用API
	TransFromJimuPage         = 2  // 基木鱼营销页
	TransFromAPIActivate      = 4  // API激活
	TransFromWebJS            = 5  // 网页JS布码
	TransFromLeadsAPI         = 7  // 线索API
	TransFromConsultTool      = 8  // 咨询工具授权
	TransFromAppSDK           = 13 // 应用SDK
	TransFromBaiduHealth      = 25 // 百度健康商城
	TransFromBaiduMiniProgram = 30 // 百度小程序导入
	TransFromBaijiaHao        = 32 // 百家号自动对接
)

// 目标转化类型枚举
const (
	TransTypeConsultClick        = 1  // 咨询按钮点击
	TransTypePhoneClick          = 2  // 电话按钮点击
	TransTypeFormSubmit          = 3  // 表单提交成功
	TransTypeActivate            = 4  // 激活
	TransTypeFormButtonClick     = 5  // 表单按钮点击
	TransTypeDownloadClick       = 6  // 下载（预约）按钮点击（小流量）
	TransTypePurchaseSuccess     = 10 // 购买成功
	TransTypeOrderSubmit         = 14 // 订单提交成功
	TransTypeThreeMsgConsult     = 17 // 三句话咨询
	TransTypeLeaveLeads          = 18 // 留线索
	TransTypeOneMsgConsult       = 19 // 一句话咨询
	TransTypeKeyPageView         = 20 // 关键页面浏览
	TransTypeRegister            = 25 // 注册（小流量）
	TransTypePayment             = 26 // 付费（小流量）
	TransTypePhoneConnected      = 30 // 电话拨通
	TransTypeWxCopyClick         = 35 // 微信复制按钮点击（小流量）
	TransTypeApply               = 41 // 申请（小流量）
	TransTypeCreditGrant         = 42 // 授信（小流量）
	TransTypeProductOrderSuccess = 45 // 商品下单成功
)

// ProjectFeedReq 查询项目请求
type ProjectFeedReq struct {
	ProjectFeedFields []string `json:"projectFeedFields"`        // 需要查询的项目属性（必填）
	ProjectFeedIds    []int64  `json:"projectFeedIds,omitempty"` // 待查询项目ID集合（传空返回全部）
}

// Format 格式化请求参数
func (r *ProjectFeedReq) Format() {}

// Validate 校验请求参数
func (r *ProjectFeedReq) Validate() error {
	return nil
}

// AppInfoType 推广app信息
type AppInfoType struct {
	AppName   string `json:"appName,omitempty"`   // 推广APP名称 [1, 40]
	ApkName   string `json:"apkName,omitempty"`   // 推广APP包名 [1, 1024]（仅Android）
	AppURL    string `json:"appUrl,omitempty"`    // 推广APP链接（仅iOS，iTunes链接）
	ChannelID int64  `json:"channelId,omitempty"` // 渠道包ID（仅Android）
}

// OcpcModel oCPC设置对象
type OcpcModel struct {
	AppTransID        int64   `json:"appTransId,omitempty"`        // 转化追踪ID
	TransFrom         int     `json:"transFrom,omitempty"`         // 接入方式
	OcpcBid           float64 `json:"ocpcBid,omitempty"`           // 目标转化出价 [0.2, 99999.99]
	TransType         int     `json:"transType,omitempty"`         // 目标转化类型
	OptimizeDeepTrans bool    `json:"optimizeDeepTrans,omitempty"` // 是否开启深度转化优化
	DeepOcpcBid       float64 `json:"deepOcpcBid,omitempty"`       // 深度转化出价
	RoiRatio          float64 `json:"roiRatio,omitempty"`          // ROI系数（目标ROI出价时使用）
	DeepTransType     int     `json:"deepTransType,omitempty"`     // 深度转化类型
	OcpcLevel         int     `json:"ocpcLevel,omitempty"`         // oCPC层级
}

// LiftBudgetSchedule 一键起量设置对象
type LiftBudgetSchedule struct {
	ScheduleModel int     `json:"scheduleModel"`        // 起量生效时间：1-立即生效, 2-指定时间, 3-每周重复
	LiftBudget    float64 `json:"liftBudget,omitempty"` // 起量预算
	StartTime     string  `json:"startTime,omitempty"`  // 指定时间 (scheduleModel=2)
	StartTime1    string  `json:"startTime1,omitempty"` // 指定时间1 (scheduleModel=2)
	StartTime2    string  `json:"startTime2,omitempty"` // 指定时间2 (scheduleModel=2)
	StartTime3    string  `json:"startTime3,omitempty"` // 指定时间3 (scheduleModel=2)
	EventWeek     string  `json:"eventWeek,omitempty"`  // 每周重复时间-天 (scheduleModel=3, 如"1,3,5")
	EventHour     string  `json:"eventHour,omitempty"`  // 每周重复时间-小时 (scheduleModel=3, 如"00:00")
}

// ProjectFeedData 项目信息数据
type ProjectFeedData struct {
	ProjectFeedID     int64               `json:"projectFeedId"`     // 项目ID
	ProjectFeedName   string              `json:"projectFeedName"`   // 项目名称
	Subject           int                 `json:"subject"`           // 营销目标
	AppInfo           AppInfoType         `json:"appInfo"`           // 推广app信息
	Pause             bool                `json:"pause"`             // 项目启停：true-暂停, false-启用
	Status            int                 `json:"status"`            // 项目状态
	BidMode           int                 `json:"bidMode"`           // 出价模式
	Ocpc              OcpcModel           `json:"ocpc"`              // oCPC设置对象
	ProjectOcpxStatus int                 `json:"projectOcpxStatus"` // 项目学习状态
	BmcUserID         int64               `json:"bmcUserId"`         // 商品中心用户ID
	CatalogID         int64               `json:"catalogId"`         // 商品目录ID/产品分组ID
	CatalogSource     int                 `json:"catalogSource"`     // 产品目录来源
	AppSubType        int                 `json:"appSubType"`        // 推广场景
	ProductType       int                 `json:"productType"`       // 产品库类型
	CampaignFeedIds   []int64             `json:"campaignFeedIds"`   // 关联计划ID集合
	ProductIDs        string              `json:"productIds"`        // 产品ID
	MiniProgramType   int                 `json:"miniProgramType"`   // 小程序类型
	UseLiftBudget     int                 `json:"useLiftBudget"`     // 是否开启一键起量
	Lift              *LiftBudgetSchedule `json:"lift,omitempty"`    // 一键起量设置对象
}

// ProjectFeedDataList 项目信息数据列表
type ProjectFeedDataList struct {
	Data []ProjectFeedData `json:"data"`
}

// 智能起量枚举
const (
	AiLiftOff = 0 // 关闭智能起量
	AiLiftOn  = 1 // 开启智能起量
)

// 智能起量模式枚举
const (
	AiLiftModelExplore = 1 // 积极探索
	AiLiftModelStable  = 2 // 稳中求进
)

// 营销目标（补充新建接口特有值）
const (
	SubjectMiniProgram = 4  // 小程序（需要开通小流量名单）
	SubjectHarmonyOS   = 13 // 应用下载（harmonyos）(小流量)
)

// ProjectFeedType 新建项目中的项目对象
type ProjectFeedType struct {
	ProjectFeedName string              `json:"projectFeedName"`           // 项目名称（必填）[1, 100]
	Subject         int                 `json:"subject"`                   // 营销目标（必填）
	AppInfo         *AppInfoType        `json:"appInfo,omitempty"`         // 推广app信息（subject=2或3时有效）
	BidMode         int                 `json:"bidMode"`                   // 出价模式（必填）
	Ocpc            OcpcModel           `json:"ocpc"`                      // oCPC设置对象（必填）
	BmcUserID       int64               `json:"bmcUserId,omitempty"`       // 商品中心用户ID
	CatalogID       int64               `json:"catalogId,omitempty"`       // 商品目录ID（关联产品库时必填）
	CatalogSource   int                 `json:"catalogSource,omitempty"`   // 产品目录来源（关联产品库时必填）
	AppSubType      int                 `json:"appSubType,omitempty"`      // 推广场景
	ProductType     int                 `json:"productType,omitempty"`     // 产品库类型（关联产品库时必填）
	CampaignFeedIds []int64             `json:"campaignFeedIds,omitempty"` // 关联计划ID集合（不传表示不关联）
	ProductIDs      string              `json:"productIds,omitempty"`      // 产品ID（仅销售线索营销目标支持）
	MiniProgramType int                 `json:"miniProgramType,omitempty"` // 小程序类型（仅小程序营销目标支持）
	UseLiftBudget   int                 `json:"useLiftBudget,omitempty"`   // 是否开启一键起量
	Lift            *LiftBudgetSchedule `json:"lift,omitempty"`            // 一键起量设置对象（useLiftBudget=1时必填）
	AiLift          int                 `json:"aiLift,omitempty"`          // 智能起量-自动优选：0-关闭, 1-开启
	AiLiftModel     int                 `json:"aiLiftModel,omitempty"`     // 智能起量-起量模式：1-积极探索, 2-稳中求进
}

// ProjectFeedAddReq 新建项目请求
type ProjectFeedAddReq struct {
	ProjectFeedTypes []ProjectFeedType `json:"projectFeedTypes"` // 项目集合
}

// Format 格式化请求参数
func (r *ProjectFeedAddReq) Format() {}

// Validate 校验请求参数
func (r *ProjectFeedAddReq) Validate() error {
	return nil
}
