package model

import "errors"

// ========== 获取一个账号下的全部通话结果 ==========
// https://developers.e.qq.com/v3.0/docs/api/leads_call_records/get

// LeadsCallRecordsGetReq 获取一个账号下的全部通话结果请求
type LeadsCallRecordsGetReq struct {
	GlobalReq
	AccountID   int64  `json:"account_id"`             // 广告主账号id (必填)
	PageSize    int    `json:"page_size"`              // 页大小 (必填)，必须大于0
	Page        int    `json:"page"`                   // 页数 (必填)，第一页页码是1
	StartDate   string `json:"start_date"`             // 开始呼叫时间查询起点 (必填)，格式YYYY-MM-DD，与end_date差值不大于30天
	EndDate     string `json:"end_date"`               // 开始呼叫时间查询终点 (必填)，格式YYYY-MM-DD，与start_date差值不大于30天
	SearchAfter string `json:"search_after,omitempty"` // 上一次查询最后一条记录的开始呼叫时间，翻第二页及以后时使用
}

func (p *LeadsCallRecordsGetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证获取全部通话结果请求参数
func (p *LeadsCallRecordsGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	if p.PageSize <= 0 {
		return errors.New("page_size为必填且必须大于0")
	}

	if p.Page <= 0 {
		return errors.New("page为必填且必须大于0")
	}

	if p.StartDate == "" {
		return errors.New("start_date为必填")
	}

	if p.EndDate == "" {
		return errors.New("end_date为必填")
	}

	return nil
}

// LeadsCallRecordsGetResp 获取一个账号下的全部通话结果响应
type LeadsCallRecordsGetResp struct {
	PageInfo  *PageInfo            `json:"page_info,omitempty"`  // 分页信息
	RequestId string               `json:"request_id,omitempty"` // 唯一业务请求id
	Entities  []*CallRecordsEntity `json:"entities,omitempty"`   // 通话记录列表
}

// CallRecordsEntity 通话记录实体
type CallRecordsEntity struct {
	CallId        string `json:"call_id,omitempty"`         // 通话记录唯一id
	AccountID     int64  `json:"account_id,omitempty"`      // 广告主账号id
	LeadsId       int64  `json:"leads_id,omitempty"`        // 线索id
	Caller        string `json:"caller,omitempty"`          // 主叫号码
	Callee        string `json:"callee,omitempty"`          // 被叫号码
	CallDirection string `json:"call_direction,omitempty"`  // 呼叫方式，CALL_IN/CALL_OUT
	Duration      int64  `json:"duration,omitempty"`        // 通话时长，单位：秒
	CallStartTime string `json:"call_start_time,omitempty"` // 呼叫开始时间
	CallEndTime   string `json:"call_end_time,omitempty"`   // 呼叫结束时间
	RingTime      string `json:"ring_time,omitempty"`       // 振铃时间
	AnswerTime    string `json:"answer_time,omitempty"`     // 接听时间
	EndStatus     int    `json:"end_status,omitempty"`      // 通话结束状态
	CallRecordUrl string `json:"call_record_url,omitempty"` // 通话录音url，7天内有效
}
