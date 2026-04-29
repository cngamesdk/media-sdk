package model

import "errors"

// AdvCardListReq 获取高级创意列表请求
type AdvCardListReq struct {
	accessTokenReq
	AdvertiserId int64 `json:"advertiser_id"`        // 广告主ID，必填
	CardType     int   `json:"card_type"`            // 卡片类型，必填：100=图片卡片 101=多利益卡-图文 102=多利益卡-多标签 103=电商促销样式 107=倒计时卡 108=优惠券卡 131=商品卡-汽车样式 132=商品卡-电商样式 133=商品卡-小说样式 134=商品卡-房产样式 200=推广位
	StyleType    int   `json:"style_type,omitempty"` // 高级创意样式类型：1=普通卡片 2=推广位
	Page         int   `json:"page,omitempty"`       // 查询页码，默认1
	PageSize     int   `json:"page_size,omitempty"`  // 单页行数，默认10，最大200
}

func (receiver *AdvCardListReq) Format() {
	receiver.accessTokenReq.Format()
}

func (receiver *AdvCardListReq) Validate() (err error) {
	if validateErr := receiver.accessTokenReq.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	if receiver.AdvertiserId <= 0 {
		err = errors.New("advertiser_id is empty")
		return
	}
	if receiver.CardType <= 0 {
		err = errors.New("card_type is empty")
		return
	}
	return
}

// SdpaCardContent SDPA商品卡样式内容（card_type=131/132/133/134专用）
type SdpaCardContent struct {
	Title           string  `json:"title"`            // 商品卡标题
	SubTitle        string  `json:"sub_title"`        // 商品卡副标题（描述/简介）
	Icon            string  `json:"icon"`             // 商品卡图片
	Description     string  `json:"description"`      // 商品卡价格类型描述
	Symbol          string  `json:"symbol"`           // 商品卡价格符号描述
	Price           float64 `json:"price"`            // 商品卡价格（现价）
	TailDescription string  `json:"tail_description"` // 商品卡价格单位描述
	SalePrice       float64 `json:"sale_price"`       // 商品卡价格（原价）
	TagText         string  `json:"tag_text"`         // 商品卡角标
	HotText         string  `json:"hot_text"`         // 商品卡热度标签
	City            string  `json:"city"`             // 商品卡城市标签
}

// AdvCardDetail 高级创意详情
type AdvCardDetail struct {
	AdvCardId       int64            `json:"adv_card_id"`       // 卡片ID
	TemplateName    string           `json:"template_name"`     // 模版名
	UnitCount       int              `json:"unit_count"`        // 关联广告组数
	Url             string           `json:"url"`               // 图片URL
	Title           string           `json:"title"`             // 标题
	SubTitle        string           `json:"sub_title"`         // 副标题
	Price           int              `json:"price"`             // 原价格（单位：分）
	SalePrice       int              `json:"sale_price"`        // 售卖价（单位：分）
	BeginTime       int              `json:"begin_time"`        // 倒计时卡开始时间（card_type=107）
	EndTime         int              `json:"end_time"`          // 倒计时卡结束时间（card_type=107）
	SdpaCardContent *SdpaCardContent `json:"sdpa_card_content"` // SDPA商品卡样式内容（card_type=131/132/133/134）
}

// AdvCardListResp 获取高级创意列表响应数据（仅data部分）
type AdvCardListResp struct {
	TotalCount int             `json:"total_count"` // 卡片总数
	Details    []AdvCardDetail `json:"details"`     // 数据详情
}
