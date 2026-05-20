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

// TestAddAtpFeedSelf 测试新增定向包（基础字段）
func TestAddAtpFeedSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AtpFeedAddReq{
		AtpFeedTypes: []model.AtpFeedType{
			{
				AtpFeedName: "测试定向包",
				AtpFeedDesc: "测试描述",
				Ftypes:      []int{model.FtypeBaiduFeed, model.FtypeTieba},
				Subject:     1,
				Audience:    map[string]string{},
			},
		},
	}
	resp, err := factory.AddAtpFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("add result: %+v", resp))
	if len(resp.Data) > 0 {
		println(fmt.Sprintf("created atp: id=%d, name=%s", resp.Data[0].AtpFeedId, resp.Data[0].AtpFeedName))
	}
}

// TestAddAtpFeedSelfFull 测试新增定向包（完整字段含定向设置）
func TestAddAtpFeedSelfFull(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.AtpFeedAddReq{
		AtpFeedTypes: []model.AtpFeedType{
			{
				AtpFeedName: "完整定向包_测试",
				AtpFeedDesc: "包含所有定向设置的完整定向包",
				Ftypes:      []int{model.FtypeBaiduFeed},
				Subject:     1,
				Audience: map[string]string{
					"age":    "25-44",
					"sex":    "1",
					"region": "110000",
				},
				DeliveryType:    []int{model.DeliveryTypeAll},
				MiniProgramType: 3,
			},
		},
	}
	resp, err := factory.AddAtpFeedSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("add result: %+v", resp))
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
