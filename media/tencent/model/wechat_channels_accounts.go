package model

import "errors"

// ========== 视频号-获取视频号列表 ==========
// https://developers.e.qq.com/v3.0/docs/api/wechat_channels_accounts/get

// 过滤字段枚举
const (
	WechatChannelsAccountsFilterFieldAccountId     = "wechat_channels_account_id"   // 按视频号账号 id 过滤
	WechatChannelsAccountsFilterFieldIsAdAcct      = "is_ad_acct"                   // 按是否广告专用账户过滤
	WechatChannelsAccountsFilterFieldCreatedSource = "created_source"               // 按创建来源过滤
	WechatChannelsAccountsFilterFieldVideoId       = "video_id"                     // 按互选 video_id 过滤
	WechatChannelsAccountsFilterFieldAccountName   = "wechat_channels_account_name" // 按视频号名称过滤
)

// 使用场景枚举
const (
	WechatChannelsAccountSceneFeedsAd                    = "WECHAT_CHANNELS_ACCOUNT_SECNE_FEEDS_AD"                      // 视频号动态推广广告
	WechatChannelsAccountSceneLiveAd                     = "WECHAT_CHANNELS_ACCOUNT_SECNE_LIVE_AD"                       // 直播广告
	WechatChannelsAccountSceneFeedsCreative              = "WECHAT_CHANNELS_ACCOUNT_SECNE_FEEDS_CREATIVE"                // 视频号版位创意
	WechatChannelsAccountSceneVideoNativeContentCreative = "WECHAT_CHANNELS_ACCOUNT_SECNE_VIDEO_NATIVE_CONTENT_CREATIVE" // 原生内容广告/创意
)

// 分页常量
const (
	MinWechatChannelsAccountsPage         = 1     // page 最小值
	MaxWechatChannelsAccountsPage         = 99999 // page 最大值
	MinWechatChannelsAccountsPageSize     = 1     // page_size 最小值
	MaxWechatChannelsAccountsPageSize     = 100   // page_size 最大值
	DefaultWechatChannelsAccountsPage     = 1     // page 默认值
	DefaultWechatChannelsAccountsPageSize = 10    // page_size 默认值
	MaxWechatChannelsAccountsFiltering    = 3     // filtering 最大数量
)

// WechatChannelsAccountsFilter 视频号列表过滤条件
type WechatChannelsAccountsFilter struct {
	Field    string   `json:"field"`    // 过滤字段 (必填)
	Operator string   `json:"operator"` // 操作符 (必填)
	Values   []string `json:"values"`   // 字段取值 (必填)
}

// WechatChannelsAccountsGetReq 获取视频号列表请求（GET）
// https://developers.e.qq.com/v3.0/docs/api/wechat_channels_accounts/get
type WechatChannelsAccountsGetReq struct {
	GlobalReq
	AccountId int64                           `json:"account_id"`          // 广告主帐号 id (必填)，不支持代理商 id
	Filtering []*WechatChannelsAccountsFilter `json:"filtering,omitempty"` // 过滤条件，最多 3 个
	Page      int                             `json:"page,omitempty"`      // 搜索页码，1-99999，默认 1
	PageSize  int                             `json:"page_size,omitempty"` // 每页条数，1-100，默认 10
	Scene     string                          `json:"scene,omitempty"`     // 视频号使用场景（枚举）
}

func (r *WechatChannelsAccountsGetReq) Format() {
	r.GlobalReq.Format()
	if r.Page == 0 {
		r.Page = DefaultWechatChannelsAccountsPage
	}
	if r.PageSize == 0 {
		r.PageSize = DefaultWechatChannelsAccountsPageSize
	}
}

// Validate 验证获取视频号列表请求参数
func (r *WechatChannelsAccountsGetReq) Validate() error {
	if r.AccountId == 0 {
		return errors.New("account_id为必填")
	}
	if r.Page < MinWechatChannelsAccountsPage || r.Page > MaxWechatChannelsAccountsPage {
		return errors.New("page须在1-99999之间")
	}
	if r.PageSize < MinWechatChannelsAccountsPageSize || r.PageSize > MaxWechatChannelsAccountsPageSize {
		return errors.New("page_size须在1-100之间")
	}
	if len(r.Filtering) > MaxWechatChannelsAccountsFiltering {
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

// WechatChannelsAccountsItem 视频号列表项
type WechatChannelsAccountsItem struct {
	WechatChannelsAccountId   string `json:"wechat_channels_account_id"`   // 视频号账号 id
	WechatChannelsAccountName string `json:"wechat_channels_account_name"` // 视频号名称
	CreatedTime               int64  `json:"created_time"`                 // 创建时间，时间戳
	LastModifiedTime          int64  `json:"last_modified_time"`           // 最后修改时间，时间戳
	WechatChannelsSpamBlock   bool   `json:"wechat_channels_spam_block"`   // 视频号账号是否封禁
	WechatChannelsSpamSlient  bool   `json:"wechat_channels_spam_slient"`  // 视频号账号是否禁言（注：接口原文拼写为 slient）
	WechatChannelsAccountIcon string `json:"wechat_channels_account_icon"` // 视频号头像
	AuthorizationType         string `json:"authorization_type"`           // 授权类型（已废弃）
	AuthorizationScope        string `json:"authorization_scope"`          // 授权范围（枚举）
	IsBlocked                 bool   `json:"is_blocked"`                   // 是否被限流
	IsPrivate                 bool   `json:"is_private"`                   // 是否设为私密
	IsAdAcct                  bool   `json:"is_ad_acct"`                   // 是否包含广告专用视频号账户
	AuthorizationBeginTime    int64  `json:"authorization_begin_time"`     // 授权开始时间，时间戳
	AuthorizationTtl          int64  `json:"authorization_ttl"`            // 授权有效时间，单位 s
	IsDisable                 bool   `json:"is_disable"`                   // 是否可使用
	DisableMessage            string `json:"disable_message"`              // 禁用原因
	AuthorizationStatus       string `json:"authorization_status"`         // 授权状态（枚举）
}

// WechatChannelsAccountsGetResp 获取视频号列表响应
type WechatChannelsAccountsGetResp struct {
	List     []*WechatChannelsAccountsItem `json:"list"`      // 视频号列表
	PageInfo *PageInfo                     `json:"page_info"` // 分页配置信息
}
