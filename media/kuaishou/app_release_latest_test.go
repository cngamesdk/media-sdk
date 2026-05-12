package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppReleaseLatest(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppReleaseLatestReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015366
	req.AppId = 2199123262249
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppReleaseLatest(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("app_id=%d package_id=%d real_app_version=%s version_code=%d\n",
		resp.AppId, resp.PackageId, resp.RealAppVersion, resp.VersionCode)
}

func TestAppReleaseLatestValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppReleaseLatestReq{}
	req.AdvertiserId = 900015366
	req.AppId = 2199123262249
	_, err := adapter.AppReleaseLatest(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppReleaseLatestReq{}
	req2.AccessToken = "your_access_token"
	req2.AppId = 2199123262249
	_, err2 := adapter.AppReleaseLatest(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 app_id，预期返回校验错误
	req3 := &kuaishouModel.AppReleaseLatestReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015366
	_, err3 := adapter.AppReleaseLatest(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty app_id")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
