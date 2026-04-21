package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 标签广场标签获取-不传过滤条件
func TestLabelsGetNoFilter(t *testing.T) {
	ctx := context.Background()
	req := &model.LabelsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LabelsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 标签广场标签获取-按标签分组 MAP 过滤
func TestLabelsGetFilterLabelGroupMAP(t *testing.T) {
	ctx := context.Background()
	req := &model.LabelsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Filtering = []*model.LabelsFiltering{
		{
			Field:    model.LabelsFilterFieldLabelGroup,
			Operator: "IN",
			Values:   []string{model.LabelsGroupMAP},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LabelsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 标签广场标签获取-按标签分组 POP 过滤
func TestLabelsGetFilterLabelGroupPOP(t *testing.T) {
	ctx := context.Background()
	req := &model.LabelsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Filtering = []*model.LabelsFiltering{
		{
			Field:    model.LabelsFilterFieldLabelGroup,
			Operator: "IN",
			Values:   []string{model.LabelsGroupPOP},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LabelsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 标签广场标签获取-按父级标签 id NOT_IN 过滤
func TestLabelsGetFilterParentIDNotIn(t *testing.T) {
	ctx := context.Background()
	req := &model.LabelsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Filtering = []*model.LabelsFiltering{
		{
			Field:    model.LabelsFilterFieldParentID,
			Operator: "NOT_IN",
			Values:   []string{"100001", "100002"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LabelsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 标签广场标签获取-按标签名称 CONTAINS 过滤
func TestLabelsGetFilterDisplayLabelNameContains(t *testing.T) {
	ctx := context.Background()
	req := &model.LabelsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Filtering = []*model.LabelsFiltering{
		{
			Field:    model.LabelsFilterFieldDisplayLabelName,
			Operator: "CONTAINS",
			Values:   []string{"汽车"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LabelsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 标签广场标签获取-多过滤条件组合
func TestLabelsGetFilterCombined(t *testing.T) {
	ctx := context.Background()
	req := &model.LabelsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Filtering = []*model.LabelsFiltering{
		{
			Field:    model.LabelsFilterFieldLabelGroup,
			Operator: "IN",
			Values:   []string{model.LabelsGroupMAP},
		},
		{
			Field:    model.LabelsFilterFieldDisplayLabelName,
			Operator: "CONTAINS",
			Values:   []string{"游戏"},
		},
	}
	req.Page = 1
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LabelsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 标签广场标签获取-验证 Format 默认值
func TestLabelsGetFormatDefaults(t *testing.T) {
	req := &model.LabelsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Format()
	if req.Page != 1 {
		t.Fatalf("期望 page 默认值为 1，实际为 %d", req.Page)
	}
	if req.PageSize != 10 {
		t.Fatalf("期望 page_size 默认值为 10，实际为 %d", req.PageSize)
	}
	fmt.Printf("page=%d, page_size=%d\n", req.Page, req.PageSize)
}

// 验证测试-缺少 account_id
func TestLabelsGetValidateAccountIDEmpty(t *testing.T) {
	req := &model.LabelsGetReq{}
	req.AccessToken = "123"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-filtering.field 为空
func TestLabelsGetValidateFilterFieldEmpty(t *testing.T) {
	req := &model.LabelsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Filtering = []*model.LabelsFiltering{
		{Field: "", Operator: "IN", Values: []string{"MAP"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.field为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-filtering.field 值无效
func TestLabelsGetValidateFilterFieldInvalid(t *testing.T) {
	req := &model.LabelsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Filtering = []*model.LabelsFiltering{
		{Field: "invalid_field", Operator: "IN", Values: []string{"MAP"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：field值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-filtering.operator 与 field 不匹配
func TestLabelsGetValidateFilterOperatorMismatch(t *testing.T) {
	req := &model.LabelsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Filtering = []*model.LabelsFiltering{
		{
			Field:    model.LabelsFilterFieldLabelGroup,
			Operator: "CONTAINS", // label_group 需要 IN
			Values:   []string{model.LabelsGroupMAP},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：operator与field不匹配")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-filtering.values 为空
func TestLabelsGetValidateFilterValuesEmpty(t *testing.T) {
	req := &model.LabelsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Filtering = []*model.LabelsFiltering{
		{
			Field:    model.LabelsFilterFieldLabelGroup,
			Operator: "IN",
			Values:   []string{},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：values为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-label_group values 包含无效值
func TestLabelsGetValidateFilterLabelGroupInvalidValue(t *testing.T) {
	req := &model.LabelsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Filtering = []*model.LabelsFiltering{
		{
			Field:    model.LabelsFilterFieldLabelGroup,
			Operator: "IN",
			Values:   []string{"INVALID_GROUP"},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：label_group values 可选值为 MAP、POP")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-page_size 超出范围
func TestLabelsGetValidatePageSizeOutOfRange(t *testing.T) {
	req := &model.LabelsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Page = 1
	req.PageSize = 200
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size必须在1-100之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}
