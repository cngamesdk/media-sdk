package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestImageUploadComplete(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.ImageUploadCompleteReq{}
	req.Endpoint = "upload.kuaishouzt.com"
	req.UploadToken = "your_upload_token"
	req.FragmentCount = 3
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ImageUploadComplete(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result=%d\n", resp.Result)
}

func TestImageUploadCompleteValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 endpoint，预期返回校验错误
	req := &kuaishouModel.ImageUploadCompleteReq{}
	req.UploadToken = "your_upload_token"
	req.FragmentCount = 3
	_, err := adapter.ImageUploadComplete(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty endpoint")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 upload_token，预期返回校验错误
	req2 := &kuaishouModel.ImageUploadCompleteReq{}
	req2.Endpoint = "upload.kuaishouzt.com"
	req2.FragmentCount = 3
	_, err2 := adapter.ImageUploadComplete(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty upload_token")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// fragment_count 为0，预期返回校验错误
	req3 := &kuaishouModel.ImageUploadCompleteReq{}
	req3.Endpoint = "upload.kuaishouzt.com"
	req3.UploadToken = "your_upload_token"
	req3.FragmentCount = 0
	_, err3 := adapter.ImageUploadComplete(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for zero fragment_count")
	}
	fmt.Printf("got expected error: %v\n", err3)
}
