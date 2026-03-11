package utils

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-resty/resty/v2"
	"golang.org/x/time/rate"
	"io"
	"net/http"
	"net/url"
	"time"
)

// HTTPConfig HTTP配置
type HTTPConfig struct {
	Timeout    time.Duration
	Proxy      string
	MaxRetries int
	RetryWait  time.Duration
	Debug      bool
	RateLimit  int // QPS
}

// HTTPClient HTTP客户端
type HTTPClient struct {
	client  *resty.Client
	limiter *rate.Limiter
	config  *HTTPConfig
}

// NewHTTPClient 创建HTTP客户端
func NewHTTPClient(config *HTTPConfig) *HTTPClient {
	client := resty.New()

	// 设置超时
	client.SetTimeout(config.Timeout)

	// 设置重试
	client.SetRetryCount(config.MaxRetries)
	client.SetRetryWaitTime(config.RetryWait)
	client.SetRetryMaxWaitTime(10 * time.Second)

	// 设置代理
	if config.Proxy != "" {
		proxyURL, err := url.Parse(config.Proxy)
		if err == nil {
			client.SetProxy(proxyURL.String())
		}
	}

	// 设置TLS
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	// 调试模式
	if config.Debug {
		client.SetDebug(true)
	}

	// 限流器
	var limiter *rate.Limiter
	if config.RateLimit > 0 {
		limiter = rate.NewLimiter(rate.Limit(config.RateLimit), config.RateLimit)
	}

	return &HTTPClient{
		client:  client,
		limiter: limiter,
		config:  config,
	}
}

// Request 发送请求
func (c *HTTPClient) Request(ctx context.Context, method, url string, body io.Reader, headers map[string]string) ([]byte, error) {
	// 限流
	if c.limiter != nil {
		err := c.limiter.Wait(ctx)
		if err != nil {
			return nil, fmt.Errorf("rate limit wait failed: %w", err)
		}
	}

	// 读取body
	var bodyBytes []byte
	if body != nil {
		bodyBytes, _ = io.ReadAll(body)
	}

	// 创建请求
	req := c.client.R().
		SetContext(ctx).
		SetHeaders(headers)

	if len(bodyBytes) > 0 {
		req.SetBody(bodyBytes)
	}

	// 发送请求
	resp, err := req.Execute(method, url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	// 检查状态码
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("http error: status=%d, body=%s", resp.StatusCode(), resp.String())
	}

	return resp.Body(), nil
}

// Get 发送GET请求
func (c *HTTPClient) Get(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	return c.Request(ctx, http.MethodGet, url, nil, headers)
}

// Post 发送POST请求
func (c *HTTPClient) Post(ctx context.Context, url string, body []byte, headers map[string]string) ([]byte, error) {
	return c.Request(ctx, http.MethodPost, url, bytes.NewReader(body), headers)
}

// Put 发送PUT请求
func (c *HTTPClient) Put(ctx context.Context, url string, body []byte, headers map[string]string) ([]byte, error) {
	return c.Request(ctx, http.MethodPut, url, bytes.NewReader(body), headers)
}

// Delete 发送DELETE请求
func (c *HTTPClient) Delete(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	return c.Request(ctx, http.MethodDelete, url, nil, headers)
}
