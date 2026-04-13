package model

import "errors"

// ========== 短剧售卖策略-获取售卖策略列表 ==========
// https://developers.e.qq.com/v3.0/docs/api/landing_page_sell_strategy/get

// 策略状态常量
const (
	LandingPageSellStrategyStatusInvalid = 0 // 无效
	LandingPageSellStrategyStatusValid   = 1 // 有效
)

// 分页常量
const (
	MinLandingPageSellStrategyPage         = 1     // page 最小值
	MaxLandingPageSellStrategyPage         = 99999 // page 最大值
	MinLandingPageSellStrategyPageSize     = 1     // page_size 最小值
	MaxLandingPageSellStrategyPageSize     = 100   // page_size 最大值
	DefaultLandingPageSellStrategyPage     = 1     // page 默认值
	DefaultLandingPageSellStrategyPageSize = 10    // page_size 默认值
)

// 字段长度常量
const (
	MinLandingPageSellStrategyNameBytes = 1  // strategy_name / full_strategy_name 最小字节数
	MaxLandingPageSellStrategyNameBytes = 60 // strategy_name / full_strategy_name 最大字节数
)

// LandingPageSellStrategyGetReq 短剧售卖策略获取列表请求（GET）
// https://developers.e.qq.com/v3.0/docs/api/landing_page_sell_strategy/get
type LandingPageSellStrategyGetReq struct {
	GlobalReq
	AccountId        int64  `json:"account_id"`                   // 广告主帐号 id (必填)
	StrategyId       int64  `json:"strategy_id,omitempty"`        // 策略 id
	StrategyStatus   *int   `json:"strategy_status,omitempty"`    // 策略状态，0：无效，1：有效（使用指针区分"未传"与"传0"）
	StrategyName     string `json:"strategy_name,omitempty"`      // 策略名称，模糊查询，1-60字节
	FullStrategyName string `json:"full_strategy_name,omitempty"` // 策略名称，精准查询，1-60字节
	Page             int    `json:"page,omitempty"`               // 搜索页码，1-99999，默认1
	PageSize         int    `json:"page_size,omitempty"`          // 每页条数，1-100，默认10
}

func (r *LandingPageSellStrategyGetReq) Format() {
	r.GlobalReq.Format()
	if r.Page == 0 {
		r.Page = DefaultLandingPageSellStrategyPage
	}
	if r.PageSize == 0 {
		r.PageSize = DefaultLandingPageSellStrategyPageSize
	}
}

// Validate 验证获取售卖策略列表请求参数
func (r *LandingPageSellStrategyGetReq) Validate() error {
	if r.AccountId == 0 {
		return errors.New("account_id为必填")
	}
	if r.StrategyStatus != nil && (*r.StrategyStatus < 0 || *r.StrategyStatus > 1) {
		return errors.New("strategy_status须为0或1")
	}
	if r.StrategyName != "" && (len(r.StrategyName) < MinLandingPageSellStrategyNameBytes || len(r.StrategyName) > MaxLandingPageSellStrategyNameBytes) {
		return errors.New("strategy_name长度须在1-60字节之间")
	}
	if r.FullStrategyName != "" && (len(r.FullStrategyName) < MinLandingPageSellStrategyNameBytes || len(r.FullStrategyName) > MaxLandingPageSellStrategyNameBytes) {
		return errors.New("full_strategy_name长度须在1-60字节之间")
	}
	if r.Page < MinLandingPageSellStrategyPage || r.Page > MaxLandingPageSellStrategyPage {
		return errors.New("page须在1-99999之间")
	}
	if r.PageSize < MinLandingPageSellStrategyPageSize || r.PageSize > MaxLandingPageSellStrategyPageSize {
		return errors.New("page_size须在1-100之间")
	}
	return r.GlobalReq.Validate()
}

// LandingPageSellStrategyItem 售卖策略列表项
type LandingPageSellStrategyItem struct {
	StrategyId      int64   `json:"strategy_id"`       // 策略 id
	StrategyName    string  `json:"strategy_name"`     // 策略名称
	EpisodePrice    float64 `json:"episode_price"`     // 单集价格，单位元，精确到 0.01
	MinRechargeTier float64 `json:"min_recharge_tier"` // 最低充值档位，单位元，精确到 0.01
	RechargeNum     int     `json:"recharge_num"`      // 起充集数
	StrategyStatus  int     `json:"strategy_status"`   // 策略状态，0：无效，1：有效
	AccountId       int64   `json:"account_id"`        // 广告主帐号 id
}

// LandingPageSellStrategyGetResp 获取售卖策略列表响应
type LandingPageSellStrategyGetResp struct {
	List     []*LandingPageSellStrategyItem `json:"list"`      // 售卖策略列表数据
	PageInfo *PageInfo                      `json:"page_info"` // 分页配置信息
}
