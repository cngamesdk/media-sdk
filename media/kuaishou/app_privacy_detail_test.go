package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAppPrivacyDetail(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AppPrivacyDetailReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 900015364
	req.PrivacyId = 4801146
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AppPrivacyDetail(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("account_id=%d privacy_id=%d url=%s\n",
		resp.AccountId, resp.PrivacyId, resp.Url)
}

func TestAppPrivacyDetailValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 access_token，预期返回校验错误
	req := &kuaishouModel.AppPrivacyDetailReq{}
	req.AdvertiserId = 900015364
	req.PrivacyId = 4801146
	_, err := adapter.AppPrivacyDetail(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty access_token")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 advertiser_id，预期返回校验错误
	req2 := &kuaishouModel.AppPrivacyDetailReq{}
	req2.AccessToken = "your_access_token"
	req2.PrivacyId = 4801146
	_, err2 := adapter.AppPrivacyDetail(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty advertiser_id")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 privacy_id，预期返回校验错误
	req3 := &kuaishouModel.AppPrivacyDetailReq{}
	req3.AccessToken = "your_access_token"
	req3.AdvertiserId = 900015364
	_, err3 := adapter.AppPrivacyDetail(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty privacy_id")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
