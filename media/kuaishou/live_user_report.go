package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// LiveUserReport 直播间报表
func (a *KuaishouAdapter) LiveUserReport(ctx context.Context, req *kuaishouModel.LiveUserReportReq) (resp *kuaishouModel.LiveUserReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.LiveUserReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/v1/report/live_user_report", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
