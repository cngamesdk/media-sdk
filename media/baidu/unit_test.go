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

// TestAddAdgroupFeedSelf 测试新建单元（基础字段）
func TestAddAdgroupFeedSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	pauseFalse := false
	req := &model.AdgroupFeedAddReq{
		AdgroupFeedTypes: []model.AdgroupFeedType{
			{
				CampaignFeedId:  12387113,
				AdgroupFeedName: "信息流推广单元_测试",
				Pause:           &pauseFalse,
				Bid:             100.0,
				Ftypes:          []int{},
				Bidtype:         model.BidTypeCPC,
			},
		},
	}
	resp, err := factory.AddAdgroupFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("add result: %+v", resp))
	if len(resp.Data) > 0 {
		println(fmt.Sprintf("created unit: id=%d, name=%s", resp.Data[0].AdgroupFeedId, resp.Data[0].AdgroupFeedName))
	}
}

// TestAddAdgroupFeedSelfWithOcpc 测试新建单元（oCPC出价）
func TestAddAdgroupFeedSelfWithOcpc(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AdgroupFeedAddReq{
		AdgroupFeedTypes: []model.AdgroupFeedType{
			{
				CampaignFeedId:  12387113,
				AdgroupFeedName: "信息流推广单元_oCPC",
				Bid:             1.5,
				Ftypes:          []int{model.FtypeBaiduFeed},
				Bidtype:         model.BidTypeOCPC,
				Ocpc: &model.AdgroupFeedOcpcType{
					AppTransID: 23415,
					TransFrom:  model.TransFromLeadsAPI,
					OcpcBid:    200.0,
					LpUrl:      "http://www.baidu.com",
					TransType:  model.TransTypeLeaveLeads,
				},
			},
		},
	}
	resp, err := factory.AddAdgroupFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("add result: %+v", resp))
	if len(resp.Data) > 0 && resp.Data[0].Ocpc != nil {
		println(fmt.Sprintf("ocpc: transType=%d, ocpcBid=%.2f", resp.Data[0].Ocpc.TransType, resp.Data[0].Ocpc.OcpcBid))
	}
}

// TestAddAdgroupFeedSelfFull 测试新建单元（完整字段）
func TestAddAdgroupFeedSelfFull(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	pauseFalse := false
	req := &model.AdgroupFeedAddReq{
		AdgroupFeedTypes: []model.AdgroupFeedType{
			{
				CampaignFeedId:  12387113,
				AdgroupFeedName: "信息流推广单元_完整测试",
				Pause:           &pauseFalse,
				Audience:        map[string]string{},
				Bid:             100.0,
				Ftypes:          []int{model.FtypeBaiduFeed},
				Bidtype:         model.BidTypeOCPC,
				Ocpc: &model.AdgroupFeedOcpcType{
					AppTransID:        23415,
					TransFrom:         model.TransFromLeadsAPI,
					OcpcBid:           200.0,
					LpUrl:             "http://www.baidu.com",
					TransType:         model.TransTypeLeaveLeads,
					OptimizeDeepTrans: false,
					DeepOcpcBid:       0.0,
					DeepTransType:     model.TransTypePurchaseSuccess,
					UrlType:           model.UrlTypeNormal,
					UseRoi:            false,
					RoiRatio:          0.0,
					MiniProgramType:   model.MiniProgramTypeMini,
					AppKey:            "32",
					PagePath:          "example/page",
					BroadCastMode:     model.BroadCastModeContinuous,
					AnchorId:          1,
				},
				DeliveryType:   []int{model.DeliveryTypeAll},
				ProductSetId:   12345678,
				FtypeSelection: model.FtypeSelectionUnit,
				BidSource:      model.BidSourceUnit,
				UrlType:        model.UrlTypeNormal,
				Url:            "http://www.baidu.com",
			},
		},
	}
	resp, err := factory.AddAdgroupFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("add result: %+v", resp))
}

