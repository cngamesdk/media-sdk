package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestWordInfoFileQuery(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.WordInfoFileQueryReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 90000344
	req.FileId = 12345678
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.WordInfoFileQuery(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
