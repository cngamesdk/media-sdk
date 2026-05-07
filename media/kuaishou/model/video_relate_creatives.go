package model

import "errors"

// VideoRelateCreativesReq 视频关联创意数查询请求
type VideoRelateCreativesReq struct {
	accessTokenReq
	AdvertiserId int64    `json:"advertiser_id"` // 广告主ID，必填
	PhotoIds     []string `json:"photo_ids"`     // 视频id列表，必填，最多20个
}

func (receiver *VideoRelateCreativesReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *VideoRelateCreativesReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if len(receiver.PhotoIds) == 0 {
		err = errors.New("photo_ids is empty")
		return
	}
	return
}

// VideoRelateCreativesResp 视频关联创意数查询响应数据（仅data部分）
type VideoRelateCreativesResp struct {
	RelatedCreatives []VideoRelateCreativesItem `json:"related_creatives"` // 与视频相关联的创意信息
}

// VideoRelateCreativesItem 单个视频关联创意信息
type VideoRelateCreativesItem struct {
	PhotoId               string                    `json:"photo_id"`                // 视频ID
	CreativeCount         int64                     `json:"creative_count"`          // 自定义创意数量
	AdvancedCreativeCount int64                     `json:"advanced_creative_count"` // 程序化创意数量
	AdvancedCreativeIds   []int64                   `json:"advanced_creative_ids"`   // 程序化创意关联的unitid
	SmartCreativeCount    int64                     `json:"smart_creative_count"`    // 关联的智能创意数量
	SmartCreativeIds      []int64                   `json:"smart_creative_ids"`      // 智能创意关联的unitid
	Creatives             []VideoRelateCreativeItem `json:"creatives"`               // 与视频关联的自定义创意信息
}

// VideoRelateCreativeItem 自定义创意信息
type VideoRelateCreativeItem struct {
	CreativeId   int64  `json:"creative_id"`   // 自定义创意ID
	CreativeName string `json:"creative_name"` // 创意名称
}
