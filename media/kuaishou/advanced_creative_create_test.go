package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAdvancedCreativeCreate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AdvancedCreativeCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 417173
	req.UnitId = 21456613
	req.PackageName = "测试程序化创意"
	req.ActionBar = "立即下载"
	req.Captions = []string{
		"【网红爆款】国妆特证，明星同款，改善头皮环境，控油防脱",
		"【国家认证】这款生姜洗发水，明星同款，实力防脱育发，改善头皮",
	}
	req.PhotoList = []kuaishouModel.AdvancedCreativePhoto{
		{
			PhotoId:              520390941748682323,
			CreativeMaterialType: 1, // 竖版视频
			CoverImageToken:      "7b27744f637d4e06aa15d2fa830d1bb.jpg",
		},
	}
	req.CreativeCategory = 701
	req.CreativeTag = []string{"快手", "优惠"}
	req.ClickUrl = "https://ad.e.kuaishou.com/create?__CALLBACK__"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AdvancedCreativeCreate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
