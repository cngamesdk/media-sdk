package model

const (
	// CampaignFeedServiceURL 查询计划API端点
	CampaignFeedServiceURL       = "/json/feed/v1/CampaignFeedService/getCampaignFeed"
	CampaignFeedAddServiceURL    = "/json/feed/v1/CampaignFeedService/addCampaignFeed"
	CampaignFeedUpdateServiceURL = "/json/feed/v1/CampaignFeedService/updateCampaignFeed"
	CampaignFeedDeleteServiceURL = "/json/feed/v1/CampaignFeedService/deleteCampaignFeed"
)

// 计划状态枚举
const (
	CampaignStatusActive                 = 0  // 有效
	CampaignStatusInSchedule             = 1  // 处于暂停时段
	CampaignStatusPaused                 = 2  // 暂停推广
	CampaignStatusBudgetInsufficient     = 3  // 推广计划预算不足
	CampaignStatusPendingActive          = 4  // 账户待激活
	CampaignStatusAcctBudgetInsufficient = 11 // 账户预算不足
	CampaignStatusZeroBalance            = 20 // 账户余额为零
	CampaignStatusBanned                 = 23 // 被禁推
	CampaignStatusAppOffline             = 24 // app已下线
	CampaignStatusAppAuditing            = 25 // 应用审核中
	CampaignStatusRTAPaused              = 26 // RTA计划暂停
	CampaignStatusPreorderExpired        = 27 // 计划绑定的新游预约包预约过期
	CampaignStatusProjectPaused          = 28 // 项目暂停
)

// 计划物料类型枚举
const (
	BsTypeNormal    = 1 // 普通计划
	BsTypeProduct   = 3 // 商品计划
	BsTypeNativeRTA = 7 // 原生RTA
)

// 计划类型枚举
const (
	CampaignTypeNormal    = 1 // 普通模式
	CampaignTypeVolume    = 4 // 放量模式
	CampaignTypeTargetROI = 8 // 目标ROI出价模式
)

// CampaignFeedFilter 计划查询过滤条件
type CampaignFeedFilter struct {
	BsType []int `json:"bstype,omitempty"` // 计划类型：1-普通, 3-商品, 7-原生RTA，不填返回全部
}

// CampaignFeedReq 查询计划请求
type CampaignFeedReq struct {
	CampaignFeedFields []string            `json:"campaignFeedFields"`           // 需要查询的计划属性（必填）
	CampaignFeedIds    []int64             `json:"campaignFeedIds,omitempty"`    // 查询计划ID集合 [0, 100]，空=全部
	CampaignFeedFilter *CampaignFeedFilter `json:"campaignFeedFilter,omitempty"` // 计划查询过滤条件
}

func (r *CampaignFeedReq) Format()         {}
func (r *CampaignFeedReq) Validate() error { return nil }

// ScheduleType 推广计划暂停时段
type ScheduleType struct {
	StartHour int `json:"startHour,omitempty"` // 开始小时
	StartMin  int `json:"startMin,omitempty"`  // 开始分钟
	EndHour   int `json:"endHour,omitempty"`   // 结束小时
	EndMin    int `json:"endMin,omitempty"`    // 结束分钟
	WeekDay   int `json:"weekDay,omitempty"`   // 星期几
}

// AppInfoShadowType 应用影子计划信息
type AppInfoShadowType struct {
	AppInfo AppInfoType `json:"appinfo,omitempty"` // 影子APP信息
	Status  int         `json:"status,omitempty"`  // 影子状态
}

