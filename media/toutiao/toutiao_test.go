package toutiao

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/toutiao/model"
	model2 "github.com/cngamesdk/media-sdk/model"
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

// AdvertiserInfoSelf 测试-获取客户信息
func TestAdvertiserInfoSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewToutiaoAdapter(config.DefaultConfig())
	req := &model.AccountReq{}
	req.AccessToken = "test"
	req.AdvertiserIds = []int64{123}
	resp, err := factory.AdvertiserInfoSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// GetAccount 测试-统一接口-获取客户信息
func TestGetAccount(t *testing.T) {
	ctx := context.Background()
	factory := NewToutiaoAdapter(config.DefaultConfig())
	req := &model2.AccountReq{}
	req.AccessToken = "test"
	req.AdvertiserID = 123
	resp, err := factory.GetAccount(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// EbpAdvertiserListSelf 测试-自己-获取升级版巨量引擎工作台下账户列表
func TestEbpAdvertiserListSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewToutiaoAdapter(config.DefaultConfig())
	req := &model.EbpAdvertiserListReq{}
	req.AccessToken = "test"
	req.EnterpriseOrganizationID = 123
	resp, err := factory.EbpAdvertiserListSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}
