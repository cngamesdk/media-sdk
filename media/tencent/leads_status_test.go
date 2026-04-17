package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 更新线索状态-通过外部线索id匹配
func TestLeadsStatusUpdateSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsStatusUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.LeadsConversionStatusList = []*model.LeadsConversionStatusItem{
		{
			OuterLeadsId:     "ext_001",
			LeadsConvertType: model.LeadsConvertStatusPotentialCustomer,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsStatusUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 更新线索状态-通过线索id匹配
func TestLeadsStatusUpdateByLeadsIdSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsStatusUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.MatchType = model.LeadsStatusMatchTypeLeadsId
	req.LeadsConversionStatusList = []*model.LeadsConversionStatusItem{
		{
			LeadsId:          123456,
			LeadsConvertType: model.LeadsConvertStatusHighIntentionCustomer,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsStatusUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 更新线索状态-通过联系方式匹配
func TestLeadsStatusUpdateByContactSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsStatusUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.MatchType = model.LeadsStatusMatchTypeContact
	req.LeadsConversionStatusList = []*model.LeadsConversionStatusItem{
		{
			LeadsTel:            "13800138000",
			LeadsConvertType:    model.LeadsConvertStatusDeprecated,
			LeadsIneffectReason: model.LeadsIneffectReasonTelNotConnected,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsStatusUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 更新线索状态-带自定义标签和外部状态
func TestLeadsStatusUpdateWithTagsSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsStatusUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.MatchType = model.LeadsStatusMatchTypeClickId
	req.LeadsConversionStatusList = []*model.LeadsConversionStatusItem{
		{
			ClickId:                  "click_123",
			LeadsConvertType:         model.LeadsConvertStatusTransCompleted,
			OuterLeadsConvertType:    "已签约",
			OuterLeadsIneffectReason: "",
			CustomizedTags: []*model.LeadsCustomizedTag{
				{
					TagGroupName: "意向等级",
					TagNameList:  []string{"高意向", "已付款"},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsStatusUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 更新线索状态-批量更新
func TestLeadsStatusUpdateBatchSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsStatusUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.LeadsConversionStatusList = []*model.LeadsConversionStatusItem{
		{
			OuterLeadsId:     "ext_001",
			LeadsConvertType: model.LeadsConvertStatusPotentialCustomer,
		},
		{
			OuterLeadsId:        "ext_002",
			LeadsConvertType:    model.LeadsConvertStatusDeprecated,
			LeadsIneffectReason: model.LeadsIneffectReasonDataDuplication,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsStatusUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少account_id
func TestLeadsStatusUpdateValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.LeadsStatusUpdateReq{}
	req.AccessToken = "123"
	req.LeadsConversionStatusList = []*model.LeadsConversionStatusItem{
		{OuterLeadsId: "ext_001", LeadsConvertType: model.LeadsConvertStatusPotentialCustomer},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少leads_conversion_status_list
func TestLeadsStatusUpdateValidateListEmptySelf(t *testing.T) {
	req := &model.LeadsStatusUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：leads_conversion_status_list为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-match_type值无效
func TestLeadsStatusUpdateValidateMatchTypeInvalidSelf(t *testing.T) {
	req := &model.LeadsStatusUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.MatchType = "INVALID"
	req.LeadsConversionStatusList = []*model.LeadsConversionStatusItem{
		{OuterLeadsId: "ext_001"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：match_type值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}
