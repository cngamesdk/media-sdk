package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取元素申诉复审结果-基本查询
func TestElementAppealReviewGet(t *testing.T) {
	ctx := context.Background()
	req := &model.ElementAppealReviewGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.ComponentID = 987654321
	req.ElementID = 111222333
	req.ElementFingerPrint = "abc123def456"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ElementAppealReviewGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取元素申诉复审结果-不传 element_finger_print
func TestElementAppealReviewGetNoFingerPrint(t *testing.T) {
	ctx := context.Background()
	req := &model.ElementAppealReviewGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.ComponentID = 987654321
	req.ElementID = 111222333
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ElementAppealReviewGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestElementAppealReviewGetValidateAccountIDEmpty(t *testing.T) {
	req := &model.ElementAppealReviewGetReq{}
	req.AccessToken = "123"
	req.DynamicCreativeID = 123456789
	req.ComponentID = 987654321
	req.ElementID = 111222333
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 dynamic_creative_id
func TestElementAppealReviewGetValidateDynamicCreativeIDEmpty(t *testing.T) {
	req := &model.ElementAppealReviewGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.ComponentID = 987654321
	req.ElementID = 111222333
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：dynamic_creative_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 component_id
func TestElementAppealReviewGetValidateComponentIDEmpty(t *testing.T) {
	req := &model.ElementAppealReviewGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.ElementID = 111222333
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：component_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 element_id
func TestElementAppealReviewGetValidateElementIDEmpty(t *testing.T) {
	req := &model.ElementAppealReviewGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.ComponentID = 987654321
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：element_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-element_finger_print 超过最大长度
func TestElementAppealReviewGetValidateFingerPrintTooLong(t *testing.T) {
	req := &model.ElementAppealReviewGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.ComponentID = 987654321
	req.ElementID = 111222333
	// 生成超过 128 字节的字符串
	longStr := ""
	for i := 0; i < 130; i++ {
		longStr += "a"
	}
	req.ElementFingerPrint = longStr
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：element_finger_print长度最大128字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 access_token
func TestElementAppealReviewGetValidateAccessTokenEmpty(t *testing.T) {
	req := &model.ElementAppealReviewGetReq{}
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.ComponentID = 987654321
	req.ElementID = 111222333
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：access_token为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
