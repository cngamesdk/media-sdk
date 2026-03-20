package toutiao

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/toutiao/model"
	"testing"
)

func TestClickMaros(t *testing.T) {
	macros := model.ClickMacros
	println(macros.BuildQueryString())

	macros.Reset(model.ProjectID, "first_level_id")
	println(macros.BuildQueryString())

	macros.Add("ext1", "EXT1")
	println(macros.BuildQueryString())

	macros.Add("ext2", "EXT2")
	println(macros.BuildUrl("https://www.xxx.com/?game_id=123&agent_id=123&site_id=123"))
}

// TestConversionEventReport 转化事件上报
func TestConversionEventReport(t *testing.T) {
	ctx := context.Background()
	factory := NewToutiaoAdapter(config.DefaultConfig())
	req := &model.ConversionEventReq{}
	req.CallbackUrl = "http://ad.toutiao.com/track/activate/?callback=B.rqB7I6enkYhGUseoWmFQtyWfsFfsrOmxU"
	req.OAID = "123"
	resp, err := factory.DataConversionEventReportSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}
