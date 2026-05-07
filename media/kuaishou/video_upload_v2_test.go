package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"os"
	"testing"
)

func TestVideoUploadV2ByFile(t *testing.T) {
	ctx := context.Background()
	fileData, err := os.ReadFile("/path/to/your/video.mp4")
	if err != nil {
		t.Fatal(err)
	}
	req := &kuaishouModel.VideoUploadV2Req{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.Signature = "36ffe2e21cb5b12bde752021331e2614"
	req.File = fileData
	req.FileName = "video.mp4"
	req.PhotoName = "测试视频"
	req.Type = 1
	req.Sync = 0
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoUploadV2(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("photo_id=%s signature=%s\n", resp.PhotoId, resp.Signature)
}

func TestVideoUploadV2ByFileWithOptions(t *testing.T) {
	ctx := context.Background()
	fileData, err := os.ReadFile("/path/to/your/video.mp4")
	if err != nil {
		t.Fatal(err)
	}
	shieldSwitch := false
	nativePlc := true
	req := &kuaishouModel.VideoUploadV2Req{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.Signature = "36ffe2e21cb5b12bde752021331e2614"
	req.File = fileData
	req.FileName = "video.mp4"
	req.PhotoName = "测试视频全参数"
	req.PhotoTag = "tag1"
	req.Type = 1
	req.Sync = 1
	req.ShieldBackwardSwitch = &shieldSwitch
	req.NativePlcSwitch = &nativePlc
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoUploadV2(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestVideoUploadV2ByBlobStoreKey(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.VideoUploadV2Req{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.Signature = "36ffe2e21cb5b12bde752021331e2614"
	req.BlobStoreKey = "your_blob_store_key"
	req.PhotoName = "流式上传测试"
	req.Type = 1
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.VideoUploadV2(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestVideoUploadV2Validation(t *testing.T) {
	ctx := context.Background()
	adapter := NewKuaishouAdapter(config.DefaultConfig())

	// 缺少 signature，预期返回校验错误
	req := &kuaishouModel.VideoUploadV2Req{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 11311124
	req.File = []byte("fake")
	req.FileName = "video.mp4"
	_, err := adapter.VideoUploadV2(ctx, req)
	if err == nil {
		t.Fatal("expected validation error for empty signature")
	}
	fmt.Printf("got expected error: %v\n", err)

	// 缺少 file 和 blob_store_key，预期返回校验错误
	req2 := &kuaishouModel.VideoUploadV2Req{}
	req2.AccessToken = "your_access_token"
	req2.AdvertiserId = 11311124
	req2.Signature = "36ffe2e21cb5b12bde752021331e2614"
	_, err2 := adapter.VideoUploadV2(ctx, req2)
	if err2 == nil {
		t.Fatal("expected validation error for missing file and blob_store_key")
	}
	fmt.Printf("got expected error: %v\n", err2)
}
