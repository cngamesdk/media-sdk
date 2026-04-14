package model

import (
	"errors"
	"github.com/cngamesdk/media-sdk/media"
	"slices"
	"strings"
)

const (
	Fix = "__"
)

const (
	// 基础ID宏
	PromotionID = "PROMOTION_ID" // 巨量营销体验版中特有的宏参，代表巨量营销体验版的营销ID
	ProjectID   = "PROJECT_ID"   // 巨量营销体验版中特有的宏参，代表巨量营销体验版的项目ID
	// 名称宏
	PromotionName = "PROMOTION_NAME" // 巨量营销体验版中的营销名称
	ProjectName   = "PROJECT_NAME"   // 巨量营销体验版中的项目名称

	// 素材宏 (MID1-MID6)
	MID1          = "MID1"           // 针对巨量营销体验版，图片素材宏参数（下发原始素材id）
	MID2          = "MID2"           // 针对巨量营销体验版，标题素材宏参数（下发原始素材id）
	MID3          = "MID3"           // 针对巨量营销体验版，视频素材宏参数（下发原始素材id）
	MID4          = "MID4"           // 针对巨量营销体验版，搭配playable试玩落地页素材宏参数（下发原始素材id）
	MID5          = "MID5"           // 针对巨量营销体验版，落地页素材宏参数（下发原始素材id）
	MID6          = "MID6"           // 针对巨量营销体验版，安卓下载详情页素材宏参数（下发原始素材id）
	AID           = "AID"            // 营销计划id
	AIDName       = "AID_NAME"       // 营销计划名称
	CID           = "CID"            // 营销创意id，长整型
	CIDName       = "CID_NAME"       // 营销创意名称
	CampaignID    = "CAMPAIGN_ID"    // 计划组id
	CampaignName  = "CAMPAIGN_NAME"  // 计划组名称
	CType         = "CTYPE"          // 创意样式
	AdvertiserID  = "ADVERTISER_ID"  // 客户id
	CSITE         = "CSITE"          // 营销投放位置
	ConvertID     = "CONVERT_ID"     // 转化id
	RequestID     = "REQUEST_ID"     // 请求下发的id
	TrackID       = "TRACK_ID"       // 请求下发的id&创意id的md5,16位
	SL            = "SL"             // 这次请求的语言
	IDFA          = "IDFA"           // IOS 6+的设备id字段，32位
	IDFAMD5       = "IDFA_MD5"       // IOS 6+的设备id的md5摘要，32位
	AndroidID     = "ANDROIDID"      // 安卓id原值的md5，32位
	OAID          = "OAID"           // Android Q及更高版本的设备号，32位
	OAIDMD5       = "OAID_MD5"       // Android Q及更高版本的设备号的md5摘要，32位
	OS            = "OS"             // 操作系统平台
	IPv4          = "IPV4"           // IPv4地址
	IPv6          = "IPV6"           // IPv6地址
	IP            = "IP"             // IP地址
	UA            = "UA"             // 用户代理
	GEO           = "GEO"            // 位置信息
	CityCode      = "CITY_CODE"      // 城市代码
	TS            = "TS"             // 时间戳
	CallbackParam = "CALLBACK_PARAM" // 回调参数
	CallbackURL   = "CALLBACK_URL"   // 回调URL
	Model         = "MODEL"          // 手机型号
	UnionSite     = "UNION_SITE"     // 对外投放版位编码
	CAID          = "CAID"           // 中国营销协会互联网营销标识
	CAIDMD5       = "CAID_MD5"       // 中国营销协会互联网营销标识MD5版本
	ProductID     = "PRODUCTID"      // 商品id，仅支持站内（不支持穿山甲）
	OuterID       = "OUTERID"        // 商品id，同时支持站内和穿山甲
)

// 所有宏参数
var AllMacros = map[string]string{
	PromotionID:   PromotionID,
	ProjectID:     ProjectID,
	PromotionName: PromotionName,
	ProjectName:   ProjectName,
	MID1:          MID1,
	MID2:          MID2,
	MID3:          MID3,
	MID4:          MID4,
	MID5:          MID5,
	MID6:          MID6,
	AID:           AID,
	AIDName:       AIDName,
	CID:           CID,
	CIDName:       CIDName,
	CampaignID:    CampaignID,
	CampaignName:  CampaignName,
	CType:         CType,
	AdvertiserID:  AdvertiserID,
	CSITE:         CSITE,
	ConvertID:     ConvertID,
	RequestID:     RequestID,
	TrackID:       TrackID,
	SL:            SL,
	IDFA:          IDFA,
	IDFAMD5:       IDFAMD5,
	AndroidID:     AndroidID,
	OAID:          OAID,
	OAIDMD5:       OAIDMD5,
	OS:            OS,
	IPv4:          IPv4,
	IPv6:          IPv6,
	IP:            IP,
	UA:            UA,
	GEO:           GEO,
	CityCode:      CityCode,
	TS:            TS,
	CallbackParam: CallbackParam,
	CallbackURL:   CallbackURL,
	Model:         Model,
	UnionSite:     UnionSite,
	CAID:          CAID,
	CAIDMD5:       CAIDMD5,
	ProductID:     ProductID,
	OuterID:       OuterID,
}

// 展示宏参数中排除
var ShowMacrosExclude = []string{
	AdvertiserID,
	IPv4,
	IPv6,
	ProductID,
	OuterID,
}

// 点击宏参数中排除
var ClickMacrosExclude []string

// 有效触点（点击）宏参数
var ClickMacros media.CustomMacros

// 展示 宏参数
var ShowMacros media.CustomMacros

