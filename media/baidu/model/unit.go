package model

const (
	// AdgroupFeedServiceURL 查询单元API端点
	AdgroupFeedServiceURL = "/json/feed/v1/AdgroupFeedService/getAdgroupFeed"
	// AdgroupFeedAddServiceURL 新建单元API端点
	AdgroupFeedAddServiceURL = "/json/feed/v1/AdgroupFeedService/addAdgroupFeed"
	// AdgroupFeedUpdateServiceURL 更新单元API端点
	AdgroupFeedUpdateServiceURL = "/json/feed/v1/AdgroupFeedService/updateAdgroupFeed"
	// AdgroupFeedDeleteServiceURL 删除单元API端点
	AdgroupFeedDeleteServiceURL = "/json/feed/v1/AdgroupFeedService/deleteAdgroupFeed"
)

// ID类型枚举
const (
	IdTypeCampaign = 1 // 计划类型
	IdTypeUnit     = 2 // 单元类型
)

// 单元状态枚举
const (
	UnitStatusActive              = 0 // 有效
	UnitStatusPaused              = 1 // 暂停推广
	UnitStatusCampaignPaused      = 2 // 推广计划暂停推广
	UnitStatusLiveEnded           = 3 // 直播结束后暂停
	UnitStatusAccountUnbindAnchor = 6 // 账户解绑主播暂停
)

// 单元学习状态枚举
const (
	UnitOcpxStatusLearning = 1 // 正在学习
	UnitOcpxStatusSuccess  = 2 // 学习成功
	UnitOcpxStatusFailed   = 3 // 学习失败
)

// 低效单元标识枚举
const (
	UnefficientAdgroupNormal        = 0 // 非低效单元
	UnefficientAdgroupLowEfficiency = 1 // 低效单元
)

// 流量来源枚举（单元级别）
const (
	FtypeSelectionUnit = 1 // 单元单独设置流量
	FtypeSelectionPlan = 2 // 使用计划流量设置
)

// 出价来源枚举（单元级别）
const (
	BidSourceUnit = 1 // 单元单独设置出价
	BidSourcePlan = 2 // 使用计划出价
)

// 投放范围枚举（单元级别）
const (
	FtypeBaiduFeed   = 1  // 自定义类-百度信息流
	FtypeTieba       = 2  // 自定义类-贴吧
	FtypeBaiqingteng = 4  // 百青藤
	FtypeHaokanVideo = 8  // 自定义类-好看视频
	FtypeBaiduNovel  = 64 // 自定义类-百度小说
)

// 投放场景枚举
const (
	DeliveryTypeAll    = 0 // 不限
	DeliveryTypeSplash = 1 // 开屏
	DeliveryTypeReward = 2 // 激励
	DeliveryTypeNative = 4 // 原生
)

// 落地页类型枚举
const (
	UrlTypeNormal      = 1 // 普通落地页
	UrlTypeMiniProgram = 2 // 百度小程序
	UrlTypeLiveRoom    = 3 // 直播间
	UrlTypeBaijiaHao   = 4 // 百家号
)

// 百度小程序类型枚举
const (
	MiniProgramTypeMini = 1 // 小程序
	MiniProgramTypeGame = 2 // 小游戏
)

// 直播间投放模式枚举
const (
	BroadCastModeDefault    = 1 // 默认（已下线）
	BroadCastModeLiveOnly   = 2 // 仅直播时投放（已下线）
	BroadCastModeContinuous = 3 // 连续投放
)

// 付费次数优化枚举
const (
	TransTypeAttrPayCount = 1 // 按照付费人数优化
	TransTypeAttrPayNum   = 2 // 按照付费次数优化
)

// 操作符枚举（商品组规则）
const (
	OperationEqual        = "EQUAL"         // 等于
	OperationNotEqual     = "NOT_EQUAL"     // 不等于
	OperationContain      = "CONTAIN"       // 包含
	OperationNotContain   = "NOT_CONTAIN"   // 不包含
	OperationGreater      = "GREATER"       // 大于
	OperationLess         = "LESS"          // 小于
	OperationGreaterEqual = "GREATER_EQUAL" // 大于等于
	OperationLessEqual    = "LESS_EQUAL"    // 小于等于
)

// AdgroupFeedReq 查询单元请求
type AdgroupFeedReq struct {
	AdgroupFeedFields []string `json:"adgroupFeedFields"` // 需要查询的单元属性（必填）
	Ids               []int64  `json:"ids"`               // 待查询计划/单元ID集合（必填）长度限制 [0, 100]
	IdType            int      `json:"idType"`            // ID类型（必填）1-计划类型, 2-单元类型
}

// Format 格式化请求参数
func (r *AdgroupFeedReq) Format() {}

// Validate 校验请求参数
func (r *AdgroupFeedReq) Validate() error {
	return nil
}

// ProductSetRule 商品组规则
type ProductSetRule struct {
	Field     string `json:"field"`     // 字段名
	Operation string `json:"operation"` // 操作符
	Value     string `json:"value"`     // 操作值
}

