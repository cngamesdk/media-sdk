package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取中间号-通过线索id
func TestLeadsCallVirtualNumberGetByLeadsIdSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsCallVirtualNumberGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123456
	req.LeadsId = 218000154
	req.Caller = "13810433402"
	req.Callee = "13717983965"
	req.RequestId = "223255a1-2d02-44d0-8c1b-7217302de746"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsCallVirtualNumberGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取中间号-通过外部线索id
func TestLeadsCallVirtualNumberGetByOuterLeadsIdSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsCallVirtualNumberGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123456
	req.OuterLeadsId = "ext_001"
	req.Caller = "13810433402"
	req.Callee = "13717983965"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsCallVirtualNumberGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取中间号-不带request_id
func TestLeadsCallVirtualNumberGetWithoutRequestIdSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsCallVirtualNumberGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123456
	req.LeadsId = 218000154
	req.Caller = "01085170811"
	req.Callee = "13717983965"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsCallVirtualNumberGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少account_id
func TestLeadsCallVirtualNumberGetValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.LeadsCallVirtualNumberGetReq{}
	req.AccessToken = "123"
	req.LeadsId = 218000154
	req.Caller = "13810433402"
	req.Callee = "13717983965"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-leads_id和outer_leads_id都未填
func TestLeadsCallVirtualNumberGetValidateLeadsIdMissingSelf(t *testing.T) {
	req := &model.LeadsCallVirtualNumberGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123456
	req.Caller = "13810433402"
	req.Callee = "13717983965"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：leads_id和outer_leads_id二选一必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少caller
func TestLeadsCallVirtualNumberGetValidateCallerEmptySelf(t *testing.T) {
	req := &model.LeadsCallVirtualNumberGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123456
	req.LeadsId = 218000154
	req.Callee = "13717983965"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：caller为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少callee
func TestLeadsCallVirtualNumberGetValidateCalleeEmptySelf(t *testing.T) {
	req := &model.LeadsCallVirtualNumberGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123456
	req.LeadsId = 218000154
	req.Caller = "13810433402"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：callee为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
