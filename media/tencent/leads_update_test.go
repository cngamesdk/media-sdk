package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 更新线索基本信息-通过外部线索id匹配
func TestLeadsUpdateSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.LeadsContactList = []*model.LeadsContactItem{
		{
			OuterLeadsId: "ext_001",
			LeadsTel:     "13800138000",
			LeadsName:    "张三",
			LeadsGender:  model.LeadsGenderTypeMale,
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 更新线索基本信息-通过线索id匹配
func TestLeadsUpdateByLeadsIdSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.MatchType = model.LeadsUpdateMatchTypeLeadsId
	req.LeadsContactList = []*model.LeadsContactItem{
		{
			LeadsId:    123456,
			LeadsTel:   "13800138001",
			LeadsName:  "李四",
			LeadsEmail: "lisi@example.com",
			LeadsArea:  "广东省深圳市",
			ShopName:   "深圳旗舰店",
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 更新线索基本信息-通过点击id匹配
func TestLeadsUpdateByClickIdSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.MatchType = model.LeadsUpdateMatchTypeClickId
	req.LeadsContactList = []*model.LeadsContactItem{
		{
			ClickId:     "click_123",
			LeadsWechat: "wx_user_001",
			LeadsName:   "王五",
			Bundle:      `{"interest":"游戏","level":"VIP"}`,
			Memo:        "高意向客户",
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 更新线索基本信息-批量更新
func TestLeadsUpdateBatchSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.LeadsContactList = []*model.LeadsContactItem{
		{
			OuterLeadsId: "ext_001",
			LeadsTel:     "13800138000",
			LeadsName:    "用户A",
		},
		{
			OuterLeadsId: "ext_002",
			LeadsQq:      123456789,
			LeadsName:    "用户B",
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少account_id
func TestLeadsUpdateValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.LeadsUpdateReq{}
	req.AccessToken = "123"
	req.LeadsContactList = []*model.LeadsContactItem{
		{OuterLeadsId: "ext_001", LeadsTel: "13800138000"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少leads_contact_list
func TestLeadsUpdateValidateListEmptySelf(t *testing.T) {
	req := &model.LeadsUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：leads_contact_list为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-match_type值无效
func TestLeadsUpdateValidateMatchTypeInvalidSelf(t *testing.T) {
	req := &model.LeadsUpdateReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.MatchType = "CONTACT"
	req.LeadsContactList = []*model.LeadsContactItem{
		{OuterLeadsId: "ext_001"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：match_type值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}
