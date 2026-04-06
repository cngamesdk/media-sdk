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
