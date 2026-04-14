package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// buildChannelsFinderobjectBaseReq 构建基础获取视频号动态详情请求
func buildChannelsFinderobjectBaseReq() *model.ChannelsFinderobjectGetReq {
	req := &model.ChannelsFinderobjectGetReq{}
	req.AccessToken = "123"
	req.AccountId = 111111
	req.ExportId = "fake_export_id_001"
	return req
}

// ========== 获取视频号动态详情接口调用测试用例 ==========

// TestChannelsFinderobjectGetSelf 测试获取视频号动态详情
func TestChannelsFinderobjectGetSelf(t *testing.T) {
	ctx := context.Background()
	req := buildChannelsFinderobjectBaseReq()
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ChannelsFinderobjectGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 获取视频号动态详情参数验证测试用例 ==========

// TestChannelsFinderobjectGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestChannelsFinderobjectGetValidateMissingAccountIDSelf(t *testing.T) {
	req := buildChannelsFinderobjectBaseReq()
	req.AccountId = 0
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestChannelsFinderobjectGetValidateMissingExportIDSelf 测试缺少 export_id
func TestChannelsFinderobjectGetValidateMissingExportIDSelf(t *testing.T) {
	req := buildChannelsFinderobjectBaseReq()
	req.ExportId = ""
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：export_id长度须在1-256字节之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestChannelsFinderobjectGetValidateExportIDTooLongSelf 测试 export_id 超出最大字节数
func TestChannelsFinderobjectGetValidateExportIDTooLongSelf(t *testing.T) {
	req := buildChannelsFinderobjectBaseReq()
	// 构造超过 256 字节的字符串
	longId := ""
	for i := 0; i < 260; i++ {
		longId += "a"
	}
	req.ExportId = longId
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：export_id长度须在1-256字节之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestChannelsFinderobjectGetValidateFullParamsSelf 测试完整合法参数通过验证
func TestChannelsFinderobjectGetValidateFullParamsSelf(t *testing.T) {
	req := buildChannelsFinderobjectBaseReq()
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}
