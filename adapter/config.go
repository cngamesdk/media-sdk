package adapter

import "time"

type AdapterConfig struct {
	AccessToken  string
	AdvertiserId int64
	BaseURL      string        `json:"base_url"`
	Timeout      time.Duration `json:"timeout"`
}
