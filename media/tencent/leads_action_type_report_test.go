package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 线索上报DMP平台-通过外部线索id
func TestLeadsActionTypeReportAddSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsActionTypeReportAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.LeadsActionTypeReportList = []*model.LeadsActionTypeReportItem{
		{
			OuterLeadsId: "ext_001",
			ActionType:   "FORM_SUBMIT",
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsActionTypeReportAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 线索上报DMP平台-电话接通带通话时长
func TestLeadsActionTypeReportAddPhoneConnectedSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsActionTypeReportAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.LeadsActionTypeReportList = []*model.LeadsActionTypeReportItem{
		{
			OuterLeadsId: "ext_002",
			ActionType:   model.ActionTypePhoneConnected,
			CallDuration: 120,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsActionTypeReportAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 线索上报DMP平台-通过线索id匹配
func TestLeadsActionTypeReportAddByLeadsIdSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsActionTypeReportAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.MatchType = model.LeadsActionReportMatchTypeLeadsId
	req.LeadsActionTypeReportList = []*model.LeadsActionTypeReportItem{
		{
			LeadsId:    123456,
			ActionType: "CONSULTATION",
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsActionTypeReportAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 线索上报DMP平台-通过联系方式匹配
func TestLeadsActionTypeReportAddByContactSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsActionTypeReportAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.MatchType = model.LeadsActionReportMatchTypeContact
	req.LeadsActionTypeReportList = []*model.LeadsActionTypeReportItem{
		{
			LeadsTel:     "13800138000",
			ActionType:   model.ActionTypePhoneConnected,
			CallDuration: 60,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsActionTypeReportAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 线索上报DMP平台-批量上报
func TestLeadsActionTypeReportAddBatchSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsActionTypeReportAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.LeadsActionTypeReportList = []*model.LeadsActionTypeReportItem{
		{
			OuterLeadsId: "ext_001",
			ActionType:   "FORM_SUBMIT",
		},
		{
			OuterLeadsId: "ext_002",
			ActionType:   model.ActionTypePhoneConnected,
			CallDuration: 180,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsActionTypeReportAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少account_id
func TestLeadsActionTypeReportAddValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.LeadsActionTypeReportAddReq{}
	req.AccessToken = "123"
	req.LeadsActionTypeReportList = []*model.LeadsActionTypeReportItem{
		{OuterLeadsId: "ext_001", ActionType: "FORM_SUBMIT"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少leads_action_type_report_list
func TestLeadsActionTypeReportAddValidateListEmptySelf(t *testing.T) {
	req := &model.LeadsActionTypeReportAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：leads_action_type_report_list为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-action_type为空
func TestLeadsActionTypeReportAddValidateActionTypeEmptySelf(t *testing.T) {
	req := &model.LeadsActionTypeReportAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.LeadsActionTypeReportList = []*model.LeadsActionTypeReportItem{
		{OuterLeadsId: "ext_001"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：leads_action_type_report_list.action_type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-PHONE_CONNECTED时缺少call_duration
func TestLeadsActionTypeReportAddValidateCallDurationMissingSelf(t *testing.T) {
	req := &model.LeadsActionTypeReportAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.LeadsActionTypeReportList = []*model.LeadsActionTypeReportItem{
		{OuterLeadsId: "ext_001", ActionType: model.ActionTypePhoneConnected},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：action_type为PHONE_CONNECTED时call_duration必填且大于0")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-match_type值无效
func TestLeadsActionTypeReportAddValidateMatchTypeInvalidSelf(t *testing.T) {
	req := &model.LeadsActionTypeReportAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.MatchType = "INVALID"
	req.LeadsActionTypeReportList = []*model.LeadsActionTypeReportItem{
		{OuterLeadsId: "ext_001", ActionType: "FORM_SUBMIT"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：match_type值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}
