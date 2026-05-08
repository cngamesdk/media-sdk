package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"os"
	"testing"
)

func TestImageUploadFragment(t *testing.T) {
	ctx := context.Background()
	fileData, err := os.ReadFile("/path/to/your/image.jpg")
	if err != nil {
		t.Fatal(err)
	}
	req := &kuaishouModel.ImageUploadFragmentReq{}
	req.Endpoint = "upload.kuaishouzt.com"
	req.UploadToken = "your_upload_token"
	req.FragmentId = 0
	req.Fragment = fileData
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.ImageUploadFragment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result=%d checksum=%s size=%d\n", resp.Result, resp.Checksum, resp.Size)
}

func TestImageUploadFragmentMultiChunk(t *testing.T) {
	ctx := context.Background()
	fileData, err := os.ReadFile("/path/to/your/large_image.jpg")
	if err != nil {
		t.Fatal(err)
	}
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	const chunkSize = 1024 * 1024 // 1MB per chunk
	total := len(fileData)
	for i := 0; i < total; i += chunkSize {
		end := i + chunkSize
		if end > total {
			end = total
		}
		req := &kuaishouModel.ImageUploadFragmentReq{}
		req.Endpoint = "upload.kuaishouzt.com"
		req.UploadToken = "your_upload_token"
		req.FragmentId = i / chunkSize
		req.Fragment = fileData[i:end]
		resp, uploadErr := adapter.ImageUploadFragment(ctx, req)
		if uploadErr != nil {
			t.Fatalf("fragment %d upload failed: %v", req.FragmentId, uploadErr)
		}
		fmt.Printf("fragment_id=%d result=%d checksum=%s size=%d\n", req.FragmentId, resp.Result, resp.Checksum, resp.Size)
	}
}

func TestImageUploadFragmentValidation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 endpoint，预期返回校验错误
	req := &kuaishouModel.ImageUploadFragmentReq{}
	req.UploadToken = "your_upload_token"
	req.FragmentId = 0
	req.Fragment = []byte("mock")
	_, err := adapter.ImageUploadFragment(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty endpoint")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 upload_token，预期返回校验错误
	req2 := &kuaishouModel.ImageUploadFragmentReq{}
	req2.Endpoint = "upload.kuaishouzt.com"
	req2.FragmentId = 0
	req2.Fragment = []byte("mock")
	_, err2 := adapter.ImageUploadFragment(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for empty upload_token")
	}
	fmt.Printf("got expected error: %v\n", err2)

	// fragment_id 为负数，预期返回校验错误
	req3 := &kuaishouModel.ImageUploadFragmentReq{}
	req3.Endpoint = "upload.kuaishouzt.com"
	req3.UploadToken = "your_upload_token"
	req3.FragmentId = -1
	req3.Fragment = []byte("mock")
	_, err3 := adapter.ImageUploadFragment(ctx, req3)
	if err3 == nil {
		t.Fatal("expected validation error for negative fragment_id")
	}
	fmt.Printf("got expected error: %v\n", err3)

	// 缺少 fragment，预期返回校验错误
	req4 := &kuaishouModel.ImageUploadFragmentReq{}
	req4.Endpoint = "upload.kuaishouzt.com"
	req4.UploadToken = "your_upload_token"
	req4.FragmentId = 0
	_, err4 := adapter.ImageUploadFragment(ctx, req4)
	if err4 == nil {
		t.Fatal("expected validation error for empty fragment")
	}
	fmt.Printf("got expected error: %v\n", err4)
}
