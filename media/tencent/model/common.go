package model

import (
	"errors"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/random"
	"time"
)

const (
	DevelopersUrl = "https://developers.e.qq.com"
	ApiUrl        = "https://api.e.qq.com"
)

// GlobalReq 全局参数
type GlobalReq struct {
	AccessToken string `json:"access_token"` // 授权令牌 (必填)
	Timestamp   int64  `json:"timestamp"`    // 当前时间戳，单位为秒 (必填)
	Nonce       string `json:"nonce"`        // 随机字符串标识，不超过32个字符 (必填)
}

// 常量定义
const (
	MaxTimestampDiff = 300         // 最大时间误差（秒）
	MaxNonceLength   = 32          // 随机字符串最大长度
	TimezoneOffset   = 8 * 60 * 60 // GMT+8时区偏移（秒）
)

func (p *GlobalReq) Format() {
	if p.Nonce == "" {
		uuid, _ := random.UUIdV4()
		p.Nonce = cryptor.Md5String(uuid)
	}
}

// Validate 验证API请求公共参数
func (p *GlobalReq) Validate() error {
	// 1. 验证access_token
	if p.AccessToken == "" {
		return errors.New("access_token为必填")
	}

	// 2. 验证timestamp
	if p.Timestamp == 0 {
		return errors.New("timestamp为必填")
	}
	if err := validateTimestamp(p.Timestamp); err != nil {
		return err
	}

	// 3. 验证nonce
	if p.Nonce == "" {
		return errors.New("nonce为必填")
	}
	if len(p.Nonce) > MaxNonceLength {
		return errors.New("nonce长度不能超过32个字符")
	}

	return nil
}

// validateTimestamp 验证时间戳
func validateTimestamp(timestamp int64) error {
	now := time.Now().Unix()
	diff := now - timestamp
	if diff < 0 {
		diff = -diff
	}
	if diff > MaxTimestampDiff {
		return errors.New("timestamp与服务器时间误差超过300秒")
	}
	return nil
}

// TimestampToTime 将时间戳转换为GMT+8时区的时间
func TimestampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0).In(time.FixedZone("GMT+8", TimezoneOffset))
}

// GetCurrentTimestamp 获取当前GMT+8时区的秒级时间戳
func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

type BaseResp struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	MessageCn string      `json:"message_cn"`
	Data      interface{} `json:"data"`
	RequestId string      `json:"request_id"`
}
