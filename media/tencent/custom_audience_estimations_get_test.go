package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 人群覆盖数预估-仅 include（OR 关系）
func TestCustomAudienceEstimationsGetIncludeOnly(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudienceEstimationsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Type = "COMBINE"
	req.AudienceSpec = &model.EstimationAudienceSpec{
		CombineSpec: &model.CombineSpec{
			Include: [][]model.CombineAudienceItem{
				{
					{AudienceID: 1234567},
					{AudienceID: 7654321},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudienceEstimationsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 人群覆盖数预估-include AND 关系（多组）
func TestCustomAudienceEstimationsGetIncludeAnd(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudienceEstimationsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Type = "COMBINE"
	req.AudienceSpec = &model.EstimationAudienceSpec{
		CombineSpec: &model.CombineSpec{
			Include: [][]model.CombineAudienceItem{
				{{AudienceID: 1234567}},
				{{AudienceID: 7654321}},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudienceEstimationsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 人群覆盖数预估-include + exclude
func TestCustomAudienceEstimationsGetIncludeExclude(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudienceEstimationsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Type = "COMBINE"
	req.AudienceSpec = &model.EstimationAudienceSpec{
		CombineSpec: &model.CombineSpec{
			Include: [][]model.CombineAudienceItem{
				{
					{AudienceID: 1234567},
					{AudienceID: 7654321},
				},
			},
			Exclude: [][]model.CombineAudienceItem{
				{
					{AudienceID: 9999999},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudienceEstimationsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 人群覆盖数预估-含 time_window（客户标签）
func TestCustomAudienceEstimationsGetWithTimeWindow(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudienceEstimationsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Type = "COMBINE"
	req.AudienceSpec = &model.EstimationAudienceSpec{
		CombineSpec: &model.CombineSpec{
			Include: [][]model.CombineAudienceItem{
				{
					{AudienceID: 1234567, TimeWindow: 30},
					{AudienceID: 7654321, TimeWindow: 7},
				},
			},
			Exclude: [][]model.CombineAudienceItem{
				{
					{AudienceID: 9999999, TimeWindow: 90},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudienceEstimationsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestCustomAudienceEstimationsGetValidateAccountIDEmpty(t *testing.T) {
	req := &model.CustomAudienceEstimationsGetReq{}
	req.AccessToken = "123"
	req.Type = "COMBINE"
	req.AudienceSpec = &model.EstimationAudienceSpec{
		CombineSpec: &model.CombineSpec{
			Include: [][]model.CombineAudienceItem{
				{{AudienceID: 1234567}},
			},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 type
func TestCustomAudienceEstimationsGetValidateTypeEmpty(t *testing.T) {
	req := &model.CustomAudienceEstimationsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceSpec = &model.EstimationAudienceSpec{
		CombineSpec: &model.CombineSpec{
			Include: [][]model.CombineAudienceItem{
				{{AudienceID: 1234567}},
			},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-type 值不支持
func TestCustomAudienceEstimationsGetValidateTypeInvalid(t *testing.T) {
	req := &model.CustomAudienceEstimationsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Type = "LOOKALIKE"
	req.AudienceSpec = &model.EstimationAudienceSpec{
		CombineSpec: &model.CombineSpec{
			Include: [][]model.CombineAudienceItem{
				{{AudienceID: 1234567}},
			},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：type目前仅支持COMBINE")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 audience_spec
func TestCustomAudienceEstimationsGetValidateAudienceSpecNil(t *testing.T) {
	req := &model.CustomAudienceEstimationsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Type = "COMBINE"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：audience_spec为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-type=COMBINE 时缺少 combine_spec
func TestCustomAudienceEstimationsGetValidateCombineSpecNil(t *testing.T) {
	req := &model.CustomAudienceEstimationsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Type = "COMBINE"
	req.AudienceSpec = &model.EstimationAudienceSpec{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：combine_spec为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-combine_spec.include 为空
func TestCustomAudienceEstimationsGetValidateIncludeEmpty(t *testing.T) {
	req := &model.CustomAudienceEstimationsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Type = "COMBINE"
	req.AudienceSpec = &model.EstimationAudienceSpec{
		CombineSpec: &model.CombineSpec{},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：combine_spec.include为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
