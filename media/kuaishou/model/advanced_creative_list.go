package model

import "errors"

// AdvancedCreativeListReq 查询程序化创意请求
type AdvancedCreativeListReq struct {
	accessTokenReq
	AdvertiserId   int64    `json:"advertiser_id"`              // 广告主ID，必填
	EndDate        string   `json:"end_date,omitempty"`         // 结束日期，格式：yyyy-MM-dd
	PackageName    string   `json:"package_name,omitempty"`     // 程序化创意包名称，0-100字符
	Page           int      `json:"page,omitempty"`             // 页数，默认1
	PageSize       int      `json:"page_size,omitempty"`        // 每页行数，默认20
	PutStatusList  []string `json:"put_status_list,omitempty"`  // 创意投放状态：1=投放 2=暂停 3=删除；传则覆盖status参数
	StartDate      string   `json:"start_date,omitempty"`       // 起始日期，格式：yyyy-MM-dd
	Status         int      `json:"status,omitempty"`           // 程序化创意状态：-2=所有（含已删除）、40=只含已删除，不传=所有（不含已删除）
	TimeFilterType int      `json:"time_filter_type,omitempty"` // 时间筛选类型：1=按创建时间筛选，0/不传=按更新时间筛选
	UnitIds        []string `json:"unit_ids,omitempty"`         // 广告组ID集，不超过100个
}

func (receiver *AdvancedCreativeListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AdvancedCreativeListReq) Validate() (err error) {
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

// AdvancedCreativeListCreative 程序化创意ID信息
type AdvancedCreativeListCreative struct {
	CreativeId int64 `json:"creative_id"` // 创意ID
	UnitId     int64 `json:"unit_id"`     // 广告组ID
}

// AdvancedCreativeListNewExposeTag 广告标签2期
type AdvancedCreativeListNewExposeTag struct {
	Text string `json:"text"` // 广告标签文字
	Url  string `json:"url"`  // 广告标签链接
}

// AdvancedCreativeListPhotoItem 素材列表项
type AdvancedCreativeListPhotoItem struct {
	CoverImageToken      string   `json:"cover_image_token"`      // 封面图片token
	CoverImageUrl        string   `json:"cover_image_url"`        // 封面图片URL
	CreativeMaterialType int      `json:"creative_material_type"` // 素材类型：1=竖版视频 2=横版视频
	PhotoId              int64    `json:"photo_id"`               // 视频ID
	PicIds               []string `json:"pic_ids"`                // 图片库图片ID
	PicList              []string `json:"pic_list"`               // 联盟图片image_token（横版/竖版）
	PicUrlList           []string `json:"pic_url_list"`           // 联盟图片url（横版/竖版）
}

// AdvancedCreativeDetail 程序化创意详情
type AdvancedCreativeDetail struct {
	ActionBar                   string                             `json:"action_bar"`                    // 行动号召按钮
	ActionbarClickUrl           string                             `json:"actionbar_click_url"`           // 第三方ActionBar点击监控链接
	AdPhotoPlayedT3sUrl         string                             `json:"ad_photo_played_t3s_url"`       // 第三方有效播放监测链接
	AppGradeType                int                                `json:"app_grade_type"`                // 审核分级类型：0=默认 1=审核降级
	Captions                    []string                           `json:"captions"`                      // 作品广告语，1-3个
	ClickUrl                    string                             `json:"click_url"`                     // 第三方点击检测链接
	CoverImageTokens            []string                           `json:"cover_image_tokens"`            // 封面image_token，1-4个
	CoverImageUrls              []string                           `json:"cover_image_urls"`              // 封面链接地址
	CoverSlogans                []string                           `json:"cover_slogans"`                 // 封面广告语
	CreateTime                  string                             `json:"create_time"`                   // 创建时间，格式：2019-06-11 15:17:25
	CreativeCategory            int                                `json:"creative_category"`             // 创意分类
	CreativeTag                 []string                           `json:"creative_tag"`                  // 创意标签
	Creatives                   []AdvancedCreativeListCreative     `json:"creatives"`                     // 生成的程序化创意ID列表
	HorizontalPhotoIds          []string                           `json:"horizontal_photo_ids"`          // 横版视频id列表
	ImpressionUrl               string                             `json:"impression_url"`                // 第三方曝光检测链接
	KolUserId                   int64                              `json:"kol_user_id"`                   // 开启原生时的达人ID
	KolUserType                 int                                `json:"kol_user_type"`                 // 达人类型：2=服务号原生 3=聚星达人原生
	MaterialIntelligentOptimize int                                `json:"material_intelligent_optimize"` // 素材智能优化开关：0=关闭 1=开启
	MicroChangeSwitch           int                                `json:"micro_change_switch"`           // 微改白盒化开关：0=关闭 1=打开
	NewExposeTag                []AdvancedCreativeListNewExposeTag `json:"new_expose_tag"`                // 广告标签2期
	OpenAccountNative           int                                `json:"open_account_native"`           // 是否为原生扩量：0=否 1=是
	OuterLoopNative             int                                `json:"outer_loop_native"`             // 是否开启原生：0=关闭 1=开启
	PackageName                 string                             `json:"package_name"`                  // 程序化创意包名称
	PhotoList                   []AdvancedCreativeListPhotoItem    `json:"photo_list"`                    // 素材列表
	PutStatus                   int                                `json:"put_status"`                    // 操作状态：1=投放 2=暂停 3=删除
	Recommendation              string                             `json:"recommendation"`                // 开启原生时的plc描述语
	StickerStyles               []string                           `json:"sticker_styles"`                // 封面贴纸
	UnitId                      int64                              `json:"unit_id"`                       // 广告组ID
	UpdateTime                  string                             `json:"update_time"`                   // 更新时间，格式：2019-06-11 15:17:25
	VerticalPhotoIds            []string                           `json:"vertical_photo_ids"`            // 竖版视频id列表
	ViewStatus                  int                                `json:"view_status"`                   // 程序化创意状态：-1=不限 1=计划已暂停 3=计划超预算 5=计划已删除 6=余额不足 11=组审核中 12=组审核未通过 14=已结束 15=组已暂停 17=组超预算 19=未达投放时间 22=不在投放时段 40=创意已删除 41=审核中 42=审核未通过 46=已暂停 52=投放中 53=作品异常 55=部分素材审核通过 56=部分审核失败 62=待送审
	ViewStatusReason            string                             `json:"view_status_reason"`            // 程序化创意状态描述
}

// AdvancedCreativeListResp 查询程序化创意响应数据（仅data部分）
type AdvancedCreativeListResp struct {
	TotalCount int64                    `json:"total_count"` // 总数
	Details    []AdvancedCreativeDetail `json:"details"`     // 详情列表
}
