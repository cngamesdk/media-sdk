package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestVideoGet(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoGetReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 139494
	req.PhotoIds = []string{"5196591116855324734"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range *resp {
		fmt.Printf("photo_id=%s photo_name=%s new_status=%d\n", d.PhotoId, d.PhotoName, d.NewStatus)
	}
}

func TestVideoGetMultiple(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoGetReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 139494
	req.PhotoIds = []string{"5196591116855324734", "5208131608549181870"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("count=%d\n", len(*resp))
	for _, d := range *resp {
		fmt.Printf("%+v\n", d)
	}
}

func TestVideoGetValuateInfo(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoGetReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 139494
	req.PhotoIds = []string{"5208131608549181870"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoGet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range *resp {
		if d.AdPhotoValuateInfo != nil {
			fmt.Printf("quality=%s running_score=%d is_dup=%v\n",
				d.AdPhotoValuateInfo.QualityLabel,
				d.AdPhotoValuateInfo.RunningScore,
				d.AdPhotoValuateInfo.IsDupPhoto)
		}
	}
}

func TestVideoGetValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 photo_ids，预期返回校验错误
	req := &kuaishouModel.VideoGetReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 139494
	_, err := adapter.VideoGet(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty photo_ids")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.VideoGetReq{}
	req2.AccessToken = "your_access_token"
	req2.PhotoIds = []string{"5196591116855324734"}
	_, err2 := adapter.VideoGet(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)
}
