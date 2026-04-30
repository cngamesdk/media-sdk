package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAuthorizationShareAccount(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AuthorizationShareAccountReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12345678
	req.KolUserId = 2022036857
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AuthorizationShareAccount(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
