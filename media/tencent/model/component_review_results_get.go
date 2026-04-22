package model

import "errors"

// ========== 查询组件审核结果 ==========
// https://developers.e.qq.com/v3.0/docs/api/component_review_results/get

// ComponentReviewResultsGetReq 查询组件审核结果请求
type ComponentReviewResultsGetReq struct {
	GlobalReq
	AccountID       int64   `json:"account_id"`        // 广告主帐号 id，有操作权限的帐号 id，不支持代理商 id (必填)
	ComponentIDList []int64 `json:"component_id_list"` // 创意组件 id 列表，数组最小长度 1，最大长度 100 (必填)
}

func (p *ComponentReviewResultsGetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证查询组件审核结果请求参数
func (p *ComponentReviewResultsGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if len(p.ComponentIDList) == 0 {
		return errors.New("component_id_list为必填，最小长度1")
	}
	if len(p.ComponentIDList) > 100 {
		return errors.New("component_id_list最大长度100")
	}
	return nil
}

// ComponentReviewResultsGetResp 查询组件审核结果响应
type ComponentReviewResultsGetResp struct {
	List []*ComponentReviewResultItem `json:"list,omitempty"` // 返回信息列表
}

// ComponentReviewResultItem 组件审核结果
type ComponentReviewResultItem struct {
	ComponentInfo     *ReviewComponentInfo            `json:"component_info,omitempty"`      // 创意组件信息
	ElementResultList []*ComponentReviewElementResult `json:"element_result_list,omitempty"` // 组件元素粒度审核结果列表
}

// ComponentReviewElementResult 组件元素粒度审核结果
type ComponentReviewElementResult struct {
	ElementID               int64                  `json:"element_id,omitempty"`                 // 元素 id
	ImageID                 string                 `json:"image_id,omitempty"`                   // 图片 id
	VideoID                 string                 `json:"video_id,omitempty"`                   // 视频 id
	ElementName             string                 `json:"element_name,omitempty"`               // 元素名称
	ElementValue            string                 `json:"element_value,omitempty"`              // 元素值
	ElementType             string                 `json:"element_type,omitempty"`               // 元素类型
	ReviewStatus            string                 `json:"review_status,omitempty"`              // 审核结果状态
	ElementRejectDetailInfo []*ElementRejectDetail `json:"element_reject_detail_info,omitempty"` // 元素拒绝原因详情列表
}
