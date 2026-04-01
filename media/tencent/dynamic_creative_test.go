package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取创意（基础）
func TestDynamicCreativesGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 按创意ID过滤
func TestDynamicCreativesGetByIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.DynamicCreativeQueryFilter{
		{
			Field:    model.CreativeFieldDynamicCreativeID,
			Operator: model.OperatorIn,
			Values:   []string{"40958977", "40958978"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 按广告ID过滤
func TestDynamicCreativesGetByAdgroupIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.DynamicCreativeQueryFilter{
		{
			Field:    model.CreativeFieldAdgroupID,
			Operator: model.OperatorEquals,
			Values:   []string{"456"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 按配置状态过滤
func TestDynamicCreativesGetByStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.DynamicCreativeQueryFilter{
		{
			Field:    model.CreativeFieldConfiguredStatus,
			Operator: model.OperatorEquals,
			Values:   []string{model.CreativeConfiguredStatusNormal},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 指定返回字段 + 游标分页
func TestDynamicCreativesGetWithCursorSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.DynamicCreativesGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Fields = []string{"dynamic_creative_id", "dynamic_creative_name", "adgroup_id", "configured_status", "creative_components"}
	req.PaginationMode = model.CreativePaginationModeCursor
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.DynamicCreativesGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}
