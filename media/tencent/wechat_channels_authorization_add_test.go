package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// buildWechatChannelsAuthorizationAddBaseReq 构建基础创建视频号授权请求
func buildWechatChannelsAuthorizationAddBaseReq() *model.WechatChannelsAuthorizationAddReq {
	req := &model.WechatChannelsAuthorizationAddReq{}
	req.AccessToken = "123"
	req.AccountId = 111111
	return req
}

// ========== 创建视频号授权接口调用测试用例 ==========

// TestWechatChannelsAuthorizationAddSelf 测试创建视频号授权（仅 account_id）
func TestWechatChannelsAuthorizationAddSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAuthorizationAddBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAuthorizationAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAuthorizationAddWithAccountNameSelf 测试带视频号名称
func TestWechatChannelsAuthorizationAddWithAccountNameSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAuthorizationAddBaseReq()
	req.WechatChannelsAccountName = "测试视频号"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAuthorizationAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAuthorizationAddWithCertificationListSelf 测试带资质列表
func TestWechatChannelsAuthorizationAddWithCertificationListSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAuthorizationAddBaseReq()
	req.AuthorizationCertificationList = []*model.WechatChannelsAuthorizationCertification{
		{
			CertificationCode:    "BUSINESS_LICENSE",
			CertificationImage:   "https://example.com/cert.png",
			CertificationNumber:  "91110000123456789X",
			CertificationImageId: "fake_image_id_001",
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAuthorizationAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAuthorizationAddWithScopeSelf 测试带授权范围
func TestWechatChannelsAuthorizationAddWithScopeSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAuthorizationAddBaseReq()
	req.AuthorizationScope = "ALL"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAuthorizationAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestWechatChannelsAuthorizationAddWithTimeSelf 测试带授权时间参数
func TestWechatChannelsAuthorizationAddWithTimeSelf(t *testing.T) {
	ctx := context.Background()
	req := buildWechatChannelsAuthorizationAddBaseReq()
	req.AuthorizationBeginTime = 1700000000
	req.AuthorizationTtl = 3122064000
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.WechatChannelsAuthorizationAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 创建视频号授权参数验证测试用例 ==========

// TestWechatChannelsAuthorizationAddValidateMissingAccountIDSelf 测试缺少 account_id
func TestWechatChannelsAuthorizationAddValidateMissingAccountIDSelf(t *testing.T) {
	req := buildWechatChannelsAuthorizationAddBaseReq()
	req.AccountId = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAuthorizationAddValidateCertificationListTooLongSelf 测试资质列表超出最大长度
func TestWechatChannelsAuthorizationAddValidateCertificationListTooLongSelf(t *testing.T) {
	req := buildWechatChannelsAuthorizationAddBaseReq()
	list := make([]*model.WechatChannelsAuthorizationCertification, 256)
	for i := range list {
		list[i] = &model.WechatChannelsAuthorizationCertification{CertificationCode: "BUSINESS_LICENSE"}
	}
	req.AuthorizationCertificationList = list
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：authorization_certification_list最大长度为255")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestWechatChannelsAuthorizationAddValidateMissingCertificationCodeSelf 测试资质列表中缺少 certification_code
func TestWechatChannelsAuthorizationAddValidateMissingCertificationCodeSelf(t *testing.T) {
	req := buildWechatChannelsAuthorizationAddBaseReq()
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

// TestWechatChannelsAuthorizationAddValidateFullParamsSelf 测试完整合法参数通过验证
func TestWechatChannelsAuthorizationAddValidateFullParamsSelf(t *testing.T) {
	req := buildWechatChannelsAuthorizationAddBaseReq()
	req.AuthorizationCertificationList = []*model.WechatChannelsAuthorizationCertification{
		{CertificationCode: "BUSINESS_LICENSE"},
	}
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}
