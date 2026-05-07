package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestImageList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.ImageListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.Page = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ImageList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestImageListWithFilter(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.ImageListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.StartDate = "2023-11-01"
	req.EndDate = "2023-11-18"
	req.PicTypes = []int{5, 6}
	req.Page = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ImageList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestImageListByTokens(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.ImageListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.ImageTokens = []string{"token1", "token2"}
	req.Page = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ImageList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
