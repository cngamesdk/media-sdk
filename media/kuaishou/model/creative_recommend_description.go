package model

// CreativeRecommendDescriptionReq 获取推荐广告语请求
type CreativeRecommendDescriptionReq struct {
	accessTokenReq
	AppId        int64 `json:"app_id,omitempty"`         // 应用ID
	CourseId     int64 `json:"course_id,omitempty"`      // 付费课堂课程ID
	DpaProductId int64 `json:"dpa_product_id,omitempty"` // 商品库产品ID
	KwaiBookId   int64 `json:"kwai_book_id,omitempty"`   // 小说ID
	SeriesId     int64 `json:"series_id,omitempty"`      // 短剧ID
}

func (receiver *CreativeRecommendDescriptionReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *CreativeRecommendDescriptionReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	return
}

// CreativeRecommendDescriptionResp 获取推荐广告语响应数据（仅data部分）
type CreativeRecommendDescriptionResp struct {
	RecommendDescription []string `json:"recommend_description"` // 广告语列表
	Count                int      `json:"count"`                 // 广告语数量
}
