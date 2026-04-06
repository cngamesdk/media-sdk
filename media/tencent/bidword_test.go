package tencent

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ========== 创建关键词测试用例 ==========

// TestBidwordAddBasicSelf 测试创建单个关键词（最简参数）
func TestBidwordAddBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordListItem{
		{
			AdgroupID: 456,
			Bidword:   "游戏推荐",
			MatchType: model.BidwordMatchTypeWide,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordAddWithBidPriceSelf 测试创建关键词（带出价）
func TestBidwordAddWithBidPriceSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordListItem{
		{
			AdgroupID: 456,
			Bidword:   "手机游戏",
			MatchType: model.BidwordMatchTypeExact,
			BidPrice:  200, // 2元
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordAddWithGroupPriceSelf 测试创建关键词（使用组出价）
func TestBidwordAddWithGroupPriceSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordListItem{
		{
			AdgroupID:     456,
			Bidword:       "网络游戏",
			MatchType:     model.BidwordMatchTypePhrase,
			UseGroupPrice: model.BidwordUseGroupPrice,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordAddWithSuspendStatusSelf 测试创建关键词（暂停状态）
func TestBidwordAddWithSuspendStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordListItem{
		{
			AdgroupID:        456,
			Bidword:          "策略游戏",
			MatchType:        model.BidwordMatchTypeWord,
			ConfiguredStatus: model.BidwordStatusSuspend,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordAddWithDynamicCreativeSelf 测试创建关键词（带创意id）
func TestBidwordAddWithDynamicCreativeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordListItem{
		{
			AdgroupID:         456,
			Bidword:           "角色扮演游戏",
			MatchType:         model.BidwordMatchTypeWide,
			DynamicCreativeID: 789,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordAddWithH5LandingPageSelf 测试创建关键词（H5落地页）
func TestBidwordAddWithH5LandingPageSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordListItem{
		{
			AdgroupID: 456,
			Bidword:   "休闲游戏",
			MatchType: model.BidwordMatchTypeWide,
			PcLandingPageInfo: &model.PcLandingPageInfo{
				PageType: model.PageTypeH5,
				PageSpec: &model.PageSpec{
					H5Spec: &model.H5Spec{
						PageURL: "https://www.example.com/game",
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordAddWithMiniProgramLandingPageSelf 测试创建关键词（微信小程序落地页）
func TestBidwordAddWithMiniProgramLandingPageSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordListItem{
		{
			AdgroupID: 456,
			Bidword:   "益智游戏",
			MatchType: model.BidwordMatchTypeExact,
			PcLandingPageInfo: &model.PcLandingPageInfo{
				PageType: model.PageTypeWechatMiniProgram,
				PageSpec: &model.PageSpec{
					WechatMiniProgramSpec: &model.WechatMiniProgramSpec{
						MiniProgramID:   "wx1234567890",
						MiniProgramPath: "pages/index/index",
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordAddWithLandingPageListSelf 测试创建关键词（含兜底落地页列表）
func TestBidwordAddWithLandingPageListSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordListItem{
		{
			AdgroupID: 456,
			Bidword:   "动作游戏",
			MatchType: model.BidwordMatchTypeWide,
			PcLandingPageInfo: &model.PcLandingPageInfo{
				PageType: model.PageTypeAppDeepLink,
				PageSpec: &model.PageSpec{
					AppDeepLinkSpec: &model.AppDeepLinkSpec{
						AndroidDeepLinkURL: "myapp://game",
						IosDeepLinkURL:     "myapp://game",
					},
				},
				LandingPageList: []*model.PcLandingPageItem{
					{
						PageType: model.PageTypeH5,
						PageSpec: &model.PageSpec{
							H5Spec: &model.H5Spec{
								PageURL: "https://www.example.com/fallback",
							},
						},
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordAddMultipleKeywordsSelf 测试批量创建多个关键词
func TestBidwordAddMultipleKeywordsSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordListItem{
		{
			AdgroupID: 456,
			Bidword:   "游戏下载",
			MatchType: model.BidwordMatchTypeWide,
			BidPrice:  100,
		},
		{
			AdgroupID: 456,
			Bidword:   "手游推荐",
			MatchType: model.BidwordMatchTypeExact,
			BidPrice:  150,
		},
		{
			AdgroupID:        456,
			Bidword:          "免费游戏",
			MatchType:        model.BidwordMatchTypePhrase,
			ConfiguredStatus: model.BidwordStatusNormal,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordAddAllMatchTypesSelf 测试四种匹配方式
func TestBidwordAddAllMatchTypesSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordListItem{
		{AdgroupID: 456, Bidword: "精确匹配游戏", MatchType: model.BidwordMatchTypeExact},
		{AdgroupID: 456, Bidword: "广泛匹配游戏", MatchType: model.BidwordMatchTypeWide},
		{AdgroupID: 456, Bidword: "词语匹配游戏", MatchType: model.BidwordMatchTypeWord},
		{AdgroupID: 456, Bidword: "短语匹配游戏", MatchType: model.BidwordMatchTypePhrase},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 参数校验测试用例 ==========

// TestBidwordAddValidateAccountIDSelf 测试缺少account_id时的校验
func TestBidwordAddValidateAccountIDSelf(t *testing.T) {
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.List = []*model.BidwordListItem{
		{AdgroupID: 456, Bidword: "游戏", MatchType: model.BidwordMatchTypeWide},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordAddValidateEmptyListSelf 测试list为空时的校验
func TestBidwordAddValidateEmptyListSelf(t *testing.T) {
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordListItem{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordAddValidateMissingBidwordSelf 测试缺少bidword时的校验
func TestBidwordAddValidateMissingBidwordSelf(t *testing.T) {
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordListItem{
		{AdgroupID: 456, MatchType: model.BidwordMatchTypeWide},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordAddValidateMissingMatchTypeSelf 测试缺少match_type时的校验
func TestBidwordAddValidateMissingMatchTypeSelf(t *testing.T) {
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordListItem{
		{AdgroupID: 456, Bidword: "游戏"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordAddValidateInvalidMatchTypeSelf 测试无效match_type时的校验
func TestBidwordAddValidateInvalidMatchTypeSelf(t *testing.T) {
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordListItem{
		{AdgroupID: 456, Bidword: "游戏", MatchType: "INVALID_MATCH"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordAddValidateBidPriceTooHighSelf 测试bid_price超过最大值时的校验
func TestBidwordAddValidateBidPriceTooHighSelf(t *testing.T) {
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordListItem{
		{AdgroupID: 456, Bidword: "游戏", MatchType: model.BidwordMatchTypeWide, BidPrice: 100000},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordAddValidateInvalidConfiguredStatusSelf 测试无效configured_status时的校验
func TestBidwordAddValidateInvalidConfiguredStatusSelf(t *testing.T) {
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordListItem{
		{AdgroupID: 456, Bidword: "游戏", MatchType: model.BidwordMatchTypeWide, ConfiguredStatus: "INVALID_STATUS"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordAddValidateLandingPageMissingPageTypeSelf 测试落地页缺少page_type时的校验
func TestBidwordAddValidateLandingPageMissingPageTypeSelf(t *testing.T) {
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordListItem{
		{
			AdgroupID: 456,
			Bidword:   "游戏",
			MatchType: model.BidwordMatchTypeWide,
			PcLandingPageInfo: &model.PcLandingPageInfo{
				PageSpec: &model.PageSpec{
					H5Spec: &model.H5Spec{PageURL: "https://example.com"},
				},
			},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordAddWithAndroidAppLandingPageSelf 测试创建关键词（Android应用落地页）
func TestBidwordAddWithAndroidAppLandingPageSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordAddReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordListItem{
		{
			AdgroupID: 456,
			Bidword:   "安卓游戏",
			MatchType: model.BidwordMatchTypeWide,
			PcLandingPageInfo: &model.PcLandingPageInfo{
				PageType: model.PageTypeAndroidApp,
				PageSpec: &model.PageSpec{
					AndroidAppSpec: &model.AndroidAppSpec{
						AndroidAppID: "com.example.game",
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 更新关键词测试用例 ==========

// TestBidwordUpdateBasicSelf 测试更新关键词出价（最简参数）
func TestBidwordUpdateBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{
			BidwordID: 2502973,
			BidPrice:  200,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordUpdateMatchTypeSelf 测试更新关键词匹配方式
func TestBidwordUpdateMatchTypeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{
			BidwordID: 2502973,
			MatchType: model.BidwordMatchTypeExact,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordUpdateConfiguredStatusSelf 测试更新关键词暂停状态
func TestBidwordUpdateConfiguredStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{
			BidwordID:        2502973,
			ConfiguredStatus: model.BidwordStatusSuspend,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordUpdateWithBidModeSelf 测试更新关键词出价方式
func TestBidwordUpdateWithBidModeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{
			BidwordID: 2502973,
			BidPrice:  300,
			BidMode:   model.BidModeCPC,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordUpdateWithOCPMBidModeSelf 测试更新关键词为OCPM出价方式
func TestBidwordUpdateWithOCPMBidModeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{
			BidwordID: 2502973,
			BidPrice:  500,
			BidMode:   model.BidModeOCPM,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordUpdateWithRaisePriceValueSelf 测试按数值修改出价
func TestBidwordUpdateWithRaisePriceValueSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{
			BidwordID:       2502973,
			PriceUpdateType: model.PriceUpdateTypeRaiseValue,
			RaisePrice:      50, // 出价上调0.5元
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordUpdateWithRaisePricePercentSelf 测试按百分比修改出价
func TestBidwordUpdateWithRaisePricePercentSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{
			BidwordID:       2502973,
			PriceUpdateType: model.PriceUpdateTypeRaisePercent,
			RaisePrice:      10, // 出价上调10%
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordUpdateWithNegativeRaisePriceSelf 测试按负值修改出价（降价）
func TestBidwordUpdateWithNegativeRaisePriceSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{
			BidwordID:       2502973,
			PriceUpdateType: model.PriceUpdateTypeRaiseValue,
			RaisePrice:      -30, // 出价下调0.3元
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordUpdateUseGroupPriceSelf 测试更新关键词为使用组出价
func TestBidwordUpdateUseGroupPriceSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{
			BidwordID:     2502973,
			UseGroupPrice: model.BidwordUseGroupPrice,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordUpdateWithDynamicCreativeSelf 测试更新关键词绑定创意
func TestBidwordUpdateWithDynamicCreativeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{
			BidwordID:         2502973,
			DynamicCreativeID: 789,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordUpdateWithH5LandingPageSelf 测试更新关键词落地页（H5）
func TestBidwordUpdateWithH5LandingPageSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{
			BidwordID: 2502973,
			PcLandingPageInfo: &model.PcLandingPageInfo{
				PageType: model.PageTypeH5,
				PageSpec: &model.PageSpec{
					H5Spec: &model.H5Spec{
						PageURL: "https://www.example.com/new-landing",
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordUpdateWithMiniProgramLandingPageSelf 测试更新关键词落地页（微信小程序）
func TestBidwordUpdateWithMiniProgramLandingPageSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{
			BidwordID: 2502973,
			PcLandingPageInfo: &model.PcLandingPageInfo{
				PageType: model.PageTypeWechatMiniProgram,
				PageSpec: &model.PageSpec{
					WechatMiniProgramSpec: &model.WechatMiniProgramSpec{
						MiniProgramID:   "wx1234567890",
						MiniProgramPath: "pages/game/index",
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordUpdateMultipleKeywordsSelf 测试批量更新多个关键词
func TestBidwordUpdateMultipleKeywordsSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{
			BidwordID: 2502973,
			BidPrice:  100,
			MatchType: model.BidwordMatchTypeExact,
		},
		{
			BidwordID:        2502974,
			ConfiguredStatus: model.BidwordStatusNormal,
		},
		{
			BidwordID:       2502975,
			PriceUpdateType: model.PriceUpdateTypeRaisePercent,
			RaisePrice:      20,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 更新关键词参数校验测试 ==========

// TestBidwordUpdateValidateAccountIDSelf 测试缺少account_id时的校验
func TestBidwordUpdateValidateAccountIDSelf(t *testing.T) {
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.List = []*model.BidwordUpdateListItem{
		{BidwordID: 2502973, BidPrice: 100},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordUpdateValidateEmptyListSelf 测试list为空时的校验
func TestBidwordUpdateValidateEmptyListSelf(t *testing.T) {
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordUpdateValidateMissingBidwordIDSelf 测试缺少bidword_id时的校验
func TestBidwordUpdateValidateMissingBidwordIDSelf(t *testing.T) {
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{BidPrice: 100},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordUpdateValidateInvalidBidModeSelf 测试无效bid_mode时的校验
func TestBidwordUpdateValidateInvalidBidModeSelf(t *testing.T) {
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{BidwordID: 2502973, BidMode: "INVALID_BID_MODE"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordUpdateValidateInvalidPriceUpdateTypeSelf 测试无效price_update_type时的校验
func TestBidwordUpdateValidateInvalidPriceUpdateTypeSelf(t *testing.T) {
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{BidwordID: 2502973, PriceUpdateType: "INVALID_TYPE"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordUpdateValidateRaisePriceOutOfRangeSelf 测试raise_price超出范围时的校验
func TestBidwordUpdateValidateRaisePriceOutOfRangeSelf(t *testing.T) {
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{BidwordID: 2502973, RaisePrice: 100000},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordUpdateValidateInvalidMatchTypeSelf 测试无效match_type时的校验
func TestBidwordUpdateValidateInvalidMatchTypeSelf(t *testing.T) {
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{BidwordID: 2502973, MatchType: "INVALID_MATCH"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordUpdateValidateLandingPageMissingPageTypeSelf 测试落地页缺少page_type时的校验
func TestBidwordUpdateValidateLandingPageMissingPageTypeSelf(t *testing.T) {
	req := &model.BidwordUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []*model.BidwordUpdateListItem{
		{
			BidwordID: 2502973,
			PcLandingPageInfo: &model.PcLandingPageInfo{
				PageSpec: &model.PageSpec{
					H5Spec: &model.H5Spec{PageURL: "https://example.com"},
				},
			},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// ========== 删除关键词测试用例 ==========

// TestBidwordDeleteSingleSelf 测试删除单个关键词
func TestBidwordDeleteSingleSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []int64{2502973}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordDeleteMultipleSelf 测试批量删除多个关键词
func TestBidwordDeleteMultipleSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []int64{2502973, 2502974, 2502975}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordDeleteMaxBatchSelf 测试删除接近上限数量的关键词（1000个）
func TestBidwordDeleteMaxBatchSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	ids := make([]int64, 1000)
	for i := range ids {
		ids[i] = int64(2500000 + i)
	}
	req.List = ids
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordDeleteSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 删除关键词参数校验测试 ==========

// TestBidwordDeleteValidateAccountIDSelf 测试缺少account_id时的校验
func TestBidwordDeleteValidateAccountIDSelf(t *testing.T) {
	req := &model.BidwordDeleteReq{}
	req.AccessToken = "123"
	req.List = []int64{2502973}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordDeleteValidateEmptyListSelf 测试list为空时的校验
func TestBidwordDeleteValidateEmptyListSelf(t *testing.T) {
	req := &model.BidwordDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []int64{}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordDeleteValidateNilListSelf 测试list为nil时的校验
func TestBidwordDeleteValidateNilListSelf(t *testing.T) {
	req := &model.BidwordDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordDeleteValidateExceedMaxListSelf 测试list超过最大长度时的校验
func TestBidwordDeleteValidateExceedMaxListSelf(t *testing.T) {
	req := &model.BidwordDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	ids := make([]int64, 1001)
	for i := range ids {
		ids[i] = int64(2500000 + i)
	}
	req.List = ids
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordDeleteValidateZeroIDSelf 测试list中包含无效id(0)时的校验
func TestBidwordDeleteValidateZeroIDSelf(t *testing.T) {
	req := &model.BidwordDeleteReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.List = []int64{2502973, 0, 2502975}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// ========== 查询关键词测试用例 ==========

// TestBidwordGetBasicSelf 测试查询关键词（无过滤条件）
func TestBidwordGetBasicSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordGetByAdgroupIDSelf 测试按广告id查询关键词
func TestBidwordGetByAdgroupIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.BidwordGetFilter{
		{
			Field:    model.BidwordFilterFieldAdgroupID,
			Operator: model.OperatorEquals,
			Values:   []string{"5076023598"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordGetByBidwordIDsSelf 测试按关键词id列表查询
func TestBidwordGetByBidwordIDsSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.BidwordGetFilter{
		{
			Field:    model.BidwordFilterFieldBidwordID,
			Operator: model.OperatorIn,
			Values:   []string{"2502973", "2502974", "2502975"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordGetByCampaignIDSelf 测试按推广计划id查询关键词
func TestBidwordGetByCampaignIDSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.BidwordGetFilter{
		{
			Field:    model.BidwordFilterFieldCampaignID,
			Operator: model.OperatorEquals,
			Values:   []string{"5076023595"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordGetByBidwordTextSelf 测试按关键词词面精确查询
func TestBidwordGetByBidwordTextSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.BidwordGetFilter{
		{
			Field:    model.BidwordFilterFieldBidword,
			Operator: model.OperatorEquals,
			Values:   []string{"游戏推荐"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordGetByBidwordContainsSelf 测试按关键词词面模糊查询
func TestBidwordGetByBidwordContainsSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.BidwordGetFilter{
		{
			Field:    model.BidwordFilterFieldBidword,
			Operator: model.OperatorContains,
			Values:   []string{"游戏"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordGetByMatchTypeSelf 测试按匹配方式查询关键词
func TestBidwordGetByMatchTypeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.BidwordGetFilter{
		{
			Field:    model.BidwordFilterFieldMatchType,
			Operator: model.OperatorEquals,
			Values:   []string{model.BidwordMatchTypeExact},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordGetByConfiguredStatusSelf 测试按暂停状态查询关键词
func TestBidwordGetByConfiguredStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.BidwordGetFilter{
		{
			Field:    model.BidwordFilterFieldConfiguredStatus,
			Operator: model.OperatorEquals,
			Values:   []string{model.BidwordStatusNormal},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordGetByBidwordStatusSelf 测试按关键词状态查询（审核通过）
func TestBidwordGetByBidwordStatusSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.BidwordGetFilter{
		{
			Field:    model.BidwordFilterFieldBidwordStatus,
			Operator: model.OperatorEquals,
			Values:   []string{model.BidwordStatusApprovalPassed},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordGetByCreatedTimeAfterSelf 测试按创建时间过滤（某时间之后）
func TestBidwordGetByCreatedTimeAfterSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.BidwordGetFilter{
		{
			Field:    model.BidwordFilterFieldCreatedTime,
			Operator: model.OperatorGreaterEquals,
			Values:   []string{"1711382400"}, // 2024-03-25 00:00:00
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordGetByLastModifiedTimeSelf 测试按最后修改时间过滤
func TestBidwordGetByLastModifiedTimeSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.BidwordGetFilter{
		{
			Field:    model.BidwordFilterFieldLastModifiedTime,
			Operator: model.OperatorGreater,
			Values:   []string{"1711382400"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordGetDeletedSelf 测试查询已删除的关键词
func TestBidwordGetDeletedSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.IsDeleted = true
	req.Filtering = []*model.BidwordGetFilter{
		{
			Field:    model.BidwordFilterFieldDeleteTime,
			Operator: model.OperatorGreaterEquals,
			Values:   []string{"1711382400"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordGetWithPaginationSelf 测试分页查询关键词
func TestBidwordGetWithPaginationSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Page = 2
	req.PageSize = 50
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordGetMultipleFiltersSelf 测试多过滤条件组合查询
func TestBidwordGetMultipleFiltersSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.BidwordGetFilter{
		{
			Field:    model.BidwordFilterFieldAdgroupID,
			Operator: model.OperatorEquals,
			Values:   []string{"5076023598"},
		},
		{
			Field:    model.BidwordFilterFieldMatchType,
			Operator: model.OperatorEquals,
			Values:   []string{model.BidwordMatchTypeWide},
		},
		{
			Field:    model.BidwordFilterFieldConfiguredStatus,
			Operator: model.OperatorEquals,
			Values:   []string{model.BidwordStatusNormal},
		},
	}
	req.PageSize = 100
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// TestBidwordGetByApprovalPendingSelf 测试查询审核中的关键词
func TestBidwordGetByApprovalPendingSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.BidwordGetFilter{
		{
			Field:    model.BidwordFilterFieldBidwordStatus,
			Operator: model.OperatorEquals,
			Values:   []string{model.BidwordStatusApprovalPending},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.BidwordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// ========== 查询关键词参数校验测试 ==========

// TestBidwordGetValidateAccountIDSelf 测试缺少account_id时的校验
func TestBidwordGetValidateAccountIDSelf(t *testing.T) {
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordGetValidateInvalidFieldSelf 测试无效过滤字段时的校验
func TestBidwordGetValidateInvalidFieldSelf(t *testing.T) {
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.BidwordGetFilter{
		{Field: "invalid_field", Operator: model.OperatorEquals, Values: []string{"1"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordGetValidateInvalidOperatorSelf 测试字段不支持的操作符时的校验
func TestBidwordGetValidateInvalidOperatorSelf(t *testing.T) {
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.BidwordGetFilter{
		// match_type 不支持 CONTAINS
		{Field: model.BidwordFilterFieldMatchType, Operator: model.OperatorContains, Values: []string{model.BidwordMatchTypeWide}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordGetValidateInvalidMatchTypeValueSelf 测试match_type值无效时的校验
func TestBidwordGetValidateInvalidMatchTypeValueSelf(t *testing.T) {
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.BidwordGetFilter{
		{Field: model.BidwordFilterFieldMatchType, Operator: model.OperatorEquals, Values: []string{"INVALID_MATCH"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordGetValidateInvalidBidwordStatusSelf 测试bidword_status值无效时的校验
func TestBidwordGetValidateInvalidBidwordStatusSelf(t *testing.T) {
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.BidwordGetFilter{
		{Field: model.BidwordFilterFieldBidwordStatus, Operator: model.OperatorEquals, Values: []string{"INVALID_STATUS"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordGetValidateTimeFieldLengthSelf 测试时间字段长度不为10时的校验
func TestBidwordGetValidateTimeFieldLengthSelf(t *testing.T) {
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Filtering = []*model.BidwordGetFilter{
		{Field: model.BidwordFilterFieldCreatedTime, Operator: model.OperatorEquals, Values: []string{"202403"}},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordGetValidatePageOutOfRangeSelf 测试page超出范围时的校验
func TestBidwordGetValidatePageOutOfRangeSelf(t *testing.T) {
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.Page = 100000 // 超出最大值 99999
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}

// TestBidwordGetValidatePageSizeOutOfRangeSelf 测试page_size超出范围时的校验
func TestBidwordGetValidatePageSizeOutOfRangeSelf(t *testing.T) {
	req := &model.BidwordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	req.PageSize = 1001 // 超出最大值 1000
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回校验错误，但未返回")
	}
	fmt.Printf("expected error: %v\n", err)
}
