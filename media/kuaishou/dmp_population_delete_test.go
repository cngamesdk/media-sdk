package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestDmpPopulationDelete(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.DmpPopulationDeleteReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.OrientationId = 110750443
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpPopulationDelete(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("delete_msg=%s\n", resp.DeleteMsg)
}

func TestDmpPopulationDeleteValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.DmpPopulationDeleteReq{}
	req.AdvertiserId = 20000681
	req.OrientationId = 110750443
	_, err := adapter.DmpPopulationDelete(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.DmpPopulationDeleteReq{}
	req2.AccessToken = "your_access_token"
	req2.OrientationId = 110750443
	_, err2 := adapter.DmpPopulationDelete(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 orientation_id，预期返回校验错误
	req3 := &kuaishouModel.DmpPopulationDeleteReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 20000681
	_, err3 := adapter.DmpPopulationDelete(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty orientation_id")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
