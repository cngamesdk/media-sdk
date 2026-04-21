package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 删除客户人群-基本删除
func TestCustomAudiencesDelete(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudiencesDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudiencesDelete(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestCustomAudiencesDeleteValidateAccountIDEmpty(t *testing.T) {
	req := &model.CustomAudiencesDeleteReq{}
	req.AccessToken = "123"
	req.AudienceID = 123456789
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 audience_id
func TestCustomAudiencesDeleteValidateAudienceIDEmpty(t *testing.T) {
	req := &model.CustomAudiencesDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：audience_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 access_token
func TestCustomAudiencesDeleteValidateAccessTokenEmpty(t *testing.T) {
	req := &model.CustomAudiencesDeleteReq{}
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：access_token为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
