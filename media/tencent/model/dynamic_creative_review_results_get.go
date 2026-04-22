package model

import "errors"

// ========== 查询动态创意审核结果 ==========
// https://developers.e.qq.com/v3.0/docs/api/dynamic_creative_review_results/get

// DynamicCreativeReviewResultsGetReq 查询动态创意审核结果请求
type DynamicCreativeReviewResultsGetReq struct {
	GlobalReq
	AccountID             int64   `json:"account_id"`               // 广告主帐号 id，有操作权限的帐号 id，不支持代理商 id (必填)
	DynamicCreativeIDList []int64 `json:"dynamic_creative_id_list"` // 创意 id 列表，数组最小长度 1，最大长度 100 (必填)
}

func (p *DynamicCreativeReviewResultsGetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证查询动态创意审核结果请求参数
func (p *DynamicCreativeReviewResultsGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if len(p.DynamicCreativeIDList) == 0 {
		return errors.New("dynamic_creative_id_list为必填，最小长度1")
	}
	if len(p.DynamicCreativeIDList) > 100 {
		return errors.New("dynamic_creative_id_list最大长度100")
	}
	return nil
}

// DynamicCreativeReviewResultsGetResp 查询动态创意审核结果响应
type DynamicCreativeReviewResultsGetResp struct {
	List []*DynamicCreativeReviewResultItem `json:"list,omitempty"` // 返回信息列表
}

// DynamicCreativeReviewResultItem 创意审核结果
type DynamicCreativeReviewResultItem struct {
	DynamicCreativeID              int64                         `json:"dynamic_creative_id,omitempty"`                // 广告创意 id
	ElementResultList              []*ElementResult              `json:"element_result_list,omitempty"`                // 元素粒度审核结果
	SiteSetResultList              []*SiteSetResult              `json:"site_set_result_list,omitempty"`               // 版位粒度审核结果
	RejectMessageList              []string                      `json:"reject_message_list,omitempty"`                // 拒绝原因列表
	DelayMessageList               []string                      `json:"delay_message_list,omitempty"`                 // 被延迟审核的信息
	IsAllComponentComposePending   bool                          `json:"is_all_component_compose_pending,omitempty"`   // 是否所有组件组合均为待审核
	TotalComponentComposeCount     int64                         `json:"total_component_compose_count,omitempty"`      // 总组件组合数量
	RejectComponentComposeCount    int64                         `json:"reject_component_compose_count,omitempty"`     // 审核拒绝组件组合数量
	PassComponentComposeCount      int64                         `json:"pass_component_compose_count,omitempty"`       // 审核通过组件组合数量
	RejectComponentComposeInfoList []*RejectComponentComposeInfo `json:"reject_component_compose_info_list,omitempty"` // 审核拒绝的组件组合元素明细
}

// ElementResult 元素粒度审核结果
type ElementResult struct {
	ElementID               int64                  `json:"element_id,omitempty"`                 // 元素 id
	ImageID                 string                 `json:"image_id,omitempty"`                   // 图片 id
	VideoID                 string                 `json:"video_id,omitempty"`                   // 视频 id
	ElementName             string                 `json:"element_name,omitempty"`               // 元素名称
	ElementValue            string                 `json:"element_value,omitempty"`              // 元素值
	ElementFingerprint      string                 `json:"element_fingerprint,omitempty"`        // 元素指纹
	ComponentInfo           *ReviewComponentInfo   `json:"component_info,omitempty"`             // 元素所属创意组件信息，可能为空
	ElementType             string                 `json:"element_type,omitempty"`               // 元素类型
	ReviewStatus            string                 `json:"review_status,omitempty"`              // 审核结果状态
	ElementRejectDetailInfo []*ElementRejectDetail `json:"element_reject_detail_info,omitempty"` // 元素拒绝原因详情列表
}

// ReviewComponentInfo 元素所属创意组件信息
type ReviewComponentInfo struct {
	ComponentID   int64  `json:"component_id,omitempty"`   // 创意组件 id
	ComponentType string `json:"component_type,omitempty"` // 创意组件类型
	ReviewStatus  string `json:"review_status,omitempty"`  // 审核结果状态
}

// ElementRejectDetail 元素拒绝原因详情
type ElementRejectDetail struct {
	Reason             string                `json:"reason,omitempty"`               // 拒绝原因
	SiteSetList        []*ReviewSiteSetItem  `json:"site_set_list,omitempty"`        // 影响版位列表
	RejectInfoLocation []*RejectInfoLocation `json:"reject_info_location,omitempty"` // 拒绝原因可视化信息列表
}

// ReviewSiteSetItem 版位信息
type ReviewSiteSetItem struct {
	SiteSet string `json:"site_set,omitempty"` // 版位
}

// RejectInfoLocation 拒绝原因可视化信息
type RejectInfoLocation struct {
	X              int64   `json:"x,omitempty"`                // x 轴位置
	Y              int64   `json:"y,omitempty"`                // y 轴位置
	Height         int64   `json:"height,omitempty"`           // 高度
	Width          int64   `json:"width,omitempty"`            // 宽度
	ImgUrl         string  `json:"img_url,omitempty"`          // 帧图片 url
	LocationImgUrl string  `json:"location_img_url,omitempty"` // 标注结果图 url
	RelatedImgUrl  string  `json:"related_img_url,omitempty"`  // 种子图 url
	TimeSecond     float64 `json:"time_second,omitempty"`      // 时间戳，视频时间点（非视频元素本字段为空）
}

// SiteSetResult 版位粒度审核结果
type SiteSetResult struct {
	SiteSet                 string                        `json:"site_set,omitempty"`                   // 版位
	SystemStatus            string                        `json:"system_status,omitempty"`              // 审核结果状态
	RejectMessage           string                        `json:"reject_message,omitempty"`             // 拒绝原因
	ElementRejectDetailInfo []*SiteSetElementRejectDetail `json:"element_reject_detail_info,omitempty"` // 元素审核结果列表
}

// SiteSetElementRejectDetail 版位粒度下的元素审核结果
type SiteSetElementRejectDetail struct {
	ElementName        string                `json:"element_name,omitempty"`         // 元素名称
	ElementType        string                `json:"element_type,omitempty"`         // 元素类型
	ElementValue       string                `json:"element_value,omitempty"`        // 元素值
	ComponentInfo      *ReviewComponentInfo  `json:"component_info,omitempty"`       // 元素所属创意组件信息，可能为空
	Reason             string                `json:"reason,omitempty"`               // 拒绝原因
	ReviewStatus       string                `json:"review_status,omitempty"`        // 审核结果状态
	RejectInfoLocation []*RejectInfoLocation `json:"reject_info_location,omitempty"` // 拒绝原因可视化信息列表
}

// RejectComponentComposeInfo 审核拒绝的组件组合元素明细
type RejectComponentComposeInfo struct {
	RejectMessage               string                     `json:"reject_message,omitempty"`                 // 拒绝原因
	ComponentComposeElementList []*ComponentComposeElement `json:"component_compose_element_list,omitempty"` // 组件组合的元素明细
}

// ComponentComposeElement 组件组合的元素明细
type ComponentComposeElement struct {
	ElementResultList []*ElementResult `json:"element_result_list,omitempty"` // 元素粒度审核结果
}
