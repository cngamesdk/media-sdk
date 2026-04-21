package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取人群授权信息-基本查询（不带过滤条件）
func TestAudienceGrantRelationsGet(t *testing.T) {
	ctx := context.Background()
	req := &model.AudienceGrantRelationsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Page = 1
	req.PageSize = 10
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AudienceGrantRelationsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取人群授权信息-按 audience_id 过滤
func TestAudienceGrantRelationsGetByAudienceID(t *testing.T) {
	ctx := context.Background()
	req := &model.AudienceGrantRelationsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Filtering = []*model.AudienceGrantRelationsFiltering{
		{
			Field:    model.AudienceGrantRelationsFilterFieldAudienceID,
			Operator: "IN",
			Values:   []string{"123", "456"},
		},
	}
	req.Page = 1
	req.PageSize = 10
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AudienceGrantRelationsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取人群授权信息-按 audience_id 过滤单个
func TestAudienceGrantRelationsGetBySingleAudienceID(t *testing.T) {
	ctx := context.Background()
	req := &model.AudienceGrantRelationsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Filtering = []*model.AudienceGrantRelationsFiltering{
		{
			Field:    model.AudienceGrantRelationsFilterFieldAudienceID,
			Operator: "IN",
			Values:   []string{"123"},
		},
	}
	req.Page = 1
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AudienceGrantRelationsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取人群授权信息-翻页查询
func TestAudienceGrantRelationsGetPage2(t *testing.T) {
	ctx := context.Background()
	req := &model.AudienceGrantRelationsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Page = 2
	req.PageSize = 50
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AudienceGrantRelationsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取人群授权信息-Format 默认值
func TestAudienceGrantRelationsGetFormatDefaults(t *testing.T) {
	req := &model.AudienceGrantRelationsGetReq{}
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

// 验证测试-缺少 account_id
func TestAudienceGrantRelationsGetValidateAccountIDEmpty(t *testing.T) {
	req := &model.AudienceGrantRelationsGetReq{}
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

// 验证测试-filtering.field 为空
func TestAudienceGrantRelationsGetValidateFilteringFieldEmpty(t *testing.T) {
	req := &model.AudienceGrantRelationsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Filtering = []*model.AudienceGrantRelationsFiltering{
		{Operator: "IN", Values: []string{"123"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.field为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-filtering.operator 为空
func TestAudienceGrantRelationsGetValidateFilteringOperatorEmpty(t *testing.T) {
	req := &model.AudienceGrantRelationsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Filtering = []*model.AudienceGrantRelationsFiltering{
		{Field: model.AudienceGrantRelationsFilterFieldAudienceID, Values: []string{"123"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.operator为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-filtering.values 为空
func TestAudienceGrantRelationsGetValidateFilteringValuesEmpty(t *testing.T) {
	req := &model.AudienceGrantRelationsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Filtering = []*model.AudienceGrantRelationsFiltering{
		{Field: model.AudienceGrantRelationsFilterFieldAudienceID, Operator: "IN"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.values为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-filtering.values 超过最大长度（audience_id 最多 100）
func TestAudienceGrantRelationsGetValidateFilteringValuesTooLong(t *testing.T) {
	req := &model.AudienceGrantRelationsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	values := make([]string, 101)
	for i := range values {
		values[i] = fmt.Sprintf("%d", i+1)
	}
	req.Filtering = []*model.AudienceGrantRelationsFiltering{
		{Field: model.AudienceGrantRelationsFilterFieldAudienceID, Operator: "IN", Values: values},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.values最大长度为100")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-page_size 超出范围
func TestAudienceGrantRelationsGetValidatePageSizeInvalid(t *testing.T) {
	req := &model.AudienceGrantRelationsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Page = 1
	req.PageSize = 101
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size必须在1-100之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}
