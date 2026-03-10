package media

import (
	"github.com/cngamesdk/media-sdk/adapter"
	"go.uber.org/zap"
)

type Media struct {
	Config adapter.AdapterConfig
	Logger *zap.Logger
}
