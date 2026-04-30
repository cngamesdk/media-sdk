package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestPhotoUpload(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.PhotoUploadReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 12345678
	req.AuthorId = 2022036857
	req.ShieldBackwardSwitch = true
	req.NativePlcSwitch = true
	req.Sync = 1
	req.PhotoCaption = "测试视频描述"
	req.FileName = "test.mp4"
	req.File = []byte("mock video content")
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.PhotoUpload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
