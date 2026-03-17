package toutiao

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/toutiao/model"
	"testing"
)

// 创建单元
func TestPromotionCreateSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewToutiaoAdapter(config.DefaultConfig())
	req := &model.PromotionCreateReq{}
	req.AccessToken = "test"
	req.AdvertiserId = 123
	req.ProjectId = 123
	req.Name = "test"
	resp, err := factory.PromotionCreateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}
