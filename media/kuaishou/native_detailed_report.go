package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// NativeDetailedReport 原生效果数据报表明细
func (a *KuaishouAdapter) NativeDetailedReport(ctx context.Context, req *kuaishouModel.NativeDetailedReportReq) (resp *kuaishouModel.NativeDetailedReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.NativeDetailedReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/effect/native/detailedReport", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
