package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取处罚指标数据-基本查询
func TestPunishMetricsGet(t *testing.T) {
	ctx := context.Background()
	req := &model.PunishMetricsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.PunishMetricsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取处罚指标数据-指定统计周期
func TestPunishMetricsGetWithPartitionTime(t *testing.T) {
	ctx := context.Background()
	req := &model.PunishMetricsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.PartitionTime = 202604
	req.PageNum = 1
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.PunishMetricsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取处罚指标数据-指定时间范围和筛选条件
func TestPunishMetricsGetWithFilters(t *testing.T) {
	ctx := context.Background()
	req := &model.PunishMetricsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.StartPartitionTime = 202601
	req.EndPartitionTime = 202604
	req.OpsAdvertiserNameList = []string{"测试广告主"}
	req.FirstLevelIndustryName = "游戏"
	req.SecondLevelIndustryName = "手机游戏"
	req.PageNum = 1
	req.PageSize = 10
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.PunishMetricsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取处罚指标数据-代理商维度查询
func TestPunishMetricsGetWithAgentGroup(t *testing.T) {
	ctx := context.Background()
	req := &model.PunishMetricsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.ZcAgentGroup = "测试政策集团"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.PunishMetricsGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestPunishMetricsGetValidateAccountIDEmpty(t *testing.T) {
	req := &model.PunishMetricsGetReq{}
	req.AccessToken = "123"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-page_size 超出范围
func TestPunishMetricsGetValidatePageSizeInvalid(t *testing.T) {
	req := &model.PunishMetricsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.PageNum = 1
	req.PageSize = 1001
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size必须在1-1000之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-Format 默认值
func TestPunishMetricsGetFormatDefaults(t *testing.T) {
	req := &model.PunishMetricsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Format()
	if req.PageNum != 1 {
		t.Fatalf("期望 page_num 默认值为 1，实际为 %d", req.PageNum)
	}
	if req.PageSize != 20 {
		t.Fatalf("期望 page_size 默认值为 20，实际为 %d", req.PageSize)
	}
	fmt.Printf("默认值验证通过: page_num=%d, page_size=%d\n", req.PageNum, req.PageSize)
}

// 验证测试-缺少 access_token
func TestPunishMetricsGetValidateAccessTokenEmpty(t *testing.T) {
	req := &model.PunishMetricsGetReq{}
	req.AccountID = 2045867
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：access_token为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
