package model

import "errors"

// CreativeListReq 查询自定义创意请求
type CreativeListReq struct {
	accessTokenReq
	AdvertiserId   int64    `json:"advertiser_id"`              // 广告主ID，必填
	CampaignId     int64    `json:"campaign_id,omitempty"`      // 广告计划ID，筛选条件，空=不限
	UnitId         int64    `json:"unit_id,omitempty"`          // 广告组ID，筛选条件，空=不限
	CreativeId     int64    `json:"creative_id,omitempty"`      // 广告创意ID，筛选条件，空=不限
	CreativeIds    []string `json:"creative_ids,omitempty"`     // 广告创意ID集，最多100个
	CreativeName   string   `json:"creative_name,omitempty"`    // 广告创意名称，支持模糊搜索精确查询
	PutStatusList  []string `json:"put_status_list,omitempty"`  // 创意投放状态：1=投放 2=暂停 3=删除；传则覆盖status参数
	Status         int      `json:"status,omitempty"`           // 广告创意状态：-2=不限 40=只含已删除 不传=所有不含已删除
	StartDate      string   `json:"start_date,omitempty"`       // 开始时间，格式：yyyy-MM-dd，需与end_date同时传
	EndDate        string   `json:"end_date,omitempty"`         // 结束时间，格式：yyyy-MM-dd，需与start_date同时传
	TimeFilterType int      `json:"time_filter_type,omitempty"` // 时间过滤类型：0/不传=更新时间 1=创建时间
	Page           int      `json:"page,omitempty"`             // 页码，默认1
	PageSize       int      `json:"page_size,omitempty"`        // 每页数量，默认20
}

func (receiver *CreativeListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *CreativeListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	return
}

// CreativeListDisplayInfo 广告展示信息
type CreativeListDisplayInfo struct {
	ActionBarText string   `json:"action_bar_text"` // 行动号召按钮文案
	Description   string   `json:"description"`     // 广告语
	DpaStyleTypes []string `json:"dpa_style_types"` // 动态商品卡样式，14001=区域服务卡
}

// CreativeListSplashPhoto 开屏视频信息
type CreativeListSplashPhoto struct {
	PhotoId  string `json:"photo_id"`  // 视频ID
	Width    int    `json:"width"`     // 视频宽度
	Height   int    `json:"height"`    // 视频高度
	PhotoMd5 string `json:"photo_md5"` // 视频MD5
}

// CreativeListSplashPicture 开屏图片信息
type CreativeListSplashPicture struct {
	CoverId   int64  `json:"cover_id"`    // 封面ID
	CoverUrl  string `json:"cover_url"`   // 封面URL
	Width     int    `json:"width"`       // 图片宽度
	Height    int    `json:"height"`      // 图片高度
	KyPhotoId int64  `json:"ky_photo_id"` // 内部图片ID
}

