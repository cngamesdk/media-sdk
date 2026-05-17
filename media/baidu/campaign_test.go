package baidu

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/baidu/model"
)

// TestGetCampaignFeedSelf 测试查询计划（指定部分字段）
func TestGetCampaignFeedSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.CampaignFeedReq{
		CampaignFeedFields: []string{
			"campaignFeedId", "campaignFeedName", "subject", "budget",
			"starttime", "endtime", "pause", "status", "bstype", "campaignType",
		},
	}
	resp, err := factory.GetCampaignFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
	if len(resp.Data) > 0 {
		println(fmt.Sprintf("campaign data[0]: %+v", resp.Data[0]))
	}
}

// TestGetCampaignFeedSelfByIDs 测试按计划ID查询
func TestGetCampaignFeedSelfByIDs(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.CampaignFeedReq{
		CampaignFeedFields: []string{"campaignFeedId", "campaignFeedName", "status"},
		CampaignFeedIds:    []int64{12341},
	}
	resp, err := factory.GetCampaignFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestGetCampaignFeedSelfWithFilter 测试带过滤条件查询
func TestGetCampaignFeedSelfWithFilter(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.CampaignFeedReq{
		CampaignFeedFields: []string{"campaignFeedId", "campaignFeedName", "bstype", "status"},
		CampaignFeedFilter: &model.CampaignFeedFilter{
			BsType: []int{model.BsTypeNormal},
		},
	}
	resp, err := factory.GetCampaignFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestGetCampaignFeedSelfAllFields 测试查询全部字段
func TestGetCampaignFeedSelfAllFields(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.CampaignFeedReq{
		CampaignFeedFields: []string{
			"campaignFeedId", "campaignFeedName", "subject", "appinfo",
			"budget", "starttime", "endtime", "schedule", "pause", "status",
			"bstype", "campaignType", "addtime", "eshopType", "shadow",
			"budgetOfflineTime", "rtaStatus", "ftypes", "bidtype", "bid",
			"ocpc", "unefficientCampaign", "campaignOcpxStatus",
			"bmcUserId", "catalogId", "productType", "projectFeedId",
			"useLiftBudget", "liftBudget", "liftStatus",
			"deliveryType", "appSubType", "miniProgramType", "bidMode",
			"productIds", "saleType", "liftBudgetSchedule",
		},
	}
	resp, err := factory.GetCampaignFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result count: %d", len(resp.Data)))
}
