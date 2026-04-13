package model

import "errors"

// ========== 官方落地页-获取落地页列表 ==========
// https://developers.e.qq.com/v3.0/docs/api/official_landing_page_list/get

// 官方落地页状态枚举
const (
	OfficialLandingPageStatusEditing  = "LANDING_PAGE_STATUS_EDITING"  // 编辑中
	OfficialLandingPageStatusPending  = "LANDING_PAGE_STATUS_PENDING"  // 审核中
	OfficialLandingPageStatusApproved = "LANDING_PAGE_STATUS_APPROVED" // 审核通过
	OfficialLandingPageStatusRejected = "LANDING_PAGE_STATUS_REJECTED" // 审核拒绝
	OfficialLandingPageStatusOffline  = "LANDING_PAGE_STATUS_OFFLINE"  // 已下线
)

// 过滤字段枚举
const (
	OfficialLandingPageFilterFieldPageId     = "page_id"     // 按落地页 id 过滤
	OfficialLandingPageFilterFieldPageName   = "page_name"   // 按落地页名称过滤
	OfficialLandingPageFilterFieldPageStatus = "page_status" // 按落地页状态过滤
)

// 分页常量
const (
	MinOfficialLandingPageListPage         = 1     // page 最小值
	MaxOfficialLandingPageListPage         = 99999 // page 最大值
	MinOfficialLandingPageListPageSize     = 1     // page_size 最小值
	MaxOfficialLandingPageListPageSize     = 100   // page_size 最大值
	DefaultOfficialLandingPageListPage     = 1     // page 默认值
	DefaultOfficialLandingPageListPageSize = 10    // page_size 默认值
	MaxOfficialLandingPageListFiltering    = 10    // filtering 最大数量
)

// OfficialLandingPageListFilter 官方落地页列表过滤条件
type OfficialLandingPageListFilter struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)，可选值：page_id, page_name, page_status
	Operator string   `json:"operator"` // 操作符 (必填)，可选值：EQUALS
	Values   []string `json:"values"`   // 字段取值 (必填)
}

// OfficialLandingPageListGetReq 官方落地页获取列表请求（GET）
// https://developers.e.qq.com/v3.0/docs/api/official_landing_page_list/get
type OfficialLandingPageListGetReq struct {
	GlobalReq
	AccountId int64                            `json:"account_id"`          // 广告主帐号 id (必填)
	Page      int                              `json:"page,omitempty"`      // 搜索页码，1-99999，默认 1
	PageSize  int                              `json:"page_size,omitempty"` // 每页条数，1-100，默认 10
	Filtering []*OfficialLandingPageListFilter `json:"filtering,omitempty"` // 过滤条件，最多 10 个
}

func (r *OfficialLandingPageListGetReq) Format() {
	r.GlobalReq.Format()
	if r.Page == 0 {
		r.Page = DefaultOfficialLandingPageListPage
	}
	if r.PageSize == 0 {
		r.PageSize = DefaultOfficialLandingPageListPageSize
	}
}

// Validate 验证获取官方落地页列表请求参数
func (r *OfficialLandingPageListGetReq) Validate() error {
	if r.AccountId == 0 {
		return errors.New("account_id为必填")
	}
	if r.Page < MinOfficialLandingPageListPage || r.Page > MaxOfficialLandingPageListPage {
		return errors.New("page须在1-99999之间")
	}
	if r.PageSize < MinOfficialLandingPageListPageSize || r.PageSize > MaxOfficialLandingPageListPageSize {
		return errors.New("page_size须在1-100之间")
	}
	if len(r.Filtering) > MaxOfficialLandingPageListFiltering {
		return errors.New("filtering数组长度不能超过10")
	}
	for i, f := range r.Filtering {
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
	}
	return r.GlobalReq.Validate()
}

// OfficialLandingPageListItem 官方落地页列表项
type OfficialLandingPageListItem struct {
	PageId        int64  `json:"page_id"`         // 落地页服务 id，用于投放端选择落地页
	LandingPageId int    `json:"landing_page_id"` // 官方落地页 id
	PageName      string `json:"page_name"`       // 落地页名称
	PageTitle     string `json:"page_title"`      // 落地页标题
	PageType      string `json:"page_type"`       // 官方落地页类型（枚举）
	PageStatus    string `json:"page_status"`     // 官方落地页状态（枚举）
}

// OfficialLandingPageListRespData 获取落地页列表响应数据
type OfficialLandingPageListRespData struct {
	List     []*OfficialLandingPageListItem `json:"list"`      // 官方落地页列表数据
	PageInfo *PageInfo                      `json:"page_info"` // 分页信息
}

// OfficialLandingPageListGetResp 获取落地页列表响应
type OfficialLandingPageListGetResp struct {
	Code      int                              `json:"code"`
	Message   string                           `json:"message"`
	MessageCn string                           `json:"message_cn"`
	Data      *OfficialLandingPageListRespData `json:"data"`
}
