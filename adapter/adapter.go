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
func CreateAdapter(mediaType config.MediaType, config *config.Config) (MediaSDK, error) {
	mu.RLock()
	factory, ok := adapters[mediaType]
	mu.RUnlock()

	if !ok {
		return nil, fmt.Errorf("unsupported media type: %s", mediaType)
	}

	return factory.Create(config)
}
