package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 查询组件审核结果-基本查询
func TestComponentReviewResultsGet(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentReviewResultsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.ComponentIDList = []int64{123456789}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentReviewResultsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 查询组件审核结果-多个组件ID
func TestComponentReviewResultsGetMultipleIDs(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentReviewResultsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.ComponentIDList = []int64{123456789, 987654321, 111222333}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentReviewResultsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestComponentReviewResultsGetValidateAccountIDEmpty(t *testing.T) {
	req := &model.ComponentReviewResultsGetReq{}
	req.AccessToken = "123"
	req.ComponentIDList = []int64{123456789}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 component_id_list
func TestComponentReviewResultsGetValidateIDListEmpty(t *testing.T) {
	req := &model.ComponentReviewResultsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：component_id_list为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-component_id_list 超过最大长度
func TestComponentReviewResultsGetValidateIDListTooLong(t *testing.T) {
	req := &model.ComponentReviewResultsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	ids := make([]int64, 101)
	for i := range ids {
		ids[i] = int64(i + 1)
	}
	req.ComponentIDList = ids
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：component_id_list最大长度100")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 access_token
func TestComponentReviewResultsGetValidateAccessTokenEmpty(t *testing.T) {
	req := &model.ComponentReviewResultsGetReq{}
	req.AccountID = 2045867
	req.ComponentIDList = []int64{123456789}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：access_token为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
