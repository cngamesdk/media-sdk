package model

import "errors"

// CreativeUpdateReq 修改自定义创意请求
type CreativeUpdateReq struct {
	accessTokenReq
	// 必填字段
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID，必填
	CreativeId   int64 `json:"creative_id"`   // 广告创意ID，必填

	// 创意基本信息
	CreativeName         string `json:"creative_name,omitempty"`          // 广告创意名称，1-100字符，不能重复
	CreativeMaterialType int    `json:"creative_material_type,omitempty"` // 素材类型：1=竖版视频 2=横版视频 5=竖版图片 6=横版图片 9=小图 10=组图 11=开屏视频 12=开屏图片
	ActionBarText        string `json:"action_bar_text,omitempty"`        // 行动号召按钮文案，直播直投不用填写
	Description          string `json:"description,omitempty"`            // 广告语，1-30字符，支持动态词包，直播直投不用填写
	Recommendation       string `json:"recommendation,omitempty"`         // 原生PLC广告语，开启原生场景下可用，直播直投不用填写

	// 素材
	PhotoId           string   `json:"photo_id,omitempty"`            // 视频ID
	ImageToken        string   `json:"image_token,omitempty"`         // 封面图片token，不传则使用视频首帧
	ImageTokens       []string `json:"image_tokens,omitempty"`        // 便利贴/图片/小图图片token；组图需上传3张
	SplashPhotoIds    []string `json:"splash_photo_ids,omitempty"`    // 开屏视频ID，creative_material_type=11时必填，需4条不同尺寸
	SplashImageTokens []string `json:"splash_image_tokens,omitempty"` // 开屏图片token，creative_material_type=12时必填，需6张不同尺寸
	DpaTemplateId     int64    `json:"dpa_template_id,omitempty"`     // DPA模板ID，通过DPA模板信息接口获取

	// 创意分类与标签
	CreativeCategory int      `json:"creative_category,omitempty"` // 创意分类，须是叶子节点，与创意标签同时传或同时不传
	CreativeTag      []string `json:"creative_tag,omitempty"`      // 创意标签，最多20个，每个不超过10字符，与创意分类同时传或同时不传

	// 广告标签
	NewExposeTag []CreativeCreateNewExposeTag `json:"new_expose_tag,omitempty"` // 广告标签2期，最多2个，直播直投不用填写

	// 监测链接
	ClickTrackUrl       string `json:"click_track_url,omitempty"`         // 第三方点击检测链接，优化目标为激活时必填
	ImpressionUrl       string `json:"impression_url,omitempty"`          // 第三方开始播放监测链接，优化目标为激活时必填
	ActionbarClickUrl   string `json:"actionbar_click_url,omitempty"`     // 第三方点击按钮监测链接
	AdPhotoPlayedT3sUrl string `json:"ad_photo_played_t3s_url,omitempty"` // 第三方有效播放监测链接，白名单功能；与impression_url不可同时使用
	LiveTrackUrl        string `json:"live_track_url,omitempty"`          // 点击监测链接，campaign_type=16时可填

	// 其他
	SiteId                      int64 `json:"site_id,omitempty"`                       // 建站ID
	MicroChangeSwitch           int   `json:"micro_change_switch,omitempty"`           // 微改白盒化开关：0=关闭 1=打开
	MaterialIntelligentOptimize int   `json:"material_intelligent_optimize,omitempty"` // 素材智能优化开关：0=关闭 1=开启，仅白名单用户可用
}

func (receiver *CreativeUpdateReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *CreativeUpdateReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.CreativeId <= 0 {
		err = errors.New("creative_id is empty")
		return
	}
	return
}

// CreativeUpdateResp 修改自定义创意响应数据（仅data部分）
type CreativeUpdateResp struct {
	CreativeId int64 `json:"creative_id"` // 创意ID
}
