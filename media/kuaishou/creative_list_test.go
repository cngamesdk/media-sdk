package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"testing"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

func TestCreativeList(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.CreativeListReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20007185
	req.CreativeId = 81810408
	req.CreativeName = "修改搜索创意2"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.CreativeList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
