package tencent

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ========== 创建批量请求测试用例 ==========

// TestBatchRequestAddSinglePostSelf 测试批量执行单个 POST 请求
func TestBatchRequestAddSinglePostSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchRequestAddReq{}
	req.AccessToken = "123"
	req.BatchRequestSpec = []*model.BatchRequestSpec{
		{
			RelativePath: "v3.0/adgroups/update",
			Body:         `{"account_id":20458,"adgroup_id":5075765587,"daily_budget":5000}`,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchRequestAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBatchRequestAddMultiplePostSelf 测试批量执行多个 POST 请求
func TestBatchRequestAddMultiplePostSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchRequestAddReq{}
	req.AccessToken = "123"
	req.BatchRequestSpec = []*model.BatchRequestSpec{
		{
			RelativePath: "v3.0/dynamic_creatives/update",
			Body:         `{"adgroup_id":5076530760,"dynamic_creative_id":41006154,"configured_status":"AD_STATUS_NORMAL","account_id":20458}`,
		},
		{
			RelativePath: "v3.0/dynamic_creatives/update",
			Body:         `{"adgroup_id":5076530876,"dynamic_creative_id":41006147,"configured_status":"AD_STATUS_SUSPEND","account_id":20458}`,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchRequestAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBatchRequestAddGetRequestSelf 测试批量执行 GET 请求（relative_path 含参数，无 body）
func TestBatchRequestAddGetRequestSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchRequestAddReq{}
	req.AccessToken = "123"
	req.BatchRequestSpec = []*model.BatchRequestSpec{
		{
			RelativePath: "v3.0/adgroups/get?account_id=20458&adgroup_ids=[5075765587]",
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchRequestAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBatchRequestAddMixedSelf 测试批量混合请求（含 body 序列化构造）
func TestBatchRequestAddMixedSelf(t *testing.T) {
	ctx := context.Background()

	// 通过 json.Marshal 构造 body 字符串，避免手写 JSON
	type adgroupUpdateBody struct {
		AccountID   int64 `json:"account_id"`
		AdgroupID   int64 `json:"adgroup_id"`
		DailyBudget int   `json:"daily_budget"`
	}
	body1, _ := json.Marshal(adgroupUpdateBody{AccountID: 20458, AdgroupID: 5075765587, DailyBudget: 5000})
	body2, _ := json.Marshal(adgroupUpdateBody{AccountID: 20458, AdgroupID: 5075765588, DailyBudget: 8000})

	req := &model.BatchRequestAddReq{}
	req.AccessToken = "123"
	req.BatchRequestSpec = []*model.BatchRequestSpec{
		{RelativePath: "v3.0/adgroups/update", Body: string(body1)},
		{RelativePath: "v3.0/adgroups/update", Body: string(body2)},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchRequestAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBatchRequestAddAdvertiserUpdateSelf 测试批量更新广告主信息
func TestBatchRequestAddAdvertiserUpdateSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BatchRequestAddReq{}
	req.AccessToken = "123"
	req.BatchRequestSpec = []*model.BatchRequestSpec{
		{
			RelativePath: "v3.0/advertiser/update",
			Body:         `{"account_id":20458,"daily_budget":100000}`,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BatchRequestAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 验证测试用例 ==========

// TestBatchRequestAddValidateEmptySpecSelf 测试 batch_request_spec 为空
func TestBatchRequestAddValidateEmptySpecSelf(t *testing.T) {
	req := &model.BatchRequestAddReq{}
	req.AccessToken = "123"
	req.BatchRequestSpec = []*model.BatchRequestSpec{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：batch_request_spec至少包含1个请求条件")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBatchRequestAddValidateNilSpecSelf 测试 batch_request_spec 为 nil
func TestBatchRequestAddValidateNilSpecSelf(t *testing.T) {
	req := &model.BatchRequestAddReq{}
	req.AccessToken = "123"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：batch_request_spec至少包含1个请求条件")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBatchRequestAddValidateMissingRelativePathSelf 测试 relative_path 为空
func TestBatchRequestAddValidateMissingRelativePathSelf(t *testing.T) {
	req := &model.BatchRequestAddReq{}
	req.AccessToken = "123"
	req.BatchRequestSpec = []*model.BatchRequestSpec{
		{Body: `{"account_id":20458}`},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：relative_path为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestBatchRequestAddValidateNilSpecItemSelf 测试 batch_request_spec 含 nil 项
func TestBatchRequestAddValidateNilSpecItemSelf(t *testing.T) {
	req := &model.BatchRequestAddReq{}
	req.AccessToken = "123"
	req.BatchRequestSpec = []*model.BatchRequestSpec{
		{RelativePath: "v3.0/adgroups/update"},
		nil,
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：spec 项不能为 nil")
	}
	fmt.Printf("验证错误: %v\n", err)
}
