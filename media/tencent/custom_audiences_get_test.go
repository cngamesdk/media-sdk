package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取客户人群-基本查询
func TestCustomAudiencesGet(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudiencesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Page = 1
	req.PageSize = 10
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudiencesGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取客户人群-按人群ID查询
func TestCustomAudiencesGetByAudienceID(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudiencesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.Page = 1
	req.PageSize = 10
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudiencesGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取客户人群-翻页查询
func TestCustomAudiencesGetPage2(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudiencesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Page = 2
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudiencesGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestCustomAudiencesGetValidateAccountIDEmpty(t *testing.T) {
	req := &model.CustomAudiencesGetReq{}
	req.AccessToken = "123"
	req.Page = 1
	req.PageSize = 10
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-page_size 超出范围
func TestCustomAudiencesGetValidatePageSizeInvalid(t *testing.T) {
	req := &model.CustomAudiencesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Page = 1
	req.PageSize = 200
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size必须在1-100之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-page 默认值
func TestCustomAudiencesGetFormatDefaults(t *testing.T) {
	req := &model.CustomAudiencesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Format()
	if req.Page != 1 {
		t.Fatalf("期望 page 默认值为 1，实际为 %d", req.Page)
	}
	if req.PageSize != 10 {
		t.Fatalf("期望 page_size 默认值为 10，实际为 %d", req.PageSize)
	}
	fmt.Printf("默认值验证通过: page=%d, page_size=%d\n", req.Page, req.PageSize)
}
