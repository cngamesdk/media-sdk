package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAsyncTaskList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AsyncTaskListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.TaskIds = []int64{1230104}
	req.Page = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AsyncTaskList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
