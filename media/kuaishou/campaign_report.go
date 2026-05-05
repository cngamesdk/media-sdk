package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// CampaignReport 广告计划数据（实时）
func (a *KuaishouAdapter) CampaignReport(ctx context.Context, req *kuaishouModel.CampaignReportReq) (resp *kuaishouModel.CampaignReportResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.CampaignReportResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/v1/report/campaign_report", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
