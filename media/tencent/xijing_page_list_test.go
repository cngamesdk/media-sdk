package tencent

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ========== 蹊径获取落地页列表接口调用测试用例 ==========

// TestXijingPageListGetBasicSelf 测试基本查询（不带过滤条件）
func TestXijingPageListGetBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageListGetByPageIDSelf 测试按 page_id 查询
func TestXijingPageListGetByPageIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageID = 523034383985764607
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageListGetByPageTypeSelf 测试按 page_type 过滤查询
func TestXijingPageListGetByPageTypeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageType = []string{
		model.XijingPageTypeAndroidAppH5,
		model.XijingPageTypeIosAppH5,
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageListGetByPublishStatusSelf 测试按发布状态过滤查询
func TestXijingPageListGetByPublishStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PagePublishStatus = []string{model.XijingPagePublishStatusPublished}
	req.PageStatus = []string{model.XijingPageStatusApproved}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageListGetGrantSourceSelf 测试查询授权落地页
func TestXijingPageListGetGrantSourceSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageSource = model.XijingPageSourceGrant
	req.PageOwnerID = 204583434
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageListGetWithPaginationSelf 测试带分页参数查询
func TestXijingPageListGetWithPaginationSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageIndex = 2
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageListGetByModifyTimeSelf 测试按最后更新时间范围查询
func TestXijingPageListGetByModifyTimeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageLastModifyStartTime = "2024-01-01 00:00:00"
	req.PageLastModifyEndTime = "2024-12-31 23:59:59"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestXijingPageListGetDeletedSelf 测试查询已删除落地页
func TestXijingPageListGetDeletedSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.XijingPageListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.QueryType = model.XijingPageQueryTypeDeleted
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.XijingPageListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 蹊径获取落地页列表参数验证测试用例 ==========

// TestXijingPageListGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestXijingPageListGetValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.XijingPageListGetReq{}
	req.AccessToken = "123"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageListGetValidatePageNameTooLongSelf 测试 page_name 超过20字节
func TestXijingPageListGetValidatePageNameTooLongSelf(t *testing.T) {
	req := &model.XijingPageListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageName = strings.Repeat("a", 21)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_name超过20字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageListGetValidateServiceIDTooLongSelf 测试 page_service_id 超过256字节
func TestXijingPageListGetValidateServiceIDTooLongSelf(t *testing.T) {
	req := &model.XijingPageListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageServiceID = strings.Repeat("a", 257)
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_service_id超过256字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageListGetValidatePageTypeTooManySelf 测试 page_type 超过8个
func TestXijingPageListGetValidatePageTypeTooManySelf(t *testing.T) {
	req := &model.XijingPageListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageType = []string{
		model.XijingPageTypeDefaultH5,
		model.XijingPageTypeAndroidAppH5,
		model.XijingPageTypeIosAppH5,
		model.XijingPageTypeWebsiteH5,
		model.XijingPageTypeAndroidAppNative,
		model.XijingPageTypeIosAppNative,
		model.XijingPageTypeWebsiteNative,
		model.XijingPageTypeFenglingLbs,
		model.XijingPageTypeDefaultH5, // 第9个
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_type超过8个")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageListGetValidatePublishStatusTooManySelf 测试 page_publish_status 超过5个
func TestXijingPageListGetValidatePublishStatusTooManySelf(t *testing.T) {
	req := &model.XijingPageListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PagePublishStatus = []string{
		model.XijingPagePublishStatusUnpublish,
		model.XijingPagePublishStatusPublished,
		model.XijingPagePublishStatusOffline,
		model.XijingPagePublishStatusDeleting,
		model.XijingPagePublishStatusDeleted,
		model.XijingPagePublishStatusUnpublish, // 第6个
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_publish_status超过5个")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageListGetValidatePageIndexOutOfRangeSelf 测试 page_index 超出范围
func TestXijingPageListGetValidatePageIndexOutOfRangeSelf(t *testing.T) {
	req := &model.XijingPageListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageIndex = 100000
	req.PageSize = 10
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_index超出范围")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageListGetValidatePageSizeOutOfRangeSelf 测试 page_size 超出范围
func TestXijingPageListGetValidatePageSizeOutOfRangeSelf(t *testing.T) {
	req := &model.XijingPageListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageIndex = 1
	req.PageSize = 101
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size超出范围")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestXijingPageListGetValidateDefaultPaginationSelf 测试 Format() 默认填充分页参数
func TestXijingPageListGetValidateDefaultPaginationSelf(t *testing.T) {
	req := &model.XijingPageListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.Format()
	if req.PageIndex != 1 {
		t.Fatalf("期望默认page_index=1，实际=%d", req.PageIndex)
	}
	if req.PageSize != 10 {
		t.Fatalf("期望默认page_size=10，实际=%d", req.PageSize)
	}
	err := req.Validate()
	if err != nil {
		t.Fatalf("默认分页参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("默认分页参数验证通过")
}

// TestXijingPageListGetValidateFullParamsSelf 测试完整合法参数通过验证
func TestXijingPageListGetValidateFullParamsSelf(t *testing.T) {
	req := &model.XijingPageListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PageName = "测试落地页"
	req.PageType = []string{model.XijingPageTypeAndroidAppH5}
	req.PagePublishStatus = []string{model.XijingPagePublishStatusPublished}
	req.PageStatus = []string{model.XijingPageStatusApproved}
	req.PageSource = model.XijingPageSourceOwner
	req.AppType = model.XijingPageAppTypeAndroid
	req.QueryType = model.XijingPageQueryTypeDefault
	req.PageIndex = 1
	req.PageSize = 10
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("完整合法参数应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("完整合法参数验证通过")
}
