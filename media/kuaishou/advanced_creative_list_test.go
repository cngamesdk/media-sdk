package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAdvancedCreativeList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AdvancedCreativeListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 417173
	req.UnitIds = []string{"21456613"}
	req.Page = 1
	req.PageSize = 20
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AdvancedCreativeList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
