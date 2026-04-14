package model

import "errors"

// ========== 视频号授权-删除视频号授权 ==========
// https://developers.e.qq.com/v3.0/docs/api/wechat_channels_authorization/delete

// WechatChannelsAuthorizationDeleteReq 删除视频号授权请求（POST JSON）
// https://developers.e.qq.com/v3.0/docs/api/wechat_channels_authorization/delete
type WechatChannelsAuthorizationDeleteReq struct {
	GlobalReq
	AccountId               int64  `json:"account_id"`                           // 广告主帐号 id (必填)
	AuthorizationId         string `json:"authorization_id,omitempty"`           // 视频号授权 id，1-20480字节
	FinderUsername          string `json:"finder_username,omitempty"`            // 视频号账号 id（已废弃），1-1024字节
	WechatChannelsAccountId string `json:"wechat_channels_account_id,omitempty"` // 视频号账号 id，1-1024字节
}

func (r *WechatChannelsAuthorizationDeleteReq) Format() {
	r.GlobalReq.Format()
}

// Validate 验证删除视频号授权请求参数
func (r *WechatChannelsAuthorizationDeleteReq) Validate() error {
	if r.AccountId == 0 {
		return errors.New("account_id为必填")
	}
	return r.GlobalReq.Validate()
}

// WechatChannelsAuthorizationDeleteResp 删除视频号授权响应
type WechatChannelsAuthorizationDeleteResp struct {
	AuthorizationId         string `json:"authorization_id"`           // 视频号授权 id
	FinderUsername          string `json:"finder_username"`            // 视频号账号 id（已废弃）
	WechatChannelsAccountId string `json:"wechat_channels_account_id"` // 视频号账号 id
}
