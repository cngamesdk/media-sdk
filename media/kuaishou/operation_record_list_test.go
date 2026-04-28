package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestOperationRecordList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.OperationRecordListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.Page = 1
	req.PageSize = 20
	req.StartDate = "2026-01-01"
	req.EndDate = "2026-04-28"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.OperationRecordList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
