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

// TestAddCampaignFeedSelf 测试新建计划
func TestAddCampaignFeedSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.CampaignFeedAddReq{
		CampaignFeedTypes: []model.CampaignFeedType{
			{
				CampaignFeedName: "销售线索_计划测试",
				Subject:          model.SubjectSalesLeads,
				BidType:          model.BidTypeOCPC,
				Bid:              1.5,
				Ocpc: &model.OcpcModel{
					AppTransID: 5431211,
					TransFrom:  model.TransFromJimuPage,
					TransType:  model.TransTypeFormSubmit,
					OcpcBid:    100,
				},
			},
		},
	}
	resp, err := factory.AddCampaignFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestAddCampaignFeedSelfFull 测试新建计划（完整字段）
func TestAddCampaignFeedSelfFull(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.CampaignFeedAddReq{
		CampaignFeedTypes: []model.CampaignFeedType{
			{
				CampaignFeedName: "应用下载_计划_完整",
				Subject:          model.SubjectAppDownloadAndroid,
				AppInfo: &model.AppInfoType{
					AppName:   "测试APP",
					ApkName:   "com.test.app",
					ChannelID: 12345,
				},
				Budget:    5000,
				StartTime: "2026-06-01",
				EndTime:   "2026-12-31",
				BidType:   model.BidTypeOCPC,
				Bid:       2.0,
				Ocpc: &model.OcpcModel{
					AppTransID: 5431211,
					TransFrom:  model.TransFromAppSDK,
					TransType:  model.TransTypeActivate,
					OcpcBid:    50,
				},
				BsType:        model.BsTypeNormal,
				CampaignType:  model.CampaignTypeNormal,
				ProjectFeedID: 123123,
				AppSubType:    model.AppSubTypeDownload,
			},
		},
	}
	resp, err := factory.AddCampaignFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestAddCampaignFeedSelfWithLift 测试新建计划（带一键起量）
func TestAddCampaignFeedSelfWithLift(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.CampaignFeedAddReq{
		CampaignFeedTypes: []model.CampaignFeedType{
			{
				CampaignFeedName: "一键起量计划",
				Subject:          model.SubjectSalesLeads,
				BidType:          model.BidTypeOCPC,
				Bid:              1.5,
				Ocpc: &model.OcpcModel{
					AppTransID: 5431211,
					TransFrom:  model.TransFromJimuPage,
					TransType:  model.TransTypeFormSubmit,
					OcpcBid:    100,
				},
				UseLiftBudget: model.UseLiftBudgetOn,
				LiftBudget:    200,
			},
		},
	}
	resp, err := factory.AddCampaignFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestUpdateCampaignFeedSelf 测试更新计划
func TestUpdateCampaignFeedSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.CampaignFeedUpdateReq{
		CampaignFeedTypes: []model.CampaignFeedType{
			{
				CampaignFeedID:   12818191,
				CampaignFeedName: "更新计划名称",
				Budget:           3000,
			},
		},
	}
	resp, err := factory.UpdateCampaignFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestUpdateCampaignFeedSelfPause 测试暂停/启用计划
func TestUpdateCampaignFeedSelfPause(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	pauseTrue := true
	req := &model.CampaignFeedUpdateReq{
		CampaignFeedTypes: []model.CampaignFeedType{
			{
				CampaignFeedID: 12818191,
				Pause:          &pauseTrue,
			},
		},
	}
	resp, err := factory.UpdateCampaignFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestUpdateCampaignFeedSelfBid 测试更新出价
func TestUpdateCampaignFeedSelfBid(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.CampaignFeedUpdateReq{
		CampaignFeedTypes: []model.CampaignFeedType{
			{
				CampaignFeedID: 12818191,
				Bid:            2.5,
				Ocpc: &model.OcpcModel{
					OcpcBid:   80,
					TransType: model.TransTypeFormSubmit,
				},
			},
		},
	}
	resp, err := factory.UpdateCampaignFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}
