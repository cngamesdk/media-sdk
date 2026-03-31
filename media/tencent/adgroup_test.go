package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取广告
func TestAdgroupsGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 创建广告
func TestAdgroupsAddSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupName = "test"
	req.MarketingGoal = model.MarketingGoalUserGrowth
	req.MarketingCarrierType = model.MarketingCarrierTypeAppAndroid
	req.BeginDate = "2026-03-28"
	req.EndDate = "2026-03-28"
	req.BidAmount = 1
	req.TimeSeries = "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
	req.SiteSet = []string{model.SiteSetChannels}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 删除广告
func TestAdgroupsDeleteSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupsDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 123
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupsDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}

// 更新广告
func TestAdgroupsUpdateSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupsUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.AdgroupID = 123
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupsUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}
