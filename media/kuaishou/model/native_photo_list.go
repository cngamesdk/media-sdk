package model

import "errors"

// NativePhotoListReq 获取达人原生视频列表请求
type NativePhotoListReq struct {
	accessTokenReq
	AdvertiserId      int64  `json:"advertiser_id"`       // 广告主id，必填
	AuthorId          int64  `json:"author_id"`           // 达人用户id，必填
	Count             int    `json:"count"`               // 每次获取个数，最大50，必填
	KolUserType       int    `json:"kol_user_type"`       // 达人类型：1-普通，2-服务号，3-聚星达人，必填
	Pcursor           string `json:"pcursor"`             // 游标，首次不传，后续传返回值
	TabType           int    `json:"tab_type"`            // 0-非隐藏视频，1-隐藏视频
	CampaignType      int    `json:"campaign_type"`       // 计划类型
	Keyword           string `json:"keyword"`             // 搜索关键词（搜索场景必填）
	SearchKeywordType int    `json:"search_keyword_type"` // 关键词类型：0-未知，1-photoId，2-关键词
}

func (receiver *NativePhotoListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *NativePhotoListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.AuthorId <= 0 {
		err = errors.New("author_id is empty")
		return
	}
	if receiver.Count <= 0 || receiver.Count > 50 {
		err = errors.New("count must be 1-50")
		return
	}
	if receiver.KolUserType < 1 || receiver.KolUserType > 3 {
		err = errors.New("kol_user_type must be 1, 2 or 3")
		return
	}
	return
}

// CdnUrlInfo CDN URL信息
type CdnUrlInfo struct {
	Cdn string `json:"cdn"` // CDN信息
	Url string `json:"url"` // URL信息
}

// NativePhotoItem 达人原生视频条目
type NativePhotoItem struct {
	PhotoId              string       `json:"photo_id"`               // 加密后的photoId
	Caption              string       `json:"caption"`                // 视频标题
	CoverUrl             []CdnUrlInfo `json:"cover_url"`              // 封面url
	MovieUrl             []CdnUrlInfo `json:"movie_url"`              // 视频url
	Duration             int64        `json:"duration"`               // 视频时长，毫秒
	Height               int          `json:"height"`                 // 视频高度
	Width                int          `json:"width"`                  // 视频宽度
	CreativeMaterialType int          `json:"creative_material_type"` // 视频横竖版
	AdSocialOrderId      string       `json:"ad_social_order_id"`     // 聚星订单id
	UploadTimeMills      int64        `json:"upload_time_mills"`      // 视频上传时间，毫秒
	Signature            string       `json:"signature"`              // 视频md5（无用字段）
	NativeGoodType       int          `json:"native_good_type"`       // 素材质量：1-良好，2-优质
	PhotoDupStatus       int          `json:"photo_dup_status"`       // 素材创新度：0-原创，1-重复
}

// NativePhotoListResp 获取达人原生视频列表响应数据（仅data部分）
type NativePhotoListResp struct {
	Photos  []NativePhotoItem `json:"photos"`  // 视频列表
	Pcursor string            `json:"pcursor"` // 下次游标，"no_more"=无更多
}
