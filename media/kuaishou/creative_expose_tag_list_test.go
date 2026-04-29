package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestCreativeExposeTagList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.CreativeExposeTagListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.CampaignType = 24
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.CreativeExposeTagList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
