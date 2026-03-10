package media_sdk

import (
	"errors"
	"github.com/cngamesdk/media-sdk/adapter"
	"github.com/cngamesdk/media-sdk/media/ocean_engine"
	"go.uber.org/zap"
)

// GetAdapter 获取适配器
func GetAdapter(code string, logger *zap.Logger) (resp adapter.Adapter, err error) {
	switch code {
	case ocean_engine.Code:
		resp = ocean_engine.NewOceanEngine(logger)
		return
	}
	err = errors.New("未找到媒体适配器." + code)
	return
}