// UnitProducts 单元商品筛选设置
type UnitProducts struct {
	CatalogID    int64            `json:"catalogId"`              // 商品目录id
	RuleProducts []ProductSetRule `json:"ruleProducts,omitempty"` // 商品组规则
}

// AdgroupFeedOcpcType 单元oCPC设置对象
type AdgroupFeedOcpcType struct {
	AppTransID         int64   `json:"appTransId"`         // 转化追踪ID
	TransFrom          int     `json:"transFrom"`          // 接入方式
	OcpcBid            float64 `json:"ocpcBid"`            // 目标转化出价
	LpUrl              string  `json:"lpUrl"`              // 推广URL
	TransType          int     `json:"transType"`          // 目标转化
	OptimizeDeepTrans  bool    `json:"optimizeDeepTrans"`  // 优化深度转化
	DeepOcpcBid        float64 `json:"deepOcpcBid"`        // 深度转化出价
	DeepTransType      int     `json:"deepTransType"`      // 深度转化类型
	UrlType            int     `json:"urlType"`            // 落地页类型
	UseRoi             bool    `json:"useRoi"`             // 使用ROI优化
	RoiRatio           float64 `json:"roiRatio"`           // ROI转化率
	MiniProgramType    int     `json:"miniProgramType"`    // 百度小程序类型
	AppKey             string  `json:"appKey"`             // 百度小程序appkey
	PagePath           string  `json:"pagePath"`           // 百度小程序页面路径
	BroadCastMode      int     `json:"broadCastMode"`      // 直播间投放模式
	AnchorId           int64   `json:"anchorId"`           // 主播连续投放时的主播ID
	TransTypeAttribute int     `json:"transTypeAttribute"` // 付费次数优化
}

// AdgroupFeedData 推广单元信息数据
type AdgroupFeedData struct {
	AdgroupFeedId      int64                `json:"adgroupFeedId"`          // 推广单元ID
	CampaignFeedId     int64                `json:"campaignFeedId"`         // 推广计划ID
	AdgroupFeedName    string               `json:"adgroupFeedName"`        // 推广单元名称
	Pause              bool                 `json:"pause"`                  // 暂停/启用推广单元
	Status             int                  `json:"status"`                 // 推广单元状态
	Audience           map[string]string    `json:"audience"`               // 定向设置
	Bid                float64              `json:"bid"`                    // 出价
	Ftypes             []int                `json:"ftypes"`                 // 投放范围
	Bidtype            int                  `json:"bidtype"`                // 优化目标和付费模式
	Ocpc               *AdgroupFeedOcpcType `json:"ocpc,omitempty"`         // oCPC设置对象
	AtpFeedId          int64                `json:"atpFeedId"`              // 定向包ID
	AddTime            string               `json:"addtime"`                // 添加时间
	ModTime            string               `json:"modtime"`                // 修改时间
	DeliveryType       []int                `json:"deliveryType"`           // 投放场景
	UnefficientAdgroup int                  `json:"unefficientAdgroup"`     // 低效单元标识
	ProductSetId       int64                `json:"productSetId"`           // 商品组ID（仅商品推广）
	UnitProducts       *UnitProducts        `json:"unitProducts,omitempty"` // 单元商品筛选设置
	FtypeSelection     int                  `json:"ftypeSelection"`         // 是否使用计划流量
	BidSource          int                  `json:"bidSource"`              // 是否使用计划出价
	UrlType            int                  `json:"urlType"`                // 落地页类型
	MiniProgram        string               `json:"miniProgram"`            // 小程序信息（json字符串）
	BroadCastInfo      string               `json:"broadCastInfo"`          // 直播间信息（json字符串）
	Url                string               `json:"url"`                    // 落地页
	UnitOcpxStatus     int                  `json:"unitOcpxStatus"`         // 单元学习状态
	AtpName            string               `json:"atpName"`                // 定向包名称
}

// AdgroupFeedDataList 推广单元信息数据列表
type AdgroupFeedDataList struct {
	Data []AdgroupFeedData `json:"data"`
}

// BjhProgram 百家号短剧合集信息
type BjhProgram struct {
	Field         string `json:"field,omitempty"`         // 字段名
	BjhCollectId  int64  `json:"bjhCollectId,omitempty"`  // 百家号短剧合集ID
	BjhVideoId    int64  `json:"bjhVideoId,omitempty"`    // 百家号短剧合集第一集视频ID
	BjhTemplateId string `json:"bjhTemplateId,omitempty"` // 百家号付费面板ID
}

// AutoIdeaOptiType 自动创意优化设置
type AutoIdeaOptiType struct {
	TextOpti  int `json:"textOpti,omitempty"`  // 自动文案优化 0:关闭, 1:开启
	VideoOpti int `json:"videoOpti,omitempty"` // 自动视频优化 0:关闭, 1:开启
	ImageOpti int `json:"imageOpti,omitempty"` // 自动图片优化 0:关闭, 1:开启
}

