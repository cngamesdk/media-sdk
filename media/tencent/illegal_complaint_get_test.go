package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取直客广告主违规申述列表-基本查询
func TestIllegalComplaintGet(t *testing.T) {
	ctx := context.Background()
	req := &model.IllegalComplaintGetReq{}
	req.AccessToken = "123"
	req.AccountIDList = []int64{2045867}
	req.Page = 1
	req.PageSize = 10
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.IllegalComplaintGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取直客广告主违规申述列表-多账户查询
func TestIllegalComplaintGetMultipleAccounts(t *testing.T) {
	ctx := context.Background()
	req := &model.IllegalComplaintGetReq{}
	req.AccessToken = "123"
	req.AccountIDList = []int64{2045867, 2045868, 2045869}
	req.Page = 1
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.IllegalComplaintGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取直客广告主违规申述列表-带筛选条件
func TestIllegalComplaintGetWithFilters(t *testing.T) {
	ctx := context.Background()
	req := &model.IllegalComplaintGetReq{}
	req.AccessToken = "123"
	req.AccountIDList = []int64{2045867}
	req.IllegalLevelList = []interface{}{1, 2}
	req.ActionTypeList = []interface{}{1}
	req.IllegalReason = "违规素材"
	req.IllegalDateRange = &model.IllegalComplaintDateRange{
		StartDate: "2026-01-01",
		EndDate:   "2026-04-22",
	}
	req.Page = 1
	req.PageSize = 10
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.IllegalComplaintGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id_list
func TestIllegalComplaintGetValidateAccountIDListEmpty(t *testing.T) {
	req := &model.IllegalComplaintGetReq{}
	req.AccessToken = "123"
	req.Page = 1
	req.PageSize = 10
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id_list为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-account_id_list 超过最大长度
func TestIllegalComplaintGetValidateAccountIDListTooLong(t *testing.T) {
	req := &model.IllegalComplaintGetReq{}
	req.AccessToken = "123"
	ids := make([]int64, 701)
	for i := range ids {
		ids[i] = int64(i + 1)
	}
	req.AccountIDList = ids
	req.Page = 1
	req.PageSize = 10
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id_list最大长度700")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-page_size 超出范围
func TestIllegalComplaintGetValidatePageSizeInvalid(t *testing.T) {
	req := &model.IllegalComplaintGetReq{}
	req.AccessToken = "123"
	req.AccountIDList = []int64{2045867}
	req.Page = 1
	req.PageSize = 101
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size必须在1-100之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-Format 默认值
func TestIllegalComplaintGetFormatDefaults(t *testing.T) {
	req := &model.IllegalComplaintGetReq{}
	req.AccessToken = "123"
	req.AccountIDList = []int64{2045867}
	req.Format()
	if req.Page != 1 {
		t.Fatalf("期望 page 默认值为 1，实际为 %d", req.Page)
	}
	if req.PageSize != 10 {
		t.Fatalf("期望 page_size 默认值为 10，实际为 %d", req.PageSize)
	}
	fmt.Printf("默认值验证通过: page=%d, page_size=%d\n", req.Page, req.PageSize)
}

// 验证测试-缺少 access_token
func TestIllegalComplaintGetValidateAccessTokenEmpty(t *testing.T) {
	req := &model.IllegalComplaintGetReq{}
	req.AccountIDList = []int64{2045867}
	req.Page = 1
	req.PageSize = 10
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：access_token为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
