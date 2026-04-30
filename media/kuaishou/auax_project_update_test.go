package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAuaxProjectUpdate(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AuaxProjectUpdateReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12345678
	req.AuaxProjectId = 1234567890
	req.OcpxActionType = 191
	req.RoiRatio = 1.5
	req.PhotoPackageInfo = []int64{1234567890}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AuaxProjectUpdate(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
