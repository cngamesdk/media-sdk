package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestCreativeUpdateStatus(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.CreativeUpdateStatusReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.CreativeId = 2960188
	req.PutStatus = 2 // 暂停
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.CreativeUpdateStatus(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
