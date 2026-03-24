package model

const (
	DevelopersUrl = "https://developers.e.qq.com"
	ApiUrl        = "https://api.e.qq.com"
)

type BaseResp struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	MessageCn string      `json:"message_cn"`
	Data      interface{} `json:"data"`
	RequestId string      `json:"request_id"`
}
