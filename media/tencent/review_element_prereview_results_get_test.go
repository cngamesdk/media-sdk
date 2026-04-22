package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 查询元素的预审结果-图片元素
func TestReviewElementPrereviewResultsGetImage(t *testing.T) {
	ctx := context.Background()
	req := &model.ReviewElementPrereviewResultsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Elements = []*model.PreReviewElement{
		{
			ElementType:    model.PreReviewElementTypeImage,
			ElementContent: "image_id_123456",
			ElementKey:     "IMG",
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ReviewElementPrereviewResultsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 查询元素的预审结果-多种元素类型
func TestReviewElementPrereviewResultsGetMultipleElements(t *testing.T) {
	ctx := context.Background()
	req := &model.ReviewElementPrereviewResultsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AdgroupID = 987654321
	req.Elements = []*model.PreReviewElement{
		{
			ElementType:    model.PreReviewElementTypeImage,
			ElementContent: "image_id_123456",
		},
		{
			ElementType:    model.PreReviewElementTypeTxt,
			ElementContent: "测试广告文案",
		},
		{
			ElementType:    model.PreReviewElementTypeDestUrl,
			ElementContent: "https://example.com/landing",
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ReviewElementPrereviewResultsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 查询元素的预审结果-带补充信息
func TestReviewElementPrereviewResultsGetWithSupplement(t *testing.T) {
	ctx := context.Background()
	req := &model.ReviewElementPrereviewResultsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Elements = []*model.PreReviewElement{
		{
			ElementType:    model.PreReviewElementTypeVideo,
			ElementContent: "video_id_789012",
		},
	}
	req.Supplement = []*model.PreReviewSupplement{
		{
			Field:    "site_set",
			Operator: "IN",
			Values:   []string{"SITE_SET_WECHAT", "SITE_SET_MOMENTS"},
		},
		{
			Field:    "is_dynamic_creative",
			Operator: "EQUALS",
			Values:   []string{"NOT_DYNAMIC_CREATIVE"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ReviewElementPrereviewResultsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestReviewElementPrereviewResultsGetValidateAccountIDEmpty(t *testing.T) {
	req := &model.ReviewElementPrereviewResultsGetReq{}
	req.AccessToken = "123"
	req.Elements = []*model.PreReviewElement{
		{ElementType: model.PreReviewElementTypeImage, ElementContent: "img_123"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 elements
func TestReviewElementPrereviewResultsGetValidateElementsEmpty(t *testing.T) {
	req := &model.ReviewElementPrereviewResultsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：elements为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-elements 超过最大长度
func TestReviewElementPrereviewResultsGetValidateElementsTooLong(t *testing.T) {
	req := &model.ReviewElementPrereviewResultsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	elements := make([]*model.PreReviewElement, 21)
	for i := range elements {
		elements[i] = &model.PreReviewElement{
			ElementType:    model.PreReviewElementTypeTxt,
			ElementContent: fmt.Sprintf("text_%d", i),
		}
	}
	req.Elements = elements
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：elements最大长度20")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-elements.element_type 为空
func TestReviewElementPrereviewResultsGetValidateElementTypeEmpty(t *testing.T) {
	req := &model.ReviewElementPrereviewResultsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Elements = []*model.PreReviewElement{
		{ElementType: "", ElementContent: "test"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：elements.element_type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-elements.element_content 为空
func TestReviewElementPrereviewResultsGetValidateElementContentEmpty(t *testing.T) {
	req := &model.ReviewElementPrereviewResultsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Elements = []*model.PreReviewElement{
		{ElementType: model.PreReviewElementTypeImage, ElementContent: ""},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：elements.element_content为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-supplement.values 为空
func TestReviewElementPrereviewResultsGetValidateSupplementValuesEmpty(t *testing.T) {
	req := &model.ReviewElementPrereviewResultsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Elements = []*model.PreReviewElement{
		{ElementType: model.PreReviewElementTypeImage, ElementContent: "img_123"},
	}
	req.Supplement = []*model.PreReviewSupplement{
		{Field: "site_set", Operator: "IN", Values: []string{}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：supplement.values为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 access_token
func TestReviewElementPrereviewResultsGetValidateAccessTokenEmpty(t *testing.T) {
	req := &model.ReviewElementPrereviewResultsGetReq{}
	req.AccountID = 2045867
	req.Elements = []*model.PreReviewElement{
		{ElementType: model.PreReviewElementTypeImage, ElementContent: "img_123"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：access_token为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
