package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestPhotoPushToProfile(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.PhotoPushToProfileReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12345678
	req.PhotoIds = "5189272782789606404,5189272782789606405"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.PhotoPushToProfile(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
