package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// NativeChartReport 原生效果数据整体概览
func (a *KuaishouAdapter) NativeChartReport(ctx context.Context, req *kuaishouModel.NativeChartReportReq) (resp *kuaishouModel.NativeChartReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.NativeChartReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/effect/native/chartReport", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
