package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestWordInfoUpdateStatus(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.WordInfoUpdateStatusReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 90000344
	req.WordInfoIds = []int64{12345678, 23456789}
	req.PutStatus = 2 // 暂停
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.WordInfoUpdateStatus(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
