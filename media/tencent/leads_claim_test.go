package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 更新线索归因信息-基础用法
func TestLeadsClaimUpdateSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsClaimUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.LeadsClaimList = []*model.LeadsClaimItem{
		{
			OuterLeadsId: "ext_001",
			CampaignId:   100001,
			AdgroupId:    200001,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsClaimUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 更新线索归因信息-带用户信息
func TestLeadsClaimUpdateWithUserInfoSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsClaimUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.LeadsClaimList = []*model.LeadsClaimItem{
		{
			OuterLeadsId:         "ext_002",
			LeadsUserType:        "USER_TYPE_WX_OPENID",
			LeadsUserWechatAppid: "wx1234567890",
			LeadsUserId:          "openid_abc123",
			CampaignId:           100002,
			AdgroupId:            200002,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsClaimUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 更新线索归因信息-带微信服务商id
func TestLeadsClaimUpdateWithAgencyIdSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsClaimUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.LeadsClaimList = []*model.LeadsClaimItem{
		{
			OuterLeadsId:   "ext_003",
			CampaignId:     100003,
			AdgroupId:      200003,
			WechatAgencyId: "agency_001",
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsClaimUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 更新线索归因信息-批量更新
func TestLeadsClaimUpdateBatchSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsClaimUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.LeadsClaimList = []*model.LeadsClaimItem{
		{
			OuterLeadsId: "ext_001",
			CampaignId:   100001,
			AdgroupId:    200001,
		},
		{
			OuterLeadsId: "ext_002",
			CampaignId:   100002,
			AdgroupId:    200002,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsClaimUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少account_id
func TestLeadsClaimUpdateValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.LeadsClaimUpdateReq{}
	req.AccessToken = "123"
	req.LeadsClaimList = []*model.LeadsClaimItem{
		{OuterLeadsId: "ext_001", CampaignId: 100001},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少leads_claim_list
func TestLeadsClaimUpdateValidateListEmptySelf(t *testing.T) {
	req := &model.LeadsClaimUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：leads_claim_list为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-outer_leads_id为空
func TestLeadsClaimUpdateValidateOuterLeadsIdEmptySelf(t *testing.T) {
	req := &model.LeadsClaimUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.LeadsClaimList = []*model.LeadsClaimItem{
		{CampaignId: 100001, AdgroupId: 200001},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：leads_claim_list.outer_leads_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
