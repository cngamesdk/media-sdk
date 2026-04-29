package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestCreativeRecommendDescription(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.CreativeRecommendDescriptionReq{}
	req.AccessToken = "your_access_token"
	req.AppId = 123456
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.CreativeRecommendDescription(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
