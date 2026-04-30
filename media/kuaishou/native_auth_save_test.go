package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestNativeAuthSave(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.NativeAuthSaveReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12345678
	req.BatchUserIds = []int64{2022036857}
	req.ValidType = "1"
	req.ValidStartTime = 1719744000000
	req.ValidEndTime = 1719820800000
	req.KolUserType = 2
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.NativeAuthSave(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
