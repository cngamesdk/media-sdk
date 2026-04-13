package model

import "errors"

// ========== 短剧售卖策略-创建售卖策略 ==========
// https://developers.e.qq.com/v3.0/docs/api/landing_page_sell_strategy/add

// 价格和集数常量
const (
	MinLandingPageSellStrategyEpisodePrice    = 0.01   // episode_price 最小值，单位元
	MaxLandingPageSellStrategyEpisodePrice    = 999999 // episode_price 最大值，单位元
	MinLandingPageSellStrategyMinRechargeTier = 0.01   // min_recharge_tier 最小值，单位元
	MaxLandingPageSellStrategyMinRechargeTier = 999999 // min_recharge_tier 最大值，单位元
	MinLandingPageSellStrategyRechargeNum     = 0      // recharge_num 最小值
	MaxLandingPageSellStrategyRechargeNum     = 999999 // recharge_num 最大值
)

// LandingPageSellStrategyAddReq 短剧售卖策略创建请求（POST JSON）
// https://developers.e.qq.com/v3.0/docs/api/landing_page_sell_strategy/add
type LandingPageSellStrategyAddReq struct {
	GlobalReq
	AccountId       int64   `json:"account_id"`        // 推广帐号 id (必填)
	StrategyName    string  `json:"strategy_name"`     // 策略名称，1-60字节 (必填)
	EpisodePrice    float64 `json:"episode_price"`     // 单集价格，>=0.01，精确到0.01，单位元 (必填)
	MinRechargeTier float64 `json:"min_recharge_tier"` // 最低充值档位，>=0.01，精确到0.01，单位元 (必填)
	RechargeNum     int     `json:"recharge_num"`      // 起充集数，>=0 (必填)
}

func (r *LandingPageSellStrategyAddReq) Format() {
	r.GlobalReq.Format()
}

// Validate 验证创建售卖策略请求参数
func (r *LandingPageSellStrategyAddReq) Validate() error {
	if r.AccountId == 0 {
		return errors.New("account_id为必填")
	}
	if len(r.StrategyName) < MinLandingPageSellStrategyNameBytes || len(r.StrategyName) > MaxLandingPageSellStrategyNameBytes {
		return errors.New("strategy_name长度须在1-60字节之间")
	}
	if r.EpisodePrice < MinLandingPageSellStrategyEpisodePrice || r.EpisodePrice > MaxLandingPageSellStrategyEpisodePrice {
		return errors.New("episode_price须在0.01-999999之间")
	}
	if r.MinRechargeTier < MinLandingPageSellStrategyMinRechargeTier || r.MinRechargeTier > MaxLandingPageSellStrategyMinRechargeTier {
		return errors.New("min_recharge_tier须在0.01-999999之间")
	}
	if r.RechargeNum < MinLandingPageSellStrategyRechargeNum || r.RechargeNum > MaxLandingPageSellStrategyRechargeNum {
		return errors.New("recharge_num须在0-999999之间")
	}
	return r.GlobalReq.Validate()
}

// LandingPageSellStrategyAddResp 创建售卖策略响应
type LandingPageSellStrategyAddResp struct {
	AccountId       int64   `json:"account_id"`        // 推广帐号 id
	StrategyId      int64   `json:"strategy_id"`       // 策略 id
	StrategyName    string  `json:"strategy_name"`     // 策略名称
	StrategyType    string  `json:"strategy_type"`     // 策略类型，当前枚举值：SELL
	StrategyStatus  int     `json:"strategy_status"`   // 策略状态，0：无效，1：有效
	EpisodePrice    float64 `json:"episode_price"`     // 单集价格，单位元
	MinRechargeTier float64 `json:"min_recharge_tier"` // 最低充值档位，单位元
	RechargeNum     int     `json:"recharge_num"`      // 起充集数
}
