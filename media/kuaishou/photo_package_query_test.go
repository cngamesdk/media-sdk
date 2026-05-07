package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestPhotoPackageQuery(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.PhotoPackageQueryReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.Page = 1
	req.PageSize = 5
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.PhotoPackageQuery(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d details=%d\n", resp.TotalCount, len(resp.Details))
	for _, d := range resp.Details {
		fmt.Printf("  photo_package_id=%d name=%s status=%d\n", d.PhotoPackageId, d.Name, d.Status)
	}
}

func TestPhotoPackageQueryById(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.PhotoPackageQueryReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.PhotoPackageId = 11135761637
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.PhotoPackageQuery(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestPhotoPackageQueryByName(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.PhotoPackageQueryReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 16859321
	req.NameLike = "MAPI"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.PhotoPackageQuery(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("total_count=%d\n", resp.TotalCount)
	for _, d := range resp.Details {
		fmt.Printf("  photo_package_id=%d name=%s photo_ids=%v\n", d.PhotoPackageId, d.Name, d.PhotoIds)
	}
}

func TestPhotoPackageQueryValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 advertiser_id，预期返回校验错误
	req := &kuaishouModel.PhotoPackageQueryReq{}
	req.AccessToken = "your_access_token"
	_, err := adapter.PhotoPackageQuery(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err)
}
