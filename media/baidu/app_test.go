package baidu

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
)

// TestGetAppFeedSelf 测试查询APP信息
func TestGetAppFeedSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	resp, err := factory.GetAppFeedSelf(ctx, "test_user", "test_token")
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
	if len(resp.Data) > 0 {
		println(fmt.Sprintf("app data[0]: %+v", resp.Data[0]))
	}
}
