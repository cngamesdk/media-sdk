package model

import "errors"

// ExploreReportParam 探索工具报表请求参数
type ExploreReportParam struct {
	Id               int64 `json:"id"`                 // 报表记录主键id，必填
	ExploreType      int   `json:"explore_type"`       // 探索类型：1=加速探索 2=一键复苏 3=辅助探索，必填
	ExploreStartTime int64 `json:"explore_start_time"` // 开始时间（毫秒时间戳），必填
	ExploreEndTime   int64 `json:"explore_end_time"`   // 结束时间（毫秒时间戳），必填
}

// ExploreReportReq 探索工具关键报表数据请求
type ExploreReportReq struct {
	accessTokenReq
	AdvertiserId int64              `json:"advertiser_id"` // 广告主ID，必填
	UnitId       int64              `json:"unit_id"`       // 广告组ID，必填
	Param        ExploreReportParam `json:"param"`         // 请求参数，必填
}

func (receiver *ExploreReportReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *ExploreReportReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.UnitId <= 0 {
		err = errors.New("unit_id is empty")
		return
	}
	if receiver.Param.Id <= 0 {
		err = errors.New("param.id is empty")
		return
	}
	return
}

// ExploreReportResp 探索工具关键报表数据响应（仅data部分）
type ExploreReportResp struct {
	ExploreBudget  int64   `json:"explore_budget"`  // 探索预算
	TotalCharge    int64   `json:"total_charge"`    // 消耗
	Click          int64   `json:"click"`           // 素材曝光数
	ActionbarClick int64   `json:"actionbar_click"` // 行为数
	ActionRatio    float64 `json:"action_ratio"`    // 素材点击率
}
