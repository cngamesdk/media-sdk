package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAdvCardShare(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AdvCardShareReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900312980
	req.ShareAdvertiserIds = []int64{900313084}
	req.AdvCardIds = []int64{18602650}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AdvCardShare(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
