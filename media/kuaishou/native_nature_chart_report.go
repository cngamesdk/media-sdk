package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// NativeNatureChartReport 原生报表披露自然流量的数据整体概览
func (a *KuaishouAdapter) NativeNatureChartReport(ctx context.Context, req *kuaishouModel.NativeNatureChartReportReq) (resp *kuaishouModel.NativeNatureChartReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.NativeNatureChartReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/effect/native/natureChartReport", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
