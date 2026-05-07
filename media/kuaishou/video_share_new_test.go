package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestVideoShareNew(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoShareNewReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.PhotoIds = []int64{5228679230762349823}
	req.AccountIds = []int64{6882090}
	req.ShareAccountType = 1
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoShareNew(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("share_status=%d success=%d retry=%d duplicated=%d\n",
		resp.ShareStatus, len(resp.ShareSuccessList), len(resp.NeedToRetryList), len(resp.DuplicatedPhotoList))
}

func TestVideoShareNewWithSyncProfile(t *testing.T) {
	ctx := context.Background()
	syncProfile := true
	req := &kuaishouModel.VideoShareNewReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.PhotoIds = []int64{5228679230762349823, 5251760176597172988}
	req.AccountIds = []int64{6882090, 6882091}
	req.ShareAccountType = 1
	req.SyncProfile = &syncProfile
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoShareNew(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
	for _, item := range resp.ShareSuccessList {
		fmt.Printf("success: account_id=%d photo_id=%s\n", item.AccountId, item.PhotoId)
	}
	for _, item := range resp.NeedToRetryList {
		fmt.Printf("retry: account_id=%d photo_id=%s result=%s\n", item.AccountId, item.PhotoId, item.ShareResult)
	}
}

func TestVideoShareNewValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 photo_ids，预期返回校验错误
	req := &kuaishouModel.VideoShareNewReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.AccountIds = []int64{6882090}
	req.ShareAccountType = 1
	_, err := adapter.VideoShareNew(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty photo_ids")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 account_ids，预期返回校验错误
	req2 := &kuaishouModel.VideoShareNewReq{}
	req2.AccessToken = "your_access_token"
	req2.AdvertiserId = 11311124
	req2.PhotoIds = []int64{5228679230762349823}
	req2.ShareAccountType = 1
	_, err2 := adapter.VideoShareNew(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty account_ids")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 share_account_type，预期返回校验错误
	req3 := &kuaishouModel.VideoShareNewReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 11311124
	req3.PhotoIds = []int64{5228679230762349823}
	req3.AccountIds = []int64{6882090}
	_, err3 := adapter.VideoShareNew(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty share_account_type")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
