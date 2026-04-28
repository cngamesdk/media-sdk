package model

import "errors"

// CreativeBatchUpdateItem 批量修改自定义创意单条创意参数
type CreativeBatchUpdateItem struct {
	CreativeId           int64  `json:"creative_id,omitempty"`            // 创意ID，传则为更新，不传则为创建
	UnitId               int64  `json:"unit_id,omitempty"`                // 广告组ID
	CreativeName         string `json:"creative_name,omitempty"`          // 创意名称，1-100字符，同一广告组下不能重复
	CreativeMaterialType int    `json:"creative_material_type,omitempty"` // 素材类型：1=竖版视频 2=横版视频 5=竖版图片 6=横版图片 9=小图 10=组图 11=开屏视频 12=开屏图片 14=DPA模板
	ActionBarText        string `json:"action_bar_text,omitempty"`        // 行动号召按钮文案，直播直投不用填写
	Description          string `json:"description,omitempty"`            // 广告语，1-30字符，支持动态词包，直播直投不用填写
	Recommendation       string `json:"recommendation,omitempty"`         // PLC描述语，开启原生时可用，直播直投不用填写

	// 素材
	PhotoId           string   `json:"photo_id,omitempty"`            // 视频ID
	ImageToken        string   `json:"image_token,omitempty"`         // 封面图片token，不传则使用视频首帧
	ImageTokens       []string `json:"image_tokens,omitempty"`        // 便利贴/图片token，便利贴创意必填，只支持一张
	SplashPhotoIds    []string `json:"splash_photo_ids,omitempty"`    // 开屏视频ID，creative_material_type=11时必填，需4条不同尺寸
	SplashImageTokens []string `json:"splash_image_tokens,omitempty"` // 开屏图片token，creative_material_type=12时必填，需6张不同尺寸
	DpaTemplateId     int64    `json:"dpa_template_id,omitempty"`     // DPA模板ID，creative_material_type=14时必填
	DpaStyleTypes     []string `json:"dpa_style_types,omitempty"`     // 动态商品卡样式ID，动态商品卡必填：14001=区域服务卡

	// 创意分类与标签
	CreativeCategory int      `json:"creative_category,omitempty"` // 创意分类，须是叶子节点，与创意标签同时传或同时不传
	CreativeTag      []string `json:"creative_tag,omitempty"`      // 创意标签，最多20个，每个不超过10字符

	// 广告标签
	NewExposeTag []CreativeCreateNewExposeTag `json:"new_expose_tag,omitempty"` // 广告标签2期，最多2个，直播直投不用填写

	// 监测链接
	ClickTrackUrl       string `json:"click_track_url,omitempty"`         // 第三方点击检测链接，优化目标为激活时必填
	ImpressionUrl       string `json:"impression_url,omitempty"`          // 第三方开始播放监测链接，优化目标为激活时必填
	ActionbarClickUrl   string `json:"actionbar_click_url,omitempty"`     // 第三方点击按钮监测链接
	AdPhotoPlayedT3sUrl string `json:"ad_photo_played_t3s_url,omitempty"` // 第三方有效播放监测链接，与impression_url不可同时使用
	LiveTrackUrl        string `json:"live_track_url,omitempty"`          // 点击监测链接，campaign_type=16时可填

	// 原生投放
	OuterLoopNative int   `json:"outer_loop_native,omitempty"` // 是否开启原生：0=关闭 1=开启
	KolUserId       int64 `json:"kol_user_id,omitempty"`       // 达人ID，outer_loop_native=1时必填
	KolUserType     int   `json:"kol_user_type,omitempty"`     // 达人用户类型：1=普通快手号 2=服务号 3=聚星达人，outer_loop_native=1时必填

	// 直播
	LiveCreativeType int `json:"live_creative_type,omitempty"` // 直播类型：3=直投直播 4=视频引流直播，campaign_type=16时必填
}

// CreativeBatchUpdateReq 批量修改自定义创意请求
type CreativeBatchUpdateReq struct {
	accessTokenReq
	// 必填字段
	AdvertiserId int64                     `json:"advertiser_id"` // 广告主ID，必填
	UnitId       int64                     `json:"unit_id"`       // 广告组ID，必填，一个unit最多创建15个创意
	Creatives    []CreativeBatchUpdateItem `json:"creatives"`     // 创意列表，必填

	// 公共可选字段（优先级低于creatives内的同名字段）
	CreativeCategory            int      `json:"creative_category,omitempty"`             // 创意分类，须是叶子节点，与创意标签同时传或同时不传
	CreativeTag                 []string `json:"creative_tag,omitempty"`                  // 创意标签，最多20个，每个不超过10字符
	ClickTrackUrl               string   `json:"click_track_url,omitempty"`               // 第三方点击检测链接
	ImpressionUrl               string   `json:"impression_url,omitempty"`                // 第三方开始播放监测链接
	ActionbarClickUrl           string   `json:"actionbar_click_url,omitempty"`           // 第三方点击按钮监测链接
	AdPhotoPlayedT3sUrl         string   `json:"ad_photo_played_t3s_url,omitempty"`       // 第三方有效播放监测链接，与impression_url不可同时使用
	HighLightFlash              int      `json:"high_light_flash,omitempty"`              // 高光智投开关：0=关闭 1=打开
	MicroChangeSwitch           int      `json:"micro_change_switch,omitempty"`           // 微改白盒化开关：0=关闭 1=打开
	MaterialIntelligentOptimize int      `json:"material_intelligent_optimize,omitempty"` // 素材智能优化开关：0=关闭 1=开启，仅白名单用户可用
}

func (receiver *CreativeBatchUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *CreativeBatchUpdateReq) Validate() (err error) {
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
	if len(receiver.Creatives) == 0 {
		err = errors.New("creatives is empty")
		return
	}
	return
}

// CreativeBatchUpdateResp 批量修改自定义创意响应数据（仅data部分）
type CreativeBatchUpdateResp struct {
	UpdateCreativeIds []int64 `json:"update_creative_ids"` // 更新的创意ID列表
	AddCreativeIds    []int64 `json:"add_creative_ids"`    // 新建的创意ID列表
}
