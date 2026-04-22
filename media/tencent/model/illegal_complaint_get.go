package model

import "errors"

// ========== 获取直客广告主违规申述列表 ==========
// https://developers.e.qq.com/v3.0/docs/api/illegal_complaint/get

// IllegalComplaintGetReq 获取直客广告主违规申述列表请求
type IllegalComplaintGetReq struct {
	GlobalReq
	AccountIDList    []int64                    `json:"account_id_list"`              // 广告主账号列表，数组最小长度 1，最大长度 700 (必填)
	IllegalLevelList []interface{}              `json:"illegal_level_list,omitempty"` // 违规等级列表，数组最小长度 1，最大长度 100
	ActionTypeList   []interface{}              `json:"action_type_list,omitempty"`   // 处罚动作列表，数组最小长度 1，最大长度 100
	IllegalReason    string                     `json:"illegal_reason,omitempty"`     // 处罚原因
	IllegalDateRange *IllegalComplaintDateRange `json:"illegal_date_range,omitempty"` // 违规时间范围
	Page             int                        `json:"page"`                         // 搜索页码，最小值 1，最大值 99999，默认值 1 (必填)
	PageSize         int                        `json:"page_size"`                    // 一页显示的数据条数，最小值 1，最大值 100，默认值 10 (必填)
}

// IllegalComplaintDateRange 违规时间范围
type IllegalComplaintDateRange struct {
	StartDate string `json:"start_date,omitempty"` // 开始日期，格式 YYYY-MM-DD，小于等于 end_date
	EndDate   string `json:"end_date,omitempty"`   // 结束日期，格式 YYYY-MM-DD，大于等于 start_date
}

func (p *IllegalComplaintGetReq) Format() {
	p.GlobalReq.Format()
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
}

// Validate 验证获取直客广告主违规申述列表请求参数
func (p *IllegalComplaintGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}
	if len(p.AccountIDList) == 0 {
		return errors.New("account_id_list为必填，最小长度1")
	}
	if len(p.AccountIDList) > 700 {
		return errors.New("account_id_list最大长度700")
	}
	if p.Page < 1 {
		return errors.New("page最小值为1")
	}
	if p.PageSize < 1 || p.PageSize > 100 {
		return errors.New("page_size必须在1-100之间")
	}
	return nil
}

// IllegalComplaintGetResp 获取直客广告主违规申述列表响应
type IllegalComplaintGetResp struct {
	PageInfo *IllegalComplaintPageInfo `json:"page_info,omitempty"` // 分页配置信息
	List     []*IllegalComplaintItem   `json:"list,omitempty"`      // 返回广告主违规申述信息列表
}

// IllegalComplaintPageInfo 分页配置信息
type IllegalComplaintPageInfo struct {
	Page        int `json:"page,omitempty"`         // 搜索页码，默认值 1
	PageSize    int `json:"page_size,omitempty"`    // 一页显示的数据条数，默认值 10
	TotalNumber int `json:"total_number,omitempty"` // 总条数
	TotalPage   int `json:"total_page,omitempty"`   // 总页数
}

// IllegalComplaintItem 违规申述信息
type IllegalComplaintItem struct {
	IllegalOrderID      string   `json:"illegal_order_id,omitempty"`      // 违规单 id
	AdvertiserAccountID int64    `json:"advertiser_account_id,omitempty"` // 推广帐号 id
	AdvertiserName      string   `json:"advertiser_name,omitempty"`       // 广告主名称
	FirstIndustry       string   `json:"first_industry,omitempty"`        // 广告主一级行业
	SecondIndustry      string   `json:"second_industry,omitempty"`       // 广告主二级行业
	ActionType          int      `json:"action_type,omitempty"`           // 处罚动作
	IllegalTime         string   `json:"illegal_time,omitempty"`          // 时间，格式 yyyy-MM-dd HH:mm:ss
	IllegalLevel        int      `json:"illegal_level,omitempty"`         // 违规等级
	IllegalReason       string   `json:"illegal_reason,omitempty"`        // 处罚原因
	ComplaintStatus     string   `json:"complaint_status,omitempty"`      // 申述状态
	RejectReason        string   `json:"reject_reason,omitempty"`         // 驳回原因
	RejectEvidenceList  []string `json:"reject_evidence_list,omitempty"`  // 驳回凭证 URL 列表
}
