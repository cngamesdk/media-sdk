package ocean_engine

import (
	"context"
	"github.com/cngamesdk/media-sdk/adapter"
	"github.com/cngamesdk/media-sdk/media"
	"github.com/cngamesdk/media-sdk/model"
	"go.uber.org/zap"
)

const Code = "oceanengine"

func NewOceanEngine(logger *zap.Logger) adapter.Adapter {
	return &OceanEngine{media.Media{Logger: logger}}
}

type OceanEngine struct {
	media.Media
}

func (receiver *OceanEngine) Name() string {
	return "巨量引擎"
}

func (receiver *OceanEngine) Code() string {
	return Code
}

func (receiver *OceanEngine) Init(config adapter.AdapterConfig) adapter.Adapter {
	receiver.Config = config
	return receiver
}

func (receiver *OceanEngine) AuthRedirect(ctx context.Context, req model.AuthRedirectReq) {
	return
}
