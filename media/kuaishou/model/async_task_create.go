package model

import "errors"

// AsyncTaskCreateReq 创建历史数据查询任务请求
type AsyncTaskCreateReq struct {
	accessTokenReq
	AdvertiserId int64                 `json:"advertiser_id"` // 广告主ID，必填
	TaskName     string                `json:"task_name"`     // 任务名称，最大50字符，不能为空，每个账户不能重复
	TaskParams   AsyncTaskCreateParams `json:"task_params"`   // 报表任务参数
}

// AsyncTaskCreateParams 报表任务参数
type AsyncTaskCreateParams struct {
	StartDate           string   `json:"start_date"`                     // 查询开始日期，格式 yyyy-MM-dd，时间跨度不能超过6个月，必填
	EndDate             string   `json:"end_date"`                       // 查询结束日期，格式 yyyy-MM-dd，必填
	ViewType            int      `json:"view_type"`                      // 查询维度（必填）：1=账户 2=广告计划 3=广告组 4=广告创意(自定义) 5=视频报表 7=封面报表 8=便利贴报表 10=程序化创意2.0&智能创意 21=关键词报表 25=搜索词报表
	CampaignIds         []int64  `json:"campaign_ids,omitempty"`         // 广告计划ID集
	UnitIds             []int64  `json:"unit_ids,omitempty"`             // 广告组ID集
	CreativeIds         []int64  `json:"creative_ids,omitempty"`         // 广告创意ID集
	PhotoIds            []int64  `json:"photo_ids,omitempty"`            // 视频ID集，仅view_type=5,7,8可使用
	CoverIds            []int64  `json:"cover_ids,omitempty"`            // 封面ID集，仅view_type=5,7,8可使用
	VirtualCreativeIds  []int64  `json:"virtual_creative_ids,omitempty"` // 程序化创意ID集
	WordInfoIds         []int64  `json:"word_info_ids,omitempty"`        // 推广关键词ID集，仅view_type=21,25可使用，单次最多5000
	Query               []string `json:"query,omitempty"`                // 用户搜索词，仅view_type=25可使用，单次最多5000
	ReportDims          string   `json:"report_dims,omitempty"`          // 投放场景：adScene=按广告场景
	TemporalGranularity string   `json:"temporal_granularity,omitempty"` // 时间粒度：DAILY=天粒度 HOURLY=小时粒度，默认天粒度
	Status              int      `json:"status,omitempty"`               // 状态：1=投放中 2=已暂停 3=已删除
	SelectedColumns     []string `json:"selected_columns,omitempty"`     // 自定义列，仅view_type=21或25支持
}

func (receiver *AsyncTaskCreateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AsyncTaskCreateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.TaskName) <= 0 {
		err = errors.New("task_name is empty")
		return
	}
	if receiver.TaskParams.ViewType <= 0 {
		err = errors.New("task_params.view_type is empty")
		return
	}
	return
}

// AsyncTaskCreateResp 创建历史数据查询任务响应数据（仅data部分）
type AsyncTaskCreateResp struct {
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID
	TaskId       int64 `json:"task_id"`       // 任务ID
}
