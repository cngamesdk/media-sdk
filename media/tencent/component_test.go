package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// TestComponentsGetSelf 测试获取全部创意组件（不加过滤条件）
func TestComponentsGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetByComponentIDSelf 测试按组件id过滤获取创意组件
func TestComponentsGetByComponentIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.ComponentQueryFilter{
		{
			Field:    model.ComponentFieldComponentID,
			Operator: model.OperatorIn,
			Values:   []string{"111111", "222222"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetBySubTypeSelf 测试按组件子类型过滤获取创意组件
func TestComponentsGetBySubTypeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.ComponentQueryFilter{
		{
			Field:    model.ComponentFieldComponentSubType,
			Operator: model.OperatorEquals,
			Values:   []string{"COMPONENT_SUB_TYPE_IMAGE"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetByCreatedTimeSelf 测试按创建时间过滤获取创意组件
func TestComponentsGetByCreatedTimeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.ComponentQueryFilter{
		{
			Field:    model.ComponentFieldCreatedTime,
			Operator: model.OperatorGreaterEquals,
			Values:   []string{"1704067200"}, // 2024-01-01 00:00:00
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetByVideoIDSelf 测试按视频id过滤获取创意组件
func TestComponentsGetByVideoIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.ComponentQueryFilter{
		{
			Field:    model.ComponentFieldVideoID,
			Operator: model.OperatorIn,
			Values:   []string{"video_id_001", "video_id_002"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetByImageIDSelf 测试按图片id过滤获取创意组件
func TestComponentsGetByImageIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.ComponentQueryFilter{
		{
			Field:    model.ComponentFieldImageID,
			Operator: model.OperatorEquals,
			Values:   []string{"image_id_001"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetByPotentialStatusSelf 测试按潜力状态过滤获取创意组件
func TestComponentsGetByPotentialStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.ComponentQueryFilter{
		{
			Field:    model.ComponentFieldPotentialStatus,
			Operator: model.OperatorEquals,
			Values:   []string{model.ComponentPotentialStatusHigh},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetWithPaginationSelf 测试分页获取创意组件
func TestComponentsGetWithPaginationSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Page = 2
	req.PageSize = 50
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetWithFieldsSelf 测试指定返回字段获取创意组件
func TestComponentsGetWithFieldsSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Fields = []string{"component_id", "component_sub_type", "component_custom_name", "created_time"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetDeletedSelf 测试获取已删除的创意组件
func TestComponentsGetDeletedSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.IsDeleted = true
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetWithOrganizationIDSelf 测试通过业务单元获取创意组件
func TestComponentsGetWithOrganizationIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.OrganizationID = 456
	req.ComponentIDFilteringMode = model.ComponentIDFilteringModeSharingByCustomerBusinessUnit
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestComponentsGetValidateAccountID 测试缺少account_id时的校验
func TestComponentsGetValidateAccountID(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentsGetReq{}
	req.AccessToken = "123"
	adapter := NewTencentAdapter(config.DefaultConfig())
	_, err := adapter.ComponentsGetSelf(ctx, req)
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}
