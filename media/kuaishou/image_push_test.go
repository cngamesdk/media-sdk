package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestImagePushByTokens(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.ImagePushReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.ImageTokens = []string{"market484ccbeaf5a6467495cbc0248d8e83d7.jpg"}
	req.ShareAdvertiserIds = []int64{11360719}
	req.ShareAccountType = 1
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ImagePush(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestImagePushByPicIds(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.ImagePushReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.PicIds = []string{"5244160298288158171"}
	req.ShareAdvertiserIds = []int64{11360719, 11360720}
	req.ShareAccountType = 1
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ImagePush(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestImagePushMultipleTargets(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.ImagePushReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.ImageTokens = []string{
		"market484ccbeaf5a6467495cbc0248d8e83d7.jpg",
		"market123abc456def789ghi012jkl345mno678.jpg",
	}
	req.ShareAdvertiserIds = []int64{11360719, 11360720, 11360721}
	req.ShareAccountType = 2
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ImagePush(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
