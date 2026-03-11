package model

import (
	"time"
)

// AccountReq 账户请求
type AccountReq struct {
	AdvertiserID string `json:"advertiser_id"`
}

// AccountResp 账户响应
type AccountResp struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Balance      float64   `json:"balance"`
	Status       string    `json:"status"`
	Rechargeable float64   `json:"rechargeable"`
	CreateTime   time.Time `json:"create_time"`
}
