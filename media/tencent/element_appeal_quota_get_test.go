package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取元素申诉复审配额-基本查询
func TestElementAppealQuotaGet(t *testing.T) {
	ctx := context.Background()
	req := &model.ElementAppealQuotaGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ElementAppealQuotaGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestElementAppealQuotaGetValidateAccountIDEmpty(t *testing.T) {
	req := &model.ElementAppealQuotaGetReq{}
	req.AccessToken = "123"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 access_token
func TestElementAppealQuotaGetValidateAccessTokenEmpty(t *testing.T) {
	req := &model.ElementAppealQuotaGetReq{}
	req.AccountID = 2045867
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：access_token为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
