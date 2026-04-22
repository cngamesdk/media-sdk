package model

import "errors"

// ========== 查询元素的预审结果 ==========
// https://developers.e.qq.com/v3.0/docs/api/review_element_prereview_results/get

// 预审元素类型枚举
const (
	PreReviewElementTypeImage   = "IMAGE"    // 图片
	PreReviewElementTypeVideo   = "VIDEO"    // 视频
	PreReviewElementTypeTxt     = "TXT"      // 文本
	PreReviewElementTypeDestUrl = "DEST_URL" // 落地页
)

// ReviewElementPrereviewResultsGetReq 查询元素的预审结果请求
type ReviewElementPrereviewResultsGetReq struct {
	GlobalReq
	AccountID  int64                  `json:"account_id"`           // 推广帐号 id，有操作权限的帐号 id，包括代理商和广告主帐号 id (必填)
	AdgroupID  int64                  `json:"adgroup_id,omitempty"` // 广告 id
	Elements   []*PreReviewElement    `json:"elements"`             // 元素信息的列表，数组最小长度 1，最大长度 20 (必填)
	Supplement []*PreReviewSupplement `json:"supplement,omitempty"` // 补充信息，不传或传空则视为无补充信息，数组最小长度 1，最大长度 50
}

// PreReviewElement 预审元素信息
type PreReviewElement struct {
	ElementType    string `json:"element_type"`          // 元素类型 (必填)，可选值：IMAGE, VIDEO, TXT, DEST_URL
	ElementContent string `json:"element_content"`       // 元素内容 (必填)，字段长度最大 256 字节
	ElementKey     string `json:"element_key,omitempty"` // 元素 key
}

// PreReviewSupplement 预审补充信息
type PreReviewSupplement struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)
	Operator string   `json:"operator"` // 操作符 (必填)
	Values   []string `json:"values"`   // 字段取值 (必填)，数组最小长度 1，最大长度 100
}

func (p *ReviewElementPrereviewResultsGetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证查询元素的预审结果请求参数
func (p *ReviewElementPrereviewResultsGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	if len(p.Elements) == 0 {
		return errors.New("elements为必填，最小长度1")
	}
	if len(p.Elements) > 20 {
		return errors.New("elements最大长度20")
	}
	for _, e := range p.Elements {
		if e.ElementType == "" {
			return errors.New("elements.element_type为必填")
		}
		if e.ElementContent == "" {
			return errors.New("elements.element_content为必填")
		}
		if len(e.ElementContent) > 256 {
			return errors.New("elements.element_content长度最大256字节")
		}
	}
	if len(p.Supplement) > 50 {
		return errors.New("supplement最大长度50")
	}
	for _, s := range p.Supplement {
		if s.Field == "" {
			return errors.New("supplement.field为必填")
		}
		if s.Operator == "" {
			return errors.New("supplement.operator为必填")
		}
		if len(s.Values) == 0 {
			return errors.New("supplement.values为必填，最小长度1")
		}
		if len(s.Values) > 100 {
			return errors.New("supplement.values最大长度100")
		}
	}
	return nil
}

// ReviewElementPrereviewResultsGetResp 查询元素的预审结果响应
type ReviewElementPrereviewResultsGetResp struct {
	List []*PreReviewResultItem `json:"list,omitempty"` // 预审结果实体集合
}

// PreReviewResultItem 预审结果实体
type PreReviewResultItem struct {
	ElementType      string             `json:"element_type,omitempty"`       // 元素类型
	ElementContent   string             `json:"element_content,omitempty"`    // 元素内容
	RiskLevel        string             `json:"risk_level,omitempty"`         // 元素风险级别
	PreReviewDetails []*PreReviewDetail `json:"pre_review_details,omitempty"` // 元素预审结果明细集合
}

// PreReviewDetail 元素预审结果明细
type PreReviewDetail struct {
	SiteSet            string                         `json:"site_set,omitempty"`             // 投放版位集合
	PreReviewResult    string                         `json:"pre_review_result,omitempty"`    // 投放版位的预审结果
	RejectReasonDetail []*PreReviewRejectReasonDetail `json:"reject_reason_detail,omitempty"` // 拒绝原因明细集合
}

// PreReviewRejectReasonDetail 拒绝原因明细
type PreReviewRejectReasonDetail struct {
	RejectReasonID      string                         `json:"reject_reason_id,omitempty"`      // 拒绝原因 id
	RejectReasonContent string                         `json:"reject_reason_content,omitempty"` // 拒绝原因内容
	CaseDoc             string                         `json:"case_doc,omitempty"`              // 案例 url
	CaseContent         string                         `json:"case_content,omitempty"`          // 案例内容富文本
	RejectInfoLocations []*PreReviewRejectInfoLocation `json:"reject_info_locations,omitempty"` // 元素拒绝原因详情
}

// PreReviewRejectInfoLocation 预审拒绝原因可视化信息
type PreReviewRejectInfoLocation struct {
	X              int64   `json:"x,omitempty"`                // x 轴位置
	Y              int64   `json:"y,omitempty"`                // y 轴位置
	Width          int64   `json:"width,omitempty"`            // 宽度
	Height         int64   `json:"height,omitempty"`           // 高度
	TimeSecond     float64 `json:"time_second,omitempty"`      // 视频时间点
	LocationImgUrl string  `json:"location_img_url,omitempty"` // 标注结果图 url
	ImgUrl         string  `json:"img_url,omitempty"`          // 帧图片 url
	RelatedImgUrl  string  `json:"related_img_url,omitempty"`  // 种子图 url
}
