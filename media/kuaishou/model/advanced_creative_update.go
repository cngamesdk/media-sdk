package model

import "errors"

// AdvancedCreativeUpdateReq 修改程序化创意请求
type AdvancedCreativeUpdateReq struct {
	accessTokenReq
	// 必填字段
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID，必填
	UnitId       int64 `json:"unit_id"`       // 广告组ID，必填，一个组下只能有一个程序化创意

	// 创意基本信息
	PackageName string   `json:"package_name,omitempty"` // 程序化创意名称，1-100字符
	ActionBar   string   `json:"action_bar,omitempty"`   // 行动号召按钮
	Captions    []string `json:"captions,omitempty"`     // 作品广告语，每个不超过30字符，最多3个

	// 素材列表
	PhotoList []AdvancedCreativePhoto `json:"photo_list,omitempty"` // 素材列表，最多10组（unit_type=10时最多15个）

	// 文案与标签
	CoverSlogans     []string                     `json:"cover_slogans,omitempty"`     // 封面广告语，0-14字符，最多6个；unit_type=10不支持
	CreativeCategory int                          `json:"creative_category,omitempty"` // 创意分类，须是叶子节点，与创意标签同时传或同时不传
	CreativeTag      []string                     `json:"creative_tag,omitempty"`      // 创意标签，最多20个，每个不超过10字符
	NewExposeTag     []CreativeCreateNewExposeTag `json:"new_expose_tag,omitempty"`    // 广告标签2期，最多2个

	// 监测链接
	ClickUrl            string `json:"click_url,omitempty"`               // 第三方点击检测链接，不超过1024字符
	ImpressionUrl       string `json:"impression_url,omitempty"`          // 曝光监测链接
	ActionbarClickUrl   string `json:"actionbar_click_url,omitempty"`     // 第三方ActionBar点击监控链接，不超过1024字符
	AdPhotoPlayedT3sUrl string `json:"ad_photo_played_t3s_url,omitempty"` // 第三方有效播放监测链接，白名单功能；与impression_url不可同时使用

	// 原生投放
	Recommendation string `json:"recommendation,omitempty"` // 原生PLC广告语，开启原生场景下可用

	// 其他
	MicroChangeSwitch           int `json:"micro_change_switch,omitempty"`           // 微改白盒化开关：0=关闭 1=打开
	MaterialIntelligentOptimize int `json:"material_intelligent_optimize,omitempty"` // 素材智能优化开关：0=关闭 1=开启，仅白名单用户可用；unit_type=10不支持
}

func (receiver *AdvancedCreativeUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AdvancedCreativeUpdateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.UnitId <= 0 {
		err = errors.New("unit_id is empty")
		return
	}
	return
}

// AdvancedCreativeUpdateResp 修改程序化创意响应数据（仅data部分）
type AdvancedCreativeUpdateResp struct {
	UnitId int64 `json:"unit_id"` // 广告组ID
}
