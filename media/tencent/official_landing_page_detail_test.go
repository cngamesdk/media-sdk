package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// buildOfficialLandingPageDetailBaseReq 构建基础获取落地页详情请求
func buildOfficialLandingPageDetailBaseReq() *model.OfficialLandingPageDetailGetReq {
	req := &model.OfficialLandingPageDetailGetReq{}
	req.AccessToken = "123"
	req.AccountId = 111111
	req.PageId = 999999
	req.ProtoVersion = 1
	return req
}

// ========== 官方落地页获取落地页详情接口调用测试用例 ==========

// TestOfficialLandingPageDetailGetBasicSelf 测试基本获取落地页详情
func TestOfficialLandingPageDetailGetBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageDetailBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageDetailGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageDetailGetWithProtoVersion0Self 测试使用 proto_version=0 获取
func TestOfficialLandingPageDetailGetWithProtoVersion0Self(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageDetailBaseReq()
	req.ProtoVersion = 0
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageDetailGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestOfficialLandingPageDetailGetWithProtoVersion1Self 测试使用 proto_version=1 获取
func TestOfficialLandingPageDetailGetWithProtoVersion1Self(t *testing.T) {
	ctx := context.Background()
	req := buildOfficialLandingPageDetailBaseReq()
	req.ProtoVersion = 1
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.OfficialLandingPageDetailGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 官方落地页获取落地页详情参数验证测试用例 ==========

// TestOfficialLandingPageDetailGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestOfficialLandingPageDetailGetValidateMissingAccountIDSelf(t *testing.T) {
	req := buildOfficialLandingPageDetailBaseReq()
	req.AccountId = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageDetailGetValidateMissingPageIDSelf 测试缺少 page_id
func TestOfficialLandingPageDetailGetValidateMissingPageIDSelf(t *testing.T) {
	req := buildOfficialLandingPageDetailBaseReq()
	req.PageId = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageDetailGetValidateInvalidProtoVersionSelf 测试非法的 proto_version
func TestOfficialLandingPageDetailGetValidateInvalidProtoVersionSelf(t *testing.T) {
	req := buildOfficialLandingPageDetailBaseReq()
	req.ProtoVersion = 2
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：proto_version须为0或1")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestOfficialLandingPageDetailGetValidateFullParamsSelf 测试完整合法参数通过验证
func TestOfficialLandingPageDetailGetValidateFullParamsSelf(t *testing.T) {
	req := buildOfficialLandingPageDetailBaseReq()
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}

// TestOfficialLandingPageDetailGetValidateDefaultProtoVersionSelf 测试默认 proto_version=0 通过验证
func TestOfficialLandingPageDetailGetValidateDefaultProtoVersionSelf(t *testing.T) {
	req := buildOfficialLandingPageDetailBaseReq()
	req.ProtoVersion = 0
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("proto_version=0 应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("proto_version=0 验证通过")
}
