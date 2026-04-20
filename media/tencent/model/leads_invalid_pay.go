package model

import "errors"

// ========== 获取无效赔付明细 ==========
// https://developers.e.qq.com/v3.0/docs/api/leads_invalid_pay/get

// 月份字段长度
const (
	LeadsInvalidPayMonthLength = 7
)

// LeadsInvalidPayGetReq 获取无效赔付明细请求
type LeadsInvalidPayGetReq struct {
	GlobalReq
	AccountID int64  `json:"account_id"` // 推广帐号id (必填)
	Month     string `json:"month"`      // 月份 (必填)，格式YYYY-MM，7字节
}

func (p *LeadsInvalidPayGetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证获取无效赔付明细请求参数
func (p *LeadsInvalidPayGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	// 验证account_id
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	// 验证month
	if p.Month == "" {
		return errors.New("month为必填")
	}
	if len(p.Month) != LeadsInvalidPayMonthLength {
		return errors.New("month长度必须为7字节，格式YYYY-MM")
	}

	return nil
}

// LeadsInvalidPayGetResp 获取无效赔付明细响应
type LeadsInvalidPayGetResp struct {
	List []*LeadsInvalidPayItem `json:"list,omitempty"` // 确认赔付信息
}

// LeadsInvalidPayItem 赔付信息
type LeadsInvalidPayItem struct {
	AccountID      int64                    `json:"account_id,omitempty"`       // 推广帐号id
	AccountName    string                   `json:"account_name,omitempty"`     // 广告主名称
	IsRealPay      bool                     `json:"is_real_pay,omitempty"`      // 是否产生真实赔付
	NoPayReason    string                   `json:"no_pay_reason,omitempty"`    // 未产生赔付的原因
	PayTotalAmount float64                  `json:"pay_total_amount,omitempty"` // 赔付总金额，单位：元
	PayDetails     []*LeadsInvalidPayDetail `json:"pay_details,omitempty"`      // 确认赔付明细
}

// LeadsInvalidPayDetail 赔付明细
type LeadsInvalidPayDetail struct {
	LeadsId int64 `json:"leads_id,omitempty"` // 线索id
}
