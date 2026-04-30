package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAuthorizationShard(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AuthorizationShardReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12345678
	req.AuthId = 1234567890
	req.ShardAuthorizeScope = 1
	req.ShardAccountId = "90000344,90000345"
	req.ShardUserType = 2
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AuthorizationShard(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
