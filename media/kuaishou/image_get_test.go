package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestImageGet(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.ImageGetReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.ImageToken = "your_image_token"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ImageGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
