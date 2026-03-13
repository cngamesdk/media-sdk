package config

import "time"

// MediaType 媒体类型
type MediaType string

const (
	MediaToutiao  MediaType = "toutiao"  // 巨量引擎
	MediaTencent  MediaType = "tencent"  // 腾讯广告
	MediaKuaiShou MediaType = "kuaisou"  // 磁力引擎（快手）
	MediaBaidu    MediaType = "baidu"    // 百度
	MediaUC       MediaType = "uc"       // UC
	MediaIQIYI    MediaType = "iqiyi"    // 爱奇艺
	MediaBilibili MediaType = "bilibili" // 哔哩哔哩
)

// Config 通用配置
type Config struct {
	Timeout    time.Duration `json:"timeout"`
	Proxy      string        `json:"proxy"`
	RateLimit  int           `json:"rate_limit"` // QPS限制
	MaxRetries int           `json:"max_retries"`
	RetryWait  time.Duration `json:"retry_wait"`
	Debug      bool          `json:"debug"`
}

// DefaultConfig 默认配置
func DefaultConfig() *Config {
	cfg := &Config{
		Timeout:    30 * time.Second,
		RateLimit:  10,
		MaxRetries: 3,
		RetryWait:  1 * time.Second,
		Debug:      false,
	}
	return cfg
}
