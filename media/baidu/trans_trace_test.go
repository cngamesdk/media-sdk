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

// TestAddTransTraceSelfAppAPI 测试新增应用API转化追踪（iOS）
func TestAddTransTraceSelfAppAPI(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.TransTraceAddReq{
		OcpcTransFeedTypes: []model.OcpcTransFeedType{
			{
				TransFrom:   model.TransTraceAppAPI,
				TransName:   "测试应用API转化_iOS",
				TransTypes:  []int{model.AddTransTypeActivate},
				Mode:        model.MonitorModeClickOnly,
				MonitorUrl:  "https://track.example.com/callback?callback={{CALLBACK_URL}}",
				AppType:     model.AppTypeIOS,
				AppName:     "测试iOS应用",
				DownloadUrl: "https://apps.apple.com/test",
			},
		},
	}
	resp, err := factory.AddTransTraceSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("add result: %+v", resp))
	if len(resp.Data) > 0 {
		println(fmt.Sprintf("created trans: id=%d, name=%s, status=%d",
			resp.Data[0].AppTransId, resp.Data[0].TransName, resp.Data[0].TransStatus))
	}
}

// TestAddTransTraceSelfAppAPIAndroid 测试新增应用API转化追踪（Android）
func TestAddTransTraceSelfAppAPIAndroid(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.TransTraceAddReq{
		OcpcTransFeedTypes: []model.OcpcTransFeedType{
			{
				TransFrom:  model.TransTraceAppAPI,
				TransName:  "测试应用API转化_Android",
				TransTypes: []int{model.AddTransTypeActivate},
				Mode:       model.MonitorModeClickOnly,
				MonitorUrl: "https://track.example.com/callback?callback={{CALLBACK_URL}}",
				AppType:    model.AppTypeAndroid,
				ApkName:    "com.example.testapp",
			},
		},
	}
	resp, err := factory.AddTransTraceSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("add result: %+v", resp))
}

// TestAddTransTraceSelfAppSDK 测试新增应用SDK转化追踪
func TestAddTransTraceSelfAppSDK(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.TransTraceAddReq{
		OcpcTransFeedTypes: []model.OcpcTransFeedType{
			{
				TransFrom:    model.TransTraceAppSDK,
				TransName:    "测试应用SDK转化",
				TransTypes:   []int{model.AddTransTypeActivate},
				AppType:      model.AppTypeAndroid,
				AppName:      "测试SDK应用",
				ApkName:      "com.example.sdkapp",
				SdkAppId:     12345,
				SdkSecretKey: "test_secret_key",
			},
		},
	}
	resp, err := factory.AddTransTraceSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("add result: %+v", resp))
}

// TestAddTransTraceSelfFull 测试新增转化追踪（完整字段含深度转化）
func TestAddTransTraceSelfFull(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.TransTraceAddReq{
		OcpcTransFeedTypes: []model.OcpcTransFeedType{
			{
				TransFrom:      model.TransTraceAppAPI,
				TransName:      "完整转化追踪_含深度转化",
				TransTypes:     []int{model.AddTransTypeActivate},
				DeepTransTypes: []int{model.TransTypePurchaseSuccess, model.TransTypeRegister},
				Mode:           model.MonitorModeClickExposure,
				MonitorUrl:     "https://track.example.com/click?callback={{CALLBACK_URL}}",
				ExposureUrl:    "https://track.example.com/exposure?callback={{CALLBACK_URL}}",
				AppType:        model.AppTypeAndroid,
				AppName:        "完整测试应用",
				ApkName:        "com.example.fullapp",
			},
		},
	}
	resp, err := factory.AddTransTraceSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("add result: %+v", resp))
}
