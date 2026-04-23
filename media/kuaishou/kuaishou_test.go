package kuaishou

import (
	"context"
	"github.com/cngamesdk/media-sdk/config"
	model3 "github.com/cngamesdk/media-sdk/media/kuaishou/model"
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

func TestAuthSelf(t *testing.T) {
	ctx := context.Background()
	req := &model3.AuthReq{}
	req.AppId = "123"
	req.RedirectUri = "https://www.xxx.com"
	req.OauthType = model3.OauthTypeAdvertiser
	req.Scope = []string{"ad_manage"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AuthSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(cast.ToString(resp))
}
