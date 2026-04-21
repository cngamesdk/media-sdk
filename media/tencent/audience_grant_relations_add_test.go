package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 添加人群授权-授权给指定账号（含 TARGET + INSIGHT 权限）
func TestAudienceGrantRelationsAddToAccount(t *testing.T) {
	ctx := context.Background()
	req := &model.AudienceGrantRelationsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceIDList = []int64{12, 23}
	req.GrantType = model.AudienceGrantTypeGrantTypeBusiness
	req.GrantSpec = &model.GrantSpec{
		GrantToBusinessSpec: &model.GrantToBusinessSpec{
			GrantBusinessID: 12312,
			GrantScopeType:  model.AudienceGrantScopeTypeAccount,
			GrantAccountPermission: []*model.GrantAccountPermission{
				{
					AccountID:               20458,
					GrantPermissionTypeList: []string{model.AudienceGrantPermissionTypeTarget, model.AudienceGrantPermissionTypeInsight},
				},
				{
					AccountID:               20345,
					GrantPermissionTypeList: []string{model.AudienceGrantPermissionTypeTarget, model.AudienceGrantPermissionTypeInsight},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AudienceGrantRelationsAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 添加人群授权-仅授权 TARGET 权限
func TestAudienceGrantRelationsAddTargetOnly(t *testing.T) {
	ctx := context.Background()
	req := &model.AudienceGrantRelationsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceIDList = []int64{100001}
	req.GrantType = model.AudienceGrantTypeGrantTypeBusiness
	req.GrantSpec = &model.GrantSpec{
		GrantToBusinessSpec: &model.GrantToBusinessSpec{
			GrantBusinessID: 12312,
			GrantScopeType:  model.AudienceGrantScopeTypeAccount,
			GrantAccountPermission: []*model.GrantAccountPermission{
				{
					AccountID:               20458,
					GrantPermissionTypeList: []string{model.AudienceGrantPermissionTypeTarget},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AudienceGrantRelationsAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 添加人群授权-批量人群授权（20个）
func TestAudienceGrantRelationsAddBatch(t *testing.T) {
	ctx := context.Background()
	req := &model.AudienceGrantRelationsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	ids := make([]int64, 20)
	for i := range ids {
		ids[i] = int64(100001 + i)
	}
	req.AudienceIDList = ids
	req.GrantType = model.AudienceGrantTypeGrantTypeBusiness
	req.GrantSpec = &model.GrantSpec{
		GrantToBusinessSpec: &model.GrantToBusinessSpec{
			GrantBusinessID: 12312,
			GrantScopeType:  model.AudienceGrantScopeTypeAccount,
			GrantAccountPermission: []*model.GrantAccountPermission{
				{
					AccountID:               20458,
					GrantPermissionTypeList: []string{model.AudienceGrantPermissionTypeTarget, model.AudienceGrantPermissionTypeInsight},
				},
			},
		},
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AudienceGrantRelationsAdd(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id
func TestAudienceGrantRelationsAddValidateAccountIDEmpty(t *testing.T) {
	req := &model.AudienceGrantRelationsAddReq{}
	req.AccessToken = "123"
	req.AudienceIDList = []int64{12}
	req.GrantType = model.AudienceGrantTypeGrantTypeBusiness
	req.GrantSpec = &model.GrantSpec{
		GrantToBusinessSpec: &model.GrantToBusinessSpec{
			GrantBusinessID: 12312,
			GrantScopeType:  model.AudienceGrantScopeTypeAccount,
			GrantAccountPermission: []*model.GrantAccountPermission{
				{AccountID: 20458, GrantPermissionTypeList: []string{model.AudienceGrantPermissionTypeTarget}},
			},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 audience_id_list
func TestAudienceGrantRelationsAddValidateAudienceIDListEmpty(t *testing.T) {
	req := &model.AudienceGrantRelationsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.GrantType = model.AudienceGrantTypeGrantTypeBusiness
	req.GrantSpec = &model.GrantSpec{
		GrantToBusinessSpec: &model.GrantToBusinessSpec{
			GrantBusinessID: 12312,
			GrantScopeType:  model.AudienceGrantScopeTypeAccount,
			GrantAccountPermission: []*model.GrantAccountPermission{
				{AccountID: 20458, GrantPermissionTypeList: []string{model.AudienceGrantPermissionTypeTarget}},
			},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：audience_id_list为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-audience_id_list 超过最大长度
func TestAudienceGrantRelationsAddValidateAudienceIDListTooLong(t *testing.T) {
	req := &model.AudienceGrantRelationsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	ids := make([]int64, 21)
	for i := range ids {
		ids[i] = int64(i + 1)
	}
	req.AudienceIDList = ids
	req.GrantType = model.AudienceGrantTypeGrantTypeBusiness
	req.GrantSpec = &model.GrantSpec{
		GrantToBusinessSpec: &model.GrantToBusinessSpec{
			GrantBusinessID: 12312,
			GrantScopeType:  model.AudienceGrantScopeTypeAccount,
			GrantAccountPermission: []*model.GrantAccountPermission{
				{AccountID: 20458, GrantPermissionTypeList: []string{model.AudienceGrantPermissionTypeTarget}},
			},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：audience_id_list最大长度为20")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-grant_type 值无效
func TestAudienceGrantRelationsAddValidateGrantTypeInvalid(t *testing.T) {
	req := &model.AudienceGrantRelationsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceIDList = []int64{12}
	req.GrantType = "INVALID_TYPE"
	req.GrantSpec = &model.GrantSpec{
		GrantToBusinessSpec: &model.GrantToBusinessSpec{
			GrantBusinessID: 12312,
			GrantScopeType:  model.AudienceGrantScopeTypeAccount,
			GrantAccountPermission: []*model.GrantAccountPermission{
				{AccountID: 20458, GrantPermissionTypeList: []string{model.AudienceGrantPermissionTypeTarget}},
			},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：grant_type值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 grant_spec
func TestAudienceGrantRelationsAddValidateGrantSpecNil(t *testing.T) {
	req := &model.AudienceGrantRelationsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceIDList = []int64{12}
	req.GrantType = model.AudienceGrantTypeGrantTypeBusiness
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：grant_spec为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-grant_scope_type=GRANT_SCOPE_TYPE_ACCOUNT 时缺少 grant_account_permission
func TestAudienceGrantRelationsAddValidateGrantAccountPermissionEmpty(t *testing.T) {
	req := &model.AudienceGrantRelationsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceIDList = []int64{12}
	req.GrantType = model.AudienceGrantTypeGrantTypeBusiness
	req.GrantSpec = &model.GrantSpec{
		GrantToBusinessSpec: &model.GrantToBusinessSpec{
			GrantBusinessID: 12312,
			GrantScopeType:  model.AudienceGrantScopeTypeAccount,
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：grant_account_permission为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-grant_account_permission 中 account_id 为空
func TestAudienceGrantRelationsAddValidateGrantAccountPermissionAccountIDEmpty(t *testing.T) {
	req := &model.AudienceGrantRelationsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceIDList = []int64{12}
	req.GrantType = model.AudienceGrantTypeGrantTypeBusiness
	req.GrantSpec = &model.GrantSpec{
		GrantToBusinessSpec: &model.GrantToBusinessSpec{
			GrantBusinessID: 12312,
			GrantScopeType:  model.AudienceGrantScopeTypeAccount,
			GrantAccountPermission: []*model.GrantAccountPermission{
				{GrantPermissionTypeList: []string{model.AudienceGrantPermissionTypeTarget}},
			},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：grant_account_permission.account_id为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-grant_permission_type_list 值无效
func TestAudienceGrantRelationsAddValidatePermissionTypeInvalid(t *testing.T) {
	req := &model.AudienceGrantRelationsAddReq{}
	req.AccessToken = "123"
	req.AccountID = 2045867
	req.AudienceIDList = []int64{12}
	req.GrantType = model.AudienceGrantTypeGrantTypeBusiness
	req.GrantSpec = &model.GrantSpec{
		GrantToBusinessSpec: &model.GrantToBusinessSpec{
			GrantBusinessID: 12312,
			GrantScopeType:  model.AudienceGrantScopeTypeAccount,
			GrantAccountPermission: []*model.GrantAccountPermission{
				{AccountID: 20458, GrantPermissionTypeList: []string{"INVALID_PERMISSION"}},
			},
		},
	}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：grant_permission_type_list值无效")
	}
	fmt.Printf("验证错误: %v\n", err)
}
