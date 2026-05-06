package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAsyncTaskDownload(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AsyncTaskDownloadReq{}
	req.AccessToken = "your_access_token"
	req.AdvertiserId = 20000800
	req.TaskId = 1230104
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AsyncTaskDownload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("file size: %d bytes\n", len(resp.FileData))
}
