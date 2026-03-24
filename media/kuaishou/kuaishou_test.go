package kuaishou

import (
	"context"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/model"
	"github.com/spf13/cast"
	"testing"
)

func TestAuth(t *testing.T) {
	ctx := context.Background()
	req := &model.AuthReq{}
	req.AppId = 123
	req.RedirectUri = "https://www.xxx.com"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.Auth(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(cast.ToString(resp))
}
