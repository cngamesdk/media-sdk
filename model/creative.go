package model

import "encoding/json"

// CreativeReq 广告创意请求
type CreativeReq struct {
	ID           string          `json:"id,omitempty"`
	UnitID       string          `json:"unit_id"`
	AdvertiserID string          `json:"advertiser_id"`
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	ImageURL     string          `json:"image_url"`
	VideoURL     string          `json:"video_url,omitempty"`
	Destination  string          `json:"destination"`
	CallToAction string          `json:"call_to_action"`
	Status       string          `json:"status"`
	Extra        json.RawMessage `json:"extra,omitempty"`
}

// CreativeResp 广告创意响应
type CreativeResp struct {
}

type GetCreativeReq struct {
}

type GetCreativeResp struct {
}

type ListCreativesReq struct {
}

type ListCreativesResp struct {
}
