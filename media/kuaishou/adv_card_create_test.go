package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAdvCardCreate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AdvCardCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 90000446
	req.AdvList = []kuaishouModel.AdvCardCreateItem{
		{
			CardType: 100, // 图片卡片
			Url:      "http://static.yximgs.com/udata/pkg/536fd750628a4cfb97ca82758e6a7ecd.png",
			Width:    720,
			Height:   408,
		},
		{
			CardType:    104, // 快捷评论卡
			ContentType: 2,   // emoji快捷评论卡
			Title:       "11221",
			EmojiList: []kuaishouModel.AdvCardEmojiItem{
				{
					EmojiCode: "[加油]",
					EmojiUrl:  "https://js2.a.yximgs.com/bs2/emotion/app_1576120138988_5x756gefzbda399.png",
				},
				{
					EmojiCode: "[期待]",
					EmojiUrl:  "https://js2.a.yximgs.com/bs2/emotion/app_1576120138988_5x78bpjjb5d63ps.png",
				},
				{
					EmojiCode: "[红包]",
					EmojiUrl:  "https://ali2.a.yximgs.com/bs2/emotion/app_1576120138988_5x2vtpwxnm7ikbc.png",
				},
			},
		},
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AdvCardCreate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
