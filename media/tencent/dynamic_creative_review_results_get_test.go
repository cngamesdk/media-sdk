package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 查询动态创意审核结果-基本查询
func TestDynamicCreativeReviewResultsGet(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativeReviewResultsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeIDList = []int64{123456789}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativeReviewResultsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 查询动态创意审核结果-多个创意ID
func TestDynamicCreativeReviewResultsGetMultipleIDs(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativeReviewResultsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeIDList = []int64{123456789, 987654321, 111222333}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativeReviewResultsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestDynamicCreativeReviewResultsGetValidateAccountIDEmpty(t *testing.T) {
	req := &model.DynamicCreativeReviewResultsGetReq{}
	req.AccessToken = "123"
	req.DynamicCreativeIDList = []int64{123456789}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 dynamic_creative_id_list
func TestDynamicCreativeReviewResultsGetValidateIDListEmpty(t *testing.T) {
	req := &model.DynamicCreativeReviewResultsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：dynamic_creative_id_list为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-dynamic_creative_id_list 超过最大长度
func TestDynamicCreativeReviewResultsGetValidateIDListTooLong(t *testing.T) {
	req := &model.DynamicCreativeReviewResultsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	ids := make([]int64, 101)
	for i := range ids {
		ids[i] = int64(i + 1)
	}
	req.DynamicCreativeIDList = ids
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：dynamic_creative_id_list最大长度100")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 access_token
func TestDynamicCreativeReviewResultsGetValidateAccessTokenEmpty(t *testing.T) {
	req := &model.DynamicCreativeReviewResultsGetReq{}
	req.AccountID = 2045867
	req.DynamicCreativeIDList = []int64{123456789}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：access_token为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
