package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAutoProjectQuery(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AutoProjectQueryReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12345678
	req.Page = 1
	req.PageSize = 10
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AutoProjectQuery(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
