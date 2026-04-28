package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAdvancedCreativeUpdate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AdvancedCreativeUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 417173
	req.UnitId = 21456613
	req.PackageName = "修改程序化创意测试"
	req.ActionBar = "立即下载"
	req.Captions = []string{
		"【网红爆款】国妆特证，明星同款，改善头皮环境，控油防脱",
	}
	req.PhotoList = []kuaishouModel.AdvancedCreativePhoto{
		{
			PhotoId:              520390941748682323,
			CreativeMaterialType: 1, // 竖版视频
			CoverImageToken:      "7b27744f637d4e06aa15d2fa830d1bb.jpg",
		},
	}
	req.ClickUrl = "https://www.test.com/click?callback=__CALLBACK__"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AdvancedCreativeUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
