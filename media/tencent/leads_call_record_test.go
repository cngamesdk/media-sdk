package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取通话结果-通过线索id
func TestLeadsCallRecordGetByLeadsIdSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsCallRecordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123456
	req.LeadsId = 218000154
	req.RequestId = "223255a1-2d02-44d0-8c1b-7217302de746"
	req.ContactId = "xat-conat-id"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsCallRecordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取通话结果-通过外部线索id
func TestLeadsCallRecordGetByOuterLeadsIdSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsCallRecordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123456
	req.OuterLeadsId = "ext_001"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsCallRecordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取通话结果-仅通过contact_id
func TestLeadsCallRecordGetByContactIdSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsCallRecordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123456
	req.LeadsId = 218000154
	req.ContactId = "xat-conat-id"
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsCallRecordGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少account_id
func TestLeadsCallRecordGetValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.LeadsCallRecordGetReq{}
	req.AccessToken = "123"
	req.LeadsId = 218000154
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-leads_id和outer_leads_id都未填
func TestLeadsCallRecordGetValidateLeadsIdMissingSelf(t *testing.T) {
	req := &model.LeadsCallRecordGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123456
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：leads_id和outer_leads_id二选一必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
