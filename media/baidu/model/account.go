package model

const (
	// 账户API端点
	AccountFeedServiceURL       = "/json/feed/v1/AccountFeedService/getAccountFeed"
	AccountFeedUpdateServiceURL = "/json/feed/v1/AccountFeedService/updateAccountFeed"
)

// 账户状态枚举
const (
	UserStatUnpaid             = 1  // 开户金未到
	UserStatActive             = 2  // 生效
	UserStatZeroBalance        = 3  // 账户余额为0
	UserStatRejected           = 4  // 被拒绝
	UserStatAuditing           = 6  // 审核中
	UserStatDisabled           = 7  // 被禁用
	UserStatPendingActive      = 8  // 待激活
	UserStatBudgetInsufficient = 11 // 账户预算不足
)

// 资金包类型枚举
const (
	BalancePackageFeed      = 0 // 信息流资金包
	BalancePackageSearch    = 1 // 搜索推广资金包
	BalancePackageAgentFeed = 2 // 代理商信息流资金包
)

// 投放流量枚举
const (
	FlowBaiduApp    = 1  // 百度APP
	FlowTieba       = 2  // 贴吧
	FlowBaiQingTeng = 4  // 百青藤
	FlowHaoKan      = 8  // 好看视频
	FlowBaiduNovel  = 64 // 百度小说
)

// Feed产品线权限状态
const (
	UAStatusOpened     = 1 // 已开通
	UAStatusPending    = 2 // 待开通
	UAStatusNotAllowed = 3 // 不允许开通（KA客户）
)

// 资金包类型（adtype）
const (
	AdTypeKA     = 187 // KA客户标识，使用KA原生资金池
	AdTypeDirect = 188 // 中小客户中的直销/分公司客户
	AdTypeAgent  = 189 // 中小客户中的代理商客户
)

// AccountFeedReq 查询信息流账户信息请求
type AccountFeedReq struct {
	AccountFeedFields []string `json:"accountFeedFields,omitempty"` // 指定需要返回的属性，枚举值见文档
}

// Format 格式化请求参数
func (r *AccountFeedReq) Format() {}

// Validate 校验请求参数
func (r *AccountFeedReq) Validate() error {
	return nil
}

// AccountFeedData 账户信息数据
type AccountFeedData struct {
	UserID            int64               `json:"userId"`            // 账户ID
	Cid               int64               `json:"cid"`               // 账户主体ID（客户不存在主体则cid为0）
	LiceName          string              `json:"liceName"`          // 账户主体名称
	Balance           float64             `json:"balance"`           // 账户余额
	Budget            float64             `json:"budget"`            // 账户预算（默认为0，表示不限预算）
	BalancePackage    int                 `json:"balancePackage"`    // 资金包类型：0-信息流, 1-搜索推广, 2-代理商信息流
	UserStat          int                 `json:"userStat"`          // 账户状态：1-开户金未到, 2-生效, 3-余额为0, 4-被拒绝, 6-审核中, 7-被禁用, 8-待激活, 11-预算不足
	UAStatus          int                 `json:"uaStatus"`          // feed产品线权限：1-已开通, 2-待开通, 3-不允许开通(KA)
	ValidFlows        []int               `json:"validFlows"`        // 可投放流量：1-百度APP, 2-贴吧, 4-百青藤, 8-好看视频, 64-百度小说
	TradeID           int                 `json:"tradeId"`           // 用户行业ID
	BudgetOfflineTime []map[string]string `json:"budgetOfflineTime"` // 账户预算撞线时间
	AdType            int                 `json:"adtype"`            // 账户资金包类型：187-KA, 188-直销/分公司, 189-代理商
	RTAUserAdmin      bool                `json:"rtaUserAdmin"`      // 是否账户层级控制RTA
}

// AccountFeedDataList 账户信息数据列表
type AccountFeedDataList struct {
	Data []AccountFeedData `json:"data"`
}

// AccountFeedUpdateType 更新账户信息中的accountFeedType字段
type AccountFeedUpdateType struct {
	Budget float64 `json:"budget"` // 账户预算，取值范围：[50, 9999999.99]，默认为0表示不限预算
}

// AccountFeedUpdateReq 更新信息流账户信息请求
type AccountFeedUpdateReq struct {
	AccountFeedType AccountFeedUpdateType `json:"accountFeedType"` // 要更新的账户字段
}

// Format 格式化请求参数
func (r *AccountFeedUpdateReq) Format() {}

// Validate 校验请求参数
func (r *AccountFeedUpdateReq) Validate() error {
	return nil
}

// AccountFeedUpdateData 更新账户信息响应数据
type AccountFeedUpdateData struct {
	UserID int64   `json:"userId"` // 账户ID
	Budget float64 `json:"budget"` // 账户预算
}

// AccountFeedUpdateDataList 更新账户信息响应数据列表
type AccountFeedUpdateDataList struct {
	Data []AccountFeedUpdateData `json:"data"`
}
