package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取无效赔付明细
func TestLeadsInvalidPayGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsInvalidPayGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123456
	req.Month = "2024-01"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsInvalidPayGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取无效赔付明细-其他月份
func TestLeadsInvalidPayGetOtherMonthSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsInvalidPayGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123456
	req.Month = "2023-12"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsInvalidPayGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少account_id
func TestLeadsInvalidPayGetValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.LeadsInvalidPayGetReq{}
	req.AccessToken = "123"
	req.Month = "2024-01"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少month
func TestLeadsInvalidPayGetValidateMonthEmptySelf(t *testing.T) {
	req := &model.LeadsInvalidPayGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123456
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：month为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-month格式不正确
func TestLeadsInvalidPayGetValidateMonthInvalidSelf(t *testing.T) {
	req := &model.LeadsInvalidPayGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123456
	req.Month = "2024-1"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：month长度必须为7字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}