// TestAddAdgroupFeedSelfWithProducts 测试新建单元（商品推广）
func TestAddAdgroupFeedSelfWithProducts(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AdgroupFeedAddReq{
		AdgroupFeedTypes: []model.AdgroupFeedType{
			{
				CampaignFeedId:  12387113,
				AdgroupFeedName: "商品推广单元_测试",
				Bid:             100.0,
				Ftypes:          []int{},
				Bidtype:         model.BidTypeCPC,
				UnitProducts: &model.UnitProducts{
					CatalogID: 1,
					RuleProducts: []model.ProductSetRule{
						{
							Field:     "name",
							Operation: model.OperationEqual,
							Value:     "衬衫",
						},
					},
				},
			},
		},
	}
	resp, err := factory.AddAdgroupFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("add result: %+v", resp))
	if len(resp.Data) > 0 && resp.Data[0].UnitProducts != nil {
		println(fmt.Sprintf("unitProducts: catalogId=%d, rules=%+v",
			resp.Data[0].UnitProducts.CatalogID, resp.Data[0].UnitProducts.RuleProducts))
	}
}

// TestAddAdgroupFeedSelfWithBidSource 测试新建单元（使用计划出价）
func TestAddAdgroupFeedSelfWithBidSource(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AdgroupFeedAddReq{
		AdgroupFeedTypes: []model.AdgroupFeedType{
			{
				CampaignFeedId:  12387113,
				AdgroupFeedName: "使用计划出价_单元",
				Bid:             0,
				Ftypes:          []int{},
				BidSource:       model.BidSourcePlan,
				UrlType:         model.UrlTypeNormal,
				Url:             "http://www.baidu.com",
			},
		},
	}
	resp, err := factory.AddAdgroupFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("add result: %+v", resp))
}

// TestUpdateAdgroupFeedSelf 测试更新单元名称
func TestUpdateAdgroupFeedSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AdgroupFeedUpdateReq{
		AdgroupFeedTypes: []model.AdgroupFeedUpdateType{
			{
				AdgroupFeedId:   12387113,
				AdgroupFeedName: "更新后的单元名称",
			},
		},
	}
	resp, err := factory.UpdateAdgroupFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("update result: %+v", resp))
}

// TestUpdateAdgroupFeedSelfPause 测试暂停/启用单元
func TestUpdateAdgroupFeedSelfPause(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	pauseTrue := true
	req := &model.AdgroupFeedUpdateReq{
		AdgroupFeedTypes: []model.AdgroupFeedUpdateType{
			{
				AdgroupFeedId: 12387113,
				Pause:         &pauseTrue,
			},
		},
	}
	resp, err := factory.UpdateAdgroupFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("update result: %+v", resp))
}

// TestUpdateAdgroupFeedSelfBid 测试更新出价
func TestUpdateAdgroupFeedSelfBid(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AdgroupFeedUpdateReq{
		AdgroupFeedTypes: []model.AdgroupFeedUpdateType{
			{
				AdgroupFeedId: 12387113,
				Bid:           2.5,
			},
		},
	}
	resp, err := factory.UpdateAdgroupFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("update result: %+v", resp))
}

// TestUpdateAdgroupFeedSelfOcpc 测试更新oCPC设置
func TestUpdateAdgroupFeedSelfOcpc(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AdgroupFeedUpdateReq{
		AdgroupFeedTypes: []model.AdgroupFeedUpdateType{
			{
				AdgroupFeedId: 12387113,
				Ocpc: &model.AdgroupFeedOcpcType{
					AppTransID: 23415,
					TransFrom:  model.TransFromLeadsAPI,
					OcpcBid:    180.0,
					LpUrl:      "http://www.baidu.com",
					TransType:  model.TransTypeLeaveLeads,
				},
			},
		},
	}
	resp, err := factory.UpdateAdgroupFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("update result: %+v", resp))
}

// TestUpdateAdgroupFeedSelfUrl 测试更新落地页
func TestUpdateAdgroupFeedSelfUrl(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AdgroupFeedUpdateReq{
		AdgroupFeedTypes: []model.AdgroupFeedUpdateType{
			{
				AdgroupFeedId: 12387113,
				Url:           "http://www.newlandingpage.com",
			},
		},
	}
	resp, err := factory.UpdateAdgroupFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("update result: %+v", resp))
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
