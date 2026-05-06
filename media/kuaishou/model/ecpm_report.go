package model

import "errors"

// EcpmReportReq 快小游ECPM报表请求
type EcpmReportReq struct {
	accessTokenReq
	AdvertiserId int64    `json:"advertiser_id"`       // 广告主ID，必填（主要用于鉴权，无实际业务含义）
	AppId        string   `json:"app_id"`              // 游戏ID，必填
	DataHour     string   `json:"data_hour"`           // 时间范围，必填。格式：YYYY-MM-DD（天级）或 YYYY-MM-DD hh:mm:ss（小时级）
	OpenId       []string `json:"open_id,omitempty"`   // 用户open_id或union_id，单次最多200个，传空字符串查所有用户
	Page         int      `json:"page,omitempty"`      // 页码，从1开始
	PageSize     int      `json:"page_size,omitempty"` // 单页大小，最大500
}

func (receiver *EcpmReportReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *EcpmReportReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.AppId == "" {
		err = errors.New("app_id is empty")
		return
	}
	if receiver.DataHour == "" {
		err = errors.New("data_hour is empty")
		return
	}
	return
}

// EcpmReportDetail 快小游ECPM报表数据明细
type EcpmReportDetail struct {
	Id        string `json:"id"`         // 记录唯一标识（返回记录按ID递增排序）
	EventTime string `json:"event_time"` // 广告计费事件发生时间
	OpenId    string `json:"open_id"`    // 用户的open_id或union_id
	Cost      int64  `json:"cost"`       // 消耗（单位：厘）
}

// EcpmReportResp 快小游ECPM报表响应数据（仅data部分）
type EcpmReportResp struct {
	TotalCount int64              `json:"total_count"` // 总条数
	Details    []EcpmReportDetail `json:"details"`     // 数据详情
}
