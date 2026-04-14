package model

import (
	"github.com/cngamesdk/media-sdk/media"
	"slices"
)

const (
	Fix = "__"
)

const (
	ImpressionId           = "IMPRESSION_ID"            // 曝光id
	ImpressionTime         = "IMPRESSION_TIME"          // 	曝光时间
	CampaignId             = "CAMPAIGN_ID"              // 	计划id
	AdgroupId              = "ADGROUP_ID"               // 	广告组id
	AdId                   = "AD_ID"                    // 	广告id
	DynamicCreativeId      = "DYNAMIC_CREATIVE_ID"      // 	创意 ID
	DynamicCreativeName    = "dynamic_creative_name"    // 	创意名称
	CreativeComponentsInfo = "CREATIVE_COMPONENTS_INFO" // 	创意组件信息
	ElementInfo            = "ELEMENT_INFO"             // 	素材信息
	MarketingGoal          = "MARKETING_GOAL"           // 	一级营销目的
	MarketingSubGoal       = "MARKETING_SUB_GOAL"       // 	二级营销目的
	MarketingTargetId      = "MARKETING_TARGET_ID"      // 	营销对象
	MarketingCarrierId     = "MARKETING_CARRIER_ID"     // 	一级营销载体
	MarketingSubCarrierId  = "MARKETING_SUB_CARRIER_ID" // 	二级营销载体
	MarketingAssetId       = "MARKETING_ASSET_ID"       // 	营销资产ID
	MaterialPackageId      = "MATERIAL_PACKAGE_ID"      // 	素材标签ID
	AdPlatformType         = "AD_PLATFORM_TYPE"         // 	广告投放平台
	AdType                 = "AD_TYPE"                  // 	广告类型
	AccountId              = "ACCOUNT_ID"               // 	广告主id
	AgencyId               = "AGENCY_ID"                // 	代理商id
	ImpressionSkuId        = "IMPRESSION_SKU_ID"        // 	曝光sku
	BillingEvent           = "BILLING_EVENT"            // 	计费类型

	// 兼容新旧版本广告 - 应用直达链接（Android）
	DeeplinkURL = "DEEPLINK_URL"

	// 兼容新旧版本广告 - 应用直达链接（iOS）
	UniversalLink = "UNIVERSAL_LINK"

	// 兼容新旧版本广告 - 落地页地址
	PageURL = "PAGE_URL"

	// 兼容新旧版本广告 - 设备类型
	DeviceOSType = "DEVICE_OS_TYPE"

	// 兼容新旧版本广告 - 请求时间
	ProcessTime = "PROCESS_TIME"

	// 兼容新旧版本广告 - 应用id
	PromotedObjectID = "PROMOTED_OBJECT_ID"
	// 兼容新旧版本广告 - 推广类型
	PromotedObjectType = "PROMOTED_OBJECT_TYPE"

	// 兼容新旧版本广告 - 请求id
	RequestID = "REQUEST_ID"

	// 兼容新旧版本广告 - 设备id (imei或idfa的加密值)
	MUID = "MUID"

	// 兼容新旧版本广告 - 安卓id做md5加密后小写
	HashAndroidID = "HASH_ANDROID_ID"

	// 兼容新旧版本广告 - 媒体投放系统获取的用户终端的公共IPV4地址
	IP = "IP"

	// 兼容新旧版本广告 - 用户代理 (user_agent)
	UserAgent = "USER_AGENT"
	// 兼容新旧版本广告 - 回调地址（自用因ap场景必填，其他场景不需要填写）
	Callback = "CALLBACK"

	// 兼容新旧版本广告 - 联盟广告位id（依据联盟白名单判断）
	EncryptedPositionID = "ENCRYPTED_POSITION_ID"

	// 兼容新旧版本广告 - 媒体投放系统获取的用户终端的公共IPV6地址
	IPv6 = "IPV6"

	// 兼容新旧版本广告 - Android Q及更高版本的设备号，64位及以下，取原值后做md5加密
	HashOaid = "HASH_OAID"

	// 兼容新旧版本广告 - URL Encode后的JSON数组
	// 其中qaid为中广协ID（即CAID），hash_qaid为CAID原值MD5加密后的结果，
	// version为腾讯版本号，支持两个版本同时下发（即最新版和上一版）
	Caid = "QAID_CAA"

	// 兼容新旧版本广告 - 广告组名称
	AdgroupName = "ADGROUP_NAME"
	// 兼容新旧版本广告 - 广告版位
	SiteSetName = "SITE_SET_NAME"

	// 仅旧版本广告 - 创意名称
	AdName = "AD_NAME"

	// 兼容新旧版本广告 - 机型
	Model = "MODEL"

	// 兼容新旧版本广告 - 专用于ROI策略(原联合专区RuleLab)的UV分组实验信息，用于区分实验组和对照组
	BoostExpInfo = "BOOST_EXP_INFO"

	// 兼容新旧版本广告 - 专用于ROI策略(原联合专区RuleLab)，对应ROI策略的策略ID(原Rule ID)，用于定位对应的ROI策略
	BoostModelID = "BOOST_MODEL_ID"

	// 兼容新旧版本广告 - 专用于网页类小程序转化规则的点击监测下发
	// 每个用户针对小程序应用会产生一个安全的OpenID，只针对当前的小程序有效
	WechatOpenID = "WECHAT_OPEN_ID"

	// 兼容新旧版本广告 - 专用于搜索广告的关键词ID下发
	KeywordID = "KEYWORD_ID"

	// 兼容新旧版本广告 - 专用于搜索广告的关键词下发
	KeywordText = "KEYWORD_TEXT"

	// 兼容新旧版本广告 - 媒体投放系统获取的用户终端的公共IPV4地址MD5加密后转小写
	// 仅在新版转化里支持配置
	IPMd5 = "IP_MD5"

	// 兼容新旧版本广告 - 媒体投放系统获取的用户终端的公共IPV6地址MD5加密后转小写
	// 仅在新版转化里支持配置
	IPv6Md5 = "IPV6_MD5"

	// 兼容新旧版本广告 - URL Encode后的JSON数组
	// 其中caid为中广协ID（即CAID），hash_caid为CAID原值MD5加密后的结果，
	// version为中广协caid版本号，支持两个版本同时下发（即最新版和上一版）
	// 仅在新版转化归因支持
	CaidRaw = "CAID"
)

