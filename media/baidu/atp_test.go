package baidu

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/baidu/model"
)

// TestGetAtpFeedSelf 测试查询定向包（指定字段）
func TestGetAtpFeedSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AtpFeedReq{
		AtpFeedFields: []string{"atpFeedId", "atpFeedName", "atpFeedDesc", "ftypes", "subject"},
	}
	resp, err := factory.GetAtpFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result count: %d", len(resp.Data)))
	if len(resp.Data) > 0 {
		println(fmt.Sprintf("atp: id=%d, name=%s, desc=%s, subject=%d",
			resp.Data[0].AtpFeedId, resp.Data[0].AtpFeedName,
			resp.Data[0].AtpFeedDesc, resp.Data[0].Subject))
	}
}

// TestGetAtpFeedSelfByIds 测试按ID查询定向包
func TestGetAtpFeedSelfByIds(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AtpFeedReq{
		AtpFeedFields: []string{"atpFeedId", "atpFeedName", "atpFeedDesc"},
		Ids:           []int64{1},
	}
	resp, err := factory.GetAtpFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// TestGetAtpFeedSelfAllFields 测试查询定向包（全部字段含定向设置）
func TestGetAtpFeedSelfAllFields(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AtpFeedReq{
		AtpFeedFields: []string{
			"atpFeedId", "atpFeedName", "atpFeedDesc", "ftypes", "subject",
			"relatedAdgroupFeeds", "audience", "deliveryType", "miniProgramType",
		},
	}
	resp, err := factory.GetAtpFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result count: %d", len(resp.Data)))
	if len(resp.Data) > 0 {
		data := resp.Data[0]
		println(fmt.Sprintf("atp: id=%d, name=%s, desc=%s, subject=%d, ftypes=%v, deliveryType=%v",
			data.AtpFeedId, data.AtpFeedName, data.AtpFeedDesc,
			data.Subject, data.Ftypes, data.DeliveryType))
		println(fmt.Sprintf("relatedUnits count=%d, audience=%+v, miniProgramType=%d",
			len(data.RelatedAdgroupFeeds), data.Audience, data.MiniProgramType))
	}
}

// TestGetAtpFeedSelfByKey 测试按关键字查询定向包
func TestGetAtpFeedSelfByKey(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AtpFeedReq{
		AtpFeedFields: []string{"atpFeedId", "atpFeedName", "atpFeedDesc"},
		Key:           "测试",
		PageNo:        1,
		PageSize:      10,
	}
	resp, err := factory.GetAtpFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result count: %d", len(resp.Data)))
}
