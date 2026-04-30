package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestNativeAuthList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.NativeAuthListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12345678
	req.PageInfo = kuaishouModel.AuthListPageInfo{
		CurrentPage: 1,
		PageSize:    10,
	}
	req.KolUserType = []int{2}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.NativeAuthList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
