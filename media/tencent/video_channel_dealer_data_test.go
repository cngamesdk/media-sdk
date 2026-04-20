package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取经销商聚合数据-基本查询
func TestVideoChannelDealerDataGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.VideoChannelDealerDataGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.StartDate = 20240101
	req.EndDate = 20240131
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.VideoChannelDealerDataGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取经销商聚合数据-按品牌筛选
func TestVideoChannelDealerDataGetByBrandSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.VideoChannelDealerDataGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.BrandIds = []string{"brand_001", "brand_002"}
	req.BrandNames = []string{"品牌A", "品牌B"}
	req.StartDate = 20240101
	req.EndDate = 20240131
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.VideoChannelDealerDataGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取经销商聚合数据-按经销商和视频号筛选
func TestVideoChannelDealerDataGetByDealerAndChannelSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.VideoChannelDealerDataGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.DealerIds = []string{"dealer_001"}
	req.DealerNames = []string{"经销商A"}
	req.VideoChannelIds = []string{"vc_001"}
	req.VideoChannelNames = []string{"视频号A"}
	req.StartDate = 20240101
	req.EndDate = 20240131
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.VideoChannelDealerDataGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestVideoChannelDealerDataGetValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.VideoChannelDealerDataGetReq{}
	req.AccessToken = "123"
	req.StartDate = 20240101
	req.EndDate = 20240131
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 start_date
func TestVideoChannelDealerDataGetValidateStartDateEmptySelf(t *testing.T) {
	req := &model.VideoChannelDealerDataGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.EndDate = 20240131
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：start_date为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 end_date
func TestVideoChannelDealerDataGetValidateEndDateEmptySelf(t *testing.T) {
	req := &model.VideoChannelDealerDataGetReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.StartDate = 20240101
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：end_date为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
