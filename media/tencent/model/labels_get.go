package model

import "errors"

// ========== 标签广场标签获取 ==========
// https://developers.e.qq.com/v3.0/docs/api/labels/get

// 过滤字段常量
const (
	LabelsFilterFieldLabelGroup       = "label_group"        // 标签分组，可选值：MAP（标签地图）、POP（热点人群），操作符：IN
	LabelsFilterFieldParentID         = "parent_id"          // 父级标签 id，操作符：NOT_IN
	LabelsFilterFieldDisplayLabelName = "display_label_name" // 标签名称，操作符：CONTAINS
)

// 标签分组枚举
const (
	LabelsGroupMAP = "MAP" // 标签地图
	LabelsGroupPOP = "POP" // 热点人群
)

// LabelsGetReq 标签广场标签获取请求
type LabelsGetReq struct {
	GlobalReq
	AccountID int64              `json:"account_id"`          // 推广帐号 id (必填)
	Filtering []*LabelsFiltering `json:"filtering,omitempty"` // 过滤条件，不传则无限制
	Page      int                `json:"page,omitempty"`      // 当前页码，最小值 1，默认值 1
	PageSize  int                `json:"page_size,omitempty"` // 分页大小，最小值 1，最大值 100，默认值 10
}

// LabelsFiltering 过滤条件
type LabelsFiltering struct {
	Field    string   `json:"field"`    // 过滤字段，可选值：label_group、parent_id、display_label_name (必填)
	Operator string   `json:"operator"` // 操作符 (必填)
	Values   []string `json:"values"`   // 字段取值 (必填)
}

func (p *LabelsGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
}

// Validate 验证标签广场标签获取请求参数
func (p *LabelsGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}
	for i, f := range p.Filtering {
		if f == nil {
			return errors.New("filtering[" + itoa(i) + "]不能为空")
		}
		if f.Field == "" {
			return errors.New("filtering[" + itoa(i) + "].field为必填")
		}
		validFields := map[string]bool{
			LabelsFilterFieldLabelGroup:       true,
			LabelsFilterFieldParentID:         true,
			LabelsFilterFieldDisplayLabelName: true,
		}
		if !validFields[f.Field] {
			return errors.New("filtering[" + itoa(i) + "].field值无效，可选值：label_group、parent_id、display_label_name")
		}
		if f.Operator == "" {
			return errors.New("filtering[" + itoa(i) + "].operator为必填")
		}
		validOperators := map[string]string{
			LabelsFilterFieldLabelGroup:       "IN",
			LabelsFilterFieldParentID:         "NOT_IN",
			LabelsFilterFieldDisplayLabelName: "CONTAINS",
		}
		if f.Operator != validOperators[f.Field] {
			return errors.New("filtering[" + itoa(i) + "].operator与field不匹配，field=" + f.Field + " 时 operator 须为 " + validOperators[f.Field])
		}
		if len(f.Values) == 0 {
			return errors.New("filtering[" + itoa(i) + "].values为必填")
		}
		if f.Field == LabelsFilterFieldLabelGroup {
			for _, v := range f.Values {
				if v != LabelsGroupMAP && v != LabelsGroupPOP {
					return errors.New("filtering[" + itoa(i) + "].values当field=label_group时可选值为MAP、POP")
				}
			}
		}
	}
	if p.Page < 1 {
		return errors.New("page最小值为1")
	}
	if p.PageSize < 1 || p.PageSize > 100 {
		return errors.New("page_size必须在1-100之间")
	}
	return nil
}

// LabelsGetResp 标签广场标签获取响应
type LabelsGetResp struct {
	List     []*LabelsItem `json:"list,omitempty"`      // 返回数组列表
	PageInfo *PageInfo     `json:"page_info,omitempty"` // 分页信息
}

// LabelsItem 标签广场标签信息
type LabelsItem struct {
	LabelID          int64  `json:"label_id,omitempty"`           // 标签 id
	AudienceID       int64  `json:"audience_id,omitempty"`        // 自定义人群 id
	DisplayLabelName string `json:"display_label_name,omitempty"` // 标签名称
	Description      string `json:"description,omitempty"`        // 标签描述
	UserCount        int64  `json:"user_count,omitempty"`         // 用户覆盖数
}
