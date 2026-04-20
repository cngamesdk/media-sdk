package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取全部通话结果-第一页
func TestLeadsCallRecordsGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsCallRecordsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.PageSize = 10
	req.Page = 1
	req.StartDate = "2024-01-01"
	req.EndDate = "2024-01-30"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsCallRecordsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取全部通话结果-使用search_after翻页
func TestLeadsCallRecordsGetWithSearchAfterSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsCallRecordsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.PageSize = 10
	req.Page = 2
	req.StartDate = "2023-01-01"
	req.EndDate = "2023-01-30"
	req.SearchAfter = "2023-01-15 13:15:29"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsCallRecordsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少account_id
func TestLeadsCallRecordsGetValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.LeadsCallRecordsGetReq{}
	req.AccessToken = "123"
	req.PageSize = 10
	req.Page = 1
	req.StartDate = "2024-01-01"
	req.EndDate = "2024-01-30"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少page_size
func TestLeadsCallRecordsGetValidatePageSizeEmptySelf(t *testing.T) {
	req := &model.LeadsCallRecordsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Page = 1
	req.StartDate = "2024-01-01"
	req.EndDate = "2024-01-30"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size为必填且必须大于0")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少start_date
func TestLeadsCallRecordsGetValidateStartDateEmptySelf(t *testing.T) {
	req := &model.LeadsCallRecordsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.PageSize = 10
	req.Page = 1
	req.EndDate = "2024-01-30"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：start_date为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少end_date
func TestLeadsCallRecordsGetValidateEndDateEmptySelf(t *testing.T) {
	req := &model.LeadsCallRecordsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.PageSize = 10
	req.Page = 1
	req.StartDate = "2024-01-01"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：end_date为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
