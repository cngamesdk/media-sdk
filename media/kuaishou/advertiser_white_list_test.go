package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAdvertiserWhiteList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AdvertiserWhiteListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AdvertiserWhiteList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
