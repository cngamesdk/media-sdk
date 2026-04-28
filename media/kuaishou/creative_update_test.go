package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestCreativeUpdate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.CreativeUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20007185
	req.CreativeId = 81810408
	req.CreativeName = "修改搜索创意2"
	req.PhotoId = "49242450755"
	req.ActionBarText = "立即咨询"
	req.Description = "12[地区][节日][地点][男人女人]123[年龄][区县]"
	req.ClickTrackUrl = "https://www.test.com/click?imei=__IMEI__&mac=__MAC__&callback_url=__CALLBACK__"
	req.ActionbarClickUrl = "https://www.test.com/actionbar_click?idfa=__IDFA__"
	req.NewExposeTag = []kuaishouModel.CreativeCreateNewExposeTag{
		{Text: "安享晚年", Url: "https://www.test.com"},
		{Text: "幸福一生", Url: "https://www.test.com"},
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.CreativeUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
