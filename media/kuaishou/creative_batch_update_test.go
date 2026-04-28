package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestCreativeBatchUpdate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.CreativeBatchUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 4171736
	req.UnitId = 22826160
	req.Creatives = []kuaishouModel.CreativeBatchUpdateItem{
		{
			CreativeName:         "测试创意1",
			PhotoId:              "5214605445653959402",
			CreativeMaterialType: 1, // 竖版视频
			ActionBarText:        "立即下载",
			Description:          "支持建站工具的落地页",
		},
		{
			CreativeName:         "测试创意2",
			PhotoId:              "5214605445653959402",
			CreativeMaterialType: 1,
			ActionBarText:        "立即下载",
			Description:          "支持建站工具的落地页",
		},
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.CreativeBatchUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
