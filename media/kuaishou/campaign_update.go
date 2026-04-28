package kuaishou

import (
	"context"

	kuaishouModel "github.com/cngamesdk/media-sdk/media/kuaishou/model"
)

// CampaignUpdate 修改广告计划
func (a *KuaishouAdapter) CampaignUpdate(ctx context.Context, req *kuaishouModel.CampaignUpdateReq) (resp *kuaishouModel.CampaignUpdateResp, err error) {
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		err = validateErr
		return
	}
	headers := req.GetHeaders()
	var result kuaishouModel.CampaignUpdateResp
	if errRequest := a.RequestPostJson(ctx, headers, kuaishouModel.AdUrl+"/rest/openapi/gw/dsp/campaign/update", req, &result); errRequest != nil {
		err = errRequest
		return
	}
	resp = &result
	return
}
