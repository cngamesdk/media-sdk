package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// CampaignContinuePeriodUpdate 修改周期稳投计划续投状态
func (a *KuaishouAdapter) CampaignContinuePeriodUpdate(ctx context.Context, req *kuaishouModel.CampaignContinuePeriodUpdateReq) (resp *kuaishouModel.CampaignContinuePeriodUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.CampaignContinuePeriodUpdateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/campaign/continue/period/update", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
