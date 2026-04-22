package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 发起元素申诉复审-基本调用
func TestElementAppealReviewAdd(t *testing.T) {
	ctx := context.Background()
	req := &model.ElementAppealReviewAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.ComponentID = 987654321
	req.ElementID = 111222333
	req.ElementType = model.ElementTypeImage
	req.ElementValue = "test_image_value"
	req.ElementFingerPrint = "abc123def456"
	req.AppealDemand = "素材合规"
	req.AppealReason = "该素材已取得相关授权"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ElementAppealReviewAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 发起元素申诉复审-带可选字段
func TestElementAppealReviewAddWithOptional(t *testing.T) {
	ctx := context.Background()
	req := &model.ElementAppealReviewAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.ComponentID = 987654321
	req.ElementID = 111222333
	req.ElementType = model.ElementTypeText
	req.ElementValue = "test_text_value"
	req.ElementFingerPrint = "abc123"
	req.AppealDemand = "文案合规;内容真实"
	req.AppealReason = "文案内容均有据可查"
	req.HistoryApprovalComponentID = 555666777
	req.Description = "补充说明材料"
	req.ImageList = []string{"image_url_1", "image_url_2"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ElementAppealReviewAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestElementAppealReviewAddValidateAccountIDEmpty(t *testing.T) {
	req := &model.ElementAppealReviewAddReq{}
	req.AccessToken = "123"
	req.DynamicCreativeID = 123456789
	req.ComponentID = 987654321
	req.ElementID = 111222333
	req.ElementType = model.ElementTypeImage
	req.ElementValue = "test"
	req.ElementFingerPrint = "abc"
	req.AppealDemand = "素材合规"
	req.AppealReason = "已取得授权"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 element_type
func TestElementAppealReviewAddValidateElementTypeEmpty(t *testing.T) {
	req := &model.ElementAppealReviewAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.ComponentID = 987654321
	req.ElementID = 111222333
	req.ElementValue = "test"
	req.ElementFingerPrint = "abc"
	req.AppealDemand = "素材合规"
	req.AppealReason = "已取得授权"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：element_type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 appeal_demand
func TestElementAppealReviewAddValidateAppealDemandEmpty(t *testing.T) {
	req := &model.ElementAppealReviewAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.ComponentID = 987654321
	req.ElementID = 111222333
	req.ElementType = model.ElementTypeImage
	req.ElementValue = "test"
	req.ElementFingerPrint = "abc"
	req.AppealReason = "已取得授权"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：appeal_demand为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 appeal_reason
func TestElementAppealReviewAddValidateAppealReasonEmpty(t *testing.T) {
	req := &model.ElementAppealReviewAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.ComponentID = 987654321
	req.ElementID = 111222333
	req.ElementType = model.ElementTypeImage
	req.ElementValue = "test"
	req.ElementFingerPrint = "abc"
	req.AppealDemand = "素材合规"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：appeal_reason为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-image_list 超过最大长度
func TestElementAppealReviewAddValidateImageListTooLong(t *testing.T) {
	req := &model.ElementAppealReviewAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.ComponentID = 987654321
	req.ElementID = 111222333
	req.ElementType = model.ElementTypeImage
	req.ElementValue = "test"
	req.ElementFingerPrint = "abc"
	req.AppealDemand = "素材合规"
	req.AppealReason = "已取得授权"
	req.ImageList = []string{"img1", "img2", "img3", "img4"}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：image_list最大长度3")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-description 超过最大长度
func TestElementAppealReviewAddValidateDescriptionTooLong(t *testing.T) {
	req := &model.ElementAppealReviewAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.ComponentID = 987654321
	req.ElementID = 111222333
	req.ElementType = model.ElementTypeImage
	req.ElementValue = "test"
	req.ElementFingerPrint = "abc"
	req.AppealDemand = "素材合规"
	req.AppealReason = "已取得授权"
	req.Description = "这是一段超过五十字节限制的详细描述内容，用于测试字段长度验证逻辑是否正确生效"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：description长度最大50字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 access_token
func TestElementAppealReviewAddValidateAccessTokenEmpty(t *testing.T) {
	req := &model.ElementAppealReviewAddReq{}
	req.AccountID = 2045867
	req.DynamicCreativeID = 123456789
	req.ComponentID = 987654321
	req.ElementID = 111222333
	req.ElementType = model.ElementTypeImage
	req.ElementValue = "test"
	req.ElementFingerPrint = "abc"
	req.AppealDemand = "素材合规"
	req.AppealReason = "已取得授权"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：access_token为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
