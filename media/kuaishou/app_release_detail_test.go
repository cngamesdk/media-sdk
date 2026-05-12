package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppReleaseDetail(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppReleaseDetailReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.PackageId = 100009087
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppReleaseDetail(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("app_id=%d package_id=%d real_app_name=%s real_app_version=%s platform=%s review_status=%d\n",
		resp.AppId, resp.PackageId, resp.RealAppName, resp.RealAppVersion, resp.Platform, resp.ReviewStatus)
	fmt.Printf("ordinary_permissions=%d sensitive_permissions=%d permission_info_count=%d\n",
		len(resp.Permissions.OrdinaryPermissions), len(resp.Permissions.SensitivePermissions), len(resp.PermissionInformation))
}

func TestAppReleaseDetailValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppReleaseDetailReq{}
	req.AdvertiserId = 900015364
	req.PackageId = 100009087
	_, err := adapter.AppReleaseDetail(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppReleaseDetailReq{}
	req2.AccessToken = "your_access_token"
	req2.PackageId = 100009087
	_, err2 := adapter.AppReleaseDetail(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 package_id，预期返回校验错误
	req3 := &kuaishouModel.AppReleaseDetailReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015364
	_, err3 := adapter.AppReleaseDetail(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty package_id")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
