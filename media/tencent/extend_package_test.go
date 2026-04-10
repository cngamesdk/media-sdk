package tencent

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ========== 创建应用分包接口调用测试用例 ==========

// TestExtendPackageAddSingleChannelSelf 测试创建单个渠道包
func TestExtendPackageAddSingleChannelSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ExtendPackageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageChannelItem{
		{
			ChannelID:   "927684_.-abc475913",
			ChannelName: "渠道包2000000336",
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ExtendPackageAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestExtendPackageAddMultiChannelSelf 测试批量创建多个渠道包
func TestExtendPackageAddMultiChannelSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ExtendPackageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageChannelItem{
		{ChannelID: "channel001", ChannelName: "渠道包001"},
		{ChannelID: "channel002", ChannelName: "渠道包002"},
		{ChannelID: "channel003", ChannelName: "渠道包003"},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ExtendPackageAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestExtendPackageAddWithCustomizedChannelIDSelf 测试携带 customized_channel_id
func TestExtendPackageAddWithCustomizedChannelIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ExtendPackageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageChannelItem{
		{
			ChannelID:           "channel_custom_01",
			ChannelName:         "自定义渠道包",
			CustomizedChannelID: "channel_custom_01",
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ExtendPackageAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestExtendPackageAddWithDefaultChannelNameSelf 测试不传 channel_name（使用默认值）
func TestExtendPackageAddWithDefaultChannelNameSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ExtendPackageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageChannelItem{
		{ChannelID: "channel_no_name"},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ExtendPackageAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 创建应用分包参数验证测试用例 ==========

// TestExtendPackageAddValidateMissingAccountIDSelf 测试缺少 account_id
func TestExtendPackageAddValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.ExtendPackageAddReq{}
	req.AccessToken = "123"
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageChannelItem{
		{ChannelID: "channel001"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageAddValidateMissingPackageIDSelf 测试缺少 package_id（值为0时视为未填，但0是合法值）
// package_id 允许为 0，此用例测试负数场景
func TestExtendPackageAddValidateNegativePackageIDSelf(t *testing.T) {
	req := &model.ExtendPackageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = -1
	req.ChannelList = []*model.ExtendPackageChannelItem{
		{ChannelID: "channel001"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：package_id不能为负数")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageAddValidateEmptyChannelListSelf 测试 channel_list 为空
func TestExtendPackageAddValidateEmptyChannelListSelf(t *testing.T) {
	req := &model.ExtendPackageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageChannelItem{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：channel_list不能为空")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageAddValidateChannelListTooLongSelf 测试 channel_list 超过200条
func TestExtendPackageAddValidateChannelListTooLongSelf(t *testing.T) {
	req := &model.ExtendPackageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	list := make([]*model.ExtendPackageChannelItem, 201)
	for i := range list {
		list[i] = &model.ExtendPackageChannelItem{ChannelID: fmt.Sprintf("channel%03d", i)}
	}
	req.ChannelList = list
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：channel_list超过200条")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageAddValidateMissingChannelIDSelf 测试 channel_id 为空
func TestExtendPackageAddValidateMissingChannelIDSelf(t *testing.T) {
	req := &model.ExtendPackageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageChannelItem{
		{ChannelName: "渠道包001"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：channel_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageAddValidateChannelIDInvalidCharsSelf 测试 channel_id 含非法字符（中文）
func TestExtendPackageAddValidateChannelIDInvalidCharsSelf(t *testing.T) {
	req := &model.ExtendPackageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageChannelItem{
		{ChannelID: "渠道包非法字符"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：channel_id含非法字符")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageAddValidateChannelIDInvalidSpaceSelf 测试 channel_id 含空格
func TestExtendPackageAddValidateChannelIDInvalidSpaceSelf(t *testing.T) {
	req := &model.ExtendPackageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageChannelItem{
		{ChannelID: "channel id with space"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：channel_id含空格")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageAddValidateChannelIDTooLongSelf 测试 channel_id 超过200字节
func TestExtendPackageAddValidateChannelIDTooLongSelf(t *testing.T) {
	req := &model.ExtendPackageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageChannelItem{
		{ChannelID: strings.Repeat("a", 201)},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：channel_id超过200字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageAddValidateChannelNameTooLongSelf 测试 channel_name 超过255字节
func TestExtendPackageAddValidateChannelNameTooLongSelf(t *testing.T) {
	req := &model.ExtendPackageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageChannelItem{
		{
			ChannelID:   "channel001",
			ChannelName: strings.Repeat("a", 256),
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：channel_name超过255字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageAddValidateCustomizedChannelIDTooLongSelf 测试 customized_channel_id 超过256字节
func TestExtendPackageAddValidateCustomizedChannelIDTooLongSelf(t *testing.T) {
	req := &model.ExtendPackageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageChannelItem{
		{
			ChannelID:           "channel001",
			CustomizedChannelID: strings.Repeat("a", 257),
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：customized_channel_id超过256字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageAddValidateNilChannelItemSelf 测试 channel_list 中包含 nil 元素
func TestExtendPackageAddValidateNilChannelItemSelf(t *testing.T) {
	req := &model.ExtendPackageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageChannelItem{
		{ChannelID: "channel001"},
		nil,
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：channel_list中包含nil元素")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageAddValidateChannelIDValidCharsSelf 测试 channel_id 合法字符（字母、数字、_.-）均可通过
func TestExtendPackageAddValidateChannelIDValidCharsSelf(t *testing.T) {
	req := &model.ExtendPackageAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageChannelItem{
		{ChannelID: "Abc123_.-XYZ"},
	}
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("合法channel_id应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("合法channel_id验证通过")
}

// ========== 更新应用子包版本接口调用测试用例 ==========

// TestExtendPackageUpdateSingleChannelSelf 测试更新单个渠道包名称
func TestExtendPackageUpdateSingleChannelSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ExtendPackageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageUpdateChannelItem{
		{
			ChannelID:   "927684_.-abc475913",
			ChannelName: "渠道包5185813111",
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ExtendPackageUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestExtendPackageUpdateMultiChannelSelf 测试批量更新多个渠道包
func TestExtendPackageUpdateMultiChannelSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ExtendPackageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageUpdateChannelItem{
		{ChannelID: "channel001", ChannelName: "渠道包更新001"},
		{ChannelID: "channel002", ChannelName: "渠道包更新002"},
		{ChannelID: "channel003", ChannelName: "渠道包更新003"},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ExtendPackageUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestExtendPackageUpdateWithoutChannelNameSelf 测试不传 channel_name（仅更新渠道标识匹配）
func TestExtendPackageUpdateWithoutChannelNameSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ExtendPackageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageUpdateChannelItem{
		{ChannelID: "channel001"},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ExtendPackageUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 更新应用子包版本参数验证测试用例 ==========

// TestExtendPackageUpdateValidateMissingAccountIDSelf 测试缺少 account_id
func TestExtendPackageUpdateValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.ExtendPackageUpdateReq{}
	req.AccessToken = "123"
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageUpdateChannelItem{
		{ChannelID: "channel001"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageUpdateValidateNegativePackageIDSelf 测试 package_id 为负数
func TestExtendPackageUpdateValidateNegativePackageIDSelf(t *testing.T) {
	req := &model.ExtendPackageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = -1
	req.ChannelList = []*model.ExtendPackageUpdateChannelItem{
		{ChannelID: "channel001"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：package_id不能为负数")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageUpdateValidateEmptyChannelListSelf 测试 channel_list 为空
func TestExtendPackageUpdateValidateEmptyChannelListSelf(t *testing.T) {
	req := &model.ExtendPackageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageUpdateChannelItem{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：channel_list不能为空")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageUpdateValidateChannelListTooLongSelf 测试 channel_list 超过200条
func TestExtendPackageUpdateValidateChannelListTooLongSelf(t *testing.T) {
	req := &model.ExtendPackageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	list := make([]*model.ExtendPackageUpdateChannelItem, 201)
	for i := range list {
		list[i] = &model.ExtendPackageUpdateChannelItem{ChannelID: fmt.Sprintf("channel%03d", i)}
	}
	req.ChannelList = list
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：channel_list超过200条")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageUpdateValidateMissingChannelIDSelf 测试 channel_id 为空
func TestExtendPackageUpdateValidateMissingChannelIDSelf(t *testing.T) {
	req := &model.ExtendPackageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageUpdateChannelItem{
		{ChannelName: "渠道包001"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：channel_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageUpdateValidateChannelIDInvalidCharsSelf 测试 channel_id 含非法字符
func TestExtendPackageUpdateValidateChannelIDInvalidCharsSelf(t *testing.T) {
	req := &model.ExtendPackageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageUpdateChannelItem{
		{ChannelID: "渠道@非法"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：channel_id含非法字符")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageUpdateValidateChannelIDTooLongSelf 测试 channel_id 超过200字节
func TestExtendPackageUpdateValidateChannelIDTooLongSelf(t *testing.T) {
	req := &model.ExtendPackageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageUpdateChannelItem{
		{ChannelID: strings.Repeat("a", 201)},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：channel_id超过200字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageUpdateValidateChannelNameTooLongSelf 测试 channel_name 超过1024字节
func TestExtendPackageUpdateValidateChannelNameTooLongSelf(t *testing.T) {
	req := &model.ExtendPackageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageUpdateChannelItem{
		{
			ChannelID:   "channel001",
			ChannelName: strings.Repeat("a", 1025),
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：channel_name超过1024字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageUpdateValidateChannelNameMaxBoundarySelf 测试 channel_name 恰好1024字节（边界值）
func TestExtendPackageUpdateValidateChannelNameMaxBoundarySelf(t *testing.T) {
	req := &model.ExtendPackageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageUpdateChannelItem{
		{
			ChannelID:   "channel001",
			ChannelName: strings.Repeat("a", 1024), // 恰好1024字节，合法
		},
	}
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("1024字节channel_name应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("1024字节channel_name验证通过")
}

// TestExtendPackageUpdateValidateNilChannelItemSelf 测试 channel_list 中包含 nil 元素
func TestExtendPackageUpdateValidateNilChannelItemSelf(t *testing.T) {
	req := &model.ExtendPackageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageUpdateChannelItem{
		{ChannelID: "channel001"},
		nil,
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：channel_list中包含nil元素")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageUpdateValidateChannelIDValidCharsSelf 测试 channel_id 合法字符正向验证
func TestExtendPackageUpdateValidateChannelIDValidCharsSelf(t *testing.T) {
	req := &model.ExtendPackageUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.ChannelList = []*model.ExtendPackageUpdateChannelItem{
		{ChannelID: "Abc123_.-XYZ", ChannelName: "合法渠道包"},
	}
	req.Format()
	err := req.Validate()
	if err != nil {
		t.Fatalf("合法channel_id应通过验证，但返回了错误: %v", err)
	}
	fmt.Println("合法channel_id验证通过")
}

// ========== 查询应用分包列表接口调用测试用例 ==========

// TestExtendPackageGetBasicSelf 测试基本查询（不带过滤条件）
func TestExtendPackageGetBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ExtendPackageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ExtendPackageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestExtendPackageGetWithPaginationSelf 测试带分页参数查询
func TestExtendPackageGetWithPaginationSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ExtendPackageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.Page = 1
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ExtendPackageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestExtendPackageGetFilterByChannelPackageIDSelf 测试按 channel_package_id 精确查询
func TestExtendPackageGetFilterByChannelPackageIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ExtendPackageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.Filtering = []*model.ExtendPackageGetFilteringItem{
		{
			Field:    model.ExtendPackageGetFilterFieldChannelPackageID,
			Operator: model.ExtendPackageGetFilterOperatorEquals,
			Values:   []string{"12345678"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ExtendPackageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestExtendPackageGetFilterByChannelNameSelf 测试按 channel_name 模糊查询
func TestExtendPackageGetFilterByChannelNameSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ExtendPackageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.Filtering = []*model.ExtendPackageGetFilteringItem{
		{
			Field:    model.ExtendPackageGetFilterFieldChannelName,
			Operator: model.ExtendPackageGetFilterOperatorContains,
			Values:   []string{"渠道包"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ExtendPackageGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 查询应用分包列表参数验证测试用例 ==========

// TestExtendPackageGetValidateMissingAccountIDSelf 测试缺少 account_id
func TestExtendPackageGetValidateMissingAccountIDSelf(t *testing.T) {
	req := &model.ExtendPackageGetReq{}
	req.AccessToken = "123"
	req.PackageID = 2000000336
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageGetValidateNegativePackageIDSelf 测试 package_id 为负数
func TestExtendPackageGetValidateNegativePackageIDSelf(t *testing.T) {
	req := &model.ExtendPackageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = -1
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：package_id不能为负数")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageGetValidateFilteringTooLongSelf 测试 filtering 超过2条
func TestExtendPackageGetValidateFilteringTooLongSelf(t *testing.T) {
	req := &model.ExtendPackageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.Filtering = []*model.ExtendPackageGetFilteringItem{
		{Field: model.ExtendPackageGetFilterFieldChannelPackageID, Operator: model.ExtendPackageGetFilterOperatorEquals, Values: []string{"1"}},
		{Field: model.ExtendPackageGetFilterFieldChannelName, Operator: model.ExtendPackageGetFilterOperatorContains, Values: []string{"test"}},
		{Field: model.ExtendPackageGetFilterFieldChannelPackageID, Operator: model.ExtendPackageGetFilterOperatorEquals, Values: []string{"2"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering超过2条")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageGetValidateNilFilterItemSelf 测试 filtering 中包含 nil 元素
func TestExtendPackageGetValidateNilFilterItemSelf(t *testing.T) {
	req := &model.ExtendPackageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.Filtering = []*model.ExtendPackageGetFilteringItem{
		{Field: model.ExtendPackageGetFilterFieldChannelPackageID, Operator: model.ExtendPackageGetFilterOperatorEquals, Values: []string{"1"}},
		nil,
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering中包含nil元素")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageGetValidateFilterMissingFieldSelf 测试 filtering.field 为空
func TestExtendPackageGetValidateFilterMissingFieldSelf(t *testing.T) {
	req := &model.ExtendPackageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.Filtering = []*model.ExtendPackageGetFilteringItem{
		{Operator: model.ExtendPackageGetFilterOperatorEquals, Values: []string{"1"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.field为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageGetValidateFilterEmptyValuesSelf 测试 filtering.values 为空
func TestExtendPackageGetValidateFilterEmptyValuesSelf(t *testing.T) {
	req := &model.ExtendPackageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.Filtering = []*model.ExtendPackageGetFilteringItem{
		{Field: model.ExtendPackageGetFilterFieldChannelPackageID, Operator: model.ExtendPackageGetFilterOperatorEquals, Values: []string{}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.values为空")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageGetValidateFilterValuesTooManySelf 测试 filtering.values 超过1个
func TestExtendPackageGetValidateFilterValuesTooManySelf(t *testing.T) {
	req := &model.ExtendPackageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.Filtering = []*model.ExtendPackageGetFilteringItem{
		{
			Field:    model.ExtendPackageGetFilterFieldChannelPackageID,
			Operator: model.ExtendPackageGetFilterOperatorEquals,
			Values:   []string{"val1", "val2"},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.values长度须为1")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageGetValidateFilterValueTooLongSelf 测试 filtering.values[0] 超过1024字节
func TestExtendPackageGetValidateFilterValueTooLongSelf(t *testing.T) {
	req := &model.ExtendPackageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.Filtering = []*model.ExtendPackageGetFilteringItem{
		{
			Field:    model.ExtendPackageGetFilterFieldChannelName,
			Operator: model.ExtendPackageGetFilterOperatorContains,
			Values:   []string{strings.Repeat("a", 1025)},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.values[0]超过1024字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageGetValidatePageOutOfRangeSelf 测试 page 超出范围
func TestExtendPackageGetValidatePageOutOfRangeSelf(t *testing.T) {
	req := &model.ExtendPackageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.Page = 100000
	req.PageSize = 10
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page超出范围")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageGetValidatePageSizeOutOfRangeSelf 测试 page_size 超出范围
func TestExtendPackageGetValidatePageSizeOutOfRangeSelf(t *testing.T) {
	req := &model.ExtendPackageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.Page = 1
	req.PageSize = 101
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size超出范围")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// TestExtendPackageGetValidateDefaultPaginationSelf 测试 Format() 默认填充分页参数
func TestExtendPackageGetValidateDefaultPaginationSelf(t *testing.T) {
	req := &model.ExtendPackageGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111
	req.PackageID = 2000000336
	req.Format()
	if req.Page != 1 {
		t.Fatalf("期望默认page=1，实际=%d", req.Page)
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
