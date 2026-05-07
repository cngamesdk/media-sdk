package model

import "errors"

// ImageRelateCreativesReq 查询图片关联创意请求
type ImageRelateCreativesReq struct {
	accessTokenReq
	AdvertiserId int64  `json:"advertiser_id"` // 广告主ID，必填
	ImageToken   string `json:"image_token"`   // 图片token，必填
}

func (receiver *ImageRelateCreativesReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *ImageRelateCreativesReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.ImageToken == "" {
		err = errors.New("image_token is empty")
		return
	}
	return
}

// CreativesItem 创意
type CreativesItem struct {
	CreativeId   int64  `json:"creative_id"`   // 创意ID
	CreativeName string `json:"creative_name"` // 创意名称
}

// ImageRelateCreativesResp 查询图片关联创意响应数据（仅data部分）
type ImageRelateCreativesResp struct {
	CreativeCount int             `json:"creative_count"`    // 绑定创意个数
	ImageToken    string          `json:"image_token"`       // 图片token
	Creatives     []CreativesItem `json:"related_creatives"` // 创意列表
}
