package model

import "errors"

// ========== 创意组件元素催审 ==========
// https://developers.e.qq.com/v3.0/docs/api/component_element_urge_review/add

// 催审维度枚举
const (
	UrgeDimensionComponent = "URGE_DIMENSION_COMPONENT" // 组件维度催审，urge_dimension_value 为 component_id
	UrgeDimensionElement   = "URGE_DIMENSION_ELEMENT"   // 元素维度催审，urge_dimension_value 为 element_fingerprint
)

// ComponentElementUrgeReviewAddReq 创意组件元素催审请求
type ComponentElementUrgeReviewAddReq struct {
	GlobalReq
	AccountID          int64  `json:"account_id"`           // 推广帐号 id，有操作权限的帐号 id，包括代理商和广告主帐号 id (必填)
	DynamicCreativeID  int64  `json:"dynamic_creative_id"`  // 广告创意 id (必填)
	UrgeDimension      string `json:"urge_dimension"`       // 催审维度 (必填)，可选值：URGE_DIMENSION_COMPONENT, URGE_DIMENSION_ELEMENT
	UrgeDimensionValue string `json:"urge_dimension_value"` // 催审维度对应的值 (必填)，组件维度为 component_id，元素维度为 element_fingerprint，字段长度最大 128 字节
}

func (p *ComponentElementUrgeReviewAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证创意组件元素催审请求参数
func (p *ComponentElementUrgeReviewAddReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.DynamicCreativeID == 0 {
		return errors.New("dynamic_creative_id为必填")
	}
	if p.UrgeDimension == "" {
		return errors.New("urge_dimension为必填")
	}
	if p.UrgeDimension != UrgeDimensionComponent && p.UrgeDimension != UrgeDimensionElement {
		return errors.New("urge_dimension可选值为URGE_DIMENSION_COMPONENT、URGE_DIMENSION_ELEMENT")
	}
	if p.UrgeDimensionValue == "" {
		return errors.New("urge_dimension_value为必填")
	}
	if len(p.UrgeDimensionValue) > 128 {
		return errors.New("urge_dimension_value长度最大128字节")
	}
	return nil
}

// ComponentElementUrgeReviewAddResp 创意组件元素催审响应
type ComponentElementUrgeReviewAddResp struct {
	UrgeResult       string `json:"urge_result,omitempty"`        // 催审结果
	UrgeFailedReason string `json:"urge_failed_reason,omitempty"` // 催审失败原因
	ReasonMsg        string `json:"reason_msg,omitempty"`         // 催审失败结果枚举说明
}
