package toutiao

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/toutiao/model"
	"testing"
)

// 创建项目
func TestProjectCreateSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewToutiaoAdapter(config.DefaultConfig())
	req := &model.ProjectCreateReq{}
	req.AccessToken = "test"
	req.Name = "test"
	resp, err := factory.ProjectCreateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// 更新项目
func TestProjectUpdateSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewToutiaoAdapter(config.DefaultConfig())
	req := &model.ProjectUpdateReq{}
	req.AccessToken = "test"
	req.Name = "test"
	req.AdvertiserId = 123
	req.ProjectId = 123
	resp, err := factory.ProjectUpdateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}
