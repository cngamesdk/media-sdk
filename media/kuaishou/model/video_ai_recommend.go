package model

import "errors"

// VideoAiRecommendReq 获取AI推荐视频请求
type VideoAiRecommendReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"`            // 广告主ID，必填
	Page         int   `json:"page,omitempty"`           // 分页
	PageSize     int   `json:"page_size,omitempty"`      // 分页大小，最大100
	DpaProductId int64 `json:"dpa_product_id,omitempty"` // sDpaId
	SeriesId     int64 `json:"series_id,omitempty"`      // 短剧ID
}

func (receiver *VideoAiRecommendReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *VideoAiRecommendReq) Validate() (err error) {
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

// VideoAiRecommendResp 获取AI推荐视频响应数据（仅data部分）
type VideoAiRecommendResp struct {
	TotalCount int64                    `json:"total_count"` // 视频总数
	Details    []VideoAiRecommendDetail `json:"details"`     // 视频详情列表
}

// VideoAiRecommendDetail AI推荐视频详情（在VideoCursorDetail基础上增加native_good_type字段）
type VideoAiRecommendDetail struct {
	VideoCursorDetail
	NativeGoodType int `json:"native_good_type"` // 素材质量：1良好 2优质 0/3其他
}
