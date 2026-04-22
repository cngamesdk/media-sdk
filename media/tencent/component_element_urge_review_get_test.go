package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取创意组件元素催审状态-基本查询（仅必填参数）
func TestComponentElementUrgeReviewGet(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentElementUrgeReviewGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentElementUrgeReviewGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取创意组件元素催审状态-带 component_id_list
func TestComponentElementUrgeReviewGetWithComponentIDs(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentElementUrgeReviewGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.ComponentIDList = []int64{111222333, 444555666}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentElementUrgeReviewGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取创意组件元素催审状态-带 element_fingerprint_list
func TestComponentElementUrgeReviewGetWithFingerprints(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentElementUrgeReviewGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.ElementFingerprintList = []string{"fp_abc123", "fp_def456"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentElementUrgeReviewGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取创意组件元素催审状态-同时带 component_id_list 和 element_fingerprint_list
func TestComponentElementUrgeReviewGetWithBoth(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentElementUrgeReviewGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.ComponentIDList = []int64{111222333}
	req.ElementFingerprintList = []string{"fp_abc123"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentElementUrgeReviewGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestComponentElementUrgeReviewGetValidateAccountIDEmpty(t *testing.T) {
	req := &model.ComponentElementUrgeReviewGetReq{}
	req.AccessToken = "123"
	req.DynamicCreativeID = 123456789
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 dynamic_creative_id
func TestComponentElementUrgeReviewGetValidateDynamicCreativeIDEmpty(t *testing.T) {
	req := &model.ComponentElementUrgeReviewGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：dynamic_creative_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-component_id_list 超过最大长度
func TestComponentElementUrgeReviewGetValidateComponentIDListTooLong(t *testing.T) {
	req := &model.ComponentElementUrgeReviewGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	ids := make([]int64, 101)
	for i := range ids {
		ids[i] = int64(i + 1)
	}
	req.ComponentIDList = ids
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：component_id_list最大长度100")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-element_fingerprint_list 超过最大长度
func TestComponentElementUrgeReviewGetValidateFingerprintListTooLong(t *testing.T) {
	req := &model.ComponentElementUrgeReviewGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	fps := make([]string, 101)
	for i := range fps {
		fps[i] = fmt.Sprintf("fp_%d", i)
	}
	req.ElementFingerprintList = fps
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：element_fingerprint_list最大长度100")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-element_fingerprint_list 元素超过最大长度
func TestComponentElementUrgeReviewGetValidateFingerprintTooLong(t *testing.T) {
	req := &model.ComponentElementUrgeReviewGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	longStr := ""
	for i := 0; i < 130; i++ {
		longStr += "a"
	}
	req.ElementFingerprintList = []string{longStr}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：element_fingerprint_list元素长度最大128字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 access_token
func TestComponentElementUrgeReviewGetValidateAccessTokenEmpty(t *testing.T) {
	req := &model.ComponentElementUrgeReviewGetReq{}
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：access_token为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
