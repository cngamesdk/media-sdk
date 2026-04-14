package model

import "errors"

// ========== 视频号-获取视频号当前的预约直播信息 ==========
// https://developers.e.qq.com/v3.0/docs/api/channels_livenoticeinfo/get

// ChannelsLivenoticeinfoGetReq 获取视频号预约直播信息请求（GET）
// https://developers.e.qq.com/v3.0/docs/api/channels_livenoticeinfo/get
type ChannelsLivenoticeinfoGetReq struct {
	GlobalReq
	AccountId               int64  `json:"account_id"`                           // 广告主帐号 id (必填)，不支持代理商 id
	FinderUsername          string `json:"finder_username,omitempty"`            // 视频号账号 id（已废弃），1-1024字节
	Nickname                string `json:"nickname,omitempty"`                   // 视频号名称，1-1024字节
	WechatChannelsAccountId string `json:"wechat_channels_account_id,omitempty"` // 视频号账号 id，1-1024字节
}

func (r *ChannelsLivenoticeinfoGetReq) Format() {
	r.GlobalReq.Format()
}

// Validate 验证获取视频号预约直播信息请求参数
func (r *ChannelsLivenoticeinfoGetReq) Validate() error {
	if r.AccountId == 0 {
		return errors.New("account_id为必填")
	}
	return r.GlobalReq.Validate()
}

// ChannelsLivenoticeinfoRecord 直播预约记录
type ChannelsLivenoticeinfoRecord struct {
	NoticeId     string `json:"notice_id"`    // 直播预约记录 id
	Status       int    `json:"status"`       // 直播预约状态
	StartTime    int64  `json:"start_time"`   // 开始时间
	Introduction string `json:"introduction"` // 字符串类型
}

// ChannelsLivenoticeinfoGetResp 获取视频号预约直播信息响应
type ChannelsLivenoticeinfoGetResp struct {
	LiveNoticeRecordList []*ChannelsLivenoticeinfoRecord `json:"live_notice_record_list"` // 直播预约记录列表
}
