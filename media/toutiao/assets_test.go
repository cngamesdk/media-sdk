package toutiao

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/toutiao/model"
	"testing"
)

// 创建项目
func TestEventManagerAssetsCreateSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewToutiaoAdapter(config.DefaultConfig())
	req := &model.EventManagerAssetsCreateReq{}
	req.AccessToken = "test"
	req.AdvertiserID = 123
	req.AssetType = model.AssetTypeApp
	req.AppAsset = &model.AppAsset{
		Name:        "test",
		PackageName: "test",
		DownloadURL: "https://www.xxx.com",
		AppType:     model.AppTypeAndroid,
		AppID:       123456,
		PackageID:   "1232323",
	}
	resp, err := factory.EventManagerAssetsCreateSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}

// 获取账户下资产列表（新）
func TestEventAssetsListSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewToutiaoAdapter(config.DefaultConfig())
	req := &model.EventAssetsListReq{}
	req.AccessToken = "test"
	req.AdvertiserID = 123
	resp, err := factory.EventAssetsListSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}
