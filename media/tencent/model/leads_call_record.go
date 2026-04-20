package model

import "errors"

// ========== 获取通话结果 ==========
// https://developers.e.qq.com/v3.0/docs/api/leads_call_record/get

// 呼叫方式常量
const (
	CallDirectionIn  = "CALL_IN"  // 呼入
	CallDirectionOut = "CALL_OUT" // 呼出
)

// 通话结束状态常量
const (
	CallEndStatusNormal       = 0  // 正常通话
	CallEndStatusRejected     = 1  // 拒接
	CallEndStatusNoAnswer     = 2  // 无人接听
	CallEndStatusPowerOff     = 3  // 关机
	CallEndStatusOutOfService = 4  // 停机
	CallEndStatusBusy         = 5  // 正在通话中
	CallEndStatusUnavailable  = 6  // 暂时无法接通
	CallEndStatusEmptyNumber  = 7  // 空号
	CallEndStatusCallerHangUp = 8  // 主叫挂断
	CallEndStatusCallerCancel = 9  // 主叫取消
	CallEndStatusOther        = 10 // 其他
)

// LeadsCallRecordGetReq 获取通话结果请求
type LeadsCallRecordGetReq struct {
	GlobalReq
	AccountID    int64  `json:"account_id"`               // 广告主账号id (必填)
	LeadsId      int64  `json:"leads_id,omitempty"`       // 线索id，与outer_leads_id二选一必填
	OuterLeadsId string `json:"outer_leads_id,omitempty"` // 外部线索id，与leads_id二选一必填
	RequestId    string `json:"request_id,omitempty"`     // 唯一业务请求id
	ContactId    string `json:"contact_id,omitempty"`     // 标识一次外呼行为
}

func (p *LeadsCallRecordGetReq) Format() {
	p.GlobalReq.Format()
}

// Validate 验证获取通话结果请求参数
func (p *LeadsCallRecordGetReq) Validate() error {
	if validateErr := p.GlobalReq.Validate(); validateErr != nil {
		return validateErr
	}

	if p.AccountID == 0 {
		return errors.New("account_id为必填")
	}

	if p.LeadsId == 0 && p.OuterLeadsId == "" {
		return errors.New("leads_id和outer_leads_id二选一必填")
	}

	return nil
}

// LeadsCallRecordGetResp 获取通话结果响应
type LeadsCallRecordGetResp struct {
	CallRecords []*CallRecordItem `json:"call_records,omitempty"` // 通话记录列表
	RequestId   string            `json:"request_id,omitempty"`   // 唯一业务请求id
}

// CallRecordItem 通话记录
type CallRecordItem struct {
	CallId        string `json:"call_id,omitempty"`         // 通话记录唯一id
	AccountID     int64  `json:"account_id,omitempty"`      // 广告主账号id
	LeadsId       int64  `json:"leads_id,omitempty"`        // 线索id
	ContactId     string `json:"contact_id,omitempty"`      // 标识一次外呼行为
	Caller        string `json:"caller,omitempty"`          // 主叫号码
	Callee        string `json:"callee,omitempty"`          // 被叫号码
	CallDirection string `json:"call_direction,omitempty"`  // 呼叫方式，CALL_IN/CALL_OUT
	Duration      int64  `json:"duration,omitempty"`        // 通话时长，单位：秒
	CallStartTime string `json:"call_start_time,omitempty"` // 呼叫开始时间
	CallEndTime   string `json:"call_end_time,omitempty"`   // 呼叫结束时间
	RingTime      string `json:"ring_time,omitempty"`       // 振铃时间
	AnswerTime    string `json:"answer_time,omitempty"`     // 接听时间
	EndStatus     int    `json:"end_status,omitempty"`      // 通话结束状态，0正常 1拒接 2无人接听 3关机 4停机 5通话中 6无法接通 7空号 8主叫挂断 9主叫取消 10其他
	CallRecordUrl string `json:"call_record_url,omitempty"` // 通话录音url，7天内有效
}
