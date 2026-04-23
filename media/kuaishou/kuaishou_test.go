package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
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
	req := &kuaishouModel.AuthReq{}
	req.AppId = "123"
	req.RedirectUri = "https://www.xxx.com"
	req.OauthType = kuaishouModel.OauthTypeAdvertiser
	req.Scope = []string{"ad_manage"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AuthSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(cast.ToString(resp))
}

func TestAccessTokenSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AccessTokenReq{}
	req.AppId = 123
	req.Secret = "your_secret"
	req.AuthCode = "your_auth_code"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AccessTokenSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
