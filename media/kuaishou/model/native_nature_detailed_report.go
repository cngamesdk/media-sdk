package model

import "errors"

// NatureDetailedPageInfo 原生经营明细分页信息
type NatureDetailedPageInfo struct {
	CurrentPage int   `json:"current_page"` // 当前页面
	PageSize    int   `json:"page_size"`    // 当前分页大小
	TotalCount  int64 `json:"total_count"`  // 总条数
}

// NatureDetailedSearchParam 原生经营明细筛选条件
type NatureDetailedSearchParam struct {
	AuthorId       []int64 `json:"author_id"`        // 快手号ID
	CampaignType   []int   `json:"campaign_type"`    // 营销目标：2-提升应用安装，5-收集销售线索，7-提升应用活跃，19-快手小程序推广，30-短剧推广
	OcpcActionType []int   `json:"ocpc_action_type"` // 转化目标：180-激活，53-表单优化，190-付费，191-首日ROI
	KolUserType    []int   `json:"kol_user_type"`    // 原生广告类型：1-普通快手号，2-蓝V，3-聚星达人
	ViewType       int     `json:"view_type"`        // 视图类型：2-计划，3-广告组，4-广告创意（明细接口必填）
	ReportEndDay   int64   `json:"report_end_day"`   // 结束时间，时间戳毫秒，必填
	ReportStartDay int64   `json:"report_start_day"` // 开始时间，时间戳毫秒，必填
	SkipSum        bool    `json:"skip_sum"`         // 是否跳过汇总：true-不获取，false-获取
}

// NativeNatureDetailedReportReq 原生报表披露自然流量后查询原生经营明细请求
type NativeNatureDetailedReportReq struct {
	accessTokenReq
	PageInfo     NatureDetailedPageInfo    `json:"page_info"`     // 分页信息
	SearchParam  NatureDetailedSearchParam `json:"search_param"`  // 原生经营明细筛选条件，必填
	AdvertiserId int64                     `json:"advertiser_id"` // 广告主id，必填
}

func (receiver *NativeNatureDetailedReportReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *NativeNatureDetailedReportReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.SearchParam.ReportStartDay <= 0 {
		err = errors.New("search_param.report_start_day is empty")
		return
	}
	if receiver.SearchParam.ReportEndDay <= 0 {
		err = errors.New("search_param.report_end_day is empty")
		return
	}
	if receiver.SearchParam.ViewType < 2 || receiver.SearchParam.ViewType > 4 {
		err = errors.New("search_param.view_type must be 2, 3 or 4")
		return
	}
	return
}

// NativeNatureReportItem 原生经营明细条目
type NativeNatureReportItem struct {
	AdItemClick     int64   `json:"ad_item_click"`      // 行为数
	ActionRatio     float64 `json:"action_ratio"`       // 素材点击率
	ConversionCnt   int64   `json:"conversion_cnt"`     // 原生经营转化数
	CampaignId      int64   `json:"campaign_id"`        // 计划ID
	CampaignName    string  `json:"campaign_name"`      // 计划名称
	UnitId          int64   `json:"unit_id"`            // 广告组ID
	UnitName        string  `json:"unit_name"`          // 广告组名称
	CreativeId      int64   `json:"creative_id"`        // 创意ID
	CreativeName    string  `json:"creative_name"`      // 创意名称
	CreateTime      string  `json:"create_time"`        // 创建时间，时间戳毫秒
	Click           int64   `json:"click"`              // 素材曝光数
	Share           int64   `json:"share"`              // 分享数
	Comment         int64   `json:"comment"`            // 评论数
	Likes           int64   `json:"likes"`              // 点赞数
	KolUserTypeDesc string  `json:"kol_user_type_desc"` // 原生广告类型
	OcpcActionType  string  `json:"ocpc_action_type"`   // 优化目标
	CampaignType    string  `json:"campaign_type"`      // 营销目标
	AuthorId        string  `json:"author_id"`          // 快手号ID
}

// NativeNatureDetailedReportResp 原生经营明细响应数据（仅data部分）
type NativeNatureDetailedReportResp struct {
	ResultList []NativeNatureReportItem `json:"result_list"` // 原生经营明细
	Sum        []NativeNatureReportItem `json:"sum"`         // 全局汇总
	PageInfo   NatureDetailedPageInfo   `json:"page_info"`   // 分页信息
}
