package baidu

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/baidu/model"
)

// TestGetProjectFeedSelf 测试查询项目（指定部分字段）
func TestGetProjectFeedSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.ProjectFeedReq{
		ProjectFeedFields: []string{
			"projectFeedId", "projectFeedName", "subject", "pause", "status",
			"bidMode", "ocpc", "projectOcpxStatus", "campaignFeedIds",
		},
	}
	resp, err := factory.GetProjectFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
	if len(resp.Data) > 0 {
		println(fmt.Sprintf("project data[0]: %+v", resp.Data[0]))
	}
}

// TestGetProjectFeedSelfByIDs 测试按项目ID查询
func TestGetProjectFeedSelfByIDs(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.ProjectFeedReq{
		ProjectFeedFields: []string{"projectFeedId", "projectFeedName", "status"},
		ProjectFeedIds:    []int64{123123},
	}
	resp, err := factory.GetProjectFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestGetProjectFeedSelfAllFields 测试查询项目（全部字段）
func TestGetProjectFeedSelfAllFields(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.ProjectFeedReq{
		ProjectFeedFields: []string{
			"projectFeedId", "projectFeedName", "subject", "appInfo",
			"pause", "status", "bidMode", "ocpc", "projectOcpxStatus",
			"bmcUserId", "catalogId", "catalogSource", "appSubType",
			"productType", "campaignFeedIds", "productIds", "miniProgramType",
			"useLiftBudget", "lift",
		},
	}
	resp, err := factory.GetProjectFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result count: %d", len(resp.Data)))
}

// TestAddProjectFeedSelf 测试新建项目
func TestAddProjectFeedSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.ProjectFeedAddReq{
		ProjectFeedTypes: []model.ProjectFeedType{
			{
				ProjectFeedName: "销售线索_项目测试",
				Subject:         model.SubjectSalesLeads,
				BidMode:         model.BidModeTargetCPA,
				Ocpc: model.OcpcModel{
					AppTransID: 5431211,
					TransFrom:  model.TransFromJimuPage,
					TransType:  model.TransTypeFormSubmit,
					OcpcBid:    123,
				},
				BmcUserID:       32111,
				CatalogID:       1111,
				ProductType:     model.ProductTypeNovel,
				CatalogSource:   model.CatalogSourceBMC,
				AppSubType:      model.AppSubTypeDownload,
				CampaignFeedIds: []int64{12341},
			},
		},
	}
	resp, err := factory.AddProjectFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
	if len(resp.Data) > 0 {
		println(fmt.Sprintf("created project: %+v", resp.Data[0]))
	}
}

// TestAddProjectFeedSelfWithLift 测试新建项目（带一键起量）
func TestAddProjectFeedSelfWithLift(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.ProjectFeedAddReq{
		ProjectFeedTypes: []model.ProjectFeedType{
			{
				ProjectFeedName: "一键起量项目测试",
				Subject:         model.SubjectSalesLeads,
				BidMode:         model.BidModeTargetCPA,
				Ocpc: model.OcpcModel{
					AppTransID: 5431211,
					TransFrom:  model.TransFromJimuPage,
					TransType:  model.TransTypeFormSubmit,
					OcpcBid:    100,
				},
				UseLiftBudget: model.UseLiftBudgetOn,
				Lift: &model.LiftBudgetSchedule{
					ScheduleModel: model.ScheduleModelWeekly,
					LiftBudget:    22,
					EventWeek:     "1,2,3",
					EventHour:     "00:00",
				},
			},
		},
	}
	resp, err := factory.AddProjectFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestAddProjectFeedSelfWithAiLift 测试新建项目（带智能起量）
func TestAddProjectFeedSelfWithAiLift(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.ProjectFeedAddReq{
		ProjectFeedTypes: []model.ProjectFeedType{
			{
				ProjectFeedName: "智能起量项目测试",
				Subject:         model.SubjectSalesLeads,
				BidMode:         model.BidModeTargetCPA,
				Ocpc: model.OcpcModel{
					AppTransID: 5431211,
					TransFrom:  model.TransFromJimuPage,
					TransType:  model.TransTypeFormSubmit,
					OcpcBid:    150,
				},
				AiLift:      model.AiLiftOn,
				AiLiftModel: model.AiLiftModelExplore,
			},
		},
	}
	resp, err := factory.AddProjectFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestUpdateProjectFeedSelf 测试更新项目名称
func TestUpdateProjectFeedSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.ProjectFeedUpdateReq{
		ProjectFeedTypes: []model.ProjectFeedType{
			{
				ProjectFeedID:   123121,
				ProjectFeedName: "销售线索_项目测试_更新",
			},
		},
	}
	resp, err := factory.UpdateProjectFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestUpdateProjectFeedSelfPause 测试暂停项目
func TestUpdateProjectFeedSelfPause(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	pauseTrue := true
	req := &model.ProjectFeedUpdateReq{
		ProjectFeedTypes: []model.ProjectFeedType{
			{
				ProjectFeedID: 123121,
				Pause:         &pauseTrue,
			},
		},
	}
	resp, err := factory.UpdateProjectFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestUpdateProjectFeedSelfLift 测试更新一键起量
func TestUpdateProjectFeedSelfLift(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.ProjectFeedUpdateReq{
		ProjectFeedTypes: []model.ProjectFeedType{
			{
				ProjectFeedID: 123121,
				UseLiftBudget: model.UseLiftBudgetOn,
				Lift: &model.LiftBudgetSchedule{
					ScheduleModel: model.ScheduleModelImmediate,
					LiftBudget:    50,
				},
			},
		},
	}
	resp, err := factory.UpdateProjectFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestUpdateProjectFeedSelfBatch 测试批量更新（计划+智能起量）
func TestUpdateProjectFeedSelfBatch(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.ProjectFeedUpdateReq{
		ProjectFeedTypes: []model.ProjectFeedType{
			{
				ProjectFeedID:   123121,
				CampaignFeedIds: []int64{12341, 12342},
				AiLift:          model.AiLiftOn,
				AiLiftModel:     model.AiLiftModelStable,
			},
		},
	}
	resp, err := factory.UpdateProjectFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}
