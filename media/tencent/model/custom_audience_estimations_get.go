package model

import "errors"

// ========== 人群覆盖数预估 ==========
// https://developers.e.qq.com/v3.0/docs/api/custom_audience_estimations/get

// CustomAudienceEstimationsGetReq 人群覆盖数预估请求
type CustomAudienceEstimationsGetReq struct {
	GlobalReq
	AccountID    int64                   `json:"account_id"`    // 推广帐号 id (必填)
	Type         string                  `json:"type"`          // 人群类型，目前仅支持 COMBINE (必填)
	AudienceSpec *EstimationAudienceSpec `json:"audience_spec"` // 人群信息，和 type 相关 (必填)
}

// EstimationAudienceSpec 预估接口人群信息
type EstimationAudienceSpec struct {
	CombineSpec *CombineSpec `json:"combine_spec,omitempty"` // 组合人群信息，type=COMBINE 时必填
}

func (p *CustomAudienceEstimationsGetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证人群覆盖数预估请求参数
func (p *CustomAudienceEstimationsGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.Type == "" {
		return errors.New("type为必填")
	}
	if p.Type != "COMBINE" {
		return errors.New("type目前仅支持COMBINE")
	}
	if p.AudienceSpec == nil {
		return errors.New("audience_spec为必填")
	}
	if p.Type == "COMBINE" && p.AudienceSpec.CombineSpec == nil {
		return errors.New("type=COMBINE时，audience_spec.combine_spec为必填")
	}
	if p.AudienceSpec.CombineSpec != nil {
		if len(p.AudienceSpec.CombineSpec.Include) == 0 {
			return errors.New("combine_spec.include为必填")
		}
		if len(p.AudienceSpec.CombineSpec.Include) > 500 {
			return errors.New("combine_spec.include最大长度为500")
		}
		if len(p.AudienceSpec.CombineSpec.Exclude) > 500 {
			return errors.New("combine_spec.exclude最大长度为500")
		}
	}
	return nil
}

// CustomAudienceEstimationsGetResp 人群覆盖数预估响应
type CustomAudienceEstimationsGetResp struct {
	UserCount int64 `json:"user_count"` // 预估出来的人群覆盖数
}
