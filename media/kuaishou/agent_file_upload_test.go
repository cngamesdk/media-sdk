package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentFileUploadSelf(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentFileUploadReq{}
	req.AccessToken = "your_access_token"
	req.AgentId = 309
	req.FileName = "test.jpg"
	req.File = []byte("fake_file_content") // 替换为真实文件内容
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentFileUploadSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
