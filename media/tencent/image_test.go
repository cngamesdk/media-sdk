package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ========== 获取图片信息测试用例 ==========

// TestImageGetByAccountIDSelf 测试按 account_id 查询图片列表
func TestImageGetByAccountIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetByOrganizationIDSelf 测试按 organization_id 查询图片列表
func TestImageGetByOrganizationIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.OrganizationID = 222222
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetFilterByImageWidthSelf 测试按图片宽度过滤
func TestImageGetFilterByImageWidthSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldImageWidth, Operator: model.OperatorEquals, Values: []string{"640"}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetFilterByImageIDSelf 测试按 image_id IN 过滤
func TestImageGetFilterByImageIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldImageID, Operator: model.OperatorIn, Values: []string{"img001", "img002", "img003"}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetFilterByImageSignatureSelf 测试按图片签名过滤
func TestImageGetFilterByImageSignatureSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldImageSignature, Operator: model.OperatorEquals, Values: []string{"f4c8a3bc4deb305fb74cb08ed395b98c"}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetFilterByCreatedTimeSelf 测试按创建时间过滤
func TestImageGetFilterByCreatedTimeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldCreatedTime, Operator: model.OperatorGreaterEquals, Values: []string{"2024-01-01"}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetFilterBySourceTypeSelf 测试按来源类型过滤
func TestImageGetFilterBySourceTypeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldSourceType, Operator: model.OperatorEquals, Values: []string{model.ImageSourceTypeLocal}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetFilterByStatusSelf 测试按状态过滤
func TestImageGetFilterByStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldStatus, Operator: model.OperatorEquals, Values: []string{model.ImageStatusNormal}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetFilterByAigcFlagSelf 测试按 AIGC 标记过滤
func TestImageGetFilterByAigcFlagSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldAigcFlag, Operator: model.OperatorEquals, Values: []string{model.ImageAigcFlagUseMuseAI}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetWithMultipleFiltersSelf 测试多个过滤条件（最多4个）
func TestImageGetWithMultipleFiltersSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldImageWidth, Operator: model.OperatorEquals, Values: []string{"1920"}},
		{Field: model.ImageFilterFieldImageHeight, Operator: model.OperatorEquals, Values: []string{"1080"}},
		{Field: model.ImageFilterFieldSourceType, Operator: model.OperatorEquals, Values: []string{model.ImageSourceTypeAPI}},
		{Field: model.ImageFilterFieldStatus, Operator: model.OperatorEquals, Values: []string{model.ImageStatusNormal}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetWithPaginationSelf 测试自定义分页
func TestImageGetWithPaginationSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 2
	req.PageSize = 50
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetWithLabelIDSelf 测试携带 label_id 参数
func TestImageGetWithLabelIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.LabelID = 5001
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestImageGetWithBusinessScenarioSelf 测试携带 business_scenario 参数
func TestImageGetWithBusinessScenarioSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.BusinessScenario = 2 // 投放素材包
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ImageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 验证测试用例 ==========

// TestImageGetValidateMissingAccountAndOrgSelf 测试 account_id 和 organization_id 均未填写
func TestImageGetValidateMissingAccountAndOrgSelf(t *testing.T) {
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id 和 organization_id 需必填其一")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageGetValidateFilteringExceedMaxSelf 测试 filtering 超过4条
func TestImageGetValidateFilteringExceedMaxSelf(t *testing.T) {
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldImageWidth, Operator: model.OperatorEquals, Values: []string{"640"}},
		{Field: model.ImageFilterFieldImageHeight, Operator: model.OperatorEquals, Values: []string{"360"}},
		{Field: model.ImageFilterFieldSourceType, Operator: model.OperatorEquals, Values: []string{model.ImageSourceTypeLocal}},
		{Field: model.ImageFilterFieldStatus, Operator: model.OperatorEquals, Values: []string{model.ImageStatusNormal}},
		{Field: model.ImageFilterFieldAigcFlag, Operator: model.OperatorEquals, Values: []string{model.ImageAigcFlagNotAI}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering超过4条")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageGetValidateFilteringMissingFieldSelf 测试 filtering 缺少 field
func TestImageGetValidateFilteringMissingFieldSelf(t *testing.T) {
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Operator: model.OperatorEquals, Values: []string{"640"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：field为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageGetValidateFilteringMissingOperatorSelf 测试 filtering 缺少 operator
func TestImageGetValidateFilteringMissingOperatorSelf(t *testing.T) {
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldImageWidth, Values: []string{"640"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：operator为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageGetValidateFilteringEmptyValuesSelf 测试 filtering values 为空
func TestImageGetValidateFilteringEmptyValuesSelf(t *testing.T) {
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.ImageFilteringItem{
		{Field: model.ImageFilterFieldStatus, Operator: model.OperatorEquals, Values: []string{}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：values为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageGetValidatePageTooLargeSelf 测试 page 超过最大值 99999
func TestImageGetValidatePageTooLargeSelf(t *testing.T) {
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 1
	req.PageSize = 10
	req.Format()
	req.Page = 100000
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page超过99999")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestImageGetValidatePageSizeTooLargeSelf 测试 page_size 超过最大值 100
func TestImageGetValidatePageSizeTooLargeSelf(t *testing.T) {
	req := &model.ImageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 1
	req.PageSize = 101
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size超过100")
	}
	fmt.Printf("验证错误: %v\n", err)
}
