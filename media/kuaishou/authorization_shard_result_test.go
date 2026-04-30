package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAuthorizationShardResult(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AuthorizationShardResultReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12345678
	req.ShardAuthId = 1234567890
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AuthorizationShardResult(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
