package model

import "errors"

// ========== 获取线索数据 ==========
// https://developers.e.qq.com/v3.0/docs/api/video_channel_leads_data/get

// VideoChannelLeadsDataGetReq 获取线索数据请求
type VideoChannelLeadsDataGetReq struct {
	GlobalReq
	AccountID         int64    `json:"account_id"`                    // 广告主账号 id，直客账号或子客账号 (必填)
	BrandIds          []string `json:"brand_ids,omitempty"`           // 品牌 id 列表，非经销商管理员必填
	BrandNames        []string `json:"brand_names,omitempty"`         // 品牌名称列表
	DealerIds         []string `json:"dealer_ids,omitempty"`          // 经销商 id 列表
	DealerNames       []string `json:"dealer_names,omitempty"`        // 经销商名称列表
	VideoChannelIds   []string `json:"video_channel_ids,omitempty"`   // 视频号账号 id 列表
	VideoChannelNames []string `json:"video_channel_names,omitempty"` // 视频号账号名称列表
	StartDate         int      `json:"start_date"`                    // 开始日期，例如 20220704 (必填)
	EndDate           int      `json:"end_date"`                      // 结束日期，例如 20220704 (必填)
	Page              int      `json:"page"`                          // 页数，第一页页码是 1 (必填)
	PageSize          int      `json:"page_size"`                     // 页大小，必须大于 0 (必填)
}

func (p *VideoChannelLeadsDataGetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证获取线索数据请求参数
func (p *VideoChannelLeadsDataGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	if p.StartDate == 0 {
		return errors.New("start_date为必填")
	}

	if p.EndDate == 0 {
		return errors.New("end_date为必填")
	}

	if p.Page <= 0 {
		return errors.New("page为必填且必须大于0")
	}

	if p.PageSize <= 0 {
		return errors.New("page_size为必填且必须大于0")
	}

	return nil
}

// VideoChannelLeadsDataGetResp 获取线索数据响应
type VideoChannelLeadsDataGetResp struct {
	LeadsInfoList []*VideoChannelLeadsInfoItem `json:"leads_info_list,omitempty"` // 线索数据列表
	PageInfo      *PageInfo                    `json:"page_info,omitempty"`       // 分页信息
}

// VideoChannelLeadsInfoItem 线索数据实体
type VideoChannelLeadsInfoItem struct {
	VideoChannelId   string `json:"video_channel_id,omitempty"`   // 视频号 id
	VideoChannelName string `json:"video_channel_name,omitempty"` // 视频号名称
	Date             int    `json:"date,omitempty"`               // 时间，int 格式，例如 20231101
	DealerId         int64  `json:"dealer_id,omitempty"`          // 经销商 id
	DealerName       string `json:"dealer_name,omitempty"`        // 经销商名称
	LeadsCnt         int64  `json:"leads_cnt,omitempty"`          // 线索总数
	FormLeadsCnt     int64  `json:"form_leads_cnt,omitempty"`     // 表单预约线索数
	WecomLeadsCnt    int64  `json:"wecom_leads_cnt,omitempty"`    // 加企微线索数
	ConsultLeadsCnt  int64  `json:"consult_leads_cnt,omitempty"`  // 在线咨询线索数
	AdLeadsCnt       int64  `json:"ad_leads_cnt,omitempty"`       // 广告流量线索数
	NaturalLeadsCnt  int64  `json:"natural_leads_cnt,omitempty"`  // 自然线索数
}
