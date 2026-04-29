package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestWordInfoList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.WordInfoListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 90000344
	req.UnitId = 3833404
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.WordInfoList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
