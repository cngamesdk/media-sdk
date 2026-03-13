package model

const (
	BaseUrlOpen = "https://open.oceanengine.com"
	BaseUrlApi  = "https://api.oceanengine.com"
	BaseUrlAd   = "https://ad.oceanengine.com"
)

type BaseResp struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	RequestId string      `json:"request_id"`
}
