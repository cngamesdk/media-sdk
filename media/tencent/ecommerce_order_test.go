package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取订单数据
func TestEcommerceOrderGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.EcommerceOrderGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.DateRange = &model.EcommerceOrderDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-31",
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.EcommerceOrderGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取订单数据-带过滤条件
func TestEcommerceOrderGetWithFilteringSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.EcommerceOrderGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.DateRange = &model.EcommerceOrderDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-31",
	}
	req.Filtering = []*model.EcommerceOrderFilter{
		{
			Field:    model.EcommerceOrderFilterFieldOrderId,
			Operator: model.EcommerceOrderOperatorEquals,
			Values:   []string{"B503186974486037"},
		},
		{
			Field:    model.EcommerceOrderFilterFieldOrderStatus,
			Operator: model.EcommerceOrderOperatorIn,
			Values:   []string{model.EcommerceOrderStatusShipped, model.EcommerceOrderStatusDelivered},
		},
	}
	req.Page = 1
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.EcommerceOrderGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取订单数据-带日期参数
func TestEcommerceOrderGetWithDateSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.EcommerceOrderGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.Date = "2024-01-15"
	req.DateRange = &model.EcommerceOrderDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-31",
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.EcommerceOrderGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少account_id
func TestEcommerceOrderGetValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.EcommerceOrderGetReq{}
	req.AccessToken = "123"
	req.DateRange = &model.EcommerceOrderDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-31",
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少date_range
func TestEcommerceOrderGetValidateDateRangeNilSelf(t *testing.T) {
	req := &model.EcommerceOrderGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：date_range为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-date_range.start_date为空
func TestEcommerceOrderGetValidateStartDateEmptySelf(t *testing.T) {
	req := &model.EcommerceOrderGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.DateRange = &model.EcommerceOrderDateRange{
		EndDate: "2024-01-31",
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：date_range.start_date为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-filtering.field值无效
func TestEcommerceOrderGetValidateFilterFieldInvalidSelf(t *testing.T) {
	req := &model.EcommerceOrderGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.DateRange = &model.EcommerceOrderDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-31",
	}
	req.Filtering = []*model.EcommerceOrderFilter{
		{
			Field:    "invalid_field",
			Operator: model.EcommerceOrderOperatorEquals,
			Values:   []string{"test"},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering.field值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-filtering数组超出长度限制
func TestEcommerceOrderGetValidateFilteringCountExceedSelf(t *testing.T) {
	req := &model.EcommerceOrderGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.DateRange = &model.EcommerceOrderDateRange{
		StartDate: "2024-01-01",
		EndDate:   "2024-01-31",
	}
	req.Filtering = []*model.EcommerceOrderFilter{
		{Field: model.EcommerceOrderFilterFieldOrderId, Operator: model.EcommerceOrderOperatorEquals, Values: []string{"1"}},
		{Field: model.EcommerceOrderFilterFieldOrderStatus, Operator: model.EcommerceOrderOperatorIn, Values: []string{"SHIPPED"}},
		{Field: model.EcommerceOrderFilterFieldOrderId, Operator: model.EcommerceOrderOperatorEquals, Values: []string{"2"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：filtering数组长度必须在1-2之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}
