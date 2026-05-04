package kuaishou

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
	"testing"
)

func TestAgentReport(t *testing.T) {
	ctx := context.Background()
	req := &kuaishouModel.AgentReportReq{}
	req.AccessToken = "your_access_token"
	req.AgentId = 123
	req.StartDate = "2020-07-09"
	req.EndDate = "2020-07-09"
	req.Page = 1
	req.PageSize = 1000
	adapter := NewKuaishouAdapter(config.DefaultConfig())
	resp, err := adapter.AgentReport(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
