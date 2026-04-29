package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAdvCardList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AdvCardListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.CardType = 100 // 图片卡片
	req.Page = 1
	req.PageSize = 10
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AdvCardList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
