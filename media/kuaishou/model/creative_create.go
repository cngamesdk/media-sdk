package model

import "errors"

// CreativeCreateNewExposeTag 广告标签
type CreativeCreateNewExposeTag struct {
	Text string `json:"text"` // 广告标签文字，必填
	Url  string `json:"url"`  // 广告标签链接，必填
}

// CreativeCreateReq 创建自定义创意请求
type CreativeCreateReq struct {
	accessTokenReq
	// 必填字段
	AdvertiserId         int64  `json:"advertiser_id"`             // 广告主ID，必填
	UnitId               int64  `json:"unit_id"`                   // 广告组ID，必填
	CreativeName         string `json:"creative_name"`             // 创意名称，必填，1-100字符，同一广告组下不能重复
	CreativeMaterialType int    `json:"creative_material_type"`    // 素材类型，必填：1=竖版视频 2=横版视频 5=竖版图片 6=横版图片 9=小图 10=组图 11=开屏视频 12=开屏图片 14=DPA模板
	ActionBarText        string `json:"action_bar_text,omitempty"` // 行动号召按钮文案，直播直投创意不用填写

	// 素材
	PhotoId           string   `json:"photo_id,omitempty"`            // 视频ID
	ImageToken        string   `json:"image_token,omitempty"`         // 封面图片token，不传则使用视频首帧
	ImageTokens       []string `json:"image_tokens,omitempty"`        // 便利贴/图片/小图/组图图片token；组图需上传3张
	SplashPhotoIds    []string `json:"splash_photo_ids,omitempty"`    // 开屏视频ID，creative_material_type=11时必填，需4条不同尺寸
	SplashImageTokens []string `json:"splash_image_tokens,omitempty"` // 开屏图片token，creative_material_type=12时必填，需6张不同尺寸
	DpaTemplateId     int64    `json:"dpa_template_id,omitempty"`     // DPA模板ID，creative_material_type=14时必填
	DpaStyleTypes     []string `json:"dpa_style_types,omitempty"`     // 动态商品卡样式，如：14001=区域服务卡

	// 文案
	Description    string `json:"description,omitempty"`    // 广告语，1-30字符，支持动态词包，开屏/直播直投不用填
	Recommendation string `json:"recommendation,omitempty"` // PLC自定义文案，直播直投不用填

	// 创意分类与标签
	CreativeCategory int      `json:"creative_category,omitempty"` // 创意分类，金融/教育/游戏/小说/电商行业必填，须是叶子节点
	CreativeTag      []string `json:"creative_tag,omitempty"`      // 创意标签，与创意分类同时传或同时不传，最多10个，每个不超过10字符

	// 监测链接
	ClickTrackUrl       string `json:"click_track_url,omitempty"`         // 第三方点击检测链接，scene_id=1/2/6/7/10时可选；优化目标为激活时必填
	ImpressionUrl       string `json:"impression_url,omitempty"`          // 第三方开始播放监测链接，优化目标为激活时必填
	ActionbarClickUrl   string `json:"actionbar_click_url,omitempty"`     // 第三方点击按钮监测链接
	AdPhotoPlayedT3sUrl string `json:"ad_photo_played_t3s_url,omitempty"` // 第三方有效播放监测链接，白名单功能；与impression_url不可同时使用
	LiveTrackUrl        string `json:"live_track_url,omitempty"`          // 点击监测链接，campaign_type=16时可填

	// 广告标签
	NewExposeTag []CreativeCreateNewExposeTag `json:"new_expose_tag,omitempty"` // 广告标签2期，最多2个，直播直投不用填写

	// 原生投放
	OuterLoopNative int   `json:"outer_loop_native,omitempty"` // 是否开启原生：0=关闭 1=开启
	KolUserId       int64 `json:"kol_user_id,omitempty"`       // 原生投放目标达人ID，outer_loop_native=1时必填；短剧推广时为短剧作者ID
	KolUserType     int   `json:"kol_user_type,omitempty"`     // 原生达人用户类型：1=普通快手号 2=服务号 3=聚星达人，outer_loop_native=1时必填

	// 直播
	LiveCreativeType int `json:"live_creative_type,omitempty"` // 直播类型，campaign_type=16时必填：3=直投直播 4=视频引流直播

	// 其他
	SiteId                      int64 `json:"site_id,omitempty"`                       // 建站ID
	HighLightFlash              int   `json:"high_light_flash,omitempty"`              // 高光创意状态：0=关闭 1=开启
	MicroChangeSwitch           int   `json:"micro_change_switch,omitempty"`           // 爆款裂变开关：0=关闭 1=打开
	MaterialIntelligentOptimize int   `json:"material_intelligent_optimize,omitempty"` // 素材智能优化开关：0=关闭 1=开启，仅白名单用户可用
}

func (receiver *CreativeCreateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *CreativeCreateReq) Validate() (err error) {
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
	if len(receiver.CreativeName) == 0 {
		err = errors.New("creative_name is empty")
		return
	}
	if receiver.CreativeMaterialType <= 0 {
		err = errors.New("creative_material_type is empty")
		return
	}
	return
}

// CreativeCreateResp 创建自定义创意响应数据（仅data部分）
type CreativeCreateResp struct {
	CreativeId int64 `json:"creative_id"` // 创意ID
}
