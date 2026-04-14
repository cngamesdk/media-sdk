package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// buildWechatChannelsAuthorizationUpdateBaseReq 构建基础更新视频号授权请求
func buildWechatChannelsAuthorizationUpdateBaseReq() *model.WechatChannelsAuthorizationUpdateReq {
	req := &model.WechatChannelsAuthorizationUpdateReq{}
	req.AccessToken = "123"
	req.AccountId = 111111
	req.AuthorizationId = "fake_authorization_id_001"
	req.AuthorizationBeginTime = 1700000000
	req.AuthorizationTtl = 3122064000
	req.AuthorizationRelationship = "RELATIONSHIP_CORPORATION"
	req.AuthorizationCertificationList = []*model.WechatChannelsAuthorizationCertification{
		{CertificationCode: "BUSINESS_LICENSE"},
	}
	return req
}

// ========== 更新视频号授权接口调用测试用例 ==========

// TestWechatChannelsAuthorizationUpdateSelf 测试更新视频号授权（基础必填参数）
func TestWechatChannelsAuthorizationUpdateSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAuthorizationUpdateBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAuthorizationUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAuthorizationUpdateWithScopeSelf 测试带授权范围
func TestWechatChannelsAuthorizationUpdateWithScopeSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAuthorizationUpdateBaseReq()
	req.AuthorizationScope = "ALL"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAuthorizationUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAuthorizationUpdateWithMultipleCertsSelf 测试多条资质
func TestWechatChannelsAuthorizationUpdateWithMultipleCertsSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAuthorizationUpdateBaseReq()
	req.AuthorizationCertificationList = []*model.WechatChannelsAuthorizationCertification{
		{
			CertificationCode:    "BUSINESS_LICENSE",
			CertificationImage:   "https://example.com/cert1.png",
			CertificationNumber:  "91110000123456789X",
			CertificationImageId: "fake_image_id_001",
		},
		{
			CertificationCode: "OTHER_CERT",
			CertificationName: "其他资质",
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAuthorizationUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 更新视频号授权参数验证测试用例 ==========

// TestWechatChannelsAuthorizationUpdateValidateMissingAccountIDSelf 测试缺少 account_id
func TestWechatChannelsAuthorizationUpdateValidateMissingAccountIDSelf(t *testing.T) {
	req := buildWechatChannelsAuthorizationUpdateBaseReq()
	req.AccountId = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAuthorizationUpdateValidateMissingAuthorizationIDSelf 测试缺少 authorization_id
func TestWechatChannelsAuthorizationUpdateValidateMissingAuthorizationIDSelf(t *testing.T) {
	req := buildWechatChannelsAuthorizationUpdateBaseReq()
	req.AuthorizationId = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：authorization_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAuthorizationUpdateValidateMissingRelationshipSelf 测试缺少 authorization_relationship
func TestWechatChannelsAuthorizationUpdateValidateMissingRelationshipSelf(t *testing.T) {
	req := buildWechatChannelsAuthorizationUpdateBaseReq()
	req.AuthorizationRelationship = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：authorization_relationship为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAuthorizationUpdateValidateMissingCertificationListSelf 测试缺少 authorization_certification_list
func TestWechatChannelsAuthorizationUpdateValidateMissingCertificationListSelf(t *testing.T) {
	req := buildWechatChannelsAuthorizationUpdateBaseReq()
	req.AuthorizationCertificationList = nil
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：authorization_certification_list为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAuthorizationUpdateValidateMissingCertificationCodeSelf 测试资质列表中缺少 certification_code
func TestWechatChannelsAuthorizationUpdateValidateMissingCertificationCodeSelf(t *testing.T) {
	req := buildWechatChannelsAuthorizationUpdateBaseReq()
	req.AuthorizationCertificationList = []*model.WechatChannelsAuthorizationCertification{
		{CertificationImage: "https://example.com/cert.png"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：certification_code为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAuthorizationUpdateValidateFullParamsSelf 测试完整合法参数通过验证
func TestWechatChannelsAuthorizationUpdateValidateFullParamsSelf(t *testing.T) {
	req := buildWechatChannelsAuthorizationUpdateBaseReq()
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}
