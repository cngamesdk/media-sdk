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
