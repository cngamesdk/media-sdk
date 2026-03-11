package model

import "encoding/json"

// UnitReq 广告组请求
type UnitReq struct {
	ID           string          `json:"id,omitempty"`
	CampaignID   string          `json:"campaign_id"`
	AdvertiserID string          `json:"advertiser_id"`
	Name         string          `json:"name"`
	Pricing      string          `json:"pricing"` // CPM, CPC, OCPM
	BidAmount    float64         `json:"bid_amount"`
	DailyBudget  float64         `json:"daily_budget"`
	Status       string          `json:"status"`
	Target       *Targeting      `json:"targeting"`
	Extra        json.RawMessage `json:"extra,omitempty"`
}

// Targeting 定向
type Targeting struct {
	Gender    []string               `json:"gender"`
	Age       []string               `json:"age"`
	Region    []string               `json:"region"`
	Interests []string               `json:"interests"`
	Device    []string               `json:"device"`
	Network   []string               `json:"network"`
	Os        []string               `json:"os"`
	Custom    map[string]interface{} `json:"custom,omitempty"`
}

// UnitResp 广告组响应
type UnitResp struct {
}
