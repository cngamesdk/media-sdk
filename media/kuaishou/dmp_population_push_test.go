package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestDmpPopulationPush(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.DmpPopulationPushReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.OrientationId = 110750443
	req.DestAccountIds = []int64{20000682, 20000683}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpPopulationPush(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("success=%v fail=%v\n", resp.Success, resp.Fail)
}

func TestDmpPopulationPushBatch(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.DmpPopulationPushReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000681
	req.OrientationId = 110839912
	req.DestAccountIds = []int64{
		20000682, 20000683, 20000684, 20000685, 20000686,
		20000687, 20000688, 20000689, 20000690, 20000691,
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.DmpPopulationPush(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("success_count=%d fail_count=%d\n", len(resp.Success), len(resp.Fail))
	if len(resp.Fail) > 0 {
		fmt.Printf("failed_account_ids=%v\n", resp.Fail)
	}
}

func TestDmpPopulationPushValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.DmpPopulationPushReq{}
	req.AdvertiserId = 20000681
	req.OrientationId = 110750443
	req.DestAccountIds = []int64{20000682}
	_, err := adapter.DmpPopulationPush(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.DmpPopulationPushReq{}
	req2.AccessToken = "your_access_token"
	req2.OrientationId = 110750443
	req2.DestAccountIds = []int64{20000682}
	_, err2 := adapter.DmpPopulationPush(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 orientation_id，预期返回校验错误
	req3 := &kuaishouModel.DmpPopulationPushReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 20000681
	req3.DestAccountIds = []int64{20000682}
	_, err3 := adapter.DmpPopulationPush(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty orientation_id")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// 缺少 dest_account_ids，预期返回校验错误
	req4 := &kuaishouModel.DmpPopulationPushReq{}
	req4.AccessToken = "your_access_token"
	req4.AdvertiserId = 20000681
	req4.OrientationId = 110750443
	_, err4 := adapter.DmpPopulationPush(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for empty dest_account_ids")
	}
	fmt.Printf("got expected error: %v\n", err4)
}
