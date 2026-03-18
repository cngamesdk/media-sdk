package toutiao

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/toutiao/model"
	"testing"
)

// 升级版工作台上传视频
func TestEbpVideoUploadSelf(t *testing.T) {
	ctx := context.Background()
	factory := NewToutiaoAdapter(config.DefaultConfig())
	req := &model.EbpVideoUploadReq{}
	req.AccessToken = "test"
	req.AccountID = 123
	req.VideoSignature = "123"
	req.UploadType = model.UploadTypeURL
	req.VideoURL = "https://www.xxx.com"
	resp, err := factory.EbpVideoUploadSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("get result: %+v", resp))
}
