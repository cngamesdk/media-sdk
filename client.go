package media_sdk

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/adapter"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/errors"
	_ "github.com/cngamesdk/media-sdk/media/toutiao"
	"github.com/cngamesdk/media-sdk/model"
	"sync"
)

// Client 媒体SDK客户端
type Client struct {
	config  *config.Config
	adapter adapter.MediaSDK
	mu      sync.RWMutex
}

// NewClientDefault 创建默认客户端
func NewClientDefault(mediaType config.MediaType) (*Client, error) {
	return NewClient(mediaType, config.DefaultConfig())
}

// NewClient 创建客户端
func NewClient(mediaType config.MediaType, config *config.Config) (*Client, error) {
	if config == nil {
		return nil, errors.ErrInvalidConfig
	}

	// 创建适配器
	ad, err := adapter.CreateAdapter(mediaType, config)
	if err != nil {
		return nil, fmt.Errorf("create adapter failed: %w", err)
	}

	return &Client{
		config:  config,
		adapter: ad,
	}, nil
}

// Auth 授权
func (c *Client) Auth(ctx context.Context, req *model.AuthReq) (interface{}, error) {
	if err := c.validateReq(req); err != nil {
		return nil, err
	}
	return c.adapter.Auth(ctx, req)
}

// AccessToken 获取AccessToken
func (c *Client) AccessToken(ctx context.Context, req *model.AccessTokenReq) (*model.AccessTokenResp, error) {
	if err := c.validateReq(req); err != nil {
		return nil, err
	}
	return c.adapter.AccessToken(ctx, req)
}

// RefreshToken 刷新Token
func (c *Client) RefreshToken(ctx context.Context, req *model.RefreshTokenReq) (*model.RefreshTokenResp, error) {
	if err := c.validateReq(req); err != nil {
		return nil, err
	}
	return c.adapter.RefreshToken(ctx, req)
}

// GetAccount 获取账户信息
func (c *Client) GetAccount(ctx context.Context, req *model.AccountReq) (*model.AccountResp, error) {
	if err := c.validateReq(req); err != nil {
		return nil, err
	}
	return c.adapter.GetAccount(ctx, req)
}

// CreateCampaign 创建广告计划
func (c *Client) CreateCampaign(ctx context.Context, req *model.CampaignReq) (*model.CampaignResp, error) {
	if err := c.validateReq(req); err != nil {
		return nil, err
	}
	return c.adapter.CreateCampaign(ctx, req)
}

// UpdateCampaign 更新广告计划
func (c *Client) UpdateCampaign(ctx context.Context, req *model.CampaignReq) (*model.CampaignResp, error) {
	if err := c.validateReq(req); err != nil {
		return nil, err
	}
	return c.adapter.UpdateCampaign(ctx, req)
}

// ListCampaigns 列出广告计划
func (c *Client) ListCampaigns(ctx context.Context, req *model.ListCampaignsReq) (*model.ListCampaignsResp, error) {
	if err := c.validateReq(req); err != nil {
		return nil, err
	}
	return c.adapter.ListCampaigns(ctx, req)
}

// CreateUnit 创建广告组
func (c *Client) CreateUnit(ctx context.Context, req *model.UnitReq) (*model.UnitResp, error) {
	if err := c.validateReq(req); err != nil {
		return nil, err
	}
	return c.adapter.CreateUnit(ctx, req)
}

// GetReport 获取报表
func (c *Client) GetReport(ctx context.Context, req *model.ReportReq) (*model.ReportResp, error) {
	if err := c.validateReq(req); err != nil {
		return nil, err
	}
	return c.adapter.GetReport(ctx, req)
}

// validateReq 验证请求参数
func (c *Client) validateReq(req interface{}) error {
	if req == nil {
		return errors.ErrInvalidRequest
	}
	return nil
}

// MultiClient 多客户端管理器
type MultiClient struct {
	clients map[config.MediaType]*Client
	mu      sync.RWMutex
}

// NewMultiClient 创建多客户端管理器
func NewMultiClient() *MultiClient {
	return &MultiClient{
		clients: make(map[config.MediaType]*Client),
	}
}

// RegisterClient 注册客户端
func (m *MultiClient) RegisterClient(mediaType config.MediaType, client *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.clients[mediaType] = client
}

// GetClient 获取客户端
func (m *MultiClient) GetClient(mediaType config.MediaType) (*Client, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	client, ok := m.clients[mediaType]
	if !ok {
		return nil, fmt.Errorf("client not found for media type: %s", mediaType)
	}
	return client, nil
}

// BatchExecute 批量执行
func (m *MultiClient) BatchExecute(ctx context.Context, fn func(client *Client) error) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var wg sync.WaitGroup
	errChan := make(chan error, len(m.clients))

	for _, client := range m.clients {
		wg.Add(1)
		go func(c *Client) {
			defer wg.Done()
			if err := fn(c); err != nil {
				errChan <- err
			}
		}(client)
	}

	wg.Wait()

	close(errChan)

	var errorsChan []error
	for err := range errChan {
		errorsChan = append(errorsChan, err)
	}

	if len(errorsChan) > 0 {
		return fmt.Errorf("batch execute failed with %d errors: %v", len(errorsChan), errorsChan)
	}
	return nil
}
