package baidu

import (
	"context"
	"fmt"
	"testing"

	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/baidu/model"
)

// TestGetMediaExposureSelf 测试查询百青藤媒体ID曝光量（单个）
func TestGetMediaExposureSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.MediaExposureReq{
		NewMediaids: []int64{100001},
	}
	resp, err := factory.GetMediaExposureSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
	if len(resp.Data) > 0 {
		println(fmt.Sprintf("media: id=%d, name=%s, pv=%d(万)", resp.Data[0].Id, resp.Data[0].Name, resp.Data[0].Pv))
	}
}

// TestGetMediaExposureSelfBatch 测试批量查询百青藤媒体ID曝光量
func TestGetMediaExposureSelfBatch(t *testing.T) {
	ctx := context.Background()
	factory := NewBaiduAdapter(config.DefaultConfig())
	req := &model.MediaExposureReq{
		NewMediaids: []int64{100001, 100002, 100003},
	}
	resp, err := factory.GetMediaExposureSelf(ctx, "test_user", "test_token", req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result count: %d", len(resp.Data)))
	for _, data := range resp.Data {
		println(fmt.Sprintf("media: id=%d, name=%s, pv=%d(万)", data.Id, data.Name, data.Pv))
	}
}
