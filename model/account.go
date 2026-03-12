package model

import (
	"time"
)

// AccountReq 账户请求
type AccountReq struct {
	AccessToken  string   `json:"access_token,omitempty"`
	AdvertiserID int64    `json:"advertiser_id,omitempty"`
	Fields       []string `json:"fields,omitempty"`
}

// AccountResp 账户响应
type AccountResp struct {
	AdvertiserID int64     `json:"advertiser_id"`
	Name         string    `json:"name"`
	Role         string    `json:"role"`
	Status       string    `json:"status"`
	CreateTime   time.Time `json:"create_time"`
	Extension
}