// 初始化默认值
func init() {

	ClickMacros = map[string]string{}
	ShowMacros = map[string]string{}

	for key, value := range AllMacros {
		if !slices.Contains(ClickMacrosExclude, key) {
			ClickMacros[key] = value
		}

		if !slices.Contains(ShowMacrosExclude, key) {
			ShowMacros[key] = value
		}
	}
}

// EventType 事件类型
type EventType int

// MatchType 归因方式
type MatchType int

const (
	OSAndroid                             = 0   // 安卓
	OSiOS                                 = 1   // IOS
	OSHarmony                             = 2   // 鸿蒙
	OSOther                               = 3   // 其他
	EventTypeActivate           EventType = 0   // 激活
	EventTypeRegister           EventType = 1   // 注册
	EventTypePay                EventType = 2   // 付费
	EventTypeForm               EventType = 3   // 表单
	EventTypeOnlineConsult      EventType = 4   // 在线咨询
	EventTypeEffectiveConsult   EventType = 5   // 有效咨询
	EventTypeRetention          EventType = 6   // 次留
	EventTypeAppOrder           EventType = 20  // app内下单
	EventTypeAppVisit           EventType = 21  // app内访问
	EventTypeAppAddCart         EventType = 22  // app内添加购物车
	EventTypeAppPay             EventType = 23  // app内付费
	EventTypeKeyBehavior        EventType = 25  // 关键行为
	EventTypeAuthorize          EventType = 28  // 授权
	EventTypeAppDetailPageUV    EventType = 29  // app内详情页到站uv
	EventTypeClickProduct       EventType = 179 // 点击商品
	EventTypeAddToWishlist      EventType = 128 // 加入收藏/心愿单
	EventTypeReceiveCoupon      EventType = 213 // 领取优惠券
	EventTypeBuyNow             EventType = 175 // 立即购买
	EventTypeAddDeliveryInfo    EventType = 212 // 添加/选定收货信息、电话
	EventTypeAddPaymentInfo     EventType = 127 // 添加/选定支付信息
	EventTypeSubmitOrder        EventType = 176 // 提交订单
	EventTypeOrderConfirm       EventType = 214 // 订单提交/确认收货
	EventTypeEnterLiveRoom      EventType = 202 // 进入直播间
	EventTypeLiveFollow         EventType = 204 // 直播间内点击关注按钮
	EventTypeLiveComment        EventType = 205 // 直播间内评论
	EventTypeLiveReward         EventType = 206 // 直播间内打赏
	EventTypeLiveClickCart      EventType = 207 // 直播间内点击购物车按钮
	EventTypeLiveClickProduct   EventType = 208 // 直播间内商品点击
	EventTypeLiveEnterPlantPage EventType = 209 // 直播间进入种草页跳转到第三方
	EventTypeLiveAddToCart      EventType = 210 // 直播-加购
	EventTypeLivePlaceOrder     EventType = 211 // 直播-下单

	MatchTypeClick MatchType = 0 // 点击
	MatchTypeShow  MatchType = 1 // 展示
	MatchTypePlay  MatchType = 2 // 有效播放归因

	CallbackUrl = "https://ad.oceanengine.com/track/activate/"
)

// ConversionEventReq 转化事件
type ConversionEventReq struct {
	CallbackUrl  string `json:"callback_url,omitempty"`   // 回调地址
	Callback     string `json:"callback"`                 // 点击检测下发的 callback (点击事件必填)
	IDFA         string `json:"idfa"`                     // ios 手机的 idfa 原值 (必填)
	OAID         string `json:"oaid"`                     // Android Q 版本的 oaid 原值 (必填)
	OAIDMD5      string `json:"oaid_md5,omitempty"`       // Android Q 版本的 oaid 原值的md5摘要 (可选)
	CAID1        string `json:"caid1"`                    // 最新版本的中国营销协会互联网营销标识 (必填)
	CAID2        string `json:"caid2"`                    // 老版本的中国营销协会互联网营销标识 (必填)
	OS           int    `json:"os"`                       // 客户端的操作系统类型 (必填)
	Source       string `json:"source,omitempty"`         // 数据来源，客户可自行定义 (可选)
	ConvTime     int64  `json:"conv_time,omitempty"`      // 转化发生的时间，UTC 时间戳 (可选)
	EventType    int    `json:"event_type"`               // 事件类型 (必填)
	MatchType    int    `json:"match_type,omitempty"`     // 归因方式 (可选)
	OuterEventID string `json:"outer_event_id,omitempty"` // 外部事件ID，用于去重 (可选)
}

func (p *ConversionEventReq) Format() {
	if len(p.CallbackUrl) == 0 {
		p.CallbackUrl = CallbackUrl
	}
	return
}

// Validate 验证回调参数
func (p *ConversionEventReq) Validate() error {
	// 验证必填字段
	if p.Callback == "" && strings.Index(p.CallbackUrl, "callback=") < 0 {
		return errors.New("callback为必填")
	}
	if p.OS == OSiOS {
		if p.IDFA == "" {
			return errors.New("idfa为必填")
		}
		if p.CAID1 == "" {
			return errors.New("caid1为必填")
		}
		if p.CAID2 == "" {
			return errors.New("caid2为必填")
		}
	}
	if p.OS == OSAndroid {
		if p.OAID == "" {
			return errors.New("oaid为必填")
		}
	}
	if p.EventType < 0 {
		return errors.New("event_type为必填")
	}

	if p.MatchType < 0 {
		return errors.New("match_type为必填")
	}
	return nil
}

type ConversionEventResp struct {
	Code int    `json:"code"`
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
}
