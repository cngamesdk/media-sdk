package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestWordInfoCreate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.WordInfoCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 90000344
	req.CampaignId = 12345678
	req.UnitId = 3833404
	req.WordInfos = []kuaishouModel.WordInfoItem{
		{
			Word:      "测试关键词1",
			MatchType: 1, // 精确匹配
		},
		{
			Word:      "测试关键词2",
			MatchType: 2, // 短语匹配
		},
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.WordInfoCreate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
