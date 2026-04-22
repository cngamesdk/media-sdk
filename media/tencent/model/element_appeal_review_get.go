package model

import "errors"

// ========== 获取元素申诉复审结果 ==========
// https://developers.e.qq.com/v3.0/docs/api/element_appeal_review/get

// ElementAppealReviewGetReq 获取元素申诉复审结果请求
type ElementAppealReviewGetReq struct {
	GlobalReq
	AccountID          int64  `json:"account_id"`           // 广告主帐号 id，有操作权限的帐号 id，不支持代理商 id (必填)
	DynamicCreativeID  int64  `json:"dynamic_creative_id"`  // 广告创意 id (必填)
	ComponentID        int64  `json:"component_id"`         // 创意组件 id (必填)
	ElementID          int64  `json:"element_id"`           // 元素 id (必填)
	ElementFingerPrint string `json:"element_finger_print"` // 元素指纹，字段长度最小 0 字节，长度最大 128 字节 (必填)
}

func (p *ElementAppealReviewGetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证获取元素申诉复审结果请求参数
func (p *ElementAppealReviewGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if p.DynamicCreativeID == 0 {
		return errors.New("dynamic_creative_id为必填")
	}
	if p.ComponentID == 0 {
		return errors.New("component_id为必填")
	}
	if p.ElementID == 0 {
		return errors.New("element_id为必填")
	}
	if len(p.ElementFingerPrint) > 128 {
		return errors.New("element_finger_print长度最大128字节")
	}
	return nil
}

// ElementAppealReviewGetResp 获取元素申诉复审结果响应
type ElementAppealReviewGetResp struct {
	AccountID                  int64    `json:"account_id,omitempty"`                    // 推广帐号 id
	DynamicCreativeID          int64    `json:"dynamic_creative_id,omitempty"`           // 广告创意 id
	ComponentID                int64    `json:"component_id,omitempty"`                  // 创意组件 id
	ElementID                  int64    `json:"element_id,omitempty"`                    // 元素 id
	ElementType                string   `json:"element_type,omitempty"`                  // 元素类型
	ElementValue               string   `json:"element_value,omitempty"`                 // 元素值
	ElementFingerPrint         string   `json:"element_finger_print,omitempty"`          // 元素指纹
	AppealDemand               string   `json:"appeal_demand,omitempty"`                 // 申诉复审需求，支持多选，多个申诉需求时以英文分号分割
	AppealReason               string   `json:"appeal_reason,omitempty"`                 // 申诉原因
	HistoryApprovalComponentID int64    `json:"history_approval_component_id,omitempty"` // 历史已通过组件 id
	AppealResult               string   `json:"appeal_result,omitempty"`                 // 申诉结果
	AppealStatus               string   `json:"appeal_status,omitempty"`                 // 申诉状态
	Description                string   `json:"description,omitempty"`                   // 详细描述
	ImageList                  []string `json:"image_list,omitempty"`                    // 附件图片列表
	ReplyImageUrlList          []string `json:"reply_image_url_list,omitempty"`          // 申诉复审回复的图片地址列表
}
