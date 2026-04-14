package model

import "errors"

// ========== 视频号授权-获取视频号授权记录列表 ==========
// https://developers.e.qq.com/v3.0/docs/api/wechat_channels_authorization/get

// 过滤字段枚举
const (
	WechatChannelsAuthorizationFilterFieldAccountId           = "wechat_channels_account_id" // 按视频号账号 id 过滤
	WechatChannelsAuthorizationFilterFieldAuthorizationStatus = "authorization_status"       // 按授权状态过滤
	WechatChannelsAuthorizationFilterFieldAuthorizationId     = "authorization_id"           // 按视频号授权 id 过滤
)

// 分页常量
const (
	MinWechatChannelsAuthorizationPage         = 1     // page 最小值
	MaxWechatChannelsAuthorizationPage         = 99999 // page 最大值
	MinWechatChannelsAuthorizationPageSize     = 1     // page_size 最小值
	MaxWechatChannelsAuthorizationPageSize     = 100   // page_size 最大值
	DefaultWechatChannelsAuthorizationPage     = 1     // page 默认值
	DefaultWechatChannelsAuthorizationPageSize = 10    // page_size 默认值
	MaxWechatChannelsAuthorizationFiltering    = 3     // filtering 最大数量
)

// WechatChannelsAuthorizationFilter 视频号授权列表过滤条件
type WechatChannelsAuthorizationFilter struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)：wechat_channels_account_id/authorization_status/authorization_id
	Operator string   `json:"operator"` // 操作符 (必填)：EQUALS/IN
	Values   []string `json:"values"`   // 字段取值 (必填)
}

// WechatChannelsAuthorizationGetReq 获取视频号授权记录列表请求（GET）
// https://developers.e.qq.com/v3.0/docs/api/wechat_channels_authorization/get
type WechatChannelsAuthorizationGetReq struct {
	GlobalReq
	AccountId                 int64                                `json:"account_id"`                             // 广告主帐号 id (必填)
	WechatChannelsAccountName string                               `json:"wechat_channels_account_name,omitempty"` // 视频号名称，0-255字节
	Page                      int                                  `json:"page,omitempty"`                         // 搜索页码，1-99999，默认1
	PageSize                  int                                  `json:"page_size,omitempty"`                    // 每页条数，1-100，默认10
	Filtering                 []*WechatChannelsAuthorizationFilter `json:"filtering,omitempty"`                    // 过滤条件，最多3个
}

func (r *WechatChannelsAuthorizationGetReq) Format() {
	r.GlobalReq.Format()
	if r.Page == 0 {
		r.Page = DefaultWechatChannelsAuthorizationPage
	}
	if r.PageSize == 0 {
		r.PageSize = DefaultWechatChannelsAuthorizationPageSize
	}
}

// Validate 验证获取视频号授权记录列表请求参数
func (r *WechatChannelsAuthorizationGetReq) Validate() error {
	if r.AccountId == 0 {
		return errors.New("account_id为必填")
	}
	if r.Page < MinWechatChannelsAuthorizationPage || r.Page > MaxWechatChannelsAuthorizationPage {
		return errors.New("page须在1-99999之间")
	}
	if r.PageSize < MinWechatChannelsAuthorizationPageSize || r.PageSize > MaxWechatChannelsAuthorizationPageSize {
		return errors.New("page_size须在1-100之间")
	}
	if len(r.Filtering) > MaxWechatChannelsAuthorizationFiltering {
		return errors.New("filtering数组长度不能超过3")
	}
	for i, f := range r.Filtering {
		if f == nil {
			return errors.New("filtering[" + itoa(i) + "]不能为空")
		}
		if f.Field == "" {
			return errors.New("filtering[" + itoa(i) + "].field为必填")
		}
		if f.Operator == "" {
			return errors.New("filtering[" + itoa(i) + "].operator为必填")
		}
		if len(f.Values) == 0 {
			return errors.New("filtering[" + itoa(i) + "].values为必填")
		}
	}
	return r.GlobalReq.Validate()
}

// WechatChannelsAuthorizationItem 视频号授权记录列表项
type WechatChannelsAuthorizationItem struct {
	WechatChannelsAccountName string   `json:"wechat_channels_account_name"` // 视频号名称
	FinderUsername            string   `json:"finder_username"`              // 视频号账号 id（已废弃）
	AuthorizationId           string   `json:"authorization_id"`             // 视频号授权 id
	AuthorizationBeginTime    int64    `json:"authorization_begin_time"`     // 授权开始时间（时间戳）
	AuthorizationTtl          int64    `json:"authorization_ttl"`            // 授权有效时间（单位s）
	AuthorizationStatus       string   `json:"authorization_status"`         // 授权状态（枚举）
	AuditMsg                  string   `json:"audit_msg"`                    // 审核原因
	AuthorizationType         string   `json:"authorization_type"`           // 授权类型（已废弃，枚举）
	AuthorizationScope        string   `json:"authorization_scope"`          // 授权范围（枚举）
	IsAdAcct                  bool     `json:"is_ad_acct"`                   // 是否包含广告专用视频号账户（只读）
	WechatChannelsAccountIcon string   `json:"wechat_channels_account_icon"` // 视频号头像
	WechatChannelsAccountId   string   `json:"wechat_channels_account_id"`   // 视频号账号 id
	IsBlocked                 bool     `json:"is_blocked"`                   // 是否被限流
	IsPrivate                 bool     `json:"is_private"`                   // 是否设为私密
	LogoutTimeSecond          int64    `json:"logout_time_second"`           // 预期注销时间（unix 时间戳，精确到秒）
	CreatedSourceList         []string `json:"created_source_list"`          // 视频号创建来源列表（枚举[]）
	AuthorizationQrCodeUrl    string   `json:"authorization_qr_code_url"`    // 授权码图片链接，扫码后在微信打开
	AuthorizationExpiredTime  int64    `json:"authorization_expired_time"`   // 授权流程过期时间（时间戳）
	AuthorizationAgreement    string   `json:"authorization_agreement"`      // 视频号授权协议 url
}

// WechatChannelsAuthorizationGetResp 获取视频号授权记录列表响应
type WechatChannelsAuthorizationGetResp struct {
	List     []*WechatChannelsAuthorizationItem `json:"list"`      // 授权记录列表
	PageInfo *PageInfo                          `json:"page_info"` // 分页配置信息
}
