package model

import "errors"

// ========== 发起元素申诉复审 ==========
// https://developers.e.qq.com/v3.0/docs/api/element_appeal_review/add

// 元素类型枚举
const (
	ElementTypeText       = "ELEMENT_TYPE_TEXT"        // 文本
	ElementTypeUrl        = "ELEMENT_TYPE_URL"         // URL
	ElementTypeImage      = "ELEMENT_TYPE_IMAGE"       // 图片
	ElementTypeVideo      = "ELEMENT_TYPE_VIDEO"       // 视频
	ElementTypeSelectNone = "ELEMENT_TYPE_SELECT_NONE" // 无
)

// ElementAppealReviewAddReq 发起元素申诉复审请求
type ElementAppealReviewAddReq struct {
	GlobalReq
	AccountID                  int64    `json:"account_id"`                              // 广告主帐号 id，有操作权限的帐号 id，不支持代理商 id (必填)
	DynamicCreativeID          int64    `json:"dynamic_creative_id"`                     // 广告创意 id (必填)
	ComponentID                int64    `json:"component_id"`                            // 创意组件 id (必填)
	ElementID                  int64    `json:"element_id"`                              // 元素 id (必填)
	ElementType                string   `json:"element_type"`                            // 元素类型 (必填)，可选值：ELEMENT_TYPE_TEXT, ELEMENT_TYPE_URL, ELEMENT_TYPE_IMAGE, ELEMENT_TYPE_VIDEO, ELEMENT_TYPE_SELECT_NONE
	ElementValue               string   `json:"element_value"`                           // 元素值，字段长度最小 0 字节，长度最大 512 字节 (必填)
	ElementFingerPrint         string   `json:"element_finger_print"`                    // 元素指纹，字段长度最小 0 字节，长度最大 128 字节 (必填)
	AppealDemand               string   `json:"appeal_demand"`                           // 申诉复审需求，支持多选，多个申诉需求时以英文分号分割，字段长度最小 1 字节，长度最大 512 字节 (必填)
	AppealReason               string   `json:"appeal_reason"`                           // 申诉原因，字段长度最小 1 字节，长度最大 512 字节 (必填)
	HistoryApprovalComponentID int64    `json:"history_approval_component_id,omitempty"` // 历史已通过组件 id
	Description                string   `json:"description,omitempty"`                   // 详细描述，字段长度最小 1 字节，长度最大 50 字节
	ImageList                  []string `json:"image_list,omitempty"`                    // 附件图片列表，数组最小长度 0，最大长度 3，字段长度最小 1 字节，长度最大 512 字节
}

func (p *ElementAppealReviewAddReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证发起元素申诉复审请求参数
func (p *ElementAppealReviewAddReq) Validate() error {
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
	if p.ElementType == "" {
		return errors.New("element_type为必填")
	}
	if p.ElementValue == "" {
		return errors.New("element_value为必填")
	}
	if len(p.ElementValue) > 512 {
		return errors.New("element_value长度最大512字节")
	}
	if len(p.ElementFingerPrint) > 128 {
		return errors.New("element_finger_print长度最大128字节")
	}
	if p.AppealDemand == "" {
		return errors.New("appeal_demand为必填")
	}
	if len(p.AppealDemand) > 512 {
		return errors.New("appeal_demand长度最大512字节")
	}
	if p.AppealReason == "" {
		return errors.New("appeal_reason为必填")
	}
	if len(p.AppealReason) > 512 {
		return errors.New("appeal_reason长度最大512字节")
	}
	if p.Description != "" && len(p.Description) > 50 {
		return errors.New("description长度最大50字节")
	}
	if len(p.ImageList) > 3 {
		return errors.New("image_list最大长度3")
	}
	for _, img := range p.ImageList {
		if img == "" || len(img) > 512 {
			return errors.New("image_list元素长度必须在1-512字节之间")
		}
	}
	return nil
}

// ElementAppealReviewAddResp 发起元素申诉复审响应（应答字段为无）
type ElementAppealReviewAddResp struct{}
