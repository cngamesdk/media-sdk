package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestCreativeCreate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.CreativeCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20007185
	req.UnitId = 3834639
	req.CreativeName = "测试自定义创意"
	req.CreativeMaterialType = 1 // 竖版视频
	req.ActionBarText = "免费咨询"
	req.PhotoId = "50611102611"
	req.ImageToken = "BMjAyMTExMTExMDM4NDVfMTY4NzU1NzE4NF82MDY0NjgwNjE4NV8xXzM=_B8de7da4c5428428377d184f6a10fb643.jpg"
	req.Description = "测试广告语"
	req.ClickTrackUrl = "https://www.test.com/click?callback=__CALLBACK__"
	req.NewExposeTag = []kuaishouModel.CreativeCreateNewExposeTag{
		{Text: "安享晚年", Url: "https://www.test.com"},
		{Text: "安全成长", Url: "https://www.test.com"},
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.CreativeCreate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
