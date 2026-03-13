package toutiao

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/toutiao/model"
	"testing"
)

// TestAuthAdvertiserGet 测试-获取已授权账户
func TestAuthAdvertiserGet(t *testing.T) {
	ctx := context.Background()
	factory := NewToutiaoAdapter(config.DefaultConfig())
	req := &model.AuthAdvertiserGetReq{}
	req.AccessToken = "test"
	resp, err := factory.AuthAdvertiserGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestAuthAdvertiserGet 测试-获取授权User信息
func TestAuthUserInfo(t *testing.T) {
	ctx := context.Background()
	factory := NewToutiaoAdapter(config.DefaultConfig())
	req := &model.AuthUserInfoReq{}
	req.AccessToken = "test"
	resp, err := factory.AuthUserInfoSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}
