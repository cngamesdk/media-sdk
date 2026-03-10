package adapter

import (
	"context"
	"github.com/cngamesdk/media-sdk/model"
)

type Adapter interface {
	Name() string
	Code() string
	Init(req AdapterConfig) Adapter
	AuthRedirect(ctx context.Context, req model.AuthRedirectReq)
}
