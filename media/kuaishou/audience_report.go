package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// AudienceReport 人群分析报表
func (a *KuaishouAdapter) AudienceReport(ctx context.Context, req *kuaishouModel.AudienceReportReq) (resp *kuaishouModel.AudienceReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.AudienceReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/report/audience_report", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
