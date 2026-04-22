package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 创意组件元素催审-按组件维度催审
func TestComponentElementUrgeReviewAddByComponent(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentElementUrgeReviewAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.UrgeDimension = model.UrgeDimensionComponent
	req.UrgeDimensionValue = "987654321"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentElementUrgeReviewAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 创意组件元素催审-按元素维度催审
func TestComponentElementUrgeReviewAddByElement(t *testing.T) {
	ctx := context.Background()
	req := &model.ComponentElementUrgeReviewAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.UrgeDimension = model.UrgeDimensionElement
	req.UrgeDimensionValue = "fp_abc123def456"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ComponentElementUrgeReviewAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestComponentElementUrgeReviewAddValidateAccountIDEmpty(t *testing.T) {
	req := &model.ComponentElementUrgeReviewAddReq{}
	req.AccessToken = "123"
	req.DynamicCreativeID = 123456789
	req.UrgeDimension = model.UrgeDimensionComponent
	req.UrgeDimensionValue = "987654321"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 dynamic_creative_id
func TestComponentElementUrgeReviewAddValidateDynamicCreativeIDEmpty(t *testing.T) {
	req := &model.ComponentElementUrgeReviewAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.UrgeDimension = model.UrgeDimensionComponent
	req.UrgeDimensionValue = "987654321"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：dynamic_creative_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 urge_dimension
func TestComponentElementUrgeReviewAddValidateUrgeDimensionEmpty(t *testing.T) {
	req := &model.ComponentElementUrgeReviewAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.UrgeDimensionValue = "987654321"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：urge_dimension为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-urge_dimension 值无效
func TestComponentElementUrgeReviewAddValidateUrgeDimensionInvalid(t *testing.T) {
	req := &model.ComponentElementUrgeReviewAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.UrgeDimension = "INVALID_DIMENSION"
	req.UrgeDimensionValue = "987654321"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：urge_dimension可选值为URGE_DIMENSION_COMPONENT、URGE_DIMENSION_ELEMENT")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 urge_dimension_value
func TestComponentElementUrgeReviewAddValidateUrgeDimensionValueEmpty(t *testing.T) {
	req := &model.ComponentElementUrgeReviewAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.UrgeDimension = model.UrgeDimensionComponent
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：urge_dimension_value为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-urge_dimension_value 超过最大长度
func TestComponentElementUrgeReviewAddValidateUrgeDimensionValueTooLong(t *testing.T) {
	req := &model.ComponentElementUrgeReviewAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.UrgeDimension = model.UrgeDimensionElement
	longStr := ""
	for i := 0; i < 130; i++ {
		longStr += "a"
	}
	req.UrgeDimensionValue = longStr
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：urge_dimension_value长度最大128字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 access_token
func TestComponentElementUrgeReviewAddValidateAccessTokenEmpty(t *testing.T) {
	req := &model.ComponentElementUrgeReviewAddReq{}
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.UrgeDimension = model.UrgeDimensionComponent
	req.UrgeDimensionValue = "987654321"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：access_token为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
