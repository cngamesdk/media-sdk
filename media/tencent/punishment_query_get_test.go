package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取违规处罚列表-基本查询
func TestPunishmentQueryGet(t *testing.T) {
	ctx := context.Background()
	req := &model.PunishmentQueryGetReq{}
	req.AccessToken = "123"
	req.AccountIDList = []int64{2045867}
	req.PageConf = &model.PunishmentPageConf{
		Page:     1,
		PageSize: 10,
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.PunishmentQueryGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取违规处罚列表-多账户查询
func TestPunishmentQueryGetMultipleAccounts(t *testing.T) {
	ctx := context.Background()
	req := &model.PunishmentQueryGetReq{}
	req.AccessToken = "123"
	req.AccountIDList = []int64{2045867, 2045868, 2045869}
	req.PageConf = &model.PunishmentPageConf{
		Page:     1,
		PageSize: 20,
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.PunishmentQueryGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 获取违规处罚列表-带筛选条件
func TestPunishmentQueryGetWithFilters(t *testing.T) {
	ctx := context.Background()
	req := &model.PunishmentQueryGetReq{}
	req.AccessToken = "123"
	req.AccountIDList = []int64{2045867}
	req.OrderIDList = []string{"ORDER_001"}
	req.CompanyName = "测试公司"
	req.IllegalStartTime = 1700000000000
	req.IllegalEndTime = 1700100000000
	req.ActionTypeList = []int64{1, 2}
	req.LevelList = []int64{1}
	req.IllegalSceneList = []string{"SCENE_001"}
	req.PageConf = &model.PunishmentPageConf{
		Page:     1,
		PageSize: 10,
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.PunishmentQueryGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}

// 验证测试-缺少 account_id_list
func TestPunishmentQueryGetValidateAccountIDListEmpty(t *testing.T) {
	req := &model.PunishmentQueryGetReq{}
	req.AccessToken = "123"
	req.PageConf = &model.PunishmentPageConf{Page: 1, PageSize: 10}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：account_id_list为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-缺少 page_conf
func TestPunishmentQueryGetValidatePageConfEmpty(t *testing.T) {
	req := &model.PunishmentQueryGetReq{}
	req.AccessToken = "123"
	req.AccountIDList = []int64{2045867}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_conf为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-page_size 超出范围
func TestPunishmentQueryGetValidatePageSizeInvalid(t *testing.T) {
	req := &model.PunishmentQueryGetReq{}
	req.AccessToken = "123"
	req.AccountIDList = []int64{2045867}
	req.PageConf = &model.PunishmentPageConf{Page: 1, PageSize: 1001}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：page_conf.page_size必须在1-1000之间")
	}
	fmt.Printf("验证错误: %v\n", err)
}

// 验证测试-Format 默认值
func TestPunishmentQueryGetFormatDefaults(t *testing.T) {
	req := &model.PunishmentQueryGetReq{}
	req.AccessToken = "123"
	req.AccountIDList = []int64{2045867}
	req.PageConf = &model.PunishmentPageConf{}
	req.Format()
	if req.PageConf.Page != 1 {
		t.Fatalf("期望 page 默认值为 1，实际为 %d", req.PageConf.Page)
	}
	if req.PageConf.PageSize != 10 {
		t.Fatalf("期望 page_size 默认值为 10，实际为 %d", req.PageConf.PageSize)
	}
	fmt.Printf("默认值验证通过: page=%d, page_size=%d\n", req.PageConf.Page, req.PageConf.PageSize)
}

// 验证测试-缺少 access_token
func TestPunishmentQueryGetValidateAccessTokenEmpty(t *testing.T) {
	req := &model.PunishmentQueryGetReq{}
	req.AccountIDList = []int64{2045867}
	req.PageConf = &model.PunishmentPageConf{Page: 1, PageSize: 10}
	req.Format()
	err := req.Validate()
	if err == nil {
		t.Fatal("期望返回错误：access_token为必填")
	}
	fmt.Printf("验证错误: %v\n", err)
}
