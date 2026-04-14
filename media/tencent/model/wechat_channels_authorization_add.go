package model

import "errors"

// ========== 视频号授权-创建视频号授权 ==========
// https://developers.e.qq.com/v3.0/docs/api/wechat_channels_authorization/add

// authorization_certification_list 数组最大长度
const MaxWechatChannelsAuthorizationCertificationListLen = 255

// WechatChannelsAuthorizationCertification 资质信息
type WechatChannelsAuthorizationCertification struct {
	CertificationImage       string `json:"certification_image,omitempty"`         // 营业执照/企业资质证明图片 URL，1-255字节
	CertificationImagePageNo int    `json:"certification_image_page_no,omitempty"` // 图片页码：0-正面，1-反面，2-手持半身，最小值0，最大值2
	CertificationCode        string `json:"certification_code"`                    // 资质类型 (必填)
	CertificationName        string `json:"certification_name,omitempty"`          // 资质名称
	CertificationNumber      string `json:"certification_number,omitempty"`        // 资质编码
	CertificationId          int64  `json:"certification_id,omitempty"`            // 资质 id
	CertificationImageId     string `json:"certification_image_id,omitempty"`      // 图片 id，1-64字节
}

// WechatChannelsAuthorizationAddReq 创建视频号授权请求（POST JSON）
// https://developers.e.qq.com/v3.0/docs/api/wechat_channels_authorization/add
type WechatChannelsAuthorizationAddReq struct {
	GlobalReq
	AccountId                      int64                                       `json:"account_id"`                                 // 广告主帐号 id (必填)
	WechatChannelsAccountName      string                                      `json:"wechat_channels_account_name,omitempty"`     // 视频号名称，0-255字节
	AuthorizationCertificationList []*WechatChannelsAuthorizationCertification `json:"authorization_certification_list,omitempty"` // 资质列表，最大255条
	AuthorizationRelationship      string                                      `json:"authorization_relationship,omitempty"`       // 视频号授权资质主体关系（已废弃）：RELATIONSHIP_UNKNOWN/RELATIONSHIP_CORPORATION/RELATIONSHIP_EMPLOYMENT
	AuthorizationScope             string                                      `json:"authorization_scope,omitempty"`              // 授权范围：DEFAULT/ALL/VIDEO/LIVE
	AuthorizationBeginTime         int64                                       `json:"authorization_begin_time,omitempty"`         // 授权开始时间（时间戳），0-9999999999
	AuthorizationTtl               int64                                       `json:"authorization_ttl,omitempty"`                // 授权有效时间（单位s），0-9999999999
}

func (r *WechatChannelsAuthorizationAddReq) Format() {
	r.GlobalReq.Format()
}

// Validate 验证创建视频号授权请求参数
func (r *WechatChannelsAuthorizationAddReq) Validate() error {
	if r.AccountId == 0 {
		return errors.New("account_id为必填")
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

// WechatChannelsAuthorizationAddResp 创建视频号授权响应
type WechatChannelsAuthorizationAddResp struct {
	FinderUsername           string `json:"finder_username"`            // 视频号账号 id（已废弃）
	AuthorizationQrCodeUrl   string `json:"authorization_qr_code_url"`  // 授权码图片链接，扫码后在微信打开
	AuthorizationDescription string `json:"authorization_description"`  // 授权信息描述
	AuthorizationExpiredTime int64  `json:"authorization_expired_time"` // 授权流程过期时间（时间戳）
	AuthorizationAgreement   string `json:"authorization_agreement"`    // 视频号授权协议 url
	WechatChannelsAccountId  string `json:"wechat_channels_account_id"` // 视频号账号 id
}
