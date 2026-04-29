package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestCreativePreview(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.CreativePreviewReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12312312
	req.UnitType = 4 // 自定义创意
	req.CreativeId = 12312321
	req.UserIds = []string{"7475171514"}
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.CreativePreview(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
