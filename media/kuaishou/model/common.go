package model

const (
	DevelopersUrl = "https://developers.e.kuaishou.com"
	AdUrl         = "https://ad.e.kuaishou.com"
)

type BaseResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