// CampaignFeedData 计划信息数据
type CampaignFeedData struct {
	CampaignFeedID        int64                 `json:"campaignFeedId"`                 // 信息流计划ID
	CampaignFeedName      string                `json:"campaignFeedName"`               // 计划名称 [1, 100]
	Subject               int                   `json:"subject"`                        // 营销目标
	AppInfo               []AppInfoType         `json:"appinfo,omitempty"`              // 推广app信息（subject=2或3时有效）
	Budget                float64               `json:"budget"`                         // 计划预算 [50, Min(10000000, 账户预算)]，null=不限
	StartTime             string                `json:"starttime"`                      // 推广开始时间 '2016-12-15'，null=长期投放
	EndTime               string                `json:"endtime"`                        // 推广结束时间，null=长期投放
	Schedule              []ScheduleType        `json:"schedule,omitempty"`             // 推广计划暂停时段
	Pause                 bool                  `json:"pause"`                          // 计划启停
	Status                int                   `json:"status"`                         // 推广计划状态
	BsType                int                   `json:"bstype"`                         // 物料类型：1-普通, 3-商品, 7-原生RTA
	CampaignType          int                   `json:"campaignType"`                   // 计划类型：1-普通, 4-放量, 8-目标ROI
	AddTime               string                `json:"addtime"`                        // 添加时间
	EshopType             string                `json:"eshopType,omitempty"`            // 交易所在平台
	Shadow                *AppInfoShadowType    `json:"shadow,omitempty"`               // 应用推广影子计划
	BudgetOfflineTime     string                `json:"budgetOfflineTime"`              // 当天计划预算下线最近一次时间
	RTAStatus             int                   `json:"rtaStatus"`                      // 是否开通RTA
	Ftypes                []int                 `json:"ftypes,omitempty"`               // 投放范围
	BidType               int                   `json:"bidtype"`                        // 出价方式
	Bid                   float64               `json:"bid"`                            // 出价（bidtype=3时有效）
	Ocpc                  *OcpcModel            `json:"ocpc,omitempty"`                 // oCPC信息
	UnefficientCampaign   int                   `json:"unefficientCampaign"`            // 低效计划
	CampaignOcpxStatus    int                   `json:"campaignOcpxStatus"`             // 计划学习状态
	BmcUserID             int64                 `json:"bmcUserId"`                      // 商品中心用户ID
	CatalogID             int64                 `json:"catalogId"`                      // 商品目录ID
	CatalogSource         int                   `json:"catalogSource"`                  // 产品目录来源
	ProductType           int                   `json:"productType"`                    // 产品库类型
	ProjectFeedID         int64                 `json:"projectFeedId"`                  // 项目ID
	InheritAscriptionType int                   `json:"inheritAscriptionType"`          // 继承归属
	InheritUserids        []int64               `json:"inheritUserids,omitempty"`       // 继承优质计划账户ID集合
	InheritCampaignInfos  []InheritCampaignInfo `json:"inheritCampaignInfos,omitempty"` // 继承优质计划信息集合
	UseLiftBudget         int                   `json:"useLiftBudget"`                  // 是否开启一键起量
	LiftBudget            float64               `json:"liftBudget"`                     // 起量预算
	LiftStatus            int                   `json:"liftStatus"`                     // 起量状态
	DeliveryType          int                   `json:"deliveryType"`                   // 投放场景
	AppSubType            int                   `json:"appSubType"`                     // 应用推广子类型
	MiniProgramType       int                   `json:"miniProgramType"`                // 小程序子类型
	BidMode               int                   `json:"bidMode"`                        // 出价模式
	ProductIDs            string                `json:"productIds"`                     // 产品ID
	SaleType              int                   `json:"saleType"`                       // 营销场景
	LiftBudgetSchedule    *LiftBudgetSchedule   `json:"liftBudgetSchedule,omitempty"`   // 起量生效时间
}

// InheritCampaignInfo 继承优质计划信息
type InheritCampaignInfo struct {
	CampaignID int64   `json:"campaignId,omitempty"` // 计划ID
	UserID     int64   `json:"userId,omitempty"`     // 账户ID
	Budget     float64 `json:"budget,omitempty"`     // 预算
}

// CampaignFeedDataList 计划信息数据列表
type CampaignFeedDataList struct {
	Data []CampaignFeedData `json:"data"`
}

// 出价方式枚举（新建计划）
const (
	BidTypeCPC  = 1 // 点击（CPC）
	BidTypeCPM  = 2 // 曝光（CPM）
	BidTypeOCPC = 3 // 转化（oCPC/oCPM）
)

// 下载方式枚举
const (
	DownloadTypeDirect   = 0 // 直接下载
	DownloadTypeLandPage = 1 // 落地页下载
)

// 继承归属类型枚举
const (
	InheritAscriptionCurrentAccount  = 1 // 当前账户
	InheritAscriptionSameCustomer    = 2 // 同客户中心
	InheritAscriptionCurrentCampaign = 3 // 当前账户内的计划
	InheritAscriptionSameCampaign    = 4 // 同客户中心的计划
)

