package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

func TestOAuth2AuthorizeSelf(t *testing.T) {
	ctx := context.Background()
	adapter := NewTencentAdapter(config.DefaultConfig())
	req := &model.OAuth2AuthorizeReq{}
	req.ClientID = 123
	req.RedirectURI = "https://www.xxx.com"
	resp, err := adapter.OAuth2AuthorizeSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(resp)
}

func TestOAuth2TokenSelf(t *testing.T) {
	ctx := context.Background()
	adapter := NewTencentAdapter(config.DefaultConfig())
	req := &model.OAuth2TokenReq{}
	req.ClientID = 123
	req.ClientSecret = "123"
	req.GrantType = model.GrantTypeAuthorizationCode
	req.AuthorizationCode = "123"
	req.RedirectURI = "https://www.xxx.com"
	resp, err := adapter.OAuth2TokenSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(resp)
}
