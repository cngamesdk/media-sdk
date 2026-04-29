package model

import "errors"

// WordInfoPhotoRecommendReq 创意高相关词推荐请求
type WordInfoPhotoRecommendReq struct {
	accessTokenReq
	PhotoIds     []string `json:"photo_ids"`     // 素材ids，必填；单次查询数量不超过20
	AdvertiserId int64    `json:"advertiser_id"` // 账户id，必填
}

func (receiver *WordInfoPhotoRecommendReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *WordInfoPhotoRecommendReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if len(receiver.PhotoIds) == 0 {
		err = errors.New("photo_ids is empty")
		return
	}
	if len(receiver.PhotoIds) > 20 {
		err = errors.New("photo_ids max length is 20")
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	return
}

// PhotoRecommendWord 推荐关键词
type PhotoRecommendWord struct {
	Word      string `json:"word"`       // 关键词
	Pv        int64  `json:"pv"`         // pv
	RelaValue int    `json:"rela_value"` // 相关性分数，值为 1,2,3,4
	PhotoId   string `json:"photo_id"`   // 素材id
}

// WordInfoPhotoRecommendResp 创意高相关词推荐响应数据（仅data部分）
type WordInfoPhotoRecommendResp struct {
	TotalCount int64                `json:"total_count"` // 关键词数量
	Details    []PhotoRecommendWord `json:"details"`     // 推荐关键词列表
}
