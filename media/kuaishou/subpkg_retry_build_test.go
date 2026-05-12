package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestSubpkgRetryBuild(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.SubpkgRetryBuildReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015073
	req.AppId = 2199123303321
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.SubpkgRetryBuild(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("retry_cnt=%d\n", resp.RetryCnt)
}

func TestSubpkgRetryBuildValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.SubpkgRetryBuildReq{}
	req.AdvertiserId = 900015073
	req.AppId = 2199123303321
	_, err := adapter.SubpkgRetryBuild(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.SubpkgRetryBuildReq{}
	req2.AccessToken = "your_access_token"
	req2.AppId = 2199123303321
	_, err2 := adapter.SubpkgRetryBuild(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 app_id，预期返回校验错误
	req3 := &kuaishouModel.SubpkgRetryBuildReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015073
	_, err3 := adapter.SubpkgRetryBuild(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty app_id")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
