package model

const (
	// TransTraceServiceURL 查询转化追踪API端点
	TransTraceServiceURL = "/json/feed/v1/SearchFeedService/getOcpcTransFeed"
)

// 接入方式枚举（转化追踪查询）
const (
	TransTraceAll              = 0  // 包含全部接入方式
	TransTraceAppAPI           = 1  // 应用API
	TransTraceJimuPage         = 2  // 基木鱼营销页
	TransTraceAPIActivate      = 4  // API激活
	TransTraceWebJS            = 5  // 网页JS布码
	TransTraceLeadsAPI         = 7  // 线索API
	TransTraceConsultTool      = 8  // 咨询工具授权
	TransTraceBaiduMiniProgram = 9  // 百度智能小程序
	TransTraceAppSDK           = 13 // 应用SDK
	TransTraceBaiduStatsWeb    = 23 // 百度统计网站导入
	TransTraceBaiduStatsMini   = 24 // 百度统计小程序导入
	TransTraceBaiduNovel       = 28 // 百度小说书城
)

// 落地页类型枚举（基木鱼过滤）
const (
	ShowTypeH5          = 0 // H5
	ShowTypePC          = 1 // PC
	ShowTypeMiniProgram = 2 // 小程序
)

// 平台枚举（基木鱼过滤）
const (
	PlatformJimu   = 1 // 基木鱼平台
	PlatformJimuEC = 2 // 基木鱼电商
)

// 监管方式枚举
const (
	ModeActivate = 1 // 激活
)

// 激活状态枚举
const (
	TransStatusActive   = 1 // 已激活
	TransStatusInactive = 0 // 未激活
)

// SearchFieldType 搜索字段
type SearchFieldType struct {
	PageName  string `json:"pageName,omitempty"`  // 页面名称 [1, 50]
	Id        int64  `json:"id,omitempty"`        // 通过id搜索
	SearchKey string `json:"searchKey,omitempty"` // 用户搜索框输入内容
}

// JmyPageFilter 基木鱼页面过滤条件
type JmyPageFilter struct {
	ShowType     int              `json:"showType,omitempty"`     // 落地页类型 0-H5, 1-PC, 2-小程序
	PlatformIds  []int            `json:"platformIds,omitempty"`  // 平台 1-基木鱼, 2-基木鱼电商
	SearchFields *SearchFieldType `json:"searchFields,omitempty"` // 搜索字段
}

// TransTraceReq 查询转化追踪请求
type TransTraceReq struct {
	TransFrom     int            `json:"transFrom"`               // 接入方式（必填）0=全部
	JmyPageFilter *JmyPageFilter `json:"jmyPageFilter,omitempty"` // 基木鱼过滤条件（transFrom=2时选填）
}

// Format 格式化请求参数
func (r *TransTraceReq) Format() {}

// Validate 校验请求参数
func (r *TransTraceReq) Validate() error {
	return nil
}

// TransTraceData 转化追踪信息数据
type TransTraceData struct {
	AppTransId     int64    `json:"appTransId"`               // 转化追踪ID
	TransFrom      int      `json:"transFrom"`                // 接入方式
	TransName      string   `json:"transName"`                // 转化名称
	TransTypes     []int    `json:"transTypes"`               // 转化类型
	MonitorUrl     string   `json:"monitorUrl,omitempty"`     // 点击监测地址（transFrom=1,4,13时返回）
	AppType        int      `json:"appType,omitempty"`        // 应用类型（transFrom=1,4,13时返回）
	DownloadUrl    string   `json:"downloadUrl,omitempty"`    // 下载URL（transFrom=1,4,13时返回）
	ExposureUrl    string   `json:"exposureUrl,omitempty"`    // 曝光监测地址（transFrom=1,4,13时返回）
	LpUrl          string   `json:"lpUrl,omitempty"`          // 转化URL（transFrom=2,5,7,8时返回）
	RelatedUrls    []string `json:"relatedUrls,omitempty"`    // 推广URL（transFrom=5,7,8时返回）
	Mode           int      `json:"mode,omitempty"`           // 监测方式（transFrom=1,4,13时返回）
	TransStatus    int      `json:"transStatus,omitempty"`    // 激活状态
	DeepTransTypes []int    `json:"deepTransTypes,omitempty"` // 深度转化类型
	Docid          string   `json:"docid,omitempty"`          // Android渠道包ID（transFrom=1,appType=2时返回）
	AppName        string   `json:"appName,omitempty"`        // 应用名称（transFrom=1,4,13时返回）
	ApkName        string   `json:"apkName,omitempty"`        // 应用包名（transFrom=1,4,13时返回）
	ChannelId      int64    `json:"channelId,omitempty"`      // Android渠道包ID（transFrom=1,appType=2时返回）
}

// TransTraceDataList 转化追踪信息数据列表
type TransTraceDataList struct {
	Data []TransTraceData `json:"data"`
}
