package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAsyncTaskCreate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AsyncTaskCreateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.TaskName = "test_async_task"
	req.TaskParams = kuaishouModel.AsyncTaskCreateParams{
		StartDate:           "2022-02-07",
		EndDate:             "2022-02-15",
		ViewType:            1,
		TemporalGranularity: "DAILY",
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AsyncTaskCreate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