// CreativeDetail 自定义创意详情
type CreativeDetail struct {
	CreativeId           int64    `json:"creative_id"`            // 广告创意ID
	CreativeName         string   `json:"creative_name"`          // 广告创意名称
	CreativeMaterialType int      `json:"creative_material_type"` // 素材类型：0=历史未区分 1=竖版视频 2=横版视频 11=开屏视频 12=开屏图片
	CreativeCategory     int      `json:"creative_category"`      // 创意分类
	CreativeTag          []string `json:"creative_tag"`           // 创意标签
	CreativeMode         int      `json:"creative_mode"`          // 创意生成模式：0=普通 1=AIGC生成
	AdType               int      `json:"ad_type"`                // 广告计划类型：0=信息流 1=搜索
	CampaignId           int64    `json:"campaign_id"`            // 广告计划ID
	UnitId               int64    `json:"unit_id"`                // 广告组ID
	Status               int      `json:"status"`                 // 广告创意状态：-1=不限 1=计划已暂停 3=计划超预算 6=余额不足 11=组审核中 12=组审核未通过 14=已结束 15=组已暂停 17=组超预算 19=未达投放时间 22=不在投放时段 40=已删除 41=审核中 42=审核未通过 46=已暂停 52=投放中 53=作品异常 54=视频审核通过可投放滑滑 55=部分素材审核失败
	PutStatus            int      `json:"put_status"`             // 投放状态：1=投放中 2=暂停 3=删除
	AppGradeType         int      `json:"app_grade_type"`         // 审核分级类型：0=默认 1=审核降级
	AutoDeliverType      int      `json:"auto_deliver_type"`      // 自动投放类型
	DspVersion           int      `json:"dsp_version"`            // DSP版本
	CreateTime           string   `json:"create_time"`            // 创建时间，格式：yyyy-MM-dd HH:mm:ss
	UpdateTime           string   `json:"update_time"`            // 最后修改时间，格式：yyyy-MM-dd HH:mm:ss

	// 素材
	PhotoId             string                      `json:"photo_id"`               // 视频作品ID
	PhotoMd5            string                      `json:"photo_md5"`              // 视频MD5
	CoverUrl            string                      `json:"cover_url"`              // 封面URL
	CoverWidth          int64                       `json:"cover_width"`            // 封面图宽度
	CoverHeight         int64                       `json:"cover_height"`           // 封面图高度
	ImageToken          string                      `json:"image_token"`            // 视频封面token
	ImageTokens         []string                    `json:"image_tokens"`           // 单图创意image_token列表
	MaterialUrl         []string                    `json:"material_url"`           // 单图创意URL列表
	PicId               string                      `json:"pic_id"`                 // 图片库图片ID
	OverlayType         string                      `json:"overlay_type"`           // 贴纸样式类型
	OverlayBgImageToken string                      `json:"overlay_bg_image_token"` // 动态词包原始封面图片token
	OverlayBgUrl        string                      `json:"overlay_bg_url"`         // 动态词包原始封面图片URL
	StickerTitle        string                      `json:"sticker_title"`          // 封面广告语标题
	ShortSlogan         string                      `json:"short_slogan"`           // 便利贴创意短广告语
	DpaTemplateId       int64                       `json:"dpa_template_id"`        // DPA模板ID
	SplashPhotos        []CreativeListSplashPhoto   `json:"splash_photos"`          // 开屏视频信息，creative_material_type=11时
	SplashPictures      []CreativeListSplashPicture `json:"splash_pictures"`        // 开屏图片信息，creative_material_type=12时

	// 展示信息
	DisplayInfo  *CreativeListDisplayInfo     `json:"display_info"`   // 广告展示信息
	ExposeTag    string                       `json:"expose_tag"`     // 广告标签
	NewExposeTag []CreativeCreateNewExposeTag `json:"new_expose_tag"` // 广告标签2期

	// 监测链接
	ClickTrackUrl       string `json:"click_track_url"`         // 点击监测链接
	ImpressionUrl       string `json:"impression_url"`          // 第三方开始播放监测链接
	ActionbarClickUrl   string `json:"actionbar_click_url"`     // 第三方点击按钮监测链接
	AdPhotoPlayedT3sUrl string `json:"ad_photo_played_t3s_url"` // 第三方有效播放监测链接
	LiveTrackUrl        string `json:"live_track_url"`          // 点击监测链接，campaign_type=16时

	// 审核
	ReviewDetail        string   `json:"review_detail"`         // 审核拒绝理由
	RejectVideoSnapshot []string `json:"reject_video_snapshot"` // 审核拒绝图片列表

	// 原生投放
	OuterLoopNative   int    `json:"outer_loop_native"`   // 是否开启原生：0=关闭 1=开启
	KolUserId         int64  `json:"kol_user_id"`         // 原生投放目标达人ID
	KolUserType       int    `json:"kol_user_type"`       // 原生达人用户类型：0=未开启 2=服务号 3=聚星达人
	OpenAccountNative int    `json:"open_account_native"` // 是否为原生扩量：0=否 1=是
	Recommendation    string `json:"recommendation"`      // PLC自定义文案

	// 直播
	LiveCreativeType int `json:"live_creative_type"` // 直播类型：3=直投直播 4=作品引流

	// 其他
	HighLightFlash              int    `json:"high_light_flash"`              // 高光创意状态：0=关闭 1=开启
	MicroChangeSwitch           int    `json:"micro_change_switch"`           // 微改白盒化开关：0=关闭 1=打开
	MaterialIntelligentOptimize int    `json:"material_intelligent_optimize"` // 素材智能优化开关：0=关闭 1=开启
	MerchantLibraryId           int64  `json:"merchant_library_id"`           // 商品库ID
	MerchantProductId           string `json:"merchant_product_id"`           // 商品ID
}

// CreativeListResp 查询自定义创意响应数据（仅data部分）
type CreativeListResp struct {
	TotalCount int64            `json:"total_count"` // 创意总数
	Details    []CreativeDetail `json:"details"`     // 创意详情列表
}