// AdgroupFeedType 新建单元对象
type AdgroupFeedType struct {
	CampaignFeedId  int64                `json:"campaignFeedId"`           // 推广计划ID（必填）
	AdgroupFeedName string               `json:"adgroupFeedName"`          // 推广单元名称（必填）[1, 100]
	Pause           *bool                `json:"pause,omitempty"`          // 暂停/启用推广单元，默认false
	Audience        map[string]string    `json:"audience,omitempty"`       // 定向设置
	Bid             float64              `json:"bid"`                      // 出价（必填）
	Ftypes          []int                `json:"ftypes"`                   // 投放范围（必填）
	Bidtype         int                  `json:"bidtype,omitempty"`        // 优化目标和付费模式，默认1
	Ocpc            *AdgroupFeedOcpcType `json:"ocpc,omitempty"`           // oCPC设置对象（bidtype=3时必填）
	AtpFeedId       int64                `json:"atpFeedId,omitempty"`      // 定向包ID
	DeliveryType    []int                `json:"deliveryType,omitempty"`   // 投放场景，默认0
	ProductSetId    int64                `json:"productSetId,omitempty"`   // 商品组ID
	UnitProducts    *UnitProducts        `json:"unitProducts,omitempty"`   // 单元商品筛选设置
	FtypeSelection  int                  `json:"ftypeSelection,omitempty"` // 流量来源
	BidSource       int                  `json:"bidSource,omitempty"`      // 出价来源
	UrlType         int                  `json:"urlType,omitempty"`        // 落地页类型
	MiniProgram     string               `json:"miniProgram,omitempty"`    // 小程序信息（json字符串）
	BroadCastInfo   string               `json:"broadCastInfo,omitempty"`  // 直播间信息（json字符串）
	Url             string               `json:"url,omitempty"`            // 落地页
	CreateAtp       *bool                `json:"createAtp,omitempty"`      // 是否创建定向包，默认false
	AtpName         string               `json:"atpName,omitempty"`        // 定向包名称 [0, 60]
	AtpDesc         string               `json:"atpDesc,omitempty"`        // 定向包描述 [0, 80]
	BjhProgram      *BjhProgram          `json:"bjhProgram,omitempty"`     // 百家号短剧合集信息（小流量）
	AutoIdeaOpti    *AutoIdeaOptiType    `json:"autoIdeaOpti,omitempty"`   // 自动创意优化设置
}

// AdgroupFeedAddReq 新建单元请求
type AdgroupFeedAddReq struct {
	AdgroupFeedTypes []AdgroupFeedType `json:"adgroupFeedTypes"` // 单元集合
}

// Format 格式化请求参数
func (r *AdgroupFeedAddReq) Format() {}

// Validate 校验请求参数
func (r *AdgroupFeedAddReq) Validate() error {
	return nil
}

// AdgroupFeedUpdateType 更新单元对象
type AdgroupFeedUpdateType struct {
	AdgroupFeedId   int64                `json:"adgroupFeedId"`             // 推广单元ID（必填）
	AdgroupFeedName string               `json:"adgroupFeedName,omitempty"` // 推广单元名称 [1, 100]
	Pause           *bool                `json:"pause,omitempty"`           // 暂停/启用推广单元
	Audience        map[string]string    `json:"audience,omitempty"`        // 定向设置
	Bid             float64              `json:"bid,omitempty"`             // 出价
	Ocpc            *AdgroupFeedOcpcType `json:"ocpc,omitempty"`            // oCPC设置对象
	AtpFeedId       int64                `json:"atpFeedId,omitempty"`       // 定向包ID（0=解除绑定）
	ProductSetId    int64                `json:"productSetId,omitempty"`    // 商品组ID
	UnitProducts    *UnitProducts        `json:"unitProducts,omitempty"`    // 单元商品筛选设置
	MiniProgram     string               `json:"miniProgram,omitempty"`     // 小程序信息（json字符串）
	BroadCastInfo   string               `json:"broadCastInfo,omitempty"`   // 直播间信息（json字符串）
	Url             string               `json:"url,omitempty"`             // 落地页
	BjhProgram      *BjhProgram          `json:"bjhProgram,omitempty"`      // 百家号短剧合集信息
}

// AdgroupFeedUpdateReq 更新单元请求
type AdgroupFeedUpdateReq struct {
	AdgroupFeedTypes []AdgroupFeedUpdateType `json:"adgroupFeedTypes"` // 单元集合
}

// Format 格式化请求参数
func (r *AdgroupFeedUpdateReq) Format() {}

// Validate 校验请求参数
func (r *AdgroupFeedUpdateReq) Validate() error {
	return nil
}

// AdgroupFeedDeleteReq 删除单元请求
type AdgroupFeedDeleteReq struct {
	AdgroupFeedIds []int64 `json:"adgroupFeedIds"` // 要删除的单元ID集合
}

// Format 格式化请求参数
func (r *AdgroupFeedDeleteReq) Format() {}

// Validate 校验请求参数
func (r *AdgroupFeedDeleteReq) Validate() error {
	return nil
}

// AdgroupFeedDeleteData 删除单元响应数据
type AdgroupFeedDeleteData struct {
	AdgroupFeedId int64 `json:"adgroupFeedId"` // 推广单元ID
}

// AdgroupFeedDeleteDataList 删除单元响应数据列表
type AdgroupFeedDeleteDataList struct {
	Data []AdgroupFeedDeleteData `json:"data"`
}
