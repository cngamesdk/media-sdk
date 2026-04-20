package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取线索数据-基本查询
func TestVideoChannelLeadsDataGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.VideoChannelLeadsDataGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.StartDate = 20240101
	req.EndDate = 20240131
	req.Page = 1
	req.PageSize = 10
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.VideoChannelLeadsDataGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取线索数据-按品牌筛选
func TestVideoChannelLeadsDataGetByBrandSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.VideoChannelLeadsDataGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.BrandIds = []string{"brand_001", "brand_002"}
	req.BrandNames = []string{"品牌A", "品牌B"}
	req.StartDate = 20240101
	req.EndDate = 20240131
	req.Page = 1
	req.PageSize = 10
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.VideoChannelLeadsDataGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取线索数据-按经销商和视频号筛选
func TestVideoChannelLeadsDataGetByDealerAndChannelSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.VideoChannelLeadsDataGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DealerIds = []string{"dealer_001"}
	req.DealerNames = []string{"经销商A"}
	req.VideoChannelIds = []string{"vc_001"}
	req.VideoChannelNames = []string{"视频号A"}
	req.StartDate = 20240101
	req.EndDate = 20240131
	req.Page = 1
	req.PageSize = 20
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.VideoChannelLeadsDataGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestVideoChannelLeadsDataGetValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.VideoChannelLeadsDataGetReq{}
	req.AccessToken = "123"
	req.StartDate = 20240101
	req.EndDate = 20240131
	req.Page = 1
	req.PageSize = 10
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 start_date
func TestVideoChannelLeadsDataGetValidateStartDateEmptySelf(t *testing.T) {
	req := &model.VideoChannelLeadsDataGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.EndDate = 20240131
	req.Page = 1
	req.PageSize = 10
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：start_date为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 end_date
func TestVideoChannelLeadsDataGetValidateEndDateEmptySelf(t *testing.T) {
	req := &model.VideoChannelLeadsDataGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.StartDate = 20240101
	req.Page = 1
	req.PageSize = 10
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：end_date为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 page
func TestVideoChannelLeadsDataGetValidatePageEmptySelf(t *testing.T) {
	req := &model.VideoChannelLeadsDataGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.StartDate = 20240101
	req.EndDate = 20240131
	req.PageSize = 10
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page为必填且必须大于0")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 page_size
func TestVideoChannelLeadsDataGetValidatePageSizeEmptySelf(t *testing.T) {
	req := &model.VideoChannelLeadsDataGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.StartDate = 20240101
	req.EndDate = 20240131
	req.Page = 1
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_size为必填且必须大于0")
	}
	fmt.Printf("验证错误: %v\n", err)
}
