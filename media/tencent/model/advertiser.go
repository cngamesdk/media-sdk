package model

import "errors"

// ========== 批量修改广告主日限额 ==========
// https://developers.e.qq.com/v3.0/docs/api/advertiser/update_daily_budget

// 字段限制常量
const (
	MaxAdvertiserUpdateDailyBudgetSpecCount = 100        // update_daily_budget_spec 最大长度
	MinAdvertiserDailyBudget                = 5000       // 日预算最小值（分），50元
	MaxAdvertiserDailyBudget                = 4000000000 // 日预算最大值（分），40000000元
)

// AdvertiserUpdateDailyBudgetSpec 广告主日限额更新条件
type AdvertiserUpdateDailyBudgetSpec struct {
	AccountID         int64 `json:"account_id"`                     // 广告主帐号 id (必填)
	DailyBudget       int64 `json:"daily_budget"`                   // 广告账户日预算，单位为分 (必填)，0=不限，否则 5000-4000000000
	UseMinDailyBudget bool  `json:"use_min_daily_budget,omitempty"` // 下调失败时是否自动设置为系统允许最小值，默认 false
}

// Validate 验证单个广告主日限额条件
func (s *AdvertiserUpdateDailyBudgetSpec) Validate() error {
	if s.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if s.DailyBudget != 0 && (s.DailyBudget < MinAdvertiserDailyBudget || s.DailyBudget > MaxAdvertiserDailyBudget) {
		return errors.New("daily_budget设置为0表示不限，否则须在5000-4000000000分之间")
	}
	return nil
}

// AdvertiserUpdateDailyBudgetReq 批量修改广告主日限额请求
// https://developers.e.qq.com/v3.0/docs/api/advertiser/update_daily_budget
type AdvertiserUpdateDailyBudgetReq struct {
	GlobalReq
	UpdateDailyBudgetSpec []*AdvertiserUpdateDailyBudgetSpec `json:"update_daily_budget_spec"` // 更新日限额条件列表 (必填)，最大100
}

func (p *AdvertiserUpdateDailyBudgetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证批量修改广告主日限额请求参数
func (p *AdvertiserUpdateDailyBudgetReq) Validate() error {
	if len(p.UpdateDailyBudgetSpec) == 0 {
		return errors.New("update_daily_budget_spec为必填，至少包含1个条件")
	}
	if len(p.UpdateDailyBudgetSpec) > MaxAdvertiserUpdateDailyBudgetSpecCount {
		return errors.New("update_daily_budget_spec数组长度不能超过100")
	}
	seen := make(map[int64]bool)
	for i, spec := range p.UpdateDailyBudgetSpec {
		if spec == nil {
			return errors.New("update_daily_budget_spec[" + itoa(i) + "]不能为空")
		}
		if err := spec.Validate(); err != nil {
			return errors.New("update_daily_budget_spec[" + itoa(i) + "]: " + err.Error())
		}
		if seen[spec.AccountID] {
			return errors.New("update_daily_budget_spec中account_id不允许重复：" + itoa(int(spec.AccountID)))
		}
		seen[spec.AccountID] = true
	}
	return p.GlobalReq.Validate()
}

// AdvertiserUpdateDailyBudgetResultItem 批量修改广告主日限额响应列表项
type AdvertiserUpdateDailyBudgetResultItem struct {
	Code              int    `json:"code"`                 // 返回码
	Message           string `json:"message"`              // 英文返回消息
	MessageCn         string `json:"message_cn"`           // 中文返回消息
	AccountID         int64  `json:"account_id"`           // 广告主帐号 id
	DailyBudget       int64  `json:"daily_budget"`         // 广告账户日预算，单位为分
	UseMinDailyBudget bool   `json:"use_min_daily_budget"` // 是否使用了系统允许最小值
}

// AdvertiserUpdateDailyBudgetResp 批量修改广告主日限额响应
// https://developers.e.qq.com/v3.0/docs/api/advertiser/update_daily_budget
type AdvertiserUpdateDailyBudgetResp struct {
	List       []*AdvertiserUpdateDailyBudgetResultItem `json:"list"`         // 返回信息列表，顺序与请求一致
	FailIDList []int64                                  `json:"fail_id_list"` // 失败的 id 集合
}
