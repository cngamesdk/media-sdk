package model

import "errors"

// ========== 获取粉丝数据 ==========
// https://developers.e.qq.com/v3.0/docs/api/video_channel_fans_data/get

// VideoChannelFansDataGetReq 获取粉丝数据请求
type VideoChannelFansDataGetReq struct {
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

func (p *VideoChannelFansDataGetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证获取粉丝数据请求参数
func (p *VideoChannelFansDataGetReq) Validate() error {
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

// VideoChannelFansDataGetResp 获取粉丝数据响应
type VideoChannelFansDataGetResp struct {
	FansInfoList []*FansInfoItem `json:"fans_info_list,omitempty"` // 粉丝数据列表
	PageInfo     *PageInfo       `json:"page_info,omitempty"`      // 分页信息
}

// FansInfoItem 粉丝数据实体
type FansInfoItem struct {
	VideoChannelId   string `json:"video_channel_id,omitempty"`   // 视频号 id
	VideoChannelName string `json:"video_channel_name,omitempty"` // 视频号名称
	Date             int    `json:"date,omitempty"`               // 时间，int 格式，例如 20231101
	FansCnt          int64  `json:"fans_cnt,omitempty"`           // 粉丝数
}
