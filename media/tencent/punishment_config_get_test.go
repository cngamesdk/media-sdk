package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取处罚系统配置-基本查询
func TestPunishmentConfigGet(t *testing.T) {
	ctx := context.Background()
	req := &model.PunishmentConfigGetReq{}
	req.AccessToken = "123"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.PunishmentConfigGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 access_token
func TestPunishmentConfigGetValidateAccessTokenEmpty(t *testing.T) {
	req := &model.PunishmentConfigGetReq{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：access_token为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
