package baidu

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/baidu/model"
)

// TestGetTransTraceSelf 测试查询转化追踪（全部接入方式）
func TestGetTransTraceSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.TransTraceReq{
		TransFrom: model.TransTraceAll,
	}
	resp, err := factory.GetTransTraceSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result count: %d", len(resp.Data)))
	if len(resp.Data) > 0 {
		data := resp.Data[0]
		println(fmt.Sprintf(
			"trans: appTransId=%d, transName=%s, transFrom=%d, transTypes=%v",
			data.AppTransId, data.TransName, data.TransFrom, data.TransTypes,
		))
	}
}

// TestGetTransTraceSelfByAppAPI 测试查询应用API接入方式的转化追踪
func TestGetTransTraceSelfByAppAPI(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.TransTraceReq{
		TransFrom: model.TransTraceAppAPI,
	}
	resp, err := factory.GetTransTraceSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result count: %d", len(resp.Data)))
	if len(resp.Data) > 0 {
		for _, data := range resp.Data {
			println(fmt.Sprintf(
				"trans: id=%d, name=%s, types=%v, appName=%s, apkName=%s, monitorUrl=%s",
				data.AppTransId, data.TransName, data.TransTypes,
				data.AppName, data.ApkName, data.MonitorUrl,
			))
		}
	}
}

// TestGetTransTraceSelfByJimu 测试查询基木鱼接入方式的转化追踪
func TestGetTransTraceSelfByJimu(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.TransTraceReq{
		TransFrom: model.TransTraceJimuPage,
	}
	resp, err := factory.GetTransTraceSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result count: %d", len(resp.Data)))
	if len(resp.Data) > 0 {
		for _, data := range resp.Data {
			println(fmt.Sprintf(
				"trans: id=%d, name=%s, lpUrl=%s, deepTransTypes=%v",
				data.AppTransId, data.TransName, data.LpUrl, data.DeepTransTypes,
			))
		}
	}
}

// TestGetTransTraceSelfByLeadsAPI 测试查询线索API接入方式的转化追踪
func TestGetTransTraceSelfByLeadsAPI(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.TransTraceReq{
		TransFrom: model.TransTraceLeadsAPI,
	}
	resp, err := factory.GetTransTraceSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result count: %d", len(resp.Data)))
	if len(resp.Data) > 0 {
		for _, data := range resp.Data {
			println(fmt.Sprintf(
				"trans: id=%d, name=%s, lpUrl=%s, relatedUrls=%v",
				data.AppTransId, data.TransName, data.LpUrl, data.RelatedUrls,
			))
		}
	}
}

// TestGetTransTraceSelfByAppSDK 测试查询应用SDK接入方式的转化追踪
func TestGetTransTraceSelfByAppSDK(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.TransTraceReq{
		TransFrom: model.TransTraceAppSDK,
	}
	resp, err := factory.GetTransTraceSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result count: %d", len(resp.Data)))
	if len(resp.Data) > 0 {
		for _, data := range resp.Data {
			println(fmt.Sprintf(
				"trans: id=%d, name=%s, appName=%s, apkName=%s, appType=%d, mode=%d",
				data.AppTransId, data.TransName, data.AppName, data.ApkName, data.AppType, data.Mode,
			))
		}
	}
}

// TestGetTransTraceSelfJimuFilter 测试带基木鱼过滤条件查询转化追踪
func TestGetTransTraceSelfJimuFilter(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.TransTraceReq{
		TransFrom: model.TransTraceJimuPage,
		JmyPageFilter: &model.JmyPageFilter{
			ShowType:    model.ShowTypeH5,
			PlatformIds: []int{model.PlatformJimu},
		},
	}
	resp, err := factory.GetTransTraceSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result count: %d", len(resp.Data)))
}
