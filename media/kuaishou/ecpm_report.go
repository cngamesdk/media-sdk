package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// EcpmReport 快小游ECPM报表
func (a *KuaishouAdapter) EcpmReport(ctx context.Context, req *kuaishouModel.EcpmReportReq) (resp *kuaishouModel.EcpmReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.EcpmReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/report/ecpm_report", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
