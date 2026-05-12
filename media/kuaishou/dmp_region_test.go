package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestDmpRegion(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.DmpRegionReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015366
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpRegion(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range *resp {
		fmt.Printf("id=%d name=%s children=%v\n", item.Id, item.Name, item.Children)
	}
}

func TestDmpRegionValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.DmpRegionReq{}
	req.AdvertiserId = 900015366
	_, err := adapter.DmpRegion(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.DmpRegionReq{}
	req2.AccessToken = "your_access_token"
	_, err2 := adapter.DmpRegion(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)
}
