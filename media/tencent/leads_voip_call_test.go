package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 网络电话呼叫-通过 leads_id
func TestLeadsVoipCallAddByLeadsIdSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsVoipCallAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.LeadsId = 218000154
	req.UserID = 1001
	req.CalleeNumber = "13800138000"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsVoipCallAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 网络电话呼叫-通过 outer_leads_id
func TestLeadsVoipCallAddByOuterLeadsIdSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsVoipCallAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.OuterLeadsId = "ext_001"
	req.UserID = 1001
	req.CalleeNumber = "13800138000"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsVoipCallAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 网络电话呼叫-携带 request_id 和 version
func TestLeadsVoipCallAddWithRequestIdSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsVoipCallAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.LeadsId = 218000154
	req.UserID = 1001
	req.CalleeNumber = "13800138000"
	req.RequestId = "223255a1-2d02-44d0-8c1b-7217302de746"
	req.Version = "1.0"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsVoipCallAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestLeadsVoipCallAddValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.LeadsVoipCallAddReq{}
	req.AccessToken = "123"
	req.LeadsId = 218000154
	req.UserID = 1001
	req.CalleeNumber = "13800138000"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-leads_id 和 outer_leads_id 都未填
func TestLeadsVoipCallAddValidateLeadsIdMissingSelf(t *testing.T) {
	req := &model.LeadsVoipCallAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.UserID = 1001
	req.CalleeNumber = "13800138000"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：leads_id和outer_leads_id二选一必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 user_id
func TestLeadsVoipCallAddValidateUserIdEmptySelf(t *testing.T) {
	req := &model.LeadsVoipCallAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.LeadsId = 218000154
	req.CalleeNumber = "13800138000"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：user_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 callee_number
func TestLeadsVoipCallAddValidateCalleeNumberEmptySelf(t *testing.T) {
	req := &model.LeadsVoipCallAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.LeadsId = 218000154
	req.UserID = 1001
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：callee_number为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
