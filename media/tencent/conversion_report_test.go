package tencent

import (
	"context"
	"fmt"
	"github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
	"testing"
)

func TestConversionReportSelf(t *testing.T) {
	ctx := context.Background()
	req := &model.ConversionReportReq{}
	req.Callback = "http://tracking.e.qq.com/conv?cb=YWRzX3NlcnZpY2UsMTU4NDUxMDI3OSwyNjg5MzNhMzc5MTM0YzBjMDQ4ZGZjMGQyNGYzMTk0NWYzMzJiOWNi&conv_id=10001"
	var actions []model.ConversionReportInterface
	appUserId := &model.ConversionReportActionUserIdApp{}
	appAction := &model.ConversionReportActionApp{}
	appAction.ActionType = model.ActionTypeActivateApp
	appAction.UserId = appUserId
	actions = append(actions, appAction)

	req.Data = &model.ConversionReportActions{
		Actions: actions,
	}
	adapter := NewTencentAdapter(config.DefaultConfig())
	result, err := adapter.ConversionReportSelf(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %+v\n", result)
}
