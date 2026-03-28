package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

// 获取广告
func TestAdgroupsGetSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.AdgroupsGetReq{}
	req.AccessToken = "123"
	req.AccountID = 123
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.AdgroupsGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v", result)
}
