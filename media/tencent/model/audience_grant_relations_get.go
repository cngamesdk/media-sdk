package model

import "errors"

// ========== 获取人群授权信息 ==========
// https://developers.e.qq.com/v3.0/docs/api/audience_grant_relations/get

// 过滤字段常量
const (
	AudienceGrantRelationsFilterFieldAudienceID = "audience_id"
)

// AudienceGrantRelationsGetReq 获取人群授权信息请求
type AudienceGrantRelationsGetReq struct {
	GlobalReq
	AccountID int64                              `json:"account_id"`          // 推广帐号 id (必填)
	Filtering []*AudienceGrantRelationsFiltering `json:"filtering,omitempty"` // 过滤条件，不传则无限制
	Page      int                                `json:"page,omitempty"`      // 当前页码，最小值 1，默认值 1
	PageSize  int                                `json:"page_size,omitempty"` // 分页大小，最小值 1，最大值 100，默认值 10
}

// AudienceGrantRelationsFiltering 过滤条件
type AudienceGrantRelationsFiltering struct {
	Field    string   `json:"field"`    // 过滤字段，可选值：audience_id (必填)
	Operator string   `json:"operator"` // 操作符，field=audience_id 时可选值：IN (必填)
	Values   []string `json:"values"`   // 字段取值，field=audience_id 时数组长度 1-100 (必填)
}

func (p *AudienceGrantRelationsGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
}

// Validate 验证获取人群授权信息请求参数
func (p *AudienceGrantRelationsGetReq) Validate() error {
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
		if f.Operator == "" {
			return errors.New("filtering[" + itoa(i) + "].operator为必填")
		}
		if len(f.Values) == 0 {
			return errors.New("filtering[" + itoa(i) + "].values为必填")
		}
		if f.Field == AudienceGrantRelationsFilterFieldAudienceID && len(f.Values) > 100 {
			return errors.New("filtering[" + itoa(i) + "].values当field=audience_id时数组长度最大为100")
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

// AudienceGrantRelationsGetResp 获取人群授权信息响应
type AudienceGrantRelationsGetResp struct {
	List     []*AudienceGrantRelationItem `json:"list,omitempty"`      // 返回数组列表
	PageInfo *PageInfo                    `json:"page_info,omitempty"` // 分页信息
}

// AudienceGrantRelationItem 人群授权信息条目
type AudienceGrantRelationItem struct {
	AudienceID int64                      `json:"audience_id,omitempty"` // 自定义人群 id
	GrantType  string                     `json:"grant_type,omitempty"`  // 授权类型
	GrantSpec  *AudienceGrantRelationSpec `json:"grant_spec,omitempty"`  // 人群授权信息
}

// AudienceGrantRelationSpec 人群授权信息（响应结构，复用 add 中定义的子结构）
type AudienceGrantRelationSpec struct {
	GrantToBusinessSpec *GrantToBusinessSpec `json:"grant_to_business_spec,omitempty"` // 授权给商务管家账号的授权信息
}
