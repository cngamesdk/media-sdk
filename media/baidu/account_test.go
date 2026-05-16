package baidu

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/baidu/model"
	"testing"
)

// TestGetAccountFeedSelf 测试查询信息流账户信息（指定返回字段）
func TestGetAccountFeedSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AccountFeedReq{
		AccountFeedFields: []string{"userId", "balance", "budget", "balancePackage", "userStat", "uaStatus", "validFlows", "cid", "liceName"},
	}
	resp, err := factory.GetAccountFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
	if len(resp.Data) > 0 {
		println(fmt.Sprintf("account data[0]: %+v", resp.Data[0]))
	}
}

// TestGetAccountFeedSelfAllFields 测试查询信息流账户信息（请求所有字段）
func TestGetAccountFeedSelfAllFields(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AccountFeedReq{
		AccountFeedFields: []string{
			"userId", "balance", "budget", "balancePackage", "userStat",
			"uaStatus", "validFlows", "cid", "liceName", "tradeId",
			"budgetOfflineTime", "adtype", "rtaUserAdmin",
		},
	}
	resp, err := factory.GetAccountFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestGetAccountFeedSelfEmpty 测试查询信息流账户信息（不指定字段，返回全部）
func TestGetAccountFeedSelfEmpty(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AccountFeedReq{}
	resp, err := factory.GetAccountFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestUpdateAccountFeedSelf 测试更新信息流账户信息
func TestUpdateAccountFeedSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AccountFeedUpdateReq{
		AccountFeedType: model.AccountFeedUpdateType{
			Budget: 3333.33,
		},
	}
	resp, err := factory.UpdateAccountFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
	if len(resp.Data) > 0 {
		println(fmt.Sprintf("update result data[0]: %+v", resp.Data[0]))
	}
}

// TestUpdateAccountFeedSelfBudget 测试更新账户预算边界值
func TestUpdateAccountFeedSelfBudget(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())

	// 测试最小预算值
	req := &model.AccountFeedUpdateReq{
		AccountFeedType: model.AccountFeedUpdateType{
			Budget: 50,
		},
	}
	resp, err := factory.UpdateAccountFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("min budget result: %+v", resp))

	// 测试预算为0（不限预算）
	req2 := &model.AccountFeedUpdateReq{
		AccountFeedType: model.AccountFeedUpdateType{
			Budget: 0,
		},
	}
	resp2, err := factory.UpdateAccountFeedSelf(ctx, "test_user", "test_token", req2)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("zero budget result: %+v", resp2))
}
