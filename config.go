package media_sdk

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
	MediaType    MediaType     `json:"media_type"`
	AppID        string        `json:"app_id"`
	AppSecret    string        `json:"app_secret"`
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	ExpireTime   time.Time     `json:"expire_time"`
	BaseURL      string        `json:"base_url"`
	Timeout      time.Duration `json:"timeout"`
	Proxy        string        `json:"proxy"`
	RateLimit    int           `json:"rate_limit"` // QPS限制
	MaxRetries   int           `json:"max_retries"`
	RetryWait    time.Duration `json:"retry_wait"`
	Debug        bool          `json:"debug"`
}

// DefaultConfig 默认配置
func DefaultConfig(mediaType MediaType) *Config {
	cfg := &Config{
		MediaType:  mediaType,
		Timeout:    30 * time.Second,
		RateLimit:  10,
		MaxRetries: 3,
		RetryWait:  1 * time.Second,
		Debug:      false,
	}

	// 设置各媒体默认域名
	switch mediaType {
	case MediaToutiao:
		cfg.BaseURL = "https://api.oceanengine.com"
	case MediaTencent:
		cfg.BaseURL = "https://api.e.qq.com"
	case MediaKuaiShou:
		cfg.BaseURL = "https://ad.e.kuaishou.com"
	case MediaBaidu:
		cfg.BaseURL = "https://api.baidu.com"
	case MediaUC:
		cfg.BaseURL = "https://e.uc.cn"
	case MediaIQIYI:
		cfg.BaseURL = "https://api.iqiyi.com"
	case MediaBilibili:
		cfg.BaseURL = "https://api.bilibili.com"
	}

	return cfg
}
