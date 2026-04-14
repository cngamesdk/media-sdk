package model

import "errors"

// ========== 视频号授权-更新视频号授权 ==========
// https://developers.e.qq.com/v3.0/docs/api/wechat_channels_authorization/update

// WechatChannelsAuthorizationUpdateReq 更新视频号授权请求（POST JSON）
// https://developers.e.qq.com/v3.0/docs/api/wechat_channels_authorization/update
type WechatChannelsAuthorizationUpdateReq struct {
	GlobalReq
	AccountId                      int64                                       `json:"account_id"`                       // 广告主帐号 id (必填)
	AuthorizationId                string                                      `json:"authorization_id"`                 // 视频号授权 id (必填)，1-20480字节
	AuthorizationBeginTime         int64                                       `json:"authorization_begin_time"`         // 授权开始时间（时间戳，必填），0-9999999999
	AuthorizationTtl               int64                                       `json:"authorization_ttl"`                // 授权有效时间（单位s，必填），0-9999999999
	AuthorizationRelationship      string                                      `json:"authorization_relationship"`       // 视频号授权资质主体关系（必填，已废弃）：RELATIONSHIP_UNKNOWN/RELATIONSHIP_CORPORATION/RELATIONSHIP_EMPLOYMENT
	AuthorizationCertificationList []*WechatChannelsAuthorizationCertification `json:"authorization_certification_list"` // 资质列表（必填），最大255条
	AuthorizationScope             string                                      `json:"authorization_scope,omitempty"`    // 授权范围：DEFAULT/ALL/VIDEO/LIVE
}

func (r *WechatChannelsAuthorizationUpdateReq) Format() {
	r.GlobalReq.Format()
}

// Validate 验证更新视频号授权请求参数
func (r *WechatChannelsAuthorizationUpdateReq) Validate() error {
	if r.AccountId == 0 {
		return errors.New("account_id为必填")
	}
	if r.AuthorizationId == "" {
		return errors.New("authorization_id为必填")
	}
	if r.AuthorizationRelationship == "" {
		return errors.New("authorization_relationship为必填")
	}
	if len(r.AuthorizationCertificationList) == 0 {
		return errors.New("authorization_certification_list为必填")
	}
	if len(r.AuthorizationCertificationList) > MaxWechatChannelsAuthorizationCertificationListLen {
		return errors.New("authorization_certification_list最大长度为255")
	}
	for i, cert := range r.AuthorizationCertificationList {
		if cert.CertificationCode == "" {
			return errors.New("authorization_certification_list[" + itoa(i) + "].certification_code为必填")
		}
	}
	return r.GlobalReq.Validate()
}

// WechatChannelsAuthorizationUpdateResp 更新视频号授权响应
type WechatChannelsAuthorizationUpdateResp struct {
	AuthorizationId string `json:"authorization_id"` // 视频号授权 id
}
