package toutiao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/toutiao/model"
	"testing"
)

// 获取自定义报表可用指标和维度
func TestReportCustomConfigGetSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewToutiaoAdapter(config.DefaultConfig())
	req := &model.ReportCustomConfigGetReq{}
	req.AccessToken = "xxxx"
	req.AdvertiserID = 123
	req.DataTopics = []string{
		model.DataTopicMaterial,
	}
	resp, err := factory.ReportCustomConfigGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	result, _ := json.Marshal(resp)
	println(fmt.Sprintf("get result: %s", string(result)))
}

// 获取自定义报表可用指标和维度
func TestReportCustomGetGetSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewToutiaoAdapter(config.DefaultConfig())
	req := &model.ReportCustomGetReq{}
	req.AccessToken = "xxxx"
	req.AdvertiserID = 123
	req.Dimensions = []string{
		"stat_time_hour",
	}
	req.Metrics = []string{
		"stat_cost",
	}
	req.StartTime = "2026-03-18 00:00:00"
	req.EndTime = "2026-03-19 23:59:59"
	req.Filters = []*model.FilterCondition{
		&model.FilterCondition{
			Field:    "stat_cost",
			Type:     model.FieldTypeInput,
			Operator: model.OperatorGreaterThan,
			Values: []string{
				"0",
			},
		},
	}
	req.OrderBy = []model.OrderBy{
		{Field: "stat_cost", Type: model.OrderByDesc},
	}
	resp, err := factory.ReportCustomGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}