// 所有宏参数
var AllMacros = map[string]string{
	ImpressionId:           ImpressionId,
	ImpressionTime:         ImpressionTime,
	CampaignId:             CampaignId,
	AdgroupId:              AdgroupId,
	AdId:                   AdId,
	DynamicCreativeId:      DynamicCreativeId,
	DynamicCreativeName:    DynamicCreativeName,
	CreativeComponentsInfo: CreativeComponentsInfo,
	ElementInfo:            ElementInfo,
	MarketingGoal:          MarketingGoal,
	MarketingSubGoal:       MarketingSubGoal,
	MarketingTargetId:      MarketingTargetId,
	MarketingCarrierId:     MarketingCarrierId,
	MarketingSubCarrierId:  MarketingSubCarrierId,
	MarketingAssetId:       MarketingAssetId,
	MaterialPackageId:      MaterialPackageId,
	AdPlatformType:         AdPlatformType,
	AdType:                 AdType,
	AccountId:              AccountId,
	AgencyId:               AgencyId,
	ImpressionSkuId:        ImpressionSkuId,
	BillingEvent:           BillingEvent,
	DeeplinkURL:            DeeplinkURL,
	UniversalLink:          UniversalLink,
	PageURL:                PageURL,
	DeviceOSType:           DeviceOSType,
	ProcessTime:            ProcessTime,
	PromotedObjectID:       PromotedObjectID,
	PromotedObjectType:     PromotedObjectType,
	RequestID:              RequestID,
	MUID:                   MUID,
	HashAndroidID:          HashAndroidID,
	IP:                     IP,
	UserAgent:              UserAgent,
	Callback:               Callback,
	EncryptedPositionID:    EncryptedPositionID,
	IPv6:                   IPv6,
	HashOaid:               HashOaid,
	Caid:                   Caid,
	AdgroupName:            AdgroupName,
	SiteSetName:            SiteSetName,
	AdName:                 AdName,
	Model:                  Model,
	BoostExpInfo:           BoostExpInfo,
	BoostModelID:           BoostModelID,
	WechatOpenID:           WechatOpenID,
	KeywordID:              KeywordID,
	KeywordText:            KeywordText,
	IPMd5:                  IPMd5,
	IPv6Md5:                IPv6Md5,
	CaidRaw:                CaidRaw,
}

var ShowMacrosAppAndroid media.CustomMacros

var ShowMacrosAppAndroidExclude = []string{
	UniversalLink,
	Caid,
}

var ShowMacrosAppIos media.CustomMacros

var ShowMacrosAppIosExclude = []string{
	DeeplinkURL,
	HashAndroidID,
	HashOaid,
}

var ShowMacrosWeb media.CustomMacros

var ShowMacrosWebExclude = []string{
	DeeplinkURL,
	Model,
	PromotedObjectID,
	UniversalLink,
}

func init() {
	ShowMacrosAppAndroid = make(map[string]string)
	ShowMacrosAppIos = make(map[string]string)
	ShowMacrosWeb = make(map[string]string)

	for key, value := range AllMacros {
		if !slices.Contains(ShowMacrosAppAndroidExclude, key) {
			ShowMacrosAppAndroid[key] = Fix + value + Fix
		}
	}

	for key, value := range AllMacros {
		if !slices.Contains(ShowMacrosAppIosExclude, key) {
			ShowMacrosAppIos[key] = Fix + value + Fix
		}
	}

	for key, value := range AllMacros {
		if !slices.Contains(ShowMacrosWebExclude, key) {
			ShowMacrosWeb[key] = Fix + value + Fix
		}
	}
}
