package baidu

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/baidu/model"
)

// TestGetAdgroupFeedSelf 测试查询单元（指定部分字段）
func TestGetAdgroupFeedSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AdgroupFeedReq{
		AdgroupFeedFields: []string{
			"adgroupFeedId", "campaignFeedId", "adgroupFeedName", "pause", "status",
			"bid", "ftypes", "bidtype", "ocpc", "atpFeedId",
		},
		Ids:    []int64{1},
		IdType: model.IdTypeUnit,
	}
	resp, err := factory.GetAdgroupFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
	if len(resp.Data) > 0 {
		println(fmt.Sprintf("unit data[0]: %+v", resp.Data[0]))
	}
}

// TestGetAdgroupFeedSelfByCampaignId 测试按计划ID查询单元
func TestGetAdgroupFeedSelfByCampaignId(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AdgroupFeedReq{
		AdgroupFeedFields: []string{
			"adgroupFeedId", "campaignFeedId", "adgroupFeedName", "status",
		},
		Ids:    []int64{12387113},
		IdType: model.IdTypeCampaign,
	}
	resp, err := factory.GetAdgroupFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestGetAdgroupFeedSelfByUnitIds 测试按单元ID查询
func TestGetAdgroupFeedSelfByUnitIds(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AdgroupFeedReq{
		AdgroupFeedFields: []string{"adgroupFeedId", "adgroupFeedName", "status", "bid"},
		Ids:               []int64{1, 2},
		IdType:            model.IdTypeUnit,
	}
	resp, err := factory.GetAdgroupFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestGetAdgroupFeedSelfAllFields 测试查询全部字段（完整字段）
func TestGetAdgroupFeedSelfAllFields(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AdgroupFeedReq{
		AdgroupFeedFields: []string{
			"adgroupFeedId", "campaignFeedId", "adgroupFeedName", "pause", "status",
			"audience", "bid", "ftypes", "bidtype", "ocpc",
			"atpFeedId", "addtime", "modtime", "deliveryType",
			"unefficientAdgroup", "productSetId", "unitProducts",
			"ftypeSelection", "bidSource", "urlType",
			"miniProgram", "broadCastInfo", "url",
			"unitOcpxStatus", "atpName",
		},
		Ids:    []int64{1},
		IdType: model.IdTypeUnit,
	}
	resp, err := factory.GetAdgroupFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result count: %d", len(resp.Data)))
	if len(resp.Data) > 0 {
		data := resp.Data[0]
		println(fmt.Sprintf(
			"unit: id=%d, name=%s, campaignId=%d, pause=%v, status=%d, bid=%.2f, bidtype=%d",
			data.AdgroupFeedId, data.AdgroupFeedName, data.CampaignFeedId,
			data.Pause, data.Status, data.Bid, data.Bidtype,
		))
		if data.Ocpc != nil {
			println(fmt.Sprintf(
				"ocpc: appTransId=%d, transFrom=%d, ocpcBid=%.2f, transType=%d, urlType=%d, useRoi=%v, roiRatio=%.2f",
				data.Ocpc.AppTransID, data.Ocpc.TransFrom, data.Ocpc.OcpcBid,
				data.Ocpc.TransType, data.Ocpc.UrlType, data.Ocpc.UseRoi, data.Ocpc.RoiRatio,
			))
		}
		if data.UnitProducts != nil {
			println(fmt.Sprintf("unitProducts: catalogId=%d, ruleProducts count=%d",
				data.UnitProducts.CatalogID, len(data.UnitProducts.RuleProducts)))
		}
		println(fmt.Sprintf(
			"ftypes=%v, deliveryType=%v, unefficientAdgroup=%d, productSetId=%d",
			data.Ftypes, data.DeliveryType, data.UnefficientAdgroup, data.ProductSetId,
		))
		println(fmt.Sprintf(
			"ftypeSelection=%d, bidSource=%d, urlType=%d, unitOcpxStatus=%d, atpName=%s",
			data.FtypeSelection, data.BidSource, data.UrlType, data.UnitOcpxStatus, data.AtpName,
		))
		println(fmt.Sprintf("addtime=%s, modtime=%s", data.AddTime, data.ModTime))
		println(fmt.Sprintf("miniProgram=%s, broadCastInfo=%s, url=%s", data.MiniProgram, data.BroadCastInfo, data.Url))
	}
}

// TestGetAdgroupFeedSelfWithAudience 测试查询包含定向设置的单元
func TestGetAdgroupFeedSelfWithAudience(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AdgroupFeedReq{
		AdgroupFeedFields: []string{
			"adgroupFeedId", "adgroupFeedName", "audience",
		},
		Ids:    []int64{1},
		IdType: model.IdTypeUnit,
	}
	resp, err := factory.GetAdgroupFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
	if len(resp.Data) > 0 {
		println(fmt.Sprintf("audience: %+v", resp.Data[0].Audience))
	}
}
