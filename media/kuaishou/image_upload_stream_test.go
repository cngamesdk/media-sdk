package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"os"
	"testing"
)

func TestImageUploadStream(t *testing.T) {
	ctx := context.Background()
	fileData, err := os.ReadFile("/path/to/your/image.jpg")
	if err != nil {
		t.Fatal(err)
	}
	req := &kuaishouModel.ImageUploadStreamReq{}
	req.Endpoint = "upload.kuaishouzt.com"
	req.UploadToken = "your_upload_token"
	req.File = fileData
	req.FileName = "image.jpg"
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ImageUploadStream(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result=%d\n", resp.Result)
}

func TestImageUploadStreamValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 endpoint，预期返回校验错误
	req := &kuaishouModel.ImageUploadStreamReq{}
	req.UploadToken = "your_upload_token"
	req.File = []byte("mock")
	req.FileName = "image.jpg"
	_, err := adapter.ImageUploadStream(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty endpoint")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 upload_token，预期返回校验错误
	req2 := &kuaishouModel.ImageUploadStreamReq{}
	req2.Endpoint = "upload.kuaishouzt.com"
	req2.File = []byte("mock")
	req2.FileName = "image.jpg"
	_, err2 := adapter.ImageUploadStream(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty upload_token")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// 缺少 file，预期返回校验错误
	req3 := &kuaishouModel.ImageUploadStreamReq{}
	req3.Endpoint = "upload.kuaishouzt.com"
	req3.UploadToken = "your_upload_token"
	req3.FileName = "image.jpg"
	_, err3 := adapter.ImageUploadStream(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for empty file")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// 缺少 file_name，预期返回校验错误
	req4 := &kuaishouModel.ImageUploadStreamReq{}
	req4.Endpoint = "upload.kuaishouzt.com"
	req4.UploadToken = "your_upload_token"
	req4.File = []byte("mock")
	_, err4 := adapter.ImageUploadStream(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for empty file_name")
	}
	fmt.Printf("got expected error: %v\n", err4)
}
