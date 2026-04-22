package model

import "errors"

// ========== 获取创意组件元素催审状态 ==========
// https://developers.e.qq.com/v3.0/docs/api/component_element_urge_review/get

// ComponentElementUrgeReviewGetReq 获取创意组件元素催审状态请求
type ComponentElementUrgeReviewGetReq struct {
	GlobalReq
	AccountID              int64    `json:"account_id"`                         // 推广帐号 id，有操作权限的帐号 id，包括代理商和广告主帐号 id (必填)
	DynamicCreativeID      int64    `json:"dynamic_creative_id"`                // 广告创意 id (必填)
	ComponentIDList        []int64  `json:"component_id_list,omitempty"`        // 创意组件 id 列表，数组最小长度 1，最大长度 100
	ElementFingerprintList []string `json:"element_fingerprint_list,omitempty"` // 元素指纹列表，数组最小长度 1，最大长度 100，字段长度最小 0 字节，长度最大 128 字节
}

func (p *ComponentElementUrgeReviewGetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证获取创意组件元素催审状态请求参数
func (p *ComponentElementUrgeReviewGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.DynamicCreativeID == 0 {
		return errors.New("dynamic_creative_id为必填")
	}
	if len(p.ComponentIDList) > 100 {
		return errors.New("component_id_list最大长度100")
	}
	if len(p.ElementFingerprintList) > 100 {
		return errors.New("element_fingerprint_list最大长度100")
	}
	for _, fp := range p.ElementFingerprintList {
		if len(fp) > 128 {
			return errors.New("element_fingerprint_list元素长度最大128字节")
		}
	}
	return nil
}

// ComponentElementUrgeReviewGetResp 获取创意组件元素催审状态响应
type ComponentElementUrgeReviewGetResp struct {
	AccountID             int64                `json:"account_id,omitempty"`               // 推广帐号 id
	DynamicCreativeID     int64                `json:"dynamic_creative_id,omitempty"`      // 广告创意 id
	ComponentUrgeInfoList []*ComponentUrgeInfo `json:"component_urge_info_list,omitempty"` // 组件催审状态列表，可为空
	ElementUrgeInfoList   []*ElementUrgeInfo   `json:"element_urge_info_list,omitempty"`   // 元素催审状态列表，可为空
}

// ComponentUrgeInfo 组件催审状态
type ComponentUrgeInfo struct {
	ComponentID int64     `json:"component_id,omitempty"` // 创意组件 id
	UrgeInfo    *UrgeInfo `json:"urge_info,omitempty"`    // 催审信息
}

// ElementUrgeInfo 元素催审状态
type ElementUrgeInfo struct {
	ElementFingerprint string    `json:"element_fingerprint,omitempty"` // 元素指纹
	UrgeInfo           *UrgeInfo `json:"urge_info,omitempty"`           // 催审信息
}

// UrgeInfo 催审信息
type UrgeInfo struct {
	UrgeStatusType string `json:"urge_status_type,omitempty"` // 催审状态
	NonUrgeReason  string `json:"non_urge_reason,omitempty"`  // 不可催审原因
	ReasonMsg      string `json:"reason_msg,omitempty"`       // 催审失败结果枚举说明
}
