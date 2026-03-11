package adapter

import (
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"sync"
)

var (
	adapters = make(map[config.MediaType]Factory)
	mu       sync.RWMutex
)

// Register 注册适配器工厂
func Register(mediaType config.MediaType, factory Factory) {
	mu.Lock()
	defer mu.Unlock()
	adapters[mediaType] = factory
}

// CreateAdapter 创建适配器
func CreateAdapter(config *config.Config) (MediaSDK, error) {
	mu.RLock()
	factory, ok := adapters[config.MediaType]
	mu.RUnlock()

	if !ok {
		return nil, fmt.Errorf("unsupported media type: %s", config.MediaType)
	}

	return factory.Create(config)
}
