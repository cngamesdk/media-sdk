package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestNativePhotoList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.NativePhotoListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12345678
	req.AuthorId = 2022036857
	req.Count = 10
	req.KolUserType = 2
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.NativePhotoList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
