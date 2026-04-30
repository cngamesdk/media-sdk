package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// NativeNatureDetailedReport 原生报表披露自然流量后查询原生经营明细
func (a *KuaishouAdapter) NativeNatureDetailedReport(ctx context.Context, req *kuaishouModel.NativeNatureDetailedReportReq) (resp *kuaishouModel.NativeNatureDetailedReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.NativeNatureDetailedReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/effect/native/natureDetailedReport", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
