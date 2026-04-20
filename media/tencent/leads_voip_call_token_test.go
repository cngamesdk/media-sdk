package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取网络电话 token
func TestLeadsVoipCallTokenGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsVoipCallTokenGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.UserID = 1001
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsVoipCallTokenGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取网络电话 token-携带 request_id
func TestLeadsVoipCallTokenGetWithRequestIdSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsVoipCallTokenGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.UserID = 1001
	req.RequestId = "223255a1-2d02-44d0-8c1b-7217302de746"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsVoipCallTokenGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestLeadsVoipCallTokenGetValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.LeadsVoipCallTokenGetReq{}
	req.AccessToken = "123"
	req.UserID = 1001
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 user_id
func TestLeadsVoipCallTokenGetValidateUserIdEmptySelf(t *testing.T) {
	req := &model.LeadsVoipCallTokenGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：user_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
