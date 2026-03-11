package errors

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidConfig 无效配置
	ErrInvalidConfig = errors.New("invalid config")

	// ErrInvalidRequest 无效请求
	ErrInvalidRequest = errors.New("invalid request")

	// ErrUnauthorized 未授权
	ErrUnauthorized = errors.New("unauthorized")

	// ErrTokenExpired token过期
	ErrTokenExpired = errors.New("token expired")

	// ErrRateLimit 超过限流
	ErrRateLimit = errors.New("rate limit exceeded")

	// ErrNetworkError 网络错误
	ErrNetworkError = errors.New("network error")

	// ErrInvalidResponse 无效响应
	ErrInvalidResponse = errors.New("invalid response")

	// ErrMediaNotFound 媒体不存在
	ErrMediaNotFound = errors.New("media not found")
)

// SDKError SDK错误
type SDKError struct {
	Code    int
	Message string
	Err     error
}

func (e *SDKError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("code=%d, message=%s, err=%v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("code=%d, message=%s", e.Code, e.Message)
}

// NewSDKError 创建SDK错误
func NewSDKError(code int, message string, err error) *SDKError {
	return &SDKError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
