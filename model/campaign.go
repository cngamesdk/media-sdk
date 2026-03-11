package model

import (
	"encoding/json"
	"time"
)

// CampaignReq 广告计划请求
type CampaignReq struct {
	ID           string          `json:"id,omitempty"`
	AdvertiserID string          `json:"advertiser_id"`
	Name         string          `json:"name"`
	Budget       float64         `json:"budget"`
	BudgetMode   string          `json:"budget_mode"` // DAY, TOTAL
	Status       string          `json:"status"`
	StartTime    time.Time       `json:"start_time"`
	EndTime      time.Time       `json:"end_time,omitempty"`
	Extra        json.RawMessage `json:"extra,omitempty"`
}

// CampaignResp 广告计划响应
type CampaignResp struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	Budget     float64         `json:"budget"`
	BudgetMode string          `json:"budget_mode"`
	Status     string          `json:"status"`
	StartTime  time.Time       `json:"start_time"`
	EndTime    time.Time       `json:"end_time"`
	CreateTime time.Time       `json:"create_time"`
	UpdateTime time.Time       `json:"update_time"`
	Extra      json.RawMessage `json:"extra,omitempty"`
}
