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

// ========== 更新订单状态测试 ==========

// 更新订单状态
func TestEcommerceOrderUpdateSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.EcommerceOrderUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.EcommerceOrderId = "B326518663301826"
	req.DeliveryTrackingNumber = "VB40977313484"
	req.ExpressCompany = model.ExpressCompanyZto
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.EcommerceOrderUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 更新订单状态-不带物流信息
func TestEcommerceOrderUpdateWithoutDeliverySelf(t *testing.T) {
	ctx := context.Background()
	req := &model.EcommerceOrderUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.EcommerceOrderId = "B326518663301826"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.EcommerceOrderUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-更新订单缺少account_id
func TestEcommerceOrderUpdateValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.EcommerceOrderUpdateReq{}
	req.AccessToken = "123"
	req.EcommerceOrderId = "B326518663301826"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少ecommerce_order_id
func TestEcommerceOrderUpdateValidateOrderIdEmptySelf(t *testing.T) {
	req := &model.EcommerceOrderUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：ecommerce_order_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-ecommerce_order_id超出长度
func TestEcommerceOrderUpdateValidateOrderIdTooLongSelf(t *testing.T) {
	req := &model.EcommerceOrderUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.EcommerceOrderId = "123456789012345678901" // 21 bytes
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：ecommerce_order_id长度必须在1-20字节之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-快递单号和快递公司不同时存在
func TestEcommerceOrderUpdateValidateDeliveryMismatchSelf(t *testing.T) {
	req := &model.EcommerceOrderUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.EcommerceOrderId = "B326518663301826"
	req.DeliveryTrackingNumber = "VB40977313484"
	// express_company未设置
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：delivery_tracking_number和express_company必须同时存在")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-快递公司枚举值无效
func TestEcommerceOrderUpdateValidateExpressCompanyInvalidSelf(t *testing.T) {
	req := &model.EcommerceOrderUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.EcommerceOrderId = "B326518663301826"
	req.DeliveryTrackingNumber = "VB40977313484"
	req.ExpressCompany = "INVALID_COMPANY"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：express_company值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}