// CampaignFeedType 新建/更新计划中的计划对象
type CampaignFeedType struct {
	CampaignFeedID        int64                 `json:"campaignFeedId,omitempty"`        // 计划ID（更新时必填）
	CampaignFeedName      string                `json:"campaignFeedName"`                // 计划名称（新建时必填）[1, 100]
	Subject               int                   `json:"subject"`                         // 营销目标（必填）
	AppInfo               *AppInfoType          `json:"appinfo,omitempty"`               // 推广app信息（subject=2或3时有效）
	Budget                float64               `json:"budget,omitempty"`                // 计划预算 [50, 9999999.99]，null=不限
	StartTime             string                `json:"starttime,omitempty"`             // 推广开始时间
	EndTime               string                `json:"endtime,omitempty"`               // 推广结束时间
	Schedule              []ScheduleType        `json:"schedule,omitempty"`              // 推广计划暂停时段
	Pause                 *bool                 `json:"pause,omitempty"`                 // 计划启停
	BsType                int                   `json:"bstype,omitempty"`                // 物料类型：1-普通(默认), 3-商品, 7-原生RTA
	CampaignType          int                   `json:"campaignType,omitempty"`          // 计划类型：1-普通, 4-放量, 8-目标ROI
	EshopType             string                `json:"eshopType,omitempty"`             // 交易所在平台
	Ftypes                []int                 `json:"ftypes,omitempty"`                // 流量类型
	BidType               int                   `json:"bidtype,omitempty"`               // 出价方式：1-CPC(默认), 2-CPM, 3-oCPC
	Bid                   float64               `json:"bid,omitempty"`                   // 出价
	Ocpc                  *OcpcModel            `json:"ocpc,omitempty"`                  // oCPC设置（bidtype=3时有效）
	BmcUserID             int64                 `json:"bmcUserId,omitempty"`             // 商品中心用户ID
	CatalogID             int64                 `json:"catalogId,omitempty"`             // 商品目录ID
	CatalogSource         int                   `json:"catalogSource,omitempty"`         // 产品目录来源
	ProductType           int                   `json:"productType,omitempty"`           // 产品库类型
	ProjectFeedID         int64                 `json:"projectFeedId,omitempty"`         // 项目ID（0或不传=不关联）
	MiniProgramType       int                   `json:"miniProgramType,omitempty"`       // 小程序子类型
	AppSubType            int                   `json:"appSubType,omitempty"`            // 应用推广子类型
	BidMode               int                   `json:"bidMode,omitempty"`               // 出价模式
	SaleType              int                   `json:"saleType,omitempty"`              // 营销场景
	InheritAscriptionType int                   `json:"inheritAscriptionType,omitempty"` // 继承归属类型
	InheritUserids        []int64               `json:"inheritUserids,omitempty"`        // 继承优质计划账户ID集合
	InheritCampaignInfos  []InheritCampaignInfo `json:"inheritCampaignInfos,omitempty"`  // 继承优质计划信息集合
	UseLiftBudget         int                   `json:"useLiftBudget,omitempty"`         // 是否开启一键起量
	LiftBudget            float64               `json:"liftBudget,omitempty"`            // 起量预算
	DeliveryType          []int                 `json:"deliveryType,omitempty"`          // 投放场景
	LiftBudgetSchedule    *LiftBudgetSchedule   `json:"liftBudgetSchedule,omitempty"`    // 起量生效时间
}

// CampaignFeedAddReq 新建计划请求
type CampaignFeedAddReq struct {
	CampaignFeedTypes []CampaignFeedType `json:"campaignFeedTypes"` // 计划集合
}

func (r *CampaignFeedAddReq) Format()         {}
func (r *CampaignFeedAddReq) Validate() error { return nil }

// CampaignFeedUpdateReq 更新计划请求（复用CampaignFeedType，campaignFeedId必填）
type CampaignFeedUpdateReq struct {
	CampaignFeedTypes []CampaignFeedType `json:"campaignFeedTypes"`
}

func (r *CampaignFeedUpdateReq) Format()         {}
func (r *CampaignFeedUpdateReq) Validate() error { return nil }

// CampaignFeedDeleteReq 删除计划请求
type CampaignFeedDeleteReq struct {
	CampaignFeedIds []int64 `json:"campaignFeedIds"` // 要删除的计划ID集合
}

func (r *CampaignFeedDeleteReq) Format()         {}
func (r *CampaignFeedDeleteReq) Validate() error { return nil }

// CampaignFeedDeleteData 删除计划响应数据
type CampaignFeedDeleteData struct {
	CampaignFeedID int64 `json:"campaignFeedId"` // 信息流计划ID
}

// CampaignFeedDeleteDataList 删除计划响应数据列表
type CampaignFeedDeleteDataList struct {
	Data []CampaignFeedDeleteData `json:"data"`
}
