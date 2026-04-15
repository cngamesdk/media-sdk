package tencent

import (
	"context"
	"github.com/cngamesdk/media-sdk/media/tencent/model"
)

// ConversionReportSelf 转化上报
func (a *TencentAdapter) ConversionReportSelf(ctx context.Context, req *model.ConversionReportReq) (resp *model.ConversionReportResp, err error) {
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := make(model.Headers)
	headers.Json()
	headers["cache-control"] = "no-cache"
	var result model.ConversionReportResp
	if requestErr := a.RequestPostJson(ctx, headers, req.Callback, req.Data, &result); requestErr != nil {
		err = requestErr
		return
	}
	resp = &result
	return
}
