package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 新增线索-基础表单类型
func TestLeadsAddSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.LeadsInfoList = []*model.LeadsAddInfo{
		{
			LeadsType: model.LeadsTypeForm,
			LeadsTel:  "13800138000",
			LeadsName: "张三",
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 新增线索-CLICKID匹配模式
func TestLeadsAddWithClickIdMatchSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.MatchType = model.LeadsMatchTypeClickId
	req.LeadsInfoList = []*model.LeadsAddInfo{
		{
			LeadsType:       model.LeadsTypeForm,
			ClickId:         "1234567",
			LeadsTel:        "13800138000",
			LeadsName:       "李四",
			LeadsGender:     model.LeadsGenderTypeMale,
			LeadsActionTime: "2024-01-15 10:30:00",
			OuterLeadsId:    "ext_001",
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 新增线索-带自定义标签
func TestLeadsAddWithCustomizedTagsSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.MatchType = model.LeadsMatchTypeContact
	req.LeadsInfoList = []*model.LeadsAddInfo{
		{
			LeadsType: model.LeadsTypeForm,
			LeadsTel:  "13800138000",
			LeadsName: "王五",
			LeadsArea: "北京",
			Bundle:    `{"key1":"value1","key2":"value2"}`,
			CustomizedTags: []*model.LeadsCustomizedTag{
				{
					TagGroupName: "意向等级",
					TagNameList:  []string{"高意向", "已咨询"},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 新增线索-批量导入
func TestLeadsAddBatchSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.LeadsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.LeadsInfoList = []*model.LeadsAddInfo{
		{
			LeadsType:    model.LeadsTypeForm,
			LeadsTel:     "13800138001",
			LeadsName:    "用户A",
			OuterLeadsId: "ext_001",
		},
		{
			LeadsType:    model.LeadsTypeMakePhoneCall,
			LeadsTel:     "13800138002",
			LeadsName:    "用户B",
			OuterLeadsId: "ext_002",
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.LeadsAddSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少account_id
func TestLeadsAddValidateAccountIdEmptySelf(t *testing.T) {
	req := &model.LeadsAddReq{}
	req.AccessToken = "123"
	req.LeadsInfoList = []*model.LeadsAddInfo{
		{LeadsType: model.LeadsTypeForm, LeadsTel: "13800138000"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少leads_info_list
func TestLeadsAddValidateLeadsInfoListEmptySelf(t *testing.T) {
	req := &model.LeadsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：leads_info_list为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-leads_type为空
func TestLeadsAddValidateLeadsTypeEmptySelf(t *testing.T) {
	req := &model.LeadsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.LeadsInfoList = []*model.LeadsAddInfo{
		{LeadsTel: "13800138000"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：leads_info_list.leads_type为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-leads_type值无效
func TestLeadsAddValidateLeadsTypeInvalidSelf(t *testing.T) {
	req := &model.LeadsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.LeadsInfoList = []*model.LeadsAddInfo{
		{LeadsType: "INVALID_TYPE", LeadsTel: "13800138000"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：leads_info_list.leads_type值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-match_type值无效
func TestLeadsAddValidateMatchTypeInvalidSelf(t *testing.T) {
	req := &model.LeadsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 111111111
	req.MatchType = "INVALID"
	req.LeadsInfoList = []*model.LeadsAddInfo{
		{LeadsType: model.LeadsTypeForm, LeadsTel: "13800138000"},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：match_type值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}
