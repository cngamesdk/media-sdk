package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppDetail(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppDetailReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12078245
	req.PackageId = 1099611636558
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppDetail(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("app_id=%d app_name=%s platform=%s status=%d release_type=%d\n",
		resp.AppId, resp.RealAppName, resp.Platform, resp.AppStatus, resp.ReleaseType)
	fmt.Printf("package_id=%d package_name=%s version=%s version_code=%d\n",
		resp.PackageId, resp.PackageName, resp.RealAppVersion, resp.VersionCode)
	fmt.Printf("developer=%s location=%s category=%d use_sdk=%d\n",
		resp.Developer, resp.Location, resp.Category, resp.UseSdk)
	fmt.Printf("source_type=%d share_type=%d share_account_count=%d trace_activation=%d\n",
		resp.SourceType, resp.ShareType, resp.ShareAccountCount, resp.TraceActivation)
	fmt.Printf("ordinary_permissions=%d sensitive_permissions=%d\n",
		len(resp.Permissions.OrdinaryPermissions), len(resp.Permissions.SensitivePermissions))
}

func TestAppDetailValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppDetailReq{}
	req.AdvertiserId = 12078245
	req.PackageId = 1099611636558
	_, err := adapter.AppDetail(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppDetailReq{}
	req2.AccessToken = "your_access_token"
	req2.PackageId = 1099611636558
	_, err2 := adapter.AppDetail(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 package_id，预期返回校验错误
	req3 := &kuaishouModel.AppDetailReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 12078245
	_, err3 := adapter.AppDetail(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty package_id")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
