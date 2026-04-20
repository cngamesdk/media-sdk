package model

import "errors"

// ========== 获取经销商聚合数据 ==========
// https://developers.e.qq.com/v3.0/docs/api/video_channel_dealer_data/get

// VideoChannelDealerDataGetReq 获取经销商聚合数据请求
type VideoChannelDealerDataGetReq struct {
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
}

func (p *VideoChannelDealerDataGetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证获取经销商聚合数据请求参数
func (p *VideoChannelDealerDataGetReq) Validate() error {
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

	return nil
}

// VideoChannelDealerDataGetResp 获取经销商聚合数据响应
type VideoChannelDealerDataGetResp struct {
	DealerInfoList []*DealerInfoItem `json:"dealer_info_list,omitempty"` // 经销商聚合数据列表
}

// DealerInfoItem 经销商聚合数据实体
type DealerInfoItem struct {
	DealerId         int64  `json:"dealer_id,omitempty"`          // 经销商 id
	DealerName       string `json:"dealer_name,omitempty"`        // 经销商名称
	Date             int    `json:"date,omitempty"`               // 时间，int 格式，例如 20231101
	LeadsCnt         int64  `json:"leads_cnt,omitempty"`          // 线索总数
	ShowingsCnt      int64  `json:"showings_cnt,omitempty"`       // 直播场次数量
	LivingDuration   int64  `json:"living_duration,omitempty"`    // 直播时长
	WatchingPersonUv int64  `json:"watching_person_uv,omitempty"` // 观看人数
	FansCnt          int64  `json:"fans_cnt,omitempty"`           // 粉丝数
}
