package model

import "errors"

// ========== 视频号-获取视频号动态详情 ==========
// https://developers.e.qq.com/v3.0/docs/api/channels_finderobject/get

// export_id 字段长度常量
const (
	MinChannelsFinderobjectExportIdBytes = 1   // export_id 最小字节数
	MaxChannelsFinderobjectExportIdBytes = 256 // export_id 最大字节数
)

// ChannelsFinderobjectGetReq 获取视频号动态详情请求（GET）
// https://developers.e.qq.com/v3.0/docs/api/channels_finderobject/get
type ChannelsFinderobjectGetReq struct {
	GlobalReq
	AccountId int64  `json:"account_id"` // 广告主帐号 id (必填)，不支持代理商 id
	ExportId  string `json:"export_id"`  // 互选小任务接单视频 id (必填)，1-256字节
}

func (r *ChannelsFinderobjectGetReq) Format() {
	r.GlobalReq.Format()
}

// Validate 验证获取视频号动态详情请求参数
func (r *ChannelsFinderobjectGetReq) Validate() error {
	if r.AccountId == 0 {
		return errors.New("account_id为必填")
	}
	if len(r.ExportId) < MinChannelsFinderobjectExportIdBytes || len(r.ExportId) > MaxChannelsFinderobjectExportIdBytes {
		return errors.New("export_id长度须在1-256字节之间")
	}
	return r.GlobalReq.Validate()
}

// ChannelsFinderobjectGetResp 获取视频号动态详情响应
// object 字段结构与视频号动态列表项一致，复用 ChannelsUserpageobjectsItem
type ChannelsFinderobjectGetResp struct {
	Object *ChannelsUserpageobjectsItem `json:"object"` // 动态详情
}
