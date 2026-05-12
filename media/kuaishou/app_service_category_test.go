package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppServiceCategory(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppServiceCategoryReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 139494
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppServiceCategory(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range *resp {
		fmt.Printf("id=%d name=%s children_count=%d\n", item.Id, item.Name, len(item.Children))
		for _, child := range item.Children {
			fmt.Printf("  child: id=%d name=%s\n", child.Id, child.Name)
		}
	}
}

func TestAppServiceCategoryValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppServiceCategoryReq{}
	req.AdvertiserId = 139494
	_, err := adapter.AppServiceCategory(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppServiceCategoryReq{}
	req2.AccessToken = "your_access_token"
	_, err2 := adapter.AppServiceCategory(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)
}
