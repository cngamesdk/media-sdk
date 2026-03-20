package model

import (
	"fmt"
	"slices"
	"strings"
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
var ClickMacrosExclude = []string{}

type CustomMacros map[string]string

// 有效触点（点击）宏参数
var ClickMacros CustomMacros

// 展示 宏参数
var ShowMacros CustomMacros

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

// 重置系统默认参数
func (m CustomMacros) Reset(keySrc string, keyDst string) {
	value, ok := m[keySrc]
	if ok {
		delete(m, keySrc)
		m[keyDst] = value
	}
}

// 添加新字段
func (m CustomMacros) Add(key string, value string) {
	m[key] = value
}

// 构建请求参数
func (m CustomMacros) BuildQueryString() string {
	var container []string
	for key, value := range m {
		tempValue := strings.TrimSpace(value)
		fix := "__"
		if !strings.HasPrefix(tempValue, fix) {
			tempValue = fix + tempValue
		}
		if !strings.HasSuffix(tempValue, fix) {
			tempValue = tempValue + fix
		}
		container = append(container, fmt.Sprintf("%s=%s", strings.TrimSpace(strings.ToLower(key)), tempValue))
	}
	return strings.Join(container, "&")
}

// 构建请求URL
func (m CustomMacros) BuildUrl(url string) string {
	connectStr := "?"
	if strings.Contains(url, "?") {
		connectStr = "&"
	}
	return url + connectStr + m.BuildQueryString()
}
