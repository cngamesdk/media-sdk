package toutiao

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/toutiao/model"
	"testing"
)

// 获取自定义报表可用指标和维度
func TestReportCustomConfigGetSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewToutiaoAdapter(config.DefaultConfig())
	req := &model.ReportCustomConfigGetReq{}
	req.AccessToken = "test"
	req.AdvertiserID = 123
	resp, err := factory.ReportCustomConfigGetSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}
