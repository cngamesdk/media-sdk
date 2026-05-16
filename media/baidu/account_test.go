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
