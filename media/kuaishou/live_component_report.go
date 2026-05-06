package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// LiveComponentReport 直播间组件报表
func (a *KuaishouAdapter) LiveComponentReport(ctx context.Context, req *kuaishouModel.LiveComponentReportReq) (resp *kuaishouModel.LiveComponentReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.LiveComponentReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/report/live_component_report", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
