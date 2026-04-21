package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 创建客户人群-LOOKALIKE
func TestCustomAudiencesAddLookalike(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudiencesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Name = "测试LOOKALIKE人群"
	req.Type = "LOOKALIKE"
	req.Description = "LOOKALIKE人群描述"
	req.AudienceSpec = &model.AudienceSpec{
		LookalikeSpec: &model.LookalikeSpec{
			SeedAudienceID:  1234567,
			ExpandUserCount: 1000000,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudiencesAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 创建客户人群-USER_ACTION(URL匹配)
func TestCustomAudiencesAddUserActionURL(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudiencesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Name = "测试USER_ACTION_URL人群"
	req.Type = "USER_ACTION"
	req.AudienceSpec = &model.AudienceSpec{
		UserActionSpec: &model.UserActionSpec{
			UserActionSetID: 987654,
			MatchRuleType:   "URL",
			TimeWindow:      30,
			URLMatchRule: &model.URLMatchRule{
				URLMatcherGroup: []*model.URLMatcherGroup{
					{
						URLMatcher: []*model.URLMatcher{
							{ParamValue: "https://example.com", Operator: "CONTAIN"},
						},
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudiencesAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 创建客户人群-USER_ACTION(ACTION匹配)
func TestCustomAudiencesAddUserActionAction(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudiencesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Name = "测试USER_ACTION_ACTION人群"
	req.Type = "USER_ACTION"
	req.AudienceSpec = &model.AudienceSpec{
		UserActionSpec: &model.UserActionSpec{
			UserActionSetID: 987654,
			MatchRuleType:   "ACTION",
			TimeWindow:      7,
			ActionMatchRule: &model.ActionMatchRule{
				ActionType: "PURCHASE",
				ParamMatcherGroup: []*model.ParamMatcherGroup{
					{
						ParamMatcher: []*model.ParamMatcher{
							{ParamName: "price", ParamValue: "100", Operator: "GT"},
						},
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudiencesAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 创建客户人群-USER_ACTION(ACTION聚合)
func TestCustomAudiencesAddUserActionAggregation(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudiencesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Name = "测试USER_ACTION_AGGREGATION人群"
	req.Type = "USER_ACTION"
	req.AudienceSpec = &model.AudienceSpec{
		UserActionSpec: &model.UserActionSpec{
			UserActionSetID: 987654,
			MatchRuleType:   "ACTION",
			ExtractType:     "AGGREGATION",
			TimeWindow:      30,
			ActionAggregationRule: &model.ActionAggregationRule{
				ActionType: "PURCHASE",
				AggregationGroup: []*model.AggregationGroup{
					{
						AggregationMatcher: []*model.AggregationMatcher{
							{
								AggregationType: "COUNT",
								CountType:       "BY_TIMES",
								Comparator:      "COMPARATOR_GE",
								ComparisonValue: 3,
							},
						},
					},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudiencesAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 创建客户人群-KEYWORD
func TestCustomAudiencesAddKeyword(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudiencesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Name = "测试KEYWORD人群"
	req.Type = "KEYWORD"
	req.AudienceSpec = &model.AudienceSpec{
		KeywordSpec: &model.KeywordSpec{
			IncludeKeyword: []string{"汽车", "SUV"},
			ExcludeKeyword: []string{"二手"},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudiencesAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 创建客户人群-AD
func TestCustomAudiencesAddAd(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudiencesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Name = "测试AD人群"
	req.Type = "AD"
	req.AudienceSpec = &model.AudienceSpec{
		AdRuleSpec: &model.AdRuleSpec{
			RuleType:      "CLICK",
			StartDate:     "2024-01-01",
			EndDate:       "2024-01-31",
			AdgroupIDList: []int64{100001, 100002},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudiencesAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 创建客户人群-COMBINE
func TestCustomAudiencesAddCombine(t *testing.T) {
	ctx := context.Background()
	req := &model.CustomAudiencesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Name = "测试COMBINE人群"
	req.Type = "COMBINE"
	req.AudienceSpec = &model.AudienceSpec{
		CombineSpec: &model.CombineSpec{
			Include: [][]model.CombineAudienceItem{
				{
					{AudienceID: 111111, TimeWindow: 30},
					{AudienceID: 222222},
				},
			},
			Exclude: [][]model.CombineAudienceItem{
				{
					{AudienceID: 333333},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.CustomAudiencesAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestCustomAudiencesAddValidateAccountIDEmpty(t *testing.T) {
	req := &model.CustomAudiencesAddReq{}
	req.AccessToken = "123"
	req.Name = "测试人群"
	req.Type = "LOOKALIKE"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 name
func TestCustomAudiencesAddValidateNameEmpty(t *testing.T) {
	req := &model.CustomAudiencesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Type = "LOOKALIKE"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：name为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 type
func TestCustomAudiencesAddValidateTypeEmpty(t *testing.T) {
	req := &model.CustomAudiencesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Name = "测试人群"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-type 值无效
func TestCustomAudiencesAddValidateTypeInvalid(t *testing.T) {
	req := &model.CustomAudiencesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Name = "测试人群"
	req.Type = "INVALID_TYPE"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：type值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-name 超过最大长度
func TestCustomAudiencesAddValidateNameTooLong(t *testing.T) {
	req := &model.CustomAudiencesAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.Name = "这个名称超过了三十二个字节的限制测试一下到底行不行啊"
	req.Type = "LOOKALIKE"
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：name长度不能超过32字节")
	}
	fmt.Printf("验证错误: %v\n", err)
}
