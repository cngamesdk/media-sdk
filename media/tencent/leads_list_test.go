package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取线索列表
func TestLeadsListGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.TimeRange = &model.LeadsTimeRange{
		StartTime: 1704038400,
		EndTime:   1706716800,
		TimeType:  model.TimeTypeActionTime,
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取线索列表-按入库时间查询
func TestLeadsListGetByCreatedTimeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.TimeRange = &model.LeadsTimeRange{
		StartTime: 1704038400,
		EndTime:   1706716800,
		TimeType:  model.TimeTypeCreatedTime,
	}
	req.Page = 1
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取线索列表-深度翻页
func TestLeadsListGetWithSearchAfterSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.TimeRange = &model.LeadsTimeRange{
		StartTime: 1704038400,
		EndTime:   1706716800,
		TimeType:  model.TimeTypeActionTime,
	}
	req.PageSize = 100
	req.LastSearchAfterValues = []string{"1571367160000", "111111"}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsListGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少account_id
func TestLeadsListGetValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.LeadsListGetReq{}
	req.AccessToken = "123"
	req.TimeRange = &model.LeadsTimeRange{
		StartTime: 1704038400,
		EndTime:   1706716800,
		TimeType:  model.TimeTypeActionTime,
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少time_range
func TestLeadsListGetValidateTimeRangeNilSelf(t *testing.T) {
	req := &model.LeadsListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：time_range为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-time_type值无效
func TestLeadsListGetValidateTimeTypeInvalidSelf(t *testing.T) {
	req := &model.LeadsListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.TimeRange = &model.LeadsTimeRange{
		StartTime: 1704038400,
		EndTime:   1706716800,
		TimeType:  "INVALID_TYPE",
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：time_range.time_type值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-start_time大于等于end_time
func TestLeadsListGetValidateTimeRangeInvalidSelf(t *testing.T) {
	req := &model.LeadsListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.TimeRange = &model.LeadsTimeRange{
		StartTime: 1706716800,
		EndTime:   1704038400,
		TimeType:  model.TimeTypeActionTime,
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：time_range.start_time必须小于end_time")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-last_search_after_values长度不为2
func TestLeadsListGetValidateSearchAfterValuesInvalidSelf(t *testing.T) {
	req := &model.LeadsListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.TimeRange = &model.LeadsTimeRange{
		StartTime: 1704038400,
		EndTime:   1706716800,
		TimeType:  model.TimeTypeActionTime,
	}
	req.LastSearchAfterValues = []string{"1571367160000"}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：last_search_after_values长度必须为2")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-page超出范围
func TestLeadsListGetValidatePageExceedSelf(t *testing.T) {
	req := &model.LeadsListGetReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.TimeRange = &model.LeadsTimeRange{
		StartTime: 1704038400,
		EndTime:   1706716800,
		TimeType:  model.TimeTypeActionTime,
	}
	req.Page = 1001
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page必须在1-1000之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}
