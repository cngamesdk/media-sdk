package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

func boolPtr(b bool) *bool { return &b }

// 更新客户人群-更新名称
func TestCustomAudiencesUpdateName(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudiencesUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.Name = "新人群名称"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudiencesUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 更新客户人群-更新描述
func TestCustomAudiencesUpdateDescription(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudiencesUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.Description = "更新后的人群描述"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudiencesUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 更新客户人群-更新深度数据合作
func TestCustomAudiencesUpdateCooperated(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudiencesUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.Cooperated = boolPtr(true)
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudiencesUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 更新客户人群-同时更新名称、描述、深度合作
func TestCustomAudiencesUpdateAll(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudiencesUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.Name = "全量更新人群"
	req.Description = "全量更新描述"
	req.Cooperated = boolPtr(false)
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudiencesUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestCustomAudiencesUpdateValidateAccountIDEmpty(t *testing.T) {
	req := &model.CustomAudiencesUpdateReq{}
	req.AccessToken = "123"
	req.AudienceID = 123456789
	req.Name = "测试人群"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 audience_id
func TestCustomAudiencesUpdateValidateAudienceIDEmpty(t *testing.T) {
	req := &model.CustomAudiencesUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Name = "测试人群"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：audience_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-name/description/cooperated 均未填写
func TestCustomAudiencesUpdateValidateNoUpdateField(t *testing.T) {
	req := &model.CustomAudiencesUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：name、description、cooperated至少填写一个")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-name 超过最大长度
func TestCustomAudiencesUpdateValidateNameTooLong(t *testing.T) {
	req := &model.CustomAudiencesUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.Name = "这个名称超过了三十二个字节的限制测试一下到底行不行啊"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：name长度不能超过32字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-description 超过最大长度
func TestCustomAudiencesUpdateValidateDescriptionTooLong(t *testing.T) {
	req := &model.CustomAudiencesUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceID = 123456789
	req.Description = "这段描述文字超过了一百个字节的限制，用来测试校验逻辑是否正确生效，继续补充文字直到超出限制为止，超出超出超出超出超出超出超出超出超出超出"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：description长度不能超过100字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}
