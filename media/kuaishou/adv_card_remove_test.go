package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAdvCardRemove(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AdvCardRemoveReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 90000446
	req.AdvCardId = 123456789
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AdvCardRemove(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
