package model

import "errors"

// AdvancedCreativePhoto 程序化创意素材
type AdvancedCreativePhoto struct {
	PhotoId              int64    `json:"photo_id"`                    // 视频ID，必填
	CreativeMaterialType int      `json:"creative_material_type"`      // 素材类型，必填：1=竖版视频 2=横版视频
	CoverImageToken      string   `json:"cover_image_token,omitempty"` // 封面图片token，不传则使用视频首帧
	CoverImageUrl        string   `json:"cover_image_url,omitempty"`   // 封面图片URL
	PicList              []string `json:"pic_list,omitempty"`          // 联盟图片（横版/竖版），image_token列表
	StickerStyles        []string `json:"sticker_styles,omitempty"`    // 封面贴纸，仅搜索广告支持，与cover_slogans同时传，最多6个
}

// AdvancedCreativeCreateReq 创建程序化创意请求
type AdvancedCreativeCreateReq struct {
	accessTokenReq
	// 必填字段
	AdvertiserId int64    `json:"advertiser_id"` // 广告主ID，必填
	UnitId       int64    `json:"unit_id"`       // 广告组ID，必填；unit_type=7创建程序化创意，unit_type=10创建智能创意
	PackageName  string   `json:"package_name"`  // 程序化创意名称，必填，1-100字符
	ActionBar    string   `json:"action_bar"`    // 行动号召按钮，必填
	Captions     []string `json:"captions"`      // 作品广告语，必填，每个不超过30字符，最多3个

	// 素材列表
	PhotoList []AdvancedCreativePhoto `json:"photo_list,omitempty"` // 素材列表，最多10组（unit_type=10时最多15个），优先使用此字段

	// 文案与标签
	CoverSlogans     []string                     `json:"cover_slogans,omitempty"`     // 封面广告语，仅搜索广告支持，0-14字符，最多6个；unit_type=10不支持
	CreativeCategory int                          `json:"creative_category,omitempty"` // 创意分类，金融/教育/游戏/小说/电商行业必填，须是叶子节点
	CreativeTag      []string                     `json:"creative_tag,omitempty"`      // 创意标签，与创意分类同时传或同时不传，最多10个，每个不超过10字符
	NewExposeTag     []CreativeCreateNewExposeTag `json:"new_expose_tag,omitempty"`    // 广告标签2期，最多2个

	// 监测链接
	ClickUrl            string `json:"click_url,omitempty"`               // 第三方点击检测链接，不超过1024字符
	ImpressionUrl       string `json:"impression_url,omitempty"`          // 曝光监测链接
	ActionbarClickUrl   string `json:"actionbar_click_url,omitempty"`     // 第三方ActionBar点击监控链接，不超过1024字符
	AdPhotoPlayedT3sUrl string `json:"ad_photo_played_t3s_url,omitempty"` // 第三方有效播放监测链接，白名单功能；与impression_url不可同时使用

	// 原生投放
	OuterLoopNative int    `json:"outer_loop_native,omitempty"` // 是否开启原生：0=关闭 1=开启
	KolUserId       int64  `json:"kol_user_id,omitempty"`       // 达人ID，outer_loop_native=1时必填；短剧推广时为短剧作者ID
	KolUserType     int    `json:"kol_user_type,omitempty"`     // 达人类型：1=普通快手号 2=服务号 3=聚星达人，outer_loop_native=1时必填
	Recommendation  string `json:"recommendation,omitempty"`    // PLC描述语，开启原生时可用

	// 其他
	MicroChangeSwitch           int `json:"micro_change_switch,omitempty"`           // 微改白盒化开关：0=关闭 1=打开
	MaterialIntelligentOptimize int `json:"material_intelligent_optimize,omitempty"` // 素材智能优化开关：0=关闭 1=开启，仅白名单用户可用；unit_type=10不支持
}

func (receiver *AdvancedCreativeCreateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AdvancedCreativeCreateReq) Validate() (err error) {
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
	if len(receiver.PackageName) == 0 {
		err = errors.New("package_name is empty")
		return
	}
	if len(receiver.ActionBar) == 0 {
		err = errors.New("action_bar is empty")
		return
	}
	if len(receiver.Captions) == 0 {
		err = errors.New("captions is empty")
		return
	}
	return
}

// AdvancedCreativeCreateResp 创建程序化创意响应数据（仅data部分）
type AdvancedCreativeCreateResp struct {
	UnitId int64 `json:"unit_id"` // 广告组ID
}
