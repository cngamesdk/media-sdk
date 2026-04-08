package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ========== 获取视频文件测试用例 ==========

// TestVideoGetByAccountIDSelf 测试按 account_id 查询视频列表
func TestVideoGetByAccountIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.VideoGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.VideoGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestVideoGetByOrganizationIDSelf 测试按 organization_id 查询视频列表
func TestVideoGetByOrganizationIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.VideoGetReq{}
	req.AccessToken = "123"
	req.OrganizationID = 222222
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.VideoGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestVideoGetFilterByMediaWidthSelf 测试按视频宽度过滤
func TestVideoGetFilterByMediaWidthSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.VideoGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.VideoFilteringItem{
		{Field: model.VideoFilterFieldMediaWidth, Operator: model.OperatorEquals, Values: []string{"640"}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.VideoGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestVideoGetFilterByMediaIDSelf 测试按 media_id IN 过滤
func TestVideoGetFilterByMediaIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.VideoGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.VideoFilteringItem{
		{Field: model.VideoFilterFieldMediaID, Operator: model.OperatorIn, Values: []string{"10001", "10002", "10003"}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.VideoGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestVideoGetFilterByCreatedTimeSelf 测试按创建时间过滤
func TestVideoGetFilterByCreatedTimeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.VideoGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.VideoFilteringItem{
		{Field: model.VideoFilterFieldCreatedTime, Operator: model.OperatorGreaterEquals, Values: []string{"2024-01-01"}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.VideoGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestVideoGetFilterBySourceTypeSelf 测试按来源类型过滤
func TestVideoGetFilterBySourceTypeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.VideoGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.VideoFilteringItem{
		{Field: model.VideoFilterFieldSourceType, Operator: model.OperatorEquals, Values: []string{model.VideoSourceTypeLocal}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.VideoGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestVideoGetFilterByStatusSelf 测试按视频状态过滤
func TestVideoGetFilterByStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.VideoGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.VideoFilteringItem{
		{Field: model.VideoFilterFieldStatus, Operator: model.OperatorEquals, Values: []string{model.VideoStatusNormal}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.VideoGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestVideoGetWithMultipleFiltersSelf 测试多个过滤条件（最多4个）
func TestVideoGetWithMultipleFiltersSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.VideoGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.VideoFilteringItem{
		{Field: model.VideoFilterFieldMediaWidth, Operator: model.OperatorEquals, Values: []string{"1920"}},
		{Field: model.VideoFilterFieldMediaHeight, Operator: model.OperatorEquals, Values: []string{"1080"}},
		{Field: model.VideoFilterFieldSourceType, Operator: model.OperatorEquals, Values: []string{model.VideoSourceTypeAPI}},
		{Field: model.VideoFilterFieldStatus, Operator: model.OperatorEquals, Values: []string{model.VideoStatusNormal}},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.VideoGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestVideoGetWithPaginationSelf 测试自定义分页
func TestVideoGetWithPaginationSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.VideoGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 2
	req.PageSize = 50
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.VideoGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestVideoGetWithLabelIDSelf 测试携带 label_id 参数
func TestVideoGetWithLabelIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.VideoGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.LabelID = 5001
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.VideoGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestVideoGetWithBusinessScenarioSelf 测试携带 business_scenario 参数
func TestVideoGetWithBusinessScenarioSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.VideoGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.BusinessScenario = 1 // 内容素材包
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.VideoGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 验证测试用例 ==========

// TestVideoGetValidateMissingAccountAndOrgSelf 测试 account_id 和 organization_id 均未填写
func TestVideoGetValidateMissingAccountAndOrgSelf(t *testing.T) {
	req := &model.VideoGetReq{}
	req.AccessToken = "123"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id 和 organization_id 需必填其一")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestVideoGetValidateFilteringExceedMaxSelf 测试 filtering 超过4条
func TestVideoGetValidateFilteringExceedMaxSelf(t *testing.T) {
	req := &model.VideoGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.VideoFilteringItem{
		{Field: model.VideoFilterFieldMediaWidth, Operator: model.OperatorEquals, Values: []string{"640"}},
		{Field: model.VideoFilterFieldMediaHeight, Operator: model.OperatorEquals, Values: []string{"360"}},
		{Field: model.VideoFilterFieldSourceType, Operator: model.OperatorEquals, Values: []string{model.VideoSourceTypeLocal}},
		{Field: model.VideoFilterFieldStatus, Operator: model.OperatorEquals, Values: []string{model.VideoStatusNormal}},
		{Field: model.VideoFilterFieldAigcFlag, Operator: model.OperatorEquals, Values: []string{model.AigcFlagNotAI}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering超过4条")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestVideoGetValidateFilteringMissingFieldSelf 测试 filtering 缺少 field
func TestVideoGetValidateFilteringMissingFieldSelf(t *testing.T) {
	req := &model.VideoGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.VideoFilteringItem{
		{Operator: model.OperatorEquals, Values: []string{"640"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：field为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestVideoGetValidateFilteringEmptyValuesSelf 测试 filtering values 为空
func TestVideoGetValidateFilteringEmptyValuesSelf(t *testing.T) {
	req := &model.VideoGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Filtering = []*model.VideoFilteringItem{
		{Field: model.VideoFilterFieldStatus, Operator: model.OperatorEquals, Values: []string{}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：values为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestVideoGetValidatePageTooLargeSelf 测试 page 超过最大值 99999
func TestVideoGetValidatePageTooLargeSelf(t *testing.T) {
	req := &model.VideoGetReq{}
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

// TestVideoGetValidatePageSizeTooLargeSelf 测试 page_size 超过最大值 100
func TestVideoGetValidatePageSizeTooLargeSelf(t *testing.T) {
	req := &model.VideoGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Page = 1
	req.PageSize = 101
	req.Format()
	req.PageSize = 101
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size超过100")
	}
	fmt.Printf("验证错误: %v\n", err)
}
