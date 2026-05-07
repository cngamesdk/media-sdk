package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestVideoAutoShareSwitch(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoAutoShareSwitchReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 7869843
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoAutoShareSwitch(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("user_id=%d switch_status=%v\n", resp.UserId, resp.SwitchStatus)
}

func TestVideoAutoShareSwitchValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 advertiser_id，预期返回校验错误
	req := &kuaishouModel.VideoAutoShareSwitchReq{}
	req.AccessToken = "your_access_token"
	_, err := adapter.VideoAutoShareSwitch(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err)
}
